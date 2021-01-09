// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	ethereumtypes "github.com/axelarnetwork/axelar-core/x/ethereum/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"sync"
)

// Ensure, that RPCClientMock does implement ethereumtypes.RPCClient.
// If this is not the case, regenerate this file with moq.
var _ ethereumtypes.RPCClient = &RPCClientMock{}

// RPCClientMock is a mock implementation of ethereumtypes.RPCClient.
//
//     func TestSomethingThatUsesRPCClient(t *testing.T) {
//
//         // make and configure a mocked ethereumtypes.RPCClient
//         mockedRPCClient := &RPCClientMock{
//             BlockNumberFunc: func(ctx context.Context) (uint64, error) {
// 	               panic("mock out the BlockNumber method")
//             },
//             EstimateGasFunc: func(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
// 	               panic("mock out the EstimateGas method")
//             },
//             NetworkIDFunc: func(ctx context.Context) (*big.Int, error) {
// 	               panic("mock out the NetworkID method")
//             },
//             PendingNonceAtFunc: func(ctx context.Context, account common.Address) (uint64, error) {
// 	               panic("mock out the PendingNonceAt method")
//             },
//             SendTransactionFunc: func(ctx context.Context, tx *coretypes.Transaction) error {
// 	               panic("mock out the SendTransaction method")
//             },
//             SuggestGasPriceFunc: func(ctx context.Context) (*big.Int, error) {
// 	               panic("mock out the SuggestGasPrice method")
//             },
//             TransactionByHashFunc: func(ctx context.Context, hash common.Hash) (*coretypes.Transaction, bool, error) {
// 	               panic("mock out the TransactionByHash method")
//             },
//             TransactionReceiptFunc: func(ctx context.Context, txHash common.Hash) (*coretypes.Receipt, error) {
// 	               panic("mock out the TransactionReceipt method")
//             },
//         }
//
//         // use mockedRPCClient in code that requires ethereumtypes.RPCClient
//         // and then make assertions.
//
//     }
type RPCClientMock struct {
	// BlockNumberFunc mocks the BlockNumber method.
	BlockNumberFunc func(ctx context.Context) (uint64, error)

	// EstimateGasFunc mocks the EstimateGas method.
	EstimateGasFunc func(ctx context.Context, msg ethereum.CallMsg) (uint64, error)

	// NetworkIDFunc mocks the NetworkID method.
	NetworkIDFunc func(ctx context.Context) (*big.Int, error)

	// PendingNonceAtFunc mocks the PendingNonceAt method.
	PendingNonceAtFunc func(ctx context.Context, account common.Address) (uint64, error)

	// SendTransactionFunc mocks the SendTransaction method.
	SendTransactionFunc func(ctx context.Context, tx *coretypes.Transaction) error

	// SuggestGasPriceFunc mocks the SuggestGasPrice method.
	SuggestGasPriceFunc func(ctx context.Context) (*big.Int, error)

	// TransactionByHashFunc mocks the TransactionByHash method.
	TransactionByHashFunc func(ctx context.Context, hash common.Hash) (*coretypes.Transaction, bool, error)

	// TransactionReceiptFunc mocks the TransactionReceipt method.
	TransactionReceiptFunc func(ctx context.Context, txHash common.Hash) (*coretypes.Receipt, error)

	// calls tracks calls to the methods.
	calls struct {
		// BlockNumber holds details about calls to the BlockNumber method.
		BlockNumber []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// EstimateGas holds details about calls to the EstimateGas method.
		EstimateGas []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Msg is the msg argument value.
			Msg ethereum.CallMsg
		}
		// NetworkID holds details about calls to the NetworkID method.
		NetworkID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// PendingNonceAt holds details about calls to the PendingNonceAt method.
		PendingNonceAt []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Account is the account argument value.
			Account common.Address
		}
		// SendTransaction holds details about calls to the SendTransaction method.
		SendTransaction []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tx is the tx argument value.
			Tx *coretypes.Transaction
		}
		// SuggestGasPrice holds details about calls to the SuggestGasPrice method.
		SuggestGasPrice []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// TransactionByHash holds details about calls to the TransactionByHash method.
		TransactionByHash []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Hash is the hash argument value.
			Hash common.Hash
		}
		// TransactionReceipt holds details about calls to the TransactionReceipt method.
		TransactionReceipt []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// TxHash is the txHash argument value.
			TxHash common.Hash
		}
	}
	lockBlockNumber        sync.RWMutex
	lockEstimateGas        sync.RWMutex
	lockNetworkID          sync.RWMutex
	lockPendingNonceAt     sync.RWMutex
	lockSendTransaction    sync.RWMutex
	lockSuggestGasPrice    sync.RWMutex
	lockTransactionByHash  sync.RWMutex
	lockTransactionReceipt sync.RWMutex
}

// BlockNumber calls BlockNumberFunc.
func (mock *RPCClientMock) BlockNumber(ctx context.Context) (uint64, error) {
	if mock.BlockNumberFunc == nil {
		panic("RPCClientMock.BlockNumberFunc: method is nil but RPCClient.BlockNumber was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockBlockNumber.Lock()
	mock.calls.BlockNumber = append(mock.calls.BlockNumber, callInfo)
	mock.lockBlockNumber.Unlock()
	return mock.BlockNumberFunc(ctx)
}

// BlockNumberCalls gets all the calls that were made to BlockNumber.
// Check the length with:
//     len(mockedRPCClient.BlockNumberCalls())
func (mock *RPCClientMock) BlockNumberCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockBlockNumber.RLock()
	calls = mock.calls.BlockNumber
	mock.lockBlockNumber.RUnlock()
	return calls
}

// EstimateGas calls EstimateGasFunc.
func (mock *RPCClientMock) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	if mock.EstimateGasFunc == nil {
		panic("RPCClientMock.EstimateGasFunc: method is nil but RPCClient.EstimateGas was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Msg ethereum.CallMsg
	}{
		Ctx: ctx,
		Msg: msg,
	}
	mock.lockEstimateGas.Lock()
	mock.calls.EstimateGas = append(mock.calls.EstimateGas, callInfo)
	mock.lockEstimateGas.Unlock()
	return mock.EstimateGasFunc(ctx, msg)
}

// EstimateGasCalls gets all the calls that were made to EstimateGas.
// Check the length with:
//     len(mockedRPCClient.EstimateGasCalls())
func (mock *RPCClientMock) EstimateGasCalls() []struct {
	Ctx context.Context
	Msg ethereum.CallMsg
} {
	var calls []struct {
		Ctx context.Context
		Msg ethereum.CallMsg
	}
	mock.lockEstimateGas.RLock()
	calls = mock.calls.EstimateGas
	mock.lockEstimateGas.RUnlock()
	return calls
}

// NetworkID calls NetworkIDFunc.
func (mock *RPCClientMock) NetworkID(ctx context.Context) (*big.Int, error) {
	if mock.NetworkIDFunc == nil {
		panic("RPCClientMock.NetworkIDFunc: method is nil but RPCClient.NetworkID was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockNetworkID.Lock()
	mock.calls.NetworkID = append(mock.calls.NetworkID, callInfo)
	mock.lockNetworkID.Unlock()
	return mock.NetworkIDFunc(ctx)
}

// NetworkIDCalls gets all the calls that were made to NetworkID.
// Check the length with:
//     len(mockedRPCClient.NetworkIDCalls())
func (mock *RPCClientMock) NetworkIDCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockNetworkID.RLock()
	calls = mock.calls.NetworkID
	mock.lockNetworkID.RUnlock()
	return calls
}

// PendingNonceAt calls PendingNonceAtFunc.
func (mock *RPCClientMock) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	if mock.PendingNonceAtFunc == nil {
		panic("RPCClientMock.PendingNonceAtFunc: method is nil but RPCClient.PendingNonceAt was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Account common.Address
	}{
		Ctx:     ctx,
		Account: account,
	}
	mock.lockPendingNonceAt.Lock()
	mock.calls.PendingNonceAt = append(mock.calls.PendingNonceAt, callInfo)
	mock.lockPendingNonceAt.Unlock()
	return mock.PendingNonceAtFunc(ctx, account)
}

// PendingNonceAtCalls gets all the calls that were made to PendingNonceAt.
// Check the length with:
//     len(mockedRPCClient.PendingNonceAtCalls())
func (mock *RPCClientMock) PendingNonceAtCalls() []struct {
	Ctx     context.Context
	Account common.Address
} {
	var calls []struct {
		Ctx     context.Context
		Account common.Address
	}
	mock.lockPendingNonceAt.RLock()
	calls = mock.calls.PendingNonceAt
	mock.lockPendingNonceAt.RUnlock()
	return calls
}

// SendTransaction calls SendTransactionFunc.
func (mock *RPCClientMock) SendTransaction(ctx context.Context, tx *coretypes.Transaction) error {
	if mock.SendTransactionFunc == nil {
		panic("RPCClientMock.SendTransactionFunc: method is nil but RPCClient.SendTransaction was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Tx  *coretypes.Transaction
	}{
		Ctx: ctx,
		Tx:  tx,
	}
	mock.lockSendTransaction.Lock()
	mock.calls.SendTransaction = append(mock.calls.SendTransaction, callInfo)
	mock.lockSendTransaction.Unlock()
	return mock.SendTransactionFunc(ctx, tx)
}

// SendTransactionCalls gets all the calls that were made to SendTransaction.
// Check the length with:
//     len(mockedRPCClient.SendTransactionCalls())
func (mock *RPCClientMock) SendTransactionCalls() []struct {
	Ctx context.Context
	Tx  *coretypes.Transaction
} {
	var calls []struct {
		Ctx context.Context
		Tx  *coretypes.Transaction
	}
	mock.lockSendTransaction.RLock()
	calls = mock.calls.SendTransaction
	mock.lockSendTransaction.RUnlock()
	return calls
}

// SuggestGasPrice calls SuggestGasPriceFunc.
func (mock *RPCClientMock) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	if mock.SuggestGasPriceFunc == nil {
		panic("RPCClientMock.SuggestGasPriceFunc: method is nil but RPCClient.SuggestGasPrice was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockSuggestGasPrice.Lock()
	mock.calls.SuggestGasPrice = append(mock.calls.SuggestGasPrice, callInfo)
	mock.lockSuggestGasPrice.Unlock()
	return mock.SuggestGasPriceFunc(ctx)
}

// SuggestGasPriceCalls gets all the calls that were made to SuggestGasPrice.
// Check the length with:
//     len(mockedRPCClient.SuggestGasPriceCalls())
func (mock *RPCClientMock) SuggestGasPriceCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockSuggestGasPrice.RLock()
	calls = mock.calls.SuggestGasPrice
	mock.lockSuggestGasPrice.RUnlock()
	return calls
}

// TransactionByHash calls TransactionByHashFunc.
func (mock *RPCClientMock) TransactionByHash(ctx context.Context, hash common.Hash) (*coretypes.Transaction, bool, error) {
	if mock.TransactionByHashFunc == nil {
		panic("RPCClientMock.TransactionByHashFunc: method is nil but RPCClient.TransactionByHash was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Hash common.Hash
	}{
		Ctx:  ctx,
		Hash: hash,
	}
	mock.lockTransactionByHash.Lock()
	mock.calls.TransactionByHash = append(mock.calls.TransactionByHash, callInfo)
	mock.lockTransactionByHash.Unlock()
	return mock.TransactionByHashFunc(ctx, hash)
}

// TransactionByHashCalls gets all the calls that were made to TransactionByHash.
// Check the length with:
//     len(mockedRPCClient.TransactionByHashCalls())
func (mock *RPCClientMock) TransactionByHashCalls() []struct {
	Ctx  context.Context
	Hash common.Hash
} {
	var calls []struct {
		Ctx  context.Context
		Hash common.Hash
	}
	mock.lockTransactionByHash.RLock()
	calls = mock.calls.TransactionByHash
	mock.lockTransactionByHash.RUnlock()
	return calls
}

// TransactionReceipt calls TransactionReceiptFunc.
func (mock *RPCClientMock) TransactionReceipt(ctx context.Context, txHash common.Hash) (*coretypes.Receipt, error) {
	if mock.TransactionReceiptFunc == nil {
		panic("RPCClientMock.TransactionReceiptFunc: method is nil but RPCClient.TransactionReceipt was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		TxHash common.Hash
	}{
		Ctx:    ctx,
		TxHash: txHash,
	}
	mock.lockTransactionReceipt.Lock()
	mock.calls.TransactionReceipt = append(mock.calls.TransactionReceipt, callInfo)
	mock.lockTransactionReceipt.Unlock()
	return mock.TransactionReceiptFunc(ctx, txHash)
}

// TransactionReceiptCalls gets all the calls that were made to TransactionReceipt.
// Check the length with:
//     len(mockedRPCClient.TransactionReceiptCalls())
func (mock *RPCClientMock) TransactionReceiptCalls() []struct {
	Ctx    context.Context
	TxHash common.Hash
} {
	var calls []struct {
		Ctx    context.Context
		TxHash common.Hash
	}
	mock.lockTransactionReceipt.RLock()
	calls = mock.calls.TransactionReceipt
	mock.lockTransactionReceipt.RUnlock()
	return calls
}