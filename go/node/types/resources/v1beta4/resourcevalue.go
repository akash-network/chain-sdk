package v1beta4

import (
	"errors"

	sdkmath "cosmossdk.io/math"
)

var (
	ErrOverflow       = errors.New("resource value overflow")
	ErrCannotSub      = errors.New("cannot subtract resources when lhs does not have same units as rhs")
	ErrNegativeResult = errors.New("result of subtraction is negative")
)

/*
ResourceValue the big point of this small change is to ensure math operations on resources
not resulting with negative value which panic on unsigned types as well as overflow which leads to panic too
instead reasonable error is returned.
Each resource using this type as value can take extra advantage of it to check upper bounds
For example in SDL v1 CPU units were handled as uint32 and operation like math.MaxUint32 + 2
would cause application to panic. But nowadays
	const CPULimit = math.MaxUint32

	func (c *CPU) add(rhs CPU) error {
		res, err := c.Units.add(rhs.Units)
		if err != nil {
			return err
		}

		if res.Units.Value() > CPULimit {
			return ErrOverflow
		}

		c.Units = res

		return nil
	}
*/

func NewResourceValue(val uint64) ResourceValue {
	res := ResourceValue{
		Val: sdkmath.NewIntFromUint64(val),
	}

	return res
}

func (m ResourceValue) Value() uint64 {
	return m.Val.Uint64()
}

func (m ResourceValue) Dup() ResourceValue {
	res := ResourceValue{
		Val: sdkmath.NewIntFromBigInt(m.Val.BigInt()),
	}

	return res
}
