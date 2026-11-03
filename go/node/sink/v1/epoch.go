package v1

import (
	"math"
	"time"
)

const EpochSeconds = 86400

func EpochOf(t time.Time) (uint64, error) {
	seconds := t.Unix()
	if seconds < 0 {
		return 0, ErrInvalidClock.Wrap("time before Unix epoch")
	}

	return uint64(seconds) / EpochSeconds, nil
}

func (u EpochUsage) Validate() error {
	return ValidateAmount(u.Amount)
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

	if u.EpochSent.IsGT(u.TotalSent) {
		return ErrInvalidGenesis.Wrap("account epoch usage exceeds lifetime amount")
	}

	return nil
}

func (u EpochUsage) At(epoch uint64) (EpochUsage, error) {
	if err := u.Validate(); err != nil {
		return EpochUsage{}, err
	}

	if u.Epoch > epoch {
		return EpochUsage{}, ErrInvalidClock.Wrap("future epoch usage")
	}

	if u.Epoch < epoch {
		return EpochUsage{Epoch: epoch, Amount: ZeroAmount()}, nil
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
