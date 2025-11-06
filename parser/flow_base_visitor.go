// Code generated from grammar/Flow.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Flow
import "github.com/antlr4-go/antlr/v4"

type BaseFlowVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFlowVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitLetExpression(ctx *LetExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitLetBinding(ctx *LetBindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitBinaryOperator(ctx *BinaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitPipeline(ctx *PipelineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitPipelineStep(ctx *PipelineStepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitLabelMatcher(ctx *LabelMatcherContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitAggregation(ctx *AggregationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitAligner(ctx *AlignerContext) interface{} {
	return v.VisitChildren(ctx)
}
