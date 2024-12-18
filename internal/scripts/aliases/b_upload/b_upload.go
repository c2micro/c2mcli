package bupload

import (
	"fmt"
	"os"

	merror "github.com/c2micro/c2mcli/internal/scripts/aliases/m_error"
	"github.com/c2micro/c2mcli/internal/service"
	"github.com/c2micro/c2mcli/internal/storage/beacon"
	"github.com/c2micro/c2mshr/defaults"
	commonv1 "github.com/c2micro/c2mshr/proto/gen/common/v1"
	operatorv1 "github.com/c2micro/c2mshr/proto/gen/operator/v1"
	"github.com/c2micro/mlan/pkg/engine/object"
	"github.com/go-faster/errors"
)

// имя API
const name = "b_upload"

// получение имени API
func GetApiName() string {
	return name
}

func UserBeaconUpload(args ...object.Object) (object.Object, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("expecting 3 arguments, got %d", len(args))
	}
	id, ok := args[0].(*object.Int)
	if !ok {
		return nil, fmt.Errorf("expecting 1st argument int, got '%s'", args[0].TypeName())
	}
	src, ok := args[1].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting 2nd argument str, got '%s'", args[1].TypeName())
	}
	dst, ok := args[2].(*object.Str)
	if !ok {
		return nil, fmt.Errorf("expecting 3rd argument str, got '%s'", args[2].TypeName())
	}
	if err := BackendBeaconUpload(uint32(id.GetValue().(int64)), src.GetValue().(string), dst.GetValue().(string)); err != nil {
		return nil, err
	}
	return object.NewNull(), nil
}

func BackendBeaconUpload(id uint32, src, dst string) error {
	cap := defaults.CAP_UPLOAD

	// проверка существования бикона
	b := beacon.Beacons.GetById(id)
	if b == nil {
		return fmt.Errorf("no beacon with id %d", id)
	}

	// проверка капы
	if !cap.ValidateMask(b.GetCaps()) {
		return merror.BackendMessageError(id, fmt.Sprintf("beacon doesn't support %s", cap.String()))
	}

	// получаем данные из файла по локальному пути
	data, err := os.ReadFile(src)
	if err != nil {
		// ошибка вычитывания файла с .NET программой
		if os.IsNotExist(err) {
			err = errors.New("no such file")
		} else if os.IsPermission(err) {
			err = errors.New("permission denied")
		}
		return merror.BackendMessageError(id, fmt.Sprintf("unable open local file by path %s: %s", src, err.Error()))
	}

	return service.NewTask(id, &operatorv1.NewTaskRequest{
		Cap: uint32(cap),
		Args: &operatorv1.NewTaskRequest_Upload{
			Upload: &commonv1.CapUpload{
				Path: dst,
				Blob: data,
			},
		},
	})
}
