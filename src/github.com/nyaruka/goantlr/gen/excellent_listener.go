// Generated from src/github.com/nyaruka/goantlr/gen/Excellent.g4 by ANTLR 4.7.

package gen // Excellent
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExcellentListener is a complete listener for a parse tree produced by ExcellentParser.
type ExcellentListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterExpAtom is called when entering the expAtom production.
	EnterExpAtom(c *ExpAtomContext)

	// EnterExpConcatenation is called when entering the expConcatenation production.
	EnterExpConcatenation(c *ExpConcatenationContext)

	// EnterExpEquality is called when entering the expEquality production.
	EnterExpEquality(c *ExpEqualityContext)

	// EnterExpComparison is called when entering the expComparison production.
	EnterExpComparison(c *ExpComparisonContext)

	// EnterExpAddition is called when entering the expAddition production.
	EnterExpAddition(c *ExpAdditionContext)

	// EnterExpMultiplication is called when entering the expMultiplication production.
	EnterExpMultiplication(c *ExpMultiplicationContext)

	// EnterExpExponent is called when entering the expExponent production.
	EnterExpExponent(c *ExpExponentContext)

	// EnterExpUnary is called when entering the expUnary production.
	EnterExpUnary(c *ExpUnaryContext)

	// EnterPath is called when entering the path production.
	EnterPath(c *PathContext)

	// EnterStep is called when entering the step production.
	EnterStep(c *StepContext)

	// EnterFunctionParameters is called when entering the functionParameters production.
	EnterFunctionParameters(c *FunctionParametersContext)

	// EnterConcatenation is called when entering the concatenation production.
	EnterConcatenation(c *ConcatenationContext)

	// EnterEquality is called when entering the equality production.
	EnterEquality(c *EqualityContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterAddition is called when entering the addition production.
	EnterAddition(c *AdditionContext)

	// EnterMultiplication is called when entering the multiplication production.
	EnterMultiplication(c *MultiplicationContext)

	// EnterExponent is called when entering the exponent production.
	EnterExponent(c *ExponentContext)

	// EnterNegation is called when entering the negation production.
	EnterNegation(c *NegationContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterFunctionPath is called when entering the functionPath production.
	EnterFunctionPath(c *FunctionPathContext)

	// EnterContextReference is called when entering the contextReference production.
	EnterContextReference(c *ContextReferenceContext)

	// EnterAtomFuncCall is called when entering the atomFuncCall production.
	EnterAtomFuncCall(c *AtomFuncCallContext)

	// EnterAtomFuncPath is called when entering the atomFuncPath production.
	EnterAtomFuncPath(c *AtomFuncPathContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterDecimalLiteral is called when entering the decimalLiteral production.
	EnterDecimalLiteral(c *DecimalLiteralContext)

	// EnterParentheses is called when entering the parentheses production.
	EnterParentheses(c *ParenthesesContext)

	// EnterTrue is called when entering the true production.
	EnterTrue(c *TrueContext)

	// EnterFalse is called when entering the false production.
	EnterFalse(c *FalseContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitExpAtom is called when exiting the expAtom production.
	ExitExpAtom(c *ExpAtomContext)

	// ExitExpConcatenation is called when exiting the expConcatenation production.
	ExitExpConcatenation(c *ExpConcatenationContext)

	// ExitExpEquality is called when exiting the expEquality production.
	ExitExpEquality(c *ExpEqualityContext)

	// ExitExpComparison is called when exiting the expComparison production.
	ExitExpComparison(c *ExpComparisonContext)

	// ExitExpAddition is called when exiting the expAddition production.
	ExitExpAddition(c *ExpAdditionContext)

	// ExitExpMultiplication is called when exiting the expMultiplication production.
	ExitExpMultiplication(c *ExpMultiplicationContext)

	// ExitExpExponent is called when exiting the expExponent production.
	ExitExpExponent(c *ExpExponentContext)

	// ExitExpUnary is called when exiting the expUnary production.
	ExitExpUnary(c *ExpUnaryContext)

	// ExitPath is called when exiting the path production.
	ExitPath(c *PathContext)

	// ExitStep is called when exiting the step production.
	ExitStep(c *StepContext)

	// ExitFunctionParameters is called when exiting the functionParameters production.
	ExitFunctionParameters(c *FunctionParametersContext)

	// ExitConcatenation is called when exiting the concatenation production.
	ExitConcatenation(c *ConcatenationContext)

	// ExitEquality is called when exiting the equality production.
	ExitEquality(c *EqualityContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitAddition is called when exiting the addition production.
	ExitAddition(c *AdditionContext)

	// ExitMultiplication is called when exiting the multiplication production.
	ExitMultiplication(c *MultiplicationContext)

	// ExitExponent is called when exiting the exponent production.
	ExitExponent(c *ExponentContext)

	// ExitNegation is called when exiting the negation production.
	ExitNegation(c *NegationContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitFunctionPath is called when exiting the functionPath production.
	ExitFunctionPath(c *FunctionPathContext)

	// ExitContextReference is called when exiting the contextReference production.
	ExitContextReference(c *ContextReferenceContext)

	// ExitAtomFuncCall is called when exiting the atomFuncCall production.
	ExitAtomFuncCall(c *AtomFuncCallContext)

	// ExitAtomFuncPath is called when exiting the atomFuncPath production.
	ExitAtomFuncPath(c *AtomFuncPathContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitDecimalLiteral is called when exiting the decimalLiteral production.
	ExitDecimalLiteral(c *DecimalLiteralContext)

	// ExitParentheses is called when exiting the parentheses production.
	ExitParentheses(c *ParenthesesContext)

	// ExitTrue is called when exiting the true production.
	ExitTrue(c *TrueContext)

	// ExitFalse is called when exiting the false production.
	ExitFalse(c *FalseContext)
}
