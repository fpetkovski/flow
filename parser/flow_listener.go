// Code generated from Flow.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Flow
import "github.com/antlr4-go/antlr/v4"

// FlowListener is a complete listener for a parse tree produced by FlowParser.
type FlowListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterLetExpression is called when entering the letExpression production.
	EnterLetExpression(c *LetExpressionContext)

	// EnterLetBinding is called when entering the letBinding production.
	EnterLetBinding(c *LetBindingContext)

	// EnterBinaryExpression is called when entering the binaryExpression production.
	EnterBinaryExpression(c *BinaryExpressionContext)

	// EnterBinaryOperator is called when entering the binaryOperator production.
	EnterBinaryOperator(c *BinaryOperatorContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

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

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitLetExpression is called when exiting the letExpression production.
	ExitLetExpression(c *LetExpressionContext)

	// ExitLetBinding is called when exiting the letBinding production.
	ExitLetBinding(c *LetBindingContext)

	// ExitBinaryExpression is called when exiting the binaryExpression production.
	ExitBinaryExpression(c *BinaryExpressionContext)

	// ExitBinaryOperator is called when exiting the binaryOperator production.
	ExitBinaryOperator(c *BinaryOperatorContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

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

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)
}
