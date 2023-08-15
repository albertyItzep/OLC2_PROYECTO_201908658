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