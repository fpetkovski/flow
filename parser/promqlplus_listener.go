// Code generated from grammar/PromQLPlus.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PromQLPlus
import "github.com/antlr4-go/antlr/v4"

// PromQLPlusListener is a complete listener for a parse tree produced by PromQLPlusParser.
type PromQLPlusListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterLetExpression is called when entering the letExpression production.
	EnterLetExpression(c *LetExpressionContext)

	// EnterLetBindings is called when entering the letBindings production.
	EnterLetBindings(c *LetBindingsContext)

	// EnterLetBinding is called when entering the letBinding production.
	EnterLetBinding(c *LetBindingContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPipeline is called when entering the pipeline production.
	EnterPipeline(c *PipelineContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterMetricIdentifier is called when entering the metricIdentifier production.
	EnterMetricIdentifier(c *MetricIdentifierContext)

	// EnterLabelMatchers is called when entering the labelMatchers production.
	EnterLabelMatchers(c *LabelMatchersContext)

	// EnterLabelMatcher is called when entering the labelMatcher production.
	EnterLabelMatcher(c *LabelMatcherContext)

	// EnterMatchOp is called when entering the matchOp production.
	EnterMatchOp(c *MatchOpContext)

	// EnterAggregation is called when entering the aggregation production.
	EnterAggregation(c *AggregationContext)

	// EnterLabelList is called when entering the labelList production.
	EnterLabelList(c *LabelListContext)

	// EnterAggregationOp is called when entering the aggregationOp production.
	EnterAggregationOp(c *AggregationOpContext)

	// EnterAligner is called when entering the aligner production.
	EnterAligner(c *AlignerContext)

	// EnterDuration is called when entering the duration production.
	EnterDuration(c *DurationContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitLetExpression is called when exiting the letExpression production.
	ExitLetExpression(c *LetExpressionContext)

	// ExitLetBindings is called when exiting the letBindings production.
	ExitLetBindings(c *LetBindingsContext)

	// ExitLetBinding is called when exiting the letBinding production.
	ExitLetBinding(c *LetBindingContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPipeline is called when exiting the pipeline production.
	ExitPipeline(c *PipelineContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitMetricIdentifier is called when exiting the metricIdentifier production.
	ExitMetricIdentifier(c *MetricIdentifierContext)

	// ExitLabelMatchers is called when exiting the labelMatchers production.
	ExitLabelMatchers(c *LabelMatchersContext)

	// ExitLabelMatcher is called when exiting the labelMatcher production.
	ExitLabelMatcher(c *LabelMatcherContext)

	// ExitMatchOp is called when exiting the matchOp production.
	ExitMatchOp(c *MatchOpContext)

	// ExitAggregation is called when exiting the aggregation production.
	ExitAggregation(c *AggregationContext)

	// ExitLabelList is called when exiting the labelList production.
	ExitLabelList(c *LabelListContext)

	// ExitAggregationOp is called when exiting the aggregationOp production.
	ExitAggregationOp(c *AggregationOpContext)

	// ExitAligner is called when exiting the aligner production.
	ExitAligner(c *AlignerContext)

	// ExitDuration is called when exiting the duration production.
	ExitDuration(c *DurationContext)
}
