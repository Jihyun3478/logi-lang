package ast

import (
	"github.com/Jihyun3478/logi-lang/token"
)

type Node interface {
	TokenLiteral() string
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

// let 문장을 나타내는 노드 (let x = 5;)
type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

// 식별자(변수명)를 나타내는 노드
type Identifier struct {
	Token token.Token
	Value string
}

// Program의 첫 번째 문장의 토큰 리터럴을 반환
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement가 Statement 인터페이스를 만족하도록 하는 마커 메서드
func (ls *LetStatement) statementNode() {}

// LetStatement의 토큰 리터럴을 반환
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier가 Expression 인터페이스를 만족하도록 하는 마커 메서드
func (i *Identifier) expressionNode() {}

// Identifier의 토큰 리터럴을 반환
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
