package visitor

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/c2micro/mlan/pkg/constants"
	"github.com/c2micro/mlan/pkg/engine/builtin"
	"github.com/c2micro/mlan/pkg/engine/object"
	"github.com/c2micro/mlan/pkg/engine/scope"
	"github.com/c2micro/mlan/pkg/engine/storage"
	"github.com/c2micro/mlan/pkg/engine/types"
	"github.com/c2micro/mlan/pkg/engine/utils"
	"github.com/c2micro/mlan/pkg/parser"
	"github.com/go-faster/errors"
)

func NewVisitor() *Visitor {
	return &Visitor{
		err: nil,
	}
}

type Visitor struct {
	parser.BaseMlanVisitor
	// line number of executed script
	line int
	// visitor error
	err error
}

func (v *Visitor) SetError(err error) {
	v.err = errors.Wrap(err, fmt.Sprintf("line %d", v.line))
}

func (v *Visitor) GetError() error {
	return v.err
}

func (v *Visitor) GetLine() int {
	return v.line
}

/*
func (v *BaseMlanVisitor) VisitAssignMul(ctx *AssignMulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitAssignDiv(ctx *AssignDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitAssignMod(ctx *AssignModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitAssignPow(ctx *AssignPowContext) interface{} {
	return v.VisitChildren(ctx)
}

*/

// loop entry-point for visitor
func (v *Visitor) Visit(tree antlr.ParseTree) any {
	if retValue != nil {
		// exit from AST in case of return value existance
		return types.Success
	}

	// set line number
	v.line = tree.(antlr.ParserRuleContext).GetStart().GetLine()

	if scope.CurrentScope.Depth() > constants.MaxScopeDepth {
		v.err = fmt.Errorf("max depth reached")
		return types.Failure
	}

	switch val := tree.(type) {
	case *parser.ProgContext:
		return v.VisitProg(val)
	case *parser.StmtContext:
		return v.VisitStmt(val)
	case *parser.AssignRegularContext:
		return v.VisitAssignRegular(val)
	case *parser.ExpBoolContext:
		return v.VisitExpBool(val)
	case *parser.IdentifierFnInvokeContext:
		return v.VisitIdentifierFnInvoke(val)
	case *parser.ExpIdentifierContext:
		return v.VisitExpIdentifier(val)
	case *parser.ExpIntegerContext:
		return v.VisitExpInteger(val)
	case *parser.ExpIntegerHexContext:
		return v.VisitExpIntegerHex(val)
	case *parser.ExpNullContext:
		return v.VisitExpNull(val)
	case *parser.ExpFloatContext:
		return v.VisitExpFloat(val)
	case *parser.ExpStringContext:
		return v.VisitExpString(val)
	case *parser.ExpListContext:
		return v.VisitExpList(val)
	case *parser.ListContext:
		return v.VisitList(val)
	case *parser.ExpDictContext:
		return v.VisitExpDict(val)
	case *parser.ExpLogicalNotContext:
		return v.VisitExpLogicalNot(val)
	case *parser.ExpLogicalOrContext:
		return v.VisitExpLogicalOr(val)
	case *parser.ExpLogicalAndContext:
		return v.VisitExpLogicalAnd(val)
	case *parser.ExpParenthesesContext:
		return v.VisitExpParentheses(val)
	case *parser.ExpEqualContext:
		return v.VisitExpEqual(val)
	case *parser.ExpNegContext:
		return v.VisitExpNeg(val)
	case *parser.ExpComparisonContext:
		return v.VisitExpComparison(val)
	case *parser.ExpSumSubContext:
		return v.VisitExpSumSub(val)
	case *parser.ExpPowContext:
		return v.VisitExpPow(val)
	case *parser.ExpMulDivModContext:
		return v.VisitExpMulDivMod(val)
	case *parser.ExpXorContext:
		return v.VisitExpXor(val)
	case *parser.ForStmtContext:
		return v.VisitForStmt(val)
	case *parser.BreakStmtContext:
		return v.VisitBreakStmt(val)
	case *parser.ContinueStmtContext:
		return v.VisitContinueStmt(val)
	case *parser.WhileStmtContext:
		return v.VisitWhileStmt(val)
	case *parser.IfStmtContext:
		return v.VisitIfStmt(val)
	case *parser.IfBlockStmtContext:
		return v.VisitIfBlockStmt(val)
	case *parser.ElseBlockStmtContext:
		return v.VisitElseBlockStmt(val)
	case *parser.ElifBlockStmtContext:
		return v.VisitElifBlockStmt(val)
	case *parser.ExpFnInvokeContext:
		return v.VisitExpFnInvoke(val)
	case *parser.ReturnStmtContext:
		return v.VisitReturnStmt(val)
	case *parser.ExpIdxContext:
		return v.VisitExpIdx(val)
	case *parser.AssignIdxRegularContext:
		return v.VisitAssignIdxRegular(val)
	case *parser.ExpCsInvokeContext:
		return v.VisitExpCsInvoke(val)
	case *parser.IdentifierCsInvokeContext:
		return v.VisitIdentifierCsInvoke(val)
	case *parser.IncludeContext:
		return v.VisitInclude(val)
	case *parser.FnContext:
		return v.VisitFn(val)
	case *parser.FnBodyContext:
		return v.VisitFnBody(val)
	case *parser.AssignSumContext:
		return v.VisitAssignSum(val)
	case *parser.AssignSubContext:
		return v.VisitAssignSub(val)
	case *parser.ExpMethodInvokeContext:
		return v.VisitExpMethodInvoke(val)
	case *parser.IdentifierMethodInvokeContext:
		return v.VisitIdentifierMethodInvoke(val)
	case *parser.ExpCsContext:
		return v.VisitExpCs(val)
	case *parser.ClosureContext:
		return v.VisitClosure(val)
	case *parser.IdxContext:
		return v.VisitIdx(val)
	case *parser.DictContext:
		return v.VisitDict(val)
	default:
		v.SetError(fmt.Errorf("unknown context '%s' on visit", reflect.TypeOf(val).String()))
		return types.Failure
	}
}

func (v *Visitor) VisitProg(ctx *parser.ProgContext) any {
	// резолв инклудов
	for _, item := range ctx.AllInclude() {
		tree, ok := v.Visit(item).(antlr.ParseTree)
		if !ok {
			return types.Failure
		}
		if res := v.Visit(tree); res != types.Success {
			return types.Failure
		}
	}

	// регистрация функций
	for _, item := range ctx.AllFn() {
		if ok := v.Visit(item).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	// выполнение стейтментов
	for _, item := range ctx.AllStmt() {
		if ok := v.Visit(item).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	return types.Success
}

func (v *Visitor) VisitStmt(ctx *parser.StmtContext) any {
	// присвоение
	if ctx.Assignment() != nil {
		if ok := v.Visit(ctx.Assignment()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	// if/elif/else
	if ctx.IfStmt() != nil {
		scope.CurrentScope = scope.NewScope(
			scope.CurrentScope,
			scope.CurrentScope.Depth()+1,
			false,
			false,
			make(map[string]object.Object),
		)
		// стейтменты внутри if блока
		if ok := v.Visit(ctx.IfStmt()).(types.VisitResultType); !ok {
			scope.CurrentScope = scope.CurrentScope.Parent()
			return types.Failure
		}
		// возвращаем скоуп
		scope.CurrentScope = scope.CurrentScope.Parent()
	}

	// while цикл
	if ctx.WhileStmt() != nil {
		scope.CurrentScope = scope.NewScope(
			scope.CurrentScope,
			scope.CurrentScope.Depth()+1,
			false,
			true,
			make(map[string]object.Object),
		)
		// стейтменты внутри цикла
		if ok := v.Visit(ctx.WhileStmt()).(types.VisitResultType); !ok {
			scope.CurrentScope = scope.CurrentScope.Parent()
			return types.Failure
		}
		// возвращаем скоуп
		scope.CurrentScope = scope.CurrentScope.Parent()
	}

	// for цикл
	if ctx.ForStmt() != nil {
		scope.CurrentScope = scope.NewScope(
			scope.CurrentScope,
			scope.CurrentScope.Depth()+1,
			false,
			true,
			make(map[string]object.Object),
		)
		// стейтменты внутри цикла
		if ok := v.Visit(ctx.ForStmt()).(types.VisitResultType); !ok {
			scope.CurrentScope = scope.CurrentScope.Parent()
			return types.Failure
		}
		// возвращаем скоуп
		scope.CurrentScope = scope.CurrentScope.Parent()
	}

	// вызов функции
	if ctx.FnInvoke() != nil {
		if res, ok := v.Visit(ctx.FnInvoke()).(types.VisitResultType); ok {
			if !res {
				return types.Failure
			}
		}
	}

	// вызов кложура
	if ctx.CsInvoke() != nil {
		if res, ok := v.Visit(ctx.CsInvoke()).(types.VisitResultType); ok {
			if !res {
				return types.Failure
			}
		}
	}

	// вызов метода
	if ctx.MethodInvoke() != nil {
		if res, ok := v.Visit(ctx.MethodInvoke()).(types.VisitResultType); ok {
			if !res {
				return types.Failure
			}
		}
	}

	// break для цикла
	if ctx.BreakStmt() != nil {
		if ok := v.Visit(ctx.BreakStmt()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	// continue для цикла
	if ctx.ContinueStmt() != nil {
		if ok := v.Visit(ctx.ContinueStmt()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	// return
	if ctx.ReturnStmt() != nil {
		if ok := v.Visit(ctx.ReturnStmt()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	return types.Success
}

func (v *Visitor) VisitAssignRegular(ctx *parser.AssignRegularContext) any {
	// получение значения выражения
	val, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}

	// сохранение объекта
	scope.CurrentScope.Put(ctx.GetName().GetText(), val)

	return types.Success
}

func (v *Visitor) VisitExpBool(ctx *parser.ExpBoolContext) any {
	switch ctx.GetText() {
	case "true":
		return object.NewBool(true)
	case "false":
		return object.NewBool(false)
	}
	v.SetError(fmt.Errorf("unable get bool from '%s'", ctx.GetText()))
	return types.Failure
}

func (v *Visitor) VisitIdentifierFnInvoke(ctx *parser.IdentifierFnInvokeContext) any {
	var params []object.Object

	for _, item := range ctx.AllExp() {
		// собираем аргументы для функции
		res, ok := v.Visit(item).(object.Object)
		if !ok {
			return types.Failure
		}
		params = append(params, res)
	}

	return v.invokeFunc(ctx.GetName().GetText(), params...)
}

func (v *Visitor) VisitExpIdentifier(ctx *parser.ExpIdentifierContext) any {
	val := scope.CurrentScope.Get(ctx.GetText(), true)
	if val == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetText()))
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpInteger(ctx *parser.ExpIntegerContext) any {
	val, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		v.SetError(fmt.Errorf("unable get int from '%s'", ctx.GetText()))
		return types.Failure
	}
	return object.NewInt(val)
}

func (v *Visitor) VisitExpIntegerHex(ctx *parser.ExpIntegerHexContext) any {
	// парсим в int64
	val, err := strconv.ParseInt(utils.StripHexPrefix(ctx.GetText()), 16, 64)
	if err != nil {
		v.SetError(fmt.Errorf("unable get int from '%s'", ctx.GetText()))
		return types.Failure
	}
	return object.NewInt(val)
}

func (v *Visitor) VisitExpNull(_ *parser.ExpNullContext) any {
	return object.NewNull()
}

func (v *Visitor) VisitExpFloat(ctx *parser.ExpFloatContext) any {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		v.SetError(fmt.Errorf("unable get float from '%s'", ctx.GetText()))
		return types.Failure
	}
	return object.NewFloat(val)
}

func (v *Visitor) VisitExpString(ctx *parser.ExpStringContext) any {
	val, err := strconv.Unquote(ctx.GetText())
	if err != nil {
		v.SetError(fmt.Errorf("unable get str from '%s'", ctx.GetText()))
		return types.Failure
	}
	return object.NewStr(val)
}

func (v *Visitor) VisitExpList(ctx *parser.ExpListContext) any {
	return v.Visit(ctx.List())
}

func (v *Visitor) VisitList(ctx *parser.ListContext) any {
	var list []object.Object
	for _, item := range ctx.AllExp() {
		val, ok := v.Visit(item).(object.Object)
		if !ok {
			return types.Failure
		}
		list = append(list, val)
	}
	return object.NewList(list)
}

func (v *Visitor) VisitExpDict(ctx *parser.ExpDictContext) any {
	dict, ok := v.Visit(ctx.Dict()).(map[string]object.Object)
	if !ok {
		return types.Failure
	}
	return object.NewDict(dict)
}

func (v *Visitor) VisitDict(ctx *parser.DictContext) any {
	dict := make(map[string]object.Object)
	for _, item := range ctx.AllDictUnit() {
		// получение ключа
		key, ok := v.Visit(item.AllExp()[0]).(object.Object)
		if !ok {
			return types.Failure
		}

		// проверяем что тип ялвяется строкой
		switch key.(type) {
		case *object.Str:
		default:
			v.SetError(fmt.Errorf("key of dict must be str"))
			return types.Failure
		}

		// проверка на дубликат
		if _, ok = dict[key.GetValue().(string)]; ok {
			v.SetError(fmt.Errorf("duplicated key '%s' in dict", key.GetValue().(string)))
			return types.Failure
		}

		// получение значения
		val, ok := v.Visit(item.AllExp()[1]).(object.Object)
		if !ok {
			return types.Failure
		}
		dict[key.GetValue().(string)] = val
	}
	return dict
}

func (v *Visitor) VisitExpLogicalNot(ctx *parser.ExpLogicalNotContext) any {
	val, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := val.UnaryOp(parser.MlanParserNot)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: !%s", val.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpLogicalOr(ctx *parser.ExpLogicalOrContext) any {
	lhs, ok := v.Visit(ctx.Exp(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(parser.MlanLexerOr, rhs)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s || %s", lhs.TypeName(), rhs.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpLogicalAnd(ctx *parser.ExpLogicalAndContext) any {
	lhs, ok := v.Visit(ctx.Exp(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(parser.MlanLexerAnd, rhs)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s && %s", lhs.TypeName(), rhs.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpParentheses(ctx *parser.ExpParenthesesContext) any {
	return v.Visit(ctx.Exp())
}

func (v *Visitor) VisitExpEqual(ctx *parser.ExpEqualContext) any {
	lhs, ok := v.Visit(ctx.Exp(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(ctx.GetOp().GetTokenType(), rhs)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s %s %s", lhs.TypeName(), ctx.GetOp().GetText(), rhs.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpNeg(ctx *parser.ExpNegContext) any {
	val, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := val.UnaryOp(parser.MlanLexerSubtract)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: -%s", val.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpComparison(ctx *parser.ExpComparisonContext) any {
	lhs, ok := v.Visit(ctx.Exp(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(ctx.GetOp().GetTokenType(), rhs)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s %s %s", lhs.TypeName(), ctx.GetOp().GetText(), rhs.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpSumSub(ctx *parser.ExpSumSubContext) any {
	lhs, ok := v.Visit(ctx.Exp(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(ctx.GetOp().GetTokenType(), rhs)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s %s %s", lhs.TypeName(), ctx.GetOp().GetText(), rhs.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpPow(ctx *parser.ExpPowContext) any {
	lhs, ok := v.Visit(ctx.Exp(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(parser.MlanLexerPow, rhs)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s ** %s", lhs.TypeName(), rhs.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpMulDivMod(ctx *parser.ExpMulDivModContext) any {
	lhs, ok := v.Visit(ctx.Exp(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(ctx.GetOp().GetTokenType(), rhs)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s %s %s", lhs.TypeName(), ctx.GetOp().GetText(), rhs.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpXor(ctx *parser.ExpXorContext) any {
	lhs, ok := v.Visit(ctx.Exp(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Exp(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(parser.MlanParserXor, rhs)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s ^ %s", lhs.TypeName(), rhs.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitForStmt(ctx *parser.ForStmtContext) any {
	// левая часть присвоения (изначальное значение счетчика)
	if ok := v.Visit(ctx.Assignment(0)).(types.VisitResultType); !ok {
		return types.Failure
	}

	for {
		// условие выхода из цикла
		res, ok := v.Visit(ctx.Exp()).(object.Object)
		if !ok {
			if v.GetError() != nil {
				v.SetError(fmt.Errorf("invalid expression for loop: %s", v.GetError()))
			} else {
				v.SetError(fmt.Errorf("invalid expression for loop"))
			}
			return types.Failure
		}

		// проверяем, что тип результата - булева
		if _, ok = res.(*object.Bool); !ok {
			v.SetError(fmt.Errorf("non-bool (%s) expression in for loop", res.TypeName()))
			return types.Failure
		}

		// проверяем, что условие выполняется
		if !res.(*object.Bool).GetValue().(bool) {
			// если false -> выходим из цикла
			break
		}

		// выполняем все выражения внутри блока цикла
		for _, item := range ctx.AllStmt() {
			// обработка break
			if scope.CurrentScope.IsBreak() {
				scope.CurrentScope.UnsetLoopBreak()
				return types.Success
			}
			// обработка continue
			if scope.CurrentScope.IsContinue() {
				scope.CurrentScope.UnsetLoopContinue()
				break
			}
			if ok := v.Visit(item).(types.VisitResultType); !ok {
				return types.Failure
			}
			// обработка return
			if retValue != nil {
				return types.Success
			}
		}

		// изменяем счетчик (правое присвоение)
		if ok := v.Visit(ctx.Assignment(1)).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	return types.Success
}

func (v *Visitor) VisitBreakStmt(ctx *parser.BreakStmtContext) any {
	if scope.CurrentScope.IsInLoop() {
		scope.CurrentScope.SetLoopBreak()
		return types.Success
	}
	v.SetError(fmt.Errorf("break outside of loop"))
	return types.Failure
}

func (v *Visitor) VisitContinueStmt(ctx *parser.ContinueStmtContext) any {
	if scope.CurrentScope.IsInLoop() {
		scope.CurrentScope.SetLoopContinue()
		return types.Success
	}
	v.SetError(fmt.Errorf("continue outside of loop"))
	return types.Failure
}

func (v *Visitor) VisitWhileStmt(ctx *parser.WhileStmtContext) any {
	for {
		// условие выхода из цикла
		res, ok := v.Visit(ctx.Exp()).(object.Object)
		if !ok {
			return types.Failure
		}

		// проверяем, что условие булева
		if _, ok = res.(*object.Bool); !ok {
			v.SetError(fmt.Errorf("non-bool (%s) expression in for loop", res.TypeName()))
			return types.Failure
		}

		// проверяем, что условие выполняется
		if !res.(*object.Bool).GetValue().(bool) {
			// если false -> выходим из цикла
			break
		}

		for _, item := range ctx.AllStmt() {
			if ok := v.Visit(item).(types.VisitResultType); !ok {
				return types.Failure
			}
			// обработка break
			if scope.CurrentScope.IsBreak() {
				scope.CurrentScope.UnsetLoopBreak()
				return types.Success
			}
			// обработка continue
			if scope.CurrentScope.IsContinue() {
				scope.CurrentScope.UnsetLoopContinue()
				break
			}
			// обработка return
			if retValue != nil {
				return types.Success
			}
		}
	}

	return types.Success
}

func (v *Visitor) VisitIfStmt(ctx *parser.IfStmtContext) any {
	// if
	if ctx.IfBlock() != nil {
		// результат if блока
		res, ok := v.Visit(ctx.IfBlock()).(object.Object)
		if !ok {
			return types.Failure
		}
		// если результат true -> выходим
		if res.(*object.Bool).GetValue().(bool) {
			return types.Success
		}
	}

	// elif
	for _, item := range ctx.AllElifBlock() {
		// получаем результат elif блока
		res, ok := v.Visit(item).(object.Object)
		if !ok {
			return types.Failure
		}
		// если результат true -> выходим
		if res.(*object.Bool).GetValue().(bool) {
			return types.Success
		}
	}

	// else
	if ctx.ElseBlock() != nil {
		if ok := v.Visit(ctx.ElseBlock()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	return types.Success
}

func (v *Visitor) VisitIfBlockStmt(ctx *parser.IfBlockStmtContext) any {
	res, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	// кастим в бул
	res, err := builtin.Bool(res)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}

	if res.(*object.Bool).GetValue().(bool) {
		// если булевый объект true -> выполняем блок
		for _, item := range ctx.AllStmt() {
			if ok := v.Visit(item).(types.VisitResultType); !ok {
				return types.Failure
			}
		}
	}

	return res
}

func (v *Visitor) VisitElseBlockStmt(ctx *parser.ElseBlockStmtContext) any {
	for _, item := range ctx.AllStmt() {
		if ok := v.Visit(item).(types.VisitResultType); !ok {
			return types.Failure
		}
	}
	return types.Success
}

func (v *Visitor) VisitElifBlockStmt(ctx *parser.ElifBlockStmtContext) any {
	res, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	// кастим в бул
	res, err := builtin.Bool(res)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}

	if res.(*object.Bool).GetValue().(bool) {
		// если булевый объект true -> выполняем блок
		for _, item := range ctx.AllStmt() {
			if ok := v.Visit(item).(types.VisitResultType); !ok {
				return types.Failure
			}
		}
	}

	return res
}

func (v *Visitor) VisitFn(ctx *parser.FnContext) any {
	// проверяем, что функции не существует
	if _, ok := storage.NativeFunctions[ctx.GetName().GetText()]; ok {
		v.SetError(fmt.Errorf("function '%s' already defined", ctx.GetName().GetText()))
		return types.Failure
	}

	// получение объекта нативной функции
	fn, ok := v.Visit(ctx.FnBody()).(*object.NativeFunc)
	if !ok {
		return types.Failure
	}

	// установка имени функции
	fn.SetName(ctx.GetName().GetText())

	// сохранение функции
	storage.NativeFunctions[ctx.GetName().GetText()] = fn

	return types.Success
}

func (v *Visitor) VisitFnBody(ctx *parser.FnBodyContext) any {
	// аргументы
	args := make([]string, 0)
	if ctx.FnParams() != nil {
		for _, item := range ctx.FnParams().AllIdentifier() {
			args = append(args, item.GetText())
		}
	}
	// сохдание объекта функции
	return object.NewNativeFunc(
		args,
		len(args),
		ctx.AllStmt(),
	)
}

func (v *Visitor) VisitExpFnInvoke(ctx *parser.ExpFnInvokeContext) any {
	res, ok := v.Visit(ctx.FnInvoke()).(object.Object)
	if !ok {
		// значит вернулся статус
		if v.GetError() == nil {
			return object.NewNull()
		}
		return types.Failure
	}
	return res
}

func (v *Visitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) any {
	if !scope.CurrentScope.IsInFunc() {
		v.SetError(fmt.Errorf("return outside of function"))
		return types.Failure
	}
	// получаем результат return
	if ctx.Exp() != nil {
		res, ok := v.Visit(ctx.Exp()).(object.Object)
		if !ok {
			return types.Failure
		}
		retValue = res
		return types.Success
	} else {
		retValue = object.NewNull()
		return types.Success
	}
}

func (v *Visitor) VisitExpIdx(ctx *parser.ExpIdxContext) any {
	lhs, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	idx, ok := v.Visit(ctx.Idx()).(object.Object)
	if !ok {
		return types.Failure
	}
	res, err := lhs.IndexGet(idx)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.SetError(fmt.Errorf("unsupported expression: %s[%s]", lhs.TypeName(), idx.TypeName()))
		} else {
			v.SetError(err)
		}
		return types.Failure
	}
	return res
}

func (v *Visitor) VisitIdx(ctx *parser.IdxContext) any {
	return v.Visit(ctx.Exp())
}

func (v *Visitor) VisitAssignIdxRegular(ctx *parser.AssignIdxRegularContext) any {
	// забираем из скоупа переменную
	lhs := scope.CurrentScope.Get(ctx.GetName().GetText(), true)
	if lhs == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetName().GetText()))
		return types.Failure
	}
	// получаем индекс
	idx, ok := v.Visit(ctx.Idx().Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	// получаем значение
	val, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	// проставляем значение по индексу
	if err := lhs.IndexSet(idx, val); err != nil {
		v.SetError(err)
		return types.Failure
	}
	return types.Success
}

func (v *Visitor) VisitExpCsInvoke(ctx *parser.ExpCsInvokeContext) any {
	res, ok := v.Visit(ctx.CsInvoke()).(object.Object)
	if !ok {
		// значит вернулся статус
		if v.GetError() == nil {
			return object.NewNull()
		}
		return types.Failure
	}
	return res
}

func (v *Visitor) VisitIdentifierCsInvoke(ctx *parser.IdentifierCsInvokeContext) any {
	var params []object.Object

	for _, item := range ctx.AllExp() {
		// собираем аргументы для функции
		res, ok := v.Visit(item).(object.Object)
		if !ok {
			return types.Failure
		}
		params = append(params, res)
	}

	return v.invokeClosureFunc(ctx.GetName().GetText(), params...)
}

func (v *Visitor) VisitInclude(ctx *parser.IncludeContext) any {
	val, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	path, ok := val.(*object.Str)
	if !ok {
		v.SetError(fmt.Errorf("invalid argument of type '%s'", val.TypeName()))
		return types.Failure
	}
	data, err := os.ReadFile(path.GetValue().(string))
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	tree, err := utils.CreateAST(string(data))
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	return tree
}

func (v *Visitor) VisitExpCs(ctx *parser.ExpCsContext) any {
	return v.Visit(ctx.Closure())
}

func (v *Visitor) VisitClosure(ctx *parser.ClosureContext) interface{} {
	return v.Visit(ctx.FnBody())
}

func (v *Visitor) VisitAssignSum(ctx *parser.AssignSumContext) any {
	// значение переменной из скоупа
	val := scope.CurrentScope.Get(ctx.GetName().GetText(), true)
	if val == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetName().GetText()))
		return types.Failure
	}
	// результат выражения справа
	rhs, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	// итоговый объект
	res, err := val.BinaryOp(parser.MlanLexerAssSum, rhs)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// сохраняем новый объект
	scope.CurrentScope.Put(ctx.GetName().GetText(), res)

	return types.Success
}

func (v *Visitor) VisitAssignSub(ctx *parser.AssignSubContext) any {
	// значение переменной из скоупа
	val := scope.CurrentScope.Get(ctx.GetName().GetText(), true)
	if val == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetName().GetText()))
		return types.Failure
	}
	// результат выражения справа
	rhs, ok := v.Visit(ctx.Exp()).(object.Object)
	if !ok {
		return types.Failure
	}
	// итоговый объект
	res, err := val.BinaryOp(parser.MlanLexerAssSub, rhs)
	if err != nil {
		v.SetError(err)
		return types.Failure
	}
	// сохраняем новый объект
	scope.CurrentScope.Put(ctx.GetName().GetText(), res)

	return types.Success
}

func (v *Visitor) VisitExpMethodInvoke(ctx *parser.ExpMethodInvokeContext) interface{} {
	return v.Visit(ctx.MethodInvoke())
}

func (v *Visitor) VisitIdentifierMethodInvoke(ctx *parser.IdentifierMethodInvokeContext) interface{} {
	// значение переменной из скоупа
	val := scope.CurrentScope.Get(ctx.GetVar_().GetText(), true)
	if val == nil {
		v.SetError(fmt.Errorf("undefined variable '%s'", ctx.GetVar_().GetText()))
		return types.Failure
	}
	// собираем аргументы для метода
	var params []object.Object
	for _, item := range ctx.AllExp() {
		res, ok := v.Visit(item).(object.Object)
		if !ok {
			return types.Failure
		}
		params = append(params, res)
	}
	// вызываем метод
	obj, err := val.MethodCall(ctx.GetName().GetText(), params...)
	if err != nil {
		v.err = err
		return types.Failure
	}
	return obj
}
