package terminales

import interprete "OLC2_PROYECTO_201908658/Interprete"

type T_Boolean struct {
	val     string
	linea   int
	columna int
}

// Constructor for T_Boolean
func NewT_Boolean(val string, linea int, columna int) *T_Boolean {
	o := new(T_Boolean)
	o.val = val
	o.linea = linea
	o.columna = columna
	return o
}

func (tBool *T_Boolean) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	if tBool.val == "true" {
		return interprete.NewBoolLiteral(true)
	} else {
		return interprete.NewBoolLiteral(false)
	}
}
