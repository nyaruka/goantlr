// Generated from src/github.com/nyaruka/goantlr/gen/Excellent.g4 by ANTLR 4.7.

package gen // Excellent
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExcellentListener is a complete listener for a parse tree produced by ExcellentParser.
type BaseExcellentListener struct{}

var _ ExcellentListener = &BaseExcellentListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExcellentListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExcellentListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExcellentListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExcellentListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseExcellentListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseExcellentListener) ExitParse(ctx *ParseContext) {}

// EnterExpAtom is called when production expAtom is entered.
func (s *BaseExcellentListener) EnterExpAtom(ctx *ExpAtomContext) {}

// ExitExpAtom is called when production expAtom is exited.
func (s *BaseExcellentListener) ExitExpAtom(ctx *ExpAtomContext) {}

// EnterExpConcatenation is called when production expConcatenation is entered.
func (s *BaseExcellentListener) EnterExpConcatenation(ctx *ExpConcatenationContext) {}

// ExitExpConcatenation is called when production expConcatenation is exited.
func (s *BaseExcellentListener) ExitExpConcatenation(ctx *ExpConcatenationContext) {}

// EnterExpEquality is called when production expEquality is entered.
func (s *BaseExcellentListener) EnterExpEquality(ctx *ExpEqualityContext) {}

// ExitExpEquality is called when production expEquality is exited.
func (s *BaseExcellentListener) ExitExpEquality(ctx *ExpEqualityContext) {}

// EnterExpComparison is called when production expComparison is entered.
func (s *BaseExcellentListener) EnterExpComparison(ctx *ExpComparisonContext) {}

// ExitExpComparison is called when production expComparison is exited.
func (s *BaseExcellentListener) ExitExpComparison(ctx *ExpComparisonContext) {}

// EnterExpAddition is called when production expAddition is entered.
func (s *BaseExcellentListener) EnterExpAddition(ctx *ExpAdditionContext) {}

// ExitExpAddition is called when production expAddition is exited.
func (s *BaseExcellentListener) ExitExpAddition(ctx *ExpAdditionContext) {}

// EnterExpMultiplication is called when production expMultiplication is entered.
func (s *BaseExcellentListener) EnterExpMultiplication(ctx *ExpMultiplicationContext) {}

// ExitExpMultiplication is called when production expMultiplication is exited.
func (s *BaseExcellentListener) ExitExpMultiplication(ctx *ExpMultiplicationContext) {}

// EnterExpExponent is called when production expExponent is entered.
func (s *BaseExcellentListener) EnterExpExponent(ctx *ExpExponentContext) {}

// ExitExpExponent is called when production expExponent is exited.
func (s *BaseExcellentListener) ExitExpExponent(ctx *ExpExponentContext) {}

// EnterExpUnary is called when production expUnary is entered.
func (s *BaseExcellentListener) EnterExpUnary(ctx *ExpUnaryContext) {}

// ExitExpUnary is called when production expUnary is exited.
func (s *BaseExcellentListener) ExitExpUnary(ctx *ExpUnaryContext) {}

// EnterPath is called when production path is entered.
func (s *BaseExcellentListener) EnterPath(ctx *PathContext) {}

// ExitPath is called when production path is exited.
func (s *BaseExcellentListener) ExitPath(ctx *PathContext) {}

// EnterStep is called when production step is entered.
func (s *BaseExcellentListener) EnterStep(ctx *StepContext) {}

// ExitStep is called when production step is exited.
func (s *BaseExcellentListener) ExitStep(ctx *StepContext) {}

// EnterFunctionParameters is called when production functionParameters is entered.
func (s *BaseExcellentListener) EnterFunctionParameters(ctx *FunctionParametersContext) {}

// ExitFunctionParameters is called when production functionParameters is exited.
func (s *BaseExcellentListener) ExitFunctionParameters(ctx *FunctionParametersContext) {}

// EnterConcatenation is called when production concatenation is entered.
func (s *BaseExcellentListener) EnterConcatenation(ctx *ConcatenationContext) {}

// ExitConcatenation is called when production concatenation is exited.
func (s *BaseExcellentListener) ExitConcatenation(ctx *ConcatenationContext) {}

// EnterEquality is called when production equality is entered.
func (s *BaseExcellentListener) EnterEquality(ctx *EqualityContext) {}

// ExitEquality is called when production equality is exited.
func (s *BaseExcellentListener) ExitEquality(ctx *EqualityContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseExcellentListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseExcellentListener) ExitComparison(ctx *ComparisonContext) {}

// EnterAddition is called when production addition is entered.
func (s *BaseExcellentListener) EnterAddition(ctx *AdditionContext) {}

// ExitAddition is called when production addition is exited.
func (s *BaseExcellentListener) ExitAddition(ctx *AdditionContext) {}

// EnterMultiplication is called when production multiplication is entered.
func (s *BaseExcellentListener) EnterMultiplication(ctx *MultiplicationContext) {}

// ExitMultiplication is called when production multiplication is exited.
func (s *BaseExcellentListener) ExitMultiplication(ctx *MultiplicationContext) {}

// EnterExponent is called when production exponent is entered.
func (s *BaseExcellentListener) EnterExponent(ctx *ExponentContext) {}

// ExitExponent is called when production exponent is exited.
func (s *BaseExcellentListener) ExitExponent(ctx *ExponentContext) {}

// EnterNegation is called when production negation is entered.
func (s *BaseExcellentListener) EnterNegation(ctx *NegationContext) {}

// ExitNegation is called when production negation is exited.
func (s *BaseExcellentListener) ExitNegation(ctx *NegationContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseExcellentListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseExcellentListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterFunctionPath is called when production functionPath is entered.
func (s *BaseExcellentListener) EnterFunctionPath(ctx *FunctionPathContext) {}

// ExitFunctionPath is called when production functionPath is exited.
func (s *BaseExcellentListener) ExitFunctionPath(ctx *FunctionPathContext) {}

// EnterContextReference is called when production contextReference is entered.
func (s *BaseExcellentListener) EnterContextReference(ctx *ContextReferenceContext) {}

// ExitContextReference is called when production contextReference is exited.
func (s *BaseExcellentListener) ExitContextReference(ctx *ContextReferenceContext) {}

// EnterAtomFuncCall is called when production atomFuncCall is entered.
func (s *BaseExcellentListener) EnterAtomFuncCall(ctx *AtomFuncCallContext) {}

// ExitAtomFuncCall is called when production atomFuncCall is exited.
func (s *BaseExcellentListener) ExitAtomFuncCall(ctx *AtomFuncCallContext) {}

// EnterAtomFuncPath is called when production atomFuncPath is entered.
func (s *BaseExcellentListener) EnterAtomFuncPath(ctx *AtomFuncPathContext) {}

// ExitAtomFuncPath is called when production atomFuncPath is exited.
func (s *BaseExcellentListener) ExitAtomFuncPath(ctx *AtomFuncPathContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseExcellentListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseExcellentListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterDecimalLiteral is called when production decimalLiteral is entered.
func (s *BaseExcellentListener) EnterDecimalLiteral(ctx *DecimalLiteralContext) {}

// ExitDecimalLiteral is called when production decimalLiteral is exited.
func (s *BaseExcellentListener) ExitDecimalLiteral(ctx *DecimalLiteralContext) {}

// EnterParentheses is called when production parentheses is entered.
func (s *BaseExcellentListener) EnterParentheses(ctx *ParenthesesContext) {}

// ExitParentheses is called when production parentheses is exited.
func (s *BaseExcellentListener) ExitParentheses(ctx *ParenthesesContext) {}

// EnterTrue is called when production true is entered.
func (s *BaseExcellentListener) EnterTrue(ctx *TrueContext) {}

// ExitTrue is called when production true is exited.
func (s *BaseExcellentListener) ExitTrue(ctx *TrueContext) {}

// EnterFalse is called when production false is entered.
func (s *BaseExcellentListener) EnterFalse(ctx *FalseContext) {}

// ExitFalse is called when production false is exited.
func (s *BaseExcellentListener) ExitFalse(ctx *FalseContext) {}
