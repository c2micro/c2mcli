package alias

import (
	"fmt"
	"strings"

	"github.com/c2micro/c2mcli/internal/scripts/aliases"
	"github.com/c2micro/c2mcli/internal/service"
	"github.com/c2micro/mlan/pkg/engine/object"
	"github.com/c2micro/mlan/pkg/engine/visitor"
	"github.com/fatih/color"
	"github.com/go-faster/errors"
	"github.com/google/shlex"
)

// имя API
const name = "alias"

// получение имени API
func GetApiName() string {
	return name
}

// регистрация нового алиаса
// args[0] - имя алиаса
// args[1] - closure на функцию, выполняющую алиас
// args[2] - описание алиаса
// args[3] - короткая справка по использованию алиаса
// args[4] - являются ли артефакты выполнения алиаса доступными для других операторов
func UserAlias(args ...object.Object) (object.Object, error) {
	if len(args) != 5 {
		return nil, fmt.Errorf("expecting 5 arguments, got %d", len(args))
	}
	name, ok := args[0].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 1st argument, got '%s'", args[0].TypeName())
	}
	closure, ok := args[1].(*object.NativeFunc)
	if !ok {
		return nil, fmt.Errorf("expecting closure as 2nd argument, got '%s'", args[1].TypeName())
	}
	description, ok := args[2].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 3rd argument, got '%s'", args[2].TypeName())
	}
	usage, ok := args[3].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting str as 4th argument, got '%s'", args[3].TypeName())
	}
	visible, ok := args[4].(*object.Bool)
	if !ok {
		return nil, fmt.Errorf("expecting bool as 5th argument, got '%s'", args[4].TypeName())
	}
	// создание объекта алиаса
	newAlias := &aliases.Alias{}
	newAlias.SetDescription(description.GetValue().(string))
	newAlias.SetUsage(usage.GetValue().(string))
	newAlias.SetVisible(visible.GetValue().(bool))
	newAlias.SetClosure(closure)
	// добавление алиаса в общую мапу
	aliases.Aliases[name.GetValue().(string)] = newAlias
	return object.NewNull(), nil
}

// выполнение алиаса в рамках таск группы
func BackendAlias(id uint32, cmd string) error {
	// сплит строки
	s, err := shlex.Split(cmd)
	if err != nil {
		return errors.Wrap(err, "split command")
	}

	// получение алиаса
	al, ok := aliases.Aliases[s[0]]
	if !ok {
		return fmt.Errorf("unknown alias '%s'", s[0])
	}

	// создание таск группы
	if err := service.NewTaskGroup(id, cmd, al.GetVisible()); err != nil {
		return errors.Wrap(err, "open task group")
	}
	defer func(id uint32) {
		// закрытие таск группы
		if err := service.CloseTaskGroup(id); err != nil {
			color.Red(err.Error())
		}
	}(id)

	// bid
	arg0 := object.NewInt(int64(id))
	// cmd
	arg1 := object.NewStr(s[0])
	var temp []object.Object
	if len(s) != 1 {
		for i := 1; i < len(s); i++ {
			temp = append(temp, object.NewStr(s[i]))
		}
	}
	// args
	arg2 := object.NewList(temp)
	// raw
	arg3 := object.NewStr("")
	if len(strings.Split(cmd, " ")) != 1 {
		arg3 = object.NewStr(strings.Join(strings.Split(cmd, " ")[1:], " "))
	}

	// вызов нативной функции (по факту кложура)
	v := visitor.NewVisitor()
	v.InvokeNativeFunc(al.GetClosure(), arg0, arg1, arg2, arg3)
	if v.GetError() != nil {
		return v.GetError()
	}
	return nil
}
