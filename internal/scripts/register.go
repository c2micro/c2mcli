package scripts

import (
	"embed"
	"fmt"

	"github.com/c2micro/c2mcli/internal/scripts/aliases/alias"
	bls "github.com/c2micro/c2mcli/internal/scripts/aliases/b_ls"
	bsleep "github.com/c2micro/c2mcli/internal/scripts/aliases/b_sleep"
	merror "github.com/c2micro/c2mcli/internal/scripts/aliases/m_error"
	minfo "github.com/c2micro/c2mcli/internal/scripts/aliases/m_info"
	mnotify "github.com/c2micro/c2mcli/internal/scripts/aliases/m_notify"
	mwarning "github.com/c2micro/c2mcli/internal/scripts/aliases/m_warning"
	"github.com/c2micro/mlan/pkg/engine"
	"github.com/c2micro/mlan/pkg/engine/object"
	"github.com/go-faster/errors"
)

// регистрация API для интеграции MLAN с C2
func registerApi() {
	// alias: регистрация нового алиаса
	engine.UserFunctions[alias.GetApiName()] = object.NewUserFunc(alias.GetApiName(), alias.UserAlias)
	// m_notify: сообщение с типом NOTIFY
	engine.UserFunctions[mnotify.GetApiName()] = object.NewUserFunc(mnotify.GetApiName(), mnotify.UserMessageNotify)
	// m_info: сообщение с типом INFO
	engine.UserFunctions[minfo.GetApiName()] = object.NewUserFunc(minfo.GetApiName(), minfo.UserMessageInfo)
	// m_warning: сообщение с типом WARNING
	engine.UserFunctions[mwarning.GetApiName()] = object.NewUserFunc(mwarning.GetApiName(), mwarning.UserMessageWarning)
	// m_error: сообщение с типом ERROR
	engine.UserFunctions[merror.GetApiName()] = object.NewUserFunc(merror.GetApiName(), merror.UserMessageError)
	// b_sleep: изменение параметров sleep/jitter бикона
	engine.UserFunctions[bsleep.GetApiName()] = object.NewUserFunc(bsleep.GetApiName(), bsleep.UserBeaconSleep)
	// b_ls: получение листинга директорий
	engine.UserFunctions[bls.GetApiName()] = object.NewUserFunc(bls.GetApiName(), bls.UserBeaconLs)
}

var (
	//go:embed builtin/*.c2m
	builtinScriptsFS embed.FS
)

// регистрация встроенных скриптов с базовыми командами
func registerBuilin() error {
	// список скриптов
	e, err := builtinScriptsFS.ReadDir("builtin")
	if err != nil {
		return err
	}
	for _, v := range e {
		// читаем файл со скриптом
		data, err := builtinScriptsFS.ReadFile(fmt.Sprintf("builtin/%s", v.Name()))
		if err != nil {
			return errors.Wrapf(err, "read %s", v.Name())
		}
		// строим AST дерево
		tree, err := engine.CreateAST(string(data))
		if err != nil {
			return errors.Wrap(err, v.Name())
		}
		// проходим по дереву
		visitor := engine.NewVisitor()
		if res := visitor.Visit(tree); res != engine.Success {
			return errors.Wrapf(visitor.Error, "evaluation %s", v.Name())
		}
	}
	return nil
}
