package bexecdetach

import (
	"fmt"

	merror "github.com/c2micro/c2mcli/internal/scripts/aliases/m_error"
	"github.com/c2micro/c2mcli/internal/service"
	"github.com/c2micro/c2mcli/internal/storage/beacon"
	"github.com/c2micro/c2mshr/defaults"
	commonv1 "github.com/c2micro/c2mshr/proto/gen/common/v1"
	operatorv1 "github.com/c2micro/c2mshr/proto/gen/operator/v1"
	"github.com/c2micro/mlan/pkg/engine/object"
)

// имя API
const name = "b_exec_detach"

// получение имени API
func GetApiName() string {
	return name
}

func UserBeaconExecDetach(args ...object.Object) (object.Object, error) {
	if len(args) < 2 || len(args) > 3 {
		return nil, fmt.Errorf("expecting 2 or 3 arguments, got %d", len(args))
	}
	id, ok := args[0].(*object.Int)
	if !ok {
		return nil, fmt.Errorf("expecting 1st argument int, got '%s'", args[0].TypeName())
	}
	cmd, ok := args[1].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting 2nd argument str, got '%s'", args[1].TypeName())
	}
	arg := object.NewStr("")
	if len(args) == 3 {
		arg, ok = args[2].(*object.Str)
		if !ok {
			return nil, fmt.Errorf("expecting 3rd argument str, got '%s'", args[2].TypeName())
		}
	}
	if err := BackendBeaconExecDetach(uint32(id.GetValue().(int64)), cmd.GetValue().(string), arg.GetValue().(string)); err != nil {
		return nil, err
	}
	return object.NewNull(), nil
}

func BackendBeaconExecDetach(id uint32, cmd string, args string) error {
	cap := defaults.CAP_EXEC_DETACH

	// проверка существования бикона
	b := beacon.Beacons.GetById(id)
	if b == nil {
		return fmt.Errorf("no beacon with id %d", id)
	}

	// проверка капы
	if !cap.ValidateMask(b.GetCaps()) {
		return merror.BackendMessageError(id, fmt.Sprintf("beacon doesn't support %s", cap.String()))
	}

	return service.NewTask(id, &operatorv1.NewTaskRequest{
		Cap: uint32(cap),
		Args: &operatorv1.NewTaskRequest_ExecDetach{
			ExecDetach: &commonv1.CapExecDetach{
				Cmd:  cmd,
				Args: args,
			},
		},
	})
}