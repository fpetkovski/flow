// Code generated from Flow.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

func (v *BaseFlowVisitor) VisitBinarySelector(ctx *BinarySelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitBinarySelectorLeg(ctx *BinarySelectorLegContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitBinaryLegLeaf(ctx *BinaryLegLeafContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitBinaryOperator(ctx *BinaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitAggregation(ctx *AggregationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitAligner(ctx *AlignerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFlowVisitor) VisitBinary(ctx *BinaryContext) interface{} {
	return v.VisitChildren(ctx)
}
