// Generated from Control.g4 by ANTLR 4.13.0
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link ControlParser}.
 */
public interface ControlListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link ControlParser#prog}.
	 * @param ctx the parse tree
	 */
	void enterProg(ControlParser.ProgContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#prog}.
	 * @param ctx the parse tree
	 */
	void exitProg(ControlParser.ProgContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(ControlParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(ControlParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentencias(ControlParser.SentenciasContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentencias(ControlParser.SentenciasContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#instWhile}.
	 * @param ctx the parse tree
	 */
	void enterInstWhile(ControlParser.InstWhileContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#instWhile}.
	 * @param ctx the parse tree
	 */
	void exitInstWhile(ControlParser.InstWhileContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#instFor}.
	 * @param ctx the parse tree
	 */
	void enterInstFor(ControlParser.InstForContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#instFor}.
	 * @param ctx the parse tree
	 */
	void exitInstFor(ControlParser.InstForContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#instGuard}.
	 * @param ctx the parse tree
	 */
	void enterInstGuard(ControlParser.InstGuardContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#instGuard}.
	 * @param ctx the parse tree
	 */
	void exitInstGuard(ControlParser.InstGuardContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#instBreak}.
	 * @param ctx the parse tree
	 */
	void enterInstBreak(ControlParser.InstBreakContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#instBreak}.
	 * @param ctx the parse tree
	 */
	void exitInstBreak(ControlParser.InstBreakContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#instContinue}.
	 * @param ctx the parse tree
	 */
	void enterInstContinue(ControlParser.InstContinueContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#instContinue}.
	 * @param ctx the parse tree
	 */
	void exitInstContinue(ControlParser.InstContinueContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#instReturn}.
	 * @param ctx the parse tree
	 */
	void enterInstReturn(ControlParser.InstReturnContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#instReturn}.
	 * @param ctx the parse tree
	 */
	void exitInstReturn(ControlParser.InstReturnContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#decVector}.
	 * @param ctx the parse tree
	 */
	void enterDecVector(ControlParser.DecVectorContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#decVector}.
	 * @param ctx the parse tree
	 */
	void exitDecVector(ControlParser.DecVectorContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#defVector}.
	 * @param ctx the parse tree
	 */
	void enterDefVector(ControlParser.DefVectorContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#defVector}.
	 * @param ctx the parse tree
	 */
	void exitDefVector(ControlParser.DefVectorContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#funcsVectoriales}.
	 * @param ctx the parse tree
	 */
	void enterFuncsVectoriales(ControlParser.FuncsVectorialesContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#funcsVectoriales}.
	 * @param ctx the parse tree
	 */
	void exitFuncsVectoriales(ControlParser.FuncsVectorialesContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#decMatriz}.
	 * @param ctx the parse tree
	 */
	void enterDecMatriz(ControlParser.DecMatrizContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#decMatriz}.
	 * @param ctx the parse tree
	 */
	void exitDecMatriz(ControlParser.DecMatrizContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#tipoMatriz}.
	 * @param ctx the parse tree
	 */
	void enterTipoMatriz(ControlParser.TipoMatrizContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#tipoMatriz}.
	 * @param ctx the parse tree
	 */
	void exitTipoMatriz(ControlParser.TipoMatrizContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#defMatrix}.
	 * @param ctx the parse tree
	 */
	void enterDefMatrix(ControlParser.DefMatrixContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#defMatrix}.
	 * @param ctx the parse tree
	 */
	void exitDefMatrix(ControlParser.DefMatrixContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#listaValoresMat}.
	 * @param ctx the parse tree
	 */
	void enterListaValoresMat(ControlParser.ListaValoresMatContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#listaValoresMat}.
	 * @param ctx the parse tree
	 */
	void exitListaValoresMat(ControlParser.ListaValoresMatContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#listaValoresMat2}.
	 * @param ctx the parse tree
	 */
	void enterListaValoresMat2(ControlParser.ListaValoresMat2Context ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#listaValoresMat2}.
	 * @param ctx the parse tree
	 */
	void exitListaValoresMat2(ControlParser.ListaValoresMat2Context ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#simpleMax}.
	 * @param ctx the parse tree
	 */
	void enterSimpleMax(ControlParser.SimpleMaxContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#simpleMax}.
	 * @param ctx the parse tree
	 */
	void exitSimpleMax(ControlParser.SimpleMaxContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#instStruct}.
	 * @param ctx the parse tree
	 */
	void enterInstStruct(ControlParser.InstStructContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#instStruct}.
	 * @param ctx the parse tree
	 */
	void exitInstStruct(ControlParser.InstStructContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#listaAtributos}.
	 * @param ctx the parse tree
	 */
	void enterListaAtributos(ControlParser.ListaAtributosContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#listaAtributos}.
	 * @param ctx the parse tree
	 */
	void exitListaAtributos(ControlParser.ListaAtributosContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#decStruct}.
	 * @param ctx the parse tree
	 */
	void enterDecStruct(ControlParser.DecStructContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#decStruct}.
	 * @param ctx the parse tree
	 */
	void exitDecStruct(ControlParser.DecStructContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#lDupla}.
	 * @param ctx the parse tree
	 */
	void enterLDupla(ControlParser.LDuplaContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#lDupla}.
	 * @param ctx the parse tree
	 */
	void exitLDupla(ControlParser.LDuplaContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#decFunciones}.
	 * @param ctx the parse tree
	 */
	void enterDecFunciones(ControlParser.DecFuncionesContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#decFunciones}.
	 * @param ctx the parse tree
	 */
	void exitDecFunciones(ControlParser.DecFuncionesContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#listaParametros}.
	 * @param ctx the parse tree
	 */
	void enterListaParametros(ControlParser.ListaParametrosContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#listaParametros}.
	 * @param ctx the parse tree
	 */
	void exitListaParametros(ControlParser.ListaParametrosContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#llamadaFuncion}.
	 * @param ctx the parse tree
	 */
	void enterLlamadaFuncion(ControlParser.LlamadaFuncionContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#llamadaFuncion}.
	 * @param ctx the parse tree
	 */
	void exitLlamadaFuncion(ControlParser.LlamadaFuncionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#listaLlamadaParametros}.
	 * @param ctx the parse tree
	 */
	void enterListaLlamadaParametros(ControlParser.ListaLlamadaParametrosContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#listaLlamadaParametros}.
	 * @param ctx the parse tree
	 */
	void exitListaLlamadaParametros(ControlParser.ListaLlamadaParametrosContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#instPrint}.
	 * @param ctx the parse tree
	 */
	void enterInstPrint(ControlParser.InstPrintContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#instPrint}.
	 * @param ctx the parse tree
	 */
	void exitInstPrint(ControlParser.InstPrintContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#instCasteos}.
	 * @param ctx the parse tree
	 */
	void enterInstCasteos(ControlParser.InstCasteosContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#instCasteos}.
	 * @param ctx the parse tree
	 */
	void exitInstCasteos(ControlParser.InstCasteosContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#numList}.
	 * @param ctx the parse tree
	 */
	void enterNumList(ControlParser.NumListContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#numList}.
	 * @param ctx the parse tree
	 */
	void exitNumList(ControlParser.NumListContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracion(ControlParser.DeclaracionContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracion(ControlParser.DeclaracionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion(ControlParser.AsignacionContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#asignacion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion(ControlParser.AsignacionContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr(ControlParser.ExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr(ControlParser.ExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#tipovariable}.
	 * @param ctx the parse tree
	 */
	void enterTipovariable(ControlParser.TipovariableContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#tipovariable}.
	 * @param ctx the parse tree
	 */
	void exitTipovariable(ControlParser.TipovariableContext ctx);
}