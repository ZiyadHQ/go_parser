package main

import (
	"fmt"
	"os"
)

func main() {

	lexer := Lexer{
		program:      os.Args[1],
		currentIndex: 0,
	}

	Scan(&lexer)
	printLexer(&lexer)

	parser := CreateParser(&lexer)

	node := parseExpression(parser)

	fmt.Printf("Result: %+v\n", node.Evaluate())

}
