package debugtools

type TracableError struct {
	err   error
	stack *Stack
}

func NewTE(err error) *TracableError {
	return &TracableError{
		err:   err,
		stack: trace(3),
	}
}
