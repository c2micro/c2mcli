package engine

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/c2micro/mlan/pkg/engine/object"
	"github.com/c2micro/mlan/pkg/engine/scope"
)

var BuiltinFunctions = make(map[string]*object.BuiltinFunc)
var UserFunctions = make(map[string]*object.UserFunc)
var NativeFunctions = make(map[string]*object.NativeFunc)

var retValue object.Object

// invokeFunc вызов функции
func (v *Visitor) invokeFunc(
	token antlr.Token,
	name string,
	params ...object.Object,
) any {
	var val object.Object
	// поиск во встроенных функциях
	val, ok := BuiltinFunctions[name]
	if !ok {
		// поиск в функциях пользователя
		val, ok = UserFunctions[name]
		if !ok {
			val, ok = NativeFunctions[name]
			if !ok {
				v.LineError(token, fmt.Errorf("undefined function '%s'", name))
				return Failure
			}
		}
	}

	// проверяем, что объект callable
	if !val.CanCall() {
		v.LineError(token, fmt.Errorf("unable call function '%s'", name))
		return Failure
	}

	// создание нового контекста
	scope.CurrentScope = scope.NewScope(
		scope.CurrentScope,
		scope.CurrentScope.Depth()+1,
		true,
		false,
		make(map[string]object.Object),
	)
	defer func() {
		scope.CurrentScope = scope.CurrentScope.Parent()
	}()

	switch val.(type) {
	case *object.NativeFunc:
		if val.(*object.NativeFunc).GetArgsLen() != len(params) {
			v.LineError(token, fmt.Errorf("function '%s' expected %d arguments, got %d", name, val.(*object.NativeFunc).GetArgsLen(), len(params)))
			return Failure
		}
		// сохраняем аргументы в скоуп
		args := val.(*object.NativeFunc).GetArgs()
		for i, arg := range params {
			scope.CurrentScope.Put(args[i], arg)
		}
		// выполняем
		for _, item := range val.(*object.NativeFunc).GetCode() {
			if ok := v.Visit(item).(VisitResultType); !ok {
				return Failure
			}
			if retValue != nil {
				x := retValue
				retValue = nil
				return x
			}
		}

		return Success
	}

	// вызов функции
	res, err := val.Call(params...)
	if err != nil {
		v.LineError(token, err)
		return Failure
	}
	return res
}

func (v *Visitor) invokeClosureFunc(
	token antlr.Token,
	name string,
	params ...object.Object,
) any {
	// получаем из скоупа переменную
	fn := scope.CurrentScope.Get(name, true)
	if fn == nil {
		v.LineError(token, fmt.Errorf("undefined closure '%s'", name))
		return Failure
	}

	// проверяем, что объект callable
	if !fn.CanCall() {
		v.LineError(token, fmt.Errorf("unable call closure '%s'", name))
		return Failure
	}

	// создание нового контекста
	scope.CurrentScope = scope.NewScope(
		scope.CurrentScope,
		scope.CurrentScope.Depth()+1,
		true,
		false,
		make(map[string]object.Object),
	)
	defer func() {
		scope.CurrentScope = scope.CurrentScope.Parent()
	}()

	switch fn.(type) {
	case *object.NativeFunc:
		if fn.(*object.NativeFunc).GetArgsLen() != len(params) {
			v.LineError(token, fmt.Errorf("closure '%s' expected %d arguments, got %d", name, fn.(*object.NativeFunc).GetArgsLen(), len(params)))
			return Failure
		}
		// сохраняем аргументы в скоуп
		args := fn.(*object.NativeFunc).GetArgs()
		for i, arg := range params {
			scope.CurrentScope.Put(args[i], arg)
		}
		// выполняем
		for _, item := range fn.(*object.NativeFunc).GetCode() {
			if ok := v.Visit(item).(VisitResultType); !ok {
				return Failure
			}
			if retValue != nil {
				x := retValue
				retValue = nil
				return x
			}
		}

		return Success
	}

	return Success
}

func (v *Visitor) InvokeNativeFunc(
	fn *object.NativeFunc,
	params ...object.Object,
) any {
	// создание нового контекста
	scope.CurrentScope = scope.NewScope(
		scope.CurrentScope,
		scope.CurrentScope.Depth()+1,
		true,
		false,
		make(map[string]object.Object),
	)
	defer func() {
		scope.CurrentScope = scope.CurrentScope.Parent()
	}()

	if fn.GetArgsLen() != len(params) {
		v.Error = fmt.Errorf("invalid number of args, expecting %d arguments, got %d", fn.GetArgsLen(), len(params))
		return Failure
	}
	// сохраняем аргументы в скоуп
	args := fn.GetArgs()
	for i, arg := range params {
		scope.CurrentScope.Put(args[i], arg)
	}
	// выполняем
	for _, item := range fn.GetCode() {
		if val, ok := v.Visit(item).(VisitResultType); ok {
			if !val {
				return Failure
			}
		}
		if retValue != nil {
			x := retValue
			retValue = nil
			return x
		}
	}

	return Success
}
