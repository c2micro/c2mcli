package object

// представляет собой элементарный объект для обработки
type Object interface {
	// возращение имени типа в виде строки
	TypeName() string
	// возвращение строковое представление типа
	String() string
	// выполнение бинарной операции
	BinaryOp(int, Object) (Object, error)
	// выполнение унарной операции
	UnaryOp(int) (Object, error)
	// может ли данный объект быть вызван как функция
	CanCall() bool
	// вызов объекта как функции с аргументами
	Call(...Object) (Object, error)
	// call method of object
	MethodCall(string, ...Object) (Object, error)
	// может ли данный объект итерироваться
	CanIterate() bool
	// возвращает итератор для объекта
	Iterate() Iterator
	// получение значения из объекта
	GetValue() any
	// получение значения объекта по его индексу внутри
	IndexGet(Object) (Object, error)
	// установка значения объекта по индексу
	IndexSet(Object, Object) error
	// LogicalNot реализация логического унарного НЕ
	LogicalNot() (Object, error)
	// LogicalOr реализация логического ИЛИ
	LogicalOr(Object) (Object, error)
	// LogicalAnd реализация логического И для объекта
	LogicalAnd(Object) (Object, error)
	// Equal реализация сравнения между объектами
	Equal(Object) (Object, error)
	// NotEqual реализация отрицательного сравнения между объектами
	NotEqual(Object) (Object, error)
	// Negative реализация унарного отрицания
	Negative() (Object, error)
	// GtEq реализация сравнения больше либо равно
	GtEq(Object) (Object, error)
	// Gt реализация сравнения строго больше
	Gt(Object) (Object, error)
	// LtEq реализация сравнения меньше либо равно
	LtEq(Object) (Object, error)
	// Lt реализация сравнения строго меньше
	Lt(Object) (Object, error)
	// Add сложение двух объектов
	Add(Object) (Object, error)
	// Sub вычитание двух объектов
	Sub(Object) (Object, error)
	// Pow возведение в степень
	Pow(Object) (Object, error)
	// Mul умножение объектов
	Mul(Object) (Object, error)
	// Div деление объектов
	Div(Object) (Object, error)
	// Mod получение целочисленного остатка
	Mod(Object) (Object, error)
	// Xor ксоринг объектов
	Xor(Object) (Object, error)
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

func (o *Impl) MethodCall(string, ...Object) (Object, error) {
	return nil, ErrUnknownMethod
}

func (a *Impl) LogicalNot() (Object, error) {
	return nil, ErrNotImplemented
}

func (a *Impl) LogicalAnd(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *Impl) LogicalOr(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *Impl) Equal(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *Impl) NotEqual(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *Impl) Negative() (Object, error) {
	return nil, ErrNotImplemented
}

func (a *Impl) GtEq(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *Impl) Gt(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *Impl) LtEq(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *Impl) Lt(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *Impl) Add(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *Impl) Sub(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *Impl) Pow(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *Impl) Mul(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *Impl) Div(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *Impl) Mod(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *Impl) Xor(Object) (Object, error) {
	return nil, ErrNotImplemented
}
