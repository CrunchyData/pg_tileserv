lexer grammar CqlLexer;

/*============================================================================
# Enable case-insensitive grammars
#============================================================================*/

fragment A : [aA];
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];

/*============================================================================
# Definition of COMPARISON operators
#============================================================================*/

ComparisonOperator : EQ | NEQ | LT | GT | LTEQ | GTEQ;
LT : '<';
EQ : '=';
GT : '>';
NEQ : LT GT;
GTEQ : GT EQ;
LTEQ : LT EQ;

/*============================================================================
# Definition of BOOLEAN literals
#============================================================================*/

BooleanLiteral : T R U E | F A L S E;

/*============================================================================
# Definition of LOGICAL operators
#============================================================================*/

AND : A N D;
OR : O R;
NOT : N O T;

/*============================================================================
# Definition of COMPARISON operators
#============================================================================*/

LIKE : L I K E;
ILIKE : I L I K E;
BETWEEN : B E T W E E N;
IS : I S;
NULL: N U L L;
IN: I N;

/*============================================================================
# Definition of ARITHMETIC operators
#============================================================================*/

ArithmeticOperator : PLUS | MINUS | ASTERISK | SOLIDUS | PERCENT;

/*============================================================================
# Definition of SPATIAL operators
#============================================================================*/

SpatialOperator : E Q U A L S | D I S J O I N T | T O U C H E S | W I T H I N | O V E R L A P S
                | C R O S S E S | I N T E R S E C T S | C O N T A I N S;

/*
# NOTE: The distance operator BEYOND is not currently included.
#       It is equivalent to NOT DWITHIN.
*/
DistanceOperator : D W I T H I N;

/*============================================================================
# Definition of TEMPORAL operators
#============================================================================*/

/*
TemporalOperator : A F T E R | B E F O R E | B E G I N S | B E G U N B Y | T C O N T A I N S | D U R I N G
                 | E N D E D B Y | E N D S | T E Q U A L S | M E E T S | M E T B Y | T O V E R L A P S
                 | O V E R L A P P E D B Y | A N Y I N T E R A C T S;
*/

/*============================================================================
# Definition of geometry types
#============================================================================*/

POINT: P O I N T;
LINESTRING: L I N E S T R I N G;
POLYGON: P O L Y G O N;
MULTIPOINT: M U L T I P O I N T;
MULTILINESTRING: M U L T I L I N E S T R I N G;
MULTIPOLYGON: M U L T I P O L Y G O N;
GEOMETRYCOLLECTION: G E O M E T R Y C O L L E C T I O N;
ENVELOPE: E N V E L O P E;

/*============================================================================
# Definition of numeric and text literals
#============================================================================*/

//-- NOTE: order is important!  This def must go here
NumericLiteral : UnsignedNumericLiteral | SignedNumericLiteral;

CharacterStringLiteralStart : QUOTE -> more, mode(STR);// (Character)* QUOTE;

/*============================================================================
# Definition of property identifiers
#============================================================================*/

//Identifier : IdentifierStart (COLON | PERIOD | IdentifierPart)* | DOUBLEQUOTE Identifier DOUBLEQUOTE;

Identifier : IdentifierStart IdentifierPart* | DOUBLEQUOTE Identifier DOUBLEQUOTE;
IdentifierStart : ALPHA;
IdentifierPart : ALPHA | DIGIT | UNDERSCORE | DOLLAR;

ALPHA : [A-Za-z];

DIGIT : [0-9];

OCTOTHORP : '#';
DOLLAR : '$';
UNDERSCORE : '_';
DOUBLEQUOTE : '"';
PERCENT : '%';
AMPERSAND : '&';
QUOTE : '\'';
LEFTPAREN : '(';
RIGHTPAREN : ')';
LEFTSQUAREBRACKET : '[';
RIGHTSQUAREBRACKET : ']';
ASTERISK : '*';
PLUS : '+';
COMMA : ',';
MINUS : '-';
PERIOD : '.';
SOLIDUS : '/';
COLON : ':';
SEMICOLON : ';';
QUESTIONMARK : '?';
VERTICALBAR : '|';
BIT : '0' | '1';
HEXIT : DIGIT | A | B | C | D | E | F;

/*============================================================================
# Definition of NUMERIC literals
#============================================================================*/

UnsignedNumericLiteral : ExactNumericLiteral | ApproximateNumericLiteral;
SignedNumericLiteral : (Sign)? ExactNumericLiteral | ApproximateNumericLiteral;
ExactNumericLiteral : UnsignedInteger  (PERIOD (UnsignedInteger)? )?
                        |  PERIOD UnsignedInteger;
ApproximateNumericLiteral : Mantissa 'E' Exponent;
Mantissa : ExactNumericLiteral;
Exponent : SignedInteger;
SignedInteger : (Sign)? UnsignedInteger;
UnsignedInteger : (DIGIT)+;
Sign : PLUS | MINUS;

/*============================================================================
# Definition of TEMPORAL literals
#============================================================================*/

TemporalLiteral : Instant
    // | Interval
    ;
Instant : FullDate | FullDate 'T' UtcTime | NOW LEFTPAREN RIGHTPAREN;
//Interval : (InstantInInterval)? SOLIDUS (InstantInInterval)?;
//InstantInInterval : '..' | Instant;
FullDate : DateYear '-' DateMonth '-' DateDay;
DateYear : DIGIT DIGIT DIGIT DIGIT;
DateMonth : DIGIT DIGIT;
DateDay : DIGIT DIGIT;
UtcTime : TimeHour ':' TimeMinute  (':' TimeSecond)? (TimeZoneOffset)?;
TimeZoneOffset : 'Z' | Sign TimeHour ':' TimeMinute;
TimeHour : DIGIT DIGIT;
TimeMinute : DIGIT DIGIT;
TimeSecond : DIGIT DIGIT (PERIOD (DIGIT)+)?;
NOW : N O W;

/*============================================================================
# ANTLR ignore whitespace
#============================================================================*/

WS : [ \t\r\n]+ -> skip;// channel(HIDDEN) ; // skip spaces, tabs, newlines

/*============================================================================
# ANTLR mode for CharacterStringLiteral with whitespaces
#============================================================================*/

mode STR;
CharacterStringLiteral: '\'' -> mode(DEFAULT_MODE);
QuotedQuote: '\'\'' -> more;
Character : ~['] -> more; // (ALPHA | DIGIT | SpecialCharacter | QuoteQuote | ' ')
