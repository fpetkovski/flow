// Generated from grammar/PromQLPlus.g4 by ANTLR 4.13.2
// jshint ignore: start
import antlr4 from 'antlr4';
import PromQLPlusListener from './PromQLPlusListener.js';
import PromQLPlusVisitor from './PromQLPlusVisitor.js';

const serializedATN = [4,1,31,104,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,
4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,10,7,10,1,0,1,0,1,0,1,0,1,0,
1,0,3,0,29,8,0,1,1,1,1,1,1,1,1,1,1,1,2,1,2,1,2,5,2,39,8,2,10,2,12,2,42,9,
2,1,3,1,3,1,3,1,3,1,4,1,4,1,4,1,4,1,5,1,5,1,5,5,5,55,8,5,10,5,12,5,58,9,
5,1,6,1,6,3,6,62,8,6,1,7,3,7,65,8,7,1,7,1,7,1,7,1,7,5,7,71,8,7,10,7,12,7,
74,9,7,3,7,76,8,7,1,7,3,7,79,8,7,1,8,1,8,1,8,1,8,1,9,1,9,1,9,1,9,1,9,1,9,
5,9,91,8,9,10,9,12,9,94,9,9,3,9,96,8,9,1,9,3,9,99,8,9,1,10,1,10,1,10,1,10,
0,0,11,0,2,4,6,8,10,12,14,16,18,20,0,3,1,0,2,6,1,0,12,15,1,0,16,17,103,0,
28,1,0,0,0,2,30,1,0,0,0,4,35,1,0,0,0,6,43,1,0,0,0,8,47,1,0,0,0,10,51,1,0,
0,0,12,61,1,0,0,0,14,64,1,0,0,0,16,80,1,0,0,0,18,84,1,0,0,0,20,100,1,0,0,
0,22,23,3,2,1,0,23,24,5,0,0,1,24,29,1,0,0,0,25,26,3,10,5,0,26,27,5,0,0,1,
27,29,1,0,0,0,28,22,1,0,0,0,28,25,1,0,0,0,29,1,1,0,0,0,30,31,5,20,0,0,31,
32,3,4,2,0,32,33,5,21,0,0,33,34,3,8,4,0,34,3,1,0,0,0,35,40,3,6,3,0,36,37,
5,1,0,0,37,39,3,6,3,0,38,36,1,0,0,0,39,42,1,0,0,0,40,38,1,0,0,0,40,41,1,
0,0,0,41,5,1,0,0,0,42,40,1,0,0,0,43,44,5,27,0,0,44,45,5,12,0,0,45,46,3,10,
5,0,46,7,1,0,0,0,47,48,5,27,0,0,48,49,7,0,0,0,49,50,5,27,0,0,50,9,1,0,0,
0,51,56,3,14,7,0,52,53,5,7,0,0,53,55,3,12,6,0,54,52,1,0,0,0,55,58,1,0,0,
0,56,54,1,0,0,0,56,57,1,0,0,0,57,11,1,0,0,0,58,56,1,0,0,0,59,62,3,18,9,0,
60,62,3,20,10,0,61,59,1,0,0,0,61,60,1,0,0,0,62,13,1,0,0,0,63,65,5,27,0,0,
64,63,1,0,0,0,64,65,1,0,0,0,65,78,1,0,0,0,66,75,5,8,0,0,67,72,3,16,8,0,68,
69,5,1,0,0,69,71,3,16,8,0,70,68,1,0,0,0,71,74,1,0,0,0,72,70,1,0,0,0,72,73,
1,0,0,0,73,76,1,0,0,0,74,72,1,0,0,0,75,67,1,0,0,0,75,76,1,0,0,0,76,77,1,
0,0,0,77,79,5,9,0,0,78,66,1,0,0,0,78,79,1,0,0,0,79,15,1,0,0,0,80,81,5,27,
0,0,81,82,7,1,0,0,82,83,5,28,0,0,83,17,1,0,0,0,84,85,5,18,0,0,85,98,7,2,
0,0,86,95,5,10,0,0,87,92,5,27,0,0,88,89,5,1,0,0,89,91,5,27,0,0,90,88,1,0,
0,0,91,94,1,0,0,0,92,90,1,0,0,0,92,93,1,0,0,0,93,96,1,0,0,0,94,92,1,0,0,
0,95,87,1,0,0,0,95,96,1,0,0,0,96,97,1,0,0,0,97,99,5,11,0,0,98,86,1,0,0,0,
98,99,1,0,0,0,99,19,1,0,0,0,100,101,5,27,0,0,101,102,5,19,0,0,102,21,1,0,
0,0,11,28,40,56,61,64,72,75,78,92,95,98];


const atn = new antlr4.atn.ATNDeserializer().deserialize(serializedATN);

const decisionsToDFA = atn.decisionToState.map( (ds, index) => new antlr4.dfa.DFA(ds, index) );

const sharedContextCache = new antlr4.atn.PredictionContextCache();

export default class PromQLPlusParser extends antlr4.Parser {

    static grammarFileName = "PromQLPlus.g4";
    static literalNames = [ null, "','", "'/'", "'+'", "'-'", "'*'", "'%'", 
                            "'|'", "'{'", "'}'", "'('", "')'", "'='", "'!='", 
                            "'=~'", "'!~'", "'by'", "'without'", null, null, 
                            "'let'", "'in'", "'and'", "'or'", "'unless'" ];
    static symbolicNames = [ null, null, null, null, null, null, null, null, 
                             null, null, null, null, "MATCH_EQ", "MATCH_NEQ", 
                             "MATCH_RE", "MATCH_NRE", "BY", "WITHOUT", "AGGREGATION_OP", 
                             "DURATION", "LET", "IN", "AND", "OR", "UNLESS", 
                             "NUMBER", "DURATION_UNIT", "IDENTIFIER", "STRING", 
                             "COMMENT", "MULTILINE_COMMENT", "WS" ];
    static ruleNames = [ "query", "letExpression", "letBindings", "letBinding", 
                         "expression", "pipeline", "pipelineStep", "selector", 
                         "labelMatcher", "aggregation", "aligner" ];

    constructor(input) {
        super(input);
        this._interp = new antlr4.atn.ParserATNSimulator(this, atn, decisionsToDFA, sharedContextCache);
        this.ruleNames = PromQLPlusParser.ruleNames;
        this.literalNames = PromQLPlusParser.literalNames;
        this.symbolicNames = PromQLPlusParser.symbolicNames;
    }



	query() {
	    let localctx = new QueryContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 0, PromQLPlusParser.RULE_query);
	    try {
	        this.state = 28;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 20:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 22;
	            this.letExpression();
	            this.state = 23;
	            this.match(PromQLPlusParser.EOF);
	            break;
	        case -1:
	        case 7:
	        case 8:
	        case 27:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 25;
	            this.pipeline();
	            this.state = 26;
	            this.match(PromQLPlusParser.EOF);
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	letExpression() {
	    let localctx = new LetExpressionContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 2, PromQLPlusParser.RULE_letExpression);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 30;
	        this.match(PromQLPlusParser.LET);
	        this.state = 31;
	        this.letBindings();
	        this.state = 32;
	        this.match(PromQLPlusParser.IN);
	        this.state = 33;
	        this.expression();
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	letBindings() {
	    let localctx = new LetBindingsContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 4, PromQLPlusParser.RULE_letBindings);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 35;
	        this.letBinding();
	        this.state = 40;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===1) {
	            this.state = 36;
	            this.match(PromQLPlusParser.T__0);
	            this.state = 37;
	            this.letBinding();
	            this.state = 42;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	letBinding() {
	    let localctx = new LetBindingContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 6, PromQLPlusParser.RULE_letBinding);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 43;
	        this.match(PromQLPlusParser.IDENTIFIER);
	        this.state = 44;
	        this.match(PromQLPlusParser.MATCH_EQ);
	        this.state = 45;
	        this.pipeline();
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	expression() {
	    let localctx = new ExpressionContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 8, PromQLPlusParser.RULE_expression);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 47;
	        this.match(PromQLPlusParser.IDENTIFIER);
	        this.state = 48;
	        _la = this._input.LA(1);
	        if(!((((_la) & ~0x1f) === 0 && ((1 << _la) & 124) !== 0))) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	        this.state = 49;
	        this.match(PromQLPlusParser.IDENTIFIER);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	pipeline() {
	    let localctx = new PipelineContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 10, PromQLPlusParser.RULE_pipeline);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 51;
	        this.selector();
	        this.state = 56;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===7) {
	            this.state = 52;
	            this.match(PromQLPlusParser.T__6);
	            this.state = 53;
	            this.pipelineStep();
	            this.state = 58;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	pipelineStep() {
	    let localctx = new PipelineStepContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 12, PromQLPlusParser.RULE_pipelineStep);
	    try {
	        this.state = 61;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 18:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 59;
	            this.aggregation();
	            break;
	        case 27:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 60;
	            this.aligner();
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	selector() {
	    let localctx = new SelectorContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 14, PromQLPlusParser.RULE_selector);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 64;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===27) {
	            this.state = 63;
	            this.match(PromQLPlusParser.IDENTIFIER);
	        }

	        this.state = 78;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===8) {
	            this.state = 66;
	            this.match(PromQLPlusParser.T__7);
	            this.state = 75;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===27) {
	                this.state = 67;
	                this.labelMatcher();
	                this.state = 72;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	                while(_la===1) {
	                    this.state = 68;
	                    this.match(PromQLPlusParser.T__0);
	                    this.state = 69;
	                    this.labelMatcher();
	                    this.state = 74;
	                    this._errHandler.sync(this);
	                    _la = this._input.LA(1);
	                }
	            }

	            this.state = 77;
	            this.match(PromQLPlusParser.T__8);
	        }

	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	labelMatcher() {
	    let localctx = new LabelMatcherContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 16, PromQLPlusParser.RULE_labelMatcher);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 80;
	        this.match(PromQLPlusParser.IDENTIFIER);
	        this.state = 81;
	        _la = this._input.LA(1);
	        if(!((((_la) & ~0x1f) === 0 && ((1 << _la) & 61440) !== 0))) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	        this.state = 82;
	        this.match(PromQLPlusParser.STRING);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	aggregation() {
	    let localctx = new AggregationContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 18, PromQLPlusParser.RULE_aggregation);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 84;
	        this.match(PromQLPlusParser.AGGREGATION_OP);
	        this.state = 85;
	        _la = this._input.LA(1);
	        if(!(_la===16 || _la===17)) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	        this.state = 98;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===10) {
	            this.state = 86;
	            this.match(PromQLPlusParser.T__9);
	            this.state = 95;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===27) {
	                this.state = 87;
	                this.match(PromQLPlusParser.IDENTIFIER);
	                this.state = 92;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	                while(_la===1) {
	                    this.state = 88;
	                    this.match(PromQLPlusParser.T__0);
	                    this.state = 89;
	                    this.match(PromQLPlusParser.IDENTIFIER);
	                    this.state = 94;
	                    this._errHandler.sync(this);
	                    _la = this._input.LA(1);
	                }
	            }

	            this.state = 97;
	            this.match(PromQLPlusParser.T__10);
	        }

	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	aligner() {
	    let localctx = new AlignerContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 20, PromQLPlusParser.RULE_aligner);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 100;
	        this.match(PromQLPlusParser.IDENTIFIER);
	        this.state = 101;
	        this.match(PromQLPlusParser.DURATION);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}


}

PromQLPlusParser.EOF = antlr4.Token.EOF;
PromQLPlusParser.T__0 = 1;
PromQLPlusParser.T__1 = 2;
PromQLPlusParser.T__2 = 3;
PromQLPlusParser.T__3 = 4;
PromQLPlusParser.T__4 = 5;
PromQLPlusParser.T__5 = 6;
PromQLPlusParser.T__6 = 7;
PromQLPlusParser.T__7 = 8;
PromQLPlusParser.T__8 = 9;
PromQLPlusParser.T__9 = 10;
PromQLPlusParser.T__10 = 11;
PromQLPlusParser.MATCH_EQ = 12;
PromQLPlusParser.MATCH_NEQ = 13;
PromQLPlusParser.MATCH_RE = 14;
PromQLPlusParser.MATCH_NRE = 15;
PromQLPlusParser.BY = 16;
PromQLPlusParser.WITHOUT = 17;
PromQLPlusParser.AGGREGATION_OP = 18;
PromQLPlusParser.DURATION = 19;
PromQLPlusParser.LET = 20;
PromQLPlusParser.IN = 21;
PromQLPlusParser.AND = 22;
PromQLPlusParser.OR = 23;
PromQLPlusParser.UNLESS = 24;
PromQLPlusParser.NUMBER = 25;
PromQLPlusParser.DURATION_UNIT = 26;
PromQLPlusParser.IDENTIFIER = 27;
PromQLPlusParser.STRING = 28;
PromQLPlusParser.COMMENT = 29;
PromQLPlusParser.MULTILINE_COMMENT = 30;
PromQLPlusParser.WS = 31;

PromQLPlusParser.RULE_query = 0;
PromQLPlusParser.RULE_letExpression = 1;
PromQLPlusParser.RULE_letBindings = 2;
PromQLPlusParser.RULE_letBinding = 3;
PromQLPlusParser.RULE_expression = 4;
PromQLPlusParser.RULE_pipeline = 5;
PromQLPlusParser.RULE_pipelineStep = 6;
PromQLPlusParser.RULE_selector = 7;
PromQLPlusParser.RULE_labelMatcher = 8;
PromQLPlusParser.RULE_aggregation = 9;
PromQLPlusParser.RULE_aligner = 10;

class QueryContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = PromQLPlusParser.RULE_query;
    }

	letExpression() {
	    return this.getTypedRuleContext(LetExpressionContext,0);
	};

	EOF() {
	    return this.getToken(PromQLPlusParser.EOF, 0);
	};

	pipeline() {
	    return this.getTypedRuleContext(PipelineContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.enterQuery(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.exitQuery(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof PromQLPlusVisitor ) {
	        return visitor.visitQuery(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class LetExpressionContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = PromQLPlusParser.RULE_letExpression;
    }

	LET() {
	    return this.getToken(PromQLPlusParser.LET, 0);
	};

	letBindings() {
	    return this.getTypedRuleContext(LetBindingsContext,0);
	};

	IN() {
	    return this.getToken(PromQLPlusParser.IN, 0);
	};

	expression() {
	    return this.getTypedRuleContext(ExpressionContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.enterLetExpression(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.exitLetExpression(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof PromQLPlusVisitor ) {
	        return visitor.visitLetExpression(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class LetBindingsContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = PromQLPlusParser.RULE_letBindings;
    }

	letBinding = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(LetBindingContext);
	    } else {
	        return this.getTypedRuleContext(LetBindingContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.enterLetBindings(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.exitLetBindings(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof PromQLPlusVisitor ) {
	        return visitor.visitLetBindings(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class LetBindingContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = PromQLPlusParser.RULE_letBinding;
    }

	IDENTIFIER() {
	    return this.getToken(PromQLPlusParser.IDENTIFIER, 0);
	};

	MATCH_EQ() {
	    return this.getToken(PromQLPlusParser.MATCH_EQ, 0);
	};

	pipeline() {
	    return this.getTypedRuleContext(PipelineContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.enterLetBinding(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.exitLetBinding(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof PromQLPlusVisitor ) {
	        return visitor.visitLetBinding(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class ExpressionContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = PromQLPlusParser.RULE_expression;
    }

	IDENTIFIER = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(PromQLPlusParser.IDENTIFIER);
	    } else {
	        return this.getToken(PromQLPlusParser.IDENTIFIER, i);
	    }
	};


	enterRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.enterExpression(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.exitExpression(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof PromQLPlusVisitor ) {
	        return visitor.visitExpression(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class PipelineContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = PromQLPlusParser.RULE_pipeline;
    }

	selector() {
	    return this.getTypedRuleContext(SelectorContext,0);
	};

	pipelineStep = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(PipelineStepContext);
	    } else {
	        return this.getTypedRuleContext(PipelineStepContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.enterPipeline(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.exitPipeline(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof PromQLPlusVisitor ) {
	        return visitor.visitPipeline(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class PipelineStepContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = PromQLPlusParser.RULE_pipelineStep;
    }

	aggregation() {
	    return this.getTypedRuleContext(AggregationContext,0);
	};

	aligner() {
	    return this.getTypedRuleContext(AlignerContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.enterPipelineStep(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.exitPipelineStep(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof PromQLPlusVisitor ) {
	        return visitor.visitPipelineStep(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class SelectorContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = PromQLPlusParser.RULE_selector;
    }

	IDENTIFIER() {
	    return this.getToken(PromQLPlusParser.IDENTIFIER, 0);
	};

	labelMatcher = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(LabelMatcherContext);
	    } else {
	        return this.getTypedRuleContext(LabelMatcherContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.enterSelector(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.exitSelector(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof PromQLPlusVisitor ) {
	        return visitor.visitSelector(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class LabelMatcherContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = PromQLPlusParser.RULE_labelMatcher;
    }

	IDENTIFIER() {
	    return this.getToken(PromQLPlusParser.IDENTIFIER, 0);
	};

	STRING() {
	    return this.getToken(PromQLPlusParser.STRING, 0);
	};

	MATCH_EQ() {
	    return this.getToken(PromQLPlusParser.MATCH_EQ, 0);
	};

	MATCH_NEQ() {
	    return this.getToken(PromQLPlusParser.MATCH_NEQ, 0);
	};

	MATCH_RE() {
	    return this.getToken(PromQLPlusParser.MATCH_RE, 0);
	};

	MATCH_NRE() {
	    return this.getToken(PromQLPlusParser.MATCH_NRE, 0);
	};

	enterRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.enterLabelMatcher(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.exitLabelMatcher(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof PromQLPlusVisitor ) {
	        return visitor.visitLabelMatcher(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class AggregationContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = PromQLPlusParser.RULE_aggregation;
    }

	AGGREGATION_OP() {
	    return this.getToken(PromQLPlusParser.AGGREGATION_OP, 0);
	};

	BY() {
	    return this.getToken(PromQLPlusParser.BY, 0);
	};

	WITHOUT() {
	    return this.getToken(PromQLPlusParser.WITHOUT, 0);
	};

	IDENTIFIER = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(PromQLPlusParser.IDENTIFIER);
	    } else {
	        return this.getToken(PromQLPlusParser.IDENTIFIER, i);
	    }
	};


	enterRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.enterAggregation(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.exitAggregation(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof PromQLPlusVisitor ) {
	        return visitor.visitAggregation(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class AlignerContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = PromQLPlusParser.RULE_aligner;
    }

	IDENTIFIER() {
	    return this.getToken(PromQLPlusParser.IDENTIFIER, 0);
	};

	DURATION() {
	    return this.getToken(PromQLPlusParser.DURATION, 0);
	};

	enterRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.enterAligner(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof PromQLPlusListener ) {
	        listener.exitAligner(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof PromQLPlusVisitor ) {
	        return visitor.visitAligner(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}




PromQLPlusParser.QueryContext = QueryContext; 
PromQLPlusParser.LetExpressionContext = LetExpressionContext; 
PromQLPlusParser.LetBindingsContext = LetBindingsContext; 
PromQLPlusParser.LetBindingContext = LetBindingContext; 
PromQLPlusParser.ExpressionContext = ExpressionContext; 
PromQLPlusParser.PipelineContext = PipelineContext; 
PromQLPlusParser.PipelineStepContext = PipelineStepContext; 
PromQLPlusParser.SelectorContext = SelectorContext; 
PromQLPlusParser.LabelMatcherContext = LabelMatcherContext; 
PromQLPlusParser.AggregationContext = AggregationContext; 
PromQLPlusParser.AlignerContext = AlignerContext; 
