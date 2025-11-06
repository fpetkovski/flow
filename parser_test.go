package main

import (
	"fmt"
	"testing"

	"github.com/fpetkovski/let-ql/parser"
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	cases := []struct {
		name      string
		query     string
		expectErr bool
	}{
		{
			name:  "selector",
			query: `http_requests_total`,
		},
		{
			name:      "selector",
			query:     `http_requests_total{`,
			expectErr: true,
		},
		{
			name:  "selector with matchers",
			query: `http_requests_total{method="POST"}`,
		},
		{
			name:  "single-label aggregation",
			query: `http_requests_total{method="POST"} | sum by (method)`,
		},
		{
			name:  "multi-label aggregation",
			query: `http_requests_total{method="POST"} | sum by (endpoint, method)`,
		},
		{
			name:  "combination of aggregation and aligner",
			query: `http_requests_total | sum by (endpoint) | rate 5m | sum by ()`,
		},
		{
			name: "let expression",
			query: `
let
	req_get = http_requests_total{method="GET"} | rate 5m | sum by (endpoint),
	req_total = http_requests_total | rate 5m | sum by (endpoint)
in req_get / req_total`,
		},
	}
	//p := parser.NewPromQLPlusParser(nil)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ast, errs := parseQuery(tc.query)
			if tc.expectErr {
				fmt.Println(errs)
				require.NotEmpty(t, errs)
			} else {
				require.Empty(t, errs)
				antlr.NewParseTreeWalker().Walk(&antlr.BaseParseTreeListener{}, ast)
			}
			v := parser.BasePromQLPlusVisitor{}
			v.Visit(ast)
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
	p := parser.NewPromQLPlusParser(
		antlr.NewCommonTokenStream(
			parser.NewPromQLPlusLexer(antlr.NewInputStream(query)),
			antlr.TokenDefaultChannel,
		))

	// Add custom error listener
	errorListener := &customErrorListener{}
	p.AddErrorListener(errorListener)

	return p.Query(), errorListener.errors
}

// printTree recursively prints the parse tree
func printTree(tree antlr.Tree, parser antlr.Parser, indent string) {
	switch t := tree.(type) {
	case antlr.TerminalNode:
		fmt.Printf("%s%s ", indent, t.GetText())
	case antlr.RuleNode:
		for i := 0; i < t.GetChildCount(); i++ {
			printTree(t.GetChild(i), parser, indent)
		}
	}
}
