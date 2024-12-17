package object

import (
	"math"
	"strconv"
	"strings"

	"github.com/c2micro/mlan/pkg/parser"
)

// Int целочисленный тип
type Int struct {
	Impl
	Arithmetic
	value int64
}

// NewInt создание нового объекта с типом Int
func NewInt(value int64) *Int {
	return &Int{
		value: value,
	}
}

func (o *Int) TypeName() string {
	return "int"
}

func (o *Int) String() string {
	return strconv.FormatInt(o.value, 10)
}

func (o *Int) GetValue() any {
	return o.value
}

func (o *Int) BinaryOp(op int, rhs Object) (Object, error) {
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
	case parser.MlanLexerSubtract:
		return o.Sub(rhs)
	case parser.MlanLexerPow:
		return o.Pow(rhs)
	case parser.MlanLexerMultiply:
		return o.Mul(rhs)
	case parser.MlanLexerDivision:
		return o.Div(rhs)
	case parser.MlanLexerModulus:
		return o.Mod(rhs)
	case parser.MlanParserXor:
		return o.Xor(rhs)
	}
	return nil, ErrInvalidOp
}

func (o *Int) UnaryOp(op int) (Object, error) {
	switch op {
	case parser.MlanLexerNot:
		return o.LogicalNot()
	case parser.MlanLexerSubtract:
		return o.Negative()
	}
	return nil, ErrInvalidOp
}

func (o *Int) LogicalNot() (Object, error) {
	return NewBool(!intToBool(o.value)), nil
}

func (o *Int) Negative() (Object, error) {
	o.value = -o.value
	return o, nil
}

func (o *Int) LogicalOr(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return o, nil
	}
	return nil, ErrInvalidOp
}

func (o *Int) LogicalAnd(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return rs, nil
	}
	return nil, ErrInvalidOp
}

func (o *Int) Equal(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(o.value == boolToInt(rs.(*Bool).value)), nil
	case *Dict:
		return NewBool(false), nil
	case *Float:
		return NewBool(intToFloat(o.value) == rs.(*Float).value), nil
	case *Int:
		return NewBool(o.value == rs.(*Int).value), nil
	case *List:
		return NewBool(false), nil
	case *Null:
		return NewBool(false), nil
	case *Str:
		return NewBool(false), nil
	}
	return nil, ErrInvalidOp
}

func (o *Int) NotEqual(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(o.value != boolToInt(rs.(*Bool).value)), nil
	case *Dict:
		return NewBool(true), nil
	case *Float:
		return NewBool(intToFloat(o.value) != rs.(*Float).value), nil
	case *Int:
		return NewBool(o.value != rs.(*Int).value), nil
	case *List:
		return NewBool(true), nil
	case *Null:
		return NewBool(true), nil
	case *Str:
		return NewBool(true), nil
	}
	return nil, ErrInvalidOp
}

func (o *Int) GtEq(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(o.value >= boolToInt(rs.(*Bool).value)), nil
	case *Float:
		return NewBool(intToFloat(o.value) >= rs.(*Float).value), nil
	case *Int:
		return NewBool(o.value >= rs.(*Int).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Int) Gt(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(o.value > boolToInt(rs.(*Bool).value)), nil
	case *Float:
		return NewBool(intToFloat(o.value) > rs.(*Float).value), nil
	case *Int:
		return NewBool(o.value > rs.(*Int).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Int) LtEq(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(o.value <= boolToInt(rs.(*Bool).value)), nil
	case *Float:
		return NewBool(intToFloat(o.value) <= rs.(*Float).value), nil
	case *Int:
		return NewBool(o.value <= rs.(*Int).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Int) Lt(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(o.value < boolToInt(rs.(*Bool).value)), nil
	case *Float:
		return NewBool(intToFloat(o.value) < rs.(*Float).value), nil
	case *Int:
		return NewBool(o.value < rs.(*Int).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Int) Add(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewInt(o.value + boolToInt(rs.(*Bool).value)), nil
	case *Float:
		return NewFloat(intToFloat(o.value) + rs.(*Float).value), nil
	case *Int:
		return NewInt(o.value + rs.(*Int).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Int) Sub(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewInt(o.value - boolToInt(rs.(*Bool).value)), nil
	case *Float:
		return NewFloat(intToFloat(o.value) - rs.(*Float).value), nil
	case *Int:
		return NewInt(o.value - rs.(*Int).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Int) Pow(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewInt(int64(math.Pow(intToFloat(o.value), boolToFloat(rs.(*Bool).value)))), nil
	case *Float:
		return NewFloat(math.Pow(intToFloat(o.value), rs.(*Float).value)), nil
	case *Int:
		return NewInt(int64(math.Pow(intToFloat(o.value), intToFloat(rs.(*Int).value)))), nil
	}
	return nil, ErrInvalidOp
}

func (o *Int) Mul(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewInt(o.value * boolToInt(rs.(*Bool).value)), nil
	case *Float:
		return NewFloat(intToFloat(o.value) * rs.(*Float).value), nil
	case *Int:
		return NewInt(o.value * rs.(*Int).value), nil
	case *List:
		var list []Object
		var i int64
		for i = 0; i < o.value; i++ {
			list = append(list, rs.(*List).value...)
		}
		return NewList(list), nil
	case *Str:
		return NewStr(strings.Repeat(rs.(*Str).value, int(o.value))), nil
	}
	return nil, ErrInvalidOp
}

func (o *Int) Div(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		if !rs.(*Bool).value {
			return nil, ErrDivByZero
		}
		return NewFloat(intToFloat(o.value) / boolToFloat(rs.(*Bool).value)), nil
	case *Float:
		if rs.(*Float).value == 0.0 {
			return nil, ErrDivByZero
		}
		return NewFloat(intToFloat(o.value) / rs.(*Float).value), nil
	case *Int:
		if rs.(*Int).value == 0 {
			return nil, ErrDivByZero
		}
		return NewInt(o.value / rs.(*Int).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Int) Mod(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		if !rs.(*Bool).value {
			return nil, ErrModByZero
		}
		return NewInt(o.value % boolToInt(rs.(*Bool).value)), nil
	case *Float:
		if rs.(*Float).value == 0.0 {
			return nil, ErrModByZero
		}
		return NewFloat(math.Mod(intToFloat(o.value), rs.(*Float).value)), nil
	case *Int:
		if rs.(*Int).value == 0 {
			return nil, ErrModByZero
		}
		return NewInt(o.value % rs.(*Int).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Int) Xor(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewInt(o.value ^ boolToInt(rs.(*Bool).value)), nil
	case *Int:
		return NewInt(o.value ^ rs.(*Int).value), nil
	}
	return nil, ErrInvalidOp
}
