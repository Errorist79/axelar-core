package types

// Event types
const (
	EventTypeNewChain                = "newChain"
	EventTypeChainConfirmation       = "chainConfirmation"
	EventTypeDepositConfirmation     = "depositConfirmation"
	EventTypeTokenConfirmation       = "tokenConfirmation"
	EventTypeTransferKeyConfirmation = "transferKeyConfirmation"
)

// Event attribute keys
const (
	AttributeKeyBatchedCommandsID = "batchedCommandID"
	AttributeKeyChain             = "chain"
	AttributeKeyAddress           = "address"
	AttributeKeyPoll              = "poll"
	AttributeKeyTxID              = "txID"
	AttributeKeyTransferKeyType   = "transferKeyType"
	AttributeKeyAmount            = "amount"
	AttributeKeyBurnAddress       = "burnAddress"
	AttributeKeyTokenAddress      = "tokenAddress"
	AttributeKeyGatewayAddress    = "gatewayAddress"
	AttributeKeyConfHeight        = "confHeight"
	AttributeKeyAsset             = "asset"
	AttributeKeySymbol            = "symbol"
	AttributeKeyNativeAsset       = "nativeAsset"
	AttributeKeyDestinationChain   = "destinationChain"
	AttributeKeyDestinationAddress = "destinationAddress"
)

// Event attribute values
const (
	AttributeValueUpdate  = "update"
	AttributeValueStart   = "start"
	AttributeValueReject  = "reject"
	AttributeValueConfirm = "confirm"
)
