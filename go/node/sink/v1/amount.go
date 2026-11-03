package v1

import (
	"math"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"pkg.akt.dev/go/sdkutil"
)

func ZeroAmount() sdk.Coin {
	return sdk.NewCoin(sdkutil.DenomUakt, sdkmath.ZeroInt())
}

// ValidateAmount accepts initialized nonnegative uakt Coins, including usage counters.
func ValidateAmount(coin sdk.Coin) error {
	if coin.Denom != sdkutil.DenomUakt || coin.Amount.IsNil() || coin.Amount.IsNegative() {
		return sdkerrors.ErrInvalidCoins.Wrap("expected an initialized nonnegative uakt Coin")
	}

	return nil
}

func AddAmount(a, b sdk.Coin) (sdk.Coin, error) {
	if err := ValidateAmount(a); err != nil {
		return sdk.Coin{}, err
	}

	if err := ValidateAmount(b); err != nil {
		return sdk.Coin{}, err
	}

	sum, err := a.Amount.SafeAdd(b.Amount)
	if err != nil {
		return sdk.Coin{}, ErrOverflow.Wrap("uakt amount")
	}

	return sdk.NewCoin(sdkutil.DenomUakt, sum), nil
}

func DefaultTotals() Totals {
	return Totals{PendingAmount: ZeroAmount(), MintedAmount: ZeroAmount(), RefundPendingAmount: ZeroAmount(), RefundedAmount: ZeroAmount()}
}

func (t Totals) Validate() error {
	for _, amount := range []sdk.Coin{t.PendingAmount, t.MintedAmount, t.RefundPendingAmount, t.RefundedAmount} {
		if err := ValidateAmount(amount); err != nil {
			return err
		}
	}

	return nil
}

func (t Totals) DepositedAmount() (sdk.Coin, error) {
	escrow, err := t.EscrowAmount()
	if err != nil {
		return sdk.Coin{}, err
	}

	return AddAmount(escrow, t.RefundedAmount)
}

func (t Totals) EscrowAmount() (sdk.Coin, error) {
	amount, err := AddAmount(t.PendingAmount, t.MintedAmount)
	if err != nil {
		return sdk.Coin{}, err
	}

	return AddAmount(amount, t.RefundPendingAmount)
}

func (t Totals) Equal(other Totals) bool {
	return t.PendingCount == other.PendingCount &&
		t.MintedCount == other.MintedCount &&
		t.RefundPendingCount == other.RefundPendingCount &&
		t.RefundedCount == other.RefundedCount &&
		t.PendingAmount.Equal(other.PendingAmount) &&
		t.MintedAmount.Equal(other.MintedAmount) &&
		t.RefundPendingAmount.Equal(other.RefundPendingAmount) &&
		t.RefundedAmount.Equal(other.RefundedAmount)
}

// Add adds one record to the population matching its financial status.
func (t *Totals) Add(status Status, amount sdk.Coin) error {
	sum, count, err := t.population(status)
	if err != nil {
		return err
	}

	if *count == math.MaxUint64 {
		return ErrOverflow.Wrap("status count")
	}

	next, err := AddAmount(*sum, amount)
	if err != nil {
		return err
	}

	*sum = next
	*count++

	return nil
}

// Sub removes one record from the population matching its financial status.
func (t *Totals) Sub(status Status, amount sdk.Coin) error {
	sum, count, err := t.population(status)
	if err != nil {
		return err
	}

	if err = ValidateAmount(*sum); err != nil {
		return err
	}

	if err = ValidateAmount(amount); err != nil {
		return err
	}

	if *count == 0 || sum.IsLT(amount) {
		return ErrOverflow.Wrap("status population underflow")
	}

	*sum = sum.Sub(amount)
	*count--

	return nil
}

func (t *Totals) population(status Status) (*sdk.Coin, *uint64, error) {
	switch status {
	case StatusTransferPending:
		return &t.PendingAmount, &t.PendingCount, nil
	case StatusMinted:
		return &t.MintedAmount, &t.MintedCount, nil
	case StatusRefundPending:
		return &t.RefundPendingAmount, &t.RefundPendingCount, nil
	case StatusRefunded:
		return &t.RefundedAmount, &t.RefundedCount, nil
	default:
		return nil, nil, ErrInvalidSettlement.Wrap("unknown financial status")
	}
}
