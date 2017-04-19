// Generated from src/github.com/nyaruka/goantlr/gen/Excellent.g4 by ANTLR 4.7.

package gen

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 28, 156,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 3, 2, 3, 2, 3, 2, 5, 2, 65, 10, 2, 5, 2,
	67, 10, 2, 3, 3, 6, 3, 70, 10, 3, 13, 3, 14, 3, 71, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3,
	22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 7, 25, 128, 10, 25, 12, 25,
	14, 25, 131, 11, 25, 3, 25, 3, 25, 3, 26, 6, 26, 136, 10, 26, 13, 26, 14,
	26, 137, 3, 26, 3, 26, 3, 27, 3, 27, 7, 27, 144, 10, 27, 12, 27, 14, 27,
	147, 11, 27, 3, 28, 3, 28, 3, 29, 3, 29, 5, 29, 153, 10, 29, 3, 30, 3,
	30, 2, 2, 31, 3, 3, 5, 2, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
	21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
	39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 2,
	57, 2, 59, 28, 3, 2, 14, 4, 2, 86, 86, 118, 118, 4, 2, 84, 84, 116, 116,
	4, 2, 87, 87, 119, 119, 4, 2, 71, 71, 103, 103, 4, 2, 72, 72, 104, 104,
	4, 2, 67, 67, 99, 99, 4, 2, 78, 78, 110, 110, 4, 2, 85, 85, 117, 117, 3,
	2, 36, 36, 5, 2, 11, 12, 15, 15, 34, 34, 16, 2, 67, 92, 97, 97, 99, 124,
	194, 216, 218, 248, 250, 769, 882, 895, 897, 8193, 8206, 8207, 8306, 8593,
	11266, 12273, 12291, 55297, 63746, 64977, 65010, 65535, 6, 2, 50, 59, 185,
	185, 770, 881, 8257, 8258, 2, 159, 2, 3, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2,
	2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2,
	2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2,
	2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3,
	2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39,
	3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2,
	47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2,
	2, 59, 3, 2, 2, 2, 3, 61, 3, 2, 2, 2, 5, 69, 3, 2, 2, 2, 7, 73, 3, 2, 2,
	2, 9, 78, 3, 2, 2, 2, 11, 84, 3, 2, 2, 2, 13, 86, 3, 2, 2, 2, 15, 88, 3,
	2, 2, 2, 17, 90, 3, 2, 2, 2, 19, 92, 3, 2, 2, 2, 21, 94, 3, 2, 2, 2, 23,
	96, 3, 2, 2, 2, 25, 98, 3, 2, 2, 2, 27, 100, 3, 2, 2, 2, 29, 102, 3, 2,
	2, 2, 31, 104, 3, 2, 2, 2, 33, 106, 3, 2, 2, 2, 35, 108, 3, 2, 2, 2, 37,
	110, 3, 2, 2, 2, 39, 113, 3, 2, 2, 2, 41, 116, 3, 2, 2, 2, 43, 119, 3,
	2, 2, 2, 45, 121, 3, 2, 2, 2, 47, 123, 3, 2, 2, 2, 49, 125, 3, 2, 2, 2,
	51, 135, 3, 2, 2, 2, 53, 141, 3, 2, 2, 2, 55, 148, 3, 2, 2, 2, 57, 152,
	3, 2, 2, 2, 59, 154, 3, 2, 2, 2, 61, 66, 5, 5, 3, 2, 62, 64, 7, 48, 2,
	2, 63, 65, 5, 5, 3, 2, 64, 63, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 67,
	3, 2, 2, 2, 66, 62, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 4, 3, 2, 2, 2,
	68, 70, 4, 50, 59, 2, 69, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 69, 3,
	2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 6, 3, 2, 2, 2, 73, 74, 9, 2, 2, 2, 74,
	75, 9, 3, 2, 2, 75, 76, 9, 4, 2, 2, 76, 77, 9, 5, 2, 2, 77, 8, 3, 2, 2,
	2, 78, 79, 9, 6, 2, 2, 79, 80, 9, 7, 2, 2, 80, 81, 9, 8, 2, 2, 81, 82,
	9, 9, 2, 2, 82, 83, 9, 5, 2, 2, 83, 10, 3, 2, 2, 2, 84, 85, 7, 48, 2, 2,
	85, 12, 3, 2, 2, 2, 86, 87, 7, 42, 2, 2, 87, 14, 3, 2, 2, 2, 88, 89, 7,
	43, 2, 2, 89, 16, 3, 2, 2, 2, 90, 91, 7, 93, 2, 2, 91, 18, 3, 2, 2, 2,
	92, 93, 7, 95, 2, 2, 93, 20, 3, 2, 2, 2, 94, 95, 7, 47, 2, 2, 95, 22, 3,
	2, 2, 2, 96, 97, 7, 45, 2, 2, 97, 24, 3, 2, 2, 2, 98, 99, 7, 44, 2, 2,
	99, 26, 3, 2, 2, 2, 100, 101, 7, 49, 2, 2, 101, 28, 3, 2, 2, 2, 102, 103,
	7, 46, 2, 2, 103, 30, 3, 2, 2, 2, 104, 105, 7, 62, 2, 2, 105, 32, 3, 2,
	2, 2, 106, 107, 7, 64, 2, 2, 107, 34, 3, 2, 2, 2, 108, 109, 7, 63, 2, 2,
	109, 36, 3, 2, 2, 2, 110, 111, 7, 35, 2, 2, 111, 112, 7, 63, 2, 2, 112,
	38, 3, 2, 2, 2, 113, 114, 7, 62, 2, 2, 114, 115, 7, 63, 2, 2, 115, 40,
	3, 2, 2, 2, 116, 117, 7, 64, 2, 2, 117, 118, 7, 63, 2, 2, 118, 42, 3, 2,
	2, 2, 119, 120, 7, 36, 2, 2, 120, 44, 3, 2, 2, 2, 121, 122, 7, 96, 2, 2,
	122, 46, 3, 2, 2, 2, 123, 124, 7, 40, 2, 2, 124, 48, 3, 2, 2, 2, 125, 129,
	7, 36, 2, 2, 126, 128, 10, 10, 2, 2, 127, 126, 3, 2, 2, 2, 128, 131, 3,
	2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 132, 3, 2, 2,
	2, 131, 129, 3, 2, 2, 2, 132, 133, 7, 36, 2, 2, 133, 50, 3, 2, 2, 2, 134,
	136, 9, 11, 2, 2, 135, 134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 135,
	3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 140, 8, 26,
	2, 2, 140, 52, 3, 2, 2, 2, 141, 145, 5, 55, 28, 2, 142, 144, 5, 57, 29,
	2, 143, 142, 3, 2, 2, 2, 144, 147, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145,
	146, 3, 2, 2, 2, 146, 54, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 148, 149, 9,
	12, 2, 2, 149, 56, 3, 2, 2, 2, 150, 153, 5, 55, 28, 2, 151, 153, 9, 13,
	2, 2, 152, 150, 3, 2, 2, 2, 152, 151, 3, 2, 2, 2, 153, 58, 3, 2, 2, 2,
	154, 155, 11, 2, 2, 2, 155, 60, 3, 2, 2, 2, 10, 2, 64, 66, 71, 129, 137,
	145, 152, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "'.'", "'('", "')'", "'['", "']'", "'-'", "'+'", "'*'",
	"'/'", "','", "'<'", "'>'", "'='", "'!='", "'<='", "'>='", "'\"'", "'^'",
	"'&'",
}

var lexerSymbolicNames = []string{
	"", "NUMBER", "TRUE", "FALSE", "PATHSEP", "LPAR", "RPAR", "LBRAC", "RBRAC",
	"SUB", "ADD", "MUL", "DIV", "COMMA", "LT", "GT", "EQ", "NE", "LTE", "GTE",
	"QUOT", "EXP", "AMP", "LITERAL", "Whitespace", "NAME", "ERRROR_CHAR",
}

var lexerRuleNames = []string{
	"NUMBER", "DIGITS", "TRUE", "FALSE", "PATHSEP", "LPAR", "RPAR", "LBRAC",
	"RBRAC", "SUB", "ADD", "MUL", "DIV", "COMMA", "LT", "GT", "EQ", "NE", "LTE",
	"GTE", "QUOT", "EXP", "AMP", "LITERAL", "Whitespace", "NAME", "NAME_START_CHARS",
	"NAME_CHARS", "ERRROR_CHAR",
}

type ExcellentLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewExcellentLexer(input antlr.CharStream) *ExcellentLexer {

	l := new(ExcellentLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Excellent.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExcellentLexer tokens.
const (
	ExcellentLexerNUMBER      = 1
	ExcellentLexerTRUE        = 2
	ExcellentLexerFALSE       = 3
	ExcellentLexerPATHSEP     = 4
	ExcellentLexerLPAR        = 5
	ExcellentLexerRPAR        = 6
	ExcellentLexerLBRAC       = 7
	ExcellentLexerRBRAC       = 8
	ExcellentLexerSUB         = 9
	ExcellentLexerADD         = 10
	ExcellentLexerMUL         = 11
	ExcellentLexerDIV         = 12
	ExcellentLexerCOMMA       = 13
	ExcellentLexerLT          = 14
	ExcellentLexerGT          = 15
	ExcellentLexerEQ          = 16
	ExcellentLexerNE          = 17
	ExcellentLexerLTE         = 18
	ExcellentLexerGTE         = 19
	ExcellentLexerQUOT        = 20
	ExcellentLexerEXP         = 21
	ExcellentLexerAMP         = 22
	ExcellentLexerLITERAL     = 23
	ExcellentLexerWhitespace  = 24
	ExcellentLexerNAME        = 25
	ExcellentLexerERRROR_CHAR = 26
)
