// Code generated from antlr/Yars.g4 by ANTLR 4.10.1. DO NOT EDIT.

package antlr // Yars
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseYarsListener is a complete listener for a parse tree produced by YarsParser.
type BaseYarsListener struct{}

var _ YarsListener = &BaseYarsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseYarsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseYarsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseYarsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseYarsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProject is called when production project is entered.
func (s *BaseYarsListener) EnterProject(ctx *ProjectContext) {}

// ExitProject is called when production project is exited.
func (s *BaseYarsListener) ExitProject(ctx *ProjectContext) {}

// EnterActionStatement is called when production actionStatement is entered.
func (s *BaseYarsListener) EnterActionStatement(ctx *ActionStatementContext) {}

// ExitActionStatement is called when production actionStatement is exited.
func (s *BaseYarsListener) ExitActionStatement(ctx *ActionStatementContext) {}

// EnterRecommendStatement is called when production recommendStatement is entered.
func (s *BaseYarsListener) EnterRecommendStatement(ctx *RecommendStatementContext) {}

// ExitRecommendStatement is called when production recommendStatement is exited.
func (s *BaseYarsListener) ExitRecommendStatement(ctx *RecommendStatementContext) {}

// EnterParamsDecl is called when production paramsDecl is entered.
func (s *BaseYarsListener) EnterParamsDecl(ctx *ParamsDeclContext) {}

// ExitParamsDecl is called when production paramsDecl is exited.
func (s *BaseYarsListener) ExitParamsDecl(ctx *ParamsDeclContext) {}

// EnterParamDecl is called when production paramDecl is entered.
func (s *BaseYarsListener) EnterParamDecl(ctx *ParamDeclContext) {}

// ExitParamDecl is called when production paramDecl is exited.
func (s *BaseYarsListener) ExitParamDecl(ctx *ParamDeclContext) {}

// EnterBinaryOp is called when production binaryOp is entered.
func (s *BaseYarsListener) EnterBinaryOp(ctx *BinaryOpContext) {}

// ExitBinaryOp is called when production binaryOp is exited.
func (s *BaseYarsListener) ExitBinaryOp(ctx *BinaryOpContext) {}

// EnterLvalueExpression is called when production lvalueExpression is entered.
func (s *BaseYarsListener) EnterLvalueExpression(ctx *LvalueExpressionContext) {}

// ExitLvalueExpression is called when production lvalueExpression is exited.
func (s *BaseYarsListener) ExitLvalueExpression(ctx *LvalueExpressionContext) {}

// EnterUnaryOp is called when production unaryOp is entered.
func (s *BaseYarsListener) EnterUnaryOp(ctx *UnaryOpContext) {}

// ExitUnaryOp is called when production unaryOp is exited.
func (s *BaseYarsListener) ExitUnaryOp(ctx *UnaryOpContext) {}

// EnterBrackets is called when production brackets is entered.
func (s *BaseYarsListener) EnterBrackets(ctx *BracketsContext) {}

// ExitBrackets is called when production brackets is exited.
func (s *BaseYarsListener) ExitBrackets(ctx *BracketsContext) {}

// EnterLvalue is called when production lvalue is entered.
func (s *BaseYarsListener) EnterLvalue(ctx *LvalueContext) {}

// ExitLvalue is called when production lvalue is exited.
func (s *BaseYarsListener) ExitLvalue(ctx *LvalueContext) {}

// EnterBinaryOps is called when production binaryOps is entered.
func (s *BaseYarsListener) EnterBinaryOps(ctx *BinaryOpsContext) {}

// ExitBinaryOps is called when production binaryOps is exited.
func (s *BaseYarsListener) ExitBinaryOps(ctx *BinaryOpsContext) {}

// EnterUnaryOps is called when production unaryOps is entered.
func (s *BaseYarsListener) EnterUnaryOps(ctx *UnaryOpsContext) {}

// ExitUnaryOps is called when production unaryOps is exited.
func (s *BaseYarsListener) ExitUnaryOps(ctx *UnaryOpsContext) {}

// EnterUnitStatement is called when production unitStatement is entered.
func (s *BaseYarsListener) EnterUnitStatement(ctx *UnitStatementContext) {}

// ExitUnitStatement is called when production unitStatement is exited.
func (s *BaseYarsListener) ExitUnitStatement(ctx *UnitStatementContext) {}

// EnterFieldsDecl is called when production fieldsDecl is entered.
func (s *BaseYarsListener) EnterFieldsDecl(ctx *FieldsDeclContext) {}

// ExitFieldsDecl is called when production fieldsDecl is exited.
func (s *BaseYarsListener) ExitFieldsDecl(ctx *FieldsDeclContext) {}

// EnterFieldDecl is called when production fieldDecl is entered.
func (s *BaseYarsListener) EnterFieldDecl(ctx *FieldDeclContext) {}

// ExitFieldDecl is called when production fieldDecl is exited.
func (s *BaseYarsListener) ExitFieldDecl(ctx *FieldDeclContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseYarsListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseYarsListener) ExitType_(ctx *Type_Context) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseYarsListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseYarsListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BaseYarsListener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BaseYarsListener) ExitBaseType(ctx *BaseTypeContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseYarsListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseYarsListener) ExitEos(ctx *EosContext) {}
