package object

import "fmt"

// 객체의 타입을 나타내는 문자열
type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ = "ERROR"
)

// 모든 값 객체가 구현해야 하는 인터페이스
type Object interface {
	Type() ObjectType
	Inspect() string
}

// 정수 값을 나타내는 객체
type Integer struct {
	Value int64
}

// Integer 객체의 타입을 반환
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// Integer 객체를 문자열로 변환하여 반환
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// 불리언 값을 나타내는 객체
type Boolean struct {
	Value bool
}

// Boolean 객체의 타입을 반환
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

// Boolean 객체를 문자열로 변환하여 반환
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// null 값을 나타내는 객체
type Null struct{}

// Null 객체의 타입을 반환
func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

// Null 객체를 문자열로 변환하여 반환
func (n *Null) Inspect() string {
	return "null"
}

// return 문의 반환값을 감싸는 객체
type ReturnValue struct {
	Value Object
}

// ReturnValue 객체의 타입을 반환
func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

// ReturnValue 객체를 문자열로 변환하여 반환
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// 에러를 나타내는 객체
type Error struct {
	Message string
}

// Error 객체의 타입을 반환
func (e *Error) Type() ObjectType { return ERROR_OBJ }

// Error 객체를 문자열로 변환하여 반환
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}
