package noterm

import interprete "OLC2_PROYECTO_201908658/Interprete"

type NT_BloqueContent struct {
	nSentencias []interprete.AbstrExpr
}

func (lBlock NT_BloqueContent) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	for _, sentencia := range lBlock.nSentencias {
		sentencia.Interpretar(ctx)
	}
	return interprete.NewNil()
}

func NewBlockContent() *NT_BloqueContent {
	return &NT_BloqueContent{
		nSentencias: make([]interprete.AbstrExpr, 0),
	}
}

func (lBlock *NT_BloqueContent) AddSentencia(sentencia interprete.AbstrExpr) {
	lBlock.nSentencias = append(lBlock.nSentencias, sentencia)
}
