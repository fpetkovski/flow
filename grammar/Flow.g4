grammar Flow;

// Parser rules

query
    : letExpression EOF
    | pipeline EOF
    ;

// Let expression with bindings
letExpression
    : LET letBinding (',' letBinding)* IN binaryExpression
    ;

// Single let binding
letBinding
    : IDENTIFIER MATCH_EQ pipeline
    ;

binaryExpression
    : primaryExpression (binaryOperator primaryExpression)*
    ;

binaryOperator
    : OP_ADD | OP_SUB | OP_MUL | OP_DIV
    | OP_GT | OP_LT | OP_GTE | OP_LTE | OP_EQ | MATCH_NEQ
    ;

primaryExpression
    : IDENTIFIER
    | NUMBER
    | '(' binaryExpression ')'
    ;

pipeline
    : selector ('|' pipelineStep)*
    ;

pipelineStep
    : aggregation
    | aligner
    | function
    ;

selector
    : IDENTIFIER? ('{' (labelMatcher (',' labelMatcher)*)? '}')?
    ;

labelMatcher
    : IDENTIFIER (MATCH_EQ | MATCH_NEQ | MATCH_RE | MATCH_NRE) STRING
    ;

// Match operators
MATCH_EQ  : '=';
MATCH_NEQ : '!=';
MATCH_RE  : '=~';
MATCH_NRE : '!~';

OP_ADD : '+';
OP_SUB : '-';
OP_MUL : '*';
OP_DIV : '/';

// Comparison operators
OP_GT  : '>';
OP_LT  : '<';
OP_GTE : '>=';
OP_LTE : '<=';
OP_EQ  : '==';

aggregation
    : AGGREGATION_OP (BY|WITHOUT) ('(' (IDENTIFIER (',' IDENTIFIER)*)? ')')?
    ;

BY         : 'by';
WITHOUT    : 'without';

AGGREGATION_OP
    : 'sum'
    | 'min'
    | 'max'
    | 'avg'
    | 'stddev'
    | 'stdvar'
    | 'count'
    | 'count_values'
    | 'bottomk'
    | 'topk'
    | 'quantile'
    ;

aligner
    : IDENTIFIER DURATION
    ;

function
    : IDENTIFIER
    ;

DURATION
    : NUMBER DURATION_UNIT
    ;

// Lexer rules

// Keywords
LET        : 'let';
IN         : 'in';

// Boolean operators
AND        : 'and';
OR         : 'or';
UNLESS     : 'unless';

// Numbers
NUMBER     : [0-9]+ ('.' [0-9]+)? ([eE] [+-]? [0-9]+)?;

// Duration units - must come before IDENTIFIER to get priority
DURATION_UNIT : 'ms' | 's' | 'm' | 'h' | 'd' | 'w' | 'y';

// Identifiers - allows colons for metric names like "http_requests:rate5m"
IDENTIFIER : [a-zA-Z_:][a-zA-Z0-9_:]*;

// Strings
STRING     : '"' (~["\r\n\\] | '\\' .)* '"'
           | '\'' (~['\r\n\\] | '\\' .)* '\''
           | '`' (~[`\\] | '\\' .)* '`'
           ;

// Comments
COMMENT    : '#' ~[\r\n]* -> skip;
MULTILINE_COMMENT : '/*' .*? '*/' -> skip;

// Whitespace
WS         : [ \t\r\n]+ -> skip;