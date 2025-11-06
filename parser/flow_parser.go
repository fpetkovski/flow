// Code generated from grammar/Flow.g4 by ANTLR 4.13.2. DO NOT EDIT.

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
		"", "','", "'('", "')'", "'|'", "'{'", "'}'", "'='", "'!='", "'=~'",
		"'!~'", "'+'", "'-'", "'*'", "'/'", "'by'", "'without'", "", "", "'let'",
		"'in'", "'and'", "'or'", "'unless'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "MATCH_EQ", "MATCH_NEQ", "MATCH_RE", "MATCH_NRE",
		"OP_ADD", "OP_SUB", "OP_MUL", "OP_DIV", "BY", "WITHOUT", "AGGREGATION_OP",
		"DURATION", "LET", "IN", "AND", "OR", "UNLESS", "NUMBER", "DURATION_UNIT",
		"IDENTIFIER", "STRING", "COMMENT", "MULTILINE_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"query", "letExpression", "letBinding", "binaryExpression", "binaryOperator",
		"primaryExpression", "pipeline", "pipelineStep", "selector", "labelMatcher",
		"aggregation", "aligner",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 119, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 31, 8, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 5, 1, 37, 8, 1, 10, 1, 12, 1, 40, 9, 1, 1, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 53, 8, 3, 10,
		3, 12, 3, 56, 9, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 65,
		8, 5, 1, 6, 1, 6, 1, 6, 5, 6, 70, 8, 6, 10, 6, 12, 6, 73, 9, 6, 1, 7, 1,
		7, 3, 7, 77, 8, 7, 1, 8, 3, 8, 80, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8,
		86, 8, 8, 10, 8, 12, 8, 89, 9, 8, 3, 8, 91, 8, 8, 1, 8, 3, 8, 94, 8, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10,
		106, 8, 10, 10, 10, 12, 10, 109, 9, 10, 3, 10, 111, 8, 10, 1, 10, 3, 10,
		114, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 0, 0, 12, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 0, 3, 1, 0, 11, 14, 1, 0, 7, 10, 1, 0, 15, 16, 119,
		0, 30, 1, 0, 0, 0, 2, 32, 1, 0, 0, 0, 4, 44, 1, 0, 0, 0, 6, 48, 1, 0, 0,
		0, 8, 57, 1, 0, 0, 0, 10, 64, 1, 0, 0, 0, 12, 66, 1, 0, 0, 0, 14, 76, 1,
		0, 0, 0, 16, 79, 1, 0, 0, 0, 18, 95, 1, 0, 0, 0, 20, 99, 1, 0, 0, 0, 22,
		115, 1, 0, 0, 0, 24, 25, 3, 2, 1, 0, 25, 26, 5, 0, 0, 1, 26, 31, 1, 0,
		0, 0, 27, 28, 3, 12, 6, 0, 28, 29, 5, 0, 0, 1, 29, 31, 1, 0, 0, 0, 30,
		24, 1, 0, 0, 0, 30, 27, 1, 0, 0, 0, 31, 1, 1, 0, 0, 0, 32, 33, 5, 19, 0,
		0, 33, 38, 3, 4, 2, 0, 34, 35, 5, 1, 0, 0, 35, 37, 3, 4, 2, 0, 36, 34,
		1, 0, 0, 0, 37, 40, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0,
		39, 41, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 41, 42, 5, 20, 0, 0, 42, 43, 3,
		6, 3, 0, 43, 3, 1, 0, 0, 0, 44, 45, 5, 26, 0, 0, 45, 46, 5, 7, 0, 0, 46,
		47, 3, 12, 6, 0, 47, 5, 1, 0, 0, 0, 48, 54, 3, 10, 5, 0, 49, 50, 3, 8,
		4, 0, 50, 51, 3, 10, 5, 0, 51, 53, 1, 0, 0, 0, 52, 49, 1, 0, 0, 0, 53,
		56, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 7, 1, 0, 0,
		0, 56, 54, 1, 0, 0, 0, 57, 58, 7, 0, 0, 0, 58, 9, 1, 0, 0, 0, 59, 65, 5,
		26, 0, 0, 60, 61, 5, 2, 0, 0, 61, 62, 3, 6, 3, 0, 62, 63, 5, 3, 0, 0, 63,
		65, 1, 0, 0, 0, 64, 59, 1, 0, 0, 0, 64, 60, 1, 0, 0, 0, 65, 11, 1, 0, 0,
		0, 66, 71, 3, 16, 8, 0, 67, 68, 5, 4, 0, 0, 68, 70, 3, 14, 7, 0, 69, 67,
		1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0,
		72, 13, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 77, 3, 20, 10, 0, 75, 77, 3,
		22, 11, 0, 76, 74, 1, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77, 15, 1, 0, 0, 0,
		78, 80, 5, 26, 0, 0, 79, 78, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 93, 1,
		0, 0, 0, 81, 90, 5, 5, 0, 0, 82, 87, 3, 18, 9, 0, 83, 84, 5, 1, 0, 0, 84,
		86, 3, 18, 9, 0, 85, 83, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0,
		0, 0, 87, 88, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 82,
		1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 94, 5, 6, 0, 0,
		93, 81, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 17, 1, 0, 0, 0, 95, 96, 5,
		26, 0, 0, 96, 97, 7, 1, 0, 0, 97, 98, 5, 27, 0, 0, 98, 19, 1, 0, 0, 0,
		99, 100, 5, 17, 0, 0, 100, 113, 7, 2, 0, 0, 101, 110, 5, 2, 0, 0, 102,
		107, 5, 26, 0, 0, 103, 104, 5, 1, 0, 0, 104, 106, 5, 26, 0, 0, 105, 103,
		1, 0, 0, 0, 106, 109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0,
		0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 110, 102, 1, 0, 0, 0,
		110, 111, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114, 5, 3, 0, 0, 113,
		101, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 21, 1, 0, 0, 0, 115, 116, 5,
		26, 0, 0, 116, 117, 5, 18, 0, 0, 117, 23, 1, 0, 0, 0, 13, 30, 38, 54, 64,
		71, 76, 79, 87, 90, 93, 107, 110, 113,
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
	FlowParserBY                = 15
	FlowParserWITHOUT           = 16
	FlowParserAGGREGATION_OP    = 17
	FlowParserDURATION          = 18
	FlowParserLET               = 19
	FlowParserIN                = 20
	FlowParserAND               = 21
	FlowParserOR                = 22
	FlowParserUNLESS            = 23
	FlowParserNUMBER            = 24
	FlowParserDURATION_UNIT     = 25
	FlowParserIDENTIFIER        = 26
	FlowParserSTRING            = 27
	FlowParserCOMMENT           = 28
	FlowParserMULTILINE_COMMENT = 29
	FlowParserWS                = 30
)

// FlowParser rules.
const (
	FlowParserRULE_query             = 0
	FlowParserRULE_letExpression     = 1
	FlowParserRULE_letBinding        = 2
	FlowParserRULE_binaryExpression  = 3
	FlowParserRULE_binaryOperator    = 4
	FlowParserRULE_primaryExpression = 5
	FlowParserRULE_pipeline          = 6
	FlowParserRULE_pipelineStep      = 7
	FlowParserRULE_selector          = 8
	FlowParserRULE_labelMatcher      = 9
	FlowParserRULE_aggregation       = 10
	FlowParserRULE_aligner           = 11
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
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FlowParserLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.LetExpression()
		}
		{
			p.SetState(25)
			p.Match(FlowParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FlowParserEOF, FlowParserT__3, FlowParserT__4, FlowParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(27)
			p.Pipeline()
		}
		{
			p.SetState(28)
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
	BinaryExpression() IBinaryExpressionContext

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

func (s *LetExpressionContext) BinaryExpression() IBinaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryExpressionContext)
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
		p.SetState(32)
		p.Match(FlowParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.LetBinding()
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FlowParserT__0 {
		{
			p.SetState(34)
			p.Match(FlowParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(35)
			p.LetBinding()
		}

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(41)
		p.Match(FlowParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.BinaryExpression()
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
		p.SetState(44)
		p.Match(FlowParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(FlowParserMATCH_EQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
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

// IBinaryExpressionContext is an interface to support dynamic dispatch.
type IBinaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPrimaryExpression() []IPrimaryExpressionContext
	PrimaryExpression(i int) IPrimaryExpressionContext
	AllBinaryOperator() []IBinaryOperatorContext
	BinaryOperator(i int) IBinaryOperatorContext

	// IsBinaryExpressionContext differentiates from other interfaces.
	IsBinaryExpressionContext()
}

type BinaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryExpressionContext() *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_binaryExpression
	return p
}

func InitEmptyBinaryExpressionContext(p *BinaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_binaryExpression
}

func (*BinaryExpressionContext) IsBinaryExpressionContext() {}

func NewBinaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_binaryExpression

	return p
}

func (s *BinaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryExpressionContext) AllPrimaryExpression() []IPrimaryExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			len++
		}
	}

	tst := make([]IPrimaryExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrimaryExpressionContext); ok {
			tst[i] = t.(IPrimaryExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryExpressionContext) PrimaryExpression(i int) IPrimaryExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
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

	return t.(IPrimaryExpressionContext)
}

func (s *BinaryExpressionContext) AllBinaryOperator() []IBinaryOperatorContext {
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

func (s *BinaryExpressionContext) BinaryOperator(i int) IBinaryOperatorContext {
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

func (s *BinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterBinaryExpression(s)
	}
}

func (s *BinaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitBinaryExpression(s)
	}
}

func (s *BinaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowVisitor:
		return t.VisitBinaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowParser) BinaryExpression() (localctx IBinaryExpressionContext) {
	localctx = NewBinaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FlowParserRULE_binaryExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.PrimaryExpression()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30720) != 0 {
		{
			p.SetState(49)
			p.BinaryOperator()
		}
		{
			p.SetState(50)
			p.PrimaryExpression()
		}

		p.SetState(56)
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
	p.EnterRule(localctx, 8, FlowParserRULE_binaryOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30720) != 0) {
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

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	BinaryExpression() IBinaryExpressionContext

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FlowParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FlowParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FlowParserIDENTIFIER, 0)
}

func (s *PrimaryExpressionContext) BinaryExpression() IBinaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinaryExpressionContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FlowListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FlowVisitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FlowParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FlowParserRULE_primaryExpression)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FlowParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Match(FlowParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case FlowParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.Match(FlowParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.BinaryExpression()
		}
		{
			p.SetState(62)
			p.Match(FlowParserT__2)
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

// IPipelineContext is an interface to support dynamic dispatch.
type IPipelineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Selector() ISelectorContext
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
	p.EnterRule(localctx, 12, FlowParserRULE_pipeline)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Selector()
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FlowParserT__3 {
		{
			p.SetState(67)
			p.Match(FlowParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.PipelineStep()
		}

		p.SetState(73)
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
	p.EnterRule(localctx, 14, FlowParserRULE_pipelineStep)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FlowParserAGGREGATION_OP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.Aggregation()
		}

	case FlowParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.Aligner()
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
	p.EnterRule(localctx, 16, FlowParserRULE_selector)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FlowParserIDENTIFIER {
		{
			p.SetState(78)
			p.Match(FlowParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FlowParserT__4 {
		{
			p.SetState(81)
			p.Match(FlowParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FlowParserIDENTIFIER {
			{
				p.SetState(82)
				p.LabelMatcher()
			}
			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == FlowParserT__0 {
				{
					p.SetState(83)
					p.Match(FlowParserT__0)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(84)
					p.LabelMatcher()
				}

				p.SetState(89)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(92)
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
	p.EnterRule(localctx, 18, FlowParserRULE_labelMatcher)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(FlowParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1920) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(97)
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
	p.EnterRule(localctx, 20, FlowParserRULE_aggregation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(FlowParserAGGREGATION_OP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FlowParserBY || _la == FlowParserWITHOUT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FlowParserT__1 {
		{
			p.SetState(101)
			p.Match(FlowParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == FlowParserIDENTIFIER {
			{
				p.SetState(102)
				p.Match(FlowParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(107)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == FlowParserT__0 {
				{
					p.SetState(103)
					p.Match(FlowParserT__0)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(104)
					p.Match(FlowParserIDENTIFIER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(109)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(112)
			p.Match(FlowParserT__2)
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
	p.EnterRule(localctx, 22, FlowParserRULE_aligner)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(FlowParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
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
