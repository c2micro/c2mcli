package object

type Iterator interface {
	Object
	// Next возвращает true если есть еще элементы для обработки
	Next() bool
	// Key возвращает ключ или индекс текущего элемента
	Key() Object
	// Value возвращает значение текущего элемента
	Value() Object
}

// ListIterator используется для итерирования по списку
type ListIterator struct {
	Impl
	// v значения в виде списка
	v []Object
	// i текущий индекс
	i int
	// l длина списка
	l int
}

func (i *ListIterator) TypeName() string {
	return "list-iter"
}

func (i *ListIterator) String() string {
	return "<list-iter>"
}

func (i *ListIterator) Next() bool {
	i.i++
	return i.i <= i.l
}

func (i *ListIterator) Key() Object {
	return &Int{value: int64(i.i - 1)}
}

func (i *ListIterator) Value() Object {
	return i.v[i.i-1]
}

type MapIterator struct {
	Impl
	// v значение мапы
	v map[string]Object
	// k ключи мапы
	k []string
	// i индекс
	i int
	// l длина мапы
	l int
}

func (m *MapIterator) TypeName() string {
	return "map-iter"
}

func (m *MapIterator) String() string {
	return "<map-iter>"
}

func (m *MapIterator) Next() bool {
	m.i++
	return m.i <= m.l
}

func (m *MapIterator) Key() Object {
	return NewStr(m.k[m.i-1])
}

func (m *MapIterator) Value() Object {
	return m.v[m.k[m.i-1]]
}
