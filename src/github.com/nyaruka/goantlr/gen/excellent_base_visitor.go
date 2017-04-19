// Generated from src/github.com/nyaruka/goantlr/gen/Excellent.g4 by ANTLR 4.7.

package gen // Excellent
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseExcellentVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseExcellentVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitExpAtom(ctx *ExpAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitExpConcatenation(ctx *ExpConcatenationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitExpEquality(ctx *ExpEqualityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitExpComparison(ctx *ExpComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitExpAddition(ctx *ExpAdditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitExpMultiplication(ctx *ExpMultiplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitExpExponent(ctx *ExpExponentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitExpUnary(ctx *ExpUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitPath(ctx *PathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitStep(ctx *StepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitFunctionParameters(ctx *FunctionParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitConcatenation(ctx *ConcatenationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitEquality(ctx *EqualityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitAddition(ctx *AdditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitMultiplication(ctx *MultiplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitExponent(ctx *ExponentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitNegation(ctx *NegationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitFunctionPath(ctx *FunctionPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitContextReference(ctx *ContextReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitAtomFuncCall(ctx *AtomFuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitAtomFuncPath(ctx *AtomFuncPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitParentheses(ctx *ParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitTrue(ctx *TrueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExcellentVisitor) VisitFalse(ctx *FalseContext) interface{} {
	return v.VisitChildren(ctx)
}
