package interprete

//interfaz de patron interprete
type AbstrExpr interface {
	Interpretar(ctx *Contexto) *Resultado //metodo heredable todas las demas clases lo ejecutaran
}
