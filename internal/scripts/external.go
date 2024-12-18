package scripts

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/c2micro/c2mcli/internal/utils"
	"github.com/c2micro/mlan/pkg/engine"
	"github.com/go-faster/errors"
	"github.com/lrita/cmap"
)

var scripts cmap.Map[string, *Script]

type Script struct {
	path    string
	tree    antlr.ParseTree
	addedAt time.Time
}

func (s *Script) GetPath() string {
	return s.path
}

func (s *Script) SetPath(data string) {
	s.path = data
}

func (s *Script) GetTree() antlr.ParseTree {
	return s.tree
}

func (s *Script) SetTree(t antlr.ParseTree) {
	s.tree = t
}

func (s *Script) GetAddedAt() time.Time {
	return s.addedAt
}

func (s *Script) SetAddedAt(t time.Time) {
	s.addedAt = t
}

// получение скриптов списком
func GetScripts() []*Script {
	temp := make([]*Script, 0)
	// сохранение скриптов в массив
	scripts.Range(func(k string, v *Script) bool {
		temp = append(temp, v)
		return true
	})
	// сортировка по дате добавления
	sort.SliceStable(temp, func(i, j int) bool {
		return temp[i].addedAt.Before(temp[j].addedAt)
	})
	return temp
}

// существует ли внешний скрипт в хранлище
func IsExternalScriptExists(path string) bool {
	_, ok := scripts.Load(path)
	return ok
}

func RemoveExternalByPath(path string) error {
	var err error

	// получаем абсолютный путь
	if path, err = utils.GetAbsPath(path); err != nil {
		return fmt.Errorf("unable get absolute path of script: %s", err)
	}

	// получение скрипта из мапы
	script, ok := scripts.Load(path)
	if !ok {
		return fmt.Errorf("script %s not registered", path)
	}

	// удаление скрипта
	scripts.Delete(script.GetPath())
	if err := Rebuild(); err != nil {
		return err
	}
	return nil
}

func ReloadExternalByPath(path string) error {
	var err error

	// получаем абсолютный путь
	if path, err = utils.GetAbsPath(path); err != nil {
		return fmt.Errorf("unable get absolute path of script: %s", err)
	}

	// получение скрипта из мапы
	script, ok := scripts.Load(path)
	if !ok {
		return fmt.Errorf("script %s not registered", path)
	}

	// процессинг скрипта
	temp, err := processExternalScript(script.GetPath())
	if err != nil {
		return err
	}

	// удаляем из мапы
	scripts.Delete(path)
	if err := Rebuild(); err != nil {
		return err
	}

	scripts.Store(path, temp)
	return nil
}

// регистрация внешнего скрипта
func RegisterExternalByPath(path string) error {
	var err error

	// получаем абсолютный путь
	if path, err = utils.GetAbsPath(path); err != nil {
		return fmt.Errorf("unable get absolute path of script: %s", err)
	}

	// проверяем, что скрипт уже зарегистрирован
	if IsExternalScriptExists(path) {
		return fmt.Errorf("script %s already registered. Reload it manually", path)
	}

	// процессинг скрипта
	temp, err := processExternalScript(path)
	if err != nil {
		return err
	}

	// добавление в сторож
	scripts.Store(path, temp)
	return nil
}

func processExternalScript(path string) (*Script, error) {
	temp := &Script{
		path:    path,
		tree:    nil,
		addedAt: time.Now(),
	}

	// читаем файл
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "read script file")
	}

	// строим AST
	temp.tree, err = engine.CreateAST(string(data))
	if err != nil {
		return nil, errors.Wrap(err, "create ast")
	}

	// проход по дереву
	v := engine.NewVisitor()
	if res := v.Visit(temp.tree); res != engine.Success {
		return nil, v.Error
	}

	return temp, nil
}
