// Code generated from Control.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Background // Control
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ControlParser.
type ControlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ControlParser#ProgInicial.
	VisitProgInicial(ctx *ProgInicialContext) interface{}

	// Visit a parse tree produced by ControlParser#BlockContent.
	VisitBlockContent(ctx *BlockContentContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaDeclaracion.
	VisitSentenciaDeclaracion(ctx *SentenciaDeclaracionContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaDeclaracionVector.
	VisitSentenciaDeclaracionVector(ctx *SentenciaDeclaracionVectorContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaDeclaracionMatrix.
	VisitSentenciaDeclaracionMatrix(ctx *SentenciaDeclaracionMatrixContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaDeclaracionFuncion.
	VisitSentenciaDeclaracionFuncion(ctx *SentenciaDeclaracionFuncionContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaAsignacion.
	VisitSentenciaAsignacion(ctx *SentenciaAsignacionContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaIf.
	VisitSentenciaIf(ctx *SentenciaIfContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaSwitch.
	VisitSentenciaSwitch(ctx *SentenciaSwitchContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaWhile.
	VisitSentenciaWhile(ctx *SentenciaWhileContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaFor.
	VisitSentenciaFor(ctx *SentenciaForContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaGuard.
	VisitSentenciaGuard(ctx *SentenciaGuardContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaBreak.
	VisitSentenciaBreak(ctx *SentenciaBreakContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaContinue.
	VisitSentenciaContinue(ctx *SentenciaContinueContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaReturn.
	VisitSentenciaReturn(ctx *SentenciaReturnContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaFuncVectoriales.
	VisitSentenciaFuncVectoriales(ctx *SentenciaFuncVectorialesContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaLLamadaFuncion.
	VisitSentenciaLLamadaFuncion(ctx *SentenciaLLamadaFuncionContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaPrint.
	VisitSentenciaPrint(ctx *SentenciaPrintContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaStruct.
	VisitSentenciaStruct(ctx *SentenciaStructContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaDeclaracionStruct.
	VisitSentenciaDeclaracionStruct(ctx *SentenciaDeclaracionStructContext) interface{}

	// Visit a parse tree produced by ControlParser#SentenciaLLamadaMetodos.
	VisitSentenciaLLamadaMetodos(ctx *SentenciaLLamadaMetodosContext) interface{}

	// Visit a parse tree produced by ControlParser#insIf.
	VisitInsIf(ctx *InsIfContext) interface{}

	// Visit a parse tree produced by ControlParser#instElseIf.
	VisitInstElseIf(ctx *InstElseIfContext) interface{}

	// Visit a parse tree produced by ControlParser#insElse.
	VisitInsElse(ctx *InsElseContext) interface{}

	// Visit a parse tree produced by ControlParser#insSwitch.
	VisitInsSwitch(ctx *InsSwitchContext) interface{}

	// Visit a parse tree produced by ControlParser#instCase.
	VisitInstCase(ctx *InstCaseContext) interface{}

	// Visit a parse tree produced by ControlParser#instDefault.
	VisitInstDefault(ctx *InstDefaultContext) interface{}

	// Visit a parse tree produced by ControlParser#instWhile.
	VisitInstWhile(ctx *InstWhileContext) interface{}

	// Visit a parse tree produced by ControlParser#instFor.
	VisitInstFor(ctx *InstForContext) interface{}

	// Visit a parse tree produced by ControlParser#instGuard.
	VisitInstGuard(ctx *InstGuardContext) interface{}

	// Visit a parse tree produced by ControlParser#instBreak.
	VisitInstBreak(ctx *InstBreakContext) interface{}

	// Visit a parse tree produced by ControlParser#instContinue.
	VisitInstContinue(ctx *InstContinueContext) interface{}

	// Visit a parse tree produced by ControlParser#instReturn.
	VisitInstReturn(ctx *InstReturnContext) interface{}

	// Visit a parse tree produced by ControlParser#decVector.
	VisitDecVector(ctx *DecVectorContext) interface{}

	// Visit a parse tree produced by ControlParser#defVector.
	VisitDefVector(ctx *DefVectorContext) interface{}

	// Visit a parse tree produced by ControlParser#funcsVectoriales.
	VisitFuncsVectoriales(ctx *FuncsVectorialesContext) interface{}

	// Visit a parse tree produced by ControlParser#decMatriz.
	VisitDecMatriz(ctx *DecMatrizContext) interface{}

	// Visit a parse tree produced by ControlParser#tipoMatriz.
	VisitTipoMatriz(ctx *TipoMatrizContext) interface{}

	// Visit a parse tree produced by ControlParser#defMatrix.
	VisitDefMatrix(ctx *DefMatrixContext) interface{}

	// Visit a parse tree produced by ControlParser#listaValoresMat.
	VisitListaValoresMat(ctx *ListaValoresMatContext) interface{}

	// Visit a parse tree produced by ControlParser#listaValoresMat2.
	VisitListaValoresMat2(ctx *ListaValoresMat2Context) interface{}

	// Visit a parse tree produced by ControlParser#simpleMax.
	VisitSimpleMax(ctx *SimpleMaxContext) interface{}

	// Visit a parse tree produced by ControlParser#instStruct.
	VisitInstStruct(ctx *InstStructContext) interface{}

	// Visit a parse tree produced by ControlParser#listaAtributos.
	VisitListaAtributos(ctx *ListaAtributosContext) interface{}

	// Visit a parse tree produced by ControlParser#decStruct.
	VisitDecStruct(ctx *DecStructContext) interface{}

	// Visit a parse tree produced by ControlParser#lDupla.
	VisitLDupla(ctx *LDuplaContext) interface{}

	// Visit a parse tree produced by ControlParser#llamAtributos.
	VisitLlamAtributos(ctx *LlamAtributosContext) interface{}

	// Visit a parse tree produced by ControlParser#llamadaMetodos.
	VisitLlamadaMetodos(ctx *LlamadaMetodosContext) interface{}

	// Visit a parse tree produced by ControlParser#decFunciones.
	VisitDecFunciones(ctx *DecFuncionesContext) interface{}

	// Visit a parse tree produced by ControlParser#listaParametros.
	VisitListaParametros(ctx *ListaParametrosContext) interface{}

	// Visit a parse tree produced by ControlParser#llamadaFuncion.
	VisitLlamadaFuncion(ctx *LlamadaFuncionContext) interface{}

	// Visit a parse tree produced by ControlParser#listaLlamadaParametros.
	VisitListaLlamadaParametros(ctx *ListaLlamadaParametrosContext) interface{}

	// Visit a parse tree produced by ControlParser#IntruccionPrint.
	VisitIntruccionPrint(ctx *IntruccionPrintContext) interface{}

	// Visit a parse tree produced by ControlParser#instCasteos.
	VisitInstCasteos(ctx *InstCasteosContext) interface{}

	// Visit a parse tree produced by ControlParser#numList.
	VisitNumList(ctx *NumListContext) interface{}

	// Visit a parse tree produced by ControlParser#expresionList.
	VisitExpresionList(ctx *ExpresionListContext) interface{}

	// Visit a parse tree produced by ControlParser#objectsList.
	VisitObjectsList(ctx *ObjectsListContext) interface{}

	// Visit a parse tree produced by ControlParser#Dec_VTipoValor.
	VisitDec_VTipoValor(ctx *Dec_VTipoValorContext) interface{}

	// Visit a parse tree produced by ControlParser#Dec_VTipo.
	VisitDec_VTipo(ctx *Dec_VTipoContext) interface{}

	// Visit a parse tree produced by ControlParser#Dec_VVal.
	VisitDec_VVal(ctx *Dec_VValContext) interface{}

	// Visit a parse tree produced by ControlParser#Dec_CVal.
	VisitDec_CVal(ctx *Dec_CValContext) interface{}

	// Visit a parse tree produced by ControlParser#Dec_CTipoValor.
	VisitDec_CTipoValor(ctx *Dec_CTipoValorContext) interface{}

	// Visit a parse tree produced by ControlParser#Asignacion_Aumento.
	VisitAsignacion_Aumento(ctx *Asignacion_AumentoContext) interface{}

	// Visit a parse tree produced by ControlParser#Asignacion_Decremento.
	VisitAsignacion_Decremento(ctx *Asignacion_DecrementoContext) interface{}

	// Visit a parse tree produced by ControlParser#Asignacion_ValorGen.
	VisitAsignacion_ValorGen(ctx *Asignacion_ValorGenContext) interface{}

	// Visit a parse tree produced by ControlParser#Asignacion_VectorAumento.
	VisitAsignacion_VectorAumento(ctx *Asignacion_VectorAumentoContext) interface{}

	// Visit a parse tree produced by ControlParser#Asignacion_VectorDecremento.
	VisitAsignacion_VectorDecremento(ctx *Asignacion_VectorDecrementoContext) interface{}

	// Visit a parse tree produced by ControlParser#Asignacion_MatrixAumento.
	VisitAsignacion_MatrixAumento(ctx *Asignacion_MatrixAumentoContext) interface{}

	// Visit a parse tree produced by ControlParser#Asignacion_MatrixDecremento.
	VisitAsignacion_MatrixDecremento(ctx *Asignacion_MatrixDecrementoContext) interface{}

	// Visit a parse tree produced by ControlParser#Asignacion_VectorGeneral.
	VisitAsignacion_VectorGeneral(ctx *Asignacion_VectorGeneralContext) interface{}

	// Visit a parse tree produced by ControlParser#Asignacion_MatrixGeneral.
	VisitAsignacion_MatrixGeneral(ctx *Asignacion_MatrixGeneralContext) interface{}

	// Visit a parse tree produced by ControlParser#Asignacion_LlAtribGeneral.
	VisitAsignacion_LlAtribGeneral(ctx *Asignacion_LlAtribGeneralContext) interface{}

	// Visit a parse tree produced by ControlParser#Asignacion_LlAtribAumento.
	VisitAsignacion_LlAtribAumento(ctx *Asignacion_LlAtribAumentoContext) interface{}

	// Visit a parse tree produced by ControlParser#Asignacion_LlAtribDecremento.
	VisitAsignacion_LlAtribDecremento(ctx *Asignacion_LlAtribDecrementoContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_Decimal.
	VisitExpr_Decimal(ctx *Expr_DecimalContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_InstCasteo.
	VisitExpr_InstCasteo(ctx *Expr_InstCasteoContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_ValidaIgualDif.
	VisitExpr_ValidaIgualDif(ctx *Expr_ValidaIgualDifContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_OpMulDiv.
	VisitExpr_OpMulDiv(ctx *Expr_OpMulDivContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_ValidaAnd.
	VisitExpr_ValidaAnd(ctx *Expr_ValidaAndContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_LlamAtributos.
	VisitExpr_LlamAtributos(ctx *Expr_LlamAtributosContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_Nil.
	VisitExpr_Nil(ctx *Expr_NilContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_PosMatrix.
	VisitExpr_PosMatrix(ctx *Expr_PosMatrixContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_Boolean.
	VisitExpr_Boolean(ctx *Expr_BooleanContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_InstRango.
	VisitExpr_InstRango(ctx *Expr_InstRangoContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_ValNumNeg.
	VisitExpr_ValNumNeg(ctx *Expr_ValNumNegContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_ValidaOr.
	VisitExpr_ValidaOr(ctx *Expr_ValidaOrContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_ValNegacion.
	VisitExpr_ValNegacion(ctx *Expr_ValNegacionContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_ID.
	VisitExpr_ID(ctx *Expr_IDContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_ValidaMayQue.
	VisitExpr_ValidaMayQue(ctx *Expr_ValidaMayQueContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_Conteo.
	VisitExpr_Conteo(ctx *Expr_ConteoContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_OpSumRes.
	VisitExpr_OpSumRes(ctx *Expr_OpSumResContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_Entero.
	VisitExpr_Entero(ctx *Expr_EnteroContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_ParExpre.
	VisitExpr_ParExpre(ctx *Expr_ParExpreContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_StringChar.
	VisitExpr_StringChar(ctx *Expr_StringCharContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_ValidaMenQue.
	VisitExpr_ValidaMenQue(ctx *Expr_ValidaMenQueContext) interface{}

	// Visit a parse tree produced by ControlParser#Expr_PosVector.
	VisitExpr_PosVector(ctx *Expr_PosVectorContext) interface{}

	// Visit a parse tree produced by ControlParser#tipovariable.
	VisitTipovariable(ctx *TipovariableContext) interface{}
}
