// Code generated from ./Mlan.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Mlan
import "github.com/antlr4-go/antlr/v4"

// BaseMlanListener is a complete listener for a parse tree produced by MlanParser.
type BaseMlanListener struct{}

var _ MlanListener = &BaseMlanListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMlanListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMlanListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMlanListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMlanListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseMlanListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseMlanListener) ExitProgram(ctx *ProgramContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseMlanListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseMlanListener) ExitBlock(ctx *BlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseMlanListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseMlanListener) ExitStatement(ctx *StatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseMlanListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseMlanListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseMlanListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseMlanListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseMlanListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseMlanListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseMlanListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseMlanListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseMlanListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseMlanListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterAssignmentRegular is called when production assignmentRegular is entered.
func (s *BaseMlanListener) EnterAssignmentRegular(ctx *AssignmentRegularContext) {}

// ExitAssignmentRegular is called when production assignmentRegular is exited.
func (s *BaseMlanListener) ExitAssignmentRegular(ctx *AssignmentRegularContext) {}

// EnterAssignmentClosure is called when production assignmentClosure is entered.
func (s *BaseMlanListener) EnterAssignmentClosure(ctx *AssignmentClosureContext) {}

// ExitAssignmentClosure is called when production assignmentClosure is exited.
func (s *BaseMlanListener) ExitAssignmentClosure(ctx *AssignmentClosureContext) {}

// EnterAssignmentSum is called when production assignmentSum is entered.
func (s *BaseMlanListener) EnterAssignmentSum(ctx *AssignmentSumContext) {}

// ExitAssignmentSum is called when production assignmentSum is exited.
func (s *BaseMlanListener) ExitAssignmentSum(ctx *AssignmentSumContext) {}

// EnterAssignmentSub is called when production assignmentSub is entered.
func (s *BaseMlanListener) EnterAssignmentSub(ctx *AssignmentSubContext) {}

// ExitAssignmentSub is called when production assignmentSub is exited.
func (s *BaseMlanListener) ExitAssignmentSub(ctx *AssignmentSubContext) {}

// EnterAssignmentMul is called when production assignmentMul is entered.
func (s *BaseMlanListener) EnterAssignmentMul(ctx *AssignmentMulContext) {}

// ExitAssignmentMul is called when production assignmentMul is exited.
func (s *BaseMlanListener) ExitAssignmentMul(ctx *AssignmentMulContext) {}

// EnterAssignmentDiv is called when production assignmentDiv is entered.
func (s *BaseMlanListener) EnterAssignmentDiv(ctx *AssignmentDivContext) {}

// ExitAssignmentDiv is called when production assignmentDiv is exited.
func (s *BaseMlanListener) ExitAssignmentDiv(ctx *AssignmentDivContext) {}

// EnterAssignmentMod is called when production assignmentMod is entered.
func (s *BaseMlanListener) EnterAssignmentMod(ctx *AssignmentModContext) {}

// ExitAssignmentMod is called when production assignmentMod is exited.
func (s *BaseMlanListener) ExitAssignmentMod(ctx *AssignmentModContext) {}

// EnterAssignmentPow is called when production assignmentPow is entered.
func (s *BaseMlanListener) EnterAssignmentPow(ctx *AssignmentPowContext) {}

// ExitAssignmentPow is called when production assignmentPow is exited.
func (s *BaseMlanListener) ExitAssignmentPow(ctx *AssignmentPowContext) {}

// EnterAssignmentIndexRegular is called when production assignmentIndexRegular is entered.
func (s *BaseMlanListener) EnterAssignmentIndexRegular(ctx *AssignmentIndexRegularContext) {}

// ExitAssignmentIndexRegular is called when production assignmentIndexRegular is exited.
func (s *BaseMlanListener) ExitAssignmentIndexRegular(ctx *AssignmentIndexRegularContext) {}

// EnterList is called when production list is entered.
func (s *BaseMlanListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseMlanListener) ExitList(ctx *ListContext) {}

// EnterDictUnit is called when production dictUnit is entered.
func (s *BaseMlanListener) EnterDictUnit(ctx *DictUnitContext) {}

// ExitDictUnit is called when production dictUnit is exited.
func (s *BaseMlanListener) ExitDictUnit(ctx *DictUnitContext) {}

// EnterDict is called when production dict is entered.
func (s *BaseMlanListener) EnterDict(ctx *DictContext) {}

// ExitDict is called when production dict is exited.
func (s *BaseMlanListener) ExitDict(ctx *DictContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseMlanListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseMlanListener) ExitIndex(ctx *IndexContext) {}

// EnterIdentifierFunctionInvoke is called when production identifierFunctionInvoke is entered.
func (s *BaseMlanListener) EnterIdentifierFunctionInvoke(ctx *IdentifierFunctionInvokeContext) {}

// ExitIdentifierFunctionInvoke is called when production identifierFunctionInvoke is exited.
func (s *BaseMlanListener) ExitIdentifierFunctionInvoke(ctx *IdentifierFunctionInvokeContext) {}

// EnterIdentifierClosureInvoke is called when production identifierClosureInvoke is entered.
func (s *BaseMlanListener) EnterIdentifierClosureInvoke(ctx *IdentifierClosureInvokeContext) {}

// ExitIdentifierClosureInvoke is called when production identifierClosureInvoke is exited.
func (s *BaseMlanListener) ExitIdentifierClosureInvoke(ctx *IdentifierClosureInvokeContext) {}

// EnterExpressionIntegerHex is called when production expressionIntegerHex is entered.
func (s *BaseMlanListener) EnterExpressionIntegerHex(ctx *ExpressionIntegerHexContext) {}

// ExitExpressionIntegerHex is called when production expressionIntegerHex is exited.
func (s *BaseMlanListener) ExitExpressionIntegerHex(ctx *ExpressionIntegerHexContext) {}

// EnterExpressionFunctionInvoke is called when production expressionFunctionInvoke is entered.
func (s *BaseMlanListener) EnterExpressionFunctionInvoke(ctx *ExpressionFunctionInvokeContext) {}

// ExitExpressionFunctionInvoke is called when production expressionFunctionInvoke is exited.
func (s *BaseMlanListener) ExitExpressionFunctionInvoke(ctx *ExpressionFunctionInvokeContext) {}

// EnterExpressionUnaryNegation is called when production expressionUnaryNegation is entered.
func (s *BaseMlanListener) EnterExpressionUnaryNegation(ctx *ExpressionUnaryNegationContext) {}

// ExitExpressionUnaryNegation is called when production expressionUnaryNegation is exited.
func (s *BaseMlanListener) ExitExpressionUnaryNegation(ctx *ExpressionUnaryNegationContext) {}

// EnterExpressionBool is called when production expressionBool is entered.
func (s *BaseMlanListener) EnterExpressionBool(ctx *ExpressionBoolContext) {}

// ExitExpressionBool is called when production expressionBool is exited.
func (s *BaseMlanListener) ExitExpressionBool(ctx *ExpressionBoolContext) {}

// EnterExpressionPow is called when production expressionPow is entered.
func (s *BaseMlanListener) EnterExpressionPow(ctx *ExpressionPowContext) {}

// ExitExpressionPow is called when production expressionPow is exited.
func (s *BaseMlanListener) ExitExpressionPow(ctx *ExpressionPowContext) {}

// EnterExpressionXor is called when production expressionXor is entered.
func (s *BaseMlanListener) EnterExpressionXor(ctx *ExpressionXorContext) {}

// ExitExpressionXor is called when production expressionXor is exited.
func (s *BaseMlanListener) ExitExpressionXor(ctx *ExpressionXorContext) {}

// EnterExpressionEqual is called when production expressionEqual is entered.
func (s *BaseMlanListener) EnterExpressionEqual(ctx *ExpressionEqualContext) {}

// ExitExpressionEqual is called when production expressionEqual is exited.
func (s *BaseMlanListener) ExitExpressionEqual(ctx *ExpressionEqualContext) {}

// EnterExpressionClosure is called when production expressionClosure is entered.
func (s *BaseMlanListener) EnterExpressionClosure(ctx *ExpressionClosureContext) {}

// ExitExpressionClosure is called when production expressionClosure is exited.
func (s *BaseMlanListener) ExitExpressionClosure(ctx *ExpressionClosureContext) {}

// EnterExpressionDict is called when production expressionDict is entered.
func (s *BaseMlanListener) EnterExpressionDict(ctx *ExpressionDictContext) {}

// ExitExpressionDict is called when production expressionDict is exited.
func (s *BaseMlanListener) ExitExpressionDict(ctx *ExpressionDictContext) {}

// EnterExpressionIdentifier is called when production expressionIdentifier is entered.
func (s *BaseMlanListener) EnterExpressionIdentifier(ctx *ExpressionIdentifierContext) {}

// ExitExpressionIdentifier is called when production expressionIdentifier is exited.
func (s *BaseMlanListener) ExitExpressionIdentifier(ctx *ExpressionIdentifierContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseMlanListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseMlanListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterExpressionSumSub is called when production expressionSumSub is entered.
func (s *BaseMlanListener) EnterExpressionSumSub(ctx *ExpressionSumSubContext) {}

// ExitExpressionSumSub is called when production expressionSumSub is exited.
func (s *BaseMlanListener) ExitExpressionSumSub(ctx *ExpressionSumSubContext) {}

// EnterExpressionComparison is called when production expressionComparison is entered.
func (s *BaseMlanListener) EnterExpressionComparison(ctx *ExpressionComparisonContext) {}

// ExitExpressionComparison is called when production expressionComparison is exited.
func (s *BaseMlanListener) ExitExpressionComparison(ctx *ExpressionComparisonContext) {}

// EnterExpressionLogicalOr is called when production expressionLogicalOr is entered.
func (s *BaseMlanListener) EnterExpressionLogicalOr(ctx *ExpressionLogicalOrContext) {}

// ExitExpressionLogicalOr is called when production expressionLogicalOr is exited.
func (s *BaseMlanListener) ExitExpressionLogicalOr(ctx *ExpressionLogicalOrContext) {}

// EnterExpressionIndex is called when production expressionIndex is entered.
func (s *BaseMlanListener) EnterExpressionIndex(ctx *ExpressionIndexContext) {}

// ExitExpressionIndex is called when production expressionIndex is exited.
func (s *BaseMlanListener) ExitExpressionIndex(ctx *ExpressionIndexContext) {}

// EnterExpressionLogicalNot is called when production expressionLogicalNot is entered.
func (s *BaseMlanListener) EnterExpressionLogicalNot(ctx *ExpressionLogicalNotContext) {}

// ExitExpressionLogicalNot is called when production expressionLogicalNot is exited.
func (s *BaseMlanListener) ExitExpressionLogicalNot(ctx *ExpressionLogicalNotContext) {}

// EnterExpressionClosureInvoke is called when production expressionClosureInvoke is entered.
func (s *BaseMlanListener) EnterExpressionClosureInvoke(ctx *ExpressionClosureInvokeContext) {}

// ExitExpressionClosureInvoke is called when production expressionClosureInvoke is exited.
func (s *BaseMlanListener) ExitExpressionClosureInvoke(ctx *ExpressionClosureInvokeContext) {}

// EnterExpressionParentheses is called when production expressionParentheses is entered.
func (s *BaseMlanListener) EnterExpressionParentheses(ctx *ExpressionParenthesesContext) {}

// ExitExpressionParentheses is called when production expressionParentheses is exited.
func (s *BaseMlanListener) ExitExpressionParentheses(ctx *ExpressionParenthesesContext) {}

// EnterExpressionMulDivMod is called when production expressionMulDivMod is entered.
func (s *BaseMlanListener) EnterExpressionMulDivMod(ctx *ExpressionMulDivModContext) {}

// ExitExpressionMulDivMod is called when production expressionMulDivMod is exited.
func (s *BaseMlanListener) ExitExpressionMulDivMod(ctx *ExpressionMulDivModContext) {}

// EnterExpressionLogicalAnd is called when production expressionLogicalAnd is entered.
func (s *BaseMlanListener) EnterExpressionLogicalAnd(ctx *ExpressionLogicalAndContext) {}

// ExitExpressionLogicalAnd is called when production expressionLogicalAnd is exited.
func (s *BaseMlanListener) ExitExpressionLogicalAnd(ctx *ExpressionLogicalAndContext) {}

// EnterExpressionFloat is called when production expressionFloat is entered.
func (s *BaseMlanListener) EnterExpressionFloat(ctx *ExpressionFloatContext) {}

// ExitExpressionFloat is called when production expressionFloat is exited.
func (s *BaseMlanListener) ExitExpressionFloat(ctx *ExpressionFloatContext) {}

// EnterExpressionInteger is called when production expressionInteger is entered.
func (s *BaseMlanListener) EnterExpressionInteger(ctx *ExpressionIntegerContext) {}

// ExitExpressionInteger is called when production expressionInteger is exited.
func (s *BaseMlanListener) ExitExpressionInteger(ctx *ExpressionIntegerContext) {}

// EnterExpressionNull is called when production expressionNull is entered.
func (s *BaseMlanListener) EnterExpressionNull(ctx *ExpressionNullContext) {}

// ExitExpressionNull is called when production expressionNull is exited.
func (s *BaseMlanListener) ExitExpressionNull(ctx *ExpressionNullContext) {}

// EnterExpressionString is called when production expressionString is entered.
func (s *BaseMlanListener) EnterExpressionString(ctx *ExpressionStringContext) {}

// ExitExpressionString is called when production expressionString is exited.
func (s *BaseMlanListener) ExitExpressionString(ctx *ExpressionStringContext) {}

// EnterIfBlockStatement is called when production ifBlockStatement is entered.
func (s *BaseMlanListener) EnterIfBlockStatement(ctx *IfBlockStatementContext) {}

// ExitIfBlockStatement is called when production ifBlockStatement is exited.
func (s *BaseMlanListener) ExitIfBlockStatement(ctx *IfBlockStatementContext) {}

// EnterElifBlockStatement is called when production elifBlockStatement is entered.
func (s *BaseMlanListener) EnterElifBlockStatement(ctx *ElifBlockStatementContext) {}

// ExitElifBlockStatement is called when production elifBlockStatement is exited.
func (s *BaseMlanListener) ExitElifBlockStatement(ctx *ElifBlockStatementContext) {}

// EnterElseBlockStatement is called when production elseBlockStatement is entered.
func (s *BaseMlanListener) EnterElseBlockStatement(ctx *ElseBlockStatementContext) {}

// ExitElseBlockStatement is called when production elseBlockStatement is exited.
func (s *BaseMlanListener) ExitElseBlockStatement(ctx *ElseBlockStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseMlanListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseMlanListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterFunctionParameters is called when production functionParameters is entered.
func (s *BaseMlanListener) EnterFunctionParameters(ctx *FunctionParametersContext) {}

// ExitFunctionParameters is called when production functionParameters is exited.
func (s *BaseMlanListener) ExitFunctionParameters(ctx *FunctionParametersContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BaseMlanListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BaseMlanListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterClosureDefinition is called when production closureDefinition is entered.
func (s *BaseMlanListener) EnterClosureDefinition(ctx *ClosureDefinitionContext) {}

// ExitClosureDefinition is called when production closureDefinition is exited.
func (s *BaseMlanListener) ExitClosureDefinition(ctx *ClosureDefinitionContext) {}

// EnterIncludeSubmodule is called when production includeSubmodule is entered.
func (s *BaseMlanListener) EnterIncludeSubmodule(ctx *IncludeSubmoduleContext) {}

// ExitIncludeSubmodule is called when production includeSubmodule is exited.
func (s *BaseMlanListener) ExitIncludeSubmodule(ctx *IncludeSubmoduleContext) {}
