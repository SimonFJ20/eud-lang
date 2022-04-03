package main

import (
	"eud-lang/bytecode"
	"eud-lang/parser"
	"fmt"
)

func main() {

	text := "(3 + 4) * 5"

	fmt.Printf("\033[1;36mInput:\033[0m\n\"%s\"\n", text)

	println("\033[1;36mTokenizing text:\033[0m")

	tokens := parser.TokenizeString(text)

	current_token := tokens
	for current_token != nil {
		println(current_token.Type.String())
		current_token = current_token.Next
	}

	println("\033[1;36mParsing tokens:\033[0m")

	p := parser.Parser{}

	ast := p.Parse(tokens)

	fmt.Printf("%s\n", ast)

	println("\033[1;36mCompiling AST:\033[0m")

	program, _ := bytecode.Compile(ast)

	for i := range program.Instructions {
		fmt.Println(program.Instructions[i].String())
	}

	println("\033[1;36mRunning bytecode:\033[0m")

	runtime := bytecode.Run(program)

	fmt.Printf("\033[1;36mResult:\033[0m\n%s = %d\n", text, runtime.Stack[0].(bytecode.I32Value).Value)

	// cli.Cli(
	// panic("not implemented i guess")
}
