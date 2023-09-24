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

func (m *Memoria) GetMemoria() map[string]*Symbol {
	return m.variables
}

// metodos para crear variables
func (m *Memoria) CreateSimbol(nombre string, tipo TipoD, tipoCompuesto TipoCompuesto, NumPar int, ListaPar []Symbol, Resultado *Resultado, Resultados []Resultado, linea int, columna int) (bool, string) {
	if tipo != Resultado.Tipo {
		return false, "Error Declaracion: Tipos incompatibles"
	}
	m.variables[nombre] = NewSymbol(nombre, tipo, tipoCompuesto, NumPar, ListaPar, Resultado, Resultados, linea, columna)
	return true, "Ok"
}
func (m *Memoria) CreateSimbolVector(constante bool, nombre string, tipo TipoD, tipoVector TipoD, tipoCompuesto TipoCompuesto, NumPar int, ListaPar []Symbol, Resultado *Resultado, Resultados []Resultado, linea int, columna int) (bool, string) {
	m.variables[nombre] = NewSymbolVector(constante, nombre, tipo, tipoVector, tipoCompuesto, NumPar, ListaPar, Resultado, Resultados, linea, columna)
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

	if simbolo.Constante {
		return false, "Error Asignacion: No se puede cambiar el valor de una constante"
	}
	simbolo.Resultado = Resultado
	return true, "ok"
}

func (m *Memoria) SetSymbolNativFor(nombre string, Resultado *Resultado) (bool, string) {
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

func (m *Memoria) GetSimboloVector(nombre string) (bool, TipoD, []Resultado, bool) {
	simbolo, res := m.variables[nombre]
	if res {
		return simbolo.Constante, simbolo.TipoVector, simbolo.Resultados, true
	}
	return false, Nil, nil, false
}
func (m *Memoria) GetSimboloTipoD(nombre string) TipoD {
	simbolo, res := m.variables[nombre]
	if res {
		return simbolo.Tipo
	}
	return Nil
}

func (m *Memoria) SetSymbolVector(nombre string, Resultado []Resultado) (bool, string) {
	simbolo, res := m.variables[nombre]
	if !res {
		return false, "Error Asignacion: Variable No existe"
	}
	simbolo.Resultados = Resultado
	return true, "ok"
}
