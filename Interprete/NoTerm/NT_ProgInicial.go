package noterm

import interprete "OLC2_PROYECTO_201908658/Interprete"

type NT_ProgInicial struct {
	lbloques interprete.AbstrExpr
}

func (bl *NT_ProgInicial) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	return bl.lbloques.Interpretar(ctx)
}

func NewNT_ProgInicial(lbloques interprete.AbstrExpr) *NT_ProgInicial {
	return &NT_ProgInicial{lbloques: lbloques}
}
