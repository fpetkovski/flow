// Code generated from grammar/PromQLPlus.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PromQLPlus
import "github.com/antlr4-go/antlr/v4"

// BasePromQLPlusListener is a complete listener for a parse tree produced by PromQLPlusParser.
type BasePromQLPlusListener struct{}

var _ PromQLPlusListener = &BasePromQLPlusListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePromQLPlusListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePromQLPlusListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePromQLPlusListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePromQLPlusListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BasePromQLPlusListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BasePromQLPlusListener) ExitQuery(ctx *QueryContext) {}

// EnterLetExpression is called when production letExpression is entered.
func (s *BasePromQLPlusListener) EnterLetExpression(ctx *LetExpressionContext) {}

// ExitLetExpression is called when production letExpression is exited.
func (s *BasePromQLPlusListener) ExitLetExpression(ctx *LetExpressionContext) {}

// EnterLetBindings is called when production letBindings is entered.
func (s *BasePromQLPlusListener) EnterLetBindings(ctx *LetBindingsContext) {}

// ExitLetBindings is called when production letBindings is exited.
func (s *BasePromQLPlusListener) ExitLetBindings(ctx *LetBindingsContext) {}

// EnterLetBinding is called when production letBinding is entered.
func (s *BasePromQLPlusListener) EnterLetBinding(ctx *LetBindingContext) {}

// ExitLetBinding is called when production letBinding is exited.
func (s *BasePromQLPlusListener) ExitLetBinding(ctx *LetBindingContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePromQLPlusListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePromQLPlusListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPipeline is called when production pipeline is entered.
func (s *BasePromQLPlusListener) EnterPipeline(ctx *PipelineContext) {}

// ExitPipeline is called when production pipeline is exited.
func (s *BasePromQLPlusListener) ExitPipeline(ctx *PipelineContext) {}

// EnterPipelineStep is called when production pipelineStep is entered.
func (s *BasePromQLPlusListener) EnterPipelineStep(ctx *PipelineStepContext) {}

// ExitPipelineStep is called when production pipelineStep is exited.
func (s *BasePromQLPlusListener) ExitPipelineStep(ctx *PipelineStepContext) {}

// EnterSelector is called when production selector is entered.
func (s *BasePromQLPlusListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BasePromQLPlusListener) ExitSelector(ctx *SelectorContext) {}

// EnterLabelMatcher is called when production labelMatcher is entered.
func (s *BasePromQLPlusListener) EnterLabelMatcher(ctx *LabelMatcherContext) {}

// ExitLabelMatcher is called when production labelMatcher is exited.
func (s *BasePromQLPlusListener) ExitLabelMatcher(ctx *LabelMatcherContext) {}

// EnterAggregation is called when production aggregation is entered.
func (s *BasePromQLPlusListener) EnterAggregation(ctx *AggregationContext) {}

// ExitAggregation is called when production aggregation is exited.
func (s *BasePromQLPlusListener) ExitAggregation(ctx *AggregationContext) {}

// EnterAligner is called when production aligner is entered.
func (s *BasePromQLPlusListener) EnterAligner(ctx *AlignerContext) {}

// ExitAligner is called when production aligner is exited.
func (s *BasePromQLPlusListener) ExitAligner(ctx *AlignerContext) {}
