package interprete

import (
	"fmt"
)

type Contexto struct {
	Memoria   *Memoria
	numeroRef int
	Consola   string
	Errores   []string
	Conversor *Conversor
}

func NewContexto() *Contexto {
	c := &Contexto{
		Memoria:   NewMemoria(nil),
		Consola:   "",
		Errores:   make([]string, 0),
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

func (c *Contexto) AddError(entrada string) {
	c.Errores = append(c.Errores, entrada)
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

// solo agregamos valores para nativos int,float,char,string
func (c *Contexto) ReasignarValorNativo(nombre string, res *Resultado) bool {
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
					c.AddError(statRes)
					return false
				}
			}
			aux_Memoria = aux_Memoria.Anterior
		}
	}
	return false
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
		}
		return false, nil, false
	}
	return c.Memoria.GetSimbolo(nombre)
}
