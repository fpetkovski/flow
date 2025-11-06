grammar PromQLPlus;

// Parser rules

query
    : letExpression EOF
    | pipeline EOF
    ;

// Let expression with bindings
letExpression
    : LET letBindings IN expression
    ;

// One or more let bindings
letBindings
    : letBinding (',' letBinding)*
    ;

// Single let binding
letBinding
    : IDENTIFIER '=' pipeline
    ;

expression
    : IDENTIFIER ('/' | '+' | '-' | '*' | '%') IDENTIFIER
    ;

pipeline
    : selector ('|' (aggregation|aligner))*
    ;

selector
    : metricIdentifier labelMatchers?
    ;

metricIdentifier
    : IDENTIFIER
    ;

labelMatchers
    : '{' labelMatcher (',' labelMatcher)* ','? '}'
    ;

labelMatcher
    : IDENTIFIER matchOp STRING
    ;

matchOp
    : '='
    | '!='
    | '=~'
    | '!~'
    ;

aggregation
    : aggregationOp (BY|WITHOUT) ('(' labelList? ')')?
    ;

labelList
    : IDENTIFIER (',' IDENTIFIER)*
    ;

aggregationOp
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
    : IDENTIFIER duration
    ;

duration
    : NUMBER DURATION_UNIT
    ;

// Lexer rules

// Keywords
LET        : 'let';
IN         : 'in';
BY         : 'by';
WITHOUT    : 'without';

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