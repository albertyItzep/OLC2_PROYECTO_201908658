package terminales

import (
	interprete "OLC2_PROYECTO_201908658/Interprete"
	"strconv"
)

type T_IntNum struct {
	Num     string
	Linea   int
	Columna int
}

func (tNum *T_IntNum) Interpretar(ctx *interprete.Contexto) *interprete.Resultado {
	num, _ := strconv.Atoi(tNum.Num)
	return interprete.NewIntLiteral(num)
}

// Constructor for T_Num
func NewT_IntNum(num string, linea int, columna int) *T_IntNum {
	o := new(T_IntNum)
	o.Num = num
	o.Linea = linea
	o.Columna = columna
	return o
}
