// Code generated from antlr/Yars.g4 by ANTLR 4.10.1. DO NOT EDIT.

package antlr // Yars
import "github.com/antlr/antlr4/runtime/Go/antlr"

// YarsListener is a complete listener for a parse tree produced by YarsParser.
type YarsListener interface {
	antlr.ParseTreeListener

	// EnterProject is called when entering the project production.
	EnterProject(c *ProjectContext)

	// EnterActionStatement is called when entering the actionStatement production.
	EnterActionStatement(c *ActionStatementContext)

	// EnterRecommendStatement is called when entering the recommendStatement production.
	EnterRecommendStatement(c *RecommendStatementContext)

	// EnterParamsDecl is called when entering the paramsDecl production.
	EnterParamsDecl(c *ParamsDeclContext)

	// EnterParamDecl is called when entering the paramDecl production.
	EnterParamDecl(c *ParamDeclContext)

	// EnterBinaryOp is called when entering the binaryOp production.
	EnterBinaryOp(c *BinaryOpContext)

	// EnterLvalueExpression is called when entering the lvalueExpression production.
	EnterLvalueExpression(c *LvalueExpressionContext)

	// EnterUnaryOp is called when entering the unaryOp production.
	EnterUnaryOp(c *UnaryOpContext)

	// EnterBrackets is called when entering the brackets production.
	EnterBrackets(c *BracketsContext)

	// EnterLvalue is called when entering the lvalue production.
	EnterLvalue(c *LvalueContext)

	// EnterBinaryOps is called when entering the binaryOps production.
	EnterBinaryOps(c *BinaryOpsContext)

	// EnterUnaryOps is called when entering the unaryOps production.
	EnterUnaryOps(c *UnaryOpsContext)

	// EnterUnitStatement is called when entering the unitStatement production.
	EnterUnitStatement(c *UnitStatementContext)

	// EnterFieldsDecl is called when entering the fieldsDecl production.
	EnterFieldsDecl(c *FieldsDeclContext)

	// EnterFieldDecl is called when entering the fieldDecl production.
	EnterFieldDecl(c *FieldDeclContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterBaseType is called when entering the baseType production.
	EnterBaseType(c *BaseTypeContext)

	// EnterEos is called when entering the eos production.
	EnterEos(c *EosContext)

	// ExitProject is called when exiting the project production.
	ExitProject(c *ProjectContext)

	// ExitActionStatement is called when exiting the actionStatement production.
	ExitActionStatement(c *ActionStatementContext)

	// ExitRecommendStatement is called when exiting the recommendStatement production.
	ExitRecommendStatement(c *RecommendStatementContext)

	// ExitParamsDecl is called when exiting the paramsDecl production.
	ExitParamsDecl(c *ParamsDeclContext)

	// ExitParamDecl is called when exiting the paramDecl production.
	ExitParamDecl(c *ParamDeclContext)

	// ExitBinaryOp is called when exiting the binaryOp production.
	ExitBinaryOp(c *BinaryOpContext)

	// ExitLvalueExpression is called when exiting the lvalueExpression production.
	ExitLvalueExpression(c *LvalueExpressionContext)

	// ExitUnaryOp is called when exiting the unaryOp production.
	ExitUnaryOp(c *UnaryOpContext)

	// ExitBrackets is called when exiting the brackets production.
	ExitBrackets(c *BracketsContext)

	// ExitLvalue is called when exiting the lvalue production.
	ExitLvalue(c *LvalueContext)

	// ExitBinaryOps is called when exiting the binaryOps production.
	ExitBinaryOps(c *BinaryOpsContext)

	// ExitUnaryOps is called when exiting the unaryOps production.
	ExitUnaryOps(c *UnaryOpsContext)

	// ExitUnitStatement is called when exiting the unitStatement production.
	ExitUnitStatement(c *UnitStatementContext)

	// ExitFieldsDecl is called when exiting the fieldsDecl production.
	ExitFieldsDecl(c *FieldsDeclContext)

	// ExitFieldDecl is called when exiting the fieldDecl production.
	ExitFieldDecl(c *FieldDeclContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitBaseType is called when exiting the baseType production.
	ExitBaseType(c *BaseTypeContext)

	// ExitEos is called when exiting the eos production.
	ExitEos(c *EosContext)
}
