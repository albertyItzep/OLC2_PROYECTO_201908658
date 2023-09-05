grammar Control;

// tokens
ID      : [a-zA-Z_][a-zA-Z0-9_]*;
INT     : [0-9]+ ;
DECIMAL : [0-9]+'.'[0-9]+;
RANGO   : [0-9]+ '...' [0-9]+;
STRING  :  '"' ('\\"'|~["])* '"' ;
WS      :[ \t\r\n] -> skip;
LINEALCOMENT: '//' ~[\r\n]* -> channel(HIDDEN)  ;
BlockComment: '/*' .*? '*/' -> skip;
//rules

prog
    :block EOF              #ProgInicial
;

block: (sentencias)*        #BlockContent       
   ;

sentencias
    : declaracion           #SentenciaDeclaracion
    | decVector             #SentenciaDeclaracionVector
    | decMatriz             #SentenciaDeclaracionMatrix
    | decFunciones          #SentenciaDeclaracionFuncion
    | asignacion            #SentenciaAsignacion
    | insIf                 #SentenciaIf
    | insSwitch             #SentenciaSwitch
    | instWhile             #SentenciaWhile
    | instFor               #SentenciaFor
    | instGuard             #SentenciaGuard
    | instBreak             #SentenciaBreak
    | instContinue          #SentenciaContinue
    | instReturn            #SentenciaReturn
    | funcsVectoriales      #SentenciaFuncVectoriales
    | llamadaFuncion        #SentenciaLLamadaFuncion
    | instPrint             #SentenciaPrint
    | instStruct            #SentenciaStruct
    | decStruct             #SentenciaDeclaracionStruct
    | llamadaMetodos        #SentenciaLLamadaMetodos
    ;

//if
insIf
    : 'if' expr '{' block '}' (instElseIf)* (insElse)*
    ;
instElseIf
    : 'else' 'if' expr '{' block '}'
    ;
insElse
    : 'else' '{' block '}'
    ;

//switch
insSwitch
    : 'switch' expr '{' instCase* instDefault? '}';

instCase
    : 'case' expr ':' block
    ;
instDefault
    : 'default' ':' block
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
    : 'var' ID (':' '[' (tipovariable|ID) ']')? '=' defVector
    | 'let' ID ':' '[' (tipovariable|ID) ']' '=' defVector
    ;

//definicion de un vector
defVector
    : '[' numList ']'
    | '[' expresionList ']'
    | '[' objectsList ']'
    | '[' ']'
    | expr
    ;
funcsVectoriales
    : ID '.' 'append' '('(expr | (ID '(' (lDupla)? ')'))')'
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
    | expresionList
    | objectsList
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

//llamada atributos dentro de la funcion con self y fuera
llamAtributos
    : ID ('[' INT ']')? ('.' ID ('[' INT ']')?)+
    | 'self' '.' ID ('[' INT ']')? ('.' ID ('[' INT ']')?)*
    ;
llamadaMetodos
    : ID ('[' INT ']')? ('.' ID('[' INT ']')?)+ '(' listaLlamadaParametros? ')'
    | 'self' '.' ID ('[' INT ']')?('.' ID ('[' INT ']')?)* '(' listaLlamadaParametros? ')'
    ;


//funciones
decFunciones
    : 'func' ID '(' listaParametros? ')' ('->' (tipovariable|ID))? '{' block '}'
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
    : 'print' '(' (expr|llamadaFuncion) (',' (expr|llamadaFuncion))* ')' #IntruccionPrint
    ;

instCasteos
    : tipovariable '(' expr ')'
    ;

numList
    : numList ',' INT
    | INT
    ;
expresionList
    : expresionList ',' expr
    | expr
    ;

objectsList
    : objectsList ',' ID '(' (lDupla)? ')'
    | ID '(' (lDupla)? ')'
    ;

declaracion
    : 'var' ID ':' tipovariable '=' expr    #Dec_VTipoValor
    | 'var' ID ':' tipovariable '?'         #Dec_VTipo
    | 'var' ID '=' expr                     #Dec_VVal
    | 'let' ID '=' expr                     #Dec_CVal
    | 'let' ID ':' tipovariable '=' expr    #Dec_CTipoValor
    ;
asignacion
    : ID '+=' expr                                              #Asignacion_Aumento
    | ID '-=' expr                                              #Asignacion_Decremento
    | ID '=' expr                                               #Asignacion_ValorGen
    | ID '[' (INT|ID) ']' '+=' expr                             #Asignacion_VectorAumento
    | ID '[' (INT|ID) ']' '-=' expr                             #Asignacion_VectorDecremento
    | ID '[' (INT|ID) ']' ('[' (INT|ID) ']')+ '-=' expr         #Asignacion_MatrixAumento
    | ID '[' (INT|ID) ']' ('[' (INT|ID) ']')+ '+=' expr         #Asignacion_MatrixDecremento
    | ID '[' (INT|ID) ']' '=' expr                              #Asignacion_VectorGeneral
    | ID '[' (INT|ID) ']' ('[' (INT|ID) ']')+ '=' expr          #Asignacion_MatrixGeneral
    | llamAtributos '=' expr                                    #Asignacion_LlAtribGeneral
    | llamAtributos '+=' expr                                   #Asignacion_LlAtribAumento
    | llamAtributos '-=' expr                                   #Asignacion_LlAtribDecremento
    ;
expr
    : '-' expr                                  #Expr_ValNumNeg
    | '!' rigth = expr                          #Expr_ValNegacion
    | left=expr op=('*'|'/') rigth=expr         #Expr_OpMulDiv
    | left=expr op=('+'| '-'|'%') rigth=expr    #Expr_OpSumRes
    | left=expr op=('>='|'>') right=expr        #Expr_ValidaMayQue
    | left=expr op=('<='|'<') right=expr        #Expr_ValidaMenQue
    | left=expr op=('=='|'!=') right=expr       #Expr_ValidaIgualDif
    | left=expr '&&' right=expr                 #Expr_ValidaAnd
    | left=expr '||' right=expr                 #Expr_ValidaOr
    | '(' expr ')'                              #Expr_ParExpre
    | INT                                       #Expr_Entero
    | DECIMAL                                   #Expr_Decimal
    | STRING                                    #Expr_StringChar
    | ('true' | 'false')                        #Expr_Boolean
    | 'nil'                                     #Expr_Nil
    | ID                                        #Expr_ID
    | ID '.' 'count'                            #Expr_Conteo
    | ID '[' (INT|ID) ']'                       #Expr_PosVector        
    | ID '[' (INT|ID) ']' ('[' (INT|ID) ']')+   #Expr_PosMatrix
    | RANGO                                     #Expr_InstRango
    | instCasteos                               #Expr_InstCasteo
    | llamAtributos                             #Expr_LlamAtributos
    ;

tipovariable
    : 'String'
    | 'Float'
    | 'Int'
    | 'Bool'
    | 'Character'
    ;
