// Generated from parser/grammar/Flow.g4 by ANTLR 4.13.2
// jshint ignore: start
import antlr4 from 'antlr4';
import FlowListener from './FlowListener.js';
import FlowVisitor from './FlowVisitor.js';

const serializedATN = [4,1,34,137,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,
4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,10,7,10,2,11,7,11,2,12,7,12,
2,13,7,13,2,14,7,14,1,0,1,0,1,0,1,0,1,0,1,0,3,0,37,8,0,1,1,1,1,1,1,1,1,5,
1,43,8,1,10,1,12,1,46,9,1,1,1,1,1,1,1,1,2,1,2,1,2,1,2,1,3,1,3,3,3,57,8,3,
1,3,1,3,5,3,61,8,3,10,3,12,3,64,9,3,1,4,1,4,1,4,1,4,3,4,70,8,4,1,5,3,5,73,
8,5,1,5,1,5,1,5,1,5,5,5,79,8,5,10,5,12,5,82,9,5,3,5,84,8,5,1,5,3,5,87,8,
5,1,6,1,6,1,6,1,6,1,7,1,7,1,7,1,7,5,7,97,8,7,10,7,12,7,100,9,7,1,8,1,8,1,
8,1,8,1,8,3,8,107,8,8,1,9,1,9,1,10,1,10,1,11,1,11,1,11,1,11,1,11,1,11,5,
11,119,8,11,10,11,12,11,122,9,11,3,11,124,8,11,1,11,3,11,127,8,11,1,12,1,
12,1,12,1,13,1,13,1,14,1,14,1,14,1,14,0,0,15,0,2,4,6,8,10,12,14,16,18,20,
22,24,26,28,0,4,1,0,7,10,1,0,29,30,2,0,8,8,11,19,1,0,20,21,137,0,36,1,0,
0,0,2,38,1,0,0,0,4,50,1,0,0,0,6,56,1,0,0,0,8,69,1,0,0,0,10,72,1,0,0,0,12,
88,1,0,0,0,14,92,1,0,0,0,16,106,1,0,0,0,18,108,1,0,0,0,20,110,1,0,0,0,22,
112,1,0,0,0,24,128,1,0,0,0,26,131,1,0,0,0,28,133,1,0,0,0,30,31,3,2,1,0,31,
32,5,0,0,1,32,37,1,0,0,0,33,34,3,6,3,0,34,35,5,0,0,1,35,37,1,0,0,0,36,30,
1,0,0,0,36,33,1,0,0,0,37,1,1,0,0,0,38,39,5,24,0,0,39,44,3,4,2,0,40,41,5,
1,0,0,41,43,3,4,2,0,42,40,1,0,0,0,43,46,1,0,0,0,44,42,1,0,0,0,44,45,1,0,
0,0,45,47,1,0,0,0,46,44,1,0,0,0,47,48,5,25,0,0,48,49,3,6,3,0,49,3,1,0,0,
0,50,51,5,30,0,0,51,52,5,7,0,0,52,53,3,6,3,0,53,5,1,0,0,0,54,57,3,10,5,0,
55,57,3,14,7,0,56,54,1,0,0,0,56,55,1,0,0,0,57,62,1,0,0,0,58,59,5,2,0,0,59,
61,3,8,4,0,60,58,1,0,0,0,61,64,1,0,0,0,62,60,1,0,0,0,62,63,1,0,0,0,63,7,
1,0,0,0,64,62,1,0,0,0,65,70,3,22,11,0,66,70,3,24,12,0,67,70,3,26,13,0,68,
70,3,28,14,0,69,65,1,0,0,0,69,66,1,0,0,0,69,67,1,0,0,0,69,68,1,0,0,0,70,
9,1,0,0,0,71,73,5,30,0,0,72,71,1,0,0,0,72,73,1,0,0,0,73,86,1,0,0,0,74,83,
5,3,0,0,75,80,3,12,6,0,76,77,5,1,0,0,77,79,3,12,6,0,78,76,1,0,0,0,79,82,
1,0,0,0,80,78,1,0,0,0,80,81,1,0,0,0,81,84,1,0,0,0,82,80,1,0,0,0,83,75,1,
0,0,0,83,84,1,0,0,0,84,85,1,0,0,0,85,87,5,4,0,0,86,74,1,0,0,0,86,87,1,0,
0,0,87,11,1,0,0,0,88,89,5,30,0,0,89,90,7,0,0,0,90,91,5,31,0,0,91,13,1,0,
0,0,92,98,3,16,8,0,93,94,3,20,10,0,94,95,3,16,8,0,95,97,1,0,0,0,96,93,1,
0,0,0,97,100,1,0,0,0,98,96,1,0,0,0,98,99,1,0,0,0,99,15,1,0,0,0,100,98,1,
0,0,0,101,107,3,18,9,0,102,103,5,5,0,0,103,104,3,14,7,0,104,105,5,6,0,0,
105,107,1,0,0,0,106,101,1,0,0,0,106,102,1,0,0,0,107,17,1,0,0,0,108,109,7,
1,0,0,109,19,1,0,0,0,110,111,7,2,0,0,111,21,1,0,0,0,112,126,5,22,0,0,113,
114,7,3,0,0,114,123,5,5,0,0,115,120,5,30,0,0,116,117,5,1,0,0,117,119,5,30,
0,0,118,116,1,0,0,0,119,122,1,0,0,0,120,118,1,0,0,0,120,121,1,0,0,0,121,
124,1,0,0,0,122,120,1,0,0,0,123,115,1,0,0,0,123,124,1,0,0,0,124,125,1,0,
0,0,125,127,5,6,0,0,126,113,1,0,0,0,126,127,1,0,0,0,127,23,1,0,0,0,128,129,
5,30,0,0,129,130,5,23,0,0,130,25,1,0,0,0,131,132,5,30,0,0,132,27,1,0,0,0,
133,134,3,20,10,0,134,135,3,18,9,0,135,29,1,0,0,0,14,36,44,56,62,69,72,80,
83,86,98,106,120,123,126];


const atn = new antlr4.atn.ATNDeserializer().deserialize(serializedATN);

const decisionsToDFA = atn.decisionToState.map( (ds, index) => new antlr4.dfa.DFA(ds, index) );

const sharedContextCache = new antlr4.atn.PredictionContextCache();

export default class FlowParser extends antlr4.Parser {

    static grammarFileName = "Flow.g4";
    static literalNames = [ null, "','", "'|'", "'{'", "'}'", "'('", "')'", 
                            "'='", "'!='", "'=~'", "'!~'", "'+'", "'-'", 
                            "'*'", "'/'", "'>'", "'<'", "'>='", "'<='", 
                            "'=='", "'by'", "'without'", null, null, "'let'", 
                            "'in'", "'and'", "'or'", "'unless'" ];
    static symbolicNames = [ null, null, null, null, null, null, null, "MATCH_EQ", 
                             "MATCH_NEQ", "MATCH_RE", "MATCH_NRE", "OP_ADD", 
                             "OP_SUB", "OP_MUL", "OP_DIV", "OP_GT", "OP_LT", 
                             "OP_GTE", "OP_LTE", "OP_EQ", "BY", "WITHOUT", 
                             "AGGREGATION_OP", "DURATION", "LET", "IN", 
                             "AND", "OR", "UNLESS", "NUMBER", "IDENTIFIER", 
                             "STRING", "COMMENT", "MULTILINE_COMMENT", "WS" ];
    static ruleNames = [ "query", "letExpression", "letBinding", "pipeline", 
                         "pipelineStep", "selector", "labelMatcher", "binarySelector", 
                         "binarySelectorLeg", "binaryLegLeaf", "binaryOperator", 
                         "aggregation", "aligner", "function", "binary" ];

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
	        this.state = 36;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 24:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 30;
	            this.letExpression();
	            this.state = 31;
	            this.match(FlowParser.EOF);
	            break;
	        case -1:
	        case 2:
	        case 3:
	        case 5:
	        case 29:
	        case 30:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 33;
	            this.pipeline();
	            this.state = 34;
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
	        this.state = 38;
	        this.match(FlowParser.LET);
	        this.state = 39;
	        this.letBinding();
	        this.state = 44;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===1) {
	            this.state = 40;
	            this.match(FlowParser.T__0);
	            this.state = 41;
	            this.letBinding();
	            this.state = 46;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 47;
	        this.match(FlowParser.IN);
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



	letBinding() {
	    let localctx = new LetBindingContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 4, FlowParser.RULE_letBinding);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 50;
	        this.match(FlowParser.IDENTIFIER);
	        this.state = 51;
	        this.match(FlowParser.MATCH_EQ);
	        this.state = 52;
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



	pipeline() {
	    let localctx = new PipelineContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 6, FlowParser.RULE_pipeline);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 56;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,2,this._ctx);
	        switch(la_) {
	        case 1:
	            this.state = 54;
	            this.selector();
	            break;

	        case 2:
	            this.state = 55;
	            this.binarySelector();
	            break;

	        }
	        this.state = 62;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===2) {
	            this.state = 58;
	            this.match(FlowParser.T__1);
	            this.state = 59;
	            this.pipelineStep();
	            this.state = 64;
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
	    this.enterRule(localctx, 8, FlowParser.RULE_pipelineStep);
	    try {
	        this.state = 69;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,4,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 65;
	            this.aggregation();
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 66;
	            this.aligner();
	            break;

	        case 3:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 67;
	            this.function_();
	            break;

	        case 4:
	            this.enterOuterAlt(localctx, 4);
	            this.state = 68;
	            this.binary();
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
	    this.enterRule(localctx, 10, FlowParser.RULE_selector);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 72;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===30) {
	            this.state = 71;
	            this.match(FlowParser.IDENTIFIER);
	        }

	        this.state = 86;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===3) {
	            this.state = 74;
	            this.match(FlowParser.T__2);
	            this.state = 83;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===30) {
	                this.state = 75;
	                this.labelMatcher();
	                this.state = 80;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	                while(_la===1) {
	                    this.state = 76;
	                    this.match(FlowParser.T__0);
	                    this.state = 77;
	                    this.labelMatcher();
	                    this.state = 82;
	                    this._errHandler.sync(this);
	                    _la = this._input.LA(1);
	                }
	            }

	            this.state = 85;
	            this.match(FlowParser.T__3);
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
	    this.enterRule(localctx, 12, FlowParser.RULE_labelMatcher);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 88;
	        this.match(FlowParser.IDENTIFIER);
	        this.state = 89;
	        _la = this._input.LA(1);
	        if(!((((_la) & ~0x1f) === 0 && ((1 << _la) & 1920) !== 0))) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	        this.state = 90;
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



	binarySelector() {
	    let localctx = new BinarySelectorContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 14, FlowParser.RULE_binarySelector);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 92;
	        this.binarySelectorLeg();
	        this.state = 98;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while((((_la) & ~0x1f) === 0 && ((1 << _la) & 1046784) !== 0)) {
	            this.state = 93;
	            this.binaryOperator();
	            this.state = 94;
	            this.binarySelectorLeg();
	            this.state = 100;
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



	binarySelectorLeg() {
	    let localctx = new BinarySelectorLegContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 16, FlowParser.RULE_binarySelectorLeg);
	    try {
	        this.state = 106;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 29:
	        case 30:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 101;
	            this.binaryLegLeaf();
	            break;
	        case 5:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 102;
	            this.match(FlowParser.T__4);
	            this.state = 103;
	            this.binarySelector();
	            this.state = 104;
	            this.match(FlowParser.T__5);
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



	binaryLegLeaf() {
	    let localctx = new BinaryLegLeafContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 18, FlowParser.RULE_binaryLegLeaf);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 108;
	        _la = this._input.LA(1);
	        if(!(_la===29 || _la===30)) {
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



	binaryOperator() {
	    let localctx = new BinaryOperatorContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 20, FlowParser.RULE_binaryOperator);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 110;
	        _la = this._input.LA(1);
	        if(!((((_la) & ~0x1f) === 0 && ((1 << _la) & 1046784) !== 0))) {
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



	aggregation() {
	    let localctx = new AggregationContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 22, FlowParser.RULE_aggregation);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 112;
	        this.match(FlowParser.AGGREGATION_OP);
	        this.state = 126;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===20 || _la===21) {
	            this.state = 113;
	            _la = this._input.LA(1);
	            if(!(_la===20 || _la===21)) {
	            this._errHandler.recoverInline(this);
	            }
	            else {
	            	this._errHandler.reportMatch(this);
	                this.consume();
	            }

	            this.state = 114;
	            this.match(FlowParser.T__4);
	            this.state = 123;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===30) {
	                this.state = 115;
	                this.match(FlowParser.IDENTIFIER);
	                this.state = 120;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	                while(_la===1) {
	                    this.state = 116;
	                    this.match(FlowParser.T__0);
	                    this.state = 117;
	                    this.match(FlowParser.IDENTIFIER);
	                    this.state = 122;
	                    this._errHandler.sync(this);
	                    _la = this._input.LA(1);
	                }
	            }

	            this.state = 125;
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



	aligner() {
	    let localctx = new AlignerContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 24, FlowParser.RULE_aligner);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 128;
	        this.match(FlowParser.IDENTIFIER);
	        this.state = 129;
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



	function_() {
	    let localctx = new FunctionContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 26, FlowParser.RULE_function);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 131;
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



	binary() {
	    let localctx = new BinaryContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 28, FlowParser.RULE_binary);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 133;
	        this.binaryOperator();
	        this.state = 134;
	        this.binaryLegLeaf();
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
FlowParser.OP_GT = 15;
FlowParser.OP_LT = 16;
FlowParser.OP_GTE = 17;
FlowParser.OP_LTE = 18;
FlowParser.OP_EQ = 19;
FlowParser.BY = 20;
FlowParser.WITHOUT = 21;
FlowParser.AGGREGATION_OP = 22;
FlowParser.DURATION = 23;
FlowParser.LET = 24;
FlowParser.IN = 25;
FlowParser.AND = 26;
FlowParser.OR = 27;
FlowParser.UNLESS = 28;
FlowParser.NUMBER = 29;
FlowParser.IDENTIFIER = 30;
FlowParser.STRING = 31;
FlowParser.COMMENT = 32;
FlowParser.MULTILINE_COMMENT = 33;
FlowParser.WS = 34;

FlowParser.RULE_query = 0;
FlowParser.RULE_letExpression = 1;
FlowParser.RULE_letBinding = 2;
FlowParser.RULE_pipeline = 3;
FlowParser.RULE_pipelineStep = 4;
FlowParser.RULE_selector = 5;
FlowParser.RULE_labelMatcher = 6;
FlowParser.RULE_binarySelector = 7;
FlowParser.RULE_binarySelectorLeg = 8;
FlowParser.RULE_binaryLegLeaf = 9;
FlowParser.RULE_binaryOperator = 10;
FlowParser.RULE_aggregation = 11;
FlowParser.RULE_aligner = 12;
FlowParser.RULE_function = 13;
FlowParser.RULE_binary = 14;

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

	pipeline() {
	    return this.getTypedRuleContext(PipelineContext,0);
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

	binarySelector() {
	    return this.getTypedRuleContext(BinarySelectorContext,0);
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

	binary() {
	    return this.getTypedRuleContext(BinaryContext,0);
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



class BinarySelectorContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = FlowParser.RULE_binarySelector;
    }

	binarySelectorLeg = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(BinarySelectorLegContext);
	    } else {
	        return this.getTypedRuleContext(BinarySelectorLegContext,i);
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
	        listener.enterBinarySelector(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.exitBinarySelector(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof FlowVisitor ) {
	        return visitor.visitBinarySelector(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class BinarySelectorLegContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = FlowParser.RULE_binarySelectorLeg;
    }

	binaryLegLeaf() {
	    return this.getTypedRuleContext(BinaryLegLeafContext,0);
	};

	binarySelector() {
	    return this.getTypedRuleContext(BinarySelectorContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.enterBinarySelectorLeg(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.exitBinarySelectorLeg(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof FlowVisitor ) {
	        return visitor.visitBinarySelectorLeg(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class BinaryLegLeafContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = FlowParser.RULE_binaryLegLeaf;
    }

	IDENTIFIER() {
	    return this.getToken(FlowParser.IDENTIFIER, 0);
	};

	NUMBER() {
	    return this.getToken(FlowParser.NUMBER, 0);
	};

	enterRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.enterBinaryLegLeaf(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.exitBinaryLegLeaf(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof FlowVisitor ) {
	        return visitor.visitBinaryLegLeaf(this);
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

	OP_GT() {
	    return this.getToken(FlowParser.OP_GT, 0);
	};

	OP_LT() {
	    return this.getToken(FlowParser.OP_LT, 0);
	};

	OP_GTE() {
	    return this.getToken(FlowParser.OP_GTE, 0);
	};

	OP_LTE() {
	    return this.getToken(FlowParser.OP_LTE, 0);
	};

	OP_EQ() {
	    return this.getToken(FlowParser.OP_EQ, 0);
	};

	MATCH_NEQ() {
	    return this.getToken(FlowParser.MATCH_NEQ, 0);
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



class BinaryContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = FlowParser.RULE_binary;
    }

	binaryOperator() {
	    return this.getTypedRuleContext(BinaryOperatorContext,0);
	};

	binaryLegLeaf() {
	    return this.getTypedRuleContext(BinaryLegLeafContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.enterBinary(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof FlowListener ) {
	        listener.exitBinary(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof FlowVisitor ) {
	        return visitor.visitBinary(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}




FlowParser.QueryContext = QueryContext; 
FlowParser.LetExpressionContext = LetExpressionContext; 
FlowParser.LetBindingContext = LetBindingContext; 
FlowParser.PipelineContext = PipelineContext; 
FlowParser.PipelineStepContext = PipelineStepContext; 
FlowParser.SelectorContext = SelectorContext; 
FlowParser.LabelMatcherContext = LabelMatcherContext; 
FlowParser.BinarySelectorContext = BinarySelectorContext; 
FlowParser.BinarySelectorLegContext = BinarySelectorLegContext; 
FlowParser.BinaryLegLeafContext = BinaryLegLeafContext; 
FlowParser.BinaryOperatorContext = BinaryOperatorContext; 
FlowParser.AggregationContext = AggregationContext; 
FlowParser.AlignerContext = AlignerContext; 
FlowParser.FunctionContext = FunctionContext; 
FlowParser.BinaryContext = BinaryContext; 
