grammar Control;

// tokens
ID      : [a-zA-Z_][a-zA-Z0-9_]*;
INT     : [0-9]+ ;
DECIMAL : [0-9]+'.'[0-9]+;
RANGO   : [0-9]+ '...' [0-9]+;
STRING  : '"' (~["\r\n] | '""')* '"' ;
WS      :[ \t\r\n] -> skip;
LINEALCOMENT: '//' ~[\r\n]* -> channel(HIDDEN)  ;
//rules

prog: block EOF;

block: (sentencias)*                           
   ;

sentencias
    : declaracion
    | decVector
    | decMatriz
    | decFunciones
    | asignacion          
    | instWhile
    | instFor
    | instGuard
    | instBreak
    | instContinue
    | instReturn
    | funcsVectoriales
    | llamadaFuncion
    | instPrint
    | instStruct
    | decStruct
    ;

//while
instWhile
    : 'while' expr '{' block '}'
    ;
//for
instFor
    : 'for' ID 'in' expr '{' block '}'
    ;
//guard
instGuard
    :'guard' expr 'else' '{' block '}'
    ;
//break
instBreak
    : 'break'
    ;
//continue
instContinue
    : 'continue'
    ;
//return
instReturn
    : 'return' 
    | 'return' expr
    ;
//vectores
decVector
    : 'var' ID ':' '[' tipovariable ']' '=' defVector
    ;

//definicion de un vector
defVector
    : '[' numList ']'
    | '[' ']'
    | expr
    ;
funcsVectoriales
    : ID '.' 'append' '('expr')'
    | ID '.' 'remove' '(''at' ':' expr')'
    | ID '.' 'removeLast' '(' ')'
    | ID '.' 'isEmpty' '(' ')'
    //| ID '.' ID '('expr')' //cualquier funcion con punto
    ;

// matrices
decMatriz 
    : 'var' ID ':' tipoMatriz '=' defMatrix
    ;

tipoMatriz
    : '[' tipoMatriz ']'
    | '[' tipovariable ']'
    ;

defMatrix
    : listaValoresMat
    | simpleMax
    ;

listaValoresMat
    : '[' listaValoresMat2 ']'
    ;

listaValoresMat2
    : listaValoresMat2 ',' listaValoresMat
    | listaValoresMat
    | numList
    ;

simpleMax
    : tipoMatriz '(' 'repeating' ':' simpleMax ',' 'count' ':' INT ')'
    | tipoMatriz '(' 'repeating' ':' expr ',' 'count' ':' INT ')'
    ;

//structs
instStruct
    : 'struct' ID '{' listaAtributos* '}'
    ;

listaAtributos
    : ('let'|'var') ID (':' (tipovariable|ID))? ('=' expr)? ';'?
    | 'mutating'? decFunciones
    ;

//creacion de structs
decStruct
    : ('var'|'let') ID (':'ID)? '=' ID '(' (lDupla)? ')'
    | ('var'|'let') ID (':'ID)? '=' ID?
    | ('var'|'let') ID (':'ID)? '=' llamadaFuncion
    ;
lDupla
    : (','? ID ':' (expr|llamadaFuncion))+
    ;
//funciones
decFunciones
    : 'func' ID '(' listaParametros? ')' ('->' tipovariable)? '{' block '}'
    ;

listaParametros
    : listaParametros ',' (ID| '_')? ID ':' 'inout'? (tipovariable | '[' tipovariable ']' | ID)
    | (ID| '_')? ID ':' 'inout'? (tipovariable |  '[' tipovariable ']' | ID)
    ;


//Llamada de funciones
llamadaFuncion
    : ID '(' listaLlamadaParametros ')'
    ;

listaLlamadaParametros
    : listaLlamadaParametros ',' (ID ':')? '&'? expr
    | (ID ':')? '&'? expr
    ;

//funciones embebidas
instPrint
    : 'print' '(' (','? expr|llamadaFuncion)+ ')'
    ;

instCasteos
    : tipovariable '(' expr ')'
    ;

numList
    : numList ',' INT
    | INT
    ;
declaracion
    : 'var' ID ':' tipovariable '=' expr
    | 'var' ID ':' tipovariable '?'
    | 'var' ID '=' expr
    | 'let' ID ':' tipovariable '=' expr
    | 'let' ID '=' expr
    ;
asignacion
    : ID '+=' expr
    | ID '-=' expr
    | ID '=' expr
    | ID '[' (INT|ID) ']' '+=' expr
    | ID '[' (INT|ID) ']' '-=' expr
    | ID '[' (INT|ID) ']' ('[' (INT|ID) ']')+ '-=' expr
    | ID '[' (INT|ID) ']' ('[' (INT|ID) ']')+ '+=' expr
    | ID '[' (INT|ID) ']' '=' expr
    | ID '[' (INT|ID) ']' ('[' (INT|ID) ']')+ '=' expr
    | 'self' '.' asignacion
    ;
expr
    : '-' expr
    | '!' rigth = expr                          
    | left=expr op=('*'|'/') rigth=expr         
    | left=expr op=('+'| '-'|'%') rigth=expr   
    | left=expr op=('>='|'>') right=expr
    | left=expr op=('<='|'<') right=expr 
    | left=expr op=('=='|'!=') right=expr
    | left=expr '&&' right=expr
    | left=expr '||' right=expr
    | '(' expr ')'     
    | INT                                       
    | DECIMAL                                   
    | STRING
    | ('true' | 'false')
    | 'nil'
    | ID
    | ID '.' 'count'
    | ID '[' (INT|ID) ']'//para un valor dentro de un vector
    | ID '[' (INT|ID) ']' ('[' (INT|ID) ']')+
    | RANGO
    | 'self' '.' ID
    | ID '.' ID
    | instCasteos                                 
    ;

tipovariable
    : 'String'
    | 'Float'
    | 'Int'
    | 'Bool'
    | 'Character'
    ;