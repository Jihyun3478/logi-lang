package evaluator

import (
	"github.com/Jihyun3478/logi-lang/ast"
	"github.com/Jihyun3478/logi-lang/object"
)

var (
	NULL = &object.Null{}
	TRUE = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
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

	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)

	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
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

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

// 전위 연산자 표현식을 평가 (!, -)
func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return NULL
	}
}

// 논리 부정 연산자(!)를 평가
func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

// 음수 연산자(-)를 평가
func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return NULL
	}
	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}
