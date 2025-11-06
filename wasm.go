//go:build wasm
// +build wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/antlr4-go/antlr/v4"
	"github.com/fpetkovski/let-ql/parser"
	"github.com/fpetkovski/let-ql/transpiler"
)

// Make the transpiler available globally
func main() {
	js.Global().Set("transpileFlow", js.FuncOf(transpileWrapper))

	// Keep the program running
	select {}
}

// transpileWrapper is the JavaScript-callable function
func transpileWrapper(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return createErrorResponse("transpileFlow requires exactly one argument")
	}

	input := args[0].String()

	result := transpile(input)

	// Return a JavaScript object
	return result
}

func transpile(input string) js.Value {
	defer func() {
		if r := recover(); r != nil {
			// Return error on panic
			js.Global().Get("console").Call("error", "Panic:", r)
		}
	}()

	// Create lexer and parser
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewFlowLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewFlowParser(tokenStream)

	// Add error listener
	errorListener := &jsErrorListener{errors: []string{}}
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	// Parse
	tree := p.Query()

	// Check for errors
	if len(errorListener.errors) > 0 {
		return createErrorResponse(errorListener.errors[0])
	}

	// Transpile
	t := transpiler.NewPromQLTranspiler()
	promql := t.Visit(tree).(string)

	return createSuccessResponse(promql)
}

func createSuccessResponse(result string) js.Value {
	obj := js.Global().Get("Object").New()
	obj.Set("success", true)
	obj.Set("result", result)
	return obj
}

func createErrorResponse(error string) js.Value {
	obj := js.Global().Get("Object").New()
	obj.Set("success", false)
	obj.Set("error", error)
	return obj
}

// Custom error listener for JavaScript
type jsErrorListener struct {
	*antlr.DefaultErrorListener
	errors []string
}

func (c *jsErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	err := fmt.Sprintf("Line %d:%d %s", line, column, msg)
	c.errors = append(c.errors, err)
}
