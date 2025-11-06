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

	// EnterPipelineStep is called when entering the pipelineStep production.
	EnterPipelineStep(c *PipelineStepContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterLabelMatcher is called when entering the labelMatcher production.
	EnterLabelMatcher(c *LabelMatcherContext)

	// EnterAggregation is called when entering the aggregation production.
	EnterAggregation(c *AggregationContext)

	// EnterAligner is called when entering the aligner production.
	EnterAligner(c *AlignerContext)

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

	// ExitPipelineStep is called when exiting the pipelineStep production.
	ExitPipelineStep(c *PipelineStepContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitLabelMatcher is called when exiting the labelMatcher production.
	ExitLabelMatcher(c *LabelMatcherContext)

	// ExitAggregation is called when exiting the aggregation production.
	ExitAggregation(c *AggregationContext)

	// ExitAligner is called when exiting the aligner production.
	ExitAligner(c *AlignerContext)
}
