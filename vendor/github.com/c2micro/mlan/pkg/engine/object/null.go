package object

import (
	"github.com/c2micro/mlan/pkg/parser"
)

// Null нулевой тип
type Null struct {
	Impl
	Arithmetic
}

// NewNull создание нового объекта
func NewNull() *Null {
	return &Null{}
}

func (o *Null) TypeName() string {
	return "null"
}

func (o *Null) String() string {
	return "<null>"
}

func (o *Null) BinaryOp(op int, rhs Object) (Object, error) {
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

func (o *Null) UnaryOp(op int) (Object, error) {
	switch op {
	case parser.MlanLexerNot:
		return o.LogicalNot()
	}
	return nil, ErrInvalidOp
}

func (o *Null) LogicalNot() (Object, error) {
	return NewBool(true), nil
}

func (o *Null) LogicalOr(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return rs, nil
	case *Dict:
		return rs, nil
	case *Float:
		return rs, nil
	case *Int:
		return rs, nil
	case *List:
		return rs, nil
	case *Null:
		return rs, nil
	case *Str:
		return rs, nil
	}
	return nil, ErrInvalidOp
}

func (o *Null) LogicalAnd(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return o, nil
	case *Dict:
		return o, nil
	case *Float:
		return o, nil
	case *Int:
		return o, nil
	case *List:
		return o, nil
	case *Null:
		return o, nil
	case *Str:
		return o, nil
	}
	return nil, ErrInvalidOp
}

func (o *Null) Equal(rs Object) (Object, error) {
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
		return NewBool(true), nil
	case *Str:
		return NewBool(false), nil
	}
	return nil, ErrInvalidOp
}

func (o *Null) NotEqual(rs Object) (Object, error) {
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
		return NewBool(false), nil
	case *Str:
		return NewBool(true), nil
	}
	return nil, ErrInvalidOp
}
