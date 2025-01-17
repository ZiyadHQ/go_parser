package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")

	lexer := Lexer{
		program:      "3 ** (2 * (3 / (4 + (5 - (6)))))",
		currentIndex: 0,
	}

	str := "Ahmad"
	str += " Khaled"
	str += " Saleh"

	fmt.Println(str)

	Scan(&lexer)
	printLexer(&lexer)

	parser := CreateParser(&lexer)

	node := parseExpression(parser)

	fmt.Printf("%+v\n", node)

	fmt.Printf("Result: %+v\n", node.Evaluate())

}
