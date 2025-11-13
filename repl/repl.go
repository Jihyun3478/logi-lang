package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Jihyun3478/logi-lang/evaluator"
	"github.com/Jihyun3478/logi-lang/lexer"
	"github.com/Jihyun3478/logi-lang/parser"
)

const PROMPT = ">> "
const LOGI_LOGO = `
 ██╗      ██████╗  ██████╗ ██╗
 ██║     ██╔═══██╗██╔════╝ ██║
 ██║     ██║   ██║██║  ███╗██║
 ██║     ██║   ██║██║   ██║██║
 ███████╗╚██████╔╝╚██████╔╝██║
 ╚══════╝ ╚═════╝  ╚═════╝ ╚═╝
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, program.String())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, LOGI_LOGO)
	io.WriteString(out, "Woops! We Ran into some business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t" + msg + "\n")
	}
}
