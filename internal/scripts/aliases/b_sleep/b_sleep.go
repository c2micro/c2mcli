package bsleep

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
const name = "b_sleep"

// получение имени API
func GetApiName() string {
	return name
}

func UserBeaconSleep(args ...object.Object) (object.Object, error) {
	if len(args) < 2 || len(args) > 3 {
		return nil, fmt.Errorf("expecting 2 or 3 arguments, got %d", len(args))
	}
	id, ok := args[0].(*object.Int)
	if !ok {
		return nil, fmt.Errorf("expecting 1st argument int, got '%s'", args[0].TypeName())
	}
	sleep, ok := args[1].(*object.Int)
	if !ok {
		return nil, fmt.Errorf("expecting 2nd argument int, got '%s'", args[1].TypeName())
	}
	jitter := object.NewInt(0)
	if len(args) == 3 {
		jitter, ok = args[2].(*object.Int)
		if !ok {
			return nil, fmt.Errorf("expecting 3rd argument int, got '%s'", args[2].TypeName())
		}
	}
	if err := BackendSleep(uint32(id.GetValue().(int64)), uint32(sleep.GetValue().(int64)), uint32(jitter.GetValue().(int64))); err != nil {
		return nil, err
	}
	return object.NewNull(), nil
}

func BackendSleep(id uint32, sleep uint32, jitter uint32) error {
	cap := defaults.CAP_SLEEP

	// проверка существования бикона
	b := beacon.Beacons.GetById(id)
	if b == nil {
		return fmt.Errorf("no beacon with id %d", id)
	}

	// проверка капы
	if !cap.ValidateMask(b.GetCaps()) {
		return merror.BackendMessageError(id, "beacon doesn't support %s", cap.String())
	}

	return service.NewTask(id, &operatorv1.NewTaskRequest{
		Cap: uint32(cap),
		Args: &operatorv1.NewTaskRequest_Sleep{
			Sleep: &commonv1.CapSleep{
				Sleep:  sleep,
				Jitter: jitter,
			},
		},
	})
}
