// Code generated from grammar/PromQLPlus.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PromQLPlus
import "github.com/antlr4-go/antlr/v4"

type BasePromQLPlusVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePromQLPlusVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLPlusVisitor) VisitLetExpression(ctx *LetExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLPlusVisitor) VisitLetBindings(ctx *LetBindingsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLPlusVisitor) VisitLetBinding(ctx *LetBindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLPlusVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLPlusVisitor) VisitPipeline(ctx *PipelineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLPlusVisitor) VisitPipelineStep(ctx *PipelineStepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLPlusVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLPlusVisitor) VisitLabelMatcher(ctx *LabelMatcherContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLPlusVisitor) VisitAggregation(ctx *AggregationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePromQLPlusVisitor) VisitAligner(ctx *AlignerContext) interface{} {
	return v.VisitChildren(ctx)
}
