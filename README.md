
ANTLR Go Performance
=====================

This repo demonstrates either a bug in the performance of ANTLR parsing of grammars or a total misunderstanding of how to build a proper grammar. (quite likely the second)

The long expression in /nyaruka/goantlr/goantlr_test.go: `(DATEDIF(DATEVALUE("01-01-1970"), date.now, "D") * 24 * 60 * 60) + ((((HOUR(date.now)+7) * 60) + MINUTE(date.now)) * 60))` takes over 28 seconds to parse using the Excellent.g4 grammar.

This was build using Antlr 4.7 with the following steps:

```
% antlr4 -Dlanguage=Go src/github.com/nyaruka/goantlr/gen/Excellent.g4 -visitor -package gen
% go test github.com/nyaruka/goantlr -bench=.
       16.952µs LEXING "hello world"
     1.195137ms PARSING "hello world"
         2.32µs LEXING 5 + 10
      874.645µs PARSING 5 + 10
        2.075µs LEXING FIRST_WORD(WORD_SLICE(contact.blerg, 2, 4))
   846.239202ms PARSING FIRST_WORD(WORD_SLICE(contact.blerg, 2, 4))
      271.049µs LEXING ((DATEDIF(DATEVALUE("01-01-1970"), date.now, "D") * 24 * 60 * 60) + (((HOUR(date.now)+7) * 60) + MINUTE(date.now)) * 60)
  39.684332943s PARSING ((DATEDIF(DATEVALUE("01-01-1970"), date.now, "D") * 24 * 60 * 60) + (((HOUR(date.now)+7) * 60) + MINUTE(date.now)) * 60)
```

Note on Python performance is still slow (pointing to a Grammar issue) but still 20 times faster than Go:

```
0.000059s LEXING "hello world"
0.014714s PARSING "hello world"
0.000049s LEXING 5 + 10
0.015481s PARSING 5 + 10
0.000057s LEXING FIRST_WORD(WORD_SLICE(contact.blerg, 2, 4))
0.644930s PARSING FIRST_WORD(WORD_SLICE(contact.blerg, 2, 4))
0.000092s LEXING ((DATEDIF(DATEVALUE("01-01-1970"), date.now, "D") * 24 * 60 * 60) + (((HOUR(date.now)+7) * 60) + MINUTE(date.now)) * 60)
1.941638s PARSING ((DATEDIF(DATEVALUE("01-01-1970"), date.now, "D") * 24 * 60 * 60) + (((HOUR(date.now)+7) * 60) + MINUTE(date.now)) * 60)
```