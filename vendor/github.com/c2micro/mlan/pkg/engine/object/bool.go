package object

import (
	"math"
	"strings"

	"github.com/c2micro/mlan/pkg/parser"
	"github.com/c2micro/mlan/pkg/engine/utils"
)

// Bool булевый тип
type Bool struct {
	Impl
	ArithmeticImpl
	value bool
}

func NewBool(v bool) *Bool {
	return &Bool{value: v}
}

func (o *Bool) TypeName() string {
	return "bool"
}

func (o *Bool) String() string {
	if o.value {
		return "true"
	} else {
		return "false"
	}
}

func (o *Bool) GetValue() any {
	return o.value
}

func (o *Bool) BinaryOp(op int, rhs Object) (Object, error) {
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
	case parser.MlanLexerXor:
		return o.Xor(rhs)
	}
	return nil, ErrInvalidOp
}

func (o *Bool) UnaryOp(op int) (Object, error) {
	switch op {
	case parser.MlanLexerNot:
		return o.LogicalNot()
	case parser.MlanLexerSubtract:
		return o.Negative()
	}
	return nil, ErrInvalidOp
}

func (o *Bool) LogicalNot() (Object, error) {
	o.value = !o.value
	return o, nil
}

func (o *Bool) Negative() (Object, error) {
	return NewInt(-utils.BoolToInt(o.value)), nil
}

func (o *Bool) LogicalOr(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(o.value || rs.(*Bool).value), nil
	case *Null:
		return o, nil
	}
	return nil, ErrInvalidOp
}

func (o *Bool) LogicalAnd(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(o.value && rs.(*Bool).value), nil
	case *Null:
		return rs, nil
	}
	return nil, ErrInvalidOp
}

func (o *Bool) Equal(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(o.value == rs.(*Bool).value), nil
	case *Dict:
		return NewBool(false), nil
	case *Float:
		return NewBool(utils.BoolToFloat(o.value) == rs.(*Float).value), nil
	case *Int:
		return NewBool(utils.BoolToInt(o.value) == rs.(*Int).value), nil
	case *List:
		return NewBool(false), nil
	case *Null:
		return NewBool(false), nil
	case *Str:
		return NewBool(false), nil
	}
	return nil, ErrInvalidOp
}

func (o *Bool) NotEqual(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(o.value != rs.(*Bool).value), nil
	case *Dict:
		return NewBool(true), nil
	case *Float:
		return NewBool(utils.BoolToFloat(o.value) != rs.(*Float).value), nil
	case *Int:
		return NewBool(utils.BoolToInt(o.value) != rs.(*Int).value), nil
	case *List:
		return NewBool(true), nil
	case *Null:
		return NewBool(true), nil
	case *Str:
		return NewBool(true), nil
	}
	return nil, ErrInvalidOp
}

func (o *Bool) GtEq(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(utils.BoolToInt(o.value) >= utils.BoolToInt(rs.(*Bool).value)), nil
	case *Float:
		return NewBool(utils.BoolToFloat(o.value) >= rs.(*Float).value), nil
	case *Int:
		return NewBool(utils.BoolToInt(o.value) >= rs.(*Int).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Bool) Gt(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(utils.BoolToInt(o.value) > utils.BoolToInt(rs.(*Bool).value)), nil
	case *Float:
		return NewBool(utils.BoolToFloat(o.value) > rs.(*Float).value), nil
	case *Int:
		return NewBool(utils.BoolToInt(o.value) > rs.(*Int).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Bool) LtEq(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(utils.BoolToInt(o.value) <= utils.BoolToInt(rs.(*Bool).value)), nil
	case *Float:
		return NewBool(utils.BoolToFloat(o.value) <= rs.(*Float).value), nil
	case *Int:
		return NewBool(utils.BoolToInt(o.value) <= rs.(*Int).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Bool) Lt(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(utils.BoolToInt(o.value) < utils.BoolToInt(rs.(*Bool).value)), nil
	case *Float:
		return NewBool(utils.BoolToFloat(o.value) < rs.(*Float).value), nil
	case *Int:
		return NewBool(utils.BoolToInt(o.value) < rs.(*Int).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Bool) Add(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewInt(utils.BoolToInt(o.value) + utils.BoolToInt(rs.(*Bool).value)), nil
	case *Float:
		return NewFloat(utils.BoolToFloat(o.value) + rs.(*Float).value), nil
	case *Int:
		return NewInt(utils.BoolToInt(o.value) + rs.(*Int).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Bool) Sub(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewInt(utils.BoolToInt(o.value) - utils.BoolToInt(rs.(*Bool).value)), nil
	case *Float:
		return NewFloat(utils.BoolToFloat(o.value) - rs.(*Float).value), nil
	case *Int:
		return NewInt(utils.BoolToInt(o.value) - rs.(*Int).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Bool) Pow(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewInt(int64(math.Pow(utils.BoolToFloat(o.value), utils.BoolToFloat(rs.(*Bool).value)))), nil
	case *Float:
		return NewFloat(math.Pow(utils.BoolToFloat(o.value), rs.(*Float).value)), nil
	case *Int:
		return NewInt(int64(math.Pow(utils.BoolToFloat(o.value), utils.IntToFloat(rs.(*Int).value)))), nil
	}
	return nil, ErrInvalidOp
}

func (o *Bool) Mul(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewInt(utils.BoolToInt(o.value) * utils.BoolToInt(rs.(*Bool).value)), nil
	case *Float:
		return NewFloat(utils.BoolToFloat(o.value) * rs.(*Float).value), nil
	case *Int:
		return NewInt(utils.BoolToInt(o.value) * rs.(*Int).value), nil
	case *List:
		var list []Object
		var i int64
		for i = 0; i < utils.BoolToInt(o.value); i++ {
			list = append(list, rs.(*List).value...)
		}
		return NewList(list), nil
	case *Str:
		return NewStr(strings.Repeat(rs.(*Str).value, int(utils.BoolToInt(o.value)))), nil
	}
	return nil, ErrInvalidOp
}

func (o *Bool) Div(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		if !rs.(*Bool).value {
			return nil, ErrDivByZero
		}
		return NewFloat(utils.BoolToFloat(o.value) / utils.BoolToFloat(rs.(*Bool).value)), nil
	case *Float:
		if rs.(*Float).value == 0.0 {
			return nil, ErrDivByZero
		}
		return NewFloat(utils.BoolToFloat(o.value) / rs.(*Float).value), nil
	case *Int:
		if rs.(*Int).value == 0 {
			return nil, ErrDivByZero
		}
		return NewFloat(utils.BoolToFloat(o.value) / utils.IntToFloat(rs.(*Int).value)), nil
	}
	return nil, ErrInvalidOp
}

func (o *Bool) Mod(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		if !rs.(*Bool).value {
			return nil, ErrModByZero
		}
		return NewInt(utils.BoolToInt(o.value) % utils.BoolToInt(rs.(*Bool).value)), nil
	case *Float:
		if rs.(*Float).value == 0.0 {
			return nil, ErrModByZero
		}
		return NewFloat(math.Mod(utils.BoolToFloat(o.value), rs.(*Float).value)), nil
	case *Int:
		if rs.(*Int).value == 0 {
			return nil, ErrModByZero
		}
		return NewInt(utils.BoolToInt(o.value) % rs.(*Int).value), nil
	}
	return nil, ErrInvalidOp
}

func (o *Bool) Xor(rs Object) (Object, error) {
	switch rs.(type) {
	case *Bool:
		return NewBool(utils.IntToBool(utils.BoolToInt(o.value) ^ utils.BoolToInt(rs.(*Bool).value))), nil
	case *Int:
		return NewInt(utils.BoolToInt(o.value) ^ rs.(*Int).value), nil
	}
	return nil, ErrInvalidOp
}
