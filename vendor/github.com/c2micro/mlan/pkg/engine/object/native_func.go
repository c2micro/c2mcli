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
	code    []parser.IStmtContext
}

func NewNativeFunc(
	args []string,
	argsLen int,
	code []parser.IStmtContext,
) *NativeFunc {
	return &NativeFunc{
		name:    "<unknown>",
		args:    args,
		argsLen: argsLen,
		code:    code,
	}
}

func (o *NativeFunc) SetName(name string) {
	o.name = name
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

func (o *NativeFunc) GetCode() []parser.IStmtContext {
	return o.code
}

func (o *NativeFunc) GetArgsLen() int {
	return o.argsLen
}

func (o *NativeFunc) GetArgs() []string {
	return o.args
}
