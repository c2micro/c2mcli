package object

// UserFunc функция, созданная пользователем
type UserFunc struct {
	Impl
	name string
	code func(args ...Object) (Object, error)
}

// NewUserFunc новая пользовательская функция
func NewUserFunc(
	n string,
	c func(args ...Object) (Object, error),
) *UserFunc {
	return &UserFunc{
		name: n,
		code: c,
	}
}

func (o *UserFunc) TypeName() string {
	return "user-func: " + o.name
}

func (o *UserFunc) String() string {
	return "<user-func>"
}

func (o *UserFunc) BinaryOp(_ int, _ Object) (Object, error) {
	return nil, ErrInvalidOp
}

func (o *UserFunc) UnaryOp(_ int) (Object, error) {
	return nil, ErrInvalidOp
}

func (o *UserFunc) CanCall() bool {
	return true
}

func (o *UserFunc) Call(args ...Object) (Object, error) {
	return o.code(args...)
}
