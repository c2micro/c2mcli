package scripts

import (
	"github.com/c2micro/c2mcli/internal/scripts/aliases"
	"github.com/c2micro/mlan/pkg/engine"
)

// инициализация движка скриптинга
func Init() error {
	// регистрация api
	registerApi()
	// регистрация builtin скриптов
	if err := registerBuiltin(); err != nil {
		return err
	}
	return nil
}

// ребилд скриптов
func Rebuild() error {
	// очистка хранилища алиасов
	aliases.Clear()
	// очистка рантайма
	engine.Clear()
	// инициализация с нуля
	Init()
	// перерегистрация внешних скриптов
	externalScripts := make([]*Script, 0)
	scripts.Range(func(k string, v *Script) bool {
		externalScripts = append(externalScripts, v)
		scripts.Delete(k)
		return true
	})
	for _, v := range externalScripts {
		RegisterExternalByPath(v.path)
	}
	return nil
}
