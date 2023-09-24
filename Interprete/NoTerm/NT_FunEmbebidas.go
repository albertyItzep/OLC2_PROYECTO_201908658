package noterm

import interprete "OLC2_PROYECTO_201908658/Interprete"

type NT_DataCasteo struct {
	tipoFinal string
	valor     interprete.AbstrExpr
	linea     int
	columna   int
}

// Constructor for NT_DataCasteo
func NewNT_DataCasteo(tipoFinal string, valor interprete.AbstrExpr, linea int, columna int) *NT_DataCasteo {
	o := new(NT_DataCasteo)
	o.tipoFinal = tipoFinal
	o.valor = valor
	o.linea = linea
	o.columna = columna
	return o
}

func (ntCasteo NT_DataCasteo) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	//vendra un tipo de casteo final
	if ntCasteo.tipoFinal == "Int" {
		tmp := ntCasteo.valor.Interpretar(ctx)
		return ctx.Conversor.Ampliar(tmp, interprete.Integer, ntCasteo.linea, ntCasteo.columna)
	} else if ntCasteo.tipoFinal == "Float" {
		tmp := ntCasteo.valor.Interpretar(ctx)
		return ctx.Conversor.Ampliar(tmp, interprete.Float, ntCasteo.linea, ntCasteo.columna)
	} else if ntCasteo.tipoFinal == "String" {
		tmp := ntCasteo.valor.Interpretar(ctx)
		return ctx.Conversor.Ampliar(tmp, interprete.String, ntCasteo.linea, ntCasteo.columna)
	}
	ctx.AddError("Tipo de casteo no compatible", ntCasteo.linea, ntCasteo.columna)
	return interprete.NewNil()
}
