// Code generated from grammar/PromQLPlus.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PromQLPlus
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

type PromQLPlusParser struct {
	*antlr.BaseParser
}

var PromQLPlusParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func promqlplusParserInit() {
	staticData := &PromQLPlusParserStaticData
	staticData.LiteralNames = []string{
		"", "','", "'/'", "'+'", "'-'", "'*'", "'%'", "'|'", "'{'", "'}'", "'('",
		"')'", "'='", "'!='", "'=~'", "'!~'", "'by'", "'without'", "", "", "'let'",
		"'in'", "'and'", "'or'", "'unless'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "MATCH_EQ", "MATCH_NEQ",
		"MATCH_RE", "MATCH_NRE", "BY", "WITHOUT", "AGGREGATION_OP", "DURATION",
		"LET", "IN", "AND", "OR", "UNLESS", "NUMBER", "DURATION_UNIT", "IDENTIFIER",
		"STRING", "COMMENT", "MULTILINE_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"query", "letExpression", "letBindings", "letBinding", "expression",
		"pipeline", "pipelineStep", "selector", "labelMatcher", "aggregation",
		"aligner",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 31, 104, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 29, 8, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 39, 8, 2, 10, 2, 12, 2, 42, 9, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5,
		55, 8, 5, 10, 5, 12, 5, 58, 9, 5, 1, 6, 1, 6, 3, 6, 62, 8, 6, 1, 7, 3,
		7, 65, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 71, 8, 7, 10, 7, 12, 7, 74,
		9, 7, 3, 7, 76, 8, 7, 1, 7, 3, 7, 79, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 91, 8, 9, 10, 9, 12, 9, 94, 9, 9,
		3, 9, 96, 8, 9, 1, 9, 3, 9, 99, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 0, 0,
		11, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 0, 3, 1, 0, 2, 6, 1, 0, 12,
		15, 1, 0, 16, 17, 103, 0, 28, 1, 0, 0, 0, 2, 30, 1, 0, 0, 0, 4, 35, 1,
		0, 0, 0, 6, 43, 1, 0, 0, 0, 8, 47, 1, 0, 0, 0, 10, 51, 1, 0, 0, 0, 12,
		61, 1, 0, 0, 0, 14, 64, 1, 0, 0, 0, 16, 80, 1, 0, 0, 0, 18, 84, 1, 0, 0,
		0, 20, 100, 1, 0, 0, 0, 22, 23, 3, 2, 1, 0, 23, 24, 5, 0, 0, 1, 24, 29,
		1, 0, 0, 0, 25, 26, 3, 10, 5, 0, 26, 27, 5, 0, 0, 1, 27, 29, 1, 0, 0, 0,
		28, 22, 1, 0, 0, 0, 28, 25, 1, 0, 0, 0, 29, 1, 1, 0, 0, 0, 30, 31, 5, 20,
		0, 0, 31, 32, 3, 4, 2, 0, 32, 33, 5, 21, 0, 0, 33, 34, 3, 8, 4, 0, 34,
		3, 1, 0, 0, 0, 35, 40, 3, 6, 3, 0, 36, 37, 5, 1, 0, 0, 37, 39, 3, 6, 3,
		0, 38, 36, 1, 0, 0, 0, 39, 42, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 41,
		1, 0, 0, 0, 41, 5, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 43, 44, 5, 27, 0, 0,
		44, 45, 5, 12, 0, 0, 45, 46, 3, 10, 5, 0, 46, 7, 1, 0, 0, 0, 47, 48, 5,
		27, 0, 0, 48, 49, 7, 0, 0, 0, 49, 50, 5, 27, 0, 0, 50, 9, 1, 0, 0, 0, 51,
		56, 3, 14, 7, 0, 52, 53, 5, 7, 0, 0, 53, 55, 3, 12, 6, 0, 54, 52, 1, 0,
		0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 11,
		1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 59, 62, 3, 18, 9, 0, 60, 62, 3, 20, 10,
		0, 61, 59, 1, 0, 0, 0, 61, 60, 1, 0, 0, 0, 62, 13, 1, 0, 0, 0, 63, 65,
		5, 27, 0, 0, 64, 63, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 78, 1, 0, 0, 0,
		66, 75, 5, 8, 0, 0, 67, 72, 3, 16, 8, 0, 68, 69, 5, 1, 0, 0, 69, 71, 3,
		16, 8, 0, 70, 68, 1, 0, 0, 0, 71, 74, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 72,
		73, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 75, 67, 1, 0, 0,
		0, 75, 76, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 79, 5, 9, 0, 0, 78, 66,
		1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 15, 1, 0, 0, 0, 80, 81, 5, 27, 0, 0,
		81, 82, 7, 1, 0, 0, 82, 83, 5, 28, 0, 0, 83, 17, 1, 0, 0, 0, 84, 85, 5,
		18, 0, 0, 85, 98, 7, 2, 0, 0, 86, 95, 5, 10, 0, 0, 87, 92, 5, 27, 0, 0,
		88, 89, 5, 1, 0, 0, 89, 91, 5, 27, 0, 0, 90, 88, 1, 0, 0, 0, 91, 94, 1,
		0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94,
		92, 1, 0, 0, 0, 95, 87, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 1, 0, 0,
		0, 97, 99, 5, 11, 0, 0, 98, 86, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 19,
		1, 0, 0, 0, 100, 101, 5, 27, 0, 0, 101, 102, 5, 19, 0, 0, 102, 21, 1, 0,
		0, 0, 11, 28, 40, 56, 61, 64, 72, 75, 78, 92, 95, 98,
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

// PromQLPlusParserInit initializes any static state used to implement PromQLPlusParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPromQLPlusParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PromQLPlusParserInit() {
	staticData := &PromQLPlusParserStaticData
	staticData.once.Do(promqlplusParserInit)
}

// NewPromQLPlusParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPromQLPlusParser(input antlr.TokenStream) *PromQLPlusParser {
	PromQLPlusParserInit()
	this := new(PromQLPlusParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PromQLPlusParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PromQLPlus.g4"

	return this
}

// PromQLPlusParser tokens.
const (
	PromQLPlusParserEOF               = antlr.TokenEOF
	PromQLPlusParserT__0              = 1
	PromQLPlusParserT__1              = 2
	PromQLPlusParserT__2              = 3
	PromQLPlusParserT__3              = 4
	PromQLPlusParserT__4              = 5
	PromQLPlusParserT__5              = 6
	PromQLPlusParserT__6              = 7
	PromQLPlusParserT__7              = 8
	PromQLPlusParserT__8              = 9
	PromQLPlusParserT__9              = 10
	PromQLPlusParserT__10             = 11
	PromQLPlusParserMATCH_EQ          = 12
	PromQLPlusParserMATCH_NEQ         = 13
	PromQLPlusParserMATCH_RE          = 14
	PromQLPlusParserMATCH_NRE         = 15
	PromQLPlusParserBY                = 16
	PromQLPlusParserWITHOUT           = 17
	PromQLPlusParserAGGREGATION_OP    = 18
	PromQLPlusParserDURATION          = 19
	PromQLPlusParserLET               = 20
	PromQLPlusParserIN                = 21
	PromQLPlusParserAND               = 22
	PromQLPlusParserOR                = 23
	PromQLPlusParserUNLESS            = 24
	PromQLPlusParserNUMBER            = 25
	PromQLPlusParserDURATION_UNIT     = 26
	PromQLPlusParserIDENTIFIER        = 27
	PromQLPlusParserSTRING            = 28
	PromQLPlusParserCOMMENT           = 29
	PromQLPlusParserMULTILINE_COMMENT = 30
	PromQLPlusParserWS                = 31
)

// PromQLPlusParser rules.
const (
	PromQLPlusParserRULE_query         = 0
	PromQLPlusParserRULE_letExpression = 1
	PromQLPlusParserRULE_letBindings   = 2
	PromQLPlusParserRULE_letBinding    = 3
	PromQLPlusParserRULE_expression    = 4
	PromQLPlusParserRULE_pipeline      = 5
	PromQLPlusParserRULE_pipelineStep  = 6
	PromQLPlusParserRULE_selector      = 7
	PromQLPlusParserRULE_labelMatcher  = 8
	PromQLPlusParserRULE_aggregation   = 9
	PromQLPlusParserRULE_aligner       = 10
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
	p.RuleIndex = PromQLPlusParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLPlusParserRULE_query

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
	return s.GetToken(PromQLPlusParserEOF, 0)
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
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromQLPlusVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromQLPlusParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PromQLPlusParserRULE_query)
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLPlusParserLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.LetExpression()
		}
		{
			p.SetState(23)
			p.Match(PromQLPlusParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromQLPlusParserEOF, PromQLPlusParserT__6, PromQLPlusParserT__7, PromQLPlusParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Pipeline()
		}
		{
			p.SetState(26)
			p.Match(PromQLPlusParserEOF)
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
	LetBindings() ILetBindingsContext
	IN() antlr.TerminalNode
	Expression() IExpressionContext

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
	p.RuleIndex = PromQLPlusParserRULE_letExpression
	return p
}

func InitEmptyLetExpressionContext(p *LetExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_letExpression
}

func (*LetExpressionContext) IsLetExpressionContext() {}

func NewLetExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetExpressionContext {
	var p = new(LetExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLPlusParserRULE_letExpression

	return p
}

func (s *LetExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LetExpressionContext) LET() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserLET, 0)
}

func (s *LetExpressionContext) LetBindings() ILetBindingsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILetBindingsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILetBindingsContext)
}

func (s *LetExpressionContext) IN() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserIN, 0)
}

func (s *LetExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LetExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LetExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.EnterLetExpression(s)
	}
}

func (s *LetExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.ExitLetExpression(s)
	}
}

func (s *LetExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromQLPlusVisitor:
		return t.VisitLetExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromQLPlusParser) LetExpression() (localctx ILetExpressionContext) {
	localctx = NewLetExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PromQLPlusParserRULE_letExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
		p.Match(PromQLPlusParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(31)
		p.LetBindings()
	}
	{
		p.SetState(32)
		p.Match(PromQLPlusParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.Expression()
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

// ILetBindingsContext is an interface to support dynamic dispatch.
type ILetBindingsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLetBinding() []ILetBindingContext
	LetBinding(i int) ILetBindingContext

	// IsLetBindingsContext differentiates from other interfaces.
	IsLetBindingsContext()
}

type LetBindingsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetBindingsContext() *LetBindingsContext {
	var p = new(LetBindingsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_letBindings
	return p
}

func InitEmptyLetBindingsContext(p *LetBindingsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_letBindings
}

func (*LetBindingsContext) IsLetBindingsContext() {}

func NewLetBindingsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetBindingsContext {
	var p = new(LetBindingsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLPlusParserRULE_letBindings

	return p
}

func (s *LetBindingsContext) GetParser() antlr.Parser { return s.parser }

func (s *LetBindingsContext) AllLetBinding() []ILetBindingContext {
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

func (s *LetBindingsContext) LetBinding(i int) ILetBindingContext {
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

func (s *LetBindingsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetBindingsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LetBindingsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.EnterLetBindings(s)
	}
}

func (s *LetBindingsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.ExitLetBindings(s)
	}
}

func (s *LetBindingsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromQLPlusVisitor:
		return t.VisitLetBindings(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromQLPlusParser) LetBindings() (localctx ILetBindingsContext) {
	localctx = NewLetBindingsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PromQLPlusParserRULE_letBindings)
	var _la int

	p.EnterOuterAlt(localctx, 1)
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

	for _la == PromQLPlusParserT__0 {
		{
			p.SetState(36)
			p.Match(PromQLPlusParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.LetBinding()
		}

		p.SetState(42)
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
	p.RuleIndex = PromQLPlusParserRULE_letBinding
	return p
}

func InitEmptyLetBindingContext(p *LetBindingContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_letBinding
}

func (*LetBindingContext) IsLetBindingContext() {}

func NewLetBindingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetBindingContext {
	var p = new(LetBindingContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLPlusParserRULE_letBinding

	return p
}

func (s *LetBindingContext) GetParser() antlr.Parser { return s.parser }

func (s *LetBindingContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserIDENTIFIER, 0)
}

func (s *LetBindingContext) MATCH_EQ() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserMATCH_EQ, 0)
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
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.EnterLetBinding(s)
	}
}

func (s *LetBindingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.ExitLetBinding(s)
	}
}

func (s *LetBindingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromQLPlusVisitor:
		return t.VisitLetBinding(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromQLPlusParser) LetBinding() (localctx ILetBindingContext) {
	localctx = NewLetBindingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PromQLPlusParserRULE_letBinding)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(PromQLPlusParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(PromQLPlusParserMATCH_EQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLPlusParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PromQLPlusParserIDENTIFIER)
}

func (s *ExpressionContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserIDENTIFIER, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromQLPlusVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromQLPlusParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PromQLPlusParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(PromQLPlusParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&124) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(49)
		p.Match(PromQLPlusParserIDENTIFIER)
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
	p.RuleIndex = PromQLPlusParserRULE_pipeline
	return p
}

func InitEmptyPipelineContext(p *PipelineContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_pipeline
}

func (*PipelineContext) IsPipelineContext() {}

func NewPipelineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipelineContext {
	var p = new(PipelineContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLPlusParserRULE_pipeline

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
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.EnterPipeline(s)
	}
}

func (s *PipelineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.ExitPipeline(s)
	}
}

func (s *PipelineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromQLPlusVisitor:
		return t.VisitPipeline(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromQLPlusParser) Pipeline() (localctx IPipelineContext) {
	localctx = NewPipelineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PromQLPlusParserRULE_pipeline)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Selector()
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromQLPlusParserT__6 {
		{
			p.SetState(52)
			p.Match(PromQLPlusParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.PipelineStep()
		}

		p.SetState(58)
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
	p.RuleIndex = PromQLPlusParserRULE_pipelineStep
	return p
}

func InitEmptyPipelineStepContext(p *PipelineStepContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_pipelineStep
}

func (*PipelineStepContext) IsPipelineStepContext() {}

func NewPipelineStepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PipelineStepContext {
	var p = new(PipelineStepContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLPlusParserRULE_pipelineStep

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
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.EnterPipelineStep(s)
	}
}

func (s *PipelineStepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.ExitPipelineStep(s)
	}
}

func (s *PipelineStepContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromQLPlusVisitor:
		return t.VisitPipelineStep(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromQLPlusParser) PipelineStep() (localctx IPipelineStepContext) {
	localctx = NewPipelineStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PromQLPlusParserRULE_pipelineStep)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLPlusParserAGGREGATION_OP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Aggregation()
		}

	case PromQLPlusParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
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
	p.RuleIndex = PromQLPlusParserRULE_selector
	return p
}

func InitEmptySelectorContext(p *SelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_selector
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLPlusParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserIDENTIFIER, 0)
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
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.EnterSelector(s)
	}
}

func (s *SelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.ExitSelector(s)
	}
}

func (s *SelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromQLPlusVisitor:
		return t.VisitSelector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromQLPlusParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PromQLPlusParserRULE_selector)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLPlusParserIDENTIFIER {
		{
			p.SetState(63)
			p.Match(PromQLPlusParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLPlusParserT__7 {
		{
			p.SetState(66)
			p.Match(PromQLPlusParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PromQLPlusParserIDENTIFIER {
			{
				p.SetState(67)
				p.LabelMatcher()
			}
			p.SetState(72)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == PromQLPlusParserT__0 {
				{
					p.SetState(68)
					p.Match(PromQLPlusParserT__0)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(69)
					p.LabelMatcher()
				}

				p.SetState(74)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(77)
			p.Match(PromQLPlusParserT__8)
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
	p.RuleIndex = PromQLPlusParserRULE_labelMatcher
	return p
}

func InitEmptyLabelMatcherContext(p *LabelMatcherContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_labelMatcher
}

func (*LabelMatcherContext) IsLabelMatcherContext() {}

func NewLabelMatcherContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelMatcherContext {
	var p = new(LabelMatcherContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLPlusParserRULE_labelMatcher

	return p
}

func (s *LabelMatcherContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelMatcherContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserIDENTIFIER, 0)
}

func (s *LabelMatcherContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserSTRING, 0)
}

func (s *LabelMatcherContext) MATCH_EQ() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserMATCH_EQ, 0)
}

func (s *LabelMatcherContext) MATCH_NEQ() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserMATCH_NEQ, 0)
}

func (s *LabelMatcherContext) MATCH_RE() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserMATCH_RE, 0)
}

func (s *LabelMatcherContext) MATCH_NRE() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserMATCH_NRE, 0)
}

func (s *LabelMatcherContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelMatcherContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelMatcherContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.EnterLabelMatcher(s)
	}
}

func (s *LabelMatcherContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.ExitLabelMatcher(s)
	}
}

func (s *LabelMatcherContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromQLPlusVisitor:
		return t.VisitLabelMatcher(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromQLPlusParser) LabelMatcher() (localctx ILabelMatcherContext) {
	localctx = NewLabelMatcherContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PromQLPlusParserRULE_labelMatcher)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(PromQLPlusParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&61440) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(82)
		p.Match(PromQLPlusParserSTRING)
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
	p.RuleIndex = PromQLPlusParserRULE_aggregation
	return p
}

func InitEmptyAggregationContext(p *AggregationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_aggregation
}

func (*AggregationContext) IsAggregationContext() {}

func NewAggregationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregationContext {
	var p = new(AggregationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLPlusParserRULE_aggregation

	return p
}

func (s *AggregationContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregationContext) AGGREGATION_OP() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserAGGREGATION_OP, 0)
}

func (s *AggregationContext) BY() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserBY, 0)
}

func (s *AggregationContext) WITHOUT() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserWITHOUT, 0)
}

func (s *AggregationContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PromQLPlusParserIDENTIFIER)
}

func (s *AggregationContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserIDENTIFIER, i)
}

func (s *AggregationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.EnterAggregation(s)
	}
}

func (s *AggregationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.ExitAggregation(s)
	}
}

func (s *AggregationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromQLPlusVisitor:
		return t.VisitAggregation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromQLPlusParser) Aggregation() (localctx IAggregationContext) {
	localctx = NewAggregationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PromQLPlusParserRULE_aggregation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(PromQLPlusParserAGGREGATION_OP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(85)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLPlusParserBY || _la == PromQLPlusParserWITHOUT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLPlusParserT__9 {
		{
			p.SetState(86)
			p.Match(PromQLPlusParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PromQLPlusParserIDENTIFIER {
			{
				p.SetState(87)
				p.Match(PromQLPlusParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == PromQLPlusParserT__0 {
				{
					p.SetState(88)
					p.Match(PromQLPlusParserT__0)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(89)
					p.Match(PromQLPlusParserIDENTIFIER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(94)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(97)
			p.Match(PromQLPlusParserT__10)
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
	p.RuleIndex = PromQLPlusParserRULE_aligner
	return p
}

func InitEmptyAlignerContext(p *AlignerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_aligner
}

func (*AlignerContext) IsAlignerContext() {}

func NewAlignerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlignerContext {
	var p = new(AlignerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLPlusParserRULE_aligner

	return p
}

func (s *AlignerContext) GetParser() antlr.Parser { return s.parser }

func (s *AlignerContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserIDENTIFIER, 0)
}

func (s *AlignerContext) DURATION() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserDURATION, 0)
}

func (s *AlignerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlignerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlignerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.EnterAligner(s)
	}
}

func (s *AlignerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.ExitAligner(s)
	}
}

func (s *AlignerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromQLPlusVisitor:
		return t.VisitAligner(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromQLPlusParser) Aligner() (localctx IAlignerContext) {
	localctx = NewAlignerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PromQLPlusParserRULE_aligner)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(PromQLPlusParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(PromQLPlusParserDURATION)
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
