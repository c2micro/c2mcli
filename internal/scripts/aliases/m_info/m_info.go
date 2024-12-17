package minfo

import (
	"fmt"

	"github.com/c2micro/c2mcli/internal/service"
	"github.com/c2micro/c2mshr/defaults"
	"github.com/c2micro/mlan/pkg/engine/object"
)

// имя API
const name = "m_info"

// получение имени API
func GetApiName() string {
	return name
}

// отправка сообщения в группу с типом INFO
// args[0] - ID бикона
// args[1] - сообщение
func UserMessageInfo(args ...object.Object) (object.Object, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("expecting 2 arguments, got %d", len(args))
	}
	id, ok := args[0].(*object.Int)
	if !ok {
		return nil, fmt.Errorf("expecting 1st argument int, got '%s'", args[0].TypeName())
	}
	msg, ok := args[1].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expectign 2nd argument str, got '%s'", args[1].TypeName())
	}
	if err := BackendMessageInfo(uint32(id.GetValue().(int64)), msg.GetValue().(string)); err != nil {
		return nil, err
	}
	return object.NewNull(), nil
}

// отправка сообщения с типом INFO
func BackendMessageInfo(id uint32, message string) error {
	return service.NewTaskGroupMessage(id, defaults.InfoMessage, message)
}
