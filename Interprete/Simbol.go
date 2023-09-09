package interprete

//tipo nativo esta en resultado, el tipo compuesto lo vamos a agregar ahora
type TipoCompuesto uint

const (
	Nativo TipoCompuesto = 0
	Struct TipoCompuesto = 1
	Matriz TipoCompuesto = 2
)

func (t TipoCompuesto) String() string {
	switch t {
	case Nativo:
		return "Nativo"
	case Struct:
		return "Struct"
	case Matriz:
		return "Matriz"
	default:
		return "Desconocido"
	}
}

type Symbol struct {
	Constante        bool
	Nombre           string
	Tipo             TipoD //tipos nativos
	TipoVector       TipoD
	TipoCom          TipoCompuesto
	NumeroParametros int
	ListaParametros  []Symbol
	Resultado        *Resultado
	Resultados       []Resultado
	linea            int
	columna          int
}

func NewSymbol(nombre string, tipo TipoD, tipoCompuesto TipoCompuesto, NumPar int, ListaPar []Symbol, Resultado *Resultado, Resultados []Resultado, linea int, columna int) *Symbol {
	return &Symbol{
		Nombre:           nombre,
		Tipo:             tipo,
		TipoCom:          tipoCompuesto,
		NumeroParametros: NumPar,
		ListaParametros:  ListaPar,
		Resultado:        Resultado,
		Resultados:       Resultados,
		linea:            linea,
		columna:          columna,
	}
}

func NewSymbolConst(constante bool, nombre string, tipo TipoD, tipoCompuesto TipoCompuesto, NumPar int, ListaPar []Symbol, Resultado *Resultado, Resultados []Resultado, linea int, columna int) *Symbol {
	return &Symbol{
		Constante:        constante,
		Nombre:           nombre,
		Tipo:             tipo,
		TipoCom:          tipoCompuesto,
		NumeroParametros: NumPar,
		ListaParametros:  ListaPar,
		Resultado:        Resultado,
		Resultados:       Resultados,
		linea:            linea,
		columna:          columna,
	}
}

func NewSymbolVector(constante bool, nombre string, tipo TipoD, tipoVector TipoD, tipoCompuesto TipoCompuesto, NumPar int, ListaPar []Symbol, Resultado *Resultado, Resultados []Resultado, linea int, columna int) *Symbol {
	return &Symbol{
		Constante:        constante,
		Nombre:           nombre,
		Tipo:             tipo,
		TipoCom:          tipoCompuesto,
		TipoVector:       tipoVector,
		NumeroParametros: NumPar,
		ListaParametros:  ListaPar,
		Resultado:        Resultado,
		Resultados:       Resultados,
		linea:            linea,
		columna:          columna,
	}
}
