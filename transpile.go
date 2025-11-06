package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/fpetkovski/let-ql/parser"
	"github.com/prometheus/prometheus/model/labels"
	promqlparser "github.com/prometheus/prometheus/promql/parser"
)

type PromQLTranspiler struct {
	antlr.ParseTreeVisitor
}

func (v PromQLTranspiler) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v).(promqlparser.Expr).String()
}

func (v PromQLTranspiler) VisitQuery(ctx *parser.QueryContext) any {
	switch {
	case ctx.LetExpression() != nil:
		letCtx := ctx.LetExpression().(*parser.LetExpressionContext)
		return v.VisitLetExpression(letCtx)
	case ctx.Pipeline() != nil:
		pipelineCtx := ctx.Pipeline().(*parser.PipelineContext)
		return v.VisitPipeline(pipelineCtx)
	default:
		panic("unknown type")
	}
}

func (v PromQLTranspiler) VisitLetExpression(ctx *parser.LetExpressionContext) any {
	return v.VisitLetBindings(ctx.LetBindings().(*parser.LetBindingsContext))
}

func (v PromQLTranspiler) VisitLetBindings(ctx *parser.LetBindingsContext) any {
	bindings := make([]string, 0)
	for _, b := range ctx.AllLetBinding() {
		bindings = append(bindings, v.VisitLetBinding(b.(*parser.LetBindingContext)).(string))
	}
	return strings.Join(bindings, ", ")
}

func (v PromQLTranspiler) VisitLetBinding(ctx *parser.LetBindingContext) any {
	name := ctx.IDENTIFIER().GetText()
	expr := v.VisitPipeline(ctx.Pipeline().(*parser.PipelineContext))

	fmt.Println(name, expr)
	return ""
}

func (v PromQLTranspiler) VisitExpression(ctx *parser.ExpressionContext) any {
	fmt.Println(ctx.AllIDENTIFIER())
	return nil
}

func (v PromQLTranspiler) VisitPipeline(ctx *parser.PipelineContext) any {
	selector := v.VisitSelector(ctx.Selector().(*parser.SelectorContext))
	steps := ctx.AllPipelineStep()
	if len(steps) == 0 {
		return selector
	}

	var root = selector.(promqlparser.Expr)
	for _, s := range steps {
		node := v.VisitPipelineStep(s.(*parser.PipelineStepContext))
		switch n := node.(type) {
		case *promqlparser.AggregateExpr:
			n.Expr = root
		case *promqlparser.Call:
			switch m := n.Args[0].(type) {
			case *promqlparser.MatrixSelector:
				m.VectorSelector = root
			}
		}
		root = node.(promqlparser.Expr)
	}
	return root
}

func (v PromQLTranspiler) VisitPipelineStep(ctx *parser.PipelineStepContext) any {
	switch {
	case ctx.Aligner() != nil:
		return v.VisitAligner(ctx.Aligner().(*parser.AlignerContext))
	case ctx.Aggregation() != nil:
		return v.VisitAggregation(ctx.Aggregation().(*parser.AggregationContext))
	default:
		panic("unknown pipeline step")
	}
}

func (v PromQLTranspiler) VisitAggregation(ctx *parser.AggregationContext) any {
	var op promqlparser.ItemType
	switch ctx.AGGREGATION_OP().GetText() {
	case "sum":
		op = promqlparser.SUM
	case "min":
		op = promqlparser.MIN
	case "max":
		op = promqlparser.MAX
	case "avg":
		op = promqlparser.AVG
	case "stddev":
		op = promqlparser.STDDEV
	case "stdvar":
		op = promqlparser.STDVAR
	case "count":
		op = promqlparser.COUNT
	case "count_values":
		op = promqlparser.COUNT_VALUES
	case "bottomk":
		op = promqlparser.BOTTOMK
	case "topk":
		op = promqlparser.TOPK
	case "quantile":
		op = promqlparser.QUANTILE
	}

	idents := ctx.AllIDENTIFIER()
	lbls := make([]string, 0, len(idents))
	for _, id := range idents {
		lbls = append(lbls, id.GetText())
	}
	return &promqlparser.AggregateExpr{
		Op:       op,
		Grouping: lbls,
		Without:  ctx.WITHOUT() != nil,
	}
}

func (v PromQLTranspiler) VisitAligner(ctx *parser.AlignerContext) any {
	d, err := time.ParseDuration(ctx.DURATION().GetText())
	if err != nil {
		panic(err)
	}
	return &promqlparser.Call{
		Func: &promqlparser.Function{
			Name: ctx.IDENTIFIER().GetText(),
		},
		Args: promqlparser.Expressions{
			&promqlparser.MatrixSelector{
				Range: d,
			},
		},
	}
}

func (v PromQLTranspiler) VisitSelector(ctx *parser.SelectorContext) any {
	matchers := make([]*labels.Matcher, 0)
	for _, m := range ctx.AllLabelMatcher() {
		matchers = append(matchers, v.VisitLabelMatcher(m.(*parser.LabelMatcherContext)).(*labels.Matcher))
	}

	var name string
	if ctx.IDENTIFIER() != nil {
		name = ctx.IDENTIFIER().GetText()
	}

	return &promqlparser.VectorSelector{
		Name:          name,
		LabelMatchers: matchers,
	}
}

func (v PromQLTranspiler) VisitLabelMatcher(ctx *parser.LabelMatcherContext) any {
	var matchType labels.MatchType
	switch {
	case ctx.MATCH_EQ() != nil:
		matchType = labels.MatchEqual
	case ctx.MATCH_NEQ() != nil:
		matchType = labels.MatchNotEqual
	case ctx.MATCH_RE() != nil:
		matchType = labels.MatchRegexp
	case ctx.MATCH_NRE() != nil:
		matchType = labels.MatchNotRegexp
	}
	val, err := strconv.Unquote(ctx.STRING().GetText())
	if err != nil {
		panic("invalid value " + ctx.STRING().GetText())
	}
	return &labels.Matcher{
		Type:  matchType,
		Name:  ctx.IDENTIFIER().GetText(),
		Value: val,
	}
}
