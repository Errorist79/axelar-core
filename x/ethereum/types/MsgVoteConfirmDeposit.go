package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Route returns the route for this message
func (m MsgVoteConfirmDeposit) Route() string {
	return RouterKey
}

// Type returns the type of the message
func (m MsgVoteConfirmDeposit) Type() string {
	return "VoteConfirmDeposit"
}

// ValidateBasic executes a stateless message validation
func (m MsgVoteConfirmDeposit) ValidateBasic() error {
	if m.Sender == nil || len(m.Sender) != sdk.AddrLen {
		return fmt.Errorf("missing sender")
	}
	if m.TxID == "" {
		return fmt.Errorf("tx ID missing")
	}
	if m.BurnAddr == "" {
		return fmt.Errorf("burn address missing")
	}
	return m.Poll.Validate()
}

// GetSignBytes returns the message bytes that need to be signed
func (m MsgVoteConfirmDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the set of signers for this message
func (m MsgVoteConfirmDeposit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
