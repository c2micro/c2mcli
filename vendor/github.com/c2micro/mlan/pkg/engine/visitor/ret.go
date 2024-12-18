package visitor

import "github.com/c2micro/mlan/pkg/engine/object"

var retValue object.Object

func ClearRet() {
	retValue = nil
}
