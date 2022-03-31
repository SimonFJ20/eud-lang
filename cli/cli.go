package cli

import "fmt"

func Cli() {
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		panic("err")
	}
	fmt.Printf("\"%s\" yourself, fuck you, not implemented\n", input)
}
