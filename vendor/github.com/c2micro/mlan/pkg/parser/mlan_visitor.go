// Code generated from ./Mlan.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Mlan
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MlanParser.
type MlanVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MlanParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by MlanParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by MlanParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by MlanParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by MlanParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by MlanParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by MlanParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by MlanParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by MlanParser#assignmentRegular.
	VisitAssignmentRegular(ctx *AssignmentRegularContext) interface{}

	// Visit a parse tree produced by MlanParser#assignmentClosure.
	VisitAssignmentClosure(ctx *AssignmentClosureContext) interface{}

	// Visit a parse tree produced by MlanParser#assignmentSum.
	VisitAssignmentSum(ctx *AssignmentSumContext) interface{}

	// Visit a parse tree produced by MlanParser#assignmentSub.
	VisitAssignmentSub(ctx *AssignmentSubContext) interface{}

	// Visit a parse tree produced by MlanParser#assignmentMul.
	VisitAssignmentMul(ctx *AssignmentMulContext) interface{}

	// Visit a parse tree produced by MlanParser#assignmentDiv.
	VisitAssignmentDiv(ctx *AssignmentDivContext) interface{}

	// Visit a parse tree produced by MlanParser#assignmentMod.
	VisitAssignmentMod(ctx *AssignmentModContext) interface{}

	// Visit a parse tree produced by MlanParser#assignmentPow.
	VisitAssignmentPow(ctx *AssignmentPowContext) interface{}

	// Visit a parse tree produced by MlanParser#assignmentIndexRegular.
	VisitAssignmentIndexRegular(ctx *AssignmentIndexRegularContext) interface{}

	// Visit a parse tree produced by MlanParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by MlanParser#dictUnit.
	VisitDictUnit(ctx *DictUnitContext) interface{}

	// Visit a parse tree produced by MlanParser#dict.
	VisitDict(ctx *DictContext) interface{}

	// Visit a parse tree produced by MlanParser#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by MlanParser#identifierFunctionInvoke.
	VisitIdentifierFunctionInvoke(ctx *IdentifierFunctionInvokeContext) interface{}

	// Visit a parse tree produced by MlanParser#identifierClosureInvoke.
	VisitIdentifierClosureInvoke(ctx *IdentifierClosureInvokeContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionIntegerHex.
	VisitExpressionIntegerHex(ctx *ExpressionIntegerHexContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionFunctionInvoke.
	VisitExpressionFunctionInvoke(ctx *ExpressionFunctionInvokeContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionUnaryNegation.
	VisitExpressionUnaryNegation(ctx *ExpressionUnaryNegationContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionBool.
	VisitExpressionBool(ctx *ExpressionBoolContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionPow.
	VisitExpressionPow(ctx *ExpressionPowContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionXor.
	VisitExpressionXor(ctx *ExpressionXorContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionEqual.
	VisitExpressionEqual(ctx *ExpressionEqualContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionClosure.
	VisitExpressionClosure(ctx *ExpressionClosureContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionDict.
	VisitExpressionDict(ctx *ExpressionDictContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionIdentifier.
	VisitExpressionIdentifier(ctx *ExpressionIdentifierContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionSumSub.
	VisitExpressionSumSub(ctx *ExpressionSumSubContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionComparison.
	VisitExpressionComparison(ctx *ExpressionComparisonContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionLogicalOr.
	VisitExpressionLogicalOr(ctx *ExpressionLogicalOrContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionIndex.
	VisitExpressionIndex(ctx *ExpressionIndexContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionLogicalNot.
	VisitExpressionLogicalNot(ctx *ExpressionLogicalNotContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionClosureInvoke.
	VisitExpressionClosureInvoke(ctx *ExpressionClosureInvokeContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionParentheses.
	VisitExpressionParentheses(ctx *ExpressionParenthesesContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionMulDivMod.
	VisitExpressionMulDivMod(ctx *ExpressionMulDivModContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionLogicalAnd.
	VisitExpressionLogicalAnd(ctx *ExpressionLogicalAndContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionFloat.
	VisitExpressionFloat(ctx *ExpressionFloatContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionInteger.
	VisitExpressionInteger(ctx *ExpressionIntegerContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionNull.
	VisitExpressionNull(ctx *ExpressionNullContext) interface{}

	// Visit a parse tree produced by MlanParser#expressionString.
	VisitExpressionString(ctx *ExpressionStringContext) interface{}

	// Visit a parse tree produced by MlanParser#ifBlockStatement.
	VisitIfBlockStatement(ctx *IfBlockStatementContext) interface{}

	// Visit a parse tree produced by MlanParser#elifBlockStatement.
	VisitElifBlockStatement(ctx *ElifBlockStatementContext) interface{}

	// Visit a parse tree produced by MlanParser#elseBlockStatement.
	VisitElseBlockStatement(ctx *ElseBlockStatementContext) interface{}

	// Visit a parse tree produced by MlanParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by MlanParser#functionParameters.
	VisitFunctionParameters(ctx *FunctionParametersContext) interface{}

	// Visit a parse tree produced by MlanParser#functionDefinition.
	VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{}

	// Visit a parse tree produced by MlanParser#closureDefinition.
	VisitClosureDefinition(ctx *ClosureDefinitionContext) interface{}

	// Visit a parse tree produced by MlanParser#includeSubmodule.
	VisitIncludeSubmodule(ctx *IncludeSubmoduleContext) interface{}
}
