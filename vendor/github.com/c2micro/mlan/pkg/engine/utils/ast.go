package utils

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/c2micro/mlan/pkg/parser"
	"github.com/c2micro/mlan/pkg/perror"
)

// создание AST дерева
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
	tree := p.Prog()
	return tree, errorParser.Err
}
