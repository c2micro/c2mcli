package tcancel

import (
	"fmt"

	"github.com/c2micro/c2mcli/internal/service"
	"github.com/c2micro/mlan/pkg/engine/object"
)

// имя API
const name = "t_cancel"

// получение имени API
func GetApiName() string {
	return name
}

func UserBeaconCancel(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	id, ok := args[0].(*object.Int)
	if !ok {
		return nil, fmt.Errorf("expecting 1st argument int, got '%s'", args[0].TypeName())
	}
	if err := BackendBeaconCancel(uint32(id.GetValue().(int64))); err != nil {
		return nil, err
	}
	return object.NewNull(), nil
}

func BackendBeaconCancel(id uint32) error {
	return service.CancelTasks(id)
}
