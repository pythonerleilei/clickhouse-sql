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

ORDER: O R D E R;
PARTITION: P A R T I T I O N;
PRIMARY: P R I M A R Y;
SAMPLE: S A M P L E;
TTL: T T L;
SETTINGS: S E T T I N G S;

KEY: K E Y;
BY: B Y;

DEFAULT: D E F A U L T;
CODEC: C O D E C;
INDEX: I N D E X;
TYPE: T Y P E;
GRANULARITY: G R A N U L A R I T Y;

INTERVAL: I N T E R V A L ;

SECOND: S E C O N D ;
MINUTE: M I N U T E ;
HOUR: H O U R ;
DAY: D A Y ;
WEEK: W E E K ;
MONTH: M O N T H ;
QUARTER: Q U A R T E R ;
YEAR: Y E A R ;


ADD: A D D;

COLUMN: C O L U M N;


ID: [a-zA-Z_]([a-zA-Z_]|[0-9])*;

EQUAL: '=';
LEFT_P: '(';
RIGHT_P: ')';
COMMA: ',';
DOT: '.';

PLUS: '+';


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