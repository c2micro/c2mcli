package object

import (
	"fmt"
	"strings"

	"github.com/c2micro/mlan/pkg/parser"
)

// Dict мапа
type Dict struct {
	Impl
	value   map[string]Object
	methods map[string]*BuiltinFunc
}

func NewDict(v map[string]Object) *Dict {
	d := &Dict{value: v}
	d.fillMethods()
	return d
}

func (o *Dict) fillMethods() {
	o.methods = make(map[string]*BuiltinFunc)
	o.methods["len"] = NewBuiltinFunc("len", o.MethodLen)
	o.methods["pop"] = NewBuiltinFunc("pop", o.MethodPop)
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
		if len(o.value) != len(rs.(*Dict).value) {
			return NewBool(false), nil
		}
		for k, v := range o.value {
			if (rs.(*Dict).value)[k] == nil {
				return NewBool(false), nil
			}
			val, err := v.Equal(rs.(*Dict).value[k])
			if err != nil {
				return nil, err
			}
			if !val.(*Bool).value {
				return NewBool(false), nil
			}
		}
		return NewBool(true), nil
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
		if len(o.value) != len(rs.(*Dict).value) {
			return NewBool(true), nil
		}
		for k, v := range o.value {
			if (rs.(*Dict).value)[k] == nil {
				return NewBool(true), nil
			}
			val, err := v.Equal(rs.(*Dict).value[k])
			if err != nil {
				return nil, err
			}
			if !val.(*Bool).value {
				return NewBool(true), nil
			}
		}
		return NewBool(false), nil
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

func (o *Dict) MethodCall(name string, args ...Object) (Object, error) {
	m := o.methods[name]
	if m == nil {
		return nil, ErrUnknownMethod
	}
	return m.Call(args...)
}

func (o *Dict) MethodLen(args ...Object) (Object, error) {
	if len(args) > 0 {
		return nil, fmt.Errorf("expecting 0 arguments, got %d", len(args))
	}
	return NewInt(int64(len(o.value))), nil
}

func (o *Dict) MethodPop(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 arguments, got %d", len(args))
	}
	key, ok := args[0].(*Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[1].TypeName())
	}
	delete(o.value, key.GetValue().(string))
	return NewNull(), nil
}
