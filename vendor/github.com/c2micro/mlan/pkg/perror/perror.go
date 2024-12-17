package perror

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type Error struct {
	antlr.ErrorListener
	Line   int
	Column int
	Err    error
}

func (c *Error) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	c.Column = column
	c.Line = line
	c.Err = fmt.Errorf("%d:%d %s", line, column, msg)
}

func (c *Error) ReportAttemptingFullContext(_ antlr.Parser, _ *antlr.DFA, _ int, _ int, _ *antlr.BitSet, _ *antlr.ATNConfigSet) {
}

func (c *Error) ReportContextSensitivity(_ antlr.Parser, _ *antlr.DFA, _ int, _ int, _ int, _ *antlr.ATNConfigSet) {
}

func (c *Error) ReportAmbiguity(_ antlr.Parser, _ *antlr.DFA, _ int, _ int, _ bool, _ *antlr.BitSet, _ *antlr.ATNConfigSet) {
}
