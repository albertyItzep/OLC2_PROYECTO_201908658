package interprete

type Contexto struct {
	consola string
	Errores []string
}

func NewContexto() *Contexto {
	return &Contexto{
		consola: "",
		Errores: make([]string, 0, 20),
	}
}

func (c *Contexto) Print(entrada string) {
	c.consola += entrada
}

func (c *Contexto) Println(entrada string) {
	c.consola += entrada + "\n"
}

func (c *Contexto) AddError(entrada string) {
	c.Errores = append(c.Errores, entrada)
}
