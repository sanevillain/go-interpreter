package main

import (
	"os"
	"sanevillain/go-interpreter/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
