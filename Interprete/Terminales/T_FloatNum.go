package terminales

import (
	interprete "OLC2_PROYECTO_201908658/Interprete"
	"strconv"
)

type T_FloatNum struct {
	Num     string
	Linea   int
	Columna int
}

func (tNum *T_FloatNum) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	num, _ := strconv.ParseFloat(tNum.Num, 64)
	return interprete.NewFloatLiteral(num)
}

// Constructor for T_FloatNum
func NewT_FloatNum(num string, linea int, columna int) *T_FloatNum {
	o := new(T_FloatNum)
	o.Num = num
	o.Linea = linea
	o.Columna = columna
	return o
}
