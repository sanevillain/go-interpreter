package repl

import (
	"bufio"
	"fmt"
	"io"
	"sanevillain/go-interpreter/lexer"
	"sanevillain/go-interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			if line == "exit" || line == "quit" {
				return
			}

			output := fmt.Sprintf("%+v\n", tok)

			n, err := out.Write([]byte(output))
			if n != len(output) {
				return
			}

			if err != nil {
				return
			}
		}
	}
}
