package v1

import "math"

func (u EpochUsage) Validate() error {
	if u.Epoch > math.MaxInt64 {
		return ErrInvalidClock.Wrap("epoch out of range")
	}

	if err := ValidateAmount(u.Amount); err != nil {
		return err
	}

	if u.Epoch == 0 && u.Amount.IsPositive() {
		return ErrInvalidClock.Wrap("usage before epoch counting started")
	}

	return nil
}

func (u AccountUsage) Validate() error {
	if u.LastEpoch < 0 {
		return ErrInvalidClock.Wrap("negative account epoch")
	}

	if err := ValidateAmount(u.TotalSent); err != nil {
		return err
	}

	if err := ValidateAmount(u.EpochSent); err != nil {
		return err
	}

	if u.LastEpoch == 0 && u.EpochSent.IsPositive() {
		return ErrInvalidClock.Wrap("account usage before epoch counting started")
	}

	if u.EpochSent.IsGT(u.TotalSent) {
		return ErrInvalidGenesis.Wrap("account epoch usage exceeds lifetime amount")
	}

	return nil
}

func (u EpochUsage) At(epoch uint64) (EpochUsage, error) {
	if err := u.Validate(); err != nil {
		return EpochUsage{}, err
	}

	if epoch > math.MaxInt64 {
		return EpochUsage{}, ErrInvalidClock.Wrap("epoch out of range")
	}

	if u.Epoch > epoch {
		return EpochUsage{}, ErrInvalidClock.Wrap("future epoch usage")
	}

	if u.Epoch < epoch {
		return EpochUsage{
			Epoch:  epoch,
			Amount: ZeroAmount(),
		}, nil
	}

	return u, nil
}

func (u AccountUsage) At(epoch uint64) (AccountUsage, error) {
	if err := u.Validate(); err != nil {
		return AccountUsage{}, err
	}

	if epoch > math.MaxInt64 {
		return AccountUsage{}, ErrInvalidClock.Wrap("epoch out of range")
	}

	current := int64(epoch)
	if u.LastEpoch > current {
		return AccountUsage{}, ErrInvalidClock.Wrap("future account usage")
	}

	if u.LastEpoch < current {
		u.LastEpoch, u.EpochSent = current, ZeroAmount()
	}

	return u, nil
}
