package interprete

import (
	"fmt"
)

type Contexto struct {
	Memoria   *Memoria
	numeroRef int
	Consola   string
	Errores   []Errs
	Compuest  map[string]DataCompuest
	Conversor *Conversor
}

func NewContexto() *Contexto {
	c := &Contexto{
		Memoria:   NewMemoria(nil),
		Consola:   "",
		Errores:   make([]Errs, 0),
		Compuest:  make(map[string]DataCompuest),
		Conversor: nil,
	}
	c.Conversor = NewConversor(c)
	return c
}

func (c *Contexto) Print(entrada string) {
	c.Consola += entrada
}

func (c *Contexto) Println(entrada string) {
	c.Consola += entrada + "\n"
}

func (c *Contexto) AddError(entrada string, linea int, columna int) {
	c.Errores = append(c.Errores, Errs{entrada, linea, columna})
}

// funciones correspondientes a la memoria
// creando ambitos creamos memorias enlazandolas una a otra comenzando en nil - primerMemmoria y asi
func (c *Contexto) AgregarAmbito() {
	headMemoria := NewMemoria(c.Memoria)
	c.Memoria = headMemoria
}

// eliminar el ambito actual
func (c *Contexto) PopAmbito() {
	if c.Memoria != nil {
		c.Memoria = c.Memoria.Anterior
	} else {
		fmt.Println("Error de ambitos de memoria")
	}
}

// agregamos una variable al ambito
func (c *Contexto) AgregarVariable(nombre string, tipo TipoD, tipoCompuesto TipoCompuesto, NumPar int, ListaPar []Symbol, Resultado *Resultado, Resultados []Resultado, linea int, columna int) (bool, string) {
	existe := c.Memoria.Exist(nombre)
	if existe {
		return false, "Error variable" + nombre + " declarada nuevamente"
	}
	return c.Memoria.CreateSimbolTipo(nombre, tipo, tipoCompuesto, NumPar, ListaPar, Resultado, Resultados, linea, columna)
}

func (c *Contexto) AgregarVector(constante bool, nombre string, tipo TipoD, tipoVector TipoD, tipoCompuesto TipoCompuesto, NumPar int, ListaPar []Symbol, Resultado *Resultado, Resultados []Resultado, linea int, columna int) (bool, string) {
	existe := c.Memoria.Exist(nombre)
	if existe {
		return false, "Error variable" + nombre + " declarada nuevamente"
	}
	return c.Memoria.CreateSimbolVector(constante, nombre, tipo, tipoVector, tipoCompuesto, NumPar, ListaPar, Resultado, Resultados, linea, columna)
}

func (c *Contexto) AgregarVariableTipo(nombre string, tipo TipoD, tipoCompuesto TipoCompuesto, NumPar int, ListaPar []Symbol, Resultado *Resultado, Resultados []Resultado, linea int, columna int) (bool, string) {
	existe := c.Memoria.Exist(nombre)
	if existe {
		return false, "Error variable" + nombre + " declarada nuevamente"
	}
	return c.Memoria.CreateSimbolTipo(nombre, tipo, tipoCompuesto, NumPar, ListaPar, Resultado, Resultados, linea, columna)
}

// agregamos metodos para agregar constantes
func (c *Contexto) AgregarConstanteExpresion(nombre string, tipo TipoD, tipoCompuesto TipoCompuesto, NumPar int, ListaPar []Symbol, Resultado *Resultado, Resultados []Resultado, linea int, columna int) (bool, string) {
	existe := c.Memoria.Exist(nombre)
	if existe {
		return false, "Error variable" + nombre + " declarada nuevamente"
	}
	return c.Memoria.CreateSimbolConstExp(nombre, tipo, tipoCompuesto, NumPar, ListaPar, Resultado, Resultados, linea, columna)
}
func (c *Contexto) AgregarConstanteVector(nombre string, tipo TipoD, tipoVector TipoD, tipoCompuesto TipoCompuesto, NumPar int, ListaPar []Symbol, Resultado *Resultado, Resultados []Resultado, linea int, columna int) (bool, string) {
	existe := c.Memoria.Exist(nombre)
	if existe {
		return false, "Error variable" + nombre + " declarada nuevamente"
	}
	return c.Memoria.CreateSimbolVector(true, nombre, tipo, tipoVector, tipoCompuesto, NumPar, ListaPar, Resultado, Resultados, linea, columna)
}

// solo agregamos valores para nativos int,float,char,string
func (c *Contexto) ReasignarValorNativo(nombre string, res *Resultado, linea int, columna int) bool {
	existe := c.Memoria.Exist(nombre)
	if !existe {
		var aux_Memoria = c.Memoria.Anterior
		for aux_Memoria != nil {
			existe := aux_Memoria.Exist(nombre)
			if existe == true {
				tmp, statRes := aux_Memoria.SetSymbolNativ(nombre, res)
				if tmp {
					return true
				} else {
					c.AddError(statRes, linea, columna)
					return false
				}
			}
			aux_Memoria = aux_Memoria.Anterior
			if aux_Memoria.Anterior == nil {
				if aux_Memoria.Exist(nombre) {
					tmp, statRes := aux_Memoria.SetSymbolNativ(nombre, res)
					if tmp {
						return true
					} else {
						c.AddError(statRes, linea, columna)
						return false
					}
				}
			}
		}
	}
	tmp, statRes := c.Memoria.SetSymbolNativ(nombre, res)
	if tmp {
		return true
	} else {
		c.AddError(statRes, linea, columna)
		return false
	}
}

func (c *Contexto) ReasignarValorNativoFor(nombre string, res *Resultado, linea int, columna int) bool {
	existe := c.Memoria.Exist(nombre)
	if !existe {
		var aux_Memoria = c.Memoria.Anterior
		for aux_Memoria != nil {
			existe := aux_Memoria.Exist(nombre)
			if existe == true {
				tmp, statRes := aux_Memoria.SetSymbolNativFor(nombre, res)
				if tmp {
					return true
				} else {
					c.AddError(statRes, linea, columna)
					return false
				}
			}
			aux_Memoria = aux_Memoria.Anterior
			if aux_Memoria.Anterior == nil {
				if aux_Memoria.Exist(nombre) {
					tmp, statRes := aux_Memoria.SetSymbolNativ(nombre, res)
					if tmp {
						return true
					} else {
						c.AddError(statRes, linea, columna)
						return false
					}
				}
			}
		}
	}
	tmp, statRes := c.Memoria.SetSymbolNativFor(nombre, res)
	if tmp {
		return true
	} else {
		c.AddError(statRes, linea, columna)
		return false
	}
}

func (c *Contexto) GetValorNativo(nombre string) (bool, *Resultado, bool) {
	existe := c.Memoria.Exist(nombre)
	if !existe {
		var aux_Memoria = c.Memoria.Anterior
		for aux_Memoria != nil {
			existe = aux_Memoria.Exist(nombre)
			if existe == true {
				return aux_Memoria.GetSimbolo(nombre)
			}
			aux_Memoria = aux_Memoria.Anterior
			if aux_Memoria.Anterior == nil {
				if aux_Memoria.Exist(nombre) {
					return aux_Memoria.GetSimbolo(nombre)
				}
			}
		}
		return false, nil, false
	}
	return c.Memoria.GetSimbolo(nombre)
}

func (c *Contexto) GetValor(nombre string) (bool, TipoD, *Resultado, []Resultado, bool) {
	existe := c.Memoria.Exist(nombre)
	if !existe {
		var aux_Memoria = c.Memoria.Anterior
		for aux_Memoria != nil {
			existe = aux_Memoria.Exist(nombre)
			if existe == true {
				tipo := c.Memoria.GetSimboloTipoD(nombre)
				if tipo != Vector {
					tmp1, tmp2, tmp3 := c.Memoria.GetSimbolo(nombre)
					return tmp1, tipo, tmp2, nil, tmp3
				} else {
					tmp1, tmpT, tmp3, tmp4 := c.Memoria.GetSimboloVector(nombre)
					return tmp1, tmpT, nil, tmp3, tmp4
				}
			}
			aux_Memoria = aux_Memoria.Anterior
			if aux_Memoria.Anterior == nil {
				if aux_Memoria.Exist(nombre) {
					tipo := c.Memoria.GetSimboloTipoD(nombre)
					if tipo != Vector {
						tmp1, tmp2, tmp3 := c.Memoria.GetSimbolo(nombre)
						return tmp1, tipo, tmp2, nil, tmp3
					} else {
						tmp1, tmpT, tmp3, tmp4 := c.Memoria.GetSimboloVector(nombre)
						return tmp1, tmpT, nil, tmp3, tmp4
					}
				}
			}
		}
		return false, Nil, nil, nil, false
	}
	tipo := c.Memoria.GetSimboloTipoD(nombre)
	if tipo != Vector {
		tmp1, tmp2, tmp3 := c.Memoria.GetSimbolo(nombre)
		return tmp1, tipo, tmp2, nil, tmp3
	} else {
		tmp1, tmpT, tmp3, tmp4 := c.Memoria.GetSimboloVector(nombre)
		return tmp1, tmpT, nil, tmp3, tmp4
	}
}

func (c *Contexto) GetValorVector(nombre string) (bool, TipoD, []Resultado, bool) {
	existe := c.Memoria.Exist(nombre)
	if !existe {
		var aux_Memoria = c.Memoria.Anterior
		for aux_Memoria != nil {
			existe = aux_Memoria.Exist(nombre)
			if existe == true {
				return aux_Memoria.GetSimboloVector(nombre)
			}
			aux_Memoria = aux_Memoria.Anterior
			if aux_Memoria.Anterior == nil {
				if aux_Memoria.Exist(nombre) {
					return aux_Memoria.GetSimboloVector(nombre)
				}
			}
		}
		return false, Nil, nil, false
	}
	return c.Memoria.GetSimboloVector(nombre)
}

func (c *Contexto) ReasignarValorVector(nombre string, res []Resultado, linea int, columna int) bool {
	existe := c.Memoria.Exist(nombre)
	if !existe {
		var aux_Memoria = c.Memoria.Anterior
		for aux_Memoria != nil {
			existe := aux_Memoria.Exist(nombre)
			if existe == true {
				tmp, statRes := aux_Memoria.SetSymbolVector(nombre, res)
				if tmp {
					return true
				} else {
					c.AddError(statRes, linea, columna)
					return false
				}
			}
			aux_Memoria = aux_Memoria.Anterior
			if aux_Memoria.Anterior == nil {
				if aux_Memoria.Exist(nombre) {
					tmp, statRes := aux_Memoria.SetSymbolVector(nombre, res)
					if tmp {
						return true
					} else {
						c.AddError(statRes, linea, columna)
						return false
					}
				}
			}
		}
	}
	tmp, statRes := c.Memoria.SetSymbolVector(nombre, res)
	if tmp {
		return true
	} else {
		c.AddError(statRes, linea, columna)
		return false
	}
}
