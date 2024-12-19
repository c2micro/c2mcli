package scripts

import (
	"embed"
	"fmt"

	"github.com/c2micro/c2mcli/internal/scripts/aliases/alias"
	bcat "github.com/c2micro/c2mcli/internal/scripts/aliases/b_cat"
	bcd "github.com/c2micro/c2mcli/internal/scripts/aliases/b_cd"
	bcp "github.com/c2micro/c2mcli/internal/scripts/aliases/b_cp"
	bdestruct "github.com/c2micro/c2mcli/internal/scripts/aliases/b_destruct"
	bdownload "github.com/c2micro/c2mcli/internal/scripts/aliases/b_download"
	bexec "github.com/c2micro/c2mcli/internal/scripts/aliases/b_exec"
	bexecassembly "github.com/c2micro/c2mcli/internal/scripts/aliases/b_exec_assembly"
	bexecdetach "github.com/c2micro/c2mcli/internal/scripts/aliases/b_exec_detach"
	bexit "github.com/c2micro/c2mcli/internal/scripts/aliases/b_exit"
	bjobkill "github.com/c2micro/c2mcli/internal/scripts/aliases/b_jobkill"
	bjobs "github.com/c2micro/c2mcli/internal/scripts/aliases/b_jobs"
	bkill "github.com/c2micro/c2mcli/internal/scripts/aliases/b_kill"
	bls "github.com/c2micro/c2mcli/internal/scripts/aliases/b_ls"
	bmkdir "github.com/c2micro/c2mcli/internal/scripts/aliases/b_mkdir"
	bmv "github.com/c2micro/c2mcli/internal/scripts/aliases/b_mv"
	bpause "github.com/c2micro/c2mcli/internal/scripts/aliases/b_pause"
	bppid "github.com/c2micro/c2mcli/internal/scripts/aliases/b_ppid"
	bps "github.com/c2micro/c2mcli/internal/scripts/aliases/b_ps"
	bpwd "github.com/c2micro/c2mcli/internal/scripts/aliases/b_pwd"
	bshell "github.com/c2micro/c2mcli/internal/scripts/aliases/b_shell"
	bsleep "github.com/c2micro/c2mcli/internal/scripts/aliases/b_sleep"
	bupload "github.com/c2micro/c2mcli/internal/scripts/aliases/b_upload"
	bwhoami "github.com/c2micro/c2mcli/internal/scripts/aliases/b_whoami"
	isarm32 "github.com/c2micro/c2mcli/internal/scripts/aliases/is_arm32"
	isarm64 "github.com/c2micro/c2mcli/internal/scripts/aliases/is_arm64"
	islinux "github.com/c2micro/c2mcli/internal/scripts/aliases/is_linux_os"
	ismacos "github.com/c2micro/c2mcli/internal/scripts/aliases/is_macos_os"
	iswindows "github.com/c2micro/c2mcli/internal/scripts/aliases/is_windows_os"
	isx64 "github.com/c2micro/c2mcli/internal/scripts/aliases/is_x64"
	isx86 "github.com/c2micro/c2mcli/internal/scripts/aliases/is_x86"
	merror "github.com/c2micro/c2mcli/internal/scripts/aliases/m_error"
	minfo "github.com/c2micro/c2mcli/internal/scripts/aliases/m_info"
	mnotify "github.com/c2micro/c2mcli/internal/scripts/aliases/m_notify"
	mwarning "github.com/c2micro/c2mcli/internal/scripts/aliases/m_warning"
	tcancel "github.com/c2micro/c2mcli/internal/scripts/aliases/t_cancel"
	"github.com/c2micro/mlan/pkg/engine/object"
	"github.com/c2micro/mlan/pkg/engine/storage"
	"github.com/c2micro/mlan/pkg/engine/types"
	mlanUtils "github.com/c2micro/mlan/pkg/engine/utils"
	"github.com/c2micro/mlan/pkg/engine/visitor"
	"github.com/go-faster/errors"
)

// регистрация API для интеграции MLAN с C2
func registerApi() {
	// alias: регистрация нового алиаса
	storage.UserFunctions[alias.GetApiName()] = object.NewUserFunc(alias.GetApiName(), alias.UserAlias)
	// m_notify: сообщение с типом NOTIFY
	storage.UserFunctions[mnotify.GetApiName()] = object.NewUserFunc(mnotify.GetApiName(), mnotify.UserMessageNotify)
	// m_info: сообщение с типом INFO
	storage.UserFunctions[minfo.GetApiName()] = object.NewUserFunc(minfo.GetApiName(), minfo.UserMessageInfo)
	// m_warning: сообщение с типом WARNING
	storage.UserFunctions[mwarning.GetApiName()] = object.NewUserFunc(mwarning.GetApiName(), mwarning.UserMessageWarning)
	// m_error: сообщение с типом ERROR
	storage.UserFunctions[merror.GetApiName()] = object.NewUserFunc(merror.GetApiName(), merror.UserMessageError)
	// b_sleep: изменение параметров sleep/jitter бикона
	storage.UserFunctions[bsleep.GetApiName()] = object.NewUserFunc(bsleep.GetApiName(), bsleep.UserBeaconSleep)
	// b_ls: получение листинга директорий
	storage.UserFunctions[bls.GetApiName()] = object.NewUserFunc(bls.GetApiName(), bls.UserBeaconLs)
	// b_pwd: получение текущей директории (CWD)
	storage.UserFunctions[bpwd.GetApiName()] = object.NewUserFunc(bpwd.GetApiName(), bpwd.UserBeaconPwd)
	// b_cd: изменение рабочей директории
	storage.UserFunctions[bcd.GetApiName()] = object.NewUserFunc(bcd.GetApiName(), bcd.UserBeaconCd)
	// b_whoami: получение текущего пользователя/его грантов
	storage.UserFunctions[bwhoami.GetApiName()] = object.NewUserFunc(bwhoami.GetApiName(), bwhoami.UserBeaconWhoami)
	// b_ps: листинг процессов
	storage.UserFunctions[bps.GetApiName()] = object.NewUserFunc(bps.GetApiName(), bps.UserBeaconPs)
	// b_cat: вывод файла
	storage.UserFunctions[bcat.GetApiName()] = object.NewUserFunc(bcat.GetApiName(), bcat.UserBeaconCat)
	// b_exec: выполнение исполняемого файла
	storage.UserFunctions[bexec.GetApiName()] = object.NewUserFunc(bexec.GetApiName(), bexec.UserBeaconExec)
	// b_cp: копирование файлов/директорий
	storage.UserFunctions[bcp.GetApiName()] = object.NewUserFunc(bcp.GetApiName(), bcp.UserBeaconCp)
	// b_jobs: получение активных задач на биконе
	storage.UserFunctions[bjobs.GetApiName()] = object.NewUserFunc(bjobs.GetApiName(), bjobs.UserBeaconJobs)
	// b_jobkill: килл активной задачи на биконе
	storage.UserFunctions[bjobkill.GetApiName()] = object.NewUserFunc(bjobkill.GetApiName(), bjobkill.UserBeaconJobkill)
	// b_kill: килл процессса на машине
	storage.UserFunctions[bkill.GetApiName()] = object.NewUserFunc(bkill.GetApiName(), bkill.UserBeaconKill)
	// b_mv: перемещение файлов/директорий
	storage.UserFunctions[bmv.GetApiName()] = object.NewUserFunc(bmv.GetApiName(), bmv.UserBeaconMv)
	// b_mkdir: создание директории
	storage.UserFunctions[bmkdir.GetApiName()] = object.NewUserFunc(bmkdir.GetApiName(), bmkdir.UserBeaconMkdir)
	// b_exec_assembly: исполнение .NET в памяти
	storage.UserFunctions[bexecassembly.GetApiName()] = object.NewUserFunc(bexecassembly.GetApiName(), bexecassembly.UserBeaconExecuteAssembly)
	// b_download: скачивание файла
	storage.UserFunctions[bdownload.GetApiName()] = object.NewUserFunc(bdownload.GetApiName(), bdownload.UserBeaconDownload)
	// b_upload: загрузка файла на хост
	storage.UserFunctions[bupload.GetApiName()] = object.NewUserFunc(bupload.GetApiName(), bupload.UserBeaconUpload)
	// b_pause: одноразовый слип на биконе
	storage.UserFunctions[bpause.GetApiName()] = object.NewUserFunc(bpause.GetApiName(), bpause.UserBeaconPause)
	// b_destruct: самоликвидация бикона
	storage.UserFunctions[bdestruct.GetApiName()] = object.NewUserFunc(bdestruct.GetApiName(), bdestruct.UserBeaconDestruct)
	// b_exec_detach: выполнение исполняемого файла с детачем
	storage.UserFunctions[bexecdetach.GetApiName()] = object.NewUserFunc(bexec.GetApiName(), bexecdetach.UserBeaconExecDetach)
	// b_shell: выполнение shell команды
	storage.UserFunctions[bshell.GetApiName()] = object.NewUserFunc(bshell.GetApiName(), bshell.UserBeaconShell)
	// b_ppid: спуфинг parent PID
	storage.UserFunctions[bppid.GetApiName()] = object.NewUserFunc(bppid.GetApiName(), bppid.UserBeaconPpid)
	// b_exit: остановка бикона
	storage.UserFunctions[bexit.GetApiName()] = object.NewUserFunc(bexit.GetApiName(), bexit.UserBeaconExit)
	// t_cancel: отмена всех тасок в статусе NEW от оператора
	storage.UserFunctions[tcancel.GetApiName()] = object.NewUserFunc(tcancel.GetApiName(), tcancel.UserBeaconCancel)
	// is_windows: запущен ли бикон на windows
	storage.UserFunctions[iswindows.GetApiName()] = object.NewUserFunc(iswindows.GetApiName(), iswindows.UserIsWindows)
	// is_linux: запущен ли бикон на linux
	storage.UserFunctions[islinux.GetApiName()] = object.NewUserFunc(islinux.GetApiName(), islinux.UserIsLinux)
	// is_macos: запущен ли бикон на macos
	storage.UserFunctions[ismacos.GetApiName()] = object.NewUserFunc(ismacos.GetApiName(), ismacos.UserIsMacos)
	// is_x64: является ли архитектура процесса x64 (amd64)
	storage.UserFunctions[isx64.GetApiName()] = object.NewUserFunc(isx64.GetApiName(), isx64.UserIsX64)
	// is_x86: является ли архитектура процесса x86
	storage.UserFunctions[isx86.GetApiName()] = object.NewUserFunc(isx86.GetApiName(), isx86.UserIsX86)
	// is_arm64: является ли архитектура процесса arm64
	storage.UserFunctions[isarm64.GetApiName()] = object.NewUserFunc(isarm64.GetApiName(), isarm64.UserIsArm64)
	// is_arm32: является ли архитектура процесса arm32
	storage.UserFunctions[isarm32.GetApiName()] = object.NewUserFunc(isarm32.GetApiName(), isarm32.UserIsArm32)
}

var (
	//go:embed builtin/*.c2m
	builtinScriptsFS embed.FS
)

// регистрация встроенных скриптов с базовыми командами
func registerBuiltin() error {
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
		tree, err := mlanUtils.CreateAST(string(data))
		if err != nil {
			return errors.Wrap(err, v.Name())
		}
		// проходим по дереву
		visitor := visitor.NewVisitor()
		if res := visitor.Visit(tree); res != types.Success {
			return errors.Wrapf(visitor.GetError(), "evaluation %s", v.Name())
		}
	}
	return nil
}
