// Code generated from Flow.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Flow
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type FlowParser struct {
	*antlr.BaseParser
}

var FlowParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func flowParserInit() {
	staticData := &FlowParserStaticData
	staticData.LiteralNames = []string{
		"", "','", "'|'", "'{'", "'}'", "'('", "')'", "'='", "'!='", "'=~'",
		"'!~'", "'+'", "'-'", "'*'", "'/'", "'>'", "'<'", "'>='", "'<='", "'=='",
		"'by'", "'without'", "", "", "'let'", "'in'", "'and'", "'or'", "'unless'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "MATCH_EQ", "MATCH_NEQ", "MATCH_RE", "MATCH_NRE",
		"OP_ADD", "OP_SUB", "OP_MUL", "OP_DIV", "OP_GT", "OP_LT", "OP_GTE",
		"OP_LTE", "OP_EQ", "BY", "WITHOUT", "AGGREGATION_OP", "DURATION", "LET",
		"IN", "AND", "OR", "UNLESS", "NUMBER", "IDENTIFIER", "STRING", "COMMENT",
		"MULTILINE_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"query", "letExpression", "letBinding", "pipeline", "pipelineStep",
		"selector", "labelMatcher", "binarySelector", "binarySelectorLeg", "binaryLegLeaf",
		"binaryOperator", "aggregation", "aligner", "function", "binary",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 34, 137, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 37, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 43,
		8, 1, 10, 1, 12, 1, 46, 9, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 3, 3, 57, 8, 3, 1, 3, 1, 3, 5, 3, 61, 8, 3, 10, 3, 12, 3, 64,
		9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 70, 8, 4, 1, 5, 3, 5, 73, 8, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 5, 5, 79, 8, 5, 10, 5, 12, 5, 82, 9, 5, 3, 5, 84,
		8, 5, 1, 5, 3, 5, 87, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 5, 7, 97, 8, 7, 10, 7, 12, 7, 100, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 3, 8, 107, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 5, 11, 119, 8, 11, 10, 11, 12, 11, 122, 9, 11, 3, 11, 124,
		8, 11, 1, 11, 3, 11, 127, 8, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 14, 1, 14, 0, 0, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 0, 4, 1, 0, 7, 10, 1, 0, 29, 30, 2, 0, 8, 8, 11, 19, 1,
		0, 20, 21, 137, 0, 36, 1, 0, 0, 0, 2, 38, 1, 0, 0, 0, 4, 50, 1, 0, 0, 0,
		6, 56, 1, 0, 0, 0, 8, 69, 1, 0, 0, 0, 10, 72, 1, 0, 0, 0, 12, 88, 1, 0,
		0, 0, 14, 92, 1, 0, 0, 0, 16, 106, 1, 0, 0, 0, 18, 108, 1, 0, 0, 0, 20,
		110, 1, 0, 0, 0, 22, 112, 1, 0, 0, 0, 24, 128, 1, 0, 0, 0, 26, 131, 1,
		0, 0, 0, 28, 133, 1, 0, 0, 0, 30, 31, 3, 2, 1, 0, 31, 32, 5, 0, 0, 1, 32,
		37, 1, 0, 0, 0, 33, 34, 3, 6, 3, 0, 34, 35, 5, 0, 0, 1, 35, 37, 1, 0, 0,
		0, 36, 30, 1, 0, 0, 0, 36, 33, 1, 0, 0, 0, 37, 1, 1, 0, 0, 0, 38, 39, 5,
		24, 0, 0, 39, 44, 3, 4, 2, 0, 40, 41, 5, 1, 0, 0, 41, 43, 3, 4, 2, 0, 42,
		40, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45, 1, 0, 0,
		0, 45, 47, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 48, 5, 25, 0, 0, 48, 49,
		3, 6, 3, 0, 49, 3, 1, 0, 0, 0, 50, 51, 5, 30, 0, 0, 51, 52, 5, 7, 0, 0,
		52, 53, 3, 6, 3, 0, 53, 5, 1, 0, 0, 0, 54, 57, 3, 10, 5, 0, 55, 57, 3,
		14, 7, 0, 56, 54, 1, 0, 0, 0, 56, 55, 1, 0, 0, 0, 57, 62, 1, 0, 0, 0, 58,
		59, 5, 2, 0, 0, 59, 61, 3, 8, 4, 0, 60, 58, 1, 0, 0, 0, 61, 64, 1, 0, 0,
		0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 7, 1, 0, 0, 0, 64, 62, 1,
		0, 0, 0, 65, 70, 3, 22, 11, 0, 66, 70, 3, 24, 12, 0, 67, 70, 3, 26, 13,
		0, 68, 70, 3, 28, 14, 0, 69, 65, 1, 0, 0, 0, 69, 66, 1, 0, 0, 0, 69, 67,
		1, 0, 0, 0, 69, 68, 1, 0, 0, 0, 70, 9, 1, 0, 0, 0, 71, 73, 5, 30, 0, 0,
		72, 71, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 86, 1, 0, 0, 0, 74, 83, 5,
		3, 0, 0, 75, 80, 3, 12, 6, 0, 76, 77, 5, 1, 0, 0, 77, 79, 3, 12, 6, 0,
		78, 76, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1,
		0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 83, 75, 1, 0, 0, 0, 83,
		84, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 87, 5, 4, 0, 0, 86, 74, 1, 0, 0,
		0, 86, 87, 1, 0, 0, 0, 87, 11, 1, 0, 0, 0, 88, 89, 5, 30, 0, 0, 89, 90,
		7, 0, 0, 0, 90, 91, 5, 31, 0, 0, 91, 13, 1, 0, 0, 0, 92, 98, 3, 16, 8,
		0, 93, 94, 3, 20, 10, 0, 94, 95, 3, 16, 8, 0, 95, 97, 1, 0, 0, 0, 96, 93,
		1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0,
		99, 15, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 107, 3, 18, 9, 0, 102, 103,
		5, 5, 0, 0, 103, 104, 3, 14, 7, 0, 104, 105, 5, 6, 0, 0, 105, 107, 1, 0,
		0, 0, 106, 101, 1, 0, 0, 0, 106, 102, 1, 0, 0, 0, 107, 17, 1, 0, 0, 0,
		108, 109, 7, 1, 0, 0, 109, 19, 1, 0, 0, 0, 110, 111, 7, 2, 0, 0, 111, 21,
		1, 0, 0, 0, 112, 126, 5, 22, 0, 0, 113, 114, 7, 3, 0, 0, 114, 123, 5, 5,
		0, 0, 115, 120, 5, 30, 0, 0, 116, 117, 5, 1, 0, 0, 117, 119, 5, 30, 0,
		0, 118, 116, 1, 0, 0, 0, 119, 122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120,
		121, 1, 0, 0, 0, 121, 124, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 115,
		1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 127, 5, 6,
		0, 0, 126, 113, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 23, 1, 0, 0, 0,
		128, 129, 5, 30, 0, 0, 129, 130, 5, 23, 0, 0, 130, 25, 1, 0, 0, 0, 131,
		132, 5, 30, 0, 0, 132, 27, 1, 0, 0, 0, 133, 134, 3, 20, 10, 0, 134, 135,
		3, 18, 9, 0, 135, 29, 1, 0, 0, 0, 14, 36, 44, 56, 62, 69, 72, 80, 83, 86,
		98, 106, 120, 123, 126,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// FlowParserInit initializes any static state used to implement FlowParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFlowParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FlowParserInit() {
	staticData := &FlowParserStaticData
	staticData.once.Do(flowParserInit)
}

// NewFlowParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFlowParser(input antlr.TokenStream) *FlowParser {
	FlowParserInit()
	this := new(FlowParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &FlowParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Flow.g4"

	return this
}

// FlowParser tokens.
const (
	FlowParserEOF               = antlr.TokenEOF
	FlowParserT__0              = 1
	FlowParserT__1              = 2
	FlowParserT__2              = 3
	FlowParserT__3              = 4
	FlowParserT__4              = 5
	FlowParserT__5              = 6
	FlowParserMATCH_EQ          = 7
	FlowParserMATCH_NEQ         = 8
	FlowParserMATCH_RE          = 9
	FlowParserMATCH_NRE         = 10
	FlowParserOP_ADD            = 11
	FlowParserOP_SUB            = 12
	FlowParserOP_MUL            = 13
	FlowParserOP_DIV            = 14
	FlowParserOP_GT             = 15
	FlowParserOP_LT             = 16
	FlowParserOP_GTE            = 17
	FlowParserOP_LTE            = 18
	FlowParserOP_EQ             = 19
	FlowParserBY                = 20
	FlowParserWITHOUT           = 21
	FlowParserAGGREGATION_OP    = 22
	FlowParserDURATION          = 23
	FlowParserLET               = 24
	FlowParserIN                = 25
	FlowParserAND               = 26
	FlowParserOR                = 27
	FlowParserUNLESS            = 28
	FlowParserNUMBER            = 29
	FlowParserIDENTIFIER        = 30
	FlowParserSTRING            = 31
	FlowParserCOMMENT           = 32
	FlowParserMULTILINE_COMMENT = 33
	FlowParserWS                = 34
)

// FlowParser rules.
const (
	FlowParserRULE_query             = 0
	FlowParserRULE_letExpression     = 1
	FlowParserRULE_letBinding        = 2
	FlowParserRULE_pipeline          = 3
	FlowParserRULE_pipelineStep      = 4
	FlowParserRULE_selector          = 5
	FlowParserRULE_labelMatcher      = 6
	FlowParserRULE_binarySelector    = 7
	FlowParserRULE_binarySelectorLeg = 8
	FlowParserRULE_binaryLegLeaf     = 9
	FlowParserRULE_binaryOperator    = 10
	FlowParserRULE_aggregation       = 11
	FlowParserRULE_aligner           = 12
	FlowParserRULE_function          = 13
	FlowParserRULE_binary            = 14
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LetExpression() ILetExpressionContext
	EOF() antlr.TerminalNode
	Pipeline() IPipelineContext

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) LetExpression() ILetExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILetExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILetExpressionContext)
}

func (s *QueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(FlowParserEOF, 0)
}

func (s *QueryContext) Pipeline() IPipelineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipelineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipelineContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FlowParserRULE_query)
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FlowParserLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.LetExpression()
		}
		{
			p.SetState(31)
			p.Match(FlowParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FlowParserEOF, FlowParserT__1, FlowParserT__2, FlowParserT__4, FlowParserNUMBER, FlowParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.Pipeline()
		}
		{
			p.SetState(34)
			p.Match(FlowParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILetExpressionContext is an interface to support dynamic dispatch.
type ILetExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LET() antlr.TerminalNode
	AllLetBinding() []ILetBindingContext
	LetBinding(i int) ILetBindingContext
	IN() antlr.TerminalNode
	Pipeline() IPipelineContext

	// IsLetExpressionContext differentiates from other interfaces.
	IsLetExpressionContext()
}

type LetExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetExpressionContext() *LetExpressionContext {
	var p = new(LetExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_letExpression
	return p
}

func InitEmptyLetExpressionContext(p *LetExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_letExpression
}

func (*LetExpressionContext) IsLetExpressionContext() {}

func NewLetExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetExpressionContext {
	var p = new(LetExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_letExpression

	return p
}

func (s *LetExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LetExpressionContext) LET() antlr.TerminalNode {
	return s.GetToken(FlowParserLET, 0)
}

func (s *LetExpressionContext) AllLetBinding() []ILetBindingContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILetBindingContext); ok {
			len++
		}
	}

	tst := make([]ILetBindingContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILetBindingContext); ok {
			tst[i] = t.(ILetBindingContext)
			i++
		}
	}

	return tst
}

func (s *LetExpressionContext) LetBinding(i int) ILetBindingContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILetBindingContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILetBindingContext)
}

func (s *LetExpressionContext) IN() antlr.TerminalNode {
	return s.GetToken(FlowParserIN, 0)
}

func (s *LetExpressionContext) Pipeline() IPipelineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipelineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipelineContext)
}

func (s *LetExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LetExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterLetExpression(s)
	}
}

func (s *LetExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitLetExpression(s)
	}
}

func (s *LetExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowVisitor:
		return t.VisitLetExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowParser) LetExpression() (localctx ILetExpressionContext) {
	localctx = NewLetExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FlowParserRULE_letExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(FlowParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.LetBinding()
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FlowParserT__0 {
		{
			p.SetState(40)
			p.Match(FlowParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.LetBinding()
		}

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(47)
		p.Match(FlowParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Pipeline()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILetBindingContext is an interface to support dynamic dispatch.
type ILetBindingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	MATCH_EQ() antlr.TerminalNode
	Pipeline() IPipelineContext

	// IsLetBindingContext differentiates from other interfaces.
	IsLetBindingContext()
}

type LetBindingContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetBindingContext() *LetBindingContext {
	var p = new(LetBindingContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_letBinding
	return p
}

func InitEmptyLetBindingContext(p *LetBindingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_letBinding
}

func (*LetBindingContext) IsLetBindingContext() {}

func NewLetBindingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetBindingContext {
	var p = new(LetBindingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_letBinding

	return p
}

func (s *LetBindingContext) GetParser() antlr.Parser { return s.parser }

func (s *LetBindingContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FlowParserIDENTIFIER, 0)
}

func (s *LetBindingContext) MATCH_EQ() antlr.TerminalNode {
	return s.GetToken(FlowParserMATCH_EQ, 0)
}

func (s *LetBindingContext) Pipeline() IPipelineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipelineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipelineContext)
}

func (s *LetBindingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetBindingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LetBindingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterLetBinding(s)
	}
}

func (s *LetBindingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitLetBinding(s)
	}
}

func (s *LetBindingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowVisitor:
		return t.VisitLetBinding(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowParser) LetBinding() (localctx ILetBindingContext) {
	localctx = NewLetBindingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FlowParserRULE_letBinding)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(FlowParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Match(FlowParserMATCH_EQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Pipeline()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPipelineContext is an interface to support dynamic dispatch.
type IPipelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Selector() ISelectorContext
	BinarySelector() IBinarySelectorContext
	AllPipelineStep() []IPipelineStepContext
	PipelineStep(i int) IPipelineStepContext

	// IsPipelineContext differentiates from other interfaces.
	IsPipelineContext()
}

type PipelineContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipelineContext() *PipelineContext {
	var p = new(PipelineContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_pipeline
	return p
}

func InitEmptyPipelineContext(p *PipelineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_pipeline
}

func (*PipelineContext) IsPipelineContext() {}

func NewPipelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipelineContext {
	var p = new(PipelineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_pipeline

	return p
}

func (s *PipelineContext) GetParser() antlr.Parser { return s.parser }

func (s *PipelineContext) Selector() ISelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *PipelineContext) BinarySelector() IBinarySelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinarySelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinarySelectorContext)
}

func (s *PipelineContext) AllPipelineStep() []IPipelineStepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPipelineStepContext); ok {
			len++
		}
	}

	tst := make([]IPipelineStepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPipelineStepContext); ok {
			tst[i] = t.(IPipelineStepContext)
			i++
		}
	}

	return tst
}

func (s *PipelineContext) PipelineStep(i int) IPipelineStepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPipelineStepContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPipelineStepContext)
}

func (s *PipelineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipelineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipelineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterPipeline(s)
	}
}

func (s *PipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitPipeline(s)
	}
}

func (s *PipelineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowVisitor:
		return t.VisitPipeline(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowParser) Pipeline() (localctx IPipelineContext) {
	localctx = NewPipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FlowParserRULE_pipeline)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(54)
			p.Selector()
		}

	case 2:
		{
			p.SetState(55)
			p.BinarySelector()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FlowParserT__1 {
		{
			p.SetState(58)
			p.Match(FlowParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.PipelineStep()
		}

		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPipelineStepContext is an interface to support dynamic dispatch.
type IPipelineStepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Aggregation() IAggregationContext
	Aligner() IAlignerContext
	Function() IFunctionContext
	Binary() IBinaryContext

	// IsPipelineStepContext differentiates from other interfaces.
	IsPipelineStepContext()
}

type PipelineStepContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPipelineStepContext() *PipelineStepContext {
	var p = new(PipelineStepContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_pipelineStep
	return p
}

func InitEmptyPipelineStepContext(p *PipelineStepContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_pipelineStep
}

func (*PipelineStepContext) IsPipelineStepContext() {}

func NewPipelineStepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipelineStepContext {
	var p = new(PipelineStepContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_pipelineStep

	return p
}

func (s *PipelineStepContext) GetParser() antlr.Parser { return s.parser }

func (s *PipelineStepContext) Aggregation() IAggregationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregationContext)
}

func (s *PipelineStepContext) Aligner() IAlignerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlignerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAlignerContext)
}

func (s *PipelineStepContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *PipelineStepContext) Binary() IBinaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryContext)
}

func (s *PipelineStepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PipelineStepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PipelineStepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterPipelineStep(s)
	}
}

func (s *PipelineStepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitPipelineStep(s)
	}
}

func (s *PipelineStepContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowVisitor:
		return t.VisitPipelineStep(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowParser) PipelineStep() (localctx IPipelineStepContext) {
	localctx = NewPipelineStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FlowParserRULE_pipelineStep)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Aggregation()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Aligner()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Function()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
			p.Binary()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllLabelMatcher() []ILabelMatcherContext
	LabelMatcher(i int) ILabelMatcherContext

	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_selector
	return p
}

func InitEmptySelectorContext(p *SelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_selector
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FlowParserIDENTIFIER, 0)
}

func (s *SelectorContext) AllLabelMatcher() []ILabelMatcherContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILabelMatcherContext); ok {
			len++
		}
	}

	tst := make([]ILabelMatcherContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILabelMatcherContext); ok {
			tst[i] = t.(ILabelMatcherContext)
			i++
		}
	}

	return tst
}

func (s *SelectorContext) LabelMatcher(i int) ILabelMatcherContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelMatcherContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelMatcherContext)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterSelector(s)
	}
}

func (s *SelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitSelector(s)
	}
}

func (s *SelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowVisitor:
		return t.VisitSelector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FlowParserRULE_selector)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FlowParserIDENTIFIER {
		{
			p.SetState(71)
			p.Match(FlowParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FlowParserT__2 {
		{
			p.SetState(74)
			p.Match(FlowParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FlowParserIDENTIFIER {
			{
				p.SetState(75)
				p.LabelMatcher()
			}
			p.SetState(80)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == FlowParserT__0 {
				{
					p.SetState(76)
					p.Match(FlowParserT__0)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(77)
					p.LabelMatcher()
				}

				p.SetState(82)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(85)
			p.Match(FlowParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILabelMatcherContext is an interface to support dynamic dispatch.
type ILabelMatcherContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	MATCH_EQ() antlr.TerminalNode
	MATCH_NEQ() antlr.TerminalNode
	MATCH_RE() antlr.TerminalNode
	MATCH_NRE() antlr.TerminalNode

	// IsLabelMatcherContext differentiates from other interfaces.
	IsLabelMatcherContext()
}

type LabelMatcherContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelMatcherContext() *LabelMatcherContext {
	var p = new(LabelMatcherContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_labelMatcher
	return p
}

func InitEmptyLabelMatcherContext(p *LabelMatcherContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_labelMatcher
}

func (*LabelMatcherContext) IsLabelMatcherContext() {}

func NewLabelMatcherContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelMatcherContext {
	var p = new(LabelMatcherContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_labelMatcher

	return p
}

func (s *LabelMatcherContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelMatcherContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FlowParserIDENTIFIER, 0)
}

func (s *LabelMatcherContext) STRING() antlr.TerminalNode {
	return s.GetToken(FlowParserSTRING, 0)
}

func (s *LabelMatcherContext) MATCH_EQ() antlr.TerminalNode {
	return s.GetToken(FlowParserMATCH_EQ, 0)
}

func (s *LabelMatcherContext) MATCH_NEQ() antlr.TerminalNode {
	return s.GetToken(FlowParserMATCH_NEQ, 0)
}

func (s *LabelMatcherContext) MATCH_RE() antlr.TerminalNode {
	return s.GetToken(FlowParserMATCH_RE, 0)
}

func (s *LabelMatcherContext) MATCH_NRE() antlr.TerminalNode {
	return s.GetToken(FlowParserMATCH_NRE, 0)
}

func (s *LabelMatcherContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelMatcherContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelMatcherContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterLabelMatcher(s)
	}
}

func (s *LabelMatcherContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitLabelMatcher(s)
	}
}

func (s *LabelMatcherContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowVisitor:
		return t.VisitLabelMatcher(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowParser) LabelMatcher() (localctx ILabelMatcherContext) {
	localctx = NewLabelMatcherContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FlowParserRULE_labelMatcher)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(FlowParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1920) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(90)
		p.Match(FlowParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinarySelectorContext is an interface to support dynamic dispatch.
type IBinarySelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBinarySelectorLeg() []IBinarySelectorLegContext
	BinarySelectorLeg(i int) IBinarySelectorLegContext
	AllBinaryOperator() []IBinaryOperatorContext
	BinaryOperator(i int) IBinaryOperatorContext

	// IsBinarySelectorContext differentiates from other interfaces.
	IsBinarySelectorContext()
}

type BinarySelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinarySelectorContext() *BinarySelectorContext {
	var p = new(BinarySelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_binarySelector
	return p
}

func InitEmptyBinarySelectorContext(p *BinarySelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_binarySelector
}

func (*BinarySelectorContext) IsBinarySelectorContext() {}

func NewBinarySelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinarySelectorContext {
	var p = new(BinarySelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_binarySelector

	return p
}

func (s *BinarySelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *BinarySelectorContext) AllBinarySelectorLeg() []IBinarySelectorLegContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBinarySelectorLegContext); ok {
			len++
		}
	}

	tst := make([]IBinarySelectorLegContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBinarySelectorLegContext); ok {
			tst[i] = t.(IBinarySelectorLegContext)
			i++
		}
	}

	return tst
}

func (s *BinarySelectorContext) BinarySelectorLeg(i int) IBinarySelectorLegContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinarySelectorLegContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinarySelectorLegContext)
}

func (s *BinarySelectorContext) AllBinaryOperator() []IBinaryOperatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBinaryOperatorContext); ok {
			len++
		}
	}

	tst := make([]IBinaryOperatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBinaryOperatorContext); ok {
			tst[i] = t.(IBinaryOperatorContext)
			i++
		}
	}

	return tst
}

func (s *BinarySelectorContext) BinaryOperator(i int) IBinaryOperatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryOperatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryOperatorContext)
}

func (s *BinarySelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinarySelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinarySelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterBinarySelector(s)
	}
}

func (s *BinarySelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitBinarySelector(s)
	}
}

func (s *BinarySelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowVisitor:
		return t.VisitBinarySelector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowParser) BinarySelector() (localctx IBinarySelectorContext) {
	localctx = NewBinarySelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FlowParserRULE_binarySelector)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.BinarySelectorLeg()
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1046784) != 0 {
		{
			p.SetState(93)
			p.BinaryOperator()
		}
		{
			p.SetState(94)
			p.BinarySelectorLeg()
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinarySelectorLegContext is an interface to support dynamic dispatch.
type IBinarySelectorLegContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BinaryLegLeaf() IBinaryLegLeafContext
	BinarySelector() IBinarySelectorContext

	// IsBinarySelectorLegContext differentiates from other interfaces.
	IsBinarySelectorLegContext()
}

type BinarySelectorLegContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinarySelectorLegContext() *BinarySelectorLegContext {
	var p = new(BinarySelectorLegContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_binarySelectorLeg
	return p
}

func InitEmptyBinarySelectorLegContext(p *BinarySelectorLegContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_binarySelectorLeg
}

func (*BinarySelectorLegContext) IsBinarySelectorLegContext() {}

func NewBinarySelectorLegContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinarySelectorLegContext {
	var p = new(BinarySelectorLegContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_binarySelectorLeg

	return p
}

func (s *BinarySelectorLegContext) GetParser() antlr.Parser { return s.parser }

func (s *BinarySelectorLegContext) BinaryLegLeaf() IBinaryLegLeafContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryLegLeafContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryLegLeafContext)
}

func (s *BinarySelectorLegContext) BinarySelector() IBinarySelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinarySelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinarySelectorContext)
}

func (s *BinarySelectorLegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinarySelectorLegContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinarySelectorLegContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterBinarySelectorLeg(s)
	}
}

func (s *BinarySelectorLegContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitBinarySelectorLeg(s)
	}
}

func (s *BinarySelectorLegContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowVisitor:
		return t.VisitBinarySelectorLeg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowParser) BinarySelectorLeg() (localctx IBinarySelectorLegContext) {
	localctx = NewBinarySelectorLegContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FlowParserRULE_binarySelectorLeg)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FlowParserNUMBER, FlowParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.BinaryLegLeaf()
		}

	case FlowParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.Match(FlowParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.BinarySelector()
		}
		{
			p.SetState(104)
			p.Match(FlowParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinaryLegLeafContext is an interface to support dynamic dispatch.
type IBinaryLegLeafContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsBinaryLegLeafContext differentiates from other interfaces.
	IsBinaryLegLeafContext()
}

type BinaryLegLeafContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryLegLeafContext() *BinaryLegLeafContext {
	var p = new(BinaryLegLeafContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_binaryLegLeaf
	return p
}

func InitEmptyBinaryLegLeafContext(p *BinaryLegLeafContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_binaryLegLeaf
}

func (*BinaryLegLeafContext) IsBinaryLegLeafContext() {}

func NewBinaryLegLeafContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryLegLeafContext {
	var p = new(BinaryLegLeafContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_binaryLegLeaf

	return p
}

func (s *BinaryLegLeafContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryLegLeafContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FlowParserIDENTIFIER, 0)
}

func (s *BinaryLegLeafContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FlowParserNUMBER, 0)
}

func (s *BinaryLegLeafContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryLegLeafContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryLegLeafContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterBinaryLegLeaf(s)
	}
}

func (s *BinaryLegLeafContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitBinaryLegLeaf(s)
	}
}

func (s *BinaryLegLeafContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowVisitor:
		return t.VisitBinaryLegLeaf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowParser) BinaryLegLeaf() (localctx IBinaryLegLeafContext) {
	localctx = NewBinaryLegLeafContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FlowParserRULE_binaryLegLeaf)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FlowParserNUMBER || _la == FlowParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinaryOperatorContext is an interface to support dynamic dispatch.
type IBinaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_ADD() antlr.TerminalNode
	OP_SUB() antlr.TerminalNode
	OP_MUL() antlr.TerminalNode
	OP_DIV() antlr.TerminalNode
	OP_GT() antlr.TerminalNode
	OP_LT() antlr.TerminalNode
	OP_GTE() antlr.TerminalNode
	OP_LTE() antlr.TerminalNode
	OP_EQ() antlr.TerminalNode
	MATCH_NEQ() antlr.TerminalNode

	// IsBinaryOperatorContext differentiates from other interfaces.
	IsBinaryOperatorContext()
}

type BinaryOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOperatorContext() *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_binaryOperator
	return p
}

func InitEmptyBinaryOperatorContext(p *BinaryOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_binaryOperator
}

func (*BinaryOperatorContext) IsBinaryOperatorContext() {}

func NewBinaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOperatorContext {
	var p = new(BinaryOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_binaryOperator

	return p
}

func (s *BinaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOperatorContext) OP_ADD() antlr.TerminalNode {
	return s.GetToken(FlowParserOP_ADD, 0)
}

func (s *BinaryOperatorContext) OP_SUB() antlr.TerminalNode {
	return s.GetToken(FlowParserOP_SUB, 0)
}

func (s *BinaryOperatorContext) OP_MUL() antlr.TerminalNode {
	return s.GetToken(FlowParserOP_MUL, 0)
}

func (s *BinaryOperatorContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(FlowParserOP_DIV, 0)
}

func (s *BinaryOperatorContext) OP_GT() antlr.TerminalNode {
	return s.GetToken(FlowParserOP_GT, 0)
}

func (s *BinaryOperatorContext) OP_LT() antlr.TerminalNode {
	return s.GetToken(FlowParserOP_LT, 0)
}

func (s *BinaryOperatorContext) OP_GTE() antlr.TerminalNode {
	return s.GetToken(FlowParserOP_GTE, 0)
}

func (s *BinaryOperatorContext) OP_LTE() antlr.TerminalNode {
	return s.GetToken(FlowParserOP_LTE, 0)
}

func (s *BinaryOperatorContext) OP_EQ() antlr.TerminalNode {
	return s.GetToken(FlowParserOP_EQ, 0)
}

func (s *BinaryOperatorContext) MATCH_NEQ() antlr.TerminalNode {
	return s.GetToken(FlowParserMATCH_NEQ, 0)
}

func (s *BinaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterBinaryOperator(s)
	}
}

func (s *BinaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitBinaryOperator(s)
	}
}

func (s *BinaryOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowVisitor:
		return t.VisitBinaryOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowParser) BinaryOperator() (localctx IBinaryOperatorContext) {
	localctx = NewBinaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FlowParserRULE_binaryOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1046784) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAggregationContext is an interface to support dynamic dispatch.
type IAggregationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AGGREGATION_OP() antlr.TerminalNode
	BY() antlr.TerminalNode
	WITHOUT() antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

	// IsAggregationContext differentiates from other interfaces.
	IsAggregationContext()
}

type AggregationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregationContext() *AggregationContext {
	var p = new(AggregationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_aggregation
	return p
}

func InitEmptyAggregationContext(p *AggregationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_aggregation
}

func (*AggregationContext) IsAggregationContext() {}

func NewAggregationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregationContext {
	var p = new(AggregationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_aggregation

	return p
}

func (s *AggregationContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregationContext) AGGREGATION_OP() antlr.TerminalNode {
	return s.GetToken(FlowParserAGGREGATION_OP, 0)
}

func (s *AggregationContext) BY() antlr.TerminalNode {
	return s.GetToken(FlowParserBY, 0)
}

func (s *AggregationContext) WITHOUT() antlr.TerminalNode {
	return s.GetToken(FlowParserWITHOUT, 0)
}

func (s *AggregationContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(FlowParserIDENTIFIER)
}

func (s *AggregationContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(FlowParserIDENTIFIER, i)
}

func (s *AggregationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterAggregation(s)
	}
}

func (s *AggregationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitAggregation(s)
	}
}

func (s *AggregationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowVisitor:
		return t.VisitAggregation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowParser) Aggregation() (localctx IAggregationContext) {
	localctx = NewAggregationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FlowParserRULE_aggregation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(FlowParserAGGREGATION_OP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FlowParserBY || _la == FlowParserWITHOUT {
		{
			p.SetState(113)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FlowParserBY || _la == FlowParserWITHOUT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		{
			p.SetState(114)
			p.Match(FlowParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FlowParserIDENTIFIER {
			{
				p.SetState(115)
				p.Match(FlowParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == FlowParserT__0 {
				{
					p.SetState(116)
					p.Match(FlowParserT__0)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(117)
					p.Match(FlowParserIDENTIFIER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(122)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(125)
			p.Match(FlowParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAlignerContext is an interface to support dynamic dispatch.
type IAlignerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	DURATION() antlr.TerminalNode

	// IsAlignerContext differentiates from other interfaces.
	IsAlignerContext()
}

type AlignerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlignerContext() *AlignerContext {
	var p = new(AlignerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_aligner
	return p
}

func InitEmptyAlignerContext(p *AlignerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_aligner
}

func (*AlignerContext) IsAlignerContext() {}

func NewAlignerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlignerContext {
	var p = new(AlignerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_aligner

	return p
}

func (s *AlignerContext) GetParser() antlr.Parser { return s.parser }

func (s *AlignerContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FlowParserIDENTIFIER, 0)
}

func (s *AlignerContext) DURATION() antlr.TerminalNode {
	return s.GetToken(FlowParserDURATION, 0)
}

func (s *AlignerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlignerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlignerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterAligner(s)
	}
}

func (s *AlignerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitAligner(s)
	}
}

func (s *AlignerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowVisitor:
		return t.VisitAligner(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowParser) Aligner() (localctx IAlignerContext) {
	localctx = NewAlignerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FlowParserRULE_aligner)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(FlowParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(FlowParserDURATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_function
	return p
}

func InitEmptyFunctionContext(p *FunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_function
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FlowParserIDENTIFIER, 0)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FlowParserRULE_function)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(FlowParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinaryContext is an interface to support dynamic dispatch.
type IBinaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BinaryOperator() IBinaryOperatorContext
	BinaryLegLeaf() IBinaryLegLeafContext

	// IsBinaryContext differentiates from other interfaces.
	IsBinaryContext()
}

type BinaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryContext() *BinaryContext {
	var p = new(BinaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_binary
	return p
}

func InitEmptyBinaryContext(p *BinaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_binary
}

func (*BinaryContext) IsBinaryContext() {}

func NewBinaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryContext {
	var p = new(BinaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_binary

	return p
}

func (s *BinaryContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryContext) BinaryOperator() IBinaryOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryOperatorContext)
}

func (s *BinaryContext) BinaryLegLeaf() IBinaryLegLeafContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryLegLeafContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryLegLeafContext)
}

func (s *BinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterBinary(s)
	}
}

func (s *BinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitBinary(s)
	}
}

func (s *BinaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowVisitor:
		return t.VisitBinary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowParser) Binary() (localctx IBinaryContext) {
	localctx = NewBinaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FlowParserRULE_binary)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.BinaryOperator()
	}
	{
		p.SetState(134)
		p.BinaryLegLeaf()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
