package ast

import (
	"bytes"

	"github.com/Jihyun3478/logi-lang/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

// 문장 노드를 나타내는 인터페이스 (let, return 등)
type Statement interface {
	Node
	statementNode()
}

// 표현식 노드를 나타내는 인터페이스 (변수, 리터럴, 연산 등)
type Expression interface {
	Node
	expressionNode()
}

// AST의 루트 노드로 모든 문장들을 포함
type Program struct {
	Statements []Statement
}

// Program의 첫 번째 문장의 토큰 리터럴을 반환
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Program을 문자열로 변환하여 반환
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// let 문장을 나타내는 노드 (let x = 5;)
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// LetStatement가 Statement 인터페이스를 만족하도록 하는 마커 메서드
func (ls *LetStatement) statementNode() {}

// LetStatement의 토큰 리터럴을 반환
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// LetStatement를 문자열로 변환하여 반환
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// return 문장을 나타내는 노드
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// ReturnStatement가 Statement 인터페이스를 만족하도록 하는 마커 메서드
func (rs *ReturnStatement) statementNode() {}

// ReturnStatement의 토큰 리터럴을 반환
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// ReturnStatement를 문자열로 변환하여 반환
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// 표현식으로 이루어진 문장을 나타내는 노드
type ExpressionStatement struct {
	Token token.Token
	Expression Expression
}

// ExpressionStatement가 Statement 인터페이스를 만족하도록 하는 마커 메서드
func (es *ExpressionStatement) statementNode() {}

// ExpressionStatement의 토큰 리터럴을 반환
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// ExpressionStatement를 문자열로 변환하여 반환
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// 식별자(변수명)를 나타내는 노드
type Identifier struct {
	Token token.Token
	Value string
}

// Identifier가 Expression 인터페이스를 만족하도록 하는 마커 메서드
func (i *Identifier) expressionNode() {}

// Identifier의 토큰 리터럴을 반환
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// Identifier를 문자열로 변환하여 반환
func (i *Identifier) String() string {
	return i.Value
}
