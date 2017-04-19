grammar Excellent;

parse
  : expr EOF
  ;

expr
  :  atom                                                    # expAtom
  |  concatenationExpr                                       # expConcatenation
  |  equalityExpr                                            # expEquality
  |  comparisonExpr                                          # expComparison
  |  additionExpr                                            # expAddition
  |  multiplicationExpr                                      # expMultiplication
  |  exponentExpr                                            # expExponent
  |  unaryExpr                                               # expUnary
  ;

path
  :  NAME (step)*
  ;

step
  : LBRAC expr RBRAC
  | PATHSEP NAME
  | PATHSEP NUMBER
  ;

parameters
  : expr (COMMA expr)*                                       # functionParameters
  ;

concatenationExpr
  : atom (AMP concatenationExpr)?                            # concatenation
  ;

equalityExpr
  :  comparisonExpr op=(EQ|NE) comparisonExpr                # equality
  ;

comparisonExpr
  :  additionExpr (op=(LT|GT|LTE|GTE) additionExpr)?         # comparison
  ;

additionExpr
  :  multiplicationExpr (op=(ADD|SUB) multiplicationExpr)*   # addition
  ;

multiplicationExpr
  :  exponentExpr (op=(MUL|DIV) exponentExpr)*               # multiplication
  ;

exponentExpr
  :  unaryExpr (EXP exponentExpr)?                           # exponent
  ;

unaryExpr
  : SUB? atom                                                # negation
  ;

funcCall
  : function=NAME LPAR parameters? RPAR                      # functionCall
  ;

funcPath
  : function=funcCall (step)*                                # functionPath
  ;

atom
  :  path                                                    # contextReference
  |  funcCall                                                # atomFuncCall
  |  funcPath                                                # atomFuncPath
  |  LITERAL                                                 # stringLiteral
  |  NUMBER                                                  # decimalLiteral
  |  LPAR expr RPAR                                          # parentheses
  |  TRUE                                                    # true
  |  FALSE                                                   # false
  ;

NUMBER
  :  DIGITS ('.' DIGITS?)?
  ;

fragment
DIGITS
  :  ('0'..'9')+
  ;

TRUE
  : [Tt][Rr][Uu][Ee]
  ;

FALSE
  : [Ff][Aa][Ll][Ss][Ee]
  ;

PATHSEP
       :'.';
LPAR
       :'(';
RPAR
       :')';
LBRAC
       :'[';
RBRAC
       :']';
SUB
       :'-';
ADD
       :'+';
MUL
       :'*';
DIV
       :'/';
COMMA
       :',';
LT
       :'<';
GT
       :'>';
EQ
       :'=';
NE
       :'!=';
LTE
       :'<=';
GTE
       :'>=';
QUOT
       :'"';
EXP
       : '^';
AMP
       : '&';

LITERAL
  :  '"' ~'"'* '"'
  ;

Whitespace
  :  (' '|'\t'|'\n'|'\r')+ ->skip
  ;

NAME
  :  NAME_START_CHARS NAME_CHARS*
  ;

fragment
NAME_START_CHARS
  :  'A'..'Z'
  |   '_'
  |  'a'..'z'
  |  '\u00C0'..'\u00D6'
  |  '\u00D8'..'\u00F6'
  |  '\u00F8'..'\u02FF'
  |  '\u0370'..'\u037D'
  |  '\u037F'..'\u1FFF'
  |  '\u200C'..'\u200D'
  |  '\u2070'..'\u218F'
  |  '\u2C00'..'\u2FEF'
  |  '\u3001'..'\uD7FF'
  |  '\uF900'..'\uFDCF'
  |  '\uFDF0'..'\uFFFD'
  ;

fragment
NAME_CHARS
  :  NAME_START_CHARS
  | '0'..'9'
  |  '\u00B7' | '\u0300'..'\u036F'
  |  '\u203F'..'\u2040'
  ;

ERRROR_CHAR
  : .
  ;