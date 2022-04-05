package cli

import (
	"errors"
	"eud/bytecode"
	"eud/parser"
	"fmt"
	"io/ioutil"
	"os"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else if err != nil {
		panic("error checking file existance")
	}
	return true
}

func Cli() {
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("no files given")
		os.Exit(1)
	}

	for i := 0; i < len(files); i++ {
		if !fileExists(files[i]) {
			fmt.Printf("file %q does not exist\n", files[i])
			os.Exit(1)
		}
	}

	for i := 0; i < len(files); i++ {
		file_bytes, err := ioutil.ReadFile(files[i])

		if err != nil {
			panic(err)
		}

		text := string(file_bytes)

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

		program, _ := bytecode.Compile([]parser.BaseStatement{ast})

		for i := range program.Instructions {
			fmt.Println(program.Instructions[i].String())
		}

		println("\033[1;36mRunning bytecode:\033[0m")

		program.RunWithDebug = true
		runtime := bytecode.Run(program)

		fmt.Printf("\033[1;36mResult:\033[0m\n%s = %d\n", text, runtime.Stack[0].(bytecode.I32Value).Value)
	}
}
