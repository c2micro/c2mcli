package aliases

import "github.com/c2micro/mlan/pkg/engine/object"

// хранилище всех алиасов
var Aliases = make(map[string]*Alias)

// существует ли алиас в мапе
func IsAliasExists(n string) bool {
	_, ok := Aliases[n]
	return ok
}

type Alias struct {
	description string
	usage       string
	visible     bool
	closure     *object.NativeFunc
}

func (a *Alias) GetDescription() string {
	return a.description
}

func (a *Alias) SetDescription(description string) {
	a.description = description
}

func (a *Alias) GetUsage() string {
	return a.usage
}

func (a *Alias) SetUsage(usage string) {
	a.usage = usage
}

func (a *Alias) GetVisible() bool {
	return a.visible
}

func (a *Alias) SetVisible(flag bool) {
	a.visible = flag
}

func (a *Alias) GetClosure() *object.NativeFunc {
	return a.closure
}

func (a *Alias) SetClosure(closure *object.NativeFunc) {
	a.closure = closure
}
