grammar Control;

// tokens
ID      : [a-zA-Z_][a-zA-Z0-9_]*;
INT     : [0-9]+ ;
DECIMAL : [0-9]+'.'[0-9]+;
STRING  : '"' (~["\r\n] | '""')* '"' ;
WS      :[ \t\r\n] -> skip;
LINEALCOMENT: '//' ~[\r\n]* -> channel(HIDDEN)  ;
//rules

prog: block EOF;

block: (sentencias)*                           
   ;

sentencias
    : declaracion
    | asignacion                            
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
    ;
expr
    : '-' expr
    | '!' rigth = expr                          
    | left=expr op=('*'|'/') rigth=expr         
    | left=expr op=('+'| '-'|'%') rigth=expr        
    | INT                                       
    | DECIMAL                                   
    | STRING
    | ('true' | 'false')
    | 'nil'                                   
    ;

tipovariable
    : 'String'
    | 'Float'
    | 'Int'
    | 'Bool'
    | 'Character'
    ;