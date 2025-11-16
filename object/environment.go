package object

// 새로운 환경 생성
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// 변수를 저장하고 관리하는 환경
type Environment struct {
	store map[string]Object
	outer *Environment
}

// 환경에서 변수 값을 조회 (외부 스코프까지 탐색)
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// 환경에 변수 값을 저장
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

// 외부 환경을 감싸는 새로운 환경을 생성 (함수 스코프용)
func NewEnclosedEnvirionment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
