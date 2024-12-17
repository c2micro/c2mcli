package object

import (
	"errors"
)

var (
	// ErrInvalidOp невалидное использование оператора
	ErrInvalidOp = errors.New("invalid operator")

	// ErrNotImplemented ошибка, когда объект не реализует метод
	ErrNotImplemented = errors.New("not implemented")

	ErrDivByZero = errors.New("division by zero")

	ErrModByZero = errors.New("modulo by zero")

	ErrInvalidIndexType = errors.New("invalid index type")

	ErrIndexOutOfRange = errors.New("index out of range")
)
