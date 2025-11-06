package transpiler

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/fpetkovski/let-ql/parser"
	promqlparser "github.com/prometheus/prometheus/promql/parser"
	"github.com/stretchr/testify/require"
)

var (
	reSpaces = regexp.MustCompile("\\s")
)

func TestParser(t *testing.T) {
	cases := []struct {
		name      string
		query     string
		expectErr bool

		promql string
	}{
		{
			name:   "selector with metric",
			query:  `http_requests_total`,
			promql: `http_requests_total`,
		},
		{
			name:   "selector with and empty matchers",
			query:  `http_requests_total{}`,
			promql: `http_requests_total`,
		},
		{
			name:   "selector without metric",
			query:  `{job="prometheus"}`,
			promql: `{job="prometheus"}`,
		},
		{
			name:      "selector",
			query:     `http_requests_total{`,
			expectErr: true,
		},
		{
			name:   "selector with matchers",
			query:  `http_requests_total{method="POST"}`,
			promql: `http_requests_total{method="POST"}`,
		},
		{
			name:   "single-label aggregation",
			query:  `http_requests_total{method="POST"} | sum by (method)`,
			promql: `sum by (method) (http_requests_total{method="POST"})`,
		},
		{
			name:   "multi-label aggregation",
			query:  `http_requests_total{method="POST"} | sum by (endpoint, method)`,
			promql: `sum by (endpoint, method) (http_requests_total{method="POST"})`,
		},
		{
			name:   "combination of aggregation and aligner",
			query:  `http_requests_total | rate 5m | sum by (endpoint) | sum by ()`,
			promql: `sum(sum by (endpoint) (rate(http_requests_total[5m])))`,
		},
		{
			name: "let expression",
			query: `
let
	req_get = http_requests_total{method="GET"} | rate 5m | sum by (endpoint),
	req_total = http_requests_total | rate 5m | sum by (endpoint)
in req_get / req_total`,
			promql: `sum by (endpoint) (rate(http_requests_total{method="GET"}[5m])) / sum by (endpoint) (rate(http_requests_total[5m]))`,
		},
		{
			name: "let expression with multiple binops",
			query: `
let
	req_get = http_requests_total{method="GET"} | rate 5m | sum by (endpoint),
	req_total = http_requests_total | rate 5m | sum by (endpoint)
in (req_total - req_get) / req_total`,
			promql: `
(
	sum by (endpoint) (rate(http_requests_total[5m]))
	- 
	sum by (endpoint) (rate(http_requests_total{method="GET"}[5m]))
)
/
sum by (endpoint) (rate(http_requests_total[5m]))
`,
		},
		{
			name: "let expression with multiple binops and constant",
			query: `
let
	req_get = http_requests_total{method="GET"} | rate 5m | sum by (endpoint),
	req_total = http_requests_total | rate 5m | sum by (endpoint)
in 100 * (req_total - req_get) / req_total
`,
			promql: `
100 * (
	sum by (endpoint) (rate(http_requests_total[5m]))
	- 
	sum by (endpoint) (rate(http_requests_total{method="GET"}[5m]))
)
/
sum by (endpoint) (rate(http_requests_total[5m]))
`,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ast, errs := parseQuery(tc.query)
			if tc.expectErr {
				fmt.Println(errs)
				require.NotEmpty(t, errs)
			} else {
				require.Empty(t, errs)
				v := NewPromQLTranspiler()
				gotPromQL, err := promqlparser.ParseExpr(v.Visit(ast).(string))
				require.NoError(t, err)

				wantPromQL, err := promqlparser.ParseExpr(tc.promql)
				require.NoError(t, err)

				require.Equal(t, gotPromQL.String(), wantPromQL.String())
			}
		})
	}
}

// Custom error listener to capture syntax errors
type customErrorListener struct {
	*antlr.DefaultErrorListener
	errors []string
}

func (c *customErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	err := fmt.Sprintf("line %d:%d %s", line, column, msg)
	c.errors = append(c.errors, err)
}

func parseQuery(query string) (parser.IQueryContext, []string) {
	p := parser.NewFlowParser(
		antlr.NewCommonTokenStream(
			parser.NewFlowLexer(antlr.NewInputStream(query)),
			antlr.TokenDefaultChannel,
		))

	// Add custom error listener
	errorListener := &customErrorListener{}
	p.AddErrorListener(errorListener)

	return p.Query(), errorListener.errors
}
