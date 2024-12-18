package engine

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/c2micro/mlan/pkg/engine/object"
	"github.com/c2micro/mlan/pkg/engine/scope"
	"github.com/c2micro/mlan/pkg/parser"
	"github.com/c2micro/mlan/pkg/perror"
)

type VisitResultType bool

const (
	Success VisitResultType = true
	Failure VisitResultType = false
)

// init инициализация рантайма для языка
func init() {
	Init()
}

// инициализация движка
func Init() {
	// инициализация глобального скоупа
	scope.GlobalScope = scope.NewScope(
		nil,
		0,
		false,
		false,
		map[string]object.Object{})
	// установка текущего скоупа
	scope.CurrentScope = scope.GlobalScope
	// регистрация дефолтных функций
	RegisterBuiltinFunctions()
}

// очистка движка и перерегистрация всех необходимых компонентов
func Clear() {
	// зануление хендлеров функций
	BuiltinFunctions = make(map[string]*object.BuiltinFunc)
	UserFunctions = make(map[string]*object.UserFunc)
	NativeFunctions = make(map[string]*object.NativeFunc)
	// зануление рета
	retValue = nil
	// инициализация
	Init()
}

// Evaluate выполнение скрипта с указанием файла
func Evaluate(file string) error {
	// читаем файл
	data, err := ReadFile(file)
	if err != nil {
		return err
	}
	// создание дерева
	tree, err := CreateAST(string(data))
	if err != nil {
		return err
	}

	// создание визитора
	v := NewVisitor()

	// проходим по дереву
	res := v.Visit(tree)
	if res != Success {
		return fmt.Errorf("runtime failure: %v", v.Error)
	}

	return nil
}

// ReadFile вычитываем файл со скриптом
func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// CreateAST создание AST дерева
func CreateAST(data string) (antlr.ParseTree, error) {
	// создание входного стрима
	iStream := antlr.NewInputStream(data)
	// создание лексера
	lexer := parser.NewMlanLexer(iStream)
	// создание стрима из лексера
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// создание парсера
	p := parser.NewMlanParser(stream)
	// error парсер
	errorParser := &perror.Error{}
	// кастомный коллбэк для ошибок
	p.RemoveErrorListeners()
	p.AddErrorListener(errorParser)
	// создание AST древа
	tree := p.Program()
	return tree, errorParser.Err
}
