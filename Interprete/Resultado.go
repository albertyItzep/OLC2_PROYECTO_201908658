package interprete

type TipoD uint

const (
	Nil       TipoD = 0
	Character TipoD = 1
	Bool      TipoD = 2
	Integer   TipoD = 3
	Float     TipoD = 4
	String    TipoD = 5
	Break     TipoD = 6
	Continue  TipoD = 7
	Rango     TipoD = 8
	Vector    TipoD = 9
	Compuesto TipoD = 10
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
	case Break:
		return "Break"
	case Continue:
		return "Continue"
	case Rango:
		return "Rango"
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
	TipoVector   TipoD
	TipoVariable bool
	Rango        []int
	Vector       []Resultado
	Return       bool
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
		TipoVector:   Nil,
		Rango:        make([]int, 0),
		Vector:       make([]Resultado, 0),
		Return:       false,
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
		TipoVector:   Nil,
		Rango:        make([]int, 0),
		Vector:       make([]Resultado, 0),
		Return:       false,
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
		TipoVector:   Nil,
		Rango:        make([]int, 0),
		Vector:       make([]Resultado, 0),
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
		TipoVector:   Nil,
		Nil:          false,
		Rango:        make([]int, 0),
		Vector:       make([]Resultado, 0),
		TipoVariable: false,
		Return:       false,
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
		TipoVector:   Nil,
		Nil:          false,
		Rango:        make([]int, 0),
		Vector:       make([]Resultado, 0),
		TipoVariable: false,
		Return:       false,
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
		TipoVector:   Nil,
		Nil:          false,
		TipoVariable: false,
		Rango:        make([]int, 0),
		Vector:       make([]Resultado, 0),
		Return:       false,
	}
}

func NewBreakLiteral() *Resultado {
	return &Resultado{
		Valor:        0,
		ValorF:       0.0,
		ValorS:       "",
		ValorC:       "",
		ValorB:       true,
		Tipo:         Break,
		TipoVector:   Nil,
		Nil:          false,
		Rango:        make([]int, 0),
		Vector:       make([]Resultado, 0),
		TipoVariable: false,
		Return:       false,
	}
}

func NewContinueLiteral() *Resultado {
	return &Resultado{
		Valor:        0,
		ValorF:       0.0,
		ValorS:       "",
		ValorC:       "",
		ValorB:       true,
		Tipo:         Continue,
		TipoVector:   Nil,
		Nil:          false,
		TipoVariable: false,
		Rango:        make([]int, 0),
		Vector:       make([]Resultado, 0),
		Return:       false,
	}
}

func NewRangoLiteral() *Resultado {
	return &Resultado{
		Valor:        0,
		ValorF:       0.0,
		ValorS:       "",
		ValorC:       "",
		ValorB:       true,
		Tipo:         Rango,
		TipoVector:   Nil,
		Nil:          false,
		TipoVariable: false,
		Rango:        make([]int, 0),
		Vector:       make([]Resultado, 0),
		Return:       false,
	}
}

func NewVectorLiteral(Tipo TipoD) *Resultado {
	return &Resultado{
		Valor:        0,
		ValorF:       0.0,
		ValorS:       "",
		ValorC:       "",
		ValorB:       true,
		Tipo:         Vector,
		TipoVector:   Tipo,
		Nil:          false,
		TipoVariable: false,
		Rango:        make([]int, 0),
		Vector:       make([]Resultado, 0),
		Return:       false,
	}
}
func NewVectorValLiteral(Tipo TipoD, Resultados []Resultado) *Resultado {
	return &Resultado{
		Valor:        0,
		ValorF:       0.0,
		ValorS:       "",
		ValorC:       "",
		ValorB:       true,
		Tipo:         Vector,
		TipoVector:   Tipo,
		Nil:          false,
		TipoVariable: false,
		Rango:        make([]int, 0),
		Vector:       Resultados,
		Return:       false,
	}
}
