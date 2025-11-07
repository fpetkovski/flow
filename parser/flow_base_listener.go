// Code generated from Flow.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Flow
import "github.com/antlr4-go/antlr/v4"

// BaseFlowListener is a complete listener for a parse tree produced by FlowParser.
type BaseFlowListener struct{}

var _ FlowListener = &BaseFlowListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFlowListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFlowListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFlowListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFlowListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseFlowListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseFlowListener) ExitQuery(ctx *QueryContext) {}

// EnterLetExpression is called when production letExpression is entered.
func (s *BaseFlowListener) EnterLetExpression(ctx *LetExpressionContext) {}

// ExitLetExpression is called when production letExpression is exited.
func (s *BaseFlowListener) ExitLetExpression(ctx *LetExpressionContext) {}

// EnterLetBinding is called when production letBinding is entered.
func (s *BaseFlowListener) EnterLetBinding(ctx *LetBindingContext) {}

// ExitLetBinding is called when production letBinding is exited.
func (s *BaseFlowListener) ExitLetBinding(ctx *LetBindingContext) {}

// EnterPipeline is called when production pipeline is entered.
func (s *BaseFlowListener) EnterPipeline(ctx *PipelineContext) {}

// ExitPipeline is called when production pipeline is exited.
func (s *BaseFlowListener) ExitPipeline(ctx *PipelineContext) {}

// EnterPipelineStep is called when production pipelineStep is entered.
func (s *BaseFlowListener) EnterPipelineStep(ctx *PipelineStepContext) {}

// ExitPipelineStep is called when production pipelineStep is exited.
func (s *BaseFlowListener) ExitPipelineStep(ctx *PipelineStepContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseFlowListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseFlowListener) ExitSelector(ctx *SelectorContext) {}

// EnterLabelMatcher is called when production labelMatcher is entered.
func (s *BaseFlowListener) EnterLabelMatcher(ctx *LabelMatcherContext) {}

// ExitLabelMatcher is called when production labelMatcher is exited.
func (s *BaseFlowListener) ExitLabelMatcher(ctx *LabelMatcherContext) {}

// EnterBinarySelector is called when production binarySelector is entered.
func (s *BaseFlowListener) EnterBinarySelector(ctx *BinarySelectorContext) {}

// ExitBinarySelector is called when production binarySelector is exited.
func (s *BaseFlowListener) ExitBinarySelector(ctx *BinarySelectorContext) {}

// EnterBinarySelectorLeg is called when production binarySelectorLeg is entered.
func (s *BaseFlowListener) EnterBinarySelectorLeg(ctx *BinarySelectorLegContext) {}

// ExitBinarySelectorLeg is called when production binarySelectorLeg is exited.
func (s *BaseFlowListener) ExitBinarySelectorLeg(ctx *BinarySelectorLegContext) {}

// EnterBinaryLegLeaf is called when production binaryLegLeaf is entered.
func (s *BaseFlowListener) EnterBinaryLegLeaf(ctx *BinaryLegLeafContext) {}

// ExitBinaryLegLeaf is called when production binaryLegLeaf is exited.
func (s *BaseFlowListener) ExitBinaryLegLeaf(ctx *BinaryLegLeafContext) {}

// EnterBinaryOperator is called when production binaryOperator is entered.
func (s *BaseFlowListener) EnterBinaryOperator(ctx *BinaryOperatorContext) {}

// ExitBinaryOperator is called when production binaryOperator is exited.
func (s *BaseFlowListener) ExitBinaryOperator(ctx *BinaryOperatorContext) {}

// EnterAggregation is called when production aggregation is entered.
func (s *BaseFlowListener) EnterAggregation(ctx *AggregationContext) {}

// ExitAggregation is called when production aggregation is exited.
func (s *BaseFlowListener) ExitAggregation(ctx *AggregationContext) {}

// EnterAligner is called when production aligner is entered.
func (s *BaseFlowListener) EnterAligner(ctx *AlignerContext) {}

// ExitAligner is called when production aligner is exited.
func (s *BaseFlowListener) ExitAligner(ctx *AlignerContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseFlowListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseFlowListener) ExitFunction(ctx *FunctionContext) {}

// EnterBinary is called when production binary is entered.
func (s *BaseFlowListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production binary is exited.
func (s *BaseFlowListener) ExitBinary(ctx *BinaryContext) {}
