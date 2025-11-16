# Logi 언어 만들기 프로젝트
> 토르슈텐의 "밑바닥부터 만드는 인터프리터 in Go" 책을 참고하여, 직접 만드는 인터프리터 프로젝트입니다!

## 📋 프로젝트 개요
Logi는 Go 언어로 구현하는 인터프리터 프로젝트입니다. 렉서(Lexer), 파서(Parser), AST(Abstract Syntax Tree), 평가기(Evaluator)를 단계별로 구현하며 프로그래밍 언어가 어떻게 동작하는지 학습합니다.

### 주요 기능 (예정)
- ✅ 렉서: 소스 코드를 토큰으로 변환
- 🚧 파서: 토큰을 AST로 변환
- ⏳ 평가기: AST를 실행하여 결과 도출
- ⏳ 변수 선언 및 할당 (let)
- ⏳ 사칙연산 (+, -, *, /)
- ⏳ 조건문 (if-else)
- ⏳ 함수 정의 및 호출

### 기술 스택
- **언어**: Go 1.25.4
- **테스트**: Go testing package
- **참고 서적**: "Writing An Interpreter in Go" by Thorsten Ball

## 📅 진행 상황
- [x] 렉서
- [x] 파서
- [x] 평가기

## 🚀 실행 방법 (임시)
```go
go run main.go
```

## 📁 프로젝트 구조
```
logi-lang/
├── token/          # 토큰 타입 정의
│   └── token.go
├── lexer/          # 렉서 (토큰화)
│   ├── lexer.go
│   └── lexer_test.go
├── ast/            # AST 노드 정의
│   └── ast.go
├── parser/         # 파서 (구문 분석)
│   ├── parser.go
│   └── parser_test.go
├── evaluator/      # 평가기 (실행) - 진행 중
├── object/         # 내부 객체 - 예정
├── main.go         # CLI 엔트리포인트
├── go.mod
└── README.md
```

## 📚 학습 자료
- [밑바닥부터 만드는 인터프리터 in Go](https://product.kyobobook.co.kr/detail/S000001033117) - Thorsten Ball
- [Go 공식 문서](https://go.dev/doc/)

## 📌 참고 사항
- 이 프로젝트는 학습 목적으로 제작되었습니다.
- 개발 기간: 2025.11.10 ~ 2025.11.16
- 매일 커밋 및 진행 상황 기록

## 🔗 링크
- GitHub: [Jihyun3478/logi-lang](https://github.com/Jihyun3478/logi-lang)
- 블로그: [개발 과정 정리] - 예정

---

**Last Updated:** 2025-11-16
**Status:** 🚧 개발 완료
