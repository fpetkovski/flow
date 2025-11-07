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

	// EnterPipeline is called when entering the pipeline production.
	EnterPipeline(c *PipelineContext)

	// EnterPipelineStep is called when entering the pipelineStep production.
	EnterPipelineStep(c *PipelineStepContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterLabelMatcher is called when entering the labelMatcher production.
	EnterLabelMatcher(c *LabelMatcherContext)

	// EnterBinarySelector is called when entering the binarySelector production.
	EnterBinarySelector(c *BinarySelectorContext)

	// EnterBinarySelectorLeg is called when entering the binarySelectorLeg production.
	EnterBinarySelectorLeg(c *BinarySelectorLegContext)

	// EnterBinaryLegLeaf is called when entering the binaryLegLeaf production.
	EnterBinaryLegLeaf(c *BinaryLegLeafContext)

	// EnterBinaryOperator is called when entering the binaryOperator production.
	EnterBinaryOperator(c *BinaryOperatorContext)

	// EnterAggregation is called when entering the aggregation production.
	EnterAggregation(c *AggregationContext)

	// EnterAligner is called when entering the aligner production.
	EnterAligner(c *AlignerContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterBinary is called when entering the binary production.
	EnterBinary(c *BinaryContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitLetExpression is called when exiting the letExpression production.
	ExitLetExpression(c *LetExpressionContext)

	// ExitLetBinding is called when exiting the letBinding production.
	ExitLetBinding(c *LetBindingContext)

	// ExitPipeline is called when exiting the pipeline production.
	ExitPipeline(c *PipelineContext)

	// ExitPipelineStep is called when exiting the pipelineStep production.
	ExitPipelineStep(c *PipelineStepContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitLabelMatcher is called when exiting the labelMatcher production.
	ExitLabelMatcher(c *LabelMatcherContext)

	// ExitBinarySelector is called when exiting the binarySelector production.
	ExitBinarySelector(c *BinarySelectorContext)

	// ExitBinarySelectorLeg is called when exiting the binarySelectorLeg production.
	ExitBinarySelectorLeg(c *BinarySelectorLegContext)

	// ExitBinaryLegLeaf is called when exiting the binaryLegLeaf production.
	ExitBinaryLegLeaf(c *BinaryLegLeafContext)

	// ExitBinaryOperator is called when exiting the binaryOperator production.
	ExitBinaryOperator(c *BinaryOperatorContext)

	// ExitAggregation is called when exiting the aggregation production.
	ExitAggregation(c *AggregationContext)

	// ExitAligner is called when exiting the aligner production.
	ExitAligner(c *AlignerContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitBinary is called when exiting the binary production.
	ExitBinary(c *BinaryContext)
}
