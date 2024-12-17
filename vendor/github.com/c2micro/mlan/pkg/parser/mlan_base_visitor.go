// Code generated from ./Mlan.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Mlan
import "github.com/antlr4-go/antlr/v4"

type BaseMlanVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMlanVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitAssignmentRegular(ctx *AssignmentRegularContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitAssignmentClosure(ctx *AssignmentClosureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitAssignmentSum(ctx *AssignmentSumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitAssignmentSub(ctx *AssignmentSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitAssignmentMul(ctx *AssignmentMulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitAssignmentDiv(ctx *AssignmentDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitAssignmentMod(ctx *AssignmentModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitAssignmentPow(ctx *AssignmentPowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitAssignmentIndexRegular(ctx *AssignmentIndexRegularContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitDictUnit(ctx *DictUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitDict(ctx *DictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitIdentifierFunctionInvoke(ctx *IdentifierFunctionInvokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitIdentifierClosureInvoke(ctx *IdentifierClosureInvokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionIntegerHex(ctx *ExpressionIntegerHexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionFunctionInvoke(ctx *ExpressionFunctionInvokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionUnaryNegation(ctx *ExpressionUnaryNegationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionBool(ctx *ExpressionBoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionPow(ctx *ExpressionPowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionXor(ctx *ExpressionXorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionEqual(ctx *ExpressionEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionClosure(ctx *ExpressionClosureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionDict(ctx *ExpressionDictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionIdentifier(ctx *ExpressionIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionSumSub(ctx *ExpressionSumSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionComparison(ctx *ExpressionComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionLogicalOr(ctx *ExpressionLogicalOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionIndex(ctx *ExpressionIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionLogicalNot(ctx *ExpressionLogicalNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionClosureInvoke(ctx *ExpressionClosureInvokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionParentheses(ctx *ExpressionParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionMulDivMod(ctx *ExpressionMulDivModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionLogicalAnd(ctx *ExpressionLogicalAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionFloat(ctx *ExpressionFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionInteger(ctx *ExpressionIntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionNull(ctx *ExpressionNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpressionString(ctx *ExpressionStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitIfBlockStatement(ctx *IfBlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitElifBlockStatement(ctx *ElifBlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitElseBlockStatement(ctx *ElseBlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitFunctionParameters(ctx *FunctionParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitClosureDefinition(ctx *ClosureDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitIncludeSubmodule(ctx *IncludeSubmoduleContext) interface{} {
	return v.VisitChildren(ctx)
}
