package isx64

import (
	"fmt"

	"github.com/c2micro/c2mcli/internal/scripts/aliases/shared"
	"github.com/c2micro/c2mshr/defaults"
	"github.com/c2micro/mlan/pkg/engine/object"
)

// имя API
const name = "is_x64"

// получение имени API
func GetApiName() string {
	return name
}

func UserIsX64(args ...object.Object) (object.Object, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expecting 1 argument, got %d", len(args))
	}
	id, ok := args[0].(*object.Int)
	if !ok {
		return nil, fmt.Errorf("expecting 1st argument int, got '%s'", args[0].TypeName())
	}
	return object.NewBool(shared.BackendIsArch(uint32(id.GetValue().(int64)), defaults.X64Arch)), nil
}
