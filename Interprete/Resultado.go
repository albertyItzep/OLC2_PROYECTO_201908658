package interprete

type TipoD uint

const (
	Nil       TipoD = 0
	Character TipoD = 1
	Bool      TipoD = 2
	Integer   TipoD = 3
	Float     TipoD = 4
	String    TipoD = 5
	Compuesto TipoD = 6
)

func (t TipoD) String() string {
	switch t {
	case Nil:
		return "NIL"
	case Character:
		return "Character"
	case Integer:
		return "Integer"
	case Float:
		return "Float"
	case String:
		return "String"
	case Bool:
		return "Bool"
	case Compuesto:
		return "Compuesto"
	default:
		return "Desconocido"
	}
}

// valores que yo puedo regresar, cadenas, booleans, int, float no importa pero puedo retornar
type Resultado struct {
	Valor        int
	ValorF       float64
	ValorS       string
	ValorC       string
	ValorB       bool
	Nil          bool
	Tipo         TipoD
	TipoVariable bool
}

// Constructor for Resultado
func NewResultado(valor int, valorS string) *Resultado {
	o := new(Resultado)
	o.Valor = valor
	o.ValorS = valorS
	o.TipoVariable = false
	return o
}

func NewNil() *Resultado {
	return &Resultado{
		Valor:        0,
		ValorF:       0,
		ValorS:       "",
		ValorC:       "",
		ValorB:       false,
		Tipo:         Nil,
		Nil:          true,
		TipoVariable: false,
	}
}

func NewFloatLiteral(valor float64) *Resultado {
	return &Resultado{
		Valor:        0,
		ValorF:       valor,
		ValorS:       "",
		ValorC:       "",
		ValorB:       false,
		Tipo:         Float,
		Nil:          false,
		TipoVariable: false,
	}
}

func NewIntLiteral(valor int) *Resultado {
	return &Resultado{
		Valor:        valor,
		ValorF:       0.0,
		ValorS:       "",
		ValorC:       "",
		ValorB:       false,
		Tipo:         Integer,
		Nil:          false,
		TipoVariable: false,
	}
}

func NewStringLiteral(valor string) *Resultado {
	return &Resultado{
		Valor:        0,
		ValorF:       0.0,
		ValorS:       valor,
		ValorC:       "",
		ValorB:       false,
		Tipo:         String,
		Nil:          false,
		TipoVariable: false,
	}
}

func NewCharLiteral(valor string) *Resultado {
	return &Resultado{
		Valor:        0,
		ValorF:       0.0,
		ValorS:       "",
		ValorC:       valor,
		ValorB:       false,
		Tipo:         Character,
		Nil:          false,
		TipoVariable: false,
	}
}

func NewBoolLiteral(valor bool) *Resultado {
	return &Resultado{
		Valor:        0,
		ValorF:       0.0,
		ValorS:       "",
		ValorC:       "",
		ValorB:       valor,
		Tipo:         Bool,
		Nil:          false,
		TipoVariable: false,
	}
}
