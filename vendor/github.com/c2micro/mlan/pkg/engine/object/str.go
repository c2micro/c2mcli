package object

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/c2micro/mlan/pkg/engine/utils"
	"github.com/c2micro/mlan/pkg/parser"
)

// Str тип с плавающей точкой
type Str struct {
	Impl
	value   string
	methods map[string]*BuiltinFunc
}

func NewStr(v string) *Str {
	s := &Str{
		value: v,
	}
	s.fillMethods()
	return s
}

func (o *Str) fillMethods() {
	o.methods = make(map[string]*BuiltinFunc)
	o.methods["len"] = NewBuiltinFunc("len", o.MethodLen)
	o.methods["reverse"] = NewBuiltinFunc("reverse", o.MethodReverse)
	o.methods["split"] = NewBuiltinFunc("split", o.MethodSplit)
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
	if idx.value < 0 || idx.value >= int64(utf8.RuneCountInString(o.value)) {
		return nil, ErrIndexOutOfRange
	}
	// ширина руны (в количестве байт)
	w := 0
	// количество рун
	c := 0
	for i := 0; i < len(o.value); i += w {
		var r rune
		r, w = utf8.DecodeRuneInString(o.value[i:])
		if c == int(idx.value) {
			return NewStr(string(r)), nil
		}
		c++
	}
	return nil, ErrNotImplemented
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
	if idxInt.value < 0 || idxInt.value >= int64(utf8.RuneCountInString(o.value)) {
		return ErrIndexOutOfRange
	}
	if utf8.RuneCountInString(rsStr.value) != 1 {
		return fmt.Errorf("expected str with length of 1, got %d", len(rsStr.value))
	}
	res := ""
	// ширина руны (в количестве байт)
	w := 0
	// количество рун
	c := 0
	for i := 0; i < len(o.value); i += w {
		var r rune
		r, w = utf8.DecodeRuneInString(o.value[i:])
		if c == int(idxInt.value) {
			res += rsStr.value
		} else {
			res += string(r)
		}
		c++
	}
	o.value = res
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
	case parser.MlanLexerAssSum:
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
		return NewStr(strings.Repeat(o.value, int(utils.BoolToInt(rs.(*Bool).value)))), nil
	case *Int:
		// in case of negative value
		count := 0
		if int(rs.(*Int).value) > 0 {
			count = int(rs.(*Int).value)
		}
		return NewStr(strings.Repeat(o.value, count)), nil
	}
	return nil, ErrInvalidOp
}

func (o *Str) MethodCall(name string, args ...Object) (Object, error) {
	m := o.methods[name]
	if m == nil {
		return nil, ErrUnknownMethod
	}
	return m.Call(args...)
}

func (o *Str) MethodLen(args ...Object) (Object, error) {
	if len(args) > 0 {
		return nil, fmt.Errorf("expecting 0 arguments, got %d", len(args))
	}
	return NewInt(int64(utf8.RuneCountInString(o.value))), nil
}

func (o *Str) MethodReverse(args ...Object) (Object, error) {
	if len(args) > 0 {
		return nil, fmt.Errorf("expecting 0 arguments, got %d", len(args))
	}
	size := len(o.value)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(o.value[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	o.value = string(buf)
	return NewNull(), nil
}

func (o *Str) MethodSplit(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 arguments, got %d", len(args))
	}
	key, ok := args[0].(*Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[1].TypeName())
	}
	temp := strings.Split(o.value, key.value)
	result := make([]Object, 0)
	for _, v := range temp {
		result = append(result, NewStr(v))
	}
	return NewList(result), nil
}
