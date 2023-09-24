package noterm

import interprete "OLC2_PROYECTO_201908658/Interprete"

type NT_Funcion1 struct {
	idFunc       string
	parametros   []interprete.AbstrExpr
	valorRetoino interprete.AbstrExpr
	blocck       []interprete.AbstrExpr
	linea        int
	columna      int
}

// Constructor for NT_Funcion1
func NewNT_Funcion1(idFunc string, parametros []interprete.AbstrExpr, valorRetoino interprete.AbstrExpr, blocck []interprete.AbstrExpr, linea int, columna int) *NT_Funcion1 {
	o := new(NT_Funcion1)
	o.idFunc = idFunc
	o.parametros = parametros
	o.valorRetoino = valorRetoino
	o.blocck = blocck
	o.linea = linea
	o.columna = columna
	return o
}

func (ntFuncion NT_Funcion1) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	panic("not implemented") // TODO: Implement
}
