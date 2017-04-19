// Generated from src/github.com/nyaruka/goantlr/gen/Excellent.g4 by ANTLR 4.7.

package gen // Excellent
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ExcellentParser.
type ExcellentVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExcellentParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by ExcellentParser#expAtom.
	VisitExpAtom(ctx *ExpAtomContext) interface{}

	// Visit a parse tree produced by ExcellentParser#expConcatenation.
	VisitExpConcatenation(ctx *ExpConcatenationContext) interface{}

	// Visit a parse tree produced by ExcellentParser#expEquality.
	VisitExpEquality(ctx *ExpEqualityContext) interface{}

	// Visit a parse tree produced by ExcellentParser#expComparison.
	VisitExpComparison(ctx *ExpComparisonContext) interface{}

	// Visit a parse tree produced by ExcellentParser#expAddition.
	VisitExpAddition(ctx *ExpAdditionContext) interface{}

	// Visit a parse tree produced by ExcellentParser#expMultiplication.
	VisitExpMultiplication(ctx *ExpMultiplicationContext) interface{}

	// Visit a parse tree produced by ExcellentParser#expExponent.
	VisitExpExponent(ctx *ExpExponentContext) interface{}

	// Visit a parse tree produced by ExcellentParser#expUnary.
	VisitExpUnary(ctx *ExpUnaryContext) interface{}

	// Visit a parse tree produced by ExcellentParser#path.
	VisitPath(ctx *PathContext) interface{}

	// Visit a parse tree produced by ExcellentParser#step.
	VisitStep(ctx *StepContext) interface{}

	// Visit a parse tree produced by ExcellentParser#functionParameters.
	VisitFunctionParameters(ctx *FunctionParametersContext) interface{}

	// Visit a parse tree produced by ExcellentParser#concatenation.
	VisitConcatenation(ctx *ConcatenationContext) interface{}

	// Visit a parse tree produced by ExcellentParser#equality.
	VisitEquality(ctx *EqualityContext) interface{}

	// Visit a parse tree produced by ExcellentParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by ExcellentParser#addition.
	VisitAddition(ctx *AdditionContext) interface{}

	// Visit a parse tree produced by ExcellentParser#multiplication.
	VisitMultiplication(ctx *MultiplicationContext) interface{}

	// Visit a parse tree produced by ExcellentParser#exponent.
	VisitExponent(ctx *ExponentContext) interface{}

	// Visit a parse tree produced by ExcellentParser#negation.
	VisitNegation(ctx *NegationContext) interface{}

	// Visit a parse tree produced by ExcellentParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by ExcellentParser#functionPath.
	VisitFunctionPath(ctx *FunctionPathContext) interface{}

	// Visit a parse tree produced by ExcellentParser#contextReference.
	VisitContextReference(ctx *ContextReferenceContext) interface{}

	// Visit a parse tree produced by ExcellentParser#atomFuncCall.
	VisitAtomFuncCall(ctx *AtomFuncCallContext) interface{}

	// Visit a parse tree produced by ExcellentParser#atomFuncPath.
	VisitAtomFuncPath(ctx *AtomFuncPathContext) interface{}

	// Visit a parse tree produced by ExcellentParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by ExcellentParser#decimalLiteral.
	VisitDecimalLiteral(ctx *DecimalLiteralContext) interface{}

	// Visit a parse tree produced by ExcellentParser#parentheses.
	VisitParentheses(ctx *ParenthesesContext) interface{}

	// Visit a parse tree produced by ExcellentParser#true.
	VisitTrue(ctx *TrueContext) interface{}

	// Visit a parse tree produced by ExcellentParser#false.
	VisitFalse(ctx *FalseContext) interface{}
}
