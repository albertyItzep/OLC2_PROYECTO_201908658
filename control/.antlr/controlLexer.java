// Generated from c:\Users\Xorg123\go\src\OLC2_PROYECTO_201908658\control\Control.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ControlLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T__31=32, T__32=33, T__33=34, T__34=35, T__35=36, T__36=37, T__37=38, 
		T__38=39, T__39=40, T__40=41, T__41=42, T__42=43, T__43=44, T__44=45, 
		T__45=46, T__46=47, T__47=48, T__48=49, T__49=50, T__50=51, T__51=52, 
		T__52=53, T__53=54, T__54=55, T__55=56, T__56=57, T__57=58, T__58=59, 
		T__59=60, T__60=61, T__61=62, T__62=63, T__63=64, T__64=65, T__65=66, 
		T__66=67, ID=68, INT=69, DECIMAL=70, STRING=71, WS=72, LINEALCOMENT=73, 
		BlockComment=74;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
			"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24", 
			"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32", 
			"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40", 
			"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "T__48", 
			"T__49", "T__50", "T__51", "T__52", "T__53", "T__54", "T__55", "T__56", 
			"T__57", "T__58", "T__59", "T__60", "T__61", "T__62", "T__63", "T__64", 
			"T__65", "T__66", "ID", "INT", "DECIMAL", "STRING", "WS", "LINEALCOMENT", 
			"BlockComment"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'if'", "'{'", "'}'", "'else'", "'switch'", "'case'", "':'", "'default'", 
			"'while'", "'for'", "'in'", "'guard'", "'break'", "'continue'", "'return'", 
			"'var'", "'['", "']'", "'='", "'let'", "'.'", "'append'", "'('", "')'", 
			"'remove'", "'at'", "'removeLast'", "','", "'repeating'", "'count'", 
			"'struct'", "';'", "'mutating'", "'self'", "'func'", "'->'", "'_'", "'inout'", 
			"'&'", "'print'", "'?'", "'+='", "'-='", "'-'", "'!'", "'*'", "'/'", 
			"'+'", "'%'", "'>='", "'>'", "'<='", "'<'", "'=='", "'!='", "'&&'", "'||'", 
			"'true'", "'false'", "'nil'", "'isEmpty'", "'...'", "'String'", "'Float'", 
			"'Int'", "'Bool'", "'Character'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, "ID", "INT", "DECIMAL", 
			"STRING", "WS", "LINEALCOMENT", "BlockComment"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public ControlLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Control.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2L\u01f7\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\3\2\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3"+
		"\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3"+
		"\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3"+
		"\21\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3\24\3\25\3\25\3\25\3\25\3\26\3"+
		"\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\31\3\31\3\32\3\32\3"+
		"\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\34\3"+
		"\34\3\34\3\34\3\34\3\34\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3"+
		"\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3 \3!\3"+
		"!\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3$\3$\3$\3$\3$\3"+
		"%\3%\3%\3&\3&\3\'\3\'\3\'\3\'\3\'\3\'\3(\3(\3)\3)\3)\3)\3)\3)\3*\3*\3"+
		"+\3+\3+\3,\3,\3,\3-\3-\3.\3.\3/\3/\3\60\3\60\3\61\3\61\3\62\3\62\3\63"+
		"\3\63\3\63\3\64\3\64\3\65\3\65\3\65\3\66\3\66\3\67\3\67\3\67\38\38\38"+
		"\39\39\39\3:\3:\3:\3;\3;\3;\3;\3;\3<\3<\3<\3<\3<\3<\3=\3=\3=\3=\3>\3>"+
		"\3>\3>\3>\3>\3>\3>\3?\3?\3?\3?\3@\3@\3@\3@\3@\3@\3@\3A\3A\3A\3A\3A\3A"+
		"\3B\3B\3B\3B\3C\3C\3C\3C\3C\3D\3D\3D\3D\3D\3D\3D\3D\3D\3D\3E\3E\7E\u01bb"+
		"\nE\fE\16E\u01be\13E\3F\6F\u01c1\nF\rF\16F\u01c2\3G\6G\u01c6\nG\rG\16"+
		"G\u01c7\3G\3G\6G\u01cc\nG\rG\16G\u01cd\3H\3H\3H\3H\7H\u01d4\nH\fH\16H"+
		"\u01d7\13H\3H\3H\3I\3I\3I\3I\3J\3J\3J\3J\7J\u01e3\nJ\fJ\16J\u01e6\13J"+
		"\3J\3J\3K\3K\3K\3K\7K\u01ee\nK\fK\16K\u01f1\13K\3K\3K\3K\3K\3K\3\u01ef"+
		"\2L\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35"+
		"\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36"+
		";\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67"+
		"m8o9q:s;u<w=y>{?}@\177A\u0081B\u0083C\u0085D\u0087E\u0089F\u008bG\u008d"+
		"H\u008fI\u0091J\u0093K\u0095L\3\2\b\5\2C\\aac|\6\2\62;C\\aac|\3\2\62;"+
		"\3\2$$\5\2\13\f\17\17\"\"\4\2\f\f\17\17\2\u01fe\2\3\3\2\2\2\2\5\3\2\2"+
		"\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21"+
		"\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2"+
		"\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3"+
		"\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3"+
		"\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3"+
		"\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2"+
		"\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2"+
		"Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3"+
		"\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q\3\2\2"+
		"\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2\2\2}\3\2\2\2\2"+
		"\177\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2\2\2\2\u0085\3\2\2\2\2\u0087\3"+
		"\2\2\2\2\u0089\3\2\2\2\2\u008b\3\2\2\2\2\u008d\3\2\2\2\2\u008f\3\2\2\2"+
		"\2\u0091\3\2\2\2\2\u0093\3\2\2\2\2\u0095\3\2\2\2\3\u0097\3\2\2\2\5\u009a"+
		"\3\2\2\2\7\u009c\3\2\2\2\t\u009e\3\2\2\2\13\u00a3\3\2\2\2\r\u00aa\3\2"+
		"\2\2\17\u00af\3\2\2\2\21\u00b1\3\2\2\2\23\u00b9\3\2\2\2\25\u00bf\3\2\2"+
		"\2\27\u00c3\3\2\2\2\31\u00c6\3\2\2\2\33\u00cc\3\2\2\2\35\u00d2\3\2\2\2"+
		"\37\u00db\3\2\2\2!\u00e2\3\2\2\2#\u00e6\3\2\2\2%\u00e8\3\2\2\2\'\u00ea"+
		"\3\2\2\2)\u00ec\3\2\2\2+\u00f0\3\2\2\2-\u00f2\3\2\2\2/\u00f9\3\2\2\2\61"+
		"\u00fb\3\2\2\2\63\u00fd\3\2\2\2\65\u0104\3\2\2\2\67\u0107\3\2\2\29\u0112"+
		"\3\2\2\2;\u0114\3\2\2\2=\u011e\3\2\2\2?\u0124\3\2\2\2A\u012b\3\2\2\2C"+
		"\u012d\3\2\2\2E\u0136\3\2\2\2G\u013b\3\2\2\2I\u0140\3\2\2\2K\u0143\3\2"+
		"\2\2M\u0145\3\2\2\2O\u014b\3\2\2\2Q\u014d\3\2\2\2S\u0153\3\2\2\2U\u0155"+
		"\3\2\2\2W\u0158\3\2\2\2Y\u015b\3\2\2\2[\u015d\3\2\2\2]\u015f\3\2\2\2_"+
		"\u0161\3\2\2\2a\u0163\3\2\2\2c\u0165\3\2\2\2e\u0167\3\2\2\2g\u016a\3\2"+
		"\2\2i\u016c\3\2\2\2k\u016f\3\2\2\2m\u0171\3\2\2\2o\u0174\3\2\2\2q\u0177"+
		"\3\2\2\2s\u017a\3\2\2\2u\u017d\3\2\2\2w\u0182\3\2\2\2y\u0188\3\2\2\2{"+
		"\u018c\3\2\2\2}\u0194\3\2\2\2\177\u0198\3\2\2\2\u0081\u019f\3\2\2\2\u0083"+
		"\u01a5\3\2\2\2\u0085\u01a9\3\2\2\2\u0087\u01ae\3\2\2\2\u0089\u01b8\3\2"+
		"\2\2\u008b\u01c0\3\2\2\2\u008d\u01c5\3\2\2\2\u008f\u01cf\3\2\2\2\u0091"+
		"\u01da\3\2\2\2\u0093\u01de\3\2\2\2\u0095\u01e9\3\2\2\2\u0097\u0098\7k"+
		"\2\2\u0098\u0099\7h\2\2\u0099\4\3\2\2\2\u009a\u009b\7}\2\2\u009b\6\3\2"+
		"\2\2\u009c\u009d\7\177\2\2\u009d\b\3\2\2\2\u009e\u009f\7g\2\2\u009f\u00a0"+
		"\7n\2\2\u00a0\u00a1\7u\2\2\u00a1\u00a2\7g\2\2\u00a2\n\3\2\2\2\u00a3\u00a4"+
		"\7u\2\2\u00a4\u00a5\7y\2\2\u00a5\u00a6\7k\2\2\u00a6\u00a7\7v\2\2\u00a7"+
		"\u00a8\7e\2\2\u00a8\u00a9\7j\2\2\u00a9\f\3\2\2\2\u00aa\u00ab\7e\2\2\u00ab"+
		"\u00ac\7c\2\2\u00ac\u00ad\7u\2\2\u00ad\u00ae\7g\2\2\u00ae\16\3\2\2\2\u00af"+
		"\u00b0\7<\2\2\u00b0\20\3\2\2\2\u00b1\u00b2\7f\2\2\u00b2\u00b3\7g\2\2\u00b3"+
		"\u00b4\7h\2\2\u00b4\u00b5\7c\2\2\u00b5\u00b6\7w\2\2\u00b6\u00b7\7n\2\2"+
		"\u00b7\u00b8\7v\2\2\u00b8\22\3\2\2\2\u00b9\u00ba\7y\2\2\u00ba\u00bb\7"+
		"j\2\2\u00bb\u00bc\7k\2\2\u00bc\u00bd\7n\2\2\u00bd\u00be\7g\2\2\u00be\24"+
		"\3\2\2\2\u00bf\u00c0\7h\2\2\u00c0\u00c1\7q\2\2\u00c1\u00c2\7t\2\2\u00c2"+
		"\26\3\2\2\2\u00c3\u00c4\7k\2\2\u00c4\u00c5\7p\2\2\u00c5\30\3\2\2\2\u00c6"+
		"\u00c7\7i\2\2\u00c7\u00c8\7w\2\2\u00c8\u00c9\7c\2\2\u00c9\u00ca\7t\2\2"+
		"\u00ca\u00cb\7f\2\2\u00cb\32\3\2\2\2\u00cc\u00cd\7d\2\2\u00cd\u00ce\7"+
		"t\2\2\u00ce\u00cf\7g\2\2\u00cf\u00d0\7c\2\2\u00d0\u00d1\7m\2\2\u00d1\34"+
		"\3\2\2\2\u00d2\u00d3\7e\2\2\u00d3\u00d4\7q\2\2\u00d4\u00d5\7p\2\2\u00d5"+
		"\u00d6\7v\2\2\u00d6\u00d7\7k\2\2\u00d7\u00d8\7p\2\2\u00d8\u00d9\7w\2\2"+
		"\u00d9\u00da\7g\2\2\u00da\36\3\2\2\2\u00db\u00dc\7t\2\2\u00dc\u00dd\7"+
		"g\2\2\u00dd\u00de\7v\2\2\u00de\u00df\7w\2\2\u00df\u00e0\7t\2\2\u00e0\u00e1"+
		"\7p\2\2\u00e1 \3\2\2\2\u00e2\u00e3\7x\2\2\u00e3\u00e4\7c\2\2\u00e4\u00e5"+
		"\7t\2\2\u00e5\"\3\2\2\2\u00e6\u00e7\7]\2\2\u00e7$\3\2\2\2\u00e8\u00e9"+
		"\7_\2\2\u00e9&\3\2\2\2\u00ea\u00eb\7?\2\2\u00eb(\3\2\2\2\u00ec\u00ed\7"+
		"n\2\2\u00ed\u00ee\7g\2\2\u00ee\u00ef\7v\2\2\u00ef*\3\2\2\2\u00f0\u00f1"+
		"\7\60\2\2\u00f1,\3\2\2\2\u00f2\u00f3\7c\2\2\u00f3\u00f4\7r\2\2\u00f4\u00f5"+
		"\7r\2\2\u00f5\u00f6\7g\2\2\u00f6\u00f7\7p\2\2\u00f7\u00f8\7f\2\2\u00f8"+
		".\3\2\2\2\u00f9\u00fa\7*\2\2\u00fa\60\3\2\2\2\u00fb\u00fc\7+\2\2\u00fc"+
		"\62\3\2\2\2\u00fd\u00fe\7t\2\2\u00fe\u00ff\7g\2\2\u00ff\u0100\7o\2\2\u0100"+
		"\u0101\7q\2\2\u0101\u0102\7x\2\2\u0102\u0103\7g\2\2\u0103\64\3\2\2\2\u0104"+
		"\u0105\7c\2\2\u0105\u0106\7v\2\2\u0106\66\3\2\2\2\u0107\u0108\7t\2\2\u0108"+
		"\u0109\7g\2\2\u0109\u010a\7o\2\2\u010a\u010b\7q\2\2\u010b\u010c\7x\2\2"+
		"\u010c\u010d\7g\2\2\u010d\u010e\7N\2\2\u010e\u010f\7c\2\2\u010f\u0110"+
		"\7u\2\2\u0110\u0111\7v\2\2\u01118\3\2\2\2\u0112\u0113\7.\2\2\u0113:\3"+
		"\2\2\2\u0114\u0115\7t\2\2\u0115\u0116\7g\2\2\u0116\u0117\7r\2\2\u0117"+
		"\u0118\7g\2\2\u0118\u0119\7c\2\2\u0119\u011a\7v\2\2\u011a\u011b\7k\2\2"+
		"\u011b\u011c\7p\2\2\u011c\u011d\7i\2\2\u011d<\3\2\2\2\u011e\u011f\7e\2"+
		"\2\u011f\u0120\7q\2\2\u0120\u0121\7w\2\2\u0121\u0122\7p\2\2\u0122\u0123"+
		"\7v\2\2\u0123>\3\2\2\2\u0124\u0125\7u\2\2\u0125\u0126\7v\2\2\u0126\u0127"+
		"\7t\2\2\u0127\u0128\7w\2\2\u0128\u0129\7e\2\2\u0129\u012a\7v\2\2\u012a"+
		"@\3\2\2\2\u012b\u012c\7=\2\2\u012cB\3\2\2\2\u012d\u012e\7o\2\2\u012e\u012f"+
		"\7w\2\2\u012f\u0130\7v\2\2\u0130\u0131\7c\2\2\u0131\u0132\7v\2\2\u0132"+
		"\u0133\7k\2\2\u0133\u0134\7p\2\2\u0134\u0135\7i\2\2\u0135D\3\2\2\2\u0136"+
		"\u0137\7u\2\2\u0137\u0138\7g\2\2\u0138\u0139\7n\2\2\u0139\u013a\7h\2\2"+
		"\u013aF\3\2\2\2\u013b\u013c\7h\2\2\u013c\u013d\7w\2\2\u013d\u013e\7p\2"+
		"\2\u013e\u013f\7e\2\2\u013fH\3\2\2\2\u0140\u0141\7/\2\2\u0141\u0142\7"+
		"@\2\2\u0142J\3\2\2\2\u0143\u0144\7a\2\2\u0144L\3\2\2\2\u0145\u0146\7k"+
		"\2\2\u0146\u0147\7p\2\2\u0147\u0148\7q\2\2\u0148\u0149\7w\2\2\u0149\u014a"+
		"\7v\2\2\u014aN\3\2\2\2\u014b\u014c\7(\2\2\u014cP\3\2\2\2\u014d\u014e\7"+
		"r\2\2\u014e\u014f\7t\2\2\u014f\u0150\7k\2\2\u0150\u0151\7p\2\2\u0151\u0152"+
		"\7v\2\2\u0152R\3\2\2\2\u0153\u0154\7A\2\2\u0154T\3\2\2\2\u0155\u0156\7"+
		"-\2\2\u0156\u0157\7?\2\2\u0157V\3\2\2\2\u0158\u0159\7/\2\2\u0159\u015a"+
		"\7?\2\2\u015aX\3\2\2\2\u015b\u015c\7/\2\2\u015cZ\3\2\2\2\u015d\u015e\7"+
		"#\2\2\u015e\\\3\2\2\2\u015f\u0160\7,\2\2\u0160^\3\2\2\2\u0161\u0162\7"+
		"\61\2\2\u0162`\3\2\2\2\u0163\u0164\7-\2\2\u0164b\3\2\2\2\u0165\u0166\7"+
		"\'\2\2\u0166d\3\2\2\2\u0167\u0168\7@\2\2\u0168\u0169\7?\2\2\u0169f\3\2"+
		"\2\2\u016a\u016b\7@\2\2\u016bh\3\2\2\2\u016c\u016d\7>\2\2\u016d\u016e"+
		"\7?\2\2\u016ej\3\2\2\2\u016f\u0170\7>\2\2\u0170l\3\2\2\2\u0171\u0172\7"+
		"?\2\2\u0172\u0173\7?\2\2\u0173n\3\2\2\2\u0174\u0175\7#\2\2\u0175\u0176"+
		"\7?\2\2\u0176p\3\2\2\2\u0177\u0178\7(\2\2\u0178\u0179\7(\2\2\u0179r\3"+
		"\2\2\2\u017a\u017b\7~\2\2\u017b\u017c\7~\2\2\u017ct\3\2\2\2\u017d\u017e"+
		"\7v\2\2\u017e\u017f\7t\2\2\u017f\u0180\7w\2\2\u0180\u0181\7g\2\2\u0181"+
		"v\3\2\2\2\u0182\u0183\7h\2\2\u0183\u0184\7c\2\2\u0184\u0185\7n\2\2\u0185"+
		"\u0186\7u\2\2\u0186\u0187\7g\2\2\u0187x\3\2\2\2\u0188\u0189\7p\2\2\u0189"+
		"\u018a\7k\2\2\u018a\u018b\7n\2\2\u018bz\3\2\2\2\u018c\u018d\7k\2\2\u018d"+
		"\u018e\7u\2\2\u018e\u018f\7G\2\2\u018f\u0190\7o\2\2\u0190\u0191\7r\2\2"+
		"\u0191\u0192\7v\2\2\u0192\u0193\7{\2\2\u0193|\3\2\2\2\u0194\u0195\7\60"+
		"\2\2\u0195\u0196\7\60\2\2\u0196\u0197\7\60\2\2\u0197~\3\2\2\2\u0198\u0199"+
		"\7U\2\2\u0199\u019a\7v\2\2\u019a\u019b\7t\2\2\u019b\u019c\7k\2\2\u019c"+
		"\u019d\7p\2\2\u019d\u019e\7i\2\2\u019e\u0080\3\2\2\2\u019f\u01a0\7H\2"+
		"\2\u01a0\u01a1\7n\2\2\u01a1\u01a2\7q\2\2\u01a2\u01a3\7c\2\2\u01a3\u01a4"+
		"\7v\2\2\u01a4\u0082\3\2\2\2\u01a5\u01a6\7K\2\2\u01a6\u01a7\7p\2\2\u01a7"+
		"\u01a8\7v\2\2\u01a8\u0084\3\2\2\2\u01a9\u01aa\7D\2\2\u01aa\u01ab\7q\2"+
		"\2\u01ab\u01ac\7q\2\2\u01ac\u01ad\7n\2\2\u01ad\u0086\3\2\2\2\u01ae\u01af"+
		"\7E\2\2\u01af\u01b0\7j\2\2\u01b0\u01b1\7c\2\2\u01b1\u01b2\7t\2\2\u01b2"+
		"\u01b3\7c\2\2\u01b3\u01b4\7e\2\2\u01b4\u01b5\7v\2\2\u01b5\u01b6\7g\2\2"+
		"\u01b6\u01b7\7t\2\2\u01b7\u0088\3\2\2\2\u01b8\u01bc\t\2\2\2\u01b9\u01bb"+
		"\t\3\2\2\u01ba\u01b9\3\2\2\2\u01bb\u01be\3\2\2\2\u01bc\u01ba\3\2\2\2\u01bc"+
		"\u01bd\3\2\2\2\u01bd\u008a\3\2\2\2\u01be\u01bc\3\2\2\2\u01bf\u01c1\t\4"+
		"\2\2\u01c0\u01bf\3\2\2\2\u01c1\u01c2\3\2\2\2\u01c2\u01c0\3\2\2\2\u01c2"+
		"\u01c3\3\2\2\2\u01c3\u008c\3\2\2\2\u01c4\u01c6\t\4\2\2\u01c5\u01c4\3\2"+
		"\2\2\u01c6\u01c7\3\2\2\2\u01c7\u01c5\3\2\2\2\u01c7\u01c8\3\2\2\2\u01c8"+
		"\u01c9\3\2\2\2\u01c9\u01cb\7\60\2\2\u01ca\u01cc\t\4\2\2\u01cb\u01ca\3"+
		"\2\2\2\u01cc\u01cd\3\2\2\2\u01cd\u01cb\3\2\2\2\u01cd\u01ce\3\2\2\2\u01ce"+
		"\u008e\3\2\2\2\u01cf\u01d5\7$\2\2\u01d0\u01d1\7^\2\2\u01d1\u01d4\7$\2"+
		"\2\u01d2\u01d4\n\5\2\2\u01d3\u01d0\3\2\2\2\u01d3\u01d2\3\2\2\2\u01d4\u01d7"+
		"\3\2\2\2\u01d5\u01d3\3\2\2\2\u01d5\u01d6\3\2\2\2\u01d6\u01d8\3\2\2\2\u01d7"+
		"\u01d5\3\2\2\2\u01d8\u01d9\7$\2\2\u01d9\u0090\3\2\2\2\u01da\u01db\t\6"+
		"\2\2\u01db\u01dc\3\2\2\2\u01dc\u01dd\bI\2\2\u01dd\u0092\3\2\2\2\u01de"+
		"\u01df\7\61\2\2\u01df\u01e0\7\61\2\2\u01e0\u01e4\3\2\2\2\u01e1\u01e3\n"+
		"\7\2\2\u01e2\u01e1\3\2\2\2\u01e3\u01e6\3\2\2\2\u01e4\u01e2\3\2\2\2\u01e4"+
		"\u01e5\3\2\2\2\u01e5\u01e7\3\2\2\2\u01e6\u01e4\3\2\2\2\u01e7\u01e8\bJ"+
		"\3\2\u01e8\u0094\3\2\2\2\u01e9\u01ea\7\61\2\2\u01ea\u01eb\7,\2\2\u01eb"+
		"\u01ef\3\2\2\2\u01ec\u01ee\13\2\2\2\u01ed\u01ec\3\2\2\2\u01ee\u01f1\3"+
		"\2\2\2\u01ef\u01f0\3\2\2\2\u01ef\u01ed\3\2\2\2\u01f0\u01f2\3\2\2\2\u01f1"+
		"\u01ef\3\2\2\2\u01f2\u01f3\7,\2\2\u01f3\u01f4\7\61\2\2\u01f4\u01f5\3\2"+
		"\2\2\u01f5\u01f6\bK\2\2\u01f6\u0096\3\2\2\2\13\2\u01bc\u01c2\u01c7\u01cd"+
		"\u01d3\u01d5\u01e4\u01ef\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}