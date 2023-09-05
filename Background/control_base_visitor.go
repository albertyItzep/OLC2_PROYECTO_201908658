// Code generated from Control.g4 by ANTLR 4.13.0. DO NOT EDIT.

package Background // Control
import "github.com/antlr4-go/antlr/v4"

type BaseControlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseControlVisitor) VisitProgInicial(ctx *ProgInicialContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitBlockContent(ctx *BlockContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaDeclaracion(ctx *SentenciaDeclaracionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaDeclaracionVector(ctx *SentenciaDeclaracionVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaDeclaracionMatrix(ctx *SentenciaDeclaracionMatrixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaDeclaracionFuncion(ctx *SentenciaDeclaracionFuncionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaAsignacion(ctx *SentenciaAsignacionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaIf(ctx *SentenciaIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaSwitch(ctx *SentenciaSwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaWhile(ctx *SentenciaWhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaFor(ctx *SentenciaForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaGuard(ctx *SentenciaGuardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaBreak(ctx *SentenciaBreakContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaContinue(ctx *SentenciaContinueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaReturn(ctx *SentenciaReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaFuncVectoriales(ctx *SentenciaFuncVectorialesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaLLamadaFuncion(ctx *SentenciaLLamadaFuncionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaPrint(ctx *SentenciaPrintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaStruct(ctx *SentenciaStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaDeclaracionStruct(ctx *SentenciaDeclaracionStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSentenciaLLamadaMetodos(ctx *SentenciaLLamadaMetodosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitInsIf(ctx *InsIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitInstElseIf(ctx *InstElseIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitInsElse(ctx *InsElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitInsSwitch(ctx *InsSwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitInstCase(ctx *InstCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitInstDefault(ctx *InstDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitInstWhile(ctx *InstWhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitInstFor(ctx *InstForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitInstGuard(ctx *InstGuardContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitInstBreak(ctx *InstBreakContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitInstContinue(ctx *InstContinueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitInstReturn(ctx *InstReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitDecVector(ctx *DecVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitDefVector(ctx *DefVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitFuncsVectoriales(ctx *FuncsVectorialesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitDecMatriz(ctx *DecMatrizContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitTipoMatriz(ctx *TipoMatrizContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitDefMatrix(ctx *DefMatrixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitListaValoresMat(ctx *ListaValoresMatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitListaValoresMat2(ctx *ListaValoresMat2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitSimpleMax(ctx *SimpleMaxContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitInstStruct(ctx *InstStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitListaAtributos(ctx *ListaAtributosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitDecStruct(ctx *DecStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitLDupla(ctx *LDuplaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitLlamAtributos(ctx *LlamAtributosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitLlamadaMetodos(ctx *LlamadaMetodosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitDecFunciones(ctx *DecFuncionesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitListaParametros(ctx *ListaParametrosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitLlamadaFuncion(ctx *LlamadaFuncionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitListaLlamadaParametros(ctx *ListaLlamadaParametrosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitIntruccionPrint(ctx *IntruccionPrintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitInstCasteos(ctx *InstCasteosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitNumList(ctx *NumListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpresionList(ctx *ExpresionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitObjectsList(ctx *ObjectsListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitDec_VTipoValor(ctx *Dec_VTipoValorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitDec_VTipo(ctx *Dec_VTipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitDec_VVal(ctx *Dec_VValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitDec_CVal(ctx *Dec_CValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitDec_CTipoValor(ctx *Dec_CTipoValorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAsignacion_Aumento(ctx *Asignacion_AumentoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAsignacion_Decremento(ctx *Asignacion_DecrementoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAsignacion_ValorGen(ctx *Asignacion_ValorGenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAsignacion_VectorAumento(ctx *Asignacion_VectorAumentoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAsignacion_VectorDecremento(ctx *Asignacion_VectorDecrementoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAsignacion_MatrixAumento(ctx *Asignacion_MatrixAumentoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAsignacion_MatrixDecremento(ctx *Asignacion_MatrixDecrementoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAsignacion_VectorGeneral(ctx *Asignacion_VectorGeneralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAsignacion_MatrixGeneral(ctx *Asignacion_MatrixGeneralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAsignacion_LlAtribGeneral(ctx *Asignacion_LlAtribGeneralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAsignacion_LlAtribAumento(ctx *Asignacion_LlAtribAumentoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAsignacion_LlAtribDecremento(ctx *Asignacion_LlAtribDecrementoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_Decimal(ctx *Expr_DecimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_InstCasteo(ctx *Expr_InstCasteoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_ValidaIgualDif(ctx *Expr_ValidaIgualDifContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_OpMulDiv(ctx *Expr_OpMulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_ValidaAnd(ctx *Expr_ValidaAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_LlamAtributos(ctx *Expr_LlamAtributosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_Nil(ctx *Expr_NilContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_PosMatrix(ctx *Expr_PosMatrixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_Boolean(ctx *Expr_BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_InstRango(ctx *Expr_InstRangoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_ValNumNeg(ctx *Expr_ValNumNegContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_ValidaOr(ctx *Expr_ValidaOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_ValNegacion(ctx *Expr_ValNegacionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_ID(ctx *Expr_IDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_ValidaMayQue(ctx *Expr_ValidaMayQueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_Conteo(ctx *Expr_ConteoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_OpSumRes(ctx *Expr_OpSumResContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_Entero(ctx *Expr_EnteroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_ParExpre(ctx *Expr_ParExpreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_StringChar(ctx *Expr_StringCharContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_ValidaMenQue(ctx *Expr_ValidaMenQueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitExpr_PosVector(ctx *Expr_PosVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitTipovariable(ctx *TipovariableContext) interface{} {
	return v.VisitChildren(ctx)
}
