mod parser;
use crate::parser::{ast, tokenizer::tokenize_string};

fn main() {
    let token = tokenize_string(String::from("2+3+4+5"));

    println!("Parsing tokens:");

    let mut parser = ast::Parser::from(token);

    let ast = parser.parse();

    println!("{}", ast);
    /*

    println("\033[1;36mCompiling AST:\033[0m")

    program, _ := bytecode.Compile([]parser.BaseStatement{ast})

    for i := range program.Instructions {
        fmt.Println(program.Instructions[i].String())
    }

    println("\033[1;36mRunning bytecode:\033[0m")

    program.RunWithDebug = true
    runtime := bytecode.Run(program)

    fmt.Printf("\033[1;36mResult:\033[0m\n%s = %d\n", text, runtime.Stack[0].(bytecode.I32Value).Value)

    */
}
