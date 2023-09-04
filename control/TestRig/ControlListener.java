// Generated from Control.g4 by ANTLR 4.13.0
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link ControlParser}.
 */
public interface ControlListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by the {@code ProgInicial}
	 * labeled alternative in {@link ControlParser#prog}.
	 * @param ctx the parse tree
	 */
	void enterProgInicial(ControlParser.ProgInicialContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ProgInicial}
	 * labeled alternative in {@link ControlParser#prog}.
	 * @param ctx the parse tree
	 */
	void exitProgInicial(ControlParser.ProgInicialContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BlockContent}
	 * labeled alternative in {@link ControlParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlockContent(ControlParser.BlockContentContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BlockContent}
	 * labeled alternative in {@link ControlParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlockContent(ControlParser.BlockContentContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaDeclaracion}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaDeclaracion(ControlParser.SentenciaDeclaracionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaDeclaracion}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaDeclaracion(ControlParser.SentenciaDeclaracionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaDeclaracionVector}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaDeclaracionVector(ControlParser.SentenciaDeclaracionVectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaDeclaracionVector}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaDeclaracionVector(ControlParser.SentenciaDeclaracionVectorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaDeclaracionMatrix}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaDeclaracionMatrix(ControlParser.SentenciaDeclaracionMatrixContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaDeclaracionMatrix}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaDeclaracionMatrix(ControlParser.SentenciaDeclaracionMatrixContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaDeclaracionFuncion}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaDeclaracionFuncion(ControlParser.SentenciaDeclaracionFuncionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaDeclaracionFuncion}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaDeclaracionFuncion(ControlParser.SentenciaDeclaracionFuncionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaAsignacion}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaAsignacion(ControlParser.SentenciaAsignacionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaAsignacion}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaAsignacion(ControlParser.SentenciaAsignacionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaIf}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaIf(ControlParser.SentenciaIfContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaIf}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaIf(ControlParser.SentenciaIfContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaSwitch}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaSwitch(ControlParser.SentenciaSwitchContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaSwitch}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaSwitch(ControlParser.SentenciaSwitchContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaWhile}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaWhile(ControlParser.SentenciaWhileContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaWhile}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaWhile(ControlParser.SentenciaWhileContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaFor}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaFor(ControlParser.SentenciaForContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaFor}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaFor(ControlParser.SentenciaForContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaGuard}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaGuard(ControlParser.SentenciaGuardContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaGuard}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaGuard(ControlParser.SentenciaGuardContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaBreak}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaBreak(ControlParser.SentenciaBreakContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaBreak}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaBreak(ControlParser.SentenciaBreakContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaContinue}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaContinue(ControlParser.SentenciaContinueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaContinue}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaContinue(ControlParser.SentenciaContinueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaReturn}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaReturn(ControlParser.SentenciaReturnContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaReturn}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaReturn(ControlParser.SentenciaReturnContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaFuncVectoriales}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaFuncVectoriales(ControlParser.SentenciaFuncVectorialesContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaFuncVectoriales}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaFuncVectoriales(ControlParser.SentenciaFuncVectorialesContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaLLamadaFuncion}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaLLamadaFuncion(ControlParser.SentenciaLLamadaFuncionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaLLamadaFuncion}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaLLamadaFuncion(ControlParser.SentenciaLLamadaFuncionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaPrint}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaPrint(ControlParser.SentenciaPrintContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaPrint}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaPrint(ControlParser.SentenciaPrintContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaStruct}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaStruct(ControlParser.SentenciaStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaStruct}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaStruct(ControlParser.SentenciaStructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaDeclaracionStruct}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaDeclaracionStruct(ControlParser.SentenciaDeclaracionStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaDeclaracionStruct}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaDeclaracionStruct(ControlParser.SentenciaDeclaracionStructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SentenciaLLamadaMetodos}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void enterSentenciaLLamadaMetodos(ControlParser.SentenciaLLamadaMetodosContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SentenciaLLamadaMetodos}
	 * labeled alternative in {@link ControlParser#sentencias}.
	 * @param ctx the parse tree
	 */
	void exitSentenciaLLamadaMetodos(ControlParser.SentenciaLLamadaMetodosContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#insIf}.
	 * @param ctx the parse tree
	 */
	void enterInsIf(ControlParser.InsIfContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#insIf}.
	 * @param ctx the parse tree
	 */
	void exitInsIf(ControlParser.InsIfContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#instElseIf}.
	 * @param ctx the parse tree
	 */
	void enterInstElseIf(ControlParser.InstElseIfContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#instElseIf}.
	 * @param ctx the parse tree
	 */
	void exitInstElseIf(ControlParser.InstElseIfContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#insElse}.
	 * @param ctx the parse tree
	 */
	void enterInsElse(ControlParser.InsElseContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#insElse}.
	 * @param ctx the parse tree
	 */
	void exitInsElse(ControlParser.InsElseContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#insSwitch}.
	 * @param ctx the parse tree
	 */
	void enterInsSwitch(ControlParser.InsSwitchContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#insSwitch}.
	 * @param ctx the parse tree
	 */
	void exitInsSwitch(ControlParser.InsSwitchContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#instCase}.
	 * @param ctx the parse tree
	 */
	void enterInstCase(ControlParser.InstCaseContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#instCase}.
	 * @param ctx the parse tree
	 */
	void exitInstCase(ControlParser.InstCaseContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#instDefault}.
	 * @param ctx the parse tree
	 */
	void enterInstDefault(ControlParser.InstDefaultContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#instDefault}.
	 * @param ctx the parse tree
	 */
	void exitInstDefault(ControlParser.InstDefaultContext ctx);
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
	 * Enter a parse tree produced by {@link ControlParser#llamAtributos}.
	 * @param ctx the parse tree
	 */
	void enterLlamAtributos(ControlParser.LlamAtributosContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#llamAtributos}.
	 * @param ctx the parse tree
	 */
	void exitLlamAtributos(ControlParser.LlamAtributosContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#llamadaMetodos}.
	 * @param ctx the parse tree
	 */
	void enterLlamadaMetodos(ControlParser.LlamadaMetodosContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#llamadaMetodos}.
	 * @param ctx the parse tree
	 */
	void exitLlamadaMetodos(ControlParser.LlamadaMetodosContext ctx);
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
	 * Enter a parse tree produced by the {@code IntruccionPrint}
	 * labeled alternative in {@link ControlParser#instPrint}.
	 * @param ctx the parse tree
	 */
	void enterIntruccionPrint(ControlParser.IntruccionPrintContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IntruccionPrint}
	 * labeled alternative in {@link ControlParser#instPrint}.
	 * @param ctx the parse tree
	 */
	void exitIntruccionPrint(ControlParser.IntruccionPrintContext ctx);
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
	 * Enter a parse tree produced by {@link ControlParser#expresionList}.
	 * @param ctx the parse tree
	 */
	void enterExpresionList(ControlParser.ExpresionListContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#expresionList}.
	 * @param ctx the parse tree
	 */
	void exitExpresionList(ControlParser.ExpresionListContext ctx);
	/**
	 * Enter a parse tree produced by {@link ControlParser#objectsList}.
	 * @param ctx the parse tree
	 */
	void enterObjectsList(ControlParser.ObjectsListContext ctx);
	/**
	 * Exit a parse tree produced by {@link ControlParser#objectsList}.
	 * @param ctx the parse tree
	 */
	void exitObjectsList(ControlParser.ObjectsListContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Dec_VTipoValor}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDec_VTipoValor(ControlParser.Dec_VTipoValorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Dec_VTipoValor}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDec_VTipoValor(ControlParser.Dec_VTipoValorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Dec_VTipo}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDec_VTipo(ControlParser.Dec_VTipoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Dec_VTipo}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDec_VTipo(ControlParser.Dec_VTipoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Dec_VVal}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDec_VVal(ControlParser.Dec_VValContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Dec_VVal}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDec_VVal(ControlParser.Dec_VValContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Dec_CVal}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDec_CVal(ControlParser.Dec_CValContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Dec_CVal}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDec_CVal(ControlParser.Dec_CValContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Dec_CTipoValor}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDec_CTipoValor(ControlParser.Dec_CTipoValorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Dec_CTipoValor}
	 * labeled alternative in {@link ControlParser#declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDec_CTipoValor(ControlParser.Dec_CTipoValorContext ctx);
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
	 * Enter a parse tree produced by the {@code Expr_Decimal}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Decimal(ControlParser.Expr_DecimalContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Decimal}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Decimal(ControlParser.Expr_DecimalContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_InstCasteo}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_InstCasteo(ControlParser.Expr_InstCasteoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_InstCasteo}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_InstCasteo(ControlParser.Expr_InstCasteoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ValidaIgualDif}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ValidaIgualDif(ControlParser.Expr_ValidaIgualDifContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ValidaIgualDif}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ValidaIgualDif(ControlParser.Expr_ValidaIgualDifContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_OpMulDiv}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_OpMulDiv(ControlParser.Expr_OpMulDivContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_OpMulDiv}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_OpMulDiv(ControlParser.Expr_OpMulDivContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ValidaAnd}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ValidaAnd(ControlParser.Expr_ValidaAndContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ValidaAnd}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ValidaAnd(ControlParser.Expr_ValidaAndContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_LlamAtributos}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_LlamAtributos(ControlParser.Expr_LlamAtributosContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_LlamAtributos}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_LlamAtributos(ControlParser.Expr_LlamAtributosContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Nil}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Nil(ControlParser.Expr_NilContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Nil}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Nil(ControlParser.Expr_NilContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_PosMatrix}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_PosMatrix(ControlParser.Expr_PosMatrixContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_PosMatrix}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_PosMatrix(ControlParser.Expr_PosMatrixContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Boolean}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Boolean(ControlParser.Expr_BooleanContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Boolean}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Boolean(ControlParser.Expr_BooleanContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_InstRango}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_InstRango(ControlParser.Expr_InstRangoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_InstRango}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_InstRango(ControlParser.Expr_InstRangoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ValNumNeg}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ValNumNeg(ControlParser.Expr_ValNumNegContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ValNumNeg}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ValNumNeg(ControlParser.Expr_ValNumNegContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ValidaOr}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ValidaOr(ControlParser.Expr_ValidaOrContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ValidaOr}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ValidaOr(ControlParser.Expr_ValidaOrContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ValNegacion}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ValNegacion(ControlParser.Expr_ValNegacionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ValNegacion}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ValNegacion(ControlParser.Expr_ValNegacionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ID}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ID(ControlParser.Expr_IDContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ID}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ID(ControlParser.Expr_IDContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ValidaMayQue}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ValidaMayQue(ControlParser.Expr_ValidaMayQueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ValidaMayQue}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ValidaMayQue(ControlParser.Expr_ValidaMayQueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Conteo}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Conteo(ControlParser.Expr_ConteoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Conteo}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Conteo(ControlParser.Expr_ConteoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_OpSumRes}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_OpSumRes(ControlParser.Expr_OpSumResContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_OpSumRes}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_OpSumRes(ControlParser.Expr_OpSumResContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Entero}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Entero(ControlParser.Expr_EnteroContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Entero}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Entero(ControlParser.Expr_EnteroContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ParExpre}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ParExpre(ControlParser.Expr_ParExpreContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ParExpre}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ParExpre(ControlParser.Expr_ParExpreContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_StringChar}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_StringChar(ControlParser.Expr_StringCharContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_StringChar}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_StringChar(ControlParser.Expr_StringCharContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_ValidaMenQue}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_ValidaMenQue(ControlParser.Expr_ValidaMenQueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_ValidaMenQue}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_ValidaMenQue(ControlParser.Expr_ValidaMenQueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_PosVector}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpr_PosVector(ControlParser.Expr_PosVectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_PosVector}
	 * labeled alternative in {@link ControlParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpr_PosVector(ControlParser.Expr_PosVectorContext ctx);
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