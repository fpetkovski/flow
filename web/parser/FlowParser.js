// Generated from grammar/Flow.g4 by ANTLR 4.13.2
// jshint ignore: start
import antlr4 from 'antlr4';
import FlowListener from './FlowListener.js';
import FlowVisitor from './FlowVisitor.js';

const serializedATN = [4,1,30,125,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,
4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,10,7,10,2,11,7,11,2,12,7,12,
1,0,1,0,1,0,1,0,1,0,1,0,3,0,33,8,0,1,1,1,1,1,1,1,1,5,1,39,8,1,10,1,12,1,
42,9,1,1,1,1,1,1,1,1,2,1,2,1,2,1,2,1,3,1,3,1,3,1,3,5,3,55,8,3,10,3,12,3,
58,9,3,1,4,1,4,1,5,1,5,1,5,1,5,1,5,1,5,3,5,68,8,5,1,6,1,6,1,6,5,6,73,8,6,
10,6,12,6,76,9,6,1,7,1,7,1,7,3,7,81,8,7,1,8,3,8,84,8,8,1,8,1,8,1,8,1,8,5,
8,90,8,8,10,8,12,8,93,9,8,3,8,95,8,8,1,8,3,8,98,8,8,1,9,1,9,1,9,1,9,1,10,
1,10,1,10,1,10,1,10,1,10,5,10,110,8,10,10,10,12,10,113,9,10,3,10,115,8,10,
1,10,3,10,118,8,10,1,11,1,11,1,12,1,12,1,12,1,12,0,0,13,0,2,4,6,8,10,12,
14,16,18,20,22,24,0,3,1,0,11,14,1,0,7,10,1,0,15,16,126,0,32,1,0,0,0,2,34,
1,0,0,0,4,46,1,0,0,0,6,50,1,0,0,0,8,59,1,0,0,0,10,67,1,0,0,0,12,69,1,0,0,
0,14,80,1,0,0,0,16,83,1,0,0,0,18,99,1,0,0,0,20,103,1,0,0,0,22,119,1,0,0,
0,24,121,1,0,0,0,26,27,3,2,1,0,27,28,5,0,0,1,28,33,1,0,0,0,29,30,3,12,6,
0,30,31,5,0,0,1,31,33,1,0,0,0,32,26,1,0,0,0,32,29,1,0,0,0,33,1,1,0,0,0,34,
35,5,19,0,0,35,40,3,4,2,0,36,37,5,1,0,0,37,39,3,4,2,0,38,36,1,0,0,0,39,42,
1,0,0,0,40,38,1,0,0,0,40,41,1,0,0,0,41,43,1,0,0,0,42,40,1,0,0,0,43,44,5,
20,0,0,44,45,3,6,3,0,45,3,1,0,0,0,46,47,5,26,0,0,47,48,5,7,0,0,48,49,3,12,
6,0,49,5,1,0,0,0,50,56,3,10,5,0,51,52,3,8,4,0,52,53,3,10,5,0,53,55,1,0,0,
0,54,51,1,0,0,0,55,58,1,0,0,0,56,54,1,0,0,0,56,57,1,0,0,0,57,7,1,0,0,0,58,
56,1,0,0,0,59,60,7,0,0,0,60,9,1,0,0,0,61,68,5,26,0,0,62,68,5,24,0,0,63,64,
5,2,0,0,64,65,3,6,3,0,65,66,5,3,0,0,66,68,1,0,0,0,67,61,1,0,0,0,67,62,1,
0,0,0,67,63,1,0,0,0,68,11,1,0,0,0,69,74,3,16,8,0,70,71,5,4,0,0,71,73,3,14,
7,0,72,70,1,0,0,0,73,76,1,0,0,0,74,72,1,0,0,0,74,75,1,0,0,0,75,13,1,0,0,
0,76,74,1,0,0,0,77,81,3,20,10,0,78,81,3,24,12,0,79,81,3,22,11,0,80,77,1,
0,0,0,80,78,1,0,0,0,80,79,1,0,0,0,81,15,1,0,0,0,82,84,5,26,0,0,83,82,1,0,
0,0,83,84,1,0,0,0,84,97,1,0,0,0,85,94,5,5,0,0,86,91,3,18,9,0,87,88,5,1,0,
0,88,90,3,18,9,0,89,87,1,0,0,0,90,93,1,0,0,0,91,89,1,0,0,0,91,92,1,0,0,0,
92,95,1,0,0,0,93,91,1,0,0,0,94,86,1,0,0,0,94,95,1,0,0,0,95,96,1,0,0,0,96,
98,5,6,0,0,97,85,1,0,0,0,97,98,1,0,0,0,98,17,1,0,0,0,99,100,5,26,0,0,100,
101,7,1,0,0,101,102,5,27,0,0,102,19,1,0,0,0,103,104,5,17,0,0,104,117,7,2,
0,0,105,114,5,2,0,0,106,111,5,26,0,0,107,108,5,1,0,0,108,110,5,26,0,0,109,
107,1,0,0,0,110,113,1,0,0,0,111,109,1,0,0,0,111,112,1,0,0,0,112,115,1,0,
0,0,113,111,1,0,0,0,114,106,1,0,0,0,114,115,1,0,0,0,115,116,1,0,0,0,116,
118,5,3,0,0,117,105,1,0,0,0,117,118,1,0,0,0,118,21,1,0,0,0,119,120,5,26,
0,0,120,23,1,0,0,0,121,122,5,26,0,0,122,123,5,18,0,0,123,25,1,0,0,0,13,32,
40,56,67,74,80,83,91,94,97,111,114,117];


const atn = new antlr4.atn.ATNDeserializer().deserialize(serializedATN);

const decisionsToDFA = atn.decisionToState.map( (ds, index) => new antlr4.dfa.DFA(ds, index) );

const sharedContextCache = new antlr4.atn.PredictionContextCache();

export default class FlowParser extends antlr4.Parser {

    static grammarFileName = "Flow.g4";
    static literalNames = [ null, "','", "'('", "')'", "'|'", "'{'", "'}'", 
                            "'='", "'!='", "'=~'", "'!~'", "'+'", "'-'", 
                            "'*'", "'/'", "'by'", "'without'", null, null, 
                            "'let'", "'in'", "'and'", "'or'", "'unless'" ];
    static symbolicNames = [ null, null, null, null, null, null, null, "MATCH_EQ", 
                             "MATCH_NEQ", "MATCH_RE", "MATCH_NRE", "OP_ADD", 
                             "OP_SUB", "OP_MUL", "OP_DIV", "BY", "WITHOUT", 
                             "AGGREGATION_OP", "DURATION", "LET", "IN", 
                             "AND", "OR", "UNLESS", "NUMBER", "DURATION_UNIT", 
                             "IDENTIFIER", "STRING", "COMMENT", "MULTILINE_COMMENT", 
                             "WS" ];
    static ruleNames = [ "query", "letExpression", "letBinding", "binaryExpression", 
                         "binaryOperator", "primaryExpression", "pipeline", 
                         "pipelineStep", "selector", "labelMatcher", "aggregation", 
                         "function", "aligner" ];

    constructor(input) {
        super(input);
        this._interp = new antlr4.atn.ParserATNSimulator(this, atn, decisionsToDFA, sharedContextCache);
        this.ruleNames = FlowParser.ruleNames;
        this.literalNames = FlowParser.literalNames;
        this.symbolicNames = FlowParser.symbolicNames;
    }



	query() {
	    let localctx = new QueryContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 0, FlowParser.RULE_query);
	    try {
	        this.state = 32;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 19:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 26;
	            this.letExpression();
	            this.state = 27;
	            this.match(FlowParser.EOF);
	            break;
	        case -1:
	        case 4:
	        case 5:
	        case 26:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 29;
	            this.pipeline();
	            this.state = 30;
	            this.match(FlowParser.EOF);
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
	    this.enterRule(localctx, 2, FlowParser.RULE_letExpression);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 34;
	        this.match(FlowParser.LET);
	        this.state = 35;
	        this.letBinding();
	        this.state = 40;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===1) {
	            this.state = 36;
	            this.match(FlowParser.T__0);
	            this.state = 37;
	            this.letBinding();
	            this.state = 42;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 43;
	        this.match(FlowParser.IN);
	        this.state = 44;
	        this.binaryExpression();
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
	    this.enterRule(localctx, 4, FlowParser.RULE_letBinding);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 46;
	        this.match(FlowParser.IDENTIFIER);
	        this.state = 47;
	        this.match(FlowParser.MATCH_EQ);
	        this.state = 48;
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



	binaryExpression() {
	    let localctx = new BinaryExpressionContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 6, FlowParser.RULE_binaryExpression);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 50;
	        this.primaryExpression();
	        this.state = 56;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while((((_la) & ~0x1f) === 0 && ((1 << _la) & 30720) !== 0)) {
	            this.state = 51;
	            this.binaryOperator();
	            this.state = 52;
	            this.primaryExpression();
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



	binaryOperator() {
	    let localctx = new BinaryOperatorContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 8, FlowParser.RULE_binaryOperator);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 59;
	        _la = this._input.LA(1);
	        if(!((((_la) & ~0x1f) === 0 && ((1 << _la) & 30720) !== 0))) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
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



	primaryExpression() {
	    let localctx = new PrimaryExpressionContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 10, FlowParser.RULE_primaryExpression);
	    try {
	        this.state = 67;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 26:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 61;
	            this.match(FlowParser.IDENTIFIER);
	            break;
	        case 24:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 62;
	            this.match(FlowParser.NUMBER);
	            break;
	        case 2:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 63;
	            this.match(FlowParser.T__1);
	            this.state = 64;
	            this.binaryExpression();
	            this.state = 65;
	            this.match(FlowParser.T__2);
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



	pipeline() {
	    let localctx = new PipelineContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 12, FlowParser.RULE_pipeline);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 69;
	        this.selector();
	        this.state = 74;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===4) {
	            this.state = 70;
	            this.match(FlowParser.T__3);
	            this.state = 71;
	            this.pipelineStep();
	            this.state = 76;
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
	    this.enterRule(localctx, 14, FlowParser.RULE_pipelineStep);
	    try {
	        this.state = 80;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,5,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 77;
	            this.aggregation();
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 78;
	            this.aligner();
	            break;

	        case 3:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 79;
	            this.function_();
	            break;

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
	    this.enterRule(localctx, 16, FlowParser.RULE_selector);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 83;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===26) {
	            this.state = 82;
	            this.match(FlowParser.IDENTIFIER);
	        }

	        this.state = 97;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===5) {
	            this.state = 85;
	            this.match(FlowParser.T__4);
	            this.state = 94;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===26) {
	                this.state = 86;
	                this.labelMatcher();
	                this.state = 91;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	                while(_la===1) {
	                    this.state = 87;
	                    this.match(FlowParser.T__0);
	                    this.state = 88;
	                    this.labelMatcher();
	                    this.state = 93;
	                    this._errHandler.sync(this);
	                    _la = this._input.LA(1);
	                }
	            }

	            this.state = 96;
	            this.match(FlowParser.T__5);
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
	    this.enterRule(localctx, 18, FlowParser.RULE_labelMatcher);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 99;
	        this.match(FlowParser.IDENTIFIER);
	        this.state = 100;
	        _la = this._input.LA(1);
	        if(!((((_la) & ~0x1f) === 0 && ((1 << _la) & 1920) !== 0))) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	        this.state = 101;
	        this.match(FlowParser.STRING);
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
	    this.enterRule(localctx, 20, FlowParser.RULE_aggregation);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 103;
	        this.match(FlowParser.AGGREGATION_OP);
	        this.state = 104;
	        _la = this._input.LA(1);
	        if(!(_la===15 || _la===16)) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	        this.state = 117;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===2) {
	            this.state = 105;
	            this.match(FlowParser.T__1);
	            this.state = 114;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===26) {
	                this.state = 106;
	                this.match(FlowParser.IDENTIFIER);
	                this.state = 111;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	                while(_la===1) {
	                    this.state = 107;
	                    this.match(FlowParser.T__0);
	                    this.state = 108;
	                    this.match(FlowParser.IDENTIFIER);
	                    this.state = 113;
	                    this._errHandler.sync(this);
	                    _la = this._input.LA(1);
	                }
	            }

	            this.state = 116;
	            this.match(FlowParser.T__2);
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



	function_() {
	    let localctx = new FunctionContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 22, FlowParser.RULE_function);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 119;
	        this.match(FlowParser.IDENTIFIER);
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
	    this.enterRule(localctx, 24, FlowParser.RULE_aligner);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 121;
	        this.match(FlowParser.IDENTIFIER);
	        this.state = 122;
	        this.match(FlowParser.DURATION);
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

FlowParser.EOF = antlr4.Token.EOF;
FlowParser.T__0 = 1;
FlowParser.T__1 = 2;
FlowParser.T__2 = 3;
FlowParser.T__3 = 4;
FlowParser.T__4 = 5;
FlowParser.T__5 = 6;
FlowParser.MATCH_EQ = 7;
FlowParser.MATCH_NEQ = 8;
FlowParser.MATCH_RE = 9;
FlowParser.MATCH_NRE = 10;
FlowParser.OP_ADD = 11;
FlowParser.OP_SUB = 12;
FlowParser.OP_MUL = 13;
FlowParser.OP_DIV = 14;
FlowParser.BY = 15;
FlowParser.WITHOUT = 16;
FlowParser.AGGREGATION_OP = 17;
FlowParser.DURATION = 18;
FlowParser.LET = 19;
FlowParser.IN = 20;
FlowParser.AND = 21;
FlowParser.OR = 22;
FlowParser.UNLESS = 23;
FlowParser.NUMBER = 24;
FlowParser.DURATION_UNIT = 25;
FlowParser.IDENTIFIER = 26;
FlowParser.STRING = 27;
FlowParser.COMMENT = 28;
FlowParser.MULTILINE_COMMENT = 29;
FlowParser.WS = 30;

FlowParser.RULE_query = 0;
FlowParser.RULE_letExpression = 1;
FlowParser.RULE_letBinding = 2;
FlowParser.RULE_binaryExpression = 3;
FlowParser.RULE_binaryOperator = 4;
FlowParser.RULE_primaryExpression = 5;
FlowParser.RULE_pipeline = 6;
FlowParser.RULE_pipelineStep = 7;
FlowParser.RULE_selector = 8;
FlowParser.RULE_labelMatcher = 9;
FlowParser.RULE_aggregation = 10;
FlowParser.RULE_function = 11;
FlowParser.RULE_aligner = 12;

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
        this.ruleIndex = FlowParser.RULE_query;
    }

	letExpression() {
	    return this.getTypedRuleContext(LetExpressionContext,0);
	};

	EOF() {
	    return this.getToken(FlowParser.EOF, 0);
	};

	pipeline() {
	    return this.getTypedRuleContext(PipelineContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.enterQuery(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.exitQuery(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof FlowVisitor ) {
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
        this.ruleIndex = FlowParser.RULE_letExpression;
    }

	LET() {
	    return this.getToken(FlowParser.LET, 0);
	};

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

	IN() {
	    return this.getToken(FlowParser.IN, 0);
	};

	binaryExpression() {
	    return this.getTypedRuleContext(BinaryExpressionContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.enterLetExpression(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.exitLetExpression(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof FlowVisitor ) {
	        return visitor.visitLetExpression(this);
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
        this.ruleIndex = FlowParser.RULE_letBinding;
    }

	IDENTIFIER() {
	    return this.getToken(FlowParser.IDENTIFIER, 0);
	};

	MATCH_EQ() {
	    return this.getToken(FlowParser.MATCH_EQ, 0);
	};

	pipeline() {
	    return this.getTypedRuleContext(PipelineContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.enterLetBinding(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.exitLetBinding(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof FlowVisitor ) {
	        return visitor.visitLetBinding(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class BinaryExpressionContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = FlowParser.RULE_binaryExpression;
    }

	primaryExpression = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(PrimaryExpressionContext);
	    } else {
	        return this.getTypedRuleContext(PrimaryExpressionContext,i);
	    }
	};

	binaryOperator = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(BinaryOperatorContext);
	    } else {
	        return this.getTypedRuleContext(BinaryOperatorContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.enterBinaryExpression(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.exitBinaryExpression(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof FlowVisitor ) {
	        return visitor.visitBinaryExpression(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class BinaryOperatorContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = FlowParser.RULE_binaryOperator;
    }

	OP_ADD() {
	    return this.getToken(FlowParser.OP_ADD, 0);
	};

	OP_SUB() {
	    return this.getToken(FlowParser.OP_SUB, 0);
	};

	OP_MUL() {
	    return this.getToken(FlowParser.OP_MUL, 0);
	};

	OP_DIV() {
	    return this.getToken(FlowParser.OP_DIV, 0);
	};

	enterRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.enterBinaryOperator(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.exitBinaryOperator(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof FlowVisitor ) {
	        return visitor.visitBinaryOperator(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class PrimaryExpressionContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = FlowParser.RULE_primaryExpression;
    }

	IDENTIFIER() {
	    return this.getToken(FlowParser.IDENTIFIER, 0);
	};

	NUMBER() {
	    return this.getToken(FlowParser.NUMBER, 0);
	};

	binaryExpression() {
	    return this.getTypedRuleContext(BinaryExpressionContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.enterPrimaryExpression(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.exitPrimaryExpression(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof FlowVisitor ) {
	        return visitor.visitPrimaryExpression(this);
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
        this.ruleIndex = FlowParser.RULE_pipeline;
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
	    if(listener instanceof FlowListener ) {
	        listener.enterPipeline(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.exitPipeline(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof FlowVisitor ) {
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
        this.ruleIndex = FlowParser.RULE_pipelineStep;
    }

	aggregation() {
	    return this.getTypedRuleContext(AggregationContext,0);
	};

	aligner() {
	    return this.getTypedRuleContext(AlignerContext,0);
	};

	function_() {
	    return this.getTypedRuleContext(FunctionContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.enterPipelineStep(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.exitPipelineStep(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof FlowVisitor ) {
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
        this.ruleIndex = FlowParser.RULE_selector;
    }

	IDENTIFIER() {
	    return this.getToken(FlowParser.IDENTIFIER, 0);
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
	    if(listener instanceof FlowListener ) {
	        listener.enterSelector(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.exitSelector(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof FlowVisitor ) {
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
        this.ruleIndex = FlowParser.RULE_labelMatcher;
    }

	IDENTIFIER() {
	    return this.getToken(FlowParser.IDENTIFIER, 0);
	};

	STRING() {
	    return this.getToken(FlowParser.STRING, 0);
	};

	MATCH_EQ() {
	    return this.getToken(FlowParser.MATCH_EQ, 0);
	};

	MATCH_NEQ() {
	    return this.getToken(FlowParser.MATCH_NEQ, 0);
	};

	MATCH_RE() {
	    return this.getToken(FlowParser.MATCH_RE, 0);
	};

	MATCH_NRE() {
	    return this.getToken(FlowParser.MATCH_NRE, 0);
	};

	enterRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.enterLabelMatcher(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.exitLabelMatcher(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof FlowVisitor ) {
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
        this.ruleIndex = FlowParser.RULE_aggregation;
    }

	AGGREGATION_OP() {
	    return this.getToken(FlowParser.AGGREGATION_OP, 0);
	};

	BY() {
	    return this.getToken(FlowParser.BY, 0);
	};

	WITHOUT() {
	    return this.getToken(FlowParser.WITHOUT, 0);
	};

	IDENTIFIER = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(FlowParser.IDENTIFIER);
	    } else {
	        return this.getToken(FlowParser.IDENTIFIER, i);
	    }
	};


	enterRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.enterAggregation(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.exitAggregation(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof FlowVisitor ) {
	        return visitor.visitAggregation(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class FunctionContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = FlowParser.RULE_function;
    }

	IDENTIFIER() {
	    return this.getToken(FlowParser.IDENTIFIER, 0);
	};

	enterRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.enterFunction(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.exitFunction(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof FlowVisitor ) {
	        return visitor.visitFunction(this);
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
        this.ruleIndex = FlowParser.RULE_aligner;
    }

	IDENTIFIER() {
	    return this.getToken(FlowParser.IDENTIFIER, 0);
	};

	DURATION() {
	    return this.getToken(FlowParser.DURATION, 0);
	};

	enterRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.enterAligner(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.exitAligner(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof FlowVisitor ) {
	        return visitor.visitAligner(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}




FlowParser.QueryContext = QueryContext; 
FlowParser.LetExpressionContext = LetExpressionContext; 
FlowParser.LetBindingContext = LetBindingContext; 
FlowParser.BinaryExpressionContext = BinaryExpressionContext; 
FlowParser.BinaryOperatorContext = BinaryOperatorContext; 
FlowParser.PrimaryExpressionContext = PrimaryExpressionContext; 
FlowParser.PipelineContext = PipelineContext; 
FlowParser.PipelineStepContext = PipelineStepContext; 
FlowParser.SelectorContext = SelectorContext; 
FlowParser.LabelMatcherContext = LabelMatcherContext; 
FlowParser.AggregationContext = AggregationContext; 
FlowParser.FunctionContext = FunctionContext; 
FlowParser.AlignerContext = AlignerContext; 
