// Code generated from ./Mlan.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Mlan
import "github.com/antlr4-go/antlr/v4"

// MlanListener is a complete listener for a parse tree produced by MlanParser.
type MlanListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterAssignmentRegular is called when entering the assignmentRegular production.
	EnterAssignmentRegular(c *AssignmentRegularContext)

	// EnterAssignmentClosure is called when entering the assignmentClosure production.
	EnterAssignmentClosure(c *AssignmentClosureContext)

	// EnterAssignmentSum is called when entering the assignmentSum production.
	EnterAssignmentSum(c *AssignmentSumContext)

	// EnterAssignmentSub is called when entering the assignmentSub production.
	EnterAssignmentSub(c *AssignmentSubContext)

	// EnterAssignmentMul is called when entering the assignmentMul production.
	EnterAssignmentMul(c *AssignmentMulContext)

	// EnterAssignmentDiv is called when entering the assignmentDiv production.
	EnterAssignmentDiv(c *AssignmentDivContext)

	// EnterAssignmentMod is called when entering the assignmentMod production.
	EnterAssignmentMod(c *AssignmentModContext)

	// EnterAssignmentPow is called when entering the assignmentPow production.
	EnterAssignmentPow(c *AssignmentPowContext)

	// EnterAssignmentIndexRegular is called when entering the assignmentIndexRegular production.
	EnterAssignmentIndexRegular(c *AssignmentIndexRegularContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterDictUnit is called when entering the dictUnit production.
	EnterDictUnit(c *DictUnitContext)

	// EnterDict is called when entering the dict production.
	EnterDict(c *DictContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterIdentifierFunctionInvoke is called when entering the identifierFunctionInvoke production.
	EnterIdentifierFunctionInvoke(c *IdentifierFunctionInvokeContext)

	// EnterIdentifierClosureInvoke is called when entering the identifierClosureInvoke production.
	EnterIdentifierClosureInvoke(c *IdentifierClosureInvokeContext)

	// EnterExpressionIntegerHex is called when entering the expressionIntegerHex production.
	EnterExpressionIntegerHex(c *ExpressionIntegerHexContext)

	// EnterExpressionFunctionInvoke is called when entering the expressionFunctionInvoke production.
	EnterExpressionFunctionInvoke(c *ExpressionFunctionInvokeContext)

	// EnterExpressionUnaryNegation is called when entering the expressionUnaryNegation production.
	EnterExpressionUnaryNegation(c *ExpressionUnaryNegationContext)

	// EnterExpressionBool is called when entering the expressionBool production.
	EnterExpressionBool(c *ExpressionBoolContext)

	// EnterExpressionPow is called when entering the expressionPow production.
	EnterExpressionPow(c *ExpressionPowContext)

	// EnterExpressionXor is called when entering the expressionXor production.
	EnterExpressionXor(c *ExpressionXorContext)

	// EnterExpressionEqual is called when entering the expressionEqual production.
	EnterExpressionEqual(c *ExpressionEqualContext)

	// EnterExpressionClosure is called when entering the expressionClosure production.
	EnterExpressionClosure(c *ExpressionClosureContext)

	// EnterExpressionDict is called when entering the expressionDict production.
	EnterExpressionDict(c *ExpressionDictContext)

	// EnterExpressionIdentifier is called when entering the expressionIdentifier production.
	EnterExpressionIdentifier(c *ExpressionIdentifierContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterExpressionSumSub is called when entering the expressionSumSub production.
	EnterExpressionSumSub(c *ExpressionSumSubContext)

	// EnterExpressionComparison is called when entering the expressionComparison production.
	EnterExpressionComparison(c *ExpressionComparisonContext)

	// EnterExpressionLogicalOr is called when entering the expressionLogicalOr production.
	EnterExpressionLogicalOr(c *ExpressionLogicalOrContext)

	// EnterExpressionIndex is called when entering the expressionIndex production.
	EnterExpressionIndex(c *ExpressionIndexContext)

	// EnterExpressionLogicalNot is called when entering the expressionLogicalNot production.
	EnterExpressionLogicalNot(c *ExpressionLogicalNotContext)

	// EnterExpressionClosureInvoke is called when entering the expressionClosureInvoke production.
	EnterExpressionClosureInvoke(c *ExpressionClosureInvokeContext)

	// EnterExpressionParentheses is called when entering the expressionParentheses production.
	EnterExpressionParentheses(c *ExpressionParenthesesContext)

	// EnterExpressionMulDivMod is called when entering the expressionMulDivMod production.
	EnterExpressionMulDivMod(c *ExpressionMulDivModContext)

	// EnterExpressionLogicalAnd is called when entering the expressionLogicalAnd production.
	EnterExpressionLogicalAnd(c *ExpressionLogicalAndContext)

	// EnterExpressionFloat is called when entering the expressionFloat production.
	EnterExpressionFloat(c *ExpressionFloatContext)

	// EnterExpressionInteger is called when entering the expressionInteger production.
	EnterExpressionInteger(c *ExpressionIntegerContext)

	// EnterExpressionNull is called when entering the expressionNull production.
	EnterExpressionNull(c *ExpressionNullContext)

	// EnterExpressionString is called when entering the expressionString production.
	EnterExpressionString(c *ExpressionStringContext)

	// EnterIfBlockStatement is called when entering the ifBlockStatement production.
	EnterIfBlockStatement(c *IfBlockStatementContext)

	// EnterElifBlockStatement is called when entering the elifBlockStatement production.
	EnterElifBlockStatement(c *ElifBlockStatementContext)

	// EnterElseBlockStatement is called when entering the elseBlockStatement production.
	EnterElseBlockStatement(c *ElseBlockStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterFunctionParameters is called when entering the functionParameters production.
	EnterFunctionParameters(c *FunctionParametersContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterClosureDefinition is called when entering the closureDefinition production.
	EnterClosureDefinition(c *ClosureDefinitionContext)

	// EnterIncludeSubmodule is called when entering the includeSubmodule production.
	EnterIncludeSubmodule(c *IncludeSubmoduleContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitAssignmentRegular is called when exiting the assignmentRegular production.
	ExitAssignmentRegular(c *AssignmentRegularContext)

	// ExitAssignmentClosure is called when exiting the assignmentClosure production.
	ExitAssignmentClosure(c *AssignmentClosureContext)

	// ExitAssignmentSum is called when exiting the assignmentSum production.
	ExitAssignmentSum(c *AssignmentSumContext)

	// ExitAssignmentSub is called when exiting the assignmentSub production.
	ExitAssignmentSub(c *AssignmentSubContext)

	// ExitAssignmentMul is called when exiting the assignmentMul production.
	ExitAssignmentMul(c *AssignmentMulContext)

	// ExitAssignmentDiv is called when exiting the assignmentDiv production.
	ExitAssignmentDiv(c *AssignmentDivContext)

	// ExitAssignmentMod is called when exiting the assignmentMod production.
	ExitAssignmentMod(c *AssignmentModContext)

	// ExitAssignmentPow is called when exiting the assignmentPow production.
	ExitAssignmentPow(c *AssignmentPowContext)

	// ExitAssignmentIndexRegular is called when exiting the assignmentIndexRegular production.
	ExitAssignmentIndexRegular(c *AssignmentIndexRegularContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitDictUnit is called when exiting the dictUnit production.
	ExitDictUnit(c *DictUnitContext)

	// ExitDict is called when exiting the dict production.
	ExitDict(c *DictContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitIdentifierFunctionInvoke is called when exiting the identifierFunctionInvoke production.
	ExitIdentifierFunctionInvoke(c *IdentifierFunctionInvokeContext)

	// ExitIdentifierClosureInvoke is called when exiting the identifierClosureInvoke production.
	ExitIdentifierClosureInvoke(c *IdentifierClosureInvokeContext)

	// ExitExpressionIntegerHex is called when exiting the expressionIntegerHex production.
	ExitExpressionIntegerHex(c *ExpressionIntegerHexContext)

	// ExitExpressionFunctionInvoke is called when exiting the expressionFunctionInvoke production.
	ExitExpressionFunctionInvoke(c *ExpressionFunctionInvokeContext)

	// ExitExpressionUnaryNegation is called when exiting the expressionUnaryNegation production.
	ExitExpressionUnaryNegation(c *ExpressionUnaryNegationContext)

	// ExitExpressionBool is called when exiting the expressionBool production.
	ExitExpressionBool(c *ExpressionBoolContext)

	// ExitExpressionPow is called when exiting the expressionPow production.
	ExitExpressionPow(c *ExpressionPowContext)

	// ExitExpressionXor is called when exiting the expressionXor production.
	ExitExpressionXor(c *ExpressionXorContext)

	// ExitExpressionEqual is called when exiting the expressionEqual production.
	ExitExpressionEqual(c *ExpressionEqualContext)

	// ExitExpressionClosure is called when exiting the expressionClosure production.
	ExitExpressionClosure(c *ExpressionClosureContext)

	// ExitExpressionDict is called when exiting the expressionDict production.
	ExitExpressionDict(c *ExpressionDictContext)

	// ExitExpressionIdentifier is called when exiting the expressionIdentifier production.
	ExitExpressionIdentifier(c *ExpressionIdentifierContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitExpressionSumSub is called when exiting the expressionSumSub production.
	ExitExpressionSumSub(c *ExpressionSumSubContext)

	// ExitExpressionComparison is called when exiting the expressionComparison production.
	ExitExpressionComparison(c *ExpressionComparisonContext)

	// ExitExpressionLogicalOr is called when exiting the expressionLogicalOr production.
	ExitExpressionLogicalOr(c *ExpressionLogicalOrContext)

	// ExitExpressionIndex is called when exiting the expressionIndex production.
	ExitExpressionIndex(c *ExpressionIndexContext)

	// ExitExpressionLogicalNot is called when exiting the expressionLogicalNot production.
	ExitExpressionLogicalNot(c *ExpressionLogicalNotContext)

	// ExitExpressionClosureInvoke is called when exiting the expressionClosureInvoke production.
	ExitExpressionClosureInvoke(c *ExpressionClosureInvokeContext)

	// ExitExpressionParentheses is called when exiting the expressionParentheses production.
	ExitExpressionParentheses(c *ExpressionParenthesesContext)

	// ExitExpressionMulDivMod is called when exiting the expressionMulDivMod production.
	ExitExpressionMulDivMod(c *ExpressionMulDivModContext)

	// ExitExpressionLogicalAnd is called when exiting the expressionLogicalAnd production.
	ExitExpressionLogicalAnd(c *ExpressionLogicalAndContext)

	// ExitExpressionFloat is called when exiting the expressionFloat production.
	ExitExpressionFloat(c *ExpressionFloatContext)

	// ExitExpressionInteger is called when exiting the expressionInteger production.
	ExitExpressionInteger(c *ExpressionIntegerContext)

	// ExitExpressionNull is called when exiting the expressionNull production.
	ExitExpressionNull(c *ExpressionNullContext)

	// ExitExpressionString is called when exiting the expressionString production.
	ExitExpressionString(c *ExpressionStringContext)

	// ExitIfBlockStatement is called when exiting the ifBlockStatement production.
	ExitIfBlockStatement(c *IfBlockStatementContext)

	// ExitElifBlockStatement is called when exiting the elifBlockStatement production.
	ExitElifBlockStatement(c *ElifBlockStatementContext)

	// ExitElseBlockStatement is called when exiting the elseBlockStatement production.
	ExitElseBlockStatement(c *ElseBlockStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitFunctionParameters is called when exiting the functionParameters production.
	ExitFunctionParameters(c *FunctionParametersContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitClosureDefinition is called when exiting the closureDefinition production.
	ExitClosureDefinition(c *ClosureDefinitionContext)

	// ExitIncludeSubmodule is called when exiting the includeSubmodule production.
	ExitIncludeSubmodule(c *IncludeSubmoduleContext)
}
