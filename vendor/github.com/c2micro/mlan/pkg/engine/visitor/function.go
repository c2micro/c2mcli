package visitor

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/c2micro/mlan/pkg/engine/object"
	"github.com/c2micro/mlan/pkg/engine/scope"
	"github.com/c2micro/mlan/pkg/engine/storage"
	"github.com/c2micro/mlan/pkg/engine/types"
)

// invokeFunc вызов функции
func (v *Visitor) invokeFunc(
	token antlr.Token,
	name string,
	params ...object.Object,
) any {
	var val object.Object
	// поиск во встроенных функциях
	val, ok := storage.BuiltinFunctions[name]
	if !ok {
		// поиск в функциях пользователя
		val, ok = storage.UserFunctions[name]
		if !ok {
			val, ok = storage.NativeFunctions[name]
			if !ok {
				v.LineError(token, fmt.Errorf("undefined function '%s'", name))
				return types.Failure
			}
		}
	}

	// проверяем, что объект callable
	if !val.CanCall() {
		v.LineError(token, fmt.Errorf("unable call function '%s'", name))
		return types.Failure
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
			return types.Failure
		}
		// сохраняем аргументы в скоуп
		args := val.(*object.NativeFunc).GetArgs()
		for i, arg := range params {
			scope.CurrentScope.Put(args[i], arg)
		}
		// выполняем
		for _, item := range val.(*object.NativeFunc).GetCode() {
			if ok := v.Visit(item).(types.VisitResultType); !ok {
				return types.Failure
			}
			if retValue != nil {
				x := retValue
				retValue = nil
				return x
			}
		}

		return types.Success
	}

	// вызов функции
	res, err := val.Call(params...)
	if err != nil {
		v.LineError(token, err)
		return types.Failure
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
		return types.Failure
	}

	// проверяем, что объект callable
	if !fn.CanCall() {
		v.LineError(token, fmt.Errorf("unable call closure '%s'", name))
		return types.Failure
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
			return types.Failure
		}
		// сохраняем аргументы в скоуп
		args := fn.(*object.NativeFunc).GetArgs()
		for i, arg := range params {
			scope.CurrentScope.Put(args[i], arg)
		}
		// выполняем
		for _, item := range fn.(*object.NativeFunc).GetCode() {
			if ok := v.Visit(item).(types.VisitResultType); !ok {
				return types.Failure
			}
			if retValue != nil {
				x := retValue
				retValue = nil
				return x
			}
		}

		return types.Success
	}

	return types.Success
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
		return types.Failure
	}
	// сохраняем аргументы в скоуп
	args := fn.GetArgs()
	for i, arg := range params {
		scope.CurrentScope.Put(args[i], arg)
	}
	// выполняем
	for _, item := range fn.GetCode() {
		if val, ok := v.Visit(item).(types.VisitResultType); ok {
			if !val {
				return types.Failure
			}
		}
		if retValue != nil {
			x := retValue
			retValue = nil
			return x
		}
	}

	return types.Success
}
