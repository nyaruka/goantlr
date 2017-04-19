
ANTLR Go Performance
=====================

This repo demonstrates either a bug in the performance of ANTLR parsing of grammars or a total misunderstanding of how to build a proper grammar. (quite likely the second)

The long expression in /nyaruka/goantlr/goantlr_test.go: `(DATEDIF(DATEVALUE("01-01-1970"), date.now, "D") * 24 * 60 * 60) + ((((HOUR(date.now)+7) * 60) + MINUTE(date.now)) * 60))` takes over 28 seconds to parse using the Excellent.g4 grammar.

This was build using Antlr 4.7 with the following steps:

```
% antlr4 -Dlanguage=Go src/github.com/nyaruka/goantlr/gen/Excellent.g4 -visitor -package gen
% go test github.com/nyaruka/goantlr -bench=.
       19.348µs LEXING "hello world"
     1.211254ms PARSING "hello world"
        2.217µs LEXING 5 + 10
      867.589µs PARSING 5 + 10
        2.181µs LEXING FIRST_WORD(WORD_SLICE(contact.blerg, 2, 4))
      843.282ms PARSING FIRST_WORD(WORD_SLICE(contact.blerg, 2, 4))
        3.628µs LEXING (DATEDIF(DATEVALUE("01-01-1970"), date.now, "D") * 24 * 60 * 60) + ((((HOUR(date.now)+7) * 60) + MINUTE(date.now)) * 60))
line 1:120 no viable alternative at input '(DATEDIF(DATEVALUE("01-01-1970"),date.now,"D")*24*60*60)+((((HOUR(date.now)+7)*60)+MINUTE(date.now))*60))'
  26.604543359s PARSING (DATEDIF(DATEVALUE("01-01-1970"), date.now, "D") * 24 * 60 * 60) + ((((HOUR(date.now)+7) * 60) + MINUTE(date.now)) * 60))
```