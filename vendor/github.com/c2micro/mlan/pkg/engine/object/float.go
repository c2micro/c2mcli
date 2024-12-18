package object

import (
	"math"
	"strconv"

	"github.com/c2micro/mlan/pkg/parser"
	"github.com/c2micro/mlan/pkg/engine/utils"
)

// Float тип с плавающей точкой
type Float struct {
	Impl
	Arithmetic
	value float64
}

func NewFloat(value float64) *Float {
	return &Float{value: value}
}

func (o *Float) TypeName() string {
	return "float"
}

func (o *Float) String() string {
	return strconv.FormatFloat(o.value, 'g', -1, 64)
}

func (o *Float) GetValue() any {
	return o.value
}

func (o *Float) BinaryOp(op int, rhs Object) (Object, error) {
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
	}
	return nil, ErrInvalidOp
}

func (o *Float) UnaryOp(op int) (Object, error) {
	switch op {
	case parser.MlanLexerNot:
		return o.LogicalNot()
	case parser.MlanLexerSubtract:
		return o.Negative()
	}
	return nil, ErrInvalidOp
}

func (o *Float) LogicalNot() (Object, error) {
	return NewBool(!utils.FloatToBool(o.value)), nil
}

func (o *Float) Negative() (Object, error) {
	o.value = -o.value
	return o, nil
}

func (o *Float) LogicalOr(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return o, nil
	}
	return nil, ErrInvalidOp
}

func (o *Float) LogicalAnd(rs Object) (Object, error) {
	switch rs.(type) {
	case *Null:
		return rs, nil
	}
	return nil, ErrInvalidOp
}

func (o *Float) Equal(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(o.value == utils.BoolToFloat(rs.(*Bool).value)), nil
	case *Dict:
		return NewBool(false), nil
	case *Float:
		return NewBool(o.value == rs.(*Float).value), nil
	case *Int:
		return NewBool(o.value == utils.IntToFloat(rs.(*Int).value)), nil
	case *List:
		return NewBool(false), nil
	case *Null:
		return NewBool(false), nil
	case *Str:
		return NewBool(false), nil
	}
	return nil, ErrInvalidOp
}

func (o *Float) NotEqual(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(o.value != utils.BoolToFloat(rs.(*Bool).value)), nil
	case *Dict:
		return NewBool(true), nil
	case *Float:
		return NewBool(o.value != rs.(*Float).value), nil
	case *Int:
		return NewBool(o.value != utils.IntToFloat(rs.(*Int).value)), nil
	case *List:
		return NewBool(true), nil
	case *Null:
		return NewBool(true), nil
	case *Str:
		return NewBool(true), nil
	}
	return nil, ErrInvalidOp
}

func (o *Float) GtEq(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(o.value >= utils.BoolToFloat(rs.(*Bool).value)), nil
	case *Float:
		return NewBool(o.value >= rs.(*Float).value), nil
	case *Int:
		return NewBool(o.value >= utils.IntToFloat(rs.(*Int).value)), nil
	}
	return nil, ErrInvalidOp
}

func (o *Float) Gt(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(o.value > utils.BoolToFloat(rs.(*Bool).value)), nil
	case *Float:
		return NewBool(o.value > rs.(*Float).value), nil
	case *Int:
		return NewBool(o.value > utils.IntToFloat(rs.(*Int).value)), nil
	}
	return nil, ErrInvalidOp
}

func (o *Float) LtEq(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(o.value <= utils.BoolToFloat(rs.(*Bool).value)), nil
	case *Float:
		return NewBool(o.value <= rs.(*Float).value), nil
	case *Int:
		return NewBool(o.value <= utils.IntToFloat(rs.(*Int).value)), nil
	}
	return nil, ErrInvalidOp
}

func (o *Float) Lt(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(o.value < utils.BoolToFloat(rs.(*Bool).value)), nil
	case *Float:
		return NewBool(o.value < rs.(*Float).value), nil
	case *Int:
		return NewBool(o.value < utils.IntToFloat(rs.(*Int).value)), nil
	}
	return nil, ErrInvalidOp
}

func (o *Float) Add(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewFloat(o.value + utils.BoolToFloat(rs.(*Bool).value)), nil
	case *Float:
		return NewFloat(o.value + rs.(*Float).value), nil
	case *Int:
		return NewFloat(o.value + utils.IntToFloat(rs.(*Int).value)), nil
	}
	return nil, ErrInvalidOp
}

func (o *Float) Sub(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewFloat(o.value - utils.BoolToFloat(rs.(*Bool).value)), nil
	case *Float:
		return NewFloat(o.value - rs.(*Float).value), nil
	case *Int:
		return NewFloat(o.value - utils.IntToFloat(rs.(*Int).value)), nil
	}
	return nil, ErrInvalidOp
}

func (o *Float) Pow(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewFloat(math.Pow(o.value, utils.BoolToFloat(rs.(*Bool).value))), nil
	case *Float:
		return NewFloat(math.Pow(o.value, rs.(*Float).value)), nil
	case *Int:
		return NewFloat(math.Pow(o.value, utils.IntToFloat(rs.(*Int).value))), nil
	}
	return nil, ErrInvalidOp
}

func (o *Float) Mul(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewFloat(o.value * utils.BoolToFloat(rs.(*Bool).value)), nil
	case *Float:
		return NewFloat(o.value * rs.(*Float).value), nil
	case *Int:
		return NewFloat(o.value * utils.IntToFloat(rs.(*Int).value)), nil
	}
	return nil, ErrInvalidOp
}

func (o *Float) Div(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		if !rs.(*Bool).value {
			return nil, ErrDivByZero
		}
		return NewFloat(o.value / utils.BoolToFloat(rs.(*Bool).value)), nil
	case *Float:
		if rs.(*Float).value == 0.0 {
			return nil, ErrDivByZero
		}
		return NewFloat(o.value / rs.(*Float).value), nil
	case *Int:
		if rs.(*Int).value == 0 {
			return nil, ErrDivByZero
		}
		return NewFloat(o.value / utils.IntToFloat(rs.(*Int).value)), nil
	}
	return nil, ErrInvalidOp
}

func (o *Float) Mod(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		if !rs.(*Bool).value {
			return nil, ErrModByZero
		}
		return NewFloat(math.Mod(o.value, utils.BoolToFloat(rs.(*Bool).value))), nil
	case *Float:
		if rs.(*Float).value == 0.0 {
			return nil, ErrModByZero
		}
		return NewFloat(math.Mod(o.value, rs.(*Float).value)), nil
	case *Int:
		if rs.(*Int).value == 0 {
			return nil, ErrModByZero
		}
		return NewFloat(math.Mod(o.value, utils.IntToFloat(rs.(*Int).value))), nil
	}
	return nil, ErrInvalidOp
}
