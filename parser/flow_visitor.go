// Code generated from Flow.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Flow
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by FlowParser.
type FlowVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FlowParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by FlowParser#letExpression.
	VisitLetExpression(ctx *LetExpressionContext) interface{}

	// Visit a parse tree produced by FlowParser#letBinding.
	VisitLetBinding(ctx *LetBindingContext) interface{}

	// Visit a parse tree produced by FlowParser#binaryExpression.
	VisitBinaryExpression(ctx *BinaryExpressionContext) interface{}

	// Visit a parse tree produced by FlowParser#binaryOperator.
	VisitBinaryOperator(ctx *BinaryOperatorContext) interface{}

	// Visit a parse tree produced by FlowParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by FlowParser#pipeline.
	VisitPipeline(ctx *PipelineContext) interface{}

	// Visit a parse tree produced by FlowParser#pipelineStep.
	VisitPipelineStep(ctx *PipelineStepContext) interface{}

	// Visit a parse tree produced by FlowParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by FlowParser#labelMatcher.
	VisitLabelMatcher(ctx *LabelMatcherContext) interface{}

	// Visit a parse tree produced by FlowParser#aggregation.
	VisitAggregation(ctx *AggregationContext) interface{}

	// Visit a parse tree produced by FlowParser#aligner.
	VisitAligner(ctx *AlignerContext) interface{}

	// Visit a parse tree produced by FlowParser#function.
	VisitFunction(ctx *FunctionContext) interface{}
}
