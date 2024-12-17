package object

type Arithmetic interface {
	// LogicalNot реализация логического унарного НЕ
	LogicalNot() (Object, error)
	// LogicalOr реализация логического ИЛИ
	LogicalOr() (Object, error)
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

type ArithmeticImpl struct{}

func (a *ArithmeticImpl) LogicalNot() (Object, error) {
	return nil, ErrNotImplemented
}

func (a *ArithmeticImpl) LogicalAnd(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *ArithmeticImpl) LogicalOr(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *ArithmeticImpl) Equal(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *ArithmeticImpl) NotEqual(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *ArithmeticImpl) Negative() (Object, error) {
	return nil, ErrNotImplemented
}

func (a *ArithmeticImpl) GtEq(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *ArithmeticImpl) Gt(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *ArithmeticImpl) LtEq(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *ArithmeticImpl) Lt(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *ArithmeticImpl) Add(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *ArithmeticImpl) Sub(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *ArithmeticImpl) Pow(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *ArithmeticImpl) Mul(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *ArithmeticImpl) Div(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *ArithmeticImpl) Mod(Object) (Object, error) {
	return nil, ErrNotImplemented
}

func (a *ArithmeticImpl) Xor(Object) (Object, error) {
	return nil, ErrNotImplemented
}
