package evaluator

import (
	"github.com/Jihyun3478/logi-lang/ast"
	"github.com/Jihyun3478/logi-lang/object"
)

// AST 노드를 평가하여 객체로 변환
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	
	// 명령문
	case *ast.Program:
		return evalStatements(node.Statements)
	
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	
	// 표현식
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	return nil
}

// 문장 목록을 순차적으로 평가
func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}
