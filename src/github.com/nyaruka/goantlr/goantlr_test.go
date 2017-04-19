package goantlr

import (
	"fmt"
	"testing"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/nyaruka/goantlr/gen"
)

func BenchmarkAntlr(b *testing.B) {

	parsingTests := []string{
		"\"hello world\"",
		"5 + 10",
		"FIRST_WORD(WORD_SLICE(contact.blerg, 2, 4))",
		"(DATEDIF(DATEVALUE(\"01-01-1970\"), date.now, \"D\") * 24 * 60 * 60) + ((((HOUR(date.now)+7) * 60) + MINUTE(date.now)) * 60))",
	}

	for _, test := range parsingTests {
		start := time.Now()
		input := antlr.NewInputStream(test)
		lexer := gen.NewExcellentLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		fmt.Printf("%15s LEXING %s\n", time.Since(start), test)
		start = time.Now()

		p := gen.NewExcellentParser(stream)
		p.Parse()

		fmt.Printf("%15s PARSING %s\n", time.Since(start), test)
	}
}
