package main

import (
	"OLC2_PROYECTO_201908658/Background"
	interprete "OLC2_PROYECTO_201908658/Interprete"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gorilla/mux"
)

type Response struct {
	Consola  string
	Errores  []interprete.Errs
	Simbolos map[string]*interprete.Symbol
}

type Content struct {
	Content string `json:"content"`
}

func main() {
	router := mux.NewRouter()
	router.Use(corsMiddleware)
	router.HandleFunc("/", indexFunc).Methods("POST")

	fmt.Println("Servidor en ejecución en el puerto 8080...")
	http.ListenAndServe(":8000", router)
}

func indexFunc(w http.ResponseWriter, r *http.Request) {
	var content Content
	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fs := antlr.NewInputStream(content.Content)
	lexer := Background.NewControlLexer(fs)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := Background.NewControlParser(tokens)

	visitor := NewVisitorInterprete()
	arbol := parser.Prog()

	raiz := visitor.Visit(arbol).(interprete.AbstrExpr)

	ctx := interprete.NewContexto()
	raiz.Interpretar(ctx)
	fmt.Println("Ejecucion Realizada")
	fmt.Println(ctx.Memoria.GetMemoria())
	fmt.Println(ctx.Consola)

	tmp := Response{Consola: ctx.Consola, Errores: ctx.Errores, Simbolos: ctx.Memoria.GetMemoria()}
	json.NewEncoder(w).Encode(tmp)
}
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Permitir solicitudes desde cualquier origen con cualquier método y encabezado
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
