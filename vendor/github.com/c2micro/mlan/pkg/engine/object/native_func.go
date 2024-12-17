package object

import (
	"github.com/c2micro/mlan/pkg/parser"
)

// NativeFunc функция
type NativeFunc struct {
	Impl
	name    string
	args    []string
	argsLen int
	code    []parser.IStatementContext
}

func NewNativeFunc(
	name string,
	args []string,
	argsLen int,
	code []parser.IStatementContext,
) *NativeFunc {
	return &NativeFunc{
		name:    name,
		args:    args,
		argsLen: argsLen,
		code:    code,
	}
}

func (o *NativeFunc) TypeName() string {
	return "native-func: " + o.name
}

func (o *NativeFunc) String() string {
	return "<native-func>"
}

func (o *NativeFunc) CanCall() bool {
	return true
}

func (o *NativeFunc) GetCode() []parser.IStatementContext {
	return o.code
}

func (o *NativeFunc) GetArgsLen() int {
	return o.argsLen
}

func (o *NativeFunc) GetArgs() []string {
	return o.args
}
