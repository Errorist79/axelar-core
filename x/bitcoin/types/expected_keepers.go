package types

import (
	"crypto/ecdsa"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/axelarnetwork/axelar-core/x/balance/exported"
	tss "github.com/axelarnetwork/axelar-core/x/tss/exported"
	voting "github.com/axelarnetwork/axelar-core/x/vote/exported"
)

//go:generate moq -pkg mock -out ./mock/expected_keepers.go . Voter Signer Balancer

type Voter interface {
	voting.Voter
}

type Signer interface {
	GetSig(ctx sdk.Context, sigID string) (tss.Signature, bool)
	GetKey(ctx sdk.Context, keyID string) (ecdsa.PublicKey, bool)
	GetCurrentMasterKey(ctx sdk.Context, chain exported.Chain) (ecdsa.PublicKey, bool)
	GetNextMasterKey(ctx sdk.Context, chain exported.Chain) (ecdsa.PublicKey, bool)
	GetKeyForSigID(ctx sdk.Context, sigID string) (ecdsa.PublicKey, bool)
}

type Balancer interface {
	LinkAddresses(ctx sdk.Context, sender exported.CrossChainAddress, recipient exported.CrossChainAddress)
	PrepareForTransfer(ctx sdk.Context, sender exported.CrossChainAddress, amount sdk.Coin) error
	GetRecipient(ctx sdk.Context, sender exported.CrossChainAddress) (exported.CrossChainAddress, bool)
}