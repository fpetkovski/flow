// Code generated from grammar/PromQLPlus.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PromQLPlus
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PromQLPlusParser.
type PromQLPlusVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PromQLPlusParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by PromQLPlusParser#letExpression.
	VisitLetExpression(ctx *LetExpressionContext) interface{}

	// Visit a parse tree produced by PromQLPlusParser#letBindings.
	VisitLetBindings(ctx *LetBindingsContext) interface{}

	// Visit a parse tree produced by PromQLPlusParser#letBinding.
	VisitLetBinding(ctx *LetBindingContext) interface{}

	// Visit a parse tree produced by PromQLPlusParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by PromQLPlusParser#pipeline.
	VisitPipeline(ctx *PipelineContext) interface{}

	// Visit a parse tree produced by PromQLPlusParser#pipelineStep.
	VisitPipelineStep(ctx *PipelineStepContext) interface{}

	// Visit a parse tree produced by PromQLPlusParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by PromQLPlusParser#labelMatcher.
	VisitLabelMatcher(ctx *LabelMatcherContext) interface{}

	// Visit a parse tree produced by PromQLPlusParser#aggregation.
	VisitAggregation(ctx *AggregationContext) interface{}

	// Visit a parse tree produced by PromQLPlusParser#aligner.
	VisitAligner(ctx *AlignerContext) interface{}
}
