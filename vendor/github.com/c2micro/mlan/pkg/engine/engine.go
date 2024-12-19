package engine

import (
	"os"

	"github.com/c2micro/mlan/pkg/engine/builtin"
	"github.com/c2micro/mlan/pkg/engine/object"
	"github.com/c2micro/mlan/pkg/engine/scope"
	"github.com/c2micro/mlan/pkg/engine/storage"
	"github.com/c2micro/mlan/pkg/engine/types"
	"github.com/c2micro/mlan/pkg/engine/utils"
	"github.com/c2micro/mlan/pkg/engine/visitor"
	"github.com/go-faster/errors"
)

func init() {
	Init()
}

// инициализация движка
func Init() {
	scope.GlobalScope = scope.NewScope(
		nil,
		0,
		false,
		false,
		map[string]object.Object{},
	)
	scope.CurrentScope = scope.GlobalScope
	builtin.Register()
}

// очистка движка и перерегистрация всех необходимых компонентов
func Clear() {
	// зануление хендлеров функций
	storage.BuiltinFunctions = make(map[string]*object.BuiltinFunc)
	storage.UserFunctions = make(map[string]*object.UserFunc)
	storage.NativeFunctions = make(map[string]*object.NativeFunc)
	// зануление рета
	visitor.ClearRet()
	// инициализация
	Init()
}

// выполнение скрипта из файла
func Evaluate(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	tree, err := utils.CreateAST(string(data))
	if err != nil {
		return err
	}

	v := visitor.NewVisitor()
	res := v.Visit(tree)
	if res != types.Success {
		return errors.Wrap(v.GetError(), "runtime failure")
	}

	return nil
}
