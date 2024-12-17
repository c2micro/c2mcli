package object

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/c2micro/mlan/pkg/parser"
)

// Dict мапа
type Dict struct {
	Impl
	value map[string]Object
}

func NewDict(v map[string]Object) *Dict {
	return &Dict{value: v}
}

func (o *Dict) TypeName() string {
	return "dict"
}

func (o *Dict) String() string {
	var items []string
	for k, v := range o.value {
		items = append(items, fmt.Sprintf("%s: %s", k, v.String()))
	}
	return fmt.Sprintf("{%s}", strings.Join(items, ", "))
}

func (o *Dict) CanIterate() bool {
	return true
}

func (o *Dict) GetValue() any {
	return o.value
}

func (o *Dict) Iterate() Iterator {
	var keys []string
	for k := range o.value {
		keys = append(keys, k)
	}
	return &MapIterator{
		v: o.value,
		k: keys,
		i: 0,
		l: len(keys),
	}
}

func (o *Dict) BinaryOp(op int, rhs Object) (Object, error) {
	switch op {
	case parser.MlanLexerOr:
		return o.LogicalOr(rhs)
	case parser.MlanLexerAnd:
		return o.LogicalAnd(rhs)
	case parser.MlanLexerEq:
		return o.Equal(rhs)
	case parser.MlanLexerNeq:
		return o.NotEqual(rhs)
	}
	return nil, ErrInvalidOp
}

func (o *Dict) IndexGet(rs Object) (Object, error) {
	idx, ok := rs.(*Str)
	if !ok {
		return nil, ErrInvalidIndexType
	}
	res, ok := o.value[idx.value]
	if !ok {
		return NewNull(), nil
	} else {
		return res, nil
	}
}

func (o *Dict) IndexSet(idx Object, rs Object) error {
	idxStr, ok := rs.(*Str)
	if !ok {
		return ErrInvalidIndexType
	}
	o.value[idxStr.value] = rs
	return nil
}

func (o *Dict) UnaryOp(op int) (Object, error) {
	switch op {
	case parser.MlanLexerNot:
		return o.LogicalNot()
	}
	return nil, ErrInvalidOp
}

func (o *Dict) LogicalNot() (Object, error) {
	if len(o.value) == 0 {
		return NewBool(true), nil
	}
	return NewBool(false), nil
}

func (o *Dict) LogicalOr(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return o, nil
	}
	return nil, ErrInvalidOp
}

func (o *Dict) LogicalAnd(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return rs, nil
	}
	return nil, ErrInvalidOp
}

func (o *Dict) Equal(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(false), nil
	case *Dict:
		return NewBool(reflect.DeepEqual(o.value, rs.(*Dict).value)), nil
	case *Float:
		return NewBool(false), nil
	case *Int:
		return NewBool(false), nil
	case *List:
		return NewBool(false), nil
	case *Null:
		return NewBool(false), nil
	case *Str:
		return NewBool(false), nil
	}
	return nil, ErrInvalidOp
}

func (o *Dict) NotEqual(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(true), nil
	case *Dict:
		return NewBool(!reflect.DeepEqual(o.value, rs.(*Dict).value)), nil
	case *Float:
		return NewBool(true), nil
	case *Int:
		return NewBool(true), nil
	case *List:
		return NewBool(true), nil
	case *Null:
		return NewBool(true), nil
	case *Str:
		return NewBool(true), nil
	}
	return nil, ErrInvalidOp
}
