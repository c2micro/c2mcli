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

// EnterProg is called when production prog is entered.
func (s *BaseMlanListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseMlanListener) ExitProg(ctx *ProgContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseMlanListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseMlanListener) ExitStmt(ctx *StmtContext) {}

// EnterWhileStmt is called when production whileStmt is entered.
func (s *BaseMlanListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production whileStmt is exited.
func (s *BaseMlanListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *BaseMlanListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *BaseMlanListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterReturnStmt is called when production returnStmt is entered.
func (s *BaseMlanListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production returnStmt is exited.
func (s *BaseMlanListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterContinueStmt is called when production continueStmt is entered.
func (s *BaseMlanListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production continueStmt is exited.
func (s *BaseMlanListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *BaseMlanListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *BaseMlanListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterAssignRegular is called when production assignRegular is entered.
func (s *BaseMlanListener) EnterAssignRegular(ctx *AssignRegularContext) {}

// ExitAssignRegular is called when production assignRegular is exited.
func (s *BaseMlanListener) ExitAssignRegular(ctx *AssignRegularContext) {}

// EnterAssignSum is called when production assignSum is entered.
func (s *BaseMlanListener) EnterAssignSum(ctx *AssignSumContext) {}

// ExitAssignSum is called when production assignSum is exited.
func (s *BaseMlanListener) ExitAssignSum(ctx *AssignSumContext) {}

// EnterAssignSub is called when production assignSub is entered.
func (s *BaseMlanListener) EnterAssignSub(ctx *AssignSubContext) {}

// ExitAssignSub is called when production assignSub is exited.
func (s *BaseMlanListener) ExitAssignSub(ctx *AssignSubContext) {}

// EnterAssignMul is called when production assignMul is entered.
func (s *BaseMlanListener) EnterAssignMul(ctx *AssignMulContext) {}

// ExitAssignMul is called when production assignMul is exited.
func (s *BaseMlanListener) ExitAssignMul(ctx *AssignMulContext) {}

// EnterAssignDiv is called when production assignDiv is entered.
func (s *BaseMlanListener) EnterAssignDiv(ctx *AssignDivContext) {}

// ExitAssignDiv is called when production assignDiv is exited.
func (s *BaseMlanListener) ExitAssignDiv(ctx *AssignDivContext) {}

// EnterAssignMod is called when production assignMod is entered.
func (s *BaseMlanListener) EnterAssignMod(ctx *AssignModContext) {}

// ExitAssignMod is called when production assignMod is exited.
func (s *BaseMlanListener) ExitAssignMod(ctx *AssignModContext) {}

// EnterAssignPow is called when production assignPow is entered.
func (s *BaseMlanListener) EnterAssignPow(ctx *AssignPowContext) {}

// ExitAssignPow is called when production assignPow is exited.
func (s *BaseMlanListener) ExitAssignPow(ctx *AssignPowContext) {}

// EnterAssignIdxRegular is called when production assignIdxRegular is entered.
func (s *BaseMlanListener) EnterAssignIdxRegular(ctx *AssignIdxRegularContext) {}

// ExitAssignIdxRegular is called when production assignIdxRegular is exited.
func (s *BaseMlanListener) ExitAssignIdxRegular(ctx *AssignIdxRegularContext) {}

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

// EnterIdx is called when production idx is entered.
func (s *BaseMlanListener) EnterIdx(ctx *IdxContext) {}

// ExitIdx is called when production idx is exited.
func (s *BaseMlanListener) ExitIdx(ctx *IdxContext) {}

// EnterIdentifierMethodInvoke is called when production identifierMethodInvoke is entered.
func (s *BaseMlanListener) EnterIdentifierMethodInvoke(ctx *IdentifierMethodInvokeContext) {}

// ExitIdentifierMethodInvoke is called when production identifierMethodInvoke is exited.
func (s *BaseMlanListener) ExitIdentifierMethodInvoke(ctx *IdentifierMethodInvokeContext) {}

// EnterIdentifierFnInvoke is called when production identifierFnInvoke is entered.
func (s *BaseMlanListener) EnterIdentifierFnInvoke(ctx *IdentifierFnInvokeContext) {}

// ExitIdentifierFnInvoke is called when production identifierFnInvoke is exited.
func (s *BaseMlanListener) ExitIdentifierFnInvoke(ctx *IdentifierFnInvokeContext) {}

// EnterIdentifierCsInvoke is called when production identifierCsInvoke is entered.
func (s *BaseMlanListener) EnterIdentifierCsInvoke(ctx *IdentifierCsInvokeContext) {}

// ExitIdentifierCsInvoke is called when production identifierCsInvoke is exited.
func (s *BaseMlanListener) ExitIdentifierCsInvoke(ctx *IdentifierCsInvokeContext) {}

// EnterExpBool is called when production expBool is entered.
func (s *BaseMlanListener) EnterExpBool(ctx *ExpBoolContext) {}

// ExitExpBool is called when production expBool is exited.
func (s *BaseMlanListener) ExitExpBool(ctx *ExpBoolContext) {}

// EnterExpComparison is called when production expComparison is entered.
func (s *BaseMlanListener) EnterExpComparison(ctx *ExpComparisonContext) {}

// ExitExpComparison is called when production expComparison is exited.
func (s *BaseMlanListener) ExitExpComparison(ctx *ExpComparisonContext) {}

// EnterExpIdx is called when production expIdx is entered.
func (s *BaseMlanListener) EnterExpIdx(ctx *ExpIdxContext) {}

// ExitExpIdx is called when production expIdx is exited.
func (s *BaseMlanListener) ExitExpIdx(ctx *ExpIdxContext) {}

// EnterExpString is called when production expString is entered.
func (s *BaseMlanListener) EnterExpString(ctx *ExpStringContext) {}

// ExitExpString is called when production expString is exited.
func (s *BaseMlanListener) ExitExpString(ctx *ExpStringContext) {}

// EnterExpCsInvoke is called when production expCsInvoke is entered.
func (s *BaseMlanListener) EnterExpCsInvoke(ctx *ExpCsInvokeContext) {}

// ExitExpCsInvoke is called when production expCsInvoke is exited.
func (s *BaseMlanListener) ExitExpCsInvoke(ctx *ExpCsInvokeContext) {}

// EnterExpFloat is called when production expFloat is entered.
func (s *BaseMlanListener) EnterExpFloat(ctx *ExpFloatContext) {}

// ExitExpFloat is called when production expFloat is exited.
func (s *BaseMlanListener) ExitExpFloat(ctx *ExpFloatContext) {}

// EnterExpPow is called when production expPow is entered.
func (s *BaseMlanListener) EnterExpPow(ctx *ExpPowContext) {}

// ExitExpPow is called when production expPow is exited.
func (s *BaseMlanListener) ExitExpPow(ctx *ExpPowContext) {}

// EnterExpDict is called when production expDict is entered.
func (s *BaseMlanListener) EnterExpDict(ctx *ExpDictContext) {}

// ExitExpDict is called when production expDict is exited.
func (s *BaseMlanListener) ExitExpDict(ctx *ExpDictContext) {}

// EnterExpXor is called when production expXor is entered.
func (s *BaseMlanListener) EnterExpXor(ctx *ExpXorContext) {}

// ExitExpXor is called when production expXor is exited.
func (s *BaseMlanListener) ExitExpXor(ctx *ExpXorContext) {}

// EnterExpNeg is called when production expNeg is entered.
func (s *BaseMlanListener) EnterExpNeg(ctx *ExpNegContext) {}

// ExitExpNeg is called when production expNeg is exited.
func (s *BaseMlanListener) ExitExpNeg(ctx *ExpNegContext) {}

// EnterExpInteger is called when production expInteger is entered.
func (s *BaseMlanListener) EnterExpInteger(ctx *ExpIntegerContext) {}

// ExitExpInteger is called when production expInteger is exited.
func (s *BaseMlanListener) ExitExpInteger(ctx *ExpIntegerContext) {}

// EnterExpLogicalOr is called when production expLogicalOr is entered.
func (s *BaseMlanListener) EnterExpLogicalOr(ctx *ExpLogicalOrContext) {}

// ExitExpLogicalOr is called when production expLogicalOr is exited.
func (s *BaseMlanListener) ExitExpLogicalOr(ctx *ExpLogicalOrContext) {}

// EnterExpCs is called when production expCs is entered.
func (s *BaseMlanListener) EnterExpCs(ctx *ExpCsContext) {}

// ExitExpCs is called when production expCs is exited.
func (s *BaseMlanListener) ExitExpCs(ctx *ExpCsContext) {}

// EnterExpMulDivMod is called when production expMulDivMod is entered.
func (s *BaseMlanListener) EnterExpMulDivMod(ctx *ExpMulDivModContext) {}

// ExitExpMulDivMod is called when production expMulDivMod is exited.
func (s *BaseMlanListener) ExitExpMulDivMod(ctx *ExpMulDivModContext) {}

// EnterExpNull is called when production expNull is entered.
func (s *BaseMlanListener) EnterExpNull(ctx *ExpNullContext) {}

// ExitExpNull is called when production expNull is exited.
func (s *BaseMlanListener) ExitExpNull(ctx *ExpNullContext) {}

// EnterExpFnInvoke is called when production expFnInvoke is entered.
func (s *BaseMlanListener) EnterExpFnInvoke(ctx *ExpFnInvokeContext) {}

// ExitExpFnInvoke is called when production expFnInvoke is exited.
func (s *BaseMlanListener) ExitExpFnInvoke(ctx *ExpFnInvokeContext) {}

// EnterExpList is called when production expList is entered.
func (s *BaseMlanListener) EnterExpList(ctx *ExpListContext) {}

// ExitExpList is called when production expList is exited.
func (s *BaseMlanListener) ExitExpList(ctx *ExpListContext) {}

// EnterExpLogicalAnd is called when production expLogicalAnd is entered.
func (s *BaseMlanListener) EnterExpLogicalAnd(ctx *ExpLogicalAndContext) {}

// ExitExpLogicalAnd is called when production expLogicalAnd is exited.
func (s *BaseMlanListener) ExitExpLogicalAnd(ctx *ExpLogicalAndContext) {}

// EnterExpParentheses is called when production expParentheses is entered.
func (s *BaseMlanListener) EnterExpParentheses(ctx *ExpParenthesesContext) {}

// ExitExpParentheses is called when production expParentheses is exited.
func (s *BaseMlanListener) ExitExpParentheses(ctx *ExpParenthesesContext) {}

// EnterExpEqual is called when production expEqual is entered.
func (s *BaseMlanListener) EnterExpEqual(ctx *ExpEqualContext) {}

// ExitExpEqual is called when production expEqual is exited.
func (s *BaseMlanListener) ExitExpEqual(ctx *ExpEqualContext) {}

// EnterExpMethodInvoke is called when production expMethodInvoke is entered.
func (s *BaseMlanListener) EnterExpMethodInvoke(ctx *ExpMethodInvokeContext) {}

// ExitExpMethodInvoke is called when production expMethodInvoke is exited.
func (s *BaseMlanListener) ExitExpMethodInvoke(ctx *ExpMethodInvokeContext) {}

// EnterExpLogicalNot is called when production expLogicalNot is entered.
func (s *BaseMlanListener) EnterExpLogicalNot(ctx *ExpLogicalNotContext) {}

// ExitExpLogicalNot is called when production expLogicalNot is exited.
func (s *BaseMlanListener) ExitExpLogicalNot(ctx *ExpLogicalNotContext) {}

// EnterExpIntegerHex is called when production expIntegerHex is entered.
func (s *BaseMlanListener) EnterExpIntegerHex(ctx *ExpIntegerHexContext) {}

// ExitExpIntegerHex is called when production expIntegerHex is exited.
func (s *BaseMlanListener) ExitExpIntegerHex(ctx *ExpIntegerHexContext) {}

// EnterExpIdentifier is called when production expIdentifier is entered.
func (s *BaseMlanListener) EnterExpIdentifier(ctx *ExpIdentifierContext) {}

// ExitExpIdentifier is called when production expIdentifier is exited.
func (s *BaseMlanListener) ExitExpIdentifier(ctx *ExpIdentifierContext) {}

// EnterExpSumSub is called when production expSumSub is entered.
func (s *BaseMlanListener) EnterExpSumSub(ctx *ExpSumSubContext) {}

// ExitExpSumSub is called when production expSumSub is exited.
func (s *BaseMlanListener) ExitExpSumSub(ctx *ExpSumSubContext) {}

// EnterIfBlockStmt is called when production ifBlockStmt is entered.
func (s *BaseMlanListener) EnterIfBlockStmt(ctx *IfBlockStmtContext) {}

// ExitIfBlockStmt is called when production ifBlockStmt is exited.
func (s *BaseMlanListener) ExitIfBlockStmt(ctx *IfBlockStmtContext) {}

// EnterElifBlockStmt is called when production elifBlockStmt is entered.
func (s *BaseMlanListener) EnterElifBlockStmt(ctx *ElifBlockStmtContext) {}

// ExitElifBlockStmt is called when production elifBlockStmt is exited.
func (s *BaseMlanListener) ExitElifBlockStmt(ctx *ElifBlockStmtContext) {}

// EnterElseBlockStmt is called when production elseBlockStmt is entered.
func (s *BaseMlanListener) EnterElseBlockStmt(ctx *ElseBlockStmtContext) {}

// ExitElseBlockStmt is called when production elseBlockStmt is exited.
func (s *BaseMlanListener) ExitElseBlockStmt(ctx *ElseBlockStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseMlanListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseMlanListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterFnParams is called when production fnParams is entered.
func (s *BaseMlanListener) EnterFnParams(ctx *FnParamsContext) {}

// ExitFnParams is called when production fnParams is exited.
func (s *BaseMlanListener) ExitFnParams(ctx *FnParamsContext) {}

// EnterFnBody is called when production fnBody is entered.
func (s *BaseMlanListener) EnterFnBody(ctx *FnBodyContext) {}

// ExitFnBody is called when production fnBody is exited.
func (s *BaseMlanListener) ExitFnBody(ctx *FnBodyContext) {}

// EnterFn is called when production fn is entered.
func (s *BaseMlanListener) EnterFn(ctx *FnContext) {}

// ExitFn is called when production fn is exited.
func (s *BaseMlanListener) ExitFn(ctx *FnContext) {}

// EnterClosure is called when production closure is entered.
func (s *BaseMlanListener) EnterClosure(ctx *ClosureContext) {}

// ExitClosure is called when production closure is exited.
func (s *BaseMlanListener) ExitClosure(ctx *ClosureContext) {}

// EnterInclude is called when production include is entered.
func (s *BaseMlanListener) EnterInclude(ctx *IncludeContext) {}

// ExitInclude is called when production include is exited.
func (s *BaseMlanListener) ExitInclude(ctx *IncludeContext) {}
