package visitor

import (
	"errors"
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
)

// Visitor реализация базового визитора
type Visitor struct {
	parser.BaseMlanVisitor
	Error error
}

func (v *Visitor) LineError(token antlr.Token, err error) {
	v.Error = fmt.Errorf("line %d: %v", token.GetLine(), err)
}

// NewVisitor создание нового визитора
func NewVisitor() *Visitor {
	return &Visitor{
		Error: nil,
	}
}

// Visit зацикленный энтри-поинт для визитора
func (v *Visitor) Visit(tree antlr.ParseTree) any {
	if retValue != nil {
		// выходим из программы, если есть возвращаемое значение
		return types.Success
	}

	if scope.CurrentScope.Depth() > constants.MaxScopeDepth {
		// если вложенность превышает допустимый максимум
		v.Error = fmt.Errorf("max depth reached")
		return types.Failure
	}

	//fmt.Println(reflect.ValueOf(tree).String())
	switch val := tree.(type) {
	case *parser.ProgramContext:
		return v.VisitProgram(val)
	case *parser.BlockContext:
		return v.VisitBlock(val)
	case *parser.StatementContext:
		return v.VisitStatement(val)
	case *parser.AssignmentRegularContext:
		return v.VisitAssignmentRegular(val)
	case *parser.ExpressionBoolContext:
		return v.VisitExpressionBool(val)
	case *parser.IdentifierFunctionInvokeContext:
		return v.VisitIdentifierFunctionInvoke(val)
	case *parser.ExpressionIdentifierContext:
		return v.VisitExpressionIdentifier(val)
	case *parser.ExpressionIntegerContext:
		return v.VisitExpressionInteger(val)
	case *parser.ExpressionIntegerHexContext:
		return v.VisitExpressionIntegerHex(val)
	case *parser.ExpressionNullContext:
		return v.VisitExpressionNull(val)
	case *parser.ExpressionFloatContext:
		return v.VisitExpressionFloat(val)
	case *parser.ExpressionStringContext:
		return v.VisitExpressionString(val)
	case *parser.ExpressionListContext:
		return v.VisitExpressionList(val)
	case *parser.ExpressionDictContext:
		return v.VisitExpressionDict(val)
	case *parser.ExpressionLogicalNotContext:
		return v.VisitExpressionLogicalNot(val)
	case *parser.ExpressionLogicalOrContext:
		return v.VisitExpressionLogicalOr(val)
	case *parser.ExpressionLogicalAndContext:
		return v.VisitExpressionLogicalAnd(val)
	case *parser.ExpressionParenthesesContext:
		return v.VisitExpressionParentheses(val)
	case *parser.ExpressionEqualContext:
		return v.VisitExpressionEqual(val)
	case *parser.ExpressionUnaryNegationContext:
		return v.VisitExpressionUnaryNegation(val)
	case *parser.ExpressionComparisonContext:
		return v.VisitExpressionComparison(val)
	case *parser.ExpressionSumSubContext:
		return v.VisitExpressionSumSub(val)
	case *parser.ExpressionPowContext:
		return v.VisitExpressionPow(val)
	case *parser.ExpressionMulDivModContext:
		return v.VisitExpressionMulDivMod(val)
	case *parser.ExpressionXorContext:
		return v.VisitExpressionXor(val)
	case *parser.ForStatementContext:
		return v.VisitForStatement(val)
	case *parser.AssignmentSumContext:
		return v.VisitAssignmentSum(val)
	case *parser.BreakStatementContext:
		return v.VisitBreakStatement(val)
	case *parser.ContinueStatementContext:
		return v.VisitContinueStatement(val)
	case *parser.WhileStatementContext:
		return v.VisitWhileStatement(val)
	case *parser.IfStatementContext:
		return v.VisitIfStatement(val)
	case *parser.IfBlockStatementContext:
		return v.VisitIfBlockStatement(val)
	case *parser.ElseBlockStatementContext:
		return v.VisitElseBlockStatement(val)
	case *parser.ElifBlockStatementContext:
		return v.VisitElifBlockStatement(val)
	case *parser.FunctionDefinitionContext:
		return v.VisitFunctionDefinition(val)
	case *parser.ExpressionFunctionInvokeContext:
		return v.VisitExpressionFunctionInvoke(val)
	case *parser.ReturnStatementContext:
		return v.VisitReturnStatement(val)
	case *parser.ExpressionIndexContext:
		return v.VisitExpressionIndex(val)
	case *parser.AssignmentIndexRegularContext:
		return v.VisitAssignmentIndexRegular(val)
	case *parser.AssignmentClosureContext:
		return v.VisitAssignmentClosure(val)
	case *parser.ExpressionClosureInvokeContext:
		return v.VisitExpressionClosureInvoke(val)
	case *parser.IdentifierClosureInvokeContext:
		return v.VisitIdentifierClosureInvoke(val)
	case *parser.IncludeSubmoduleContext:
		return v.VisitIncludeSubmodule(val)
	case *parser.ExpressionClosureContext:
		return v.VisitExpressionClosure(val)
	default:
		v.Error = fmt.Errorf("unknown context %s on visit", reflect.TypeOf(val).String())
		return types.Failure
	}
}

func (v *Visitor) VisitProgram(ctx *parser.ProgramContext) any {
	return v.Visit(ctx.Block())
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) any {
	// резолв инклудов
	for _, item := range ctx.AllIncludeSubmodule() {
		tree, ok := v.Visit(item).(antlr.ParseTree)
		if !ok {
			return types.Failure
		}
		if res := v.Visit(tree); res != types.Success {
			return types.Failure
		}
	}

	// регистрация функций
	for _, item := range ctx.AllFunctionDefinition() {
		if ok := v.Visit(item).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	// выполнение стейтментов
	for _, item := range ctx.AllStatement() {
		if ok := v.Visit(item).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	return types.Success
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) any {
	// присвоение
	if ctx.Assignment() != nil {
		if ok := v.Visit(ctx.Assignment()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	// if/elif/else
	if ctx.IfStatement() != nil {
		scope.CurrentScope = scope.NewScope(
			scope.CurrentScope,
			scope.CurrentScope.Depth()+1,
			false,
			false,
			make(map[string]object.Object),
		)
		// стейтменты внутри if блока
		if ok := v.Visit(ctx.IfStatement()).(types.VisitResultType); !ok {
			scope.CurrentScope = scope.CurrentScope.Parent()
			return types.Failure
		}
		// возвращаем скоуп
		scope.CurrentScope = scope.CurrentScope.Parent()
	}

	// while цикл
	if ctx.WhileStatement() != nil {
		scope.CurrentScope = scope.NewScope(
			scope.CurrentScope,
			scope.CurrentScope.Depth()+1,
			false,
			true,
			make(map[string]object.Object),
		)
		// стейтменты внутри цикла
		if ok := v.Visit(ctx.WhileStatement()).(types.VisitResultType); !ok {
			scope.CurrentScope = scope.CurrentScope.Parent()
			return types.Failure
		}
		// возвращаем скоуп
		scope.CurrentScope = scope.CurrentScope.Parent()
	}

	// for цикл
	if ctx.ForStatement() != nil {
		scope.CurrentScope = scope.NewScope(
			scope.CurrentScope,
			scope.CurrentScope.Depth()+1,
			false,
			true,
			make(map[string]object.Object),
		)
		// стейтменты внутри цикла
		if ok := v.Visit(ctx.ForStatement()).(types.VisitResultType); !ok {
			scope.CurrentScope = scope.CurrentScope.Parent()
			return types.Failure
		}
		// возвращаем скоуп
		scope.CurrentScope = scope.CurrentScope.Parent()
	}

	// вызов функции
	if ctx.FunctionInvoke() != nil {
		if res, ok := v.Visit(ctx.FunctionInvoke()).(types.VisitResultType); ok {
			if !res {
				return types.Failure
			}
		}
	}

	// вызов кложура
	if ctx.ClosureInvoke() != nil {
		if res, ok := v.Visit(ctx.ClosureInvoke()).(types.VisitResultType); ok {
			if !res {
				return types.Failure
			}
		}
	}

	// break для цикла
	if ctx.BreakStatement() != nil {
		if ok := v.Visit(ctx.BreakStatement()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	// continue для цикла
	if ctx.ContinueStatement() != nil {
		if ok := v.Visit(ctx.ContinueStatement()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	// return
	if ctx.ReturnStatement() != nil {
		if ok := v.Visit(ctx.ReturnStatement()).(types.VisitResultType); !ok {
			return types.Failure
		}
	}

	return types.Success
}

func (v *Visitor) VisitAssignmentRegular(ctx *parser.AssignmentRegularContext) any {
	// получение значения выражения
	val, ok := v.Visit(ctx.Expression()).(object.Object)
	if !ok {
		return types.Failure
	}

	// сохранение объекта
	scope.CurrentScope.Put(ctx.GetVarScalarName().GetText(), val)

	return types.Success
}

func (v *Visitor) VisitExpressionBool(ctx *parser.ExpressionBoolContext) any {
	switch ctx.GetText() {
	case "true":
		return object.NewBool(true)
	case "false":
		return object.NewBool(false)
	}
	v.LineError(ctx.GetStart(), fmt.Errorf("unable get bool from %s", ctx.GetText()))
	return types.Failure
}

func (v *Visitor) VisitIdentifierFunctionInvoke(ctx *parser.IdentifierFunctionInvokeContext) any {
	var params []object.Object

	for _, item := range ctx.AllExpression() {
		// собираем аргументы для функции
		res, ok := v.Visit(item).(object.Object)
		if !ok {
			return types.Failure
		}
		params = append(params, res)
	}

	return v.invokeFunc(ctx.GetStart(), ctx.GetVarFunctionName().GetText(), params...)
}

func (v *Visitor) VisitExpressionIdentifier(ctx *parser.ExpressionIdentifierContext) any {
	val := scope.CurrentScope.Get(ctx.GetText(), true)
	if val == nil {
		v.LineError(ctx.GetStart(), fmt.Errorf("undefined variable '%s'", ctx.GetText()))
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpressionInteger(ctx *parser.ExpressionIntegerContext) any {
	val, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		v.LineError(ctx.GetStart(), fmt.Errorf("unable get int from %s", ctx.GetText()))
		return types.Failure
	}
	return object.NewInt(val)
}

func (v *Visitor) VisitExpressionIntegerHex(ctx *parser.ExpressionIntegerHexContext) any {
	// парсим в int64
	val, err := strconv.ParseInt(utils.StripHexPrefix(ctx.GetText()), 16, 64)
	if err != nil {
		v.LineError(ctx.GetStart(), fmt.Errorf("unable get int from %s", ctx.GetText()))
		return types.Failure
	}
	return object.NewInt(val)
}

func (v *Visitor) VisitExpressionNull(_ *parser.ExpressionNullContext) any {
	return object.NewNull()
}

func (v *Visitor) VisitExpressionFloat(ctx *parser.ExpressionFloatContext) any {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		v.LineError(ctx.GetStart(), fmt.Errorf("unable get float from %s", ctx.GetText()))
		return types.Failure
	}
	return object.NewFloat(val)
}

func (v *Visitor) VisitExpressionString(ctx *parser.ExpressionStringContext) any {
	val, err := strconv.Unquote(ctx.GetText())
	if err != nil {
		v.LineError(ctx.GetStart(), fmt.Errorf("unable get str from %s", ctx.GetText()))
		return types.Failure
	}
	return object.NewStr(val)
}

func (v *Visitor) VisitExpressionList(ctx *parser.ExpressionListContext) any {
	var list []object.Object
	for _, item := range ctx.List().AllExpression() {
		val, ok := v.Visit(item).(object.Object)
		if !ok {
			return types.Failure
		}
		list = append(list, val)
	}
	return object.NewList(list)
}

func (v *Visitor) VisitExpressionDict(ctx *parser.ExpressionDictContext) any {
	dict := make(map[string]object.Object)

	for _, item := range ctx.Dict().AllDictUnit() {
		// получение ключа
		key, ok := v.Visit(item.AllExpression()[0]).(object.Object)
		if !ok {
			return types.Failure
		}

		// проверяем что тип ялвяется строкой
		switch key.(type) {
		case *object.Str:
		default:
			v.LineError(ctx.GetStart(), fmt.Errorf("key of dict must be str"))
			return types.Failure
		}

		// проверка на дубликат
		if _, ok = dict[key.GetValue().(string)]; ok {
			v.LineError(ctx.GetStart(), fmt.Errorf("duplicated key '%s' in dict", key.GetValue().(string)))
			return types.Failure
		}

		// получение значения
		val, ok := v.Visit(item.AllExpression()[1]).(object.Object)
		if !ok {
			return types.Failure
		}
		dict[key.GetValue().(string)] = val
	}
	return object.NewDict(dict)
}

func (v *Visitor) VisitExpressionLogicalNot(ctx *parser.ExpressionLogicalNotContext) any {
	val, ok := v.Visit(ctx.Expression()).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := val.UnaryOp(parser.MlanParserNot)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.LineError(ctx.GetStart(), fmt.Errorf("unsupported expression: !%s", val.TypeName()))
		} else {
			v.LineError(ctx.GetStart(), err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpressionLogicalOr(ctx *parser.ExpressionLogicalOrContext) any {
	lhs, ok := v.Visit(ctx.Expression(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Expression(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(parser.MlanLexerOr, rhs)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.LineError(ctx.GetStart(), fmt.Errorf("unsupported expression: %s || %s", lhs.TypeName(), rhs.TypeName()))
		} else {
			v.LineError(ctx.GetStart(), err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpressionLogicalAnd(ctx *parser.ExpressionLogicalAndContext) any {
	lhs, ok := v.Visit(ctx.Expression(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Expression(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(parser.MlanLexerAnd, rhs)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.LineError(ctx.GetStart(), fmt.Errorf("unsupported expression: %s && %s", lhs.TypeName(), rhs.TypeName()))
		} else {
			v.LineError(ctx.GetStart(), err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpressionParentheses(ctx *parser.ExpressionParenthesesContext) any {
	return v.Visit(ctx.Expression())
}

func (v *Visitor) VisitExpressionEqual(ctx *parser.ExpressionEqualContext) any {
	lhs, ok := v.Visit(ctx.Expression(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Expression(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(ctx.GetOp().GetTokenType(), rhs)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.LineError(ctx.GetStart(), fmt.Errorf("unsupported expression: %s %s %s", lhs.TypeName(), ctx.GetOp().GetText(), rhs.TypeName()))
		} else {
			v.LineError(ctx.GetStart(), err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpressionUnaryNegation(ctx *parser.ExpressionUnaryNegationContext) any {
	val, ok := v.Visit(ctx.Expression()).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := val.UnaryOp(parser.MlanLexerSubtract)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.LineError(ctx.GetStart(), fmt.Errorf("unsupported expression: -%s", val.TypeName()))
		} else {
			v.LineError(ctx.GetStart(), err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpressionComparison(ctx *parser.ExpressionComparisonContext) any {
	lhs, ok := v.Visit(ctx.Expression(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Expression(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(ctx.GetOp().GetTokenType(), rhs)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.LineError(ctx.GetStart(), fmt.Errorf("unsupported expression: %s %s %s", lhs.TypeName(), ctx.GetOp().GetText(), rhs.TypeName()))
		} else {
			v.LineError(ctx.GetStart(), err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpressionSumSub(ctx *parser.ExpressionSumSubContext) any {
	lhs, ok := v.Visit(ctx.Expression(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Expression(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(ctx.GetOp().GetTokenType(), rhs)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.LineError(ctx.GetStart(), fmt.Errorf("unsupported expression: %s %s %s", lhs.TypeName(), ctx.GetOp().GetText(), rhs.TypeName()))
		} else {
			v.LineError(ctx.GetStart(), err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpressionPow(ctx *parser.ExpressionPowContext) any {
	lhs, ok := v.Visit(ctx.Expression(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Expression(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(parser.MlanLexerPow, rhs)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.LineError(ctx.GetStart(), fmt.Errorf("unsupported expression: %s ** %s", lhs.TypeName(), rhs.TypeName()))
		} else {
			v.LineError(ctx.GetStart(), err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpressionMulDivMod(ctx *parser.ExpressionMulDivModContext) any {
	lhs, ok := v.Visit(ctx.Expression(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Expression(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(ctx.GetOp().GetTokenType(), rhs)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.LineError(ctx.GetStart(), fmt.Errorf("unsupported expression: %s %s %s", lhs.TypeName(), ctx.GetOp().GetText(), rhs.TypeName()))
		} else {
			v.LineError(ctx.GetStart(), err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitExpressionXor(ctx *parser.ExpressionXorContext) any {
	lhs, ok := v.Visit(ctx.Expression(0)).(object.Object)
	if !ok {
		return types.Failure
	}
	rhs, ok := v.Visit(ctx.Expression(1)).(object.Object)
	if !ok {
		return types.Failure
	}
	val, err := lhs.BinaryOp(parser.MlanParserXor, rhs)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.LineError(ctx.GetStart(), fmt.Errorf("unsupported expression: %s ^ %s", lhs.TypeName(), rhs.TypeName()))
		} else {
			v.LineError(ctx.GetStart(), err)
		}
		return types.Failure
	}
	return val
}

func (v *Visitor) VisitForStatement(ctx *parser.ForStatementContext) any {
	// левая часть присвоения (изначальное значение счетчика)
	if ok := v.Visit(ctx.Assignment(0)).(types.VisitResultType); !ok {
		return types.Failure
	}

	for {
		// условие выхода из цикла
		res, ok := v.Visit(ctx.Expression()).(object.Object)
		if !ok {
			v.LineError(ctx.GetStart(), fmt.Errorf("invalid expression for loop"))
			return types.Failure
		}

		// проверяем, что тип результата - булева
		if _, ok = res.(*object.Bool); !ok {
			v.LineError(ctx.GetStart(), fmt.Errorf("non-bool (%s) expression in for loop", res.TypeName()))
			return types.Failure
		}

		// проверяем, что условие выполняется
		if !res.(*object.Bool).GetValue().(bool) {
			// если false -> выходим из цикла
			break
		}

		// выполняем все выражения внутри блока цикла
		for _, item := range ctx.AllStatement() {
			// обработка break
			if scope.CurrentScope.IsBreak() {
				scope.CurrentScope.SetLoopBreak(false)
				return types.Success
			}
			// обработка continue
			if scope.CurrentScope.IsContinue() {
				scope.CurrentScope.SetLoopContinue(false)
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

func (v *Visitor) VisitAssignmentSum(ctx *parser.AssignmentSumContext) any {
	// получение значения переменной из скоупа
	val := scope.CurrentScope.Get(ctx.GetVarScalarName().GetText(), true)
	if val == nil {
		v.LineError(ctx.GetStart(), fmt.Errorf("undefined variable '%s'", ctx.GetVarScalarName().GetText()))
		return types.Failure
	}
	// получение результата выражения (справа)
	rhs, ok := v.Visit(ctx.Expression()).(object.Object)
	if !ok {
		return types.Failure
	}
	// получаем итоговый скаляр
	res, err := val.BinaryOp(parser.MlanLexerAssignSum, rhs)
	if err != nil {
		v.LineError(ctx.GetStart(), err)
		return types.Failure
	}
	// сохраняем новый скаляр
	scope.CurrentScope.Put(ctx.GetVarScalarName().GetText(), res)

	return types.Success
}

func (v *Visitor) VisitBreakStatement(ctx *parser.BreakStatementContext) any {
	if scope.CurrentScope.IsInLoop() {
		scope.CurrentScope.SetLoopBreak(true)
		return types.Success
	}
	v.LineError(ctx.GetStart(), fmt.Errorf("break outside of loop"))
	return types.Failure
}

func (v *Visitor) VisitContinueStatement(ctx *parser.ContinueStatementContext) any {
	if scope.CurrentScope.IsInLoop() {
		scope.CurrentScope.SetLoopContinue(true)
		return types.Success
	}
	v.LineError(ctx.GetStart(), fmt.Errorf("continue outside of loop"))
	return types.Failure
}

func (v *Visitor) VisitWhileStatement(ctx *parser.WhileStatementContext) any {
	for {
		// условие выхода из цикла
		res, ok := v.Visit(ctx.Expression()).(object.Object)
		if !ok {
			return types.Failure
		}

		// проверяем, что условие булева
		if _, ok = res.(*object.Bool); !ok {
			v.LineError(ctx.GetStart(), fmt.Errorf("non-bool (%s) expression in for loop", res.TypeName()))
			return types.Failure
		}

		// проверяем, что условие выполняется
		if !res.(*object.Bool).GetValue().(bool) {
			// если false -> выходим из цикла
			break
		}

		for _, item := range ctx.AllStatement() {
			// обработка break
			if scope.CurrentScope.IsBreak() {
				scope.CurrentScope.SetLoopBreak(false)
				return types.Success
			}
			// обработка continue
			if scope.CurrentScope.IsContinue() {
				scope.CurrentScope.SetLoopContinue(false)
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
	}

	return types.Success
}

func (v *Visitor) VisitIfStatement(ctx *parser.IfStatementContext) any {
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

func (v *Visitor) VisitIfBlockStatement(ctx *parser.IfBlockStatementContext) any {
	res, ok := v.Visit(ctx.Expression()).(object.Object)
	if !ok {
		return types.Failure
	}
	// кастим в бул
	res, err := builtin.Bool(res)
	if err != nil {
		v.LineError(ctx.GetStart(), err)
		return types.Failure
	}

	if res.(*object.Bool).GetValue().(bool) {
		// если булевый объект true -> выполняем блок
		for _, item := range ctx.AllStatement() {
			if ok := v.Visit(item).(types.VisitResultType); !ok {
				return types.Failure
			}
		}
	}

	return res
}

func (v *Visitor) VisitElseBlockStatement(ctx *parser.ElseBlockStatementContext) any {
	for _, item := range ctx.AllStatement() {
		if ok := v.Visit(item).(types.VisitResultType); !ok {
			return types.Failure
		}
	}
	return types.Success
}

func (v *Visitor) VisitElifBlockStatement(ctx *parser.ElifBlockStatementContext) any {
	res, ok := v.Visit(ctx.Expression()).(object.Object)
	if !ok {
		return types.Failure
	}
	// кастим в бул
	res, err := builtin.Bool(res)
	if err != nil {
		v.LineError(ctx.GetStart(), err)
		return types.Failure
	}

	if res.(*object.Bool).GetValue().(bool) {
		// если булевый объект true -> выполняем блок
		for _, item := range ctx.AllStatement() {
			if ok := v.Visit(item).(types.VisitResultType); !ok {
				return types.Failure
			}
		}
	}

	return res
}

func (v *Visitor) VisitFunctionDefinition(ctx *parser.FunctionDefinitionContext) any {
	// получение аргументов
	var args []string
	if ctx.FunctionParameters() != nil {
		for _, item := range ctx.FunctionParameters().AllIdentifier() {
			args = append(args, item.GetText())
		}
	}

	// проверяем, что функция еще не существует
	if _, ok := storage.NativeFunctions[ctx.GetVarFunctionName().GetText()]; ok {
		v.LineError(ctx.GetStart(), fmt.Errorf("function '%s' already defined", ctx.GetVarFunctionName().GetText()))
		return types.Failure
	}

	// создаем новую нативную функцию
	storage.NativeFunctions[ctx.GetVarFunctionName().GetText()] = object.NewNativeFunc(
		ctx.GetVarFunctionName().GetText(),
		args,
		len(args),
		ctx.AllStatement(),
	)

	return types.Success
}

func (v *Visitor) VisitExpressionFunctionInvoke(ctx *parser.ExpressionFunctionInvokeContext) any {
	res, ok := v.Visit(ctx.FunctionInvoke()).(object.Object)
	if !ok {
		// значит вернулся статус
		if v.Error == nil {
			return object.NewNull()
		}
		return types.Failure
	}
	return res
}

func (v *Visitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) any {
	if !scope.CurrentScope.IsInFunc() {
		v.LineError(ctx.GetStart(), fmt.Errorf("return outside of function"))
		return types.Failure
	}
	// получаем результат return
	if ctx.Expression() != nil {
		res, ok := v.Visit(ctx.Expression()).(object.Object)
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

func (v *Visitor) VisitExpressionIndex(ctx *parser.ExpressionIndexContext) any {
	lhs, ok := v.Visit(ctx.Expression()).(object.Object)
	if !ok {
		return types.Failure
	}
	idx, ok := v.Visit(ctx.Index().Expression()).(object.Object)
	if !ok {
		return types.Failure
	}
	res, err := lhs.IndexGet(idx)
	if err != nil {
		if errors.Is(err, object.ErrInvalidOp) {
			v.LineError(ctx.GetStart(), fmt.Errorf("unsupported expression: %s[%s]", lhs.TypeName(), idx.TypeName()))
		} else {
			v.LineError(ctx.GetStart(), err)
		}
		return types.Failure
	}
	return res
}

func (v *Visitor) VisitAssignmentIndexRegular(ctx *parser.AssignmentIndexRegularContext) any {
	// забираем из скоупа переменную
	lhs := scope.CurrentScope.Get(ctx.GetVarScalarName().GetText(), true)
	if lhs == nil {
		v.LineError(ctx.GetStart(), fmt.Errorf("undefined variable '%s'", ctx.GetVarScalarName().GetText()))
		return types.Failure
	}
	// получаем индекс
	idx, ok := v.Visit(ctx.Index().Expression()).(object.Object)
	if !ok {
		return types.Failure
	}
	// получаем значение
	val, ok := v.Visit(ctx.Expression()).(object.Object)
	if !ok {
		return types.Failure
	}
	// проставляем значение по индексу
	if err := lhs.IndexSet(idx, val); err != nil {
		v.LineError(ctx.GetStart(), err)
		return types.Failure
	}
	return types.Success
}

func (v *Visitor) VisitAssignmentClosure(ctx *parser.AssignmentClosureContext) any {
	// формируем объект функции
	if ctx.ClosureDefinition() == nil {
		v.LineError(ctx.GetStart(), fmt.Errorf("no definition of closure"))
		return types.Failure
	}
	// получаем аргументы
	var args []string
	if ctx.ClosureDefinition().FunctionParameters() != nil {
		for _, item := range ctx.ClosureDefinition().FunctionParameters().AllIdentifier() {
			args = append(args, item.GetText())
		}
	}
	// формируем объект функции
	fn := object.NewNativeFunc(
		ctx.GetVarScalarName().GetText(),
		args,
		len(args),
		ctx.ClosureDefinition().AllStatement(),
	)

	// сохраняем в переменные
	scope.CurrentScope.Put(ctx.GetVarScalarName().GetText(), fn)

	return types.Success
}

func (v *Visitor) VisitExpressionClosureInvoke(ctx *parser.ExpressionClosureInvokeContext) any {
	res, ok := v.Visit(ctx.ClosureInvoke()).(object.Object)
	if !ok {
		// значит вернулся статус
		if v.Error == nil {
			return object.NewNull()
		}
		return types.Failure
	}
	return res
}

func (v *Visitor) VisitIdentifierClosureInvoke(ctx *parser.IdentifierClosureInvokeContext) any {
	var params []object.Object

	for _, item := range ctx.AllExpression() {
		// собираем аргументы для функции
		res, ok := v.Visit(item).(object.Object)
		if !ok {
			return types.Failure
		}
		params = append(params, res)
	}

	return v.invokeClosureFunc(ctx.GetStart(), ctx.GetVarClosureName().GetText(), params...)
}

func (v *Visitor) VisitIncludeSubmodule(ctx *parser.IncludeSubmoduleContext) any {
	val, ok := v.Visit(ctx.Expression()).(object.Object)
	if !ok {
		return types.Failure
	}
	path, ok := val.(*object.Str)
	if !ok {
		v.LineError(ctx.GetStart(), fmt.Errorf("invalid argument of type '%s'", val.TypeName()))
		return types.Failure
	}
	data, err := os.ReadFile(path.GetValue().(string))
	if err != nil {
		v.LineError(ctx.GetStart(), err)
		return types.Failure
	}
	tree, err := utils.CreateAST(string(data))
	if err != nil {
		v.LineError(ctx.GetStart(), err)
		return types.Failure
	}
	return tree
}

func (v *Visitor) VisitExpressionClosure(ctx *parser.ExpressionClosureContext) any {
	// формируем объект функции
	if ctx.ClosureDefinition() == nil {
		v.LineError(ctx.GetStart(), fmt.Errorf("no definition of closure"))
		return types.Failure
	}
	// получаем аргументы
	var args []string
	if ctx.ClosureDefinition().FunctionParameters() != nil {
		for _, item := range ctx.ClosureDefinition().FunctionParameters().AllIdentifier() {
			args = append(args, item.GetText())
		}
	}
	// формируем объект функции
	fn := object.NewNativeFunc(
		"unknown",
		args,
		len(args),
		ctx.ClosureDefinition().AllStatement(),
	)

	return fn
}
