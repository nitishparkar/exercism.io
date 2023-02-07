package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError error

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, err := weightFodder.FodderAmount()

	if fodderAmount < 0 && (err == nil || err == ErrScaleMalfunction) {
		return 0.0, errors.New("negative fodder")
	}

	if err == ErrScaleMalfunction && fodderAmount > 0 {
		return fodderAmount * 2.0 / float64(cows), nil
	}

	if err != nil {
		return 0.0, err
	}

	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}

	if cows < 0 {
		return 0.0, fmt.Errorf("silly nephew, there cannot be %d cows", cows)
	}

	return fodderAmount / float64(cows), nil
}
