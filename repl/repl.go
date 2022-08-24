package repl

import (
	"bufio"
	"dumpy/lexer"
	"dumpy/parser"
	"fmt"
	"io"
)

// TODO: Replace with ascii dumpster art
const DUMPSTER = `
$$$$$$$\                                              
$$  __$$\                                             
$$ |  $$ |$$\   $$\ $$$$$$\$$$$\   $$$$$$\  $$\   $$\ 
$$ |  $$ |$$ |  $$ |$$  _$$  _$$\ $$  __$$\ $$ |  $$ |
$$ |  $$ |$$ |  $$ |$$ / $$ / $$ |$$ /  $$ |$$ |  $$ |
$$ |  $$ |$$ |  $$ |$$ | $$ | $$ |$$ |  $$ |$$ |  $$ |
$$$$$$$  |\$$$$$$  |$$ | $$ | $$ |$$$$$$$  |\$$$$$$$ |
\_______/  \______/ \__| \__| \__|$$  ____/  \____$$ |
                                  $$ |      $$\   $$ |
                                  $$ |      \$$$$$$  |
                                  \__|       \______/ 
`
const PROMPT = ">> "

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

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")

		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	fmt.Fprintf(out, "%+v\n", tok)
		// }
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, DUMPSTER)
	io.WriteString(out, "Woops! We ran into some garbage here!\n")
	io.WriteString(out, " parser errors:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
