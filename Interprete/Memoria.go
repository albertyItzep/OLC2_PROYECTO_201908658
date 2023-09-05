package interprete

type Memoria struct {
	variables map[string]*Symbol
	Anterior  *Memoria
}

func NewMemoria(anterior *Memoria) *Memoria {
	return &Memoria{
		variables: make(map[string]*Symbol),
		Anterior:  anterior,
	}
}

// metodos para crear variables
func (m *Memoria) CreateSimbol(nombre string, tipo TipoD, tipoCompuesto TipoCompuesto, NumPar int, ListaPar []Symbol, Resultado *Resultado, Resultados []Resultado, linea int, columna int) (bool, string) {
	if tipo != Resultado.Tipo {
		return false, "Error Declaracion: Tipos incompatibles"
	}
	m.variables[nombre] = NewSymbol(nombre, tipo, tipoCompuesto, NumPar, ListaPar, Resultado, Resultados, linea, columna)
	return true, "Ok"
}

// creamos una variable de cualquier tipo con valor nil
func (m *Memoria) CreateSimbolTipo(nombre string, tipo TipoD, tipoCompuesto TipoCompuesto, NumPar int, ListaPar []Symbol, Resultado *Resultado, Resultados []Resultado, linea int, columna int) (bool, string) {
	m.variables[nombre] = NewSymbol(nombre, tipo, tipoCompuesto, NumPar, ListaPar, Resultado, Resultados, linea, columna)
	return true, "Ok"
}

// metodos para crear constantes
func (m *Memoria) CreateSimbolConstExp(nombre string, tipo TipoD, tipoCompuesto TipoCompuesto, NumPar int, ListaPar []Symbol, Resultado *Resultado, Resultados []Resultado, linea int, columna int) (bool, string) {
	if tipo != Resultado.Tipo {
		return false, "Error Declaracion: Tipos incompatibles"
	}
	m.variables[nombre] = NewSymbolConst(true, nombre, tipo, tipoCompuesto, NumPar, ListaPar, Resultado, Resultados, linea, columna)
	return true, "Ok"
}

func (m *Memoria) SetSymbolNativ(nombre string, Resultado *Resultado) (bool, string) {
	simbolo, res := m.variables[nombre]
	if !res {
		return false, "Error Asignacion: Variable No existe"
	}
	//crear funcion que validad el mismo tipo de dato si es int que se asigne un int
	if simbolo.Tipo != Resultado.Tipo {
		return false, "Error Asignacion: Tipos Incompatibles"
	}
	simbolo.Resultado = Resultado
	return true, "ok"
}

func (m *Memoria) Exist(nombre string) bool {
	_, res := m.variables[nombre]
	return res
}

func (m *Memoria) GetSimbolo(nombre string) (bool, *Resultado, bool) {
	simbolo, res := m.variables[nombre]
	if res {
		return simbolo.Constante, simbolo.Resultado, true
	}
	return false, NewNil(), false
}
