// Generated from grammar/Flow.g4 by ANTLR 4.13.2
// jshint ignore: start
import antlr4 from 'antlr4';
import FlowListener from './FlowListener.js';
import FlowVisitor from './FlowVisitor.js';

const serializedATN = [4,1,30,119,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,
4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,10,7,10,2,11,7,11,1,0,1,0,1,
0,1,0,1,0,1,0,3,0,31,8,0,1,1,1,1,1,1,1,1,5,1,37,8,1,10,1,12,1,40,9,1,1,1,
1,1,1,1,1,2,1,2,1,2,1,2,1,3,1,3,1,3,1,3,5,3,53,8,3,10,3,12,3,56,9,3,1,4,
1,4,1,5,1,5,1,5,1,5,1,5,3,5,65,8,5,1,6,1,6,1,6,5,6,70,8,6,10,6,12,6,73,9,
6,1,7,1,7,3,7,77,8,7,1,8,3,8,80,8,8,1,8,1,8,1,8,1,8,5,8,86,8,8,10,8,12,8,
89,9,8,3,8,91,8,8,1,8,3,8,94,8,8,1,9,1,9,1,9,1,9,1,10,1,10,1,10,1,10,1,10,
1,10,5,10,106,8,10,10,10,12,10,109,9,10,3,10,111,8,10,1,10,3,10,114,8,10,
1,11,1,11,1,11,1,11,0,0,12,0,2,4,6,8,10,12,14,16,18,20,22,0,3,1,0,11,14,
1,0,7,10,1,0,15,16,119,0,30,1,0,0,0,2,32,1,0,0,0,4,44,1,0,0,0,6,48,1,0,0,
0,8,57,1,0,0,0,10,64,1,0,0,0,12,66,1,0,0,0,14,76,1,0,0,0,16,79,1,0,0,0,18,
95,1,0,0,0,20,99,1,0,0,0,22,115,1,0,0,0,24,25,3,2,1,0,25,26,5,0,0,1,26,31,
1,0,0,0,27,28,3,12,6,0,28,29,5,0,0,1,29,31,1,0,0,0,30,24,1,0,0,0,30,27,1,
0,0,0,31,1,1,0,0,0,32,33,5,19,0,0,33,38,3,4,2,0,34,35,5,1,0,0,35,37,3,4,
2,0,36,34,1,0,0,0,37,40,1,0,0,0,38,36,1,0,0,0,38,39,1,0,0,0,39,41,1,0,0,
0,40,38,1,0,0,0,41,42,5,20,0,0,42,43,3,6,3,0,43,3,1,0,0,0,44,45,5,26,0,0,
45,46,5,7,0,0,46,47,3,12,6,0,47,5,1,0,0,0,48,54,3,10,5,0,49,50,3,8,4,0,50,
51,3,10,5,0,51,53,1,0,0,0,52,49,1,0,0,0,53,56,1,0,0,0,54,52,1,0,0,0,54,55,
1,0,0,0,55,7,1,0,0,0,56,54,1,0,0,0,57,58,7,0,0,0,58,9,1,0,0,0,59,65,5,26,
0,0,60,61,5,2,0,0,61,62,3,6,3,0,62,63,5,3,0,0,63,65,1,0,0,0,64,59,1,0,0,
0,64,60,1,0,0,0,65,11,1,0,0,0,66,71,3,16,8,0,67,68,5,4,0,0,68,70,3,14,7,
0,69,67,1,0,0,0,70,73,1,0,0,0,71,69,1,0,0,0,71,72,1,0,0,0,72,13,1,0,0,0,
73,71,1,0,0,0,74,77,3,20,10,0,75,77,3,22,11,0,76,74,1,0,0,0,76,75,1,0,0,
0,77,15,1,0,0,0,78,80,5,26,0,0,79,78,1,0,0,0,79,80,1,0,0,0,80,93,1,0,0,0,
81,90,5,5,0,0,82,87,3,18,9,0,83,84,5,1,0,0,84,86,3,18,9,0,85,83,1,0,0,0,
86,89,1,0,0,0,87,85,1,0,0,0,87,88,1,0,0,0,88,91,1,0,0,0,89,87,1,0,0,0,90,
82,1,0,0,0,90,91,1,0,0,0,91,92,1,0,0,0,92,94,5,6,0,0,93,81,1,0,0,0,93,94,
1,0,0,0,94,17,1,0,0,0,95,96,5,26,0,0,96,97,7,1,0,0,97,98,5,27,0,0,98,19,
1,0,0,0,99,100,5,17,0,0,100,113,7,2,0,0,101,110,5,2,0,0,102,107,5,26,0,0,
103,104,5,1,0,0,104,106,5,26,0,0,105,103,1,0,0,0,106,109,1,0,0,0,107,105,
1,0,0,0,107,108,1,0,0,0,108,111,1,0,0,0,109,107,1,0,0,0,110,102,1,0,0,0,
110,111,1,0,0,0,111,112,1,0,0,0,112,114,5,3,0,0,113,101,1,0,0,0,113,114,
1,0,0,0,114,21,1,0,0,0,115,116,5,26,0,0,116,117,5,18,0,0,117,23,1,0,0,0,
13,30,38,54,64,71,76,79,87,90,93,107,110,113];


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
                         "aligner" ];

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
	        this.state = 30;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 19:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 24;
	            this.letExpression();
	            this.state = 25;
	            this.match(FlowParser.EOF);
	            break;
	        case -1:
	        case 4:
	        case 5:
	        case 26:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 27;
	            this.pipeline();
	            this.state = 28;
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
	        this.state = 32;
	        this.match(FlowParser.LET);
	        this.state = 33;
	        this.letBinding();
	        this.state = 38;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===1) {
	            this.state = 34;
	            this.match(FlowParser.T__0);
	            this.state = 35;
	            this.letBinding();
	            this.state = 40;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 41;
	        this.match(FlowParser.IN);
	        this.state = 42;
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
	        this.state = 44;
	        this.match(FlowParser.IDENTIFIER);
	        this.state = 45;
	        this.match(FlowParser.MATCH_EQ);
	        this.state = 46;
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
	        this.state = 48;
	        this.primaryExpression();
	        this.state = 54;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while((((_la) & ~0x1f) === 0 && ((1 << _la) & 30720) !== 0)) {
	            this.state = 49;
	            this.binaryOperator();
	            this.state = 50;
	            this.primaryExpression();
	            this.state = 56;
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
	        this.state = 57;
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
	        this.state = 64;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 26:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 59;
	            this.match(FlowParser.IDENTIFIER);
	            break;
	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 60;
	            this.match(FlowParser.T__1);
	            this.state = 61;
	            this.binaryExpression();
	            this.state = 62;
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
	        this.state = 66;
	        this.selector();
	        this.state = 71;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===4) {
	            this.state = 67;
	            this.match(FlowParser.T__3);
	            this.state = 68;
	            this.pipelineStep();
	            this.state = 73;
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
	        this.state = 76;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 17:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 74;
	            this.aggregation();
	            break;
	        case 26:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 75;
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
	    this.enterRule(localctx, 16, FlowParser.RULE_selector);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 79;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===26) {
	            this.state = 78;
	            this.match(FlowParser.IDENTIFIER);
	        }

	        this.state = 93;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===5) {
	            this.state = 81;
	            this.match(FlowParser.T__4);
	            this.state = 90;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===26) {
	                this.state = 82;
	                this.labelMatcher();
	                this.state = 87;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	                while(_la===1) {
	                    this.state = 83;
	                    this.match(FlowParser.T__0);
	                    this.state = 84;
	                    this.labelMatcher();
	                    this.state = 89;
	                    this._errHandler.sync(this);
	                    _la = this._input.LA(1);
	                }
	            }

	            this.state = 92;
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
	        this.state = 95;
	        this.match(FlowParser.IDENTIFIER);
	        this.state = 96;
	        _la = this._input.LA(1);
	        if(!((((_la) & ~0x1f) === 0 && ((1 << _la) & 1920) !== 0))) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	        this.state = 97;
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
	        this.state = 99;
	        this.match(FlowParser.AGGREGATION_OP);
	        this.state = 100;
	        _la = this._input.LA(1);
	        if(!(_la===15 || _la===16)) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	        this.state = 113;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===2) {
	            this.state = 101;
	            this.match(FlowParser.T__1);
	            this.state = 110;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===26) {
	                this.state = 102;
	                this.match(FlowParser.IDENTIFIER);
	                this.state = 107;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	                while(_la===1) {
	                    this.state = 103;
	                    this.match(FlowParser.T__0);
	                    this.state = 104;
	                    this.match(FlowParser.IDENTIFIER);
	                    this.state = 109;
	                    this._errHandler.sync(this);
	                    _la = this._input.LA(1);
	                }
	            }

	            this.state = 112;
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



	aligner() {
	    let localctx = new AlignerContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 22, FlowParser.RULE_aligner);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 115;
	        this.match(FlowParser.IDENTIFIER);
	        this.state = 116;
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
FlowParser.RULE_aligner = 11;

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
FlowParser.AlignerContext = AlignerContext; 
