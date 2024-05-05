package cracklepop

import "fmt"

// ErrNegativeInput occurs when the starting or ending number is negative.
type ErrNegativeInput int

func (e ErrNegativeInput) Error() string {
	return fmt.Sprintf("starting and ending numbers must greater than or equal 0: %d", int(e))
}

// ErrInvalidInput occurs when the ending number is greater than the
// starting number.
type ErrInvalidInput struct {
	start int
	end   int
}

func (e ErrInvalidInput) Error() string {
	return fmt.Sprintf("ending number must be greater than or equal starting number: %d %d", e.start, e.end)
}
