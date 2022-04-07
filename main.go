package main

import (
	"errors"
	"eud/astjson"
	"eud/bytecode"
	"eud/parser"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
)

type Options struct {
	UsePythonParser bool
}

func main() {
	file := getFileFromArgs()
	options := getOptionsFromArgs()

	file_bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	text := string(file_bytes)

	fmt.Printf("\033[1;36mInput:\033[0m\n%s\n\n", text)

	var ast []parser.BaseStatement
	if options.UsePythonParser {
		ast = parseUsingPythonParser(text, file)
	} else {
		ast = parseUsingGoParser(text)
	}

	// fmt.Printf("%s\n", ast)
	for i := range ast {
		fmt.Printf("%s\n", ast[i].StringNested(1))
	}

	println("\033[1;36mCompiling AST:\033[0m")

	program, err := bytecode.Compile(ast)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range program.Instructions {
		fmt.Printf("  %s\n", program.Instructions[i].String())
	}

	println("\033[1;36mRunning bytecode:\033[0m")

	program.RunWithDebug = true
	runtime := bytecode.Run(program)

	last_useful_index := findLastUsefulIndex(runtime)

	// fmt.Printf("\033[1;36mResult:\033[0m\n%s\n\nStack: %s\n", text, runtime.Stack[:last_useful_index])

	locals_str := "["
	locals_str_first := true
	for i := range runtime.Locals {
		if !locals_str_first {
			locals_str += ", "
		} else {
			locals_str_first = false
		}
		locals_str += fmt.Sprintf("%d: %s", i, runtime.Locals[i].String())
	}
	locals_str += "]"

	fmt.Printf("\033[1;36mResult:\033[0m\n  Stack: %s\n  Locals: %s\n", runtime.Stack[:last_useful_index], locals_str)
}

func getFileFromArgs() string {
	if len(os.Args) == 1 {
		fmt.Println("no files given")
		os.Exit(1)
	}

	file := os.Args[1]

	if !fileExists(file) {
		fmt.Printf("file %q does not exist\n", file)
		os.Exit(1)
	}
	return file
}

func getOptionsFromArgs() Options {
	args := os.Args[2:]
	options := Options{
		UsePythonParser: false,
	}
	for i := range args {
		switch args[i] {
		case "--pyparse":
			options.UsePythonParser = true
		}
	}
	return options
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else if err != nil {
		log.Fatal("error checking file existance")
	}
	return true
}

func parseUsingGoParser(text string) []parser.BaseStatement {
	println("\033[1;36mTokenizing text:\033[0m")

	tokens := parser.TokenizeString(text)

	println("\033[1;36mParsing tokens:\033[0m")

	p := parser.Parser{}

	ast := p.Parse(tokens)
	return ast
}

func parseUsingPythonParser(text string, filepath string) []parser.BaseStatement {
	python_comand := "python3"
	if runtime.GOOS == "windows" {
		python_comand = "py"
	}
	fmt.Printf("\033[1;36mParsing text to AST with `%s parser.py %s -ofile`:\033[0m\n", python_comand, filepath)
	cmd := exec.Command(python_comand, "parser.py", filepath, "-ofile")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("parser.py: %s\n", err)
		os.Exit(1)
	}
	asttempjsonbytes, err := ioutil.ReadFile("ast.temp.json")
	astjsonstring := string(asttempjsonbytes)
	if err != nil {
		log.Fatal(err)
	}
	println("\033[1;36mParsing AST json to internal AST:\033[0m")
	ast := astjson.Parse(astjsonstring)
	return ast
}

func findLastUsefulIndex(runtime bytecode.Runtime) int {
	last_useful_index := 0
	for i := len(runtime.Stack) - 1; i >= 0; i-- {
		if runtime.Stack[i] != nil {
			last_useful_index = i + 1
			break
		}
	}
	return last_useful_index
}
