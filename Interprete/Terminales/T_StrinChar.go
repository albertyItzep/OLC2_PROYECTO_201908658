package terminales

import interprete "OLC2_PROYECTO_201908658/Interprete"

type T_StringChar struct {
	Cadena  string
	Linea   int
	Columna int
}

func (tcad *T_StringChar) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	return interprete.NewStringLiteral(tcad.Cadena)
}

// Constructor for T_StringChar
func NewT_StringChar(valor string, linea int, columna int) *T_StringChar {
	o := new(T_StringChar)
	o.Cadena = valor
	o.Linea = linea
	o.Columna = columna
	return o
}
