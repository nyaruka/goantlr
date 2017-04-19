// Generated from src/github.com/nyaruka/goantlr/gen/Excellent.g4 by ANTLR 4.7.

package gen // Excellent
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 28, 138,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 44, 10, 3, 3, 4, 3, 4,
	7, 4, 48, 10, 4, 12, 4, 14, 4, 51, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 5, 5, 61, 10, 5, 3, 6, 3, 6, 3, 6, 7, 6, 66, 10, 6, 12,
	6, 14, 6, 69, 11, 6, 3, 7, 3, 7, 3, 7, 5, 7, 74, 10, 7, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 9, 5, 9, 83, 10, 9, 3, 10, 3, 10, 3, 10, 7, 10,
	88, 10, 10, 12, 10, 14, 10, 91, 11, 10, 3, 11, 3, 11, 3, 11, 7, 11, 96,
	10, 11, 12, 11, 14, 11, 99, 11, 11, 3, 12, 3, 12, 3, 12, 5, 12, 104, 10,
	12, 3, 13, 5, 13, 107, 10, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 5, 14,
	114, 10, 14, 3, 14, 3, 14, 3, 15, 3, 15, 7, 15, 120, 10, 15, 12, 15, 14,
	15, 123, 11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 5, 16, 136, 10, 16, 3, 16, 2, 2, 17, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 2, 6, 3, 2, 18, 19, 4, 2, 16, 17,
	20, 21, 3, 2, 11, 12, 3, 2, 13, 14, 2, 148, 2, 32, 3, 2, 2, 2, 4, 43, 3,
	2, 2, 2, 6, 45, 3, 2, 2, 2, 8, 60, 3, 2, 2, 2, 10, 62, 3, 2, 2, 2, 12,
	70, 3, 2, 2, 2, 14, 75, 3, 2, 2, 2, 16, 79, 3, 2, 2, 2, 18, 84, 3, 2, 2,
	2, 20, 92, 3, 2, 2, 2, 22, 100, 3, 2, 2, 2, 24, 106, 3, 2, 2, 2, 26, 110,
	3, 2, 2, 2, 28, 117, 3, 2, 2, 2, 30, 135, 3, 2, 2, 2, 32, 33, 5, 4, 3,
	2, 33, 34, 7, 2, 2, 3, 34, 3, 3, 2, 2, 2, 35, 44, 5, 30, 16, 2, 36, 44,
	5, 12, 7, 2, 37, 44, 5, 14, 8, 2, 38, 44, 5, 16, 9, 2, 39, 44, 5, 18, 10,
	2, 40, 44, 5, 20, 11, 2, 41, 44, 5, 22, 12, 2, 42, 44, 5, 24, 13, 2, 43,
	35, 3, 2, 2, 2, 43, 36, 3, 2, 2, 2, 43, 37, 3, 2, 2, 2, 43, 38, 3, 2, 2,
	2, 43, 39, 3, 2, 2, 2, 43, 40, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 42,
	3, 2, 2, 2, 44, 5, 3, 2, 2, 2, 45, 49, 7, 27, 2, 2, 46, 48, 5, 8, 5, 2,
	47, 46, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50, 3,
	2, 2, 2, 50, 7, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 53, 7, 9, 2, 2, 53,
	54, 5, 4, 3, 2, 54, 55, 7, 10, 2, 2, 55, 61, 3, 2, 2, 2, 56, 57, 7, 6,
	2, 2, 57, 61, 7, 27, 2, 2, 58, 59, 7, 6, 2, 2, 59, 61, 7, 3, 2, 2, 60,
	52, 3, 2, 2, 2, 60, 56, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 61, 9, 3, 2, 2,
	2, 62, 67, 5, 4, 3, 2, 63, 64, 7, 15, 2, 2, 64, 66, 5, 4, 3, 2, 65, 63,
	3, 2, 2, 2, 66, 69, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2,
	68, 11, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 70, 73, 5, 30, 16, 2, 71, 72, 7,
	24, 2, 2, 72, 74, 5, 12, 7, 2, 73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2,
	74, 13, 3, 2, 2, 2, 75, 76, 5, 16, 9, 2, 76, 77, 9, 2, 2, 2, 77, 78, 5,
	16, 9, 2, 78, 15, 3, 2, 2, 2, 79, 82, 5, 18, 10, 2, 80, 81, 9, 3, 2, 2,
	81, 83, 5, 18, 10, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 17, 3,
	2, 2, 2, 84, 89, 5, 20, 11, 2, 85, 86, 9, 4, 2, 2, 86, 88, 5, 20, 11, 2,
	87, 85, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3,
	2, 2, 2, 90, 19, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 97, 5, 22, 12, 2,
	93, 94, 9, 5, 2, 2, 94, 96, 5, 22, 12, 2, 95, 93, 3, 2, 2, 2, 96, 99, 3,
	2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 21, 3, 2, 2, 2, 99,
	97, 3, 2, 2, 2, 100, 103, 5, 24, 13, 2, 101, 102, 7, 23, 2, 2, 102, 104,
	5, 22, 12, 2, 103, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 23, 3, 2,
	2, 2, 105, 107, 7, 11, 2, 2, 106, 105, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2,
	107, 108, 3, 2, 2, 2, 108, 109, 5, 30, 16, 2, 109, 25, 3, 2, 2, 2, 110,
	111, 7, 27, 2, 2, 111, 113, 7, 7, 2, 2, 112, 114, 5, 10, 6, 2, 113, 112,
	3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 116, 7, 8,
	2, 2, 116, 27, 3, 2, 2, 2, 117, 121, 5, 26, 14, 2, 118, 120, 5, 8, 5, 2,
	119, 118, 3, 2, 2, 2, 120, 123, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 121,
	122, 3, 2, 2, 2, 122, 29, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2, 124, 136, 5,
	6, 4, 2, 125, 136, 5, 26, 14, 2, 126, 136, 5, 28, 15, 2, 127, 136, 7, 25,
	2, 2, 128, 136, 7, 3, 2, 2, 129, 130, 7, 7, 2, 2, 130, 131, 5, 4, 3, 2,
	131, 132, 7, 8, 2, 2, 132, 136, 3, 2, 2, 2, 133, 136, 7, 4, 2, 2, 134,
	136, 7, 5, 2, 2, 135, 124, 3, 2, 2, 2, 135, 125, 3, 2, 2, 2, 135, 126,
	3, 2, 2, 2, 135, 127, 3, 2, 2, 2, 135, 128, 3, 2, 2, 2, 135, 129, 3, 2,
	2, 2, 135, 133, 3, 2, 2, 2, 135, 134, 3, 2, 2, 2, 136, 31, 3, 2, 2, 2,
	15, 43, 49, 60, 67, 73, 82, 89, 97, 103, 106, 113, 121, 135,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "'.'", "'('", "')'", "'['", "']'", "'-'", "'+'", "'*'",
	"'/'", "','", "'<'", "'>'", "'='", "'!='", "'<='", "'>='", "'\"'", "'^'",
	"'&'",
}
var symbolicNames = []string{
	"", "NUMBER", "TRUE", "FALSE", "PATHSEP", "LPAR", "RPAR", "LBRAC", "RBRAC",
	"SUB", "ADD", "MUL", "DIV", "COMMA", "LT", "GT", "EQ", "NE", "LTE", "GTE",
	"QUOT", "EXP", "AMP", "LITERAL", "Whitespace", "NAME", "ERRROR_CHAR",
}

var ruleNames = []string{
	"parse", "expr", "path", "step", "parameters", "concatenationExpr", "equalityExpr",
	"comparisonExpr", "additionExpr", "multiplicationExpr", "exponentExpr",
	"unaryExpr", "funcCall", "funcPath", "atom",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ExcellentParser struct {
	*antlr.BaseParser
}

func NewExcellentParser(input antlr.TokenStream) *ExcellentParser {
	this := new(ExcellentParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Excellent.g4"

	return this
}

// ExcellentParser tokens.
const (
	ExcellentParserEOF         = antlr.TokenEOF
	ExcellentParserNUMBER      = 1
	ExcellentParserTRUE        = 2
	ExcellentParserFALSE       = 3
	ExcellentParserPATHSEP     = 4
	ExcellentParserLPAR        = 5
	ExcellentParserRPAR        = 6
	ExcellentParserLBRAC       = 7
	ExcellentParserRBRAC       = 8
	ExcellentParserSUB         = 9
	ExcellentParserADD         = 10
	ExcellentParserMUL         = 11
	ExcellentParserDIV         = 12
	ExcellentParserCOMMA       = 13
	ExcellentParserLT          = 14
	ExcellentParserGT          = 15
	ExcellentParserEQ          = 16
	ExcellentParserNE          = 17
	ExcellentParserLTE         = 18
	ExcellentParserGTE         = 19
	ExcellentParserQUOT        = 20
	ExcellentParserEXP         = 21
	ExcellentParserAMP         = 22
	ExcellentParserLITERAL     = 23
	ExcellentParserWhitespace  = 24
	ExcellentParserNAME        = 25
	ExcellentParserERRROR_CHAR = 26
)

// ExcellentParser rules.
const (
	ExcellentParserRULE_parse              = 0
	ExcellentParserRULE_expr               = 1
	ExcellentParserRULE_path               = 2
	ExcellentParserRULE_step               = 3
	ExcellentParserRULE_parameters         = 4
	ExcellentParserRULE_concatenationExpr  = 5
	ExcellentParserRULE_equalityExpr       = 6
	ExcellentParserRULE_comparisonExpr     = 7
	ExcellentParserRULE_additionExpr       = 8
	ExcellentParserRULE_multiplicationExpr = 9
	ExcellentParserRULE_exponentExpr       = 10
	ExcellentParserRULE_unaryExpr          = 11
	ExcellentParserRULE_funcCall           = 12
	ExcellentParserRULE_funcPath           = 13
	ExcellentParserRULE_atom               = 14
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcellentParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcellentParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExcellentParserEOF, 0)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitParse(s)
	}
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExcellentParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExcellentParserRULE_parse)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.Expr()
	}
	{
		p.SetState(31)
		p.Match(ExcellentParserEOF)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcellentParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcellentParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpUnaryContext struct {
	*ExprContext
}

func NewExpUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpUnaryContext {
	var p = new(ExpUnaryContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExpUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpUnaryContext) UnaryExpr() IUnaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *ExpUnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterExpUnary(s)
	}
}

func (s *ExpUnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitExpUnary(s)
	}
}

func (s *ExpUnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitExpUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpConcatenationContext struct {
	*ExprContext
}

func NewExpConcatenationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpConcatenationContext {
	var p = new(ExpConcatenationContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExpConcatenationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpConcatenationContext) ConcatenationExpr() IConcatenationExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConcatenationExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConcatenationExprContext)
}

func (s *ExpConcatenationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterExpConcatenation(s)
	}
}

func (s *ExpConcatenationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitExpConcatenation(s)
	}
}

func (s *ExpConcatenationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitExpConcatenation(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpComparisonContext struct {
	*ExprContext
}

func NewExpComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpComparisonContext {
	var p = new(ExpComparisonContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExpComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpComparisonContext) ComparisonExpr() IComparisonExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonExprContext)
}

func (s *ExpComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterExpComparison(s)
	}
}

func (s *ExpComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitExpComparison(s)
	}
}

func (s *ExpComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitExpComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpMultiplicationContext struct {
	*ExprContext
}

func NewExpMultiplicationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpMultiplicationContext {
	var p = new(ExpMultiplicationContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExpMultiplicationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpMultiplicationContext) MultiplicationExpr() IMultiplicationExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicationExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicationExprContext)
}

func (s *ExpMultiplicationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterExpMultiplication(s)
	}
}

func (s *ExpMultiplicationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitExpMultiplication(s)
	}
}

func (s *ExpMultiplicationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitExpMultiplication(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpAtomContext struct {
	*ExprContext
}

func NewExpAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpAtomContext {
	var p = new(ExpAtomContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExpAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpAtomContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ExpAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterExpAtom(s)
	}
}

func (s *ExpAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitExpAtom(s)
	}
}

func (s *ExpAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitExpAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpEqualityContext struct {
	*ExprContext
}

func NewExpEqualityContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpEqualityContext {
	var p = new(ExpEqualityContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExpEqualityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpEqualityContext) EqualityExpr() IEqualityExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqualityExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqualityExprContext)
}

func (s *ExpEqualityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterExpEquality(s)
	}
}

func (s *ExpEqualityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitExpEquality(s)
	}
}

func (s *ExpEqualityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitExpEquality(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpAdditionContext struct {
	*ExprContext
}

func NewExpAdditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpAdditionContext {
	var p = new(ExpAdditionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExpAdditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpAdditionContext) AdditionExpr() IAdditionExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditionExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditionExprContext)
}

func (s *ExpAdditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterExpAddition(s)
	}
}

func (s *ExpAdditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitExpAddition(s)
	}
}

func (s *ExpAdditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitExpAddition(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpExponentContext struct {
	*ExprContext
}

func NewExpExponentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpExponentContext {
	var p = new(ExpExponentContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExpExponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpExponentContext) ExponentExpr() IExponentExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExponentExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExponentExprContext)
}

func (s *ExpExponentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterExpExponent(s)
	}
}

func (s *ExpExponentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitExpExponent(s)
	}
}

func (s *ExpExponentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitExpExponent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExcellentParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExcellentParserRULE_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Atom()
		}

	case 2:
		localctx = NewExpConcatenationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)
			p.ConcatenationExpr()
		}

	case 3:
		localctx = NewExpEqualityContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(35)
			p.EqualityExpr()
		}

	case 4:
		localctx = NewExpComparisonContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(36)
			p.ComparisonExpr()
		}

	case 5:
		localctx = NewExpAdditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(37)
			p.AdditionExpr()
		}

	case 6:
		localctx = NewExpMultiplicationContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(38)
			p.MultiplicationExpr()
		}

	case 7:
		localctx = NewExpExponentContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(39)
			p.ExponentExpr()
		}

	case 8:
		localctx = NewExpUnaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(40)
			p.UnaryExpr()
		}

	}

	return localctx
}

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcellentParserRULE_path
	return p
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcellentParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) NAME() antlr.TerminalNode {
	return s.GetToken(ExcellentParserNAME, 0)
}

func (s *PathContext) AllStep() []IStepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStepContext)(nil)).Elem())
	var tst = make([]IStepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStepContext)
		}
	}

	return tst
}

func (s *PathContext) Step(i int) IStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStepContext)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitPath(s)
	}
}

func (s *PathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExcellentParser) Path() (localctx IPathContext) {
	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ExcellentParserRULE_path)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(ExcellentParserNAME)
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ExcellentParserPATHSEP || _la == ExcellentParserLBRAC {
		{
			p.SetState(44)
			p.Step()
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStepContext is an interface to support dynamic dispatch.
type IStepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStepContext differentiates from other interfaces.
	IsStepContext()
}

type StepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStepContext() *StepContext {
	var p = new(StepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcellentParserRULE_step
	return p
}

func (*StepContext) IsStepContext() {}

func NewStepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StepContext {
	var p = new(StepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcellentParserRULE_step

	return p
}

func (s *StepContext) GetParser() antlr.Parser { return s.parser }

func (s *StepContext) LBRAC() antlr.TerminalNode {
	return s.GetToken(ExcellentParserLBRAC, 0)
}

func (s *StepContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StepContext) RBRAC() antlr.TerminalNode {
	return s.GetToken(ExcellentParserRBRAC, 0)
}

func (s *StepContext) PATHSEP() antlr.TerminalNode {
	return s.GetToken(ExcellentParserPATHSEP, 0)
}

func (s *StepContext) NAME() antlr.TerminalNode {
	return s.GetToken(ExcellentParserNAME, 0)
}

func (s *StepContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ExcellentParserNUMBER, 0)
}

func (s *StepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterStep(s)
	}
}

func (s *StepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitStep(s)
	}
}

func (s *StepContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitStep(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExcellentParser) Step() (localctx IStepContext) {
	localctx = NewStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ExcellentParserRULE_step)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Match(ExcellentParserLBRAC)
		}
		{
			p.SetState(51)
			p.Expr()
		}
		{
			p.SetState(52)
			p.Match(ExcellentParserRBRAC)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Match(ExcellentParserPATHSEP)
		}
		{
			p.SetState(55)
			p.Match(ExcellentParserNAME)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(56)
			p.Match(ExcellentParserPATHSEP)
		}
		{
			p.SetState(57)
			p.Match(ExcellentParserNUMBER)
		}

	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcellentParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcellentParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) CopyFrom(ctx *ParametersContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionParametersContext struct {
	*ParametersContext
}

func NewFunctionParametersContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionParametersContext {
	var p = new(FunctionParametersContext)

	p.ParametersContext = NewEmptyParametersContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ParametersContext))

	return p
}

func (s *FunctionParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParametersContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *FunctionParametersContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ExcellentParserCOMMA)
}

func (s *FunctionParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ExcellentParserCOMMA, i)
}

func (s *FunctionParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterFunctionParameters(s)
	}
}

func (s *FunctionParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitFunctionParameters(s)
	}
}

func (s *FunctionParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitFunctionParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExcellentParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ExcellentParserRULE_parameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewFunctionParametersContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Expr()
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ExcellentParserCOMMA {
		{
			p.SetState(61)
			p.Match(ExcellentParserCOMMA)
		}
		{
			p.SetState(62)
			p.Expr()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConcatenationExprContext is an interface to support dynamic dispatch.
type IConcatenationExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConcatenationExprContext differentiates from other interfaces.
	IsConcatenationExprContext()
}

type ConcatenationExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcatenationExprContext() *ConcatenationExprContext {
	var p = new(ConcatenationExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcellentParserRULE_concatenationExpr
	return p
}

func (*ConcatenationExprContext) IsConcatenationExprContext() {}

func NewConcatenationExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConcatenationExprContext {
	var p = new(ConcatenationExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcellentParserRULE_concatenationExpr

	return p
}

func (s *ConcatenationExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ConcatenationExprContext) CopyFrom(ctx *ConcatenationExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ConcatenationExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcatenationExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConcatenationContext struct {
	*ConcatenationExprContext
}

func NewConcatenationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConcatenationContext {
	var p = new(ConcatenationContext)

	p.ConcatenationExprContext = NewEmptyConcatenationExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConcatenationExprContext))

	return p
}

func (s *ConcatenationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcatenationContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ConcatenationContext) AMP() antlr.TerminalNode {
	return s.GetToken(ExcellentParserAMP, 0)
}

func (s *ConcatenationContext) ConcatenationExpr() IConcatenationExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConcatenationExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConcatenationExprContext)
}

func (s *ConcatenationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterConcatenation(s)
	}
}

func (s *ConcatenationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitConcatenation(s)
	}
}

func (s *ConcatenationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitConcatenation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExcellentParser) ConcatenationExpr() (localctx IConcatenationExprContext) {
	localctx = NewConcatenationExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ExcellentParserRULE_concatenationExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewConcatenationContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Atom()
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ExcellentParserAMP {
		{
			p.SetState(69)
			p.Match(ExcellentParserAMP)
		}
		{
			p.SetState(70)
			p.ConcatenationExpr()
		}

	}

	return localctx
}

// IEqualityExprContext is an interface to support dynamic dispatch.
type IEqualityExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqualityExprContext differentiates from other interfaces.
	IsEqualityExprContext()
}

type EqualityExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExprContext() *EqualityExprContext {
	var p = new(EqualityExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcellentParserRULE_equalityExpr
	return p
}

func (*EqualityExprContext) IsEqualityExprContext() {}

func NewEqualityExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExprContext {
	var p = new(EqualityExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcellentParserRULE_equalityExpr

	return p
}

func (s *EqualityExprContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExprContext) CopyFrom(ctx *EqualityExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *EqualityExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EqualityContext struct {
	*EqualityExprContext
	op antlr.Token
}

func NewEqualityContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityContext {
	var p = new(EqualityContext)

	p.EqualityExprContext = NewEmptyEqualityExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*EqualityExprContext))

	return p
}

func (s *EqualityContext) GetOp() antlr.Token { return s.op }

func (s *EqualityContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityContext) AllComparisonExpr() []IComparisonExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComparisonExprContext)(nil)).Elem())
	var tst = make([]IComparisonExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComparisonExprContext)
		}
	}

	return tst
}

func (s *EqualityContext) ComparisonExpr(i int) IComparisonExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComparisonExprContext)
}

func (s *EqualityContext) EQ() antlr.TerminalNode {
	return s.GetToken(ExcellentParserEQ, 0)
}

func (s *EqualityContext) NE() antlr.TerminalNode {
	return s.GetToken(ExcellentParserNE, 0)
}

func (s *EqualityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterEquality(s)
	}
}

func (s *EqualityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitEquality(s)
	}
}

func (s *EqualityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitEquality(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExcellentParser) EqualityExpr() (localctx IEqualityExprContext) {
	localctx = NewEqualityExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ExcellentParserRULE_equalityExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewEqualityContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.ComparisonExpr()
	}
	p.SetState(74)

	var _lt = p.GetTokenStream().LT(1)

	localctx.(*EqualityContext).op = _lt

	_la = p.GetTokenStream().LA(1)

	if !(_la == ExcellentParserEQ || _la == ExcellentParserNE) {
		var _ri = p.GetErrorHandler().RecoverInline(p)

		localctx.(*EqualityContext).op = _ri
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(75)
		p.ComparisonExpr()
	}

	return localctx
}

// IComparisonExprContext is an interface to support dynamic dispatch.
type IComparisonExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonExprContext differentiates from other interfaces.
	IsComparisonExprContext()
}

type ComparisonExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonExprContext() *ComparisonExprContext {
	var p = new(ComparisonExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcellentParserRULE_comparisonExpr
	return p
}

func (*ComparisonExprContext) IsComparisonExprContext() {}

func NewComparisonExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonExprContext {
	var p = new(ComparisonExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcellentParserRULE_comparisonExpr

	return p
}

func (s *ComparisonExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonExprContext) CopyFrom(ctx *ComparisonExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ComparisonExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ComparisonContext struct {
	*ComparisonExprContext
	op antlr.Token
}

func NewComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonContext {
	var p = new(ComparisonContext)

	p.ComparisonExprContext = NewEmptyComparisonExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComparisonExprContext))

	return p
}

func (s *ComparisonContext) GetOp() antlr.Token { return s.op }

func (s *ComparisonContext) SetOp(v antlr.Token) { s.op = v }

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) AllAdditionExpr() []IAdditionExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAdditionExprContext)(nil)).Elem())
	var tst = make([]IAdditionExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAdditionExprContext)
		}
	}

	return tst
}

func (s *ComparisonContext) AdditionExpr(i int) IAdditionExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditionExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAdditionExprContext)
}

func (s *ComparisonContext) LT() antlr.TerminalNode {
	return s.GetToken(ExcellentParserLT, 0)
}

func (s *ComparisonContext) GT() antlr.TerminalNode {
	return s.GetToken(ExcellentParserGT, 0)
}

func (s *ComparisonContext) LTE() antlr.TerminalNode {
	return s.GetToken(ExcellentParserLTE, 0)
}

func (s *ComparisonContext) GTE() antlr.TerminalNode {
	return s.GetToken(ExcellentParserGTE, 0)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (s *ComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExcellentParser) ComparisonExpr() (localctx IComparisonExprContext) {
	localctx = NewComparisonExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ExcellentParserRULE_comparisonExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewComparisonContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.AdditionExpr()
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ExcellentParserLT)|(1<<ExcellentParserGT)|(1<<ExcellentParserLTE)|(1<<ExcellentParserGTE))) != 0 {
		p.SetState(78)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ComparisonContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ExcellentParserLT)|(1<<ExcellentParserGT)|(1<<ExcellentParserLTE)|(1<<ExcellentParserGTE))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ComparisonContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(79)
			p.AdditionExpr()
		}

	}

	return localctx
}

// IAdditionExprContext is an interface to support dynamic dispatch.
type IAdditionExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditionExprContext differentiates from other interfaces.
	IsAdditionExprContext()
}

type AdditionExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditionExprContext() *AdditionExprContext {
	var p = new(AdditionExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcellentParserRULE_additionExpr
	return p
}

func (*AdditionExprContext) IsAdditionExprContext() {}

func NewAdditionExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditionExprContext {
	var p = new(AdditionExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcellentParserRULE_additionExpr

	return p
}

func (s *AdditionExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditionExprContext) CopyFrom(ctx *AdditionExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AdditionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AdditionContext struct {
	*AdditionExprContext
	op antlr.Token
}

func NewAdditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditionContext {
	var p = new(AdditionContext)

	p.AdditionExprContext = NewEmptyAdditionExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AdditionExprContext))

	return p
}

func (s *AdditionContext) GetOp() antlr.Token { return s.op }

func (s *AdditionContext) SetOp(v antlr.Token) { s.op = v }

func (s *AdditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionContext) AllMultiplicationExpr() []IMultiplicationExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiplicationExprContext)(nil)).Elem())
	var tst = make([]IMultiplicationExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiplicationExprContext)
		}
	}

	return tst
}

func (s *AdditionContext) MultiplicationExpr(i int) IMultiplicationExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicationExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiplicationExprContext)
}

func (s *AdditionContext) AllADD() []antlr.TerminalNode {
	return s.GetTokens(ExcellentParserADD)
}

func (s *AdditionContext) ADD(i int) antlr.TerminalNode {
	return s.GetToken(ExcellentParserADD, i)
}

func (s *AdditionContext) AllSUB() []antlr.TerminalNode {
	return s.GetTokens(ExcellentParserSUB)
}

func (s *AdditionContext) SUB(i int) antlr.TerminalNode {
	return s.GetToken(ExcellentParserSUB, i)
}

func (s *AdditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterAddition(s)
	}
}

func (s *AdditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitAddition(s)
	}
}

func (s *AdditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitAddition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExcellentParser) AdditionExpr() (localctx IAdditionExprContext) {
	localctx = NewAdditionExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ExcellentParserRULE_additionExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewAdditionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.MultiplicationExpr()
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ExcellentParserSUB || _la == ExcellentParserADD {
		p.SetState(83)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*AdditionContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == ExcellentParserSUB || _la == ExcellentParserADD) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*AdditionContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(84)
			p.MultiplicationExpr()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMultiplicationExprContext is an interface to support dynamic dispatch.
type IMultiplicationExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicationExprContext differentiates from other interfaces.
	IsMultiplicationExprContext()
}

type MultiplicationExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicationExprContext() *MultiplicationExprContext {
	var p = new(MultiplicationExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcellentParserRULE_multiplicationExpr
	return p
}

func (*MultiplicationExprContext) IsMultiplicationExprContext() {}

func NewMultiplicationExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicationExprContext {
	var p = new(MultiplicationExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcellentParserRULE_multiplicationExpr

	return p
}

func (s *MultiplicationExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicationExprContext) CopyFrom(ctx *MultiplicationExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MultiplicationExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MultiplicationContext struct {
	*MultiplicationExprContext
	op antlr.Token
}

func NewMultiplicationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicationContext {
	var p = new(MultiplicationContext)

	p.MultiplicationExprContext = NewEmptyMultiplicationExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MultiplicationExprContext))

	return p
}

func (s *MultiplicationContext) GetOp() antlr.Token { return s.op }

func (s *MultiplicationContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultiplicationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationContext) AllExponentExpr() []IExponentExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExponentExprContext)(nil)).Elem())
	var tst = make([]IExponentExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExponentExprContext)
		}
	}

	return tst
}

func (s *MultiplicationContext) ExponentExpr(i int) IExponentExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExponentExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExponentExprContext)
}

func (s *MultiplicationContext) AllMUL() []antlr.TerminalNode {
	return s.GetTokens(ExcellentParserMUL)
}

func (s *MultiplicationContext) MUL(i int) antlr.TerminalNode {
	return s.GetToken(ExcellentParserMUL, i)
}

func (s *MultiplicationContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(ExcellentParserDIV)
}

func (s *MultiplicationContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(ExcellentParserDIV, i)
}

func (s *MultiplicationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterMultiplication(s)
	}
}

func (s *MultiplicationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitMultiplication(s)
	}
}

func (s *MultiplicationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitMultiplication(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExcellentParser) MultiplicationExpr() (localctx IMultiplicationExprContext) {
	localctx = NewMultiplicationExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ExcellentParserRULE_multiplicationExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewMultiplicationContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.ExponentExpr()
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ExcellentParserMUL || _la == ExcellentParserDIV {
		p.SetState(91)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*MultiplicationContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == ExcellentParserMUL || _la == ExcellentParserDIV) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*MultiplicationContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(92)
			p.ExponentExpr()
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExponentExprContext is an interface to support dynamic dispatch.
type IExponentExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExponentExprContext differentiates from other interfaces.
	IsExponentExprContext()
}

type ExponentExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExponentExprContext() *ExponentExprContext {
	var p = new(ExponentExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcellentParserRULE_exponentExpr
	return p
}

func (*ExponentExprContext) IsExponentExprContext() {}

func NewExponentExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExponentExprContext {
	var p = new(ExponentExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcellentParserRULE_exponentExpr

	return p
}

func (s *ExponentExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExponentExprContext) CopyFrom(ctx *ExponentExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExponentExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExponentExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExponentContext struct {
	*ExponentExprContext
}

func NewExponentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExponentContext {
	var p = new(ExponentContext)

	p.ExponentExprContext = NewEmptyExponentExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExponentExprContext))

	return p
}

func (s *ExponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExponentContext) UnaryExpr() IUnaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *ExponentContext) EXP() antlr.TerminalNode {
	return s.GetToken(ExcellentParserEXP, 0)
}

func (s *ExponentContext) ExponentExpr() IExponentExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExponentExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExponentExprContext)
}

func (s *ExponentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterExponent(s)
	}
}

func (s *ExponentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitExponent(s)
	}
}

func (s *ExponentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitExponent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExcellentParser) ExponentExpr() (localctx IExponentExprContext) {
	localctx = NewExponentExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ExcellentParserRULE_exponentExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewExponentContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.UnaryExpr()
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ExcellentParserEXP {
		{
			p.SetState(99)
			p.Match(ExcellentParserEXP)
		}
		{
			p.SetState(100)
			p.ExponentExpr()
		}

	}

	return localctx
}

// IUnaryExprContext is an interface to support dynamic dispatch.
type IUnaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryExprContext differentiates from other interfaces.
	IsUnaryExprContext()
}

type UnaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExprContext() *UnaryExprContext {
	var p = new(UnaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcellentParserRULE_unaryExpr
	return p
}

func (*UnaryExprContext) IsUnaryExprContext() {}

func NewUnaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExprContext {
	var p = new(UnaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcellentParserRULE_unaryExpr

	return p
}

func (s *UnaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExprContext) CopyFrom(ctx *UnaryExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NegationContext struct {
	*UnaryExprContext
}

func NewNegationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegationContext {
	var p = new(NegationContext)

	p.UnaryExprContext = NewEmptyUnaryExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnaryExprContext))

	return p
}

func (s *NegationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegationContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *NegationContext) SUB() antlr.TerminalNode {
	return s.GetToken(ExcellentParserSUB, 0)
}

func (s *NegationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterNegation(s)
	}
}

func (s *NegationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitNegation(s)
	}
}

func (s *NegationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitNegation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExcellentParser) UnaryExpr() (localctx IUnaryExprContext) {
	localctx = NewUnaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ExcellentParserRULE_unaryExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewNegationContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ExcellentParserSUB {
		{
			p.SetState(103)
			p.Match(ExcellentParserSUB)
		}

	}
	{
		p.SetState(106)
		p.Atom()
	}

	return localctx
}

// IFuncCallContext is an interface to support dynamic dispatch.
type IFuncCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncCallContext differentiates from other interfaces.
	IsFuncCallContext()
}

type FuncCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallContext() *FuncCallContext {
	var p = new(FuncCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcellentParserRULE_funcCall
	return p
}

func (*FuncCallContext) IsFuncCallContext() {}

func NewFuncCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallContext {
	var p = new(FuncCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcellentParserRULE_funcCall

	return p
}

func (s *FuncCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallContext) CopyFrom(ctx *FuncCallContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionCallContext struct {
	*FuncCallContext
	function antlr.Token
}

func NewFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.FuncCallContext = NewEmptyFuncCallContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FuncCallContext))

	return p
}

func (s *FunctionCallContext) GetFunction() antlr.Token { return s.function }

func (s *FunctionCallContext) SetFunction(v antlr.Token) { s.function = v }

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) LPAR() antlr.TerminalNode {
	return s.GetToken(ExcellentParserLPAR, 0)
}

func (s *FunctionCallContext) RPAR() antlr.TerminalNode {
	return s.GetToken(ExcellentParserRPAR, 0)
}

func (s *FunctionCallContext) NAME() antlr.TerminalNode {
	return s.GetToken(ExcellentParserNAME, 0)
}

func (s *FunctionCallContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExcellentParser) FuncCall() (localctx IFuncCallContext) {
	localctx = NewFuncCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ExcellentParserRULE_funcCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewFunctionCallContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)

		var _m = p.Match(ExcellentParserNAME)

		localctx.(*FunctionCallContext).function = _m
	}
	{
		p.SetState(109)
		p.Match(ExcellentParserLPAR)
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ExcellentParserNUMBER)|(1<<ExcellentParserTRUE)|(1<<ExcellentParserFALSE)|(1<<ExcellentParserLPAR)|(1<<ExcellentParserSUB)|(1<<ExcellentParserLITERAL)|(1<<ExcellentParserNAME))) != 0 {
		{
			p.SetState(110)
			p.Parameters()
		}

	}
	{
		p.SetState(113)
		p.Match(ExcellentParserRPAR)
	}

	return localctx
}

// IFuncPathContext is an interface to support dynamic dispatch.
type IFuncPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncPathContext differentiates from other interfaces.
	IsFuncPathContext()
}

type FuncPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncPathContext() *FuncPathContext {
	var p = new(FuncPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcellentParserRULE_funcPath
	return p
}

func (*FuncPathContext) IsFuncPathContext() {}

func NewFuncPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncPathContext {
	var p = new(FuncPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcellentParserRULE_funcPath

	return p
}

func (s *FuncPathContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncPathContext) CopyFrom(ctx *FuncPathContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FuncPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionPathContext struct {
	*FuncPathContext
	function IFuncCallContext
}

func NewFunctionPathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionPathContext {
	var p = new(FunctionPathContext)

	p.FuncPathContext = NewEmptyFuncPathContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FuncPathContext))

	return p
}

func (s *FunctionPathContext) GetFunction() IFuncCallContext { return s.function }

func (s *FunctionPathContext) SetFunction(v IFuncCallContext) { s.function = v }

func (s *FunctionPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionPathContext) FuncCall() IFuncCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *FunctionPathContext) AllStep() []IStepContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStepContext)(nil)).Elem())
	var tst = make([]IStepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStepContext)
		}
	}

	return tst
}

func (s *FunctionPathContext) Step(i int) IStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStepContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStepContext)
}

func (s *FunctionPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterFunctionPath(s)
	}
}

func (s *FunctionPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitFunctionPath(s)
	}
}

func (s *FunctionPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitFunctionPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExcellentParser) FuncPath() (localctx IFuncPathContext) {
	localctx = NewFuncPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ExcellentParserRULE_funcPath)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewFunctionPathContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)

		var _x = p.FuncCall()

		localctx.(*FunctionPathContext).function = _x
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ExcellentParserPATHSEP || _la == ExcellentParserLBRAC {
		{
			p.SetState(116)
			p.Step()
		}

		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExcellentParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExcellentParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) CopyFrom(ctx *AtomContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DecimalLiteralContext struct {
	*AtomContext
}

func NewDecimalLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *DecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalLiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ExcellentParserNUMBER, 0)
}

func (s *DecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitDecimalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenthesesContext struct {
	*AtomContext
}

func NewParenthesesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesesContext {
	var p = new(ParenthesesContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *ParenthesesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesesContext) LPAR() antlr.TerminalNode {
	return s.GetToken(ExcellentParserLPAR, 0)
}

func (s *ParenthesesContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenthesesContext) RPAR() antlr.TerminalNode {
	return s.GetToken(ExcellentParserRPAR, 0)
}

func (s *ParenthesesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterParentheses(s)
	}
}

func (s *ParenthesesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitParentheses(s)
	}
}

func (s *ParenthesesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitParentheses(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringLiteralContext struct {
	*AtomContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(ExcellentParserLITERAL, 0)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type TrueContext struct {
	*AtomContext
}

func NewTrueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueContext {
	var p = new(TrueContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *TrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(ExcellentParserTRUE, 0)
}

func (s *TrueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterTrue(s)
	}
}

func (s *TrueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitTrue(s)
	}
}

func (s *TrueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitTrue(s)

	default:
		return t.VisitChildren(s)
	}
}

type FalseContext struct {
	*AtomContext
}

func NewFalseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseContext {
	var p = new(FalseContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *FalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseContext) FALSE() antlr.TerminalNode {
	return s.GetToken(ExcellentParserFALSE, 0)
}

func (s *FalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterFalse(s)
	}
}

func (s *FalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitFalse(s)
	}
}

func (s *FalseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitFalse(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContextReferenceContext struct {
	*AtomContext
}

func NewContextReferenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContextReferenceContext {
	var p = new(ContextReferenceContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *ContextReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContextReferenceContext) Path() IPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *ContextReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterContextReference(s)
	}
}

func (s *ContextReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitContextReference(s)
	}
}

func (s *ContextReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitContextReference(s)

	default:
		return t.VisitChildren(s)
	}
}

type AtomFuncCallContext struct {
	*AtomContext
}

func NewAtomFuncCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomFuncCallContext {
	var p = new(AtomFuncCallContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *AtomFuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomFuncCallContext) FuncCall() IFuncCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *AtomFuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterAtomFuncCall(s)
	}
}

func (s *AtomFuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitAtomFuncCall(s)
	}
}

func (s *AtomFuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitAtomFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type AtomFuncPathContext struct {
	*AtomContext
}

func NewAtomFuncPathContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomFuncPathContext {
	var p = new(AtomFuncPathContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *AtomFuncPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomFuncPathContext) FuncPath() IFuncPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncPathContext)
}

func (s *AtomFuncPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.EnterAtomFuncPath(s)
	}
}

func (s *AtomFuncPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExcellentListener); ok {
		listenerT.ExitAtomFuncPath(s)
	}
}

func (s *AtomFuncPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExcellentVisitor:
		return t.VisitAtomFuncPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ExcellentParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ExcellentParserRULE_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewContextReferenceContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Path()
		}

	case 2:
		localctx = NewAtomFuncCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.FuncCall()
		}

	case 3:
		localctx = NewAtomFuncPathContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.FuncPath()
		}

	case 4:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(125)
			p.Match(ExcellentParserLITERAL)
		}

	case 5:
		localctx = NewDecimalLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(126)
			p.Match(ExcellentParserNUMBER)
		}

	case 6:
		localctx = NewParenthesesContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(127)
			p.Match(ExcellentParserLPAR)
		}
		{
			p.SetState(128)
			p.Expr()
		}
		{
			p.SetState(129)
			p.Match(ExcellentParserRPAR)
		}

	case 7:
		localctx = NewTrueContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(131)
			p.Match(ExcellentParserTRUE)
		}

	case 8:
		localctx = NewFalseContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(132)
			p.Match(ExcellentParserFALSE)
		}

	}

	return localctx
}
