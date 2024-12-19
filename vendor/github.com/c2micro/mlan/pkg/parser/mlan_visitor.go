// Code generated from ./Mlan.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Mlan
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MlanParser.
type MlanVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MlanParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by MlanParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by MlanParser#whileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by MlanParser#forStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by MlanParser#returnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by MlanParser#continueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by MlanParser#breakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by MlanParser#assignRegular.
	VisitAssignRegular(ctx *AssignRegularContext) interface{}

	// Visit a parse tree produced by MlanParser#assignSum.
	VisitAssignSum(ctx *AssignSumContext) interface{}

	// Visit a parse tree produced by MlanParser#assignSub.
	VisitAssignSub(ctx *AssignSubContext) interface{}

	// Visit a parse tree produced by MlanParser#assignMul.
	VisitAssignMul(ctx *AssignMulContext) interface{}

	// Visit a parse tree produced by MlanParser#assignDiv.
	VisitAssignDiv(ctx *AssignDivContext) interface{}

	// Visit a parse tree produced by MlanParser#assignMod.
	VisitAssignMod(ctx *AssignModContext) interface{}

	// Visit a parse tree produced by MlanParser#assignPow.
	VisitAssignPow(ctx *AssignPowContext) interface{}

	// Visit a parse tree produced by MlanParser#assignIdxRegular.
	VisitAssignIdxRegular(ctx *AssignIdxRegularContext) interface{}

	// Visit a parse tree produced by MlanParser#list.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by MlanParser#dictUnit.
	VisitDictUnit(ctx *DictUnitContext) interface{}

	// Visit a parse tree produced by MlanParser#dict.
	VisitDict(ctx *DictContext) interface{}

	// Visit a parse tree produced by MlanParser#idx.
	VisitIdx(ctx *IdxContext) interface{}

	// Visit a parse tree produced by MlanParser#identifierMethodInvoke.
	VisitIdentifierMethodInvoke(ctx *IdentifierMethodInvokeContext) interface{}

	// Visit a parse tree produced by MlanParser#identifierFnInvoke.
	VisitIdentifierFnInvoke(ctx *IdentifierFnInvokeContext) interface{}

	// Visit a parse tree produced by MlanParser#identifierCsInvoke.
	VisitIdentifierCsInvoke(ctx *IdentifierCsInvokeContext) interface{}

	// Visit a parse tree produced by MlanParser#expBool.
	VisitExpBool(ctx *ExpBoolContext) interface{}

	// Visit a parse tree produced by MlanParser#expComparison.
	VisitExpComparison(ctx *ExpComparisonContext) interface{}

	// Visit a parse tree produced by MlanParser#expIdx.
	VisitExpIdx(ctx *ExpIdxContext) interface{}

	// Visit a parse tree produced by MlanParser#expString.
	VisitExpString(ctx *ExpStringContext) interface{}

	// Visit a parse tree produced by MlanParser#expCsInvoke.
	VisitExpCsInvoke(ctx *ExpCsInvokeContext) interface{}

	// Visit a parse tree produced by MlanParser#expFloat.
	VisitExpFloat(ctx *ExpFloatContext) interface{}

	// Visit a parse tree produced by MlanParser#expPow.
	VisitExpPow(ctx *ExpPowContext) interface{}

	// Visit a parse tree produced by MlanParser#expDict.
	VisitExpDict(ctx *ExpDictContext) interface{}

	// Visit a parse tree produced by MlanParser#expXor.
	VisitExpXor(ctx *ExpXorContext) interface{}

	// Visit a parse tree produced by MlanParser#expNeg.
	VisitExpNeg(ctx *ExpNegContext) interface{}

	// Visit a parse tree produced by MlanParser#expInteger.
	VisitExpInteger(ctx *ExpIntegerContext) interface{}

	// Visit a parse tree produced by MlanParser#expLogicalOr.
	VisitExpLogicalOr(ctx *ExpLogicalOrContext) interface{}

	// Visit a parse tree produced by MlanParser#expCs.
	VisitExpCs(ctx *ExpCsContext) interface{}

	// Visit a parse tree produced by MlanParser#expMulDivMod.
	VisitExpMulDivMod(ctx *ExpMulDivModContext) interface{}

	// Visit a parse tree produced by MlanParser#expNull.
	VisitExpNull(ctx *ExpNullContext) interface{}

	// Visit a parse tree produced by MlanParser#expFnInvoke.
	VisitExpFnInvoke(ctx *ExpFnInvokeContext) interface{}

	// Visit a parse tree produced by MlanParser#expList.
	VisitExpList(ctx *ExpListContext) interface{}

	// Visit a parse tree produced by MlanParser#expLogicalAnd.
	VisitExpLogicalAnd(ctx *ExpLogicalAndContext) interface{}

	// Visit a parse tree produced by MlanParser#expParentheses.
	VisitExpParentheses(ctx *ExpParenthesesContext) interface{}

	// Visit a parse tree produced by MlanParser#expEqual.
	VisitExpEqual(ctx *ExpEqualContext) interface{}

	// Visit a parse tree produced by MlanParser#expMethodInvoke.
	VisitExpMethodInvoke(ctx *ExpMethodInvokeContext) interface{}

	// Visit a parse tree produced by MlanParser#expLogicalNot.
	VisitExpLogicalNot(ctx *ExpLogicalNotContext) interface{}

	// Visit a parse tree produced by MlanParser#expIntegerHex.
	VisitExpIntegerHex(ctx *ExpIntegerHexContext) interface{}

	// Visit a parse tree produced by MlanParser#expIdentifier.
	VisitExpIdentifier(ctx *ExpIdentifierContext) interface{}

	// Visit a parse tree produced by MlanParser#expSumSub.
	VisitExpSumSub(ctx *ExpSumSubContext) interface{}

	// Visit a parse tree produced by MlanParser#ifBlockStmt.
	VisitIfBlockStmt(ctx *IfBlockStmtContext) interface{}

	// Visit a parse tree produced by MlanParser#elifBlockStmt.
	VisitElifBlockStmt(ctx *ElifBlockStmtContext) interface{}

	// Visit a parse tree produced by MlanParser#elseBlockStmt.
	VisitElseBlockStmt(ctx *ElseBlockStmtContext) interface{}

	// Visit a parse tree produced by MlanParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by MlanParser#fnParams.
	VisitFnParams(ctx *FnParamsContext) interface{}

	// Visit a parse tree produced by MlanParser#fnBody.
	VisitFnBody(ctx *FnBodyContext) interface{}

	// Visit a parse tree produced by MlanParser#fn.
	VisitFn(ctx *FnContext) interface{}

	// Visit a parse tree produced by MlanParser#closure.
	VisitClosure(ctx *ClosureContext) interface{}

	// Visit a parse tree produced by MlanParser#include.
	VisitInclude(ctx *IncludeContext) interface{}
}
