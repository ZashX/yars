// Yars.g4
grammar Yars;

// Tokens
INT32:   'int32';
UINT32:  'uint32';
FLOAT32: 'float32';
BOOL:    'bool';
STRING:  'string';

NOT: 'not';
AND: 'and';
OR:  'or';

IN:  'in';

PLUS:  '+';
MINUS: '-';
MUL:   '*';

UNIT:      'unit';
ACTION:    'action';
PROJECT:   'project';
RECOMMEND: 'recommend';
IDENTIFIER: LETTER (LETTER | UNICODE_DIGIT)*;

L_CURLY                : '{';
R_CURLY                : '}';
L_BRACKET              : '[';
R_BRACKET              : ']';
L_RBRACKET             : '(';
R_RBRACKET             : ')';
SEMI                   : ';';
COLON                  : ':';
AT                     : '@';
COMMA                  : ',';
DOT                    : '.';
WHITESPACE: [ \r\n\t]+ -> skip;

//entities
project: PROJECT IDENTIFIER SEMI (unitStatement | actionStatement | recommendStatement)* EOF;

actionStatement: ACTION name=IDENTIFIER actor=IDENTIFIER AT object=IDENTIFIER SEMI;

recommendStatement: RECOMMEND name=IDENTIFIER actor=IDENTIFIER AT object=IDENTIFIER paramsDecl;

paramsDecl: L_CURLY (paramDecl eos)* R_CURLY;

paramDecl: name=IDENTIFIER COLON expression;

expression
    : L_RBRACKET expression R_RBRACKET     #brackets
    | op=unaryOps expression               #unaryOp
    | expression op=binaryOps expression   #binaryOp
    | lvalue                               #lvalueExpression
    ;

lvalue: name=IDENTIFIER | (name=IDENTIFIER DOT field=IDENTIFIER);

binaryOps
    : MUL
    | PLUS
    | MINUS

    | AND
    | OR
    | IN

    | IN
    ;

unaryOps
    : NOT
    | PLUS
    | MINUS
    ;

unitStatement: UNIT IDENTIFIER fieldsDecl;

fieldsDecl: L_CURLY (fieldDecl eos)* R_CURLY;

fieldDecl: IDENTIFIER type_;
type_: (arrayType | baseType);
arrayType:  L_BRACKET R_BRACKET baseType;
baseType: INT32 | UINT32 | FLOAT32 | BOOL | STRING;

eos:
	SEMI;

fragment LETTER
    : UNICODE_LETTER
    | '_'
    ;

fragment UNICODE_DIGIT
    : [\p{Nd}];

fragment UNICODE_LETTER
    : [\p{L}];