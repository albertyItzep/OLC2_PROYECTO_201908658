package main

import (
	"OLC2_PROYECTO_201908658/Background"
	interprete "OLC2_PROYECTO_201908658/Interprete"
	"fmt"
	"log"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	fmt.Println("Hola compi")
	ejecutar("entrada.txt")
}

func ejecutar(ruta string) {
	fs, err := antlr.NewFileStream(ruta)
	if err != nil {
		log.Fatal(err)
	}

	lexer := Background.NewControlLexer(fs)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := Background.NewControlParser(tokens)

	visitor := NewVisitorInterprete()
	arbol := parser.Prog()

	raiz := visitor.Visit(arbol).(interprete.AbstrExpr)

	ctx := interprete.NewContexto()

	raiz.Interpretar(ctx)

	if len(ctx.Errores) > 0 {
		for _, v := range ctx.Errores {
			fmt.Println(v)
		}
	} else {
		fmt.Println(ctx.Consola)
	}
}
