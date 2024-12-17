package object

// Object представляет собой элементарный объект для обработки
type Object interface {
	// TypeName возращение имени типа в виде строки
	TypeName() string
	// String возвращение строковое представление типа
	String() string
	// BinaryOp выполнение бинарной операции
	BinaryOp(int, Object) (Object, error)
	// UnaryOp выполнение унарной операции
	UnaryOp(int) (Object, error)
	// CanCall может ли данный объект быть вызван как функция
	CanCall() bool
	// Call вызов объекта как функции с аргументами
	Call(...Object) (Object, error)
	// CanIterate может ли данный объект итерироваться
	CanIterate() bool
	// Iterate возвращает итератор для объекта
	Iterate() Iterator
	// GetValue получение значения из объекта
	GetValue() any
	// IndexGet получение значения объекта по его индексу внутри
	IndexGet(Object) (Object, error)
	// IndexSet установка значения объекта по индексу
	IndexSet(Object, Object) error
}

// Impl базовая реализация объекта
type Impl struct{}

func (o *Impl) TypeName() string {
	panic(ErrNotImplemented)
}

func (o *Impl) String() string {
	panic(ErrNotImplemented)
}

func (o *Impl) BinaryOp(_ int, _ Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (o *Impl) UnaryOp(_ int) (Object, error) {
	return nil, ErrNotImplemented
}

func (o *Impl) CanCall() bool {
	return false
}

func (o *Impl) Call(_ ...Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (o *Impl) CanIterate() bool {
	return false
}

func (o *Impl) Iterate() Iterator {
	return nil
}

func (o *Impl) GetValue() any {
	return nil
}

func (o *Impl) IndexGet(Object) (Object, error) {
	return nil, ErrInvalidOp
}

func (o *Impl) IndexSet(Object, Object) error {
	return ErrInvalidOp
}
