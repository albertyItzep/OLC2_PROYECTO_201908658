package interprete

import (
	"fmt"
	"strconv"
)

type Conversor struct {
	ctx *Contexto
}

func NewConversor(ctx *Contexto) *Conversor {
	return &Conversor{
		ctx: ctx,
	}
}

func (c *Conversor) Ampliar(res *Resultado, tipo TipoD) *Resultado {
	switch tipo {
	case Nil:
		c.ctx.AddError("No se puede realizar conversiones a Nil")
	case Bool:

	case Integer:
		switch res.Tipo {
		case Nil:
			c.ctx.AddError("No se puede realizar conversiones a Nil")
			return NewNil()
		case Bool:
			c.ctx.AddError("No se puede realizar conversion de Bool a Int")
			return NewNil()
		case Integer:
			return res
		case String:
			tmp, err := strconv.Atoi(res.ValorS)
			if err != nil {
				c.ctx.AddError("Error al realizar la conversion de cadena: " + res.ValorS + " a string")
				return NewNil()
			}
			res.Valor = tmp
			res.Tipo = Integer
			res.ValorS = ""
			return res
		case Character:
			c.ctx.AddError("No se puede realizar la conversion de Character a Integer")
			return NewNil()
		case Float:
			res.Valor = int(res.ValorF)
			res.Tipo = Integer
			res.ValorF = 0.0
			return res
		}
	case String:
		switch res.Tipo {
		case Nil:
			c.ctx.AddError("No se puede realizar la conversion de Nil a String")
			return NewNil()
		case Bool:
			if res.ValorB {
				res.ValorS = "true"
				res.Tipo = String
				res.ValorB = false
				return res
			} else {
				res.ValorS = "false"
				res.Tipo = String
				res.ValorB = false
				return res
			}
		case Integer:
			res.ValorS = strconv.Itoa(res.Valor)
			res.Tipo = String
			res.Valor = 0
			return res
		case String:
			return res
		case Character:
		case Float:
			res.ValorS = fmt.Sprintf("%f", res.ValorF)
			res.Tipo = String
			res.ValorF = 0
			return res
		}
	case Character:
	case Float:
		switch res.Tipo {
		case Nil:
			c.ctx.AddError("No se puede realizar la conversion de Nil a Float")
			return NewNil()
		case Bool:
			c.ctx.AddError("No se puede realizar la conversion de bool a Float")
			return NewNil()
		case Integer:
			res.ValorF = float64(res.Valor)
			res.Tipo = Float
			res.Valor = 0
			return res
		case String:
			tmpFloat, tmp := strconv.ParseFloat(res.ValorS, 64)
			if tmp != nil {
				c.ctx.AddError("Error Parseo: No se pudo realizar el casteo de string a float")
				return NewNil()
			}
			res.ValorF = tmpFloat
			res.Tipo = Float
			res.ValorS = ""
			return res
		case Character:
			c.ctx.AddError("No se puede realizar la conversion de character a Float")
			return NewNil()
		case Float:
		}
	}
	c.ctx.AddError("Conversion ilegal" +
		tipo.String() + " y" + res.Tipo.String())
	return NewNil()
}

/*
case Nil:
	case Bool:
	case Integer:
	case String:
	case Character:
	case Float:
*/
