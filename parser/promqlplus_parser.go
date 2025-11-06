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
		"", "','", "'='", "'/'", "'+'", "'-'", "'*'", "'%'", "'|'", "'{'", "'}'",
		"'!='", "'=~'", "'!~'", "'('", "')'", "'sum'", "'min'", "'max'", "'avg'",
		"'stddev'", "'stdvar'", "'count'", "'count_values'", "'bottomk'", "'topk'",
		"'quantile'", "'let'", "'in'", "'by'", "'without'", "'and'", "'or'",
		"'unless'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "LET", "IN", "BY", "WITHOUT",
		"AND", "OR", "UNLESS", "NUMBER", "DURATION_UNIT", "IDENTIFIER", "STRING",
		"COMMENT", "MULTILINE_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"query", "letExpression", "letBindings", "letBinding", "expression",
		"pipeline", "selector", "metricIdentifier", "labelMatchers", "labelMatcher",
		"matchOp", "aggregation", "labelList", "aggregationOp", "aligner", "duration",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 40, 124, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 39, 8, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 49, 8, 2, 10, 2, 12, 2, 52, 9, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5,
		66, 8, 5, 5, 5, 68, 8, 5, 10, 5, 12, 5, 71, 9, 5, 1, 6, 1, 6, 3, 6, 75,
		8, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 83, 8, 8, 10, 8, 12, 8,
		86, 9, 8, 1, 8, 3, 8, 89, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 103, 8, 11, 1, 11, 3, 11,
		106, 8, 11, 1, 12, 1, 12, 1, 12, 5, 12, 111, 8, 12, 10, 12, 12, 12, 114,
		9, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 0,
		0, 16, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 0, 4,
		1, 0, 3, 7, 2, 0, 2, 2, 11, 13, 1, 0, 29, 30, 1, 0, 16, 26, 117, 0, 38,
		1, 0, 0, 0, 2, 40, 1, 0, 0, 0, 4, 45, 1, 0, 0, 0, 6, 53, 1, 0, 0, 0, 8,
		57, 1, 0, 0, 0, 10, 61, 1, 0, 0, 0, 12, 72, 1, 0, 0, 0, 14, 76, 1, 0, 0,
		0, 16, 78, 1, 0, 0, 0, 18, 92, 1, 0, 0, 0, 20, 96, 1, 0, 0, 0, 22, 98,
		1, 0, 0, 0, 24, 107, 1, 0, 0, 0, 26, 115, 1, 0, 0, 0, 28, 117, 1, 0, 0,
		0, 30, 120, 1, 0, 0, 0, 32, 33, 3, 2, 1, 0, 33, 34, 5, 0, 0, 1, 34, 39,
		1, 0, 0, 0, 35, 36, 3, 10, 5, 0, 36, 37, 5, 0, 0, 1, 37, 39, 1, 0, 0, 0,
		38, 32, 1, 0, 0, 0, 38, 35, 1, 0, 0, 0, 39, 1, 1, 0, 0, 0, 40, 41, 5, 27,
		0, 0, 41, 42, 3, 4, 2, 0, 42, 43, 5, 28, 0, 0, 43, 44, 3, 8, 4, 0, 44,
		3, 1, 0, 0, 0, 45, 50, 3, 6, 3, 0, 46, 47, 5, 1, 0, 0, 47, 49, 3, 6, 3,
		0, 48, 46, 1, 0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51,
		1, 0, 0, 0, 51, 5, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 54, 5, 36, 0, 0,
		54, 55, 5, 2, 0, 0, 55, 56, 3, 10, 5, 0, 56, 7, 1, 0, 0, 0, 57, 58, 5,
		36, 0, 0, 58, 59, 7, 0, 0, 0, 59, 60, 5, 36, 0, 0, 60, 9, 1, 0, 0, 0, 61,
		69, 3, 12, 6, 0, 62, 65, 5, 8, 0, 0, 63, 66, 3, 22, 11, 0, 64, 66, 3, 28,
		14, 0, 65, 63, 1, 0, 0, 0, 65, 64, 1, 0, 0, 0, 66, 68, 1, 0, 0, 0, 67,
		62, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0,
		0, 70, 11, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 74, 3, 14, 7, 0, 73, 75,
		3, 16, 8, 0, 74, 73, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 13, 1, 0, 0, 0,
		76, 77, 5, 36, 0, 0, 77, 15, 1, 0, 0, 0, 78, 79, 5, 9, 0, 0, 79, 84, 3,
		18, 9, 0, 80, 81, 5, 1, 0, 0, 81, 83, 3, 18, 9, 0, 82, 80, 1, 0, 0, 0,
		83, 86, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 88, 1,
		0, 0, 0, 86, 84, 1, 0, 0, 0, 87, 89, 5, 1, 0, 0, 88, 87, 1, 0, 0, 0, 88,
		89, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 91, 5, 10, 0, 0, 91, 17, 1, 0,
		0, 0, 92, 93, 5, 36, 0, 0, 93, 94, 3, 20, 10, 0, 94, 95, 5, 37, 0, 0, 95,
		19, 1, 0, 0, 0, 96, 97, 7, 1, 0, 0, 97, 21, 1, 0, 0, 0, 98, 99, 3, 26,
		13, 0, 99, 105, 7, 2, 0, 0, 100, 102, 5, 14, 0, 0, 101, 103, 3, 24, 12,
		0, 102, 101, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104,
		106, 5, 15, 0, 0, 105, 100, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 23,
		1, 0, 0, 0, 107, 112, 5, 36, 0, 0, 108, 109, 5, 1, 0, 0, 109, 111, 5, 36,
		0, 0, 110, 108, 1, 0, 0, 0, 111, 114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0,
		112, 113, 1, 0, 0, 0, 113, 25, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 116,
		7, 3, 0, 0, 116, 27, 1, 0, 0, 0, 117, 118, 5, 36, 0, 0, 118, 119, 3, 30,
		15, 0, 119, 29, 1, 0, 0, 0, 120, 121, 5, 34, 0, 0, 121, 122, 5, 35, 0,
		0, 122, 31, 1, 0, 0, 0, 10, 38, 50, 65, 69, 74, 84, 88, 102, 105, 112,
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
	PromQLPlusParserT__11             = 12
	PromQLPlusParserT__12             = 13
	PromQLPlusParserT__13             = 14
	PromQLPlusParserT__14             = 15
	PromQLPlusParserT__15             = 16
	PromQLPlusParserT__16             = 17
	PromQLPlusParserT__17             = 18
	PromQLPlusParserT__18             = 19
	PromQLPlusParserT__19             = 20
	PromQLPlusParserT__20             = 21
	PromQLPlusParserT__21             = 22
	PromQLPlusParserT__22             = 23
	PromQLPlusParserT__23             = 24
	PromQLPlusParserT__24             = 25
	PromQLPlusParserT__25             = 26
	PromQLPlusParserLET               = 27
	PromQLPlusParserIN                = 28
	PromQLPlusParserBY                = 29
	PromQLPlusParserWITHOUT           = 30
	PromQLPlusParserAND               = 31
	PromQLPlusParserOR                = 32
	PromQLPlusParserUNLESS            = 33
	PromQLPlusParserNUMBER            = 34
	PromQLPlusParserDURATION_UNIT     = 35
	PromQLPlusParserIDENTIFIER        = 36
	PromQLPlusParserSTRING            = 37
	PromQLPlusParserCOMMENT           = 38
	PromQLPlusParserMULTILINE_COMMENT = 39
	PromQLPlusParserWS                = 40
)

// PromQLPlusParser rules.
const (
	PromQLPlusParserRULE_query            = 0
	PromQLPlusParserRULE_letExpression    = 1
	PromQLPlusParserRULE_letBindings      = 2
	PromQLPlusParserRULE_letBinding       = 3
	PromQLPlusParserRULE_expression       = 4
	PromQLPlusParserRULE_pipeline         = 5
	PromQLPlusParserRULE_selector         = 6
	PromQLPlusParserRULE_metricIdentifier = 7
	PromQLPlusParserRULE_labelMatchers    = 8
	PromQLPlusParserRULE_labelMatcher     = 9
	PromQLPlusParserRULE_matchOp          = 10
	PromQLPlusParserRULE_aggregation      = 11
	PromQLPlusParserRULE_labelList        = 12
	PromQLPlusParserRULE_aggregationOp    = 13
	PromQLPlusParserRULE_aligner          = 14
	PromQLPlusParserRULE_duration         = 15
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
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PromQLPlusParserLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.LetExpression()
		}
		{
			p.SetState(33)
			p.Match(PromQLPlusParserEOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PromQLPlusParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Pipeline()
		}
		{
			p.SetState(36)
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
		p.SetState(40)
		p.Match(PromQLPlusParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.LetBindings()
	}
	{
		p.SetState(42)
		p.Match(PromQLPlusParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
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
		p.SetState(45)
		p.LetBinding()
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromQLPlusParserT__0 {
		{
			p.SetState(46)
			p.Match(PromQLPlusParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.LetBinding()
		}

		p.SetState(52)
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
		p.SetState(53)
		p.Match(PromQLPlusParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Match(PromQLPlusParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
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
		p.SetState(57)
		p.Match(PromQLPlusParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&248) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(59)
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
	AllAggregation() []IAggregationContext
	Aggregation(i int) IAggregationContext
	AllAligner() []IAlignerContext
	Aligner(i int) IAlignerContext

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

func (s *PipelineContext) AllAggregation() []IAggregationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAggregationContext); ok {
			len++
		}
	}

	tst := make([]IAggregationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAggregationContext); ok {
			tst[i] = t.(IAggregationContext)
			i++
		}
	}

	return tst
}

func (s *PipelineContext) Aggregation(i int) IAggregationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregationContext); ok {
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

	return t.(IAggregationContext)
}

func (s *PipelineContext) AllAligner() []IAlignerContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAlignerContext); ok {
			len++
		}
	}

	tst := make([]IAlignerContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAlignerContext); ok {
			tst[i] = t.(IAlignerContext)
			i++
		}
	}

	return tst
}

func (s *PipelineContext) Aligner(i int) IAlignerContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAlignerContext); ok {
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

	return t.(IAlignerContext)
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
		p.SetState(61)
		p.Selector()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromQLPlusParserT__7 {
		{
			p.SetState(62)
			p.Match(PromQLPlusParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PromQLPlusParserT__15, PromQLPlusParserT__16, PromQLPlusParserT__17, PromQLPlusParserT__18, PromQLPlusParserT__19, PromQLPlusParserT__20, PromQLPlusParserT__21, PromQLPlusParserT__22, PromQLPlusParserT__23, PromQLPlusParserT__24, PromQLPlusParserT__25:
			{
				p.SetState(63)
				p.Aggregation()
			}

		case PromQLPlusParserIDENTIFIER:
			{
				p.SetState(64)
				p.Aligner()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(71)
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

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MetricIdentifier() IMetricIdentifierContext
	LabelMatchers() ILabelMatchersContext

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

func (s *SelectorContext) MetricIdentifier() IMetricIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetricIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetricIdentifierContext)
}

func (s *SelectorContext) LabelMatchers() ILabelMatchersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelMatchersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelMatchersContext)
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
	p.EnterRule(localctx, 12, PromQLPlusParserRULE_selector)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.MetricIdentifier()
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLPlusParserT__8 {
		{
			p.SetState(73)
			p.LabelMatchers()
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

// IMetricIdentifierContext is an interface to support dynamic dispatch.
type IMetricIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsMetricIdentifierContext differentiates from other interfaces.
	IsMetricIdentifierContext()
}

type MetricIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetricIdentifierContext() *MetricIdentifierContext {
	var p = new(MetricIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_metricIdentifier
	return p
}

func InitEmptyMetricIdentifierContext(p *MetricIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_metricIdentifier
}

func (*MetricIdentifierContext) IsMetricIdentifierContext() {}

func NewMetricIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetricIdentifierContext {
	var p = new(MetricIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLPlusParserRULE_metricIdentifier

	return p
}

func (s *MetricIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *MetricIdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserIDENTIFIER, 0)
}

func (s *MetricIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetricIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetricIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.EnterMetricIdentifier(s)
	}
}

func (s *MetricIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.ExitMetricIdentifier(s)
	}
}

func (s *MetricIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromQLPlusVisitor:
		return t.VisitMetricIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromQLPlusParser) MetricIdentifier() (localctx IMetricIdentifierContext) {
	localctx = NewMetricIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PromQLPlusParserRULE_metricIdentifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
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

// ILabelMatchersContext is an interface to support dynamic dispatch.
type ILabelMatchersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLabelMatcher() []ILabelMatcherContext
	LabelMatcher(i int) ILabelMatcherContext

	// IsLabelMatchersContext differentiates from other interfaces.
	IsLabelMatchersContext()
}

type LabelMatchersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelMatchersContext() *LabelMatchersContext {
	var p = new(LabelMatchersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_labelMatchers
	return p
}

func InitEmptyLabelMatchersContext(p *LabelMatchersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_labelMatchers
}

func (*LabelMatchersContext) IsLabelMatchersContext() {}

func NewLabelMatchersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelMatchersContext {
	var p = new(LabelMatchersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLPlusParserRULE_labelMatchers

	return p
}

func (s *LabelMatchersContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelMatchersContext) AllLabelMatcher() []ILabelMatcherContext {
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

func (s *LabelMatchersContext) LabelMatcher(i int) ILabelMatcherContext {
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

func (s *LabelMatchersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelMatchersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelMatchersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.EnterLabelMatchers(s)
	}
}

func (s *LabelMatchersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.ExitLabelMatchers(s)
	}
}

func (s *LabelMatchersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromQLPlusVisitor:
		return t.VisitLabelMatchers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromQLPlusParser) LabelMatchers() (localctx ILabelMatchersContext) {
	localctx = NewLabelMatchersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PromQLPlusParserRULE_labelMatchers)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(PromQLPlusParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.LabelMatcher()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(80)
				p.Match(PromQLPlusParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(81)
				p.LabelMatcher()
			}

		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLPlusParserT__0 {
		{
			p.SetState(87)
			p.Match(PromQLPlusParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(90)
		p.Match(PromQLPlusParserT__9)
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

// ILabelMatcherContext is an interface to support dynamic dispatch.
type ILabelMatcherContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	MatchOp() IMatchOpContext
	STRING() antlr.TerminalNode

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

func (s *LabelMatcherContext) MatchOp() IMatchOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchOpContext)
}

func (s *LabelMatcherContext) STRING() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserSTRING, 0)
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
	p.EnterRule(localctx, 18, PromQLPlusParserRULE_labelMatcher)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(PromQLPlusParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.MatchOp()
	}
	{
		p.SetState(94)
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

// IMatchOpContext is an interface to support dynamic dispatch.
type IMatchOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMatchOpContext differentiates from other interfaces.
	IsMatchOpContext()
}

type MatchOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchOpContext() *MatchOpContext {
	var p = new(MatchOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_matchOp
	return p
}

func InitEmptyMatchOpContext(p *MatchOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_matchOp
}

func (*MatchOpContext) IsMatchOpContext() {}

func NewMatchOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchOpContext {
	var p = new(MatchOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLPlusParserRULE_matchOp

	return p
}

func (s *MatchOpContext) GetParser() antlr.Parser { return s.parser }
func (s *MatchOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.EnterMatchOp(s)
	}
}

func (s *MatchOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.ExitMatchOp(s)
	}
}

func (s *MatchOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromQLPlusVisitor:
		return t.VisitMatchOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromQLPlusParser) MatchOp() (localctx IMatchOpContext) {
	localctx = NewMatchOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PromQLPlusParserRULE_matchOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14340) != 0) {
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
	AggregationOp() IAggregationOpContext
	BY() antlr.TerminalNode
	WITHOUT() antlr.TerminalNode
	LabelList() ILabelListContext

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

func (s *AggregationContext) AggregationOp() IAggregationOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregationOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregationOpContext)
}

func (s *AggregationContext) BY() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserBY, 0)
}

func (s *AggregationContext) WITHOUT() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserWITHOUT, 0)
}

func (s *AggregationContext) LabelList() ILabelListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILabelListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILabelListContext)
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
	p.EnterRule(localctx, 22, PromQLPlusParserRULE_aggregation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.AggregationOp()
	}
	{
		p.SetState(99)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PromQLPlusParserBY || _la == PromQLPlusParserWITHOUT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PromQLPlusParserT__13 {
		{
			p.SetState(100)
			p.Match(PromQLPlusParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PromQLPlusParserIDENTIFIER {
			{
				p.SetState(101)
				p.LabelList()
			}

		}
		{
			p.SetState(104)
			p.Match(PromQLPlusParserT__14)
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

// ILabelListContext is an interface to support dynamic dispatch.
type ILabelListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

	// IsLabelListContext differentiates from other interfaces.
	IsLabelListContext()
}

type LabelListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelListContext() *LabelListContext {
	var p = new(LabelListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_labelList
	return p
}

func InitEmptyLabelListContext(p *LabelListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_labelList
}

func (*LabelListContext) IsLabelListContext() {}

func NewLabelListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelListContext {
	var p = new(LabelListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLPlusParserRULE_labelList

	return p
}

func (s *LabelListContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelListContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PromQLPlusParserIDENTIFIER)
}

func (s *LabelListContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserIDENTIFIER, i)
}

func (s *LabelListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.EnterLabelList(s)
	}
}

func (s *LabelListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.ExitLabelList(s)
	}
}

func (s *LabelListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromQLPlusVisitor:
		return t.VisitLabelList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromQLPlusParser) LabelList() (localctx ILabelListContext) {
	localctx = NewLabelListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PromQLPlusParserRULE_labelList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(PromQLPlusParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PromQLPlusParserT__0 {
		{
			p.SetState(108)
			p.Match(PromQLPlusParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Match(PromQLPlusParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(114)
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

// IAggregationOpContext is an interface to support dynamic dispatch.
type IAggregationOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAggregationOpContext differentiates from other interfaces.
	IsAggregationOpContext()
}

type AggregationOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregationOpContext() *AggregationOpContext {
	var p = new(AggregationOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_aggregationOp
	return p
}

func InitEmptyAggregationOpContext(p *AggregationOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_aggregationOp
}

func (*AggregationOpContext) IsAggregationOpContext() {}

func NewAggregationOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregationOpContext {
	var p = new(AggregationOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLPlusParserRULE_aggregationOp

	return p
}

func (s *AggregationOpContext) GetParser() antlr.Parser { return s.parser }
func (s *AggregationOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregationOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregationOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.EnterAggregationOp(s)
	}
}

func (s *AggregationOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.ExitAggregationOp(s)
	}
}

func (s *AggregationOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromQLPlusVisitor:
		return t.VisitAggregationOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromQLPlusParser) AggregationOp() (localctx IAggregationOpContext) {
	localctx = NewAggregationOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PromQLPlusParserRULE_aggregationOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&134152192) != 0) {
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

// IAlignerContext is an interface to support dynamic dispatch.
type IAlignerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Duration() IDurationContext

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

func (s *AlignerContext) Duration() IDurationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurationContext)
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
	p.EnterRule(localctx, 28, PromQLPlusParserRULE_aligner)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(PromQLPlusParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(118)
		p.Duration()
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

// IDurationContext is an interface to support dynamic dispatch.
type IDurationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	DURATION_UNIT() antlr.TerminalNode

	// IsDurationContext differentiates from other interfaces.
	IsDurationContext()
}

type DurationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationContext() *DurationContext {
	var p = new(DurationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_duration
	return p
}

func InitEmptyDurationContext(p *DurationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PromQLPlusParserRULE_duration
}

func (*DurationContext) IsDurationContext() {}

func NewDurationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationContext {
	var p = new(DurationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromQLPlusParserRULE_duration

	return p
}

func (s *DurationContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserNUMBER, 0)
}

func (s *DurationContext) DURATION_UNIT() antlr.TerminalNode {
	return s.GetToken(PromQLPlusParserDURATION_UNIT, 0)
}

func (s *DurationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DurationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.EnterDuration(s)
	}
}

func (s *DurationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromQLPlusListener); ok {
		listenerT.ExitDuration(s)
	}
}

func (s *DurationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PromQLPlusVisitor:
		return t.VisitDuration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PromQLPlusParser) Duration() (localctx IDurationContext) {
	localctx = NewDurationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PromQLPlusParserRULE_duration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(PromQLPlusParserNUMBER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Match(PromQLPlusParserDURATION_UNIT)
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
