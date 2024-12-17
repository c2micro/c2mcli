package object

// BuiltinFunc функция
type BuiltinFunc struct {
	Impl
	name string
	code func(args ...Object) (Object, error)
}

// NewBuiltinFunc новая встроенная функция
func NewBuiltinFunc(
	n string,
	c func(args ...Object) (Object, error),
) *BuiltinFunc {
	return &BuiltinFunc{
		name: n,
		code: c,
	}
}

func (o *BuiltinFunc) TypeName() string {
	return "builtin-func: " + o.name
}

func (o *BuiltinFunc) String() string {
	return "<builtin-func>"
}

func (o *BuiltinFunc) BinaryOp(_ int, _ Object) (Object, error) {
	return nil, ErrInvalidOp
}

func (o *BuiltinFunc) UnaryOp(_ int) (Object, error) {
	return nil, ErrInvalidOp
}

func (o *BuiltinFunc) CanCall() bool {
	return true
}

func (o *BuiltinFunc) Call(args ...Object) (Object, error) {
	return o.code(args...)
}
