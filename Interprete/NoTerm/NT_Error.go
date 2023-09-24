package noterm

import interprete "OLC2_PROYECTO_201908658/Interprete"

type NT_Error struct {
	valError string
	linea    int
	columna  int
}

//no recordaba que tenia este objeto, volvia crear uno y aplique el otro
func (ntError *NT_Error) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	ctx.AddError(ntError.valError, ntError.linea, ntError.columna)
	return interprete.NewNil()
}

// Constructor for NT_Error
func NewNT_Error(valError string) *NT_Error {
	o := new(NT_Error)
	o.valError = valError
	return o
}
