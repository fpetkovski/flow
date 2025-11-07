package transpiler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/fpetkovski/let-ql/parser"
	"github.com/prometheus/prometheus/model/labels"
	promqlparser "github.com/prometheus/prometheus/promql/parser"
)

type PromQLTranspiler struct {
	bindings map[string]promqlparser.Expr
}

func NewPromQLTranspiler() PromQLTranspiler {
	return PromQLTranspiler{
		make(map[string]promqlparser.Expr),
	}
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
	for _, b := range ctx.AllLetBinding() {
		name := b.(*parser.LetBindingContext).IDENTIFIER().GetText()
		expr := v.VisitPipeline(b.(*parser.LetBindingContext).Pipeline().(*parser.PipelineContext)).(promqlparser.Expr)
		v.bindings[name] = expr
	}

	return v.VisitPipeline(ctx.Pipeline().(*parser.PipelineContext)).(promqlparser.Expr)
}

func (v PromQLTranspiler) VisitBinarySelector(ctx *parser.BinarySelectorContext) any {
	exprs := ctx.AllBinarySelectorLeg()

	// Build the expression from left to right
	// Since PromQL will handle precedence, we can just build it linearly
	left := v.VisitBinarySelectorLeg(exprs[0].(*parser.BinarySelectorLegContext)).(promqlparser.Expr)

	operators := ctx.AllBinaryOperator()
	for i := 0; i < len(operators); i++ {
		right := v.VisitBinarySelectorLeg(exprs[i+1].(*parser.BinarySelectorLegContext)).(promqlparser.Expr)
		op := v.VisitBinaryOperator(operators[i].(*parser.BinaryOperatorContext)).(promqlparser.ItemType)

		left = &promqlparser.BinaryExpr{
			Op:  op,
			LHS: left,
			RHS: right,
		}
	}

	return left
}

func (v PromQLTranspiler) VisitBinaryLegLeaf(ctx *parser.BinaryLegLeafContext) interface{} {
	if id := ctx.IDENTIFIER(); id != nil {
		if bound, ok := v.bindings[id.GetText()]; ok {
			return bound
		}
		// If not found in bindings, treat as a metric name
		return &promqlparser.VectorSelector{
			Name: id.GetText(),
		}
	}
	if n := ctx.NUMBER(); n != nil {
		val, err := strconv.ParseFloat(n.GetText(), 64)
		if err != nil {
			panic(err)
		}
		return &promqlparser.NumberLiteral{
			Val: val,
		}
	}
	panic("unknown leaf")
}

func (v PromQLTranspiler) VisitBinaryOperator(ctx *parser.BinaryOperatorContext) any {
	var op promqlparser.ItemType
	switch {
	case ctx.OP_ADD() != nil:
		op = promqlparser.ADD
	case ctx.OP_SUB() != nil:
		op = promqlparser.SUB
	case ctx.OP_MUL() != nil:
		op = promqlparser.MUL
	case ctx.OP_DIV() != nil:
		op = promqlparser.DIV
	case ctx.OP_GT() != nil:
		op = promqlparser.GTR
	case ctx.OP_LT() != nil:
		op = promqlparser.LSS
	case ctx.OP_GTE() != nil:
		op = promqlparser.GTE
	case ctx.OP_LTE() != nil:
		op = promqlparser.LTE
	case ctx.OP_EQ() != nil:
		op = promqlparser.EQLC
	case ctx.MATCH_NEQ() != nil:
		op = promqlparser.NEQ
	default:
		panic("unknown binary operator")
	}
	return op
}

func (v PromQLTranspiler) VisitBinarySelectorLeg(ctx *parser.BinarySelectorLegContext) any {
	if ctx.BinaryLegLeaf() != nil {
		return v.VisitBinaryLegLeaf(ctx.BinaryLegLeaf().(*parser.BinaryLegLeafContext))
	}

	// Must be a parenthesized expression - preserve the parentheses
	expr := v.VisitBinarySelector(ctx.BinarySelector().(*parser.BinarySelectorContext)).(promqlparser.Expr)
	return &promqlparser.ParenExpr{
		Expr: expr,
	}
}

func (v PromQLTranspiler) VisitLetBinding(ctx *parser.LetBindingContext) any {
	name := ctx.IDENTIFIER().GetText()
	expr := v.VisitPipeline(ctx.Pipeline().(*parser.PipelineContext))

	fmt.Println(name, expr)
	return ""
}

func (v PromQLTranspiler) VisitPipeline(ctx *parser.PipelineContext) any {
	var selector promqlparser.Expr
	switch {
	case ctx.Selector() != nil:
		selector = v.VisitSelector(ctx.Selector().(*parser.SelectorContext)).(promqlparser.Expr)
	case ctx.BinarySelector() != nil:
		selector = v.VisitBinarySelector(ctx.BinarySelector().(*parser.BinarySelectorContext)).(promqlparser.Expr)
	}

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
			// Check if this is an aligner (has MatrixSelector) or a function
			if len(n.Args) > 0 {
				switch m := n.Args[0].(type) {
				case *promqlparser.MatrixSelector:
					// This is an aligner like rate, irate, etc.
					m.VectorSelector = root
				case nil:
					// This is a function that needs the previous expression as its argument
					n.Args[0] = root
				}
			} else {
				// No args set yet, add the root as the first argument
				n.Args = append(n.Args, root)
			}
		case *promqlparser.BinaryExpr:
			n.LHS = root
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
	case ctx.Function() != nil:
		return v.VisitFunction(ctx.Function().(*parser.FunctionContext))
	case ctx.Binary() != nil:
		return v.VisitBinary(ctx.Binary().(*parser.BinaryContext))
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

func (v PromQLTranspiler) VisitFunction(ctx *parser.FunctionContext) any {
	// Functions like histogram_sum operate on the previous pipeline step
	call := &promqlparser.Call{
		Func: &promqlparser.Function{
			Name: ctx.IDENTIFIER().GetText(),
		},
		Args: promqlparser.Expressions{}, // Empty args, will be filled by pipeline
	}
	return call
}

func (v PromQLTranspiler) VisitAligner(ctx *parser.AlignerContext) any {
	call := &promqlparser.Call{
		Func: &promqlparser.Function{
			Name: ctx.IDENTIFIER().GetText(),
		},
		Args: make(promqlparser.Expressions, 1, 1),
	}

	durationStr := ctx.DURATION().GetText()
	duration, err := parseDuration(durationStr)
	if err != nil {
		panic("invalid duration: " + durationStr)
	}

	// Create a MatrixSelector with the duration
	call.Args[0] = &promqlparser.MatrixSelector{
		VectorSelector: nil, // Will be filled by pipeline
		Range:          duration,
	}

	return call
}

func (v PromQLTranspiler) VisitBinary(ctx *parser.BinaryContext) interface{} {
	return &promqlparser.BinaryExpr{
		Op:  v.VisitBinaryOperator(ctx.BinaryOperator().(*parser.BinaryOperatorContext)).(promqlparser.ItemType),
		RHS: v.VisitBinaryLegLeaf(ctx.BinaryLegLeaf().(*parser.BinaryLegLeafContext)).(promqlparser.Expr),
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

func (v PromQLTranspiler) VisitChildren(node antlr.RuleNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v PromQLTranspiler) VisitTerminal(node antlr.TerminalNode) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v PromQLTranspiler) VisitErrorNode(node antlr.ErrorNode) interface{} {
	//TODO implement me
	panic("implement me")
}

// parseDuration parses a duration string like "5m" into a time.Duration
func parseDuration(s string) (time.Duration, error) {
	// This is a simplified version - in production you'd want full parsing
	return time.ParseDuration(s)
}
