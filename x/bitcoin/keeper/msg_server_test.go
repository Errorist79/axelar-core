package keeper_test

import (
	"bytes"
	"crypto/ecdsa"
	cryptoRand "crypto/rand"
	"fmt"
	mathRand "math/rand"
	"testing"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogoprototypes "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/axelarnetwork/axelar-core/testutils"
	"github.com/axelarnetwork/axelar-core/testutils/rand"
	"github.com/axelarnetwork/axelar-core/utils"
	utilsmock "github.com/axelarnetwork/axelar-core/utils/mock"
	"github.com/axelarnetwork/axelar-core/x/bitcoin/exported"
	bitcoinKeeper "github.com/axelarnetwork/axelar-core/x/bitcoin/keeper"
	"github.com/axelarnetwork/axelar-core/x/bitcoin/types"
	"github.com/axelarnetwork/axelar-core/x/bitcoin/types/mock"
	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	snapshot "github.com/axelarnetwork/axelar-core/x/snapshot/exported"
	tss "github.com/axelarnetwork/axelar-core/x/tss/exported"
	vote "github.com/axelarnetwork/axelar-core/x/vote/exported"
	voteMock "github.com/axelarnetwork/axelar-core/x/vote/exported/mock"
)

func TestHandleMsgLink(t *testing.T) {
	var (
		server      types.MsgServiceServer
		btcKeeper   *mock.BTCKeeperMock
		signer      *mock.SignerMock
		nexusKeeper *mock.NexusMock
		ctx         sdk.Context
		msg         *types.LinkRequest
	)
	setup := func() {
		btcKeeper = &mock.BTCKeeperMock{
			GetNetworkFunc: func(ctx sdk.Context) types.Network { return types.Mainnet },
			SetAddressFunc: func(sdk.Context, types.AddressInfo) {},
			LoggerFunc:     func(sdk.Context) log.Logger { return log.TestingLogger() },
		}
		signer = &mock.SignerMock{GetCurrentKeyFunc: func(_ sdk.Context, _ nexus.Chain, keyRole tss.KeyRole) (tss.Key, bool) {
			sk, _ := ecdsa.GenerateKey(btcec.S256(), cryptoRand.Reader)
			return tss.Key{Value: sk.PublicKey, ID: rand.StrBetween(5, 20), Role: keyRole}, true
		}}
		nexusKeeper = &mock.NexusMock{
			GetChainFunc: func(_ sdk.Context, chain string) (nexus.Chain, bool) {
				return nexus.Chain{
					Name:                  chain,
					NativeAsset:           rand.StrBetween(5, 20),
					SupportsForeignAssets: true,
				}, true
			},
			IsAssetRegisteredFunc: func(sdk.Context, string, string) bool { return true },
			LinkAddressesFunc:     func(sdk.Context, nexus.CrossChainAddress, nexus.CrossChainAddress) {},
		}
		ctx = sdk.NewContext(nil, tmproto.Header{Height: rand.PosI64()}, false, log.TestingLogger())
		msg = randomMsgLink()
		server = bitcoinKeeper.NewMsgServerImpl(btcKeeper, signer, nexusKeeper, &mock.VoterMock{}, &mock.SnapshotterMock{})
	}
	repeatCount := 20

	t.Run("happy path", testutils.Func(func(t *testing.T) {
		setup()
		res, err := server.Link(sdk.WrapSDKContext(ctx), msg)
		assert.NoError(t, err)
		assert.Len(t, btcKeeper.SetAddressCalls(), 1)
		assert.Len(t, nexusKeeper.LinkAddressesCalls(), 1)
		assert.Equal(t, exported.Bitcoin, signer.GetCurrentKeyCalls()[0].Chain)
		assert.Equal(t, msg.RecipientChain, nexusKeeper.GetChainCalls()[0].Chain)
		assert.Equal(t, btcKeeper.SetAddressCalls()[0].Address.Address, res.DepositAddr)
		assert.Equal(t, types.Deposit, btcKeeper.SetAddressCalls()[0].Address.Role)
	}).Repeat(repeatCount))

	t.Run("no master key", testutils.Func(func(t *testing.T) {
		setup()
		signer.GetCurrentKeyFunc = func(sdk.Context, nexus.Chain, tss.KeyRole) (tss.Key, bool) { return tss.Key{}, false }
		_, err := server.Link(sdk.WrapSDKContext(ctx), msg)
		assert.Error(t, err)
	}).Repeat(repeatCount))

	t.Run("unknown chain", testutils.Func(func(t *testing.T) {
		setup()
		nexusKeeper.GetChainFunc = func(sdk.Context, string) (nexus.Chain, bool) { return nexus.Chain{}, false }
		_, err := server.Link(sdk.WrapSDKContext(ctx), msg)
		assert.Error(t, err)
	}).Repeat(repeatCount))

	t.Run("asset not registered", testutils.Func(func(t *testing.T) {
		setup()
		nexusKeeper.IsAssetRegisteredFunc = func(sdk.Context, string, string) bool { return false }
		_, err := server.Link(sdk.WrapSDKContext(ctx), msg)
		assert.Error(t, err)
	}).Repeat(repeatCount))
}

func TestHandleMsgConfirmOutpoint(t *testing.T) {
	var (
		btcKeeper *mock.BTCKeeperMock
		voter     *mock.VoterMock
		signer    *mock.SignerMock
		ctx       sdk.Context
		msg       *types.ConfirmOutpointRequest
		server    types.MsgServiceServer
	)
	setup := func() {
		address := randomAddress()
		btcKeeper = &mock.BTCKeeperMock{
			GetOutPointInfoFunc: func(sdk.Context, wire.OutPoint) (types.OutPointInfo, types.OutPointState, bool) {
				return types.OutPointInfo{}, 0, false
			},
			GetAddressFunc: func(sdk.Context, string) (types.AddressInfo, bool) {
				return types.AddressInfo{
					Address:      address.EncodeAddress(),
					RedeemScript: rand.Bytes(200),
					Role:         types.Deposit,
					KeyID:        rand.StrBetween(5, 20),
				}, true
			},
			GetRevoteLockingPeriodFunc:        func(sdk.Context) int64 { return int64(mathRand.Uint32()) },
			GetRequiredConfirmationHeightFunc: func(sdk.Context) uint64 { return mathRand.Uint64() },
			SetPendingOutpointInfoFunc:        func(sdk.Context, vote.PollKey, types.OutPointInfo) {},
		}
		voter = &mock.VoterMock{
			InitializePollFunc: func(sdk.Context, vote.PollKey, int64, ...vote.PollProperty) error { return nil },
		}

		signer = &mock.SignerMock{
			GetCurrentKeyIDFunc: func(ctx sdk.Context, chain nexus.Chain, keyRole tss.KeyRole) (string, bool) {
				return rand.StrBetween(5, 20), true
			},
			GetSnapshotCounterForKeyIDFunc: func(sdk.Context, string) (int64, bool) {
				return rand.PosI64(), true
			},
		}

		ctx = sdk.NewContext(nil, tmproto.Header{Height: rand.PosI64()}, false, log.TestingLogger())
		msg = randomMsgConfirmOutpoint()
		msg.OutPointInfo.Address = address.EncodeAddress()
		server = bitcoinKeeper.NewMsgServerImpl(btcKeeper, signer, &mock.NexusMock{}, voter, &mock.SnapshotterMock{})
	}

	repeatCount := 20
	t.Run("happy path deposit", testutils.Func(func(t *testing.T) {
		setup()
		_, err := server.ConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		events := ctx.EventManager().ABCIEvents()
		assert.NoError(t, err)
		assert.Len(t, testutils.Events(events).Filter(func(event abci.Event) bool { return event.Type == types.EventTypeOutpointConfirmation }), 1)
		assert.Equal(t, msg.OutPointInfo, btcKeeper.SetPendingOutpointInfoCalls()[0].Info)
		assert.Equal(t, voter.InitializePollCalls()[0].Key, btcKeeper.SetPendingOutpointInfoCalls()[0].Key)
	}).Repeat(repeatCount))
	t.Run("happy path consolidation", testutils.Func(func(t *testing.T) {
		setup()
		addr, _ := btcKeeper.GetAddress(ctx, msg.OutPointInfo.Address)
		addr.Role = types.Consolidation
		btcKeeper.GetAddressFunc = func(sdk.Context, string) (types.AddressInfo, bool) {
			return addr, true
		}

		_, err := server.ConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		events := sdk.UnwrapSDKContext(sdk.WrapSDKContext(ctx)).EventManager().ABCIEvents()
		assert.NoError(t, err)
		assert.Len(t, testutils.Events(events).Filter(func(event abci.Event) bool { return event.Type == types.EventTypeOutpointConfirmation }), 1)
		assert.Equal(t, msg.OutPointInfo, btcKeeper.SetPendingOutpointInfoCalls()[0].Info)
		assert.Equal(t, voter.InitializePollCalls()[0].Key, btcKeeper.SetPendingOutpointInfoCalls()[0].Key)
	}).Repeat(repeatCount))
	t.Run("already confirmed", testutils.Func(func(t *testing.T) {
		setup()
		btcKeeper.GetOutPointInfoFunc = func(sdk.Context, wire.OutPoint) (types.OutPointInfo, types.OutPointState, bool) {
			return msg.OutPointInfo, types.OutPointState_Confirmed, true
		}
		_, err := server.ConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		assert.Error(t, err)
	}).Repeat(repeatCount))

	t.Run("already spent", testutils.Func(func(t *testing.T) {
		setup()
		btcKeeper.GetOutPointInfoFunc = func(sdk.Context, wire.OutPoint) (types.OutPointInfo, types.OutPointState, bool) {
			return msg.OutPointInfo, types.OutPointState_Spent, true
		}
		_, err := server.ConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		assert.Error(t, err)
	}).Repeat(repeatCount))

	t.Run("address unknown", testutils.Func(func(t *testing.T) {
		setup()
		btcKeeper.GetAddressFunc = func(sdk.Context, string) (types.AddressInfo, bool) { return types.AddressInfo{}, false }
		_, err := server.ConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		assert.Error(t, err)
	}).Repeat(repeatCount))

	t.Run("init poll failed", testutils.Func(func(t *testing.T) {
		setup()

		voter.InitializePollFunc = func(sdk.Context, vote.PollKey, int64, ...vote.PollProperty) error {
			return fmt.Errorf("poll setup failed")
		}

		_, err := server.ConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		assert.Error(t, err)
	}).Repeat(repeatCount))
}

func TestHandleMsgVoteConfirmOutpoint(t *testing.T) {
	var (
		btcKeeper   *mock.BTCKeeperMock
		voter       *mock.VoterMock
		nexusKeeper *mock.NexusMock
		ctx         sdk.Context
		msg         *types.VoteConfirmOutpointRequest
		info        types.OutPointInfo
		server      types.MsgServiceServer

		currentSecondaryKey tss.Key
		depositAddressInfo  types.AddressInfo
	)
	setup := func() {
		address := randomAddress()
		info = randomOutpointInfo()
		msg = randomMsgVoteConfirmOutpoint()
		msg.OutPoint = info.OutPoint
		depositAddressInfo = types.AddressInfo{
			Address:      address.EncodeAddress(),
			RedeemScript: rand.Bytes(200),
			Role:         types.Deposit,
			KeyID:        rand.StrBetween(5, 20),
		}
		btcKeeper = &mock.BTCKeeperMock{
			GetOutPointInfoFunc: func(sdk.Context, wire.OutPoint) (types.OutPointInfo, types.OutPointState, bool) {
				return types.OutPointInfo{}, 0, false
			},
			SetConfirmedOutpointInfoFunc:  func(sdk.Context, string, types.OutPointInfo) {},
			GetPendingOutPointInfoFunc:    func(sdk.Context, vote.PollKey) (types.OutPointInfo, bool) { return info, true },
			DeletePendingOutPointInfoFunc: func(sdk.Context, vote.PollKey) {},
			GetSignedTxFunc:               func(sdk.Context, chainhash.Hash) (types.SignedTx, bool) { return types.SignedTx{}, false },
			GetAddressFunc: func(sdk.Context, string) (types.AddressInfo, bool) {
				return depositAddressInfo, true
			},
			GetUnconfirmedAmountFunc: func(sdk.Context, string) btcutil.Amount { return 0 },
			SetUnconfirmedAmountFunc: func(sdk.Context, string, btcutil.Amount) {},
		}
		voter = &mock.VoterMock{
			GetPollFunc: func(sdk.Context, vote.PollKey) vote.Poll {
				return &voteMock.PollMock{
					VoteFunc:      func(sdk.ValAddress, codec.ProtoMarshaler) error { return nil },
					GetResultFunc: func() codec.ProtoMarshaler { return &gogoprototypes.BoolValue{Value: true} },
					IsFunc: func(state vote.PollState) bool {
						return state == vote.Completed
					},
				}
			},
		}

		nexusKeeper = &mock.NexusMock{
			EnqueueForTransferFunc: func(sdk.Context, nexus.CrossChainAddress, sdk.Coin) error { return nil },
		}
		privateKey, _ := ecdsa.GenerateKey(btcec.S256(), cryptoRand.Reader)
		currentSecondaryKey = tss.Key{ID: rand.StrBetween(5, 20), Value: privateKey.PublicKey, Role: tss.MasterKey}
		signerKeeper := &mock.SignerMock{
			GetNextKeyFunc:    func(sdk.Context, nexus.Chain, tss.KeyRole) (tss.Key, bool) { return tss.Key{}, false },
			GetCurrentKeyFunc: func(sdk.Context, nexus.Chain, tss.KeyRole) (tss.Key, bool) { return currentSecondaryKey, true },
			AssignNextKeyFunc: func(sdk.Context, nexus.Chain, tss.KeyRole, string) error { return nil },
		}
		ctx = sdk.NewContext(nil, tmproto.Header{Height: rand.PosI64()}, false, log.TestingLogger())
		snapshotter := &mock.SnapshotterMock{GetOperatorFunc: func(ctx sdk.Context, proxy sdk.AccAddress) sdk.ValAddress {
			return rand.Bytes(sdk.AddrLen)
		}}
		server = bitcoinKeeper.NewMsgServerImpl(btcKeeper, signerKeeper, nexusKeeper, voter, snapshotter)
	}

	repeats := 20

	t.Run("happy path confirm deposit to deposit address", testutils.Func(func(t *testing.T) {
		setup()

		_, err := server.VoteConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		assert.NoError(t, err)
		assert.Len(t, btcKeeper.DeletePendingOutPointInfoCalls(), 1)
		assert.Equal(t, info, btcKeeper.SetConfirmedOutpointInfoCalls()[0].Info)
		assert.Equal(t, depositAddressInfo.KeyID, btcKeeper.SetConfirmedOutpointInfoCalls()[0].KeyID)
		assert.Equal(t, info.Address, nexusKeeper.EnqueueForTransferCalls()[0].Sender.Address)
		assert.Equal(t, int64(info.Amount), nexusKeeper.EnqueueForTransferCalls()[0].Amount.Amount.Int64())
	}).Repeat(repeats))

	t.Run("happy path confirm deposit to consolidation address", testutils.Func(func(t *testing.T) {
		setup()
		addr, _ := btcKeeper.GetAddress(ctx, info.Address)
		addr.Role = types.Consolidation
		btcKeeper.GetAddressFunc = func(sdk.Context, string) (types.AddressInfo, bool) {
			return addr, true
		}

		_, err := server.VoteConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		assert.NoError(t, err)
		assert.Len(t, btcKeeper.DeletePendingOutPointInfoCalls(), 1)
		assert.Equal(t, info, btcKeeper.SetConfirmedOutpointInfoCalls()[0].Info)
		assert.Equal(t, depositAddressInfo.KeyID, btcKeeper.SetConfirmedOutpointInfoCalls()[0].KeyID)
		assert.Len(t, nexusKeeper.EnqueueForTransferCalls(), 0)
	}).Repeat(repeats))

	t.Run("happy path confirm deposit to consolidation address in consolidation tx", testutils.Func(func(t *testing.T) {
		setup()
		tx := wire.NewMsgTx(wire.TxVersion)
		hash := tx.TxHash()
		op := wire.NewOutPoint(&hash, info.GetOutPoint().Index)
		info.OutPoint = op.String()
		msg.OutPoint = op.String()
		addr, _ := btcKeeper.GetAddress(ctx, info.Address)
		addr.Role = types.Consolidation
		btcKeeper.GetAddressFunc = func(sdk.Context, string) (types.AddressInfo, bool) {
			return addr, true
		}

		_, err := server.VoteConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		assert.NoError(t, err)
		assert.Len(t, btcKeeper.DeletePendingOutPointInfoCalls(), 1)
		assert.Equal(t, info, btcKeeper.SetConfirmedOutpointInfoCalls()[0].Info)
		assert.Equal(t, depositAddressInfo.KeyID, btcKeeper.SetConfirmedOutpointInfoCalls()[0].KeyID)
		assert.Len(t, nexusKeeper.EnqueueForTransferCalls(), 0)
	}).Repeat(repeats))

	t.Run("happy path confirm deposit to deposit address in consolidation tx", testutils.Func(func(t *testing.T) {
		setup()
		tx := wire.NewMsgTx(wire.TxVersion)
		hash := tx.TxHash()
		op := wire.NewOutPoint(&hash, info.GetOutPoint().Index)
		info.OutPoint = op.String()
		msg.OutPoint = op.String()

		_, err := server.VoteConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		assert.NoError(t, err)
		assert.Len(t, btcKeeper.DeletePendingOutPointInfoCalls(), 1)
		assert.Equal(t, info, btcKeeper.SetConfirmedOutpointInfoCalls()[0].Info)
		assert.Equal(t, depositAddressInfo.KeyID, btcKeeper.SetConfirmedOutpointInfoCalls()[0].KeyID)
		assert.Len(t, nexusKeeper.EnqueueForTransferCalls(), 1)
	}).Repeat(repeats))

	t.Run("happy path reject", testutils.Func(func(t *testing.T) {
		setup()
		voter.GetPollFunc = func(sdk.Context, vote.PollKey) vote.Poll {
			return &voteMock.PollMock{
				VoteFunc:      func(sdk.ValAddress, codec.ProtoMarshaler) error { return nil },
				GetResultFunc: func() codec.ProtoMarshaler { return &gogoprototypes.BoolValue{Value: false} },
				IsFunc: func(state vote.PollState) bool {
					return state == vote.Completed
				},
			}
		}

		_, err := server.VoteConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		assert.NoError(t, err)
		assert.Len(t, btcKeeper.DeletePendingOutPointInfoCalls(), 1)
		assert.Len(t, btcKeeper.SetConfirmedOutpointInfoCalls(), 0)
		assert.Len(t, nexusKeeper.EnqueueForTransferCalls(), 0)
	}).Repeat(repeats))

	t.Run("happy path no result yet", testutils.Func(func(t *testing.T) {
		setup()
		voter.GetPollFunc = func(sdk.Context, vote.PollKey) vote.Poll {
			return &voteMock.PollMock{
				VoteFunc:      func(sdk.ValAddress, codec.ProtoMarshaler) error { return nil },
				GetResultFunc: func() codec.ProtoMarshaler { return nil },
				IsFunc: func(state vote.PollState) bool {
					return state == vote.Pending
				},
			}
		}

		_, err := server.VoteConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		assert.NoError(t, err)
		assert.Len(t, btcKeeper.DeletePendingOutPointInfoCalls(), 0)
		assert.Len(t, btcKeeper.SetConfirmedOutpointInfoCalls(), 0)
		assert.Len(t, nexusKeeper.EnqueueForTransferCalls(), 0)
	}).Repeat(repeats))

	t.Run("happy path poll already completed", testutils.Func(func(t *testing.T) {
		setup()
		btcKeeper.GetPendingOutPointInfoFunc = func(sdk.Context, vote.PollKey) (types.OutPointInfo, bool) {
			return types.OutPointInfo{}, false
		}
		btcKeeper.GetOutPointInfoFunc = func(sdk.Context, wire.OutPoint) (types.OutPointInfo, types.OutPointState, bool) {
			return info, types.OutPointState_Confirmed, true
		}

		_, err := server.VoteConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		assert.NoError(t, err)
		assert.Len(t, btcKeeper.DeletePendingOutPointInfoCalls(), 0)
		assert.Len(t, btcKeeper.SetConfirmedOutpointInfoCalls(), 0)
		assert.Len(t, nexusKeeper.EnqueueForTransferCalls(), 0)
	}).Repeat(repeats))

	t.Run("happy path second poll (outpoint already confirmed)", testutils.Func(func(t *testing.T) {
		setup()
		btcKeeper.GetOutPointInfoFunc = func(sdk.Context, wire.OutPoint) (types.OutPointInfo, types.OutPointState, bool) {
			return info, types.OutPointState_Confirmed, true
		}

		_, err := server.VoteConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		assert.NoError(t, err)
		assert.Len(t, btcKeeper.DeletePendingOutPointInfoCalls(), 1)
		assert.Len(t, btcKeeper.SetConfirmedOutpointInfoCalls(), 0)
		assert.Len(t, nexusKeeper.EnqueueForTransferCalls(), 0)
	}).Repeat(repeats))

	t.Run("happy path already spent", testutils.Func(func(t *testing.T) {
		setup()
		btcKeeper.GetOutPointInfoFunc = func(sdk.Context, wire.OutPoint) (types.OutPointInfo, types.OutPointState, bool) {
			return info, types.OutPointState_Spent, true
		}

		_, err := server.VoteConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		assert.NoError(t, err)
		assert.Len(t, btcKeeper.DeletePendingOutPointInfoCalls(), 1)
		assert.Len(t, btcKeeper.SetConfirmedOutpointInfoCalls(), 0)
		assert.Len(t, nexusKeeper.EnqueueForTransferCalls(), 0)
	}).Repeat(repeats))

	t.Run("unknown outpoint", testutils.Func(func(t *testing.T) {
		setup()
		btcKeeper.GetPendingOutPointInfoFunc =
			func(sdk.Context, vote.PollKey) (types.OutPointInfo, bool) { return types.OutPointInfo{}, false }

		_, err := server.VoteConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		assert.Error(t, err)
	}).Repeat(repeats))

	t.Run("tally failed", testutils.Func(func(t *testing.T) {
		setup()
		voter.GetPollFunc = func(sdk.Context, vote.PollKey) vote.Poll {
			return &voteMock.PollMock{
				VoteFunc: func(sdk.ValAddress, codec.ProtoMarshaler) error { return fmt.Errorf("some error") },
			}
		}

		_, err := server.VoteConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		assert.Error(t, err)
	}).Repeat(repeats))

	t.Run("enqueue transfer failed", testutils.Func(func(t *testing.T) {
		setup()
		nexusKeeper.EnqueueForTransferFunc = func(sdk.Context, nexus.CrossChainAddress, sdk.Coin) error {
			return fmt.Errorf("failed")
		}

		_, err := server.VoteConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		assert.Error(t, err)
	}).Repeat(repeats))

	t.Run("outpoint does not match poll", testutils.Func(func(t *testing.T) {
		setup()
		info = randomOutpointInfo()

		_, err := server.VoteConfirmOutpoint(sdk.WrapSDKContext(ctx), msg)
		assert.Error(t, err)
	}).Repeat(repeats))
}

func TestCreateMasterTx(t *testing.T) {
	var (
		btcKeeper   *mock.BTCKeeperMock
		voter       *mock.VoterMock
		nexusKeeper *mock.NexusMock
		server      types.MsgServiceServer

		ctx                    sdk.Context
		masterKey              tss.Key
		oldMasterKey           tss.Key
		secondaryKey           tss.Key
		consolidationKey       tss.Key
		externalKeys           []tss.Key
		masterKeyRotationCount int64
		inputs                 []types.OutPointInfo
		inputTotal             btcutil.Amount
	)
	setup := func() {
		ctx = sdk.NewContext(nil, tmproto.Header{Height: rand.PosI64()}, false, log.TestingLogger())
		masterKey = createRandomKey(tss.MasterKey, time.Now())
		oldMasterKey = createRandomKey(tss.MasterKey)
		secondaryKey = createRandomKey(tss.SecondaryKey)
		consolidationKey = createRandomKey(tss.MasterKey)

		externalKeyCount := types.DefaultParams().ExternalMultisigThreshold.Denominator
		externalKeys = make([]tss.Key, externalKeyCount)
		for i := 0; i < int(externalKeyCount); i++ {
			externalKeys[i] = createRandomKey(tss.ExternalKey)
		}

		masterKeyRotationCount = rand.I64Between(100, 1000)
		oldMasterKeyRotationCount := masterKeyRotationCount - (masterKeyRotationCount-1)%types.DefaultParams().MasterKeyRetentionPeriod

		inputCount := int(types.DefaultParams().MaxInputCount)
		inputs = make([]types.OutPointInfo, inputCount)
		inputTotal = 0
		for i := 0; i < inputCount; i++ {
			inputs[i] = randomOutpointInfo()
			inputTotal += inputs[i].Amount
		}

		btcKeeper = &mock.BTCKeeperMock{
			GetConfirmedOutpointInfoQueueForKeyFunc: func(ctx sdk.Context, keyID string) utils.KVQueue {
				if keyID == masterKey.ID {
					dequeueCount := 0

					return &utilsmock.KVQueueMock{
						IsEmptyFunc: func() bool { return true },
						DequeueFunc: func(value codec.ProtoMarshaler) bool {
							if dequeueCount >= len(inputs) {
								return false
							}

							types.ModuleCdc.MustUnmarshalBinaryLengthPrefixed(
								types.ModuleCdc.MustMarshalBinaryLengthPrefixed(&inputs[dequeueCount]),
								value,
							)

							dequeueCount++
							return true
						},
					}
				}

				return &utilsmock.KVQueueMock{}
			},
			GetMinOutputAmountFunc: func(ctx sdk.Context) btcutil.Amount {
				satoshi, err := types.ToSatoshiCoin(types.DefaultParams().MinOutputAmount)
				if err != nil {
					panic(err)
				}

				return btcutil.Amount(satoshi.Amount.Int64())
			},
			GetMaxSecondaryOutputAmountFunc: func(ctx sdk.Context) btcutil.Amount {
				satoshi, err := types.ToSatoshiCoin(types.DefaultParams().MaxSecondaryOutputAmount)
				if err != nil {
					panic(err)
				}

				return btcutil.Amount(satoshi.Amount.Int64())
			},
			GetMaxInputCountFunc: func(ctx sdk.Context) int64 {
				return types.DefaultParams().MaxInputCount
			},
			GetExternalMultisigThresholdFunc: func(ctx sdk.Context) utils.Threshold {
				return types.DefaultParams().ExternalMultisigThreshold
			},
			GetAnyoneCanSpendAddressFunc: func(ctx sdk.Context) types.AddressInfo {
				return types.NewAnyoneCanSpendAddress(types.DefaultParams().Network)
			},
			GetUnsignedTxFunc: func(ctx sdk.Context, keyRole tss.KeyRole) (types.UnsignedTx, bool) {
				return types.UnsignedTx{}, false
			},
			GetMasterKeyRetentionPeriodFunc: func(ctx sdk.Context) int64 {
				return types.DefaultParams().MasterKeyRetentionPeriod
			},
			GetMasterAddressLockDurationFunc: func(ctx sdk.Context) time.Duration {
				return types.DefaultParams().MasterAddressLockDuration
			},
			GetNetworkFunc: func(ctx sdk.Context) types.Network {
				return types.DefaultParams().Network
			},
			GetAddressFunc: func(_ sdk.Context, encodedAddress string) (types.AddressInfo, bool) {
				return types.AddressInfo{
					Address:      encodedAddress,
					RedeemScript: nil,
					KeyID:        masterKey.ID,
				}, true
			},
			GetExternalKeyIDsFunc: func(ctx sdk.Context) ([]string, bool) {
				externalKeyIDs := make([]string, len(externalKeys))
				for i := 0; i < len(externalKeyIDs); i++ {
					externalKeyIDs[i] = externalKeys[i].ID
				}

				return externalKeyIDs, true
			},
			GetUnconfirmedAmountFunc: func(ctx sdk.Context, keyID string) btcutil.Amount { return 0 },
			DeleteOutpointInfoFunc:   func(ctx sdk.Context, outPoint wire.OutPoint) {},
			SetSpentOutpointInfoFunc: func(ctx sdk.Context, info types.OutPointInfo) {},
			SetAddressFunc:           func(ctx sdk.Context, address types.AddressInfo) {},
			SetUnsignedTxFunc:        func(ctx sdk.Context, keyRole tss.KeyRole, tx types.UnsignedTx) {},
		}
		voter = &mock.VoterMock{}
		nexusKeeper = &mock.NexusMock{}
		signerKeeper := &mock.SignerMock{
			GetCurrentKeyFunc: func(ctx sdk.Context, chain nexus.Chain, keyRole tss.KeyRole) (tss.Key, bool) {
				switch keyRole {
				case tss.MasterKey:
					return masterKey, true
				case tss.SecondaryKey:
					return secondaryKey, true
				default:
					return tss.Key{}, false
				}
			},
			GetKeyFunc: func(ctx sdk.Context, keyID string) (tss.Key, bool) {
				switch keyID {
				case masterKey.ID:
					return masterKey, true
				case oldMasterKey.ID:
					return masterKey, true
				case secondaryKey.ID:
					return secondaryKey, true
				case consolidationKey.ID:
					return consolidationKey, true
				default:
					for _, externalKey := range externalKeys {
						if keyID == externalKey.ID {
							return externalKey, true
						}
					}

					return tss.Key{}, false
				}
			},
			GetRotationCountFunc: func(ctx sdk.Context, chain nexus.Chain, keyRole tss.KeyRole) int64 {
				if keyRole == tss.MasterKey {
					return masterKeyRotationCount
				}

				return 0
			},
			GetKeyByRotationCountFunc: func(ctx sdk.Context, chain nexus.Chain, keyRole tss.KeyRole, rotationCount int64) (tss.Key, bool) {
				if keyRole == tss.MasterKey && rotationCount == oldMasterKeyRotationCount {
					return oldMasterKey, true
				}

				return tss.Key{}, false
			},
			GetNextKeyFunc: func(ctx sdk.Context, chain nexus.Chain, keyRole tss.KeyRole) (tss.Key, bool) {
				return tss.Key{}, false
			},
			AssertMatchesRequirementsFunc: func(ctx sdk.Context, snapshotter snapshot.Snapshotter, chain nexus.Chain, keyID string, keyRole tss.KeyRole) error {
				return nil
			},
		}
		snapshotter := &mock.SnapshotterMock{}
		server = bitcoinKeeper.NewMsgServerImpl(btcKeeper, signerKeeper, nexusKeeper, voter, snapshotter)
	}

	t.Run("should create master consolidation transaction sending no coin to the secondary key when the amount is not set", testutils.Func(func(t *testing.T) {
		setup()

		req := types.NewCreateMasterTxRequest(rand.Bytes(sdk.AddrLen), consolidationKey.ID, 0)
		_, err := server.CreateMasterTx(sdk.WrapSDKContext(ctx), req)
		assert.NoError(t, err)

		network := types.DefaultParams().Network
		expectedAnyoneCanSpendAddress := types.NewAnyoneCanSpendAddress(network).Address
		expectedMasterConsolidationAddress := types.NewMasterConsolidationAddress(consolidationKey, oldMasterKey, types.DefaultParams().ExternalMultisigThreshold.Numerator, externalKeys, masterKey.RotatedAt.Add(types.DefaultParams().MasterAddressLockDuration), network).Address
		minOutputAmount, err := types.ToSatoshiCoin(types.DefaultParams().MinOutputAmount)
		if err != nil {
			panic(err)
		}

		assert.Len(t, btcKeeper.SetUnsignedTxCalls(), 1)
		assert.Len(t, btcKeeper.DeleteOutpointInfoCalls(), len(inputs))
		assert.Len(t, btcKeeper.SetSpentOutpointInfoCalls(), len(inputs))
		assert.Len(t, btcKeeper.SetAddressCalls(), 2)
		assert.Equal(t, expectedMasterConsolidationAddress, btcKeeper.SetAddressCalls()[0].Address.Address)
		assert.Equal(t, expectedMasterConsolidationAddress, btcKeeper.SetAddressCalls()[1].Address.Address)
		actualUnsignedTx := btcKeeper.SetUnsignedTxCalls()[0].Tx
		assert.Len(t, actualUnsignedTx.GetTx().TxIn, len(inputs))
		for i, txIn := range actualUnsignedTx.GetTx().TxIn {
			assert.Equal(t, txIn.Sequence, wire.MaxTxInSequenceNum)
			assert.Equal(t, txIn.PreviousOutPoint.String(), inputs[i].OutPoint)
		}
		assertTxOutputs(t, actualUnsignedTx.GetTx(),
			types.Output{
				Recipient: types.MustDecodeAddress(expectedAnyoneCanSpendAddress, network),
				Amount:    btcutil.Amount(minOutputAmount.Amount.Int64()),
			},
			types.Output{
				Recipient: types.MustDecodeAddress(expectedMasterConsolidationAddress, network),
			},
		)
		assert.Equal(t, uint32(0), actualUnsignedTx.GetTx().LockTime)
	}))

	t.Run("should create master consolidation transaction sending coins to the secondary key when the amount is set", testutils.Func(func(t *testing.T) {
		setup()

		secondaryKeyAmount := btcutil.Amount(rand.I64Between(1000, 10000))
		req := types.NewCreateMasterTxRequest(rand.Bytes(sdk.AddrLen), consolidationKey.ID, secondaryKeyAmount)
		_, err := server.CreateMasterTx(sdk.WrapSDKContext(ctx), req)
		assert.NoError(t, err)

		network := types.DefaultParams().Network
		expectedAnyoneCanSpendAddress := types.NewAnyoneCanSpendAddress(network).Address
		expectedSecondaryConsolidationAddress := types.NewSecondaryConsolidationAddress(secondaryKey, network).Address
		expectedMasterConsolidationAddress := types.NewMasterConsolidationAddress(consolidationKey, oldMasterKey, types.DefaultParams().ExternalMultisigThreshold.Numerator, externalKeys, masterKey.RotatedAt.Add(types.DefaultParams().MasterAddressLockDuration), network).Address
		minOutputAmount, err := types.ToSatoshiCoin(types.DefaultParams().MinOutputAmount)
		if err != nil {
			panic(err)
		}

		assert.Len(t, btcKeeper.SetUnsignedTxCalls(), 1)
		assert.Len(t, btcKeeper.DeleteOutpointInfoCalls(), len(inputs))
		assert.Len(t, btcKeeper.SetSpentOutpointInfoCalls(), len(inputs))
		assert.Len(t, btcKeeper.SetAddressCalls(), 3)
		assert.Equal(t, expectedSecondaryConsolidationAddress, btcKeeper.SetAddressCalls()[0].Address.Address)
		assert.Equal(t, expectedMasterConsolidationAddress, btcKeeper.SetAddressCalls()[1].Address.Address)
		assert.Equal(t, expectedMasterConsolidationAddress, btcKeeper.SetAddressCalls()[2].Address.Address)
		actualUnsignedTx := btcKeeper.SetUnsignedTxCalls()[0].Tx
		assert.Len(t, actualUnsignedTx.GetTx().TxIn, len(inputs))
		for i, txIn := range actualUnsignedTx.GetTx().TxIn {
			assert.Equal(t, txIn.Sequence, wire.MaxTxInSequenceNum)
			assert.Equal(t, txIn.PreviousOutPoint.String(), inputs[i].OutPoint)
		}
		assertTxOutputs(t, actualUnsignedTx.GetTx(),
			types.Output{
				Recipient: types.MustDecodeAddress(expectedSecondaryConsolidationAddress, network),
				Amount:    secondaryKeyAmount,
			},
			types.Output{
				Recipient: types.MustDecodeAddress(expectedAnyoneCanSpendAddress, network),
				Amount:    btcutil.Amount(minOutputAmount.Amount.Int64()),
			},
			types.Output{
				Recipient: types.MustDecodeAddress(expectedMasterConsolidationAddress, network),
			},
		)
		assert.Equal(t, uint32(0), actualUnsignedTx.GetTx().LockTime)
	}))

	t.Run("should return error if consolidating to a new key while the current key still has UTXO", testutils.Func(func(t *testing.T) {
		setup()

		btcKeeper.GetConfirmedOutpointInfoQueueForKeyFunc = func(ctx sdk.Context, keyID string) utils.KVQueue {
			if keyID == masterKey.ID {
				dequeueCount := 0

				return &utilsmock.KVQueueMock{
					IsEmptyFunc: func() bool { return false },
					DequeueFunc: func(value codec.ProtoMarshaler) bool {
						if dequeueCount >= len(inputs) {
							return false
						}

						types.ModuleCdc.MustUnmarshalBinaryLengthPrefixed(
							types.ModuleCdc.MustMarshalBinaryLengthPrefixed(&inputs[dequeueCount]),
							value,
						)

						dequeueCount++
						return true
					},
				}
			}

			return &utilsmock.KVQueueMock{}
		}

		req := types.NewCreateMasterTxRequest(rand.Bytes(sdk.AddrLen), consolidationKey.ID, 0)
		_, err := server.CreateMasterTx(sdk.WrapSDKContext(ctx), req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "still has outpoints to be signed and therefore it cannot be rotated out yet")
	}))

	t.Run("should return error if consolidating to a new key while the current key still has unconfirmed amount", testutils.Func(func(t *testing.T) {
		setup()

		btcKeeper.GetUnconfirmedAmountFunc = func(ctx sdk.Context, keyID string) btcutil.Amount {
			return btcutil.Amount(rand.I64Between(1, 100))
		}

		req := types.NewCreateMasterTxRequest(rand.Bytes(sdk.AddrLen), consolidationKey.ID, 0)
		_, err := server.CreateMasterTx(sdk.WrapSDKContext(ctx), req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "still has unconfirmed outpoints and therefore it cannot be rotated out yet")
	}))
}

func TestCreatePendingTransfersTx(t *testing.T) {
	var (
		btcKeeper   *mock.BTCKeeperMock
		voter       *mock.VoterMock
		nexusKeeper *mock.NexusMock
		server      types.MsgServiceServer

		ctx                    sdk.Context
		masterKey              tss.Key
		oldMasterKey           tss.Key
		secondaryKey           tss.Key
		consolidationKey       tss.Key
		externalKeys           []tss.Key
		masterKeyRotationCount int64
		inputs                 []types.OutPointInfo
		inputTotal             btcutil.Amount
		transfers              []nexus.CrossChainTransfer
	)
	setup := func() {
		ctx = sdk.NewContext(nil, tmproto.Header{Height: rand.PosI64()}, false, log.TestingLogger())
		masterKey = createRandomKey(tss.MasterKey, time.Now())
		oldMasterKey = createRandomKey(tss.MasterKey)
		secondaryKey = createRandomKey(tss.SecondaryKey)
		consolidationKey = createRandomKey(tss.SecondaryKey)

		externalKeyCount := types.DefaultParams().ExternalMultisigThreshold.Denominator
		externalKeys = make([]tss.Key, externalKeyCount)
		for i := 0; i < int(externalKeyCount); i++ {
			externalKeys[i] = createRandomKey(tss.ExternalKey)
		}

		masterKeyRotationCount = rand.I64Between(100, 1000)
		oldMasterKeyRotationCount := masterKeyRotationCount - (masterKeyRotationCount-1)%types.DefaultParams().MasterKeyRetentionPeriod

		inputCount := int(types.DefaultParams().MaxInputCount)
		inputs = make([]types.OutPointInfo, inputCount)
		inputTotal = 0
		for i := 0; i < inputCount; i++ {
			inputs[i] = randomOutpointInfo()
			inputTotal += inputs[i].Amount
		}

		transfers = []nexus.CrossChainTransfer{}
		outputTotal := btcutil.Amount(0)
		for {
			transfer := randomCrossChainTransfer(int64(inputTotal))
			if transfer.Asset.Amount.AddRaw(int64(outputTotal)).Int64() > int64(inputTotal) {
				break
			}

			transfers = append(transfers, transfer)
			outputTotal += btcutil.Amount(transfer.Asset.Amount.Int64())
		}

		btcKeeper = &mock.BTCKeeperMock{
			LoggerFunc: func(ctx sdk.Context) log.Logger { return log.TestingLogger() },
			GetConfirmedOutpointInfoQueueForKeyFunc: func(ctx sdk.Context, keyID string) utils.KVQueue {
				if keyID == secondaryKey.ID {
					dequeueCount := 0

					return &utilsmock.KVQueueMock{
						IsEmptyFunc: func() bool { return true },
						DequeueFunc: func(value codec.ProtoMarshaler) bool {
							if dequeueCount >= len(inputs) {
								return false
							}

							types.ModuleCdc.MustUnmarshalBinaryLengthPrefixed(
								types.ModuleCdc.MustMarshalBinaryLengthPrefixed(&inputs[dequeueCount]),
								value,
							)

							dequeueCount++
							return true
						},
					}
				}

				return &utilsmock.KVQueueMock{}
			},
			GetMinOutputAmountFunc: func(ctx sdk.Context) btcutil.Amount {
				satoshi, err := types.ToSatoshiCoin(types.DefaultParams().MinOutputAmount)
				if err != nil {
					panic(err)
				}

				return btcutil.Amount(satoshi.Amount.Int64())
			},
			GetMaxInputCountFunc: func(ctx sdk.Context) int64 {
				return types.DefaultParams().MaxInputCount
			},
			GetExternalMultisigThresholdFunc: func(ctx sdk.Context) utils.Threshold {
				return types.DefaultParams().ExternalMultisigThreshold
			},
			GetAnyoneCanSpendAddressFunc: func(ctx sdk.Context) types.AddressInfo {
				return types.NewAnyoneCanSpendAddress(types.DefaultParams().Network)
			},
			GetUnsignedTxFunc: func(ctx sdk.Context, keyRole tss.KeyRole) (types.UnsignedTx, bool) {
				return types.UnsignedTx{}, false
			},
			GetMasterKeyRetentionPeriodFunc: func(ctx sdk.Context) int64 {
				return types.DefaultParams().MasterKeyRetentionPeriod
			},
			GetMasterAddressLockDurationFunc: func(ctx sdk.Context) time.Duration {
				return types.DefaultParams().MasterAddressLockDuration
			},
			GetNetworkFunc: func(ctx sdk.Context) types.Network {
				return types.DefaultParams().Network
			},
			GetAddressFunc: func(_ sdk.Context, encodedAddress string) (types.AddressInfo, bool) {
				return types.AddressInfo{
					Address:      encodedAddress,
					RedeemScript: nil,
					KeyID:        masterKey.ID,
				}, true
			},
			GetExternalKeyIDsFunc: func(ctx sdk.Context) ([]string, bool) {
				externalKeyIDs := make([]string, len(externalKeys))
				for i := 0; i < len(externalKeyIDs); i++ {
					externalKeyIDs[i] = externalKeys[i].ID
				}

				return externalKeyIDs, true
			},
			GetDustAmountFunc:        func(ctx sdk.Context, encodedAddress string) btcutil.Amount { return 0 },
			GetUnconfirmedAmountFunc: func(ctx sdk.Context, keyID string) btcutil.Amount { return 0 },
			DeleteDustAmountFunc:     func(ctx sdk.Context, encodedAddress string) {},
			DeleteOutpointInfoFunc:   func(ctx sdk.Context, outPoint wire.OutPoint) {},
			SetSpentOutpointInfoFunc: func(ctx sdk.Context, info types.OutPointInfo) {},
			SetAddressFunc:           func(ctx sdk.Context, address types.AddressInfo) {},
			SetUnsignedTxFunc:        func(ctx sdk.Context, keyRole tss.KeyRole, tx types.UnsignedTx) {},
		}
		voter = &mock.VoterMock{}
		nexusKeeper = &mock.NexusMock{
			GetTransfersForChainFunc: func(ctx sdk.Context, chain nexus.Chain, state nexus.TransferState) []nexus.CrossChainTransfer {
				if chain == exported.Bitcoin && state == nexus.Pending {
					return transfers
				}

				return []nexus.CrossChainTransfer{}
			},
			ArchivePendingTransferFunc: func(ctx sdk.Context, transfer nexus.CrossChainTransfer) {},
		}
		signerKeeper := &mock.SignerMock{
			GetCurrentKeyFunc: func(ctx sdk.Context, chain nexus.Chain, keyRole tss.KeyRole) (tss.Key, bool) {
				switch keyRole {
				case tss.MasterKey:
					return masterKey, true
				case tss.SecondaryKey:
					return secondaryKey, true
				default:
					return tss.Key{}, false
				}
			},
			GetKeyFunc: func(ctx sdk.Context, keyID string) (tss.Key, bool) {
				switch keyID {
				case masterKey.ID:
					return masterKey, true
				case oldMasterKey.ID:
					return masterKey, true
				case secondaryKey.ID:
					return secondaryKey, true
				case consolidationKey.ID:
					return consolidationKey, true
				default:
					for _, externalKey := range externalKeys {
						if keyID == externalKey.ID {
							return externalKey, true
						}
					}

					return tss.Key{}, false
				}
			},
			GetRotationCountFunc: func(ctx sdk.Context, chain nexus.Chain, keyRole tss.KeyRole) int64 {
				if keyRole == tss.MasterKey {
					return masterKeyRotationCount
				}

				return 0
			},
			GetKeyByRotationCountFunc: func(ctx sdk.Context, chain nexus.Chain, keyRole tss.KeyRole, rotationCount int64) (tss.Key, bool) {
				if keyRole == tss.MasterKey && rotationCount == oldMasterKeyRotationCount {
					return oldMasterKey, true
				}

				return tss.Key{}, false
			},
			GetNextKeyFunc: func(ctx sdk.Context, chain nexus.Chain, keyRole tss.KeyRole) (tss.Key, bool) {
				return tss.Key{}, false
			},
			AssertMatchesRequirementsFunc: func(ctx sdk.Context, snapshotter snapshot.Snapshotter, chain nexus.Chain, keyID string, keyRole tss.KeyRole) error {
				return nil
			},
		}
		snapshotter := &mock.SnapshotterMock{}
		server = bitcoinKeeper.NewMsgServerImpl(btcKeeper, signerKeeper, nexusKeeper, voter, snapshotter)
	}

	t.Run("should create secondary consolidation transaction sending no coin to the master key when the amount is not set", testutils.Func(func(t *testing.T) {
		setup()

		req := types.NewCreatePendingTransfersTxRequest(rand.Bytes(sdk.AddrLen), consolidationKey.ID, 0)
		_, err := server.CreatePendingTransfersTx(sdk.WrapSDKContext(ctx), req)
		assert.NoError(t, err)

		network := types.DefaultParams().Network
		expectedAnyoneCanSpendAddress := types.NewAnyoneCanSpendAddress(network).Address
		expectedSecondaryConsolidationAddress := types.NewSecondaryConsolidationAddress(consolidationKey, network).Address
		minOutputAmount, err := types.ToSatoshiCoin(types.DefaultParams().MinOutputAmount)
		if err != nil {
			panic(err)
		}

		assert.Len(t, btcKeeper.SetUnsignedTxCalls(), 1)
		assert.Len(t, btcKeeper.DeleteOutpointInfoCalls(), len(inputs))
		assert.Len(t, btcKeeper.SetSpentOutpointInfoCalls(), len(inputs))
		assert.Len(t, btcKeeper.SetAddressCalls(), 2)
		assert.Equal(t, expectedSecondaryConsolidationAddress, btcKeeper.SetAddressCalls()[0].Address.Address)
		assert.Equal(t, expectedSecondaryConsolidationAddress, btcKeeper.SetAddressCalls()[1].Address.Address)
		assert.Len(t, nexusKeeper.ArchivePendingTransferCalls(), len(transfers))
		actualUnsignedTx := btcKeeper.SetUnsignedTxCalls()[0].Tx
		assert.Len(t, actualUnsignedTx.GetTx().TxIn, len(inputs))
		for i, txIn := range actualUnsignedTx.GetTx().TxIn {
			assert.Equal(t, txIn.Sequence, wire.MaxTxInSequenceNum)
			assert.Equal(t, txIn.PreviousOutPoint.String(), inputs[i].OutPoint)
		}
		var expectedOutputs []types.Output
		for _, transfer := range transfers {
			expectedOutputs = append(expectedOutputs, types.Output{
				Recipient: types.MustDecodeAddress(transfer.Recipient.Address, network),
				Amount:    btcutil.Amount(transfer.Asset.Amount.Int64()),
			})
		}
		expectedOutputs = append(expectedOutputs, types.Output{
			Recipient: types.MustDecodeAddress(expectedAnyoneCanSpendAddress, network),
			Amount:    btcutil.Amount(minOutputAmount.Amount.Int64()),
		})
		expectedOutputs = append(expectedOutputs, types.Output{
			Recipient: types.MustDecodeAddress(expectedSecondaryConsolidationAddress, network),
		})
		assertTxOutputs(t, actualUnsignedTx.GetTx(), expectedOutputs...)
		assert.Equal(t, uint32(0), actualUnsignedTx.GetTx().LockTime)
	}))

	t.Run("should create secondary consolidation transaction sending coin to the master key when the amount is set", testutils.Func(func(t *testing.T) {
		setup()

		masterKeyAmount := btcutil.Amount(transfers[len(transfers)-1].Asset.Amount.Int64())
		transfers = transfers[:len(transfers)-1]
		req := types.NewCreatePendingTransfersTxRequest(rand.Bytes(sdk.AddrLen), consolidationKey.ID, masterKeyAmount)
		_, err := server.CreatePendingTransfersTx(sdk.WrapSDKContext(ctx), req)
		assert.NoError(t, err)

		network := types.DefaultParams().Network
		expectedAnyoneCanSpendAddress := types.NewAnyoneCanSpendAddress(network).Address
		expectedSecondaryConsolidationAddress := types.NewSecondaryConsolidationAddress(consolidationKey, network).Address
		expectedMasterConsolidationAddress := types.NewMasterConsolidationAddress(masterKey, oldMasterKey, types.DefaultParams().ExternalMultisigThreshold.Numerator, externalKeys, masterKey.RotatedAt.Add(types.DefaultParams().MasterAddressLockDuration), network).Address
		minOutputAmount, err := types.ToSatoshiCoin(types.DefaultParams().MinOutputAmount)
		if err != nil {
			panic(err)
		}

		assert.Len(t, btcKeeper.SetUnsignedTxCalls(), 1)
		assert.Len(t, btcKeeper.DeleteOutpointInfoCalls(), len(inputs))
		assert.Len(t, btcKeeper.SetSpentOutpointInfoCalls(), len(inputs))
		assert.Len(t, btcKeeper.SetAddressCalls(), 3)
		assert.Equal(t, expectedMasterConsolidationAddress, btcKeeper.SetAddressCalls()[0].Address.Address)
		assert.Equal(t, expectedSecondaryConsolidationAddress, btcKeeper.SetAddressCalls()[1].Address.Address)
		assert.Equal(t, expectedSecondaryConsolidationAddress, btcKeeper.SetAddressCalls()[2].Address.Address)
		assert.Len(t, nexusKeeper.ArchivePendingTransferCalls(), len(transfers))
		actualUnsignedTx := btcKeeper.SetUnsignedTxCalls()[0].Tx
		assert.Len(t, actualUnsignedTx.GetTx().TxIn, len(inputs))
		for i, txIn := range actualUnsignedTx.GetTx().TxIn {
			assert.Equal(t, txIn.Sequence, wire.MaxTxInSequenceNum)
			assert.Equal(t, txIn.PreviousOutPoint.String(), inputs[i].OutPoint)
		}
		var expectedOutputs []types.Output
		for _, transfer := range transfers {
			expectedOutputs = append(expectedOutputs, types.Output{
				Recipient: types.MustDecodeAddress(transfer.Recipient.Address, network),
				Amount:    btcutil.Amount(transfer.Asset.Amount.Int64()),
			})
		}
		expectedOutputs = append(expectedOutputs, types.Output{
			Recipient: types.MustDecodeAddress(expectedAnyoneCanSpendAddress, network),
			Amount:    btcutil.Amount(minOutputAmount.Amount.Int64()),
		})
		expectedOutputs = append(expectedOutputs, types.Output{
			Recipient: types.MustDecodeAddress(expectedMasterConsolidationAddress, network),
			Amount:    masterKeyAmount,
		})
		expectedOutputs = append(expectedOutputs, types.Output{
			Recipient: types.MustDecodeAddress(expectedSecondaryConsolidationAddress, network),
		})
		assertTxOutputs(t, actualUnsignedTx.GetTx(), expectedOutputs...)
		assert.Equal(t, uint32(0), actualUnsignedTx.GetTx().LockTime)
	}))

	t.Run("should return error if consolidating to a new key while the current key still has UTXO", testutils.Func(func(t *testing.T) {
		setup()

		btcKeeper.GetConfirmedOutpointInfoQueueForKeyFunc = func(ctx sdk.Context, keyID string) utils.KVQueue {
			if keyID == secondaryKey.ID {
				dequeueCount := 0

				return &utilsmock.KVQueueMock{
					IsEmptyFunc: func() bool { return false },
					DequeueFunc: func(value codec.ProtoMarshaler) bool {
						if dequeueCount >= len(inputs) {
							return false
						}

						types.ModuleCdc.MustUnmarshalBinaryLengthPrefixed(
							types.ModuleCdc.MustMarshalBinaryLengthPrefixed(&inputs[dequeueCount]),
							value,
						)

						dequeueCount++
						return true
					},
				}
			}

			return &utilsmock.KVQueueMock{}
		}

		req := types.NewCreatePendingTransfersTxRequest(rand.Bytes(sdk.AddrLen), consolidationKey.ID, 0)
		_, err := server.CreatePendingTransfersTx(sdk.WrapSDKContext(ctx), req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "still has outpoints to be signed and therefore it cannot be rotated out yet")
	}))

	t.Run("should return error if consolidating to a new key while the current key still has unconfirmed amount", testutils.Func(func(t *testing.T) {
		setup()

		btcKeeper.GetUnconfirmedAmountFunc = func(ctx sdk.Context, keyID string) btcutil.Amount {
			return btcutil.Amount(rand.I64Between(1, 100))
		}

		req := types.NewCreatePendingTransfersTxRequest(rand.Bytes(sdk.AddrLen), consolidationKey.ID, 0)
		_, err := server.CreatePendingTransfersTx(sdk.WrapSDKContext(ctx), req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "still has unconfirmed outpoints and therefore it cannot be rotated out yet")
	}))
}

func assertTxOutputs(t *testing.T, tx *wire.MsgTx, outputs ...types.Output) {
	assert.Len(t, tx.TxOut, len(outputs))

	for _, expected := range outputs {
		found := false
		pkScript, err := txscript.PayToAddrScript(expected.Recipient)
		if err != nil {
			panic(err)
		}

		for _, actual := range tx.TxOut {
			if !bytes.Equal(pkScript, actual.PkScript) {
				continue
			}

			// ignore if amount is 0
			if expected.Amount != 0 && expected.Amount != btcutil.Amount(actual.Value) {
				continue
			}

			found = true
			break
		}

		assert.True(t, found, fmt.Sprintf("expected output %s-%d not found", expected.Recipient, expected.Amount))
	}
}

func createRandomKey(keyRole tss.KeyRole, rotatedAt ...time.Time) tss.Key {
	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		panic(err)
	}

	key := tss.Key{
		ID:        rand.Str(10),
		Value:     privKey.PublicKey,
		Role:      keyRole,
		RotatedAt: nil,
	}

	if len(rotatedAt) > 0 {
		key.RotatedAt = &rotatedAt[0]
	}

	return key
}

func randomMsgLink() *types.LinkRequest {
	return types.NewLinkRequest(
		rand.Bytes(sdk.AddrLen),
		rand.StrBetween(5, 100),
		rand.StrBetween(5, 100))
}

func randomMsgConfirmOutpoint() *types.ConfirmOutpointRequest {
	return types.NewConfirmOutpointRequest(rand.Bytes(sdk.AddrLen), randomOutpointInfo())
}

func randomMsgVoteConfirmOutpoint() *types.VoteConfirmOutpointRequest {
	return types.NewVoteConfirmOutpointRequest(
		rand.Bytes(sdk.AddrLen),
		vote.PollKey{
			Module: types.ModuleName,
			ID:     rand.StrBetween(5, 20),
		},
		randomOutpointInfo().GetOutPoint(),
		rand.Bools(0.5).Next(),
	)
}

func randomOutpointInfo() types.OutPointInfo {
	txHash, err := chainhash.NewHash(rand.Bytes(chainhash.HashSize))
	if err != nil {
		panic(err)
	}
	vout := mathRand.Uint32()
	if vout == 0 {
		vout++
	}
	minOutputAmount, err := types.ToSatoshiCoin(types.DefaultParams().MinOutputAmount)
	if err != nil {
		panic(err)
	}

	return types.OutPointInfo{
		OutPoint: wire.NewOutPoint(txHash, vout).String(),
		Amount:   btcutil.Amount(rand.I64Between(minOutputAmount.Amount.Int64(), 10000000000)),
		Address:  randomAddress().EncodeAddress(),
	}
}

func randomCrossChainTransfer(maxAmount int64) nexus.CrossChainTransfer {
	minOutputAmount, err := types.ToSatoshiCoin(types.DefaultParams().MinOutputAmount)
	if err != nil {
		panic(err)
	}

	asset := types.DefaultParams().MinOutputAmount
	asset.Amount = asset.Amount.Add(sdk.NewDec(rand.I64Between(minOutputAmount.Amount.Int64(), maxAmount)))

	return nexus.CrossChainTransfer{
		ID:    uint64(rand.PosI64()),
		Asset: sdk.NewCoin(asset.Denom, asset.Amount.TruncateInt()),
		Recipient: nexus.CrossChainAddress{
			Chain:   exported.Bitcoin,
			Address: types.NewSecondaryConsolidationAddress(createRandomKey(tss.SecondaryKey), types.DefaultParams().Network).Address,
		},
	}
}

func randomAddress() *btcutil.AddressWitnessScriptHash {
	addr, err := btcutil.NewAddressWitnessScriptHash(rand.Bytes(32), types.DefaultParams().Network.Params())
	if err != nil {
		panic(err)
	}
	return addr
}
