package v1

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.HasValidateBasic = &MsgSink{}
	_ sdk.HasValidateBasic = &MsgUpdateParams{}
	_ sdk.HasValidateBasic = &MsgRequestVoid{}
	_ sdk.HasValidateBasic = &MsgReportOutcome{}
	_ sdk.HasValidateBasic = &MsgVetoVoid{}
	_ sdk.HasValidateBasic = &MsgRestartTransfer{}
	_ sdk.HasValidateBasic = &MsgConfirmMint{}
	_ sdk.HasValidateBasic = &MsgConfirmVoid{}
	_ sdk.HasValidateBasic = &MsgFinalizeRefund{}
	_ sdk.HasValidateBasic = &MsgGuardianPause{}
)

type Resolution interface {
	sdk.Msg
	sdk.HasValidateBasic
	resolution()
}

func (*MsgRestartTransfer) resolution() {}

func (*MsgConfirmMint) resolution() {}

func (*MsgConfirmVoid) resolution() {}

func (m *MsgSink) ValidateBasic() error {
	if m.DestType != DestTypeSolana {
		return ErrInvalidDestination.Wrapf("destination type must be %s", DestTypeSolana)
	}

	if _, err := ParseAccount(m.Sender, AccountLength); err != nil {
		return err
	}

	return (RecordData{Dest: m.Dest, Amount: m.Amount}).ValidateBasic()
}

func (m *MsgUpdateParams) ValidateBasic() error {
	if _, err := ParseAccount(m.Authority, AccountLength, DerivedAccountLength); err != nil {
		return err
	}

	return m.Params.ValidateBasic()
}

func (m *MsgRequestVoid) ValidateBasic() error {
	if err := validateSignerID(m.Authority, m.ID); err != nil {
		return err
	}

	if m.ExpectedRound == 0 {
		return ErrInvalidResolution.Wrap("expected_round must be positive")
	}

	return nil
}

func (m *MsgFinalizeRefund) ValidateBasic() error {
	return validateSignerID(m.Executor, m.ID)
}

func (m *MsgVetoVoid) ValidateBasic() error {
	if err := validateSignerID(m.Signer, m.ID); err != nil {
		return err
	}

	return ValidateReason(m.Reason)
}

func (m *MsgReportOutcome) ValidateBasic() error {
	if _, err := ParseAccount(m.Reporter, AccountLength); err != nil {
		return err
	}

	if err := m.ID.ValidateBasic(); err != nil {
		return err
	}

	if m.PolicyVersion == 0 || m.Round == 0 {
		return ErrStaleReport.Wrap("policy_version and round must be positive")
	}

	return m.Observation.Validate()
}

func (m *MsgRestartTransfer) ValidateBasic() error {
	return validateResolution(m.Authority, m.ID, m.Reason, m.ExpectedRound)
}

func (m *MsgConfirmMint) ValidateBasic() error {
	if err := validateResolution(m.Authority, m.ID, m.Reason, m.ExpectedRound); err != nil {
		return err
	}

	return m.Evidence.Validate()
}

func (m *MsgConfirmVoid) ValidateBasic() error {
	if err := validateResolution(m.Authority, m.ID, m.Reason, m.ExpectedRound); err != nil {
		return err
	}

	return m.Evidence.Validate()
}

func validateSignerID(signer string, id ID) error {
	if _, err := ParseAccount(signer, AccountLength, DerivedAccountLength); err != nil {
		return err
	}

	return id.ValidateBasic()
}

func validateResolution(authority string, id ID, reason string, round uint64) error {
	if err := validateSignerID(authority, id); err != nil {
		return err
	}

	if round == 0 {
		return ErrInvalidResolution.Wrap("expected_round must be positive")
	}

	if round == math.MaxUint64 {
		return ErrOverflow.Wrap("settlement round")
	}

	return ValidateReason(reason)
}

func (m *MsgGuardianPause) ValidateBasic() error {
	if err := ValidatePauseGuardian(m.Guardian); err != nil {
		return err
	}

	if !m.Deposits && !m.Refunds {
		return ErrInvalidGuardianPause
	}

	return nil
}
