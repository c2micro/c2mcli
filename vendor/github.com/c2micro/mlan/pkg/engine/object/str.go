package object

import (
	"fmt"
	"strings"

	"github.com/c2micro/mlan/pkg/parser"
)

// Str тип с плавающей точкой
type Str struct {
	Impl
	Arithmetic
	value string
}

func NewStr(v string) *Str {
	return &Str{value: v}
}

func (o *Str) TypeName() string {
	return "str"
}

func (o *Str) String() string {
	return o.value
}

func (o *Str) IndexGet(rs Object) (Object, error) {
	idx, ok := rs.(*Int)
	if !ok {
		return nil, ErrInvalidIndexType
	}
	if idx.value < 0 || idx.value >= int64(len(o.value)) {
		return nil, ErrIndexOutOfRange
	}
	return NewStr(string(o.value[idx.value])), nil
}

func (o *Str) IndexSet(idx Object, rs Object) error {
	idxInt, ok := idx.(*Int)
	if !ok {
		return ErrInvalidIndexType
	}
	rsStr, ok := rs.(*Str)
	if !ok {
		return fmt.Errorf("invalid type '%s' of assignment to str", rs.TypeName())
	}
	if idxInt.value < 0 || int(idxInt.value) >= len(o.value) {
		return ErrIndexOutOfRange
	}
	if len(rsStr.value) != 1 {
		return fmt.Errorf("expected str with length of 1, got %d", len(rsStr.value))
	}
	o.value = o.value[:idxInt.value] + rsStr.value + o.value[idxInt.value+1:]
	return nil
}

func (o *Str) BinaryOp(op int, rhs Object) (Object, error) {
	switch op {
	case parser.MlanLexerOr:
		return o.LogicalOr(rhs)
	case parser.MlanLexerAnd:
		return o.LogicalAnd(rhs)
	case parser.MlanLexerEq:
		return o.Equal(rhs)
	case parser.MlanLexerNeq:
		return o.NotEqual(rhs)
	case parser.MlanLexerGtEq:
		return o.GtEq(rhs)
	case parser.MlanLexerGt:
		return o.Gt(rhs)
	case parser.MlanLexerLtEq:
		return o.LtEq(rhs)
	case parser.MlanLexerLt:
		return o.Lt(rhs)
	case parser.MlanLexerAdd:
		return o.Add(rhs)
	case parser.MlanLexerAssignSum:
		return o.Add(rhs)
	case parser.MlanLexerMultiply:
		return o.Mul(rhs)
	}
	return nil, ErrInvalidOp
}

func (o *Str) UnaryOp(op int) (Object, error) {
	switch op {
	case parser.MlanLexerNot:
		return o.LogicalNot()
	}
	return nil, ErrInvalidOp
}

func (o *Str) GetValue() any {
	return o.value
}

func (o *Str) LogicalNot() (Object, error) {
	if len(o.value) == 0 {
		return NewBool(true), nil
	}
	return NewBool(false), nil
}

func (o *Str) LogicalOr(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return o, nil
	}
	return nil, ErrInvalidOp
}

func (o *Str) LogicalAnd(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return rs, nil
	}
	return nil, ErrInvalidOp
}

func (o *Str) Equal(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(false), nil
	case *Dict:
		return NewBool(false), nil
	case *Float:
		return NewBool(false), nil
	case *Int:
		return NewBool(false), nil
	case *List:
		return NewBool(false), nil
	case *Null:
		return NewBool(false), nil
	case *Str:
		return NewBool(o.value == rs.(*Str).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Str) NotEqual(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(true), nil
	case *Dict:
		return NewBool(true), nil
	case *Float:
		return NewBool(true), nil
	case *Int:
		return NewBool(true), nil
	case *List:
		return NewBool(true), nil
	case *Null:
		return NewBool(true), nil
	case *Str:
		return NewBool(o.value != rs.(*Str).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Str) GtEq(rs Object) (Object, error) {
	switch rs.(type) {
	case *Str:
		return NewBool(o.value >= rs.(*Str).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Str) Gt(rs Object) (Object, error) {
	switch rs.(type) {
	case *Str:
		return NewBool(o.value > rs.(*Str).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Str) LtEq(rs Object) (Object, error) {
	switch rs.(type) {
	case *Str:
		return NewBool(o.value <= rs.(*Str).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Str) Lt(rs Object) (Object, error) {
	switch rs.(type) {
	case *Str:
		return NewBool(o.value < rs.(*Str).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Str) Add(rs Object) (Object, error) {
	switch rs.(type) {
	case *Str:
		o.value += rs.(*Str).value
		return o, nil
	}
	return nil, ErrInvalidOp
}

func (o *Str) Mul(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewStr(strings.Repeat(o.value, int(boolToInt(rs.(*Bool).value)))), nil
	case *Int:
		return NewStr(strings.Repeat(o.value, int(rs.(*Int).value))), nil
	}
	return nil, ErrInvalidOp
}
