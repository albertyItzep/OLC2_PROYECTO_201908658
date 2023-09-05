package terminales

import interprete "OLC2_PROYECTO_201908658/Interprete"

type T_TipoVariable struct {
	Tipo string
}

// Constructor for T_TipoVariable
func NewT_TipoVariable(tipo string) *T_TipoVariable {
	o := new(T_TipoVariable)
	o.Tipo = tipo
	return o
}

func (tTipo *T_TipoVariable) GetTipoD() interprete.TipoD {
	switch tTipo.Tipo {
	case "String":
		return interprete.String
	case "Float":
		return interprete.Float
	case "Int":
		return interprete.Integer
	case "Bool":
		return interprete.Bool
	case "Character":
		return interprete.Character
	default:
		return interprete.Nil
	}
}
