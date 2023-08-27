lexer grammar ClickhouseLexer;

CREATE: C R E A T E;
ALTER: A L T E R;
DATABASE: D A T A B A S E;
TABLE: T A B L E;
IF: I F;
NOT: N O T;
EXISTS: E X I S T S;
ON: O N;
CLUSTER: C L U S T E R;
ENGINE: E N G I N E;
COMMENT: C O M M E N T;

ADD: A D D;

COLUMN: C O L U M N;


ID: [a-zA-Z_]([a-zA-Z_]|[0-9])*;

EQUAL: '=';
LEFT_P: '(';
RIGHT_P: ')';
COMMA: ',';
DOT: '.';

NUMBER: [0-9]+;
STRING: '\''~('\'')*?'\'';

fragment
A: [aA];
B: [bB];
C: [cC];
D: [dD];
E: [eE];
F: [fF];
G: [gG];
H: [hH];
I: [iI];
J: [jJ];
K: [kK];
L: [lL];
M: [mM];
N: [nN];
O: [oO];
P: [pP];
Q: [qQ];
R: [rR];
S: [sS];
T: [tT];
U: [uU];
V: [vV];
W: [wW];
X: [xX];
Y: [yY];
Z: [zZ];



WS: [ \t\r\n] -> skip;