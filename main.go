package main

import (
	"eud-lang/parser"
	"fmt"
)

func main() {

	text := "3 + 4 * 5"

	tokens := parser.TokenizeString(text)

	p := parser.Parser{}

	ast := p.Parse(tokens)

	fmt.Printf("%s\b", ast)

	// program, _ := bytecode.Compile(ast)

	// for i := range program.Instructions {
	// 	fmt.Printf(program.Instructions[i].String())
	// }

	// cli.Cli(
	// panic("not implemented i guess")
}
