// Code generated from ./Mlan.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Mlan
import "github.com/antlr4-go/antlr/v4"

type BaseMlanVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMlanVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitAssignRegular(ctx *AssignRegularContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitAssignSum(ctx *AssignSumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitAssignSub(ctx *AssignSubContext) interface{} {
	return v.VisitChildren(ctx)
}

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

func (v *BaseMlanVisitor) VisitAssignIdxRegular(ctx *AssignIdxRegularContext) interface{} {
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

func (v *BaseMlanVisitor) VisitIdx(ctx *IdxContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitIdentifierMethodInvoke(ctx *IdentifierMethodInvokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitIdentifierFnInvoke(ctx *IdentifierFnInvokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitIdentifierCsInvoke(ctx *IdentifierCsInvokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpBool(ctx *ExpBoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpComparison(ctx *ExpComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpIdx(ctx *ExpIdxContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpString(ctx *ExpStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpCsInvoke(ctx *ExpCsInvokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpFloat(ctx *ExpFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpPow(ctx *ExpPowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpDict(ctx *ExpDictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpXor(ctx *ExpXorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpNeg(ctx *ExpNegContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpInteger(ctx *ExpIntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpLogicalOr(ctx *ExpLogicalOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpCs(ctx *ExpCsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpMulDivMod(ctx *ExpMulDivModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpNull(ctx *ExpNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpFnInvoke(ctx *ExpFnInvokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpList(ctx *ExpListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpLogicalAnd(ctx *ExpLogicalAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpParentheses(ctx *ExpParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpEqual(ctx *ExpEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpMethodInvoke(ctx *ExpMethodInvokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpLogicalNot(ctx *ExpLogicalNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpIntegerHex(ctx *ExpIntegerHexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpIdentifier(ctx *ExpIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitExpSumSub(ctx *ExpSumSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitIfBlockStmt(ctx *IfBlockStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitElifBlockStmt(ctx *ElifBlockStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitElseBlockStmt(ctx *ElseBlockStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitFnParams(ctx *FnParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitFnBody(ctx *FnBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitFn(ctx *FnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitClosure(ctx *ClosureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMlanVisitor) VisitInclude(ctx *IncludeContext) interface{} {
	return v.VisitChildren(ctx)
}
