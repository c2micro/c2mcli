package storage

import "github.com/c2micro/mlan/pkg/engine/object"

var BuiltinFunctions = make(map[string]*object.BuiltinFunc)
var UserFunctions = make(map[string]*object.UserFunc)
var NativeFunctions = make(map[string]*object.NativeFunc)
