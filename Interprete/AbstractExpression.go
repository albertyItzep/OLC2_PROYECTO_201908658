package interprete

//valores que yo puedo regresar, cadenas, booleans, int, float no importa pero puedo retornar
type Resultado struct {
	valor  int
	valorS string
	Nil    bool
}

// Constructor for Resultado
func NewResultado(valor int, valorS string) *Resultado {
	o := new(Resultado)
	o.valor = valor
	o.valorS = valorS
	return o
}

func NewNil() *Resultado {
	return &Resultado{
		valor:  0,
		valorS: "0",
		Nil:    true,
	}
}

//interfaz de patron interprete
type AbstrExpr interface {
	Interpretar(ctx *Contexto) *Resultado
}
