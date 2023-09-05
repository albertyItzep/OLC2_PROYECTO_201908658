package main

import (
	"OLC2_PROYECTO_201908658/Background"
	interprete "OLC2_PROYECTO_201908658/Interprete"
	noterm "OLC2_PROYECTO_201908658/Interprete/NoTerm"
	terminales "OLC2_PROYECTO_201908658/Interprete/Terminales"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type VisitorInterprete struct {
	antlr.ParseTreeVisitor
	Raiz interprete.AbstrExpr
}

func NewVisitorInterprete() Background.ControlVisitor {
	return &VisitorInterprete{
		ParseTreeVisitor: &Background.BaseControlVisitor{},
		Raiz:             nil,
	}
}

// Visit a parse tree produced by ControlParser#Dec_VTipoValor.
func (v VisitorInterprete) VisitDec_VTipoValor(ctx *Background.Dec_VTipoValorContext) interface{} {
	Id := ctx.ID().GetText()
	tipo := terminales.NewT_TipoVariable(ctx.Tipovariable().GetText())
	var expr interprete.AbstrExpr = nil
	if ctx.Expr() != nil {
		expr = ctx.Expr().Accept(v).(interprete.AbstrExpr)
	}
	variable := noterm.NewNT_Variables(Id, tipo.GetTipoD(), expr, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
	return variable
}

// Visit a parse tree produced by ControlParser#Dec_VTipo.
func (v VisitorInterprete) VisitDec_VTipo(ctx *Background.Dec_VTipoContext) interface{} {
	Id := ctx.ID().GetText()
	tipo := terminales.NewT_TipoVariable(ctx.Tipovariable().GetText())
	variable := noterm.NewNT_VariablesTipo(Id, tipo.GetTipoD(), ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())

	return variable
}

// Visit a parse tree produced by ControlParser#Dec_VVal.
func (v VisitorInterprete) VisitDec_VVal(ctx *Background.Dec_VValContext) interface{} {
	Id := ctx.ID().GetText()
	expr := ctx.Expr().Accept(v).(interprete.AbstrExpr)
	variable := noterm.NewNT_VariableExpresion(Id, expr, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
	return variable
}

// Visit a parse tree produced by ControlParser#Dec_CVal.
func (v VisitorInterprete) VisitDec_CVal(ctx *Background.Dec_CValContext) interface{} {
	Id := ctx.ID().GetText()
	expr := ctx.Expr().Accept(v).(interprete.AbstrExpr)
	variable := noterm.NewNT_DecConstanteExpresion(Id, expr, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
	return variable
}

// Visit a parse tree produced by ControlParser#Dec_CTipoValor.
func (v VisitorInterprete) VisitDec_CTipoValor(ctx *Background.Dec_CTipoValorContext) interface{} {
	Id := ctx.ID().GetText()
	tipo := terminales.NewT_TipoVariable(ctx.Tipovariable().GetText())
	var expr interprete.AbstrExpr = nil
	if ctx.Expr() != nil {
		expr = ctx.Expr().Accept(v).(interprete.AbstrExpr)
	}
	variable := noterm.NewNT_DecConsComp(Id, tipo.GetTipoD(), expr, ctx.ID().GetSymbol().GetLine(), ctx.ID().GetSymbol().GetColumn())
	return variable
}

func (v VisitorInterprete) Visit(tree antlr.ParseTree) interface{} {
	switch nodo := tree.(type) {
	case *antlr.ErrorNodeImpl:
		return noterm.NewNT_Error(nodo.GetText())
	default:
		nodoActual, resStatus := tree.Accept(v).(interprete.AbstrExpr)
		if resStatus {
			return nodoActual
		}
		return noterm.NewNT_Error("Nodo desconocido")
	}
}

/*
func (v VisitorInterprete) VisitChildren(node antlr.RuleNode) interface{} {
	panic("not implemented") // TODO: Implement
}

func (v VisitorInterprete) VisitTerminal(node antlr.TerminalNode) interface{} {
	panic("not implemented") // TODO: Implement
}
*/

func (v VisitorInterprete) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return noterm.NewNT_Error(node.GetText())
}

// Visit a parse tree produced by ControlParser#ProgInicial.
func (v VisitorInterprete) VisitProgInicial(ctx *Background.ProgInicialContext) interface{} {
	nodo := ctx.Block().Accept(v).(interprete.AbstrExpr)
	v.Raiz = nodo
	return nodo
}

// Visit a parse tree produced by ControlParser#BlockContent.
func (v VisitorInterprete) VisitBlockContent(ctx *Background.BlockContentContext) interface{} {
	blockContent := noterm.NewBlockContent()
	sentencias := ctx.AllSentencias()
	for _, sentecia := range sentencias {
		nodoSentencia := sentecia.Accept(v).(interprete.AbstrExpr)
		blockContent.AddSentencia(nodoSentencia)
	}
	return blockContent
}

// Visit a parse tree produced by ControlParser#SentenciaDeclaracion.
func (v VisitorInterprete) VisitSentenciaDeclaracion(ctx *Background.SentenciaDeclaracionContext) interface{} {
	fmt.Print("10")
	return ctx.Declaracion().Accept(v).(interprete.AbstrExpr)
}

// Visit a parse tree produced by ControlParser#SentenciaDeclaracionVector.
func (v VisitorInterprete) VisitSentenciaDeclaracionVector(ctx *Background.SentenciaDeclaracionVectorContext) interface{} {
	fmt.Print("110")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaDeclaracionMatrix.
func (v VisitorInterprete) VisitSentenciaDeclaracionMatrix(ctx *Background.SentenciaDeclaracionMatrixContext) interface{} {
	fmt.Print("120")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaDeclaracionFuncion.
func (v VisitorInterprete) VisitSentenciaDeclaracionFuncion(ctx *Background.SentenciaDeclaracionFuncionContext) interface{} {
	fmt.Print("130")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaAsignacion.
func (v VisitorInterprete) VisitSentenciaAsignacion(ctx *Background.SentenciaAsignacionContext) interface{} {
	return ctx.Asignacion().Accept(v).(interprete.AbstrExpr)
}

// Visit a parse tree produced by ControlParser#SentenciaIf.
func (v VisitorInterprete) VisitSentenciaIf(ctx *Background.SentenciaIfContext) interface{} {
	fmt.Print("150")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaSwitch.
func (v VisitorInterprete) VisitSentenciaSwitch(ctx *Background.SentenciaSwitchContext) interface{} {
	fmt.Print("160")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaWhile.
func (v VisitorInterprete) VisitSentenciaWhile(ctx *Background.SentenciaWhileContext) interface{} {
	fmt.Print("170")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaFor.
func (v VisitorInterprete) VisitSentenciaFor(ctx *Background.SentenciaForContext) interface{} {
	fmt.Print("180")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaGuard.
func (v VisitorInterprete) VisitSentenciaGuard(ctx *Background.SentenciaGuardContext) interface{} {
	fmt.Print("190")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaBreak.
func (v VisitorInterprete) VisitSentenciaBreak(ctx *Background.SentenciaBreakContext) interface{} {
	fmt.Print("121")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaContinue.
func (v VisitorInterprete) VisitSentenciaContinue(ctx *Background.SentenciaContinueContext) interface{} {
	fmt.Print("122")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaReturn.
func (v VisitorInterprete) VisitSentenciaReturn(ctx *Background.SentenciaReturnContext) interface{} {
	fmt.Print("123")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaFuncVectoriales.
func (v VisitorInterprete) VisitSentenciaFuncVectoriales(ctx *Background.SentenciaFuncVectorialesContext) interface{} {
	fmt.Print("124")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaLLamadaFuncion.
func (v VisitorInterprete) VisitSentenciaLLamadaFuncion(ctx *Background.SentenciaLLamadaFuncionContext) interface{} {
	fmt.Print("125")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaPrint.
func (v VisitorInterprete) VisitSentenciaPrint(ctx *Background.SentenciaPrintContext) interface{} {
	fmt.Println("126")
	return ctx.InstPrint().Accept(v).(interprete.AbstrExpr)
}

// Visit a parse tree produced by ControlParser#SentenciaStruct.
func (v VisitorInterprete) VisitSentenciaStruct(ctx *Background.SentenciaStructContext) interface{} {
	fmt.Print("127")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaDeclaracionStruct.
func (v VisitorInterprete) VisitSentenciaDeclaracionStruct(ctx *Background.SentenciaDeclaracionStructContext) interface{} {
	fmt.Print("128")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#SentenciaLLamadaMetodos.
func (v VisitorInterprete) VisitSentenciaLLamadaMetodos(ctx *Background.SentenciaLLamadaMetodosContext) interface{} {
	fmt.Print("129")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#insIf.
func (v VisitorInterprete) VisitInsIf(ctx *Background.InsIfContext) interface{} {
	fmt.Print("13")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#instElseIf.
func (v VisitorInterprete) VisitInstElseIf(ctx *Background.InstElseIfContext) interface{} {
	fmt.Print("131")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#insElse.
func (v VisitorInterprete) VisitInsElse(ctx *Background.InsElseContext) interface{} {
	fmt.Print("132")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#insSwitch.
func (v VisitorInterprete) VisitInsSwitch(ctx *Background.InsSwitchContext) interface{} {
	fmt.Print("133")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#instCase.
func (v VisitorInterprete) VisitInstCase(ctx *Background.InstCaseContext) interface{} {
	fmt.Print("134")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#instDefault.
func (v VisitorInterprete) VisitInstDefault(ctx *Background.InstDefaultContext) interface{} {
	fmt.Print("135")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#instWhile.
func (v VisitorInterprete) VisitInstWhile(ctx *Background.InstWhileContext) interface{} {
	fmt.Print("136")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#instFor.
func (v VisitorInterprete) VisitInstFor(ctx *Background.InstForContext) interface{} {
	fmt.Print("137")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#instGuard.
func (v VisitorInterprete) VisitInstGuard(ctx *Background.InstGuardContext) interface{} {
	fmt.Print("138")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#instBreak.
func (v VisitorInterprete) VisitInstBreak(ctx *Background.InstBreakContext) interface{} {
	fmt.Print("139")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#instContinue.
func (v VisitorInterprete) VisitInstContinue(ctx *Background.InstContinueContext) interface{} {
	fmt.Print("14")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#instReturn.
func (v VisitorInterprete) VisitInstReturn(ctx *Background.InstReturnContext) interface{} {
	fmt.Print("141")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#decVector.
func (v VisitorInterprete) VisitDecVector(ctx *Background.DecVectorContext) interface{} {
	fmt.Print("142")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#defVector.
func (v VisitorInterprete) VisitDefVector(ctx *Background.DefVectorContext) interface{} {
	fmt.Print("143")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#funcsVectoriales.
func (v VisitorInterprete) VisitFuncsVectoriales(ctx *Background.FuncsVectorialesContext) interface{} {
	fmt.Print("144")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#decMatriz.
func (v VisitorInterprete) VisitDecMatriz(ctx *Background.DecMatrizContext) interface{} {
	fmt.Print("145")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#tipoMatriz.
func (v VisitorInterprete) VisitTipoMatriz(ctx *Background.TipoMatrizContext) interface{} {
	fmt.Print("146")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#defMatrix.
func (v VisitorInterprete) VisitDefMatrix(ctx *Background.DefMatrixContext) interface{} {
	fmt.Print("147")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#listaValoresMat.
func (v VisitorInterprete) VisitListaValoresMat(ctx *Background.ListaValoresMatContext) interface{} {
	fmt.Print("148")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#listaValoresMat2.
func (v VisitorInterprete) VisitListaValoresMat2(ctx *Background.ListaValoresMat2Context) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#simpleMax.
func (v VisitorInterprete) VisitSimpleMax(ctx *Background.SimpleMaxContext) interface{} {
	fmt.Print("149")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#instStruct.
func (v VisitorInterprete) VisitInstStruct(ctx *Background.InstStructContext) interface{} {
	fmt.Print("15")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#listaAtributos.
func (v VisitorInterprete) VisitListaAtributos(ctx *Background.ListaAtributosContext) interface{} {
	fmt.Print("151")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#decStruct.
func (v VisitorInterprete) VisitDecStruct(ctx *Background.DecStructContext) interface{} {
	fmt.Print("152")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#lDupla.
func (v VisitorInterprete) VisitLDupla(ctx *Background.LDuplaContext) interface{} {
	fmt.Print("153")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#llamAtributos.
func (v VisitorInterprete) VisitLlamAtributos(ctx *Background.LlamAtributosContext) interface{} {
	fmt.Print("154")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#llamadaMetodos.
func (v VisitorInterprete) VisitLlamadaMetodos(ctx *Background.LlamadaMetodosContext) interface{} {
	fmt.Print("155")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#decFunciones.
func (v VisitorInterprete) VisitDecFunciones(ctx *Background.DecFuncionesContext) interface{} {
	fmt.Print("156")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#listaParametros.
func (v VisitorInterprete) VisitListaParametros(ctx *Background.ListaParametrosContext) interface{} {
	fmt.Print("157")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#llamadaFuncion.
func (v VisitorInterprete) VisitLlamadaFuncion(ctx *Background.LlamadaFuncionContext) interface{} {
	fmt.Print("158")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#listaLlamadaParametros.
func (v VisitorInterprete) VisitListaLlamadaParametros(ctx *Background.ListaLlamadaParametrosContext) interface{} {
	fmt.Print("159")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#IntruccionPrint.
func (v VisitorInterprete) VisitIntruccionPrint(ctx *Background.IntruccionPrintContext) interface{} {
	fmt.Println("16")
	fprint := noterm.NewNT_Print()
	expresiones := ctx.AllExpr()
	for _, expr := range expresiones {
		nodoExpr := expr.Accept(v).(interprete.AbstrExpr)
		fprint.AddExpresion(nodoExpr)
	}
	return fprint
}

// Visit a parse tree produced by ControlParser#instCasteos.
func (v VisitorInterprete) VisitInstCasteos(ctx *Background.InstCasteosContext) interface{} {
	fmt.Print("161")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#numList.
func (v VisitorInterprete) VisitNumList(ctx *Background.NumListContext) interface{} {
	fmt.Print("162")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#expresionList.
func (v VisitorInterprete) VisitExpresionList(ctx *Background.ExpresionListContext) interface{} {
	fmt.Print("163")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#objectsList.
func (v VisitorInterprete) VisitObjectsList(ctx *Background.ObjectsListContext) interface{} {
	fmt.Print("164")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#declaracion.
func (v VisitorInterprete) VisitDeclaracion(ctx *Background.DeclaracionContext) interface{} {
	fmt.Print("165")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#asignacion.
func (v VisitorInterprete) VisitAsignacion(ctx *Background.AsignacionContext) interface{} {
	fmt.Print("166")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_Decimal.
func (v VisitorInterprete) VisitExpr_Decimal(ctx *Background.Expr_DecimalContext) interface{} {
	return terminales.NewT_FloatNum(ctx.DECIMAL().GetText(),
		ctx.DECIMAL().GetSymbol().GetLine(),
		ctx.DECIMAL().GetSymbol().GetColumn())
}

// Visit a parse tree produced by ControlParser#Expr_InstCasteo.
func (v VisitorInterprete) VisitExpr_InstCasteo(ctx *Background.Expr_InstCasteoContext) interface{} {
	fmt.Print("167")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_ValidaIgualDif.
func (v VisitorInterprete) VisitExpr_ValidaIgualDif(ctx *Background.Expr_ValidaIgualDifContext) interface{} {
	expIzq := ctx.Expr(0).Accept(v).(interprete.AbstrExpr)
	expDer := ctx.Expr(1).Accept(v).(interprete.AbstrExpr)

	if ctx.GetOp().GetText() == "==" {
		return noterm.NewNT_CompIgualdad(expIzq, expDer)
	} else {
		return noterm.NewNT_CompDif(expIzq, expDer)
	}
}

// Visit a parse tree produced by ControlParser#Expr_OpMulDiv.
func (v VisitorInterprete) VisitExpr_OpMulDiv(ctx *Background.Expr_OpMulDivContext) interface{} {
	expIzq := ctx.Expr(0).Accept(v).(interprete.AbstrExpr)
	expDer := ctx.Expr(1).Accept(v).(interprete.AbstrExpr)
	if ctx.GetOp().GetText() == "*" {
		return noterm.NewNT_ExpMult(expIzq, expDer)
	} else {
		return noterm.NewNT_ExpDiv(expIzq, expDer)
	}
}

// Visit a parse tree produced by ControlParser#Expr_ValidaAnd.
func (v VisitorInterprete) VisitExpr_ValidaAnd(ctx *Background.Expr_ValidaAndContext) interface{} {
	expIzq := ctx.Expr(0).Accept(v).(interprete.AbstrExpr)
	expDer := ctx.Expr(1).Accept(v).(interprete.AbstrExpr)
	return noterm.NewNT_LogAnd(expIzq, expDer)
}

// Visit a parse tree produced by ControlParser#Expr_LlamAtributos.
func (v VisitorInterprete) VisitExpr_LlamAtributos(ctx *Background.Expr_LlamAtributosContext) interface{} {
	fmt.Print("171")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_Nil.
func (v VisitorInterprete) VisitExpr_Nil(ctx *Background.Expr_NilContext) interface{} {
	fmt.Print("172")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_PosMatrix.
func (v VisitorInterprete) VisitExpr_PosMatrix(ctx *Background.Expr_PosMatrixContext) interface{} {
	fmt.Print("173")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_Boolean.
func (v VisitorInterprete) VisitExpr_Boolean(ctx *Background.Expr_BooleanContext) interface{} {
	fmt.Print("174")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_InstRango.
func (v VisitorInterprete) VisitExpr_InstRango(ctx *Background.Expr_InstRangoContext) interface{} {
	fmt.Print("175")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_ValNumNeg.
func (v VisitorInterprete) VisitExpr_ValNumNeg(ctx *Background.Expr_ValNumNegContext) interface{} {
	expr := ctx.Expr().Accept(v).(interprete.AbstrExpr)
	return terminales.NewT_NegExp(expr)
}

// Visit a parse tree produced by ControlParser#Expr_ValidaOr.
func (v VisitorInterprete) VisitExpr_ValidaOr(ctx *Background.Expr_ValidaOrContext) interface{} {
	expIzq := ctx.Expr(0).Accept(v).(interprete.AbstrExpr)
	expDer := ctx.Expr(1).Accept(v).(interprete.AbstrExpr)
	return noterm.NewNT_LogOR(expIzq, expDer)
}

// Visit a parse tree produced by ControlParser#Expr_ValNegacion.
func (v VisitorInterprete) VisitExpr_ValNegacion(ctx *Background.Expr_ValNegacionContext) interface{} {
	exp := ctx.Expr().Accept(v).(interprete.AbstrExpr)
	return noterm.NewNT_LogNot(exp)
}

// Visit a parse tree produced by ControlParser#Expr_ID.
func (v VisitorInterprete) VisitExpr_ID(ctx *Background.Expr_IDContext) interface{} {
	return terminales.NewT_IdExp(
		ctx.GetText(),
		ctx.ID().GetSymbol().GetLine(),
		ctx.ID().GetSymbol().GetColumn(),
	) // me va servir para llamar a las variables
}

// Visit a parse tree produced by ControlParser#Expr_ValidaMayQue.
func (v VisitorInterprete) VisitExpr_ValidaMayQue(ctx *Background.Expr_ValidaMayQueContext) interface{} {
	expIzq := ctx.Expr(0).Accept(v).(interprete.AbstrExpr)
	expDer := ctx.Expr(1).Accept(v).(interprete.AbstrExpr)

	if ctx.GetOp().GetText() == ">=" {
		return noterm.NewNT_RelMayIgual(expIzq, expDer)
	} else {
		return noterm.NewNT_RelMayor(expIzq, expDer)
	}
}

// Visit a parse tree produced by ControlParser#Expr_Conteo.
func (v VisitorInterprete) VisitExpr_Conteo(ctx *Background.Expr_ConteoContext) interface{} {
	fmt.Print("181")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Expr_OpSumRes.
func (v VisitorInterprete) VisitExpr_OpSumRes(ctx *Background.Expr_OpSumResContext) interface{} {
	expIzq := ctx.Expr(0).Accept(v).(interprete.AbstrExpr)
	expDer := ctx.Expr(1).Accept(v).(interprete.AbstrExpr)

	if ctx.GetOp().GetText() == "+" {
		return noterm.NewNT_ExpSum(expIzq, expDer)
	} else if ctx.GetOp().GetText() == "-" {
		return noterm.NewNT_ExpRes(expIzq, expDer)
	} else {
		return noterm.NewNT_ExpModulo(expIzq, expDer)
	}
}

// Visit a parse tree produced by ControlParser#Expr_Entero.
func (v VisitorInterprete) VisitExpr_Entero(ctx *Background.Expr_EnteroContext) interface{} {
	return terminales.NewT_IntNum(ctx.INT().GetText(),
		ctx.INT().GetSymbol().GetLine(),
		ctx.INT().GetSymbol().GetColumn())
}

// Visit a parse tree produced by ControlParser#Expr_ParExpre.
func (v VisitorInterprete) VisitExpr_ParExpre(ctx *Background.Expr_ParExpreContext) interface{} {
	return ctx.Expr().Accept(v).(interprete.AbstrExpr)
}

// Visit a parse tree produced by ControlParser#Expr_StringChar.
func (v VisitorInterprete) VisitExpr_StringChar(ctx *Background.Expr_StringCharContext) interface{} {
	return terminales.NewT_StringChar(
		ctx.STRING().GetText(),
		ctx.STRING().GetSymbol().GetLine(),
		ctx.STRING().GetSymbol().GetColumn(),
	)
}

// Visit a parse tree produced by ControlParser#Expr_ValidaMenQue.
func (v VisitorInterprete) VisitExpr_ValidaMenQue(ctx *Background.Expr_ValidaMenQueContext) interface{} {
	expIzq := ctx.Expr(0).Accept(v).(interprete.AbstrExpr)
	expDer := ctx.Expr(1).Accept(v).(interprete.AbstrExpr)

	if ctx.GetOp().GetText() == "<=" {
		return noterm.NewNT_RelMenIgual(expIzq, expDer)
	} else {
		return noterm.NewNT_RelMenor(expIzq, expDer)
	}
}

// Visit a parse tree produced by ControlParser#Expr_PosVector.
func (v VisitorInterprete) VisitExpr_PosVector(ctx *Background.Expr_PosVectorContext) interface{} {
	fmt.Print("185")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#tipovariable.
func (v VisitorInterprete) VisitTipovariable(ctx *Background.TipovariableContext) interface{} {
	fmt.Print("186")
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_Aumento.
func (v VisitorInterprete) VisitAsignacion_Aumento(ctx *Background.Asignacion_AumentoContext) interface{} {
	id := ctx.ID().GetText()
	expr := ctx.Expr().Accept(v).(interprete.AbstrExpr)

	return noterm.NewNT_AsIncremento(id, expr)
}

// Visit a parse tree produced by ControlParser#Asignacion_Decremento.
func (v VisitorInterprete) VisitAsignacion_Decremento(ctx *Background.Asignacion_DecrementoContext) interface{} {
	id := ctx.ID().GetText()
	expr := ctx.Expr().Accept(v).(interprete.AbstrExpr)

	return noterm.NewNT_AsDecremento(id, expr)
}

// Visit a parse tree produced by ControlParser#Asignacion_ValorGen.
func (v VisitorInterprete) VisitAsignacion_ValorGen(ctx *Background.Asignacion_ValorGenContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_VectorAumento.
func (v VisitorInterprete) VisitAsignacion_VectorAumento(ctx *Background.Asignacion_VectorAumentoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_VectorDecremento.
func (v VisitorInterprete) VisitAsignacion_VectorDecremento(ctx *Background.Asignacion_VectorDecrementoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_MatrixAumento.
func (v VisitorInterprete) VisitAsignacion_MatrixAumento(ctx *Background.Asignacion_MatrixAumentoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_MatrixDecremento.
func (v VisitorInterprete) VisitAsignacion_MatrixDecremento(ctx *Background.Asignacion_MatrixDecrementoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_VectorGeneral.
func (v VisitorInterprete) VisitAsignacion_VectorGeneral(ctx *Background.Asignacion_VectorGeneralContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_MatrixGeneral.
func (v VisitorInterprete) VisitAsignacion_MatrixGeneral(ctx *Background.Asignacion_MatrixGeneralContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_LlAtribGeneral.
func (v VisitorInterprete) VisitAsignacion_LlAtribGeneral(ctx *Background.Asignacion_LlAtribGeneralContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_LlAtribAumento.
func (v VisitorInterprete) VisitAsignacion_LlAtribAumento(ctx *Background.Asignacion_LlAtribAumentoContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by ControlParser#Asignacion_LlAtribDecremento.
func (v VisitorInterprete) VisitAsignacion_LlAtribDecremento(ctx *Background.Asignacion_LlAtribDecrementoContext) interface{} {
	panic("not implemented") // TODO: Implement
}