package main

import (
	"fmt"
	"log"
)

type TokenType int

const (
	Plus = iota
	Minus
	Slash
	Star
	StarStar
	LeftParen
	RightParen
	LiteralString
	Pipe
)

var TokenType_str = map[TokenType]string{
	Plus:          "Plus",
	Minus:         "Minus",
	Slash:         "Slash",
	Star:          "Star",
	StarStar:      "StarStar",
	LeftParen:     "LeftParen",
	RightParen:    "RightParen",
	LiteralString: "LiteralString",
	Pipe:          "Pipe",
}

type Token struct {
	content   string
	tokenType TokenType
}

type Lexer struct {
	program      string
	currentIndex int
	tokens       []Token
}

func isAtEnd(lexer *Lexer) bool {
	return lexer.currentIndex >= len(lexer.program)
}

// Despite the name, it checks for both whitespace AND reserved characters
func isWhiteSpace(lexer *Lexer) bool {
	if isAtEnd(lexer) {
		return true
	}
	char := lexer.program[lexer.currentIndex]
	return (char == ' ' || char == '\n' || char == '\t' || char == '\r' || char == '\v' || char == '\f' || char == '+' || char == '-' || char == '/' || char == '*' || char == '(' || char == ')' || char == '|')
}

func Next(lexer *Lexer) byte {
	char := lexer.program[lexer.currentIndex]
	lexer.currentIndex++
	return char
}

func Peek(lexer *Lexer) byte {
	if lexer.currentIndex < len(lexer.program) {
		return lexer.program[lexer.currentIndex]
	} else {
		return 0
	}
}

func Previous(lexer *Lexer) byte {
	if lexer.currentIndex >= 0 {
		return lexer.program[lexer.currentIndex]
	}
	log.Fatalf("Error returning Previous character, index less than 0, index: %d", lexer.currentIndex)
	return 0
}

func addToken(lexer *Lexer, token Token) {
	lexer.tokens = append(lexer.tokens, token)
}

func Scan(lexer *Lexer) {

	for !isAtEnd(lexer) {

		currentChar := Next(lexer)
		switch currentChar {
		case '+':
			addToken(lexer, Token{content: string(currentChar), tokenType: Plus})
		case '-':
			addToken(lexer, Token{content: string(currentChar), tokenType: Minus})
		case '/':
			addToken(lexer, Token{content: string(currentChar), tokenType: Slash})
		case '(':
			addToken(lexer, Token{content: string(currentChar), tokenType: LeftParen})
		case ')':
			addToken(lexer, Token{content: string(currentChar), tokenType: RightParen})
		case '|':
			addToken(lexer, Token{content: string(currentChar), tokenType: Pipe})
		case '*':
			if Peek(lexer) == '*' {
				addToken(lexer, Token{content: string(currentChar) + string(Next(lexer)), tokenType: StarStar})
			} else {
				addToken(lexer, Token{content: string(currentChar), tokenType: Star})
			}

		case ' ', '\n', '\t', 'r', '\v', '\f':
		default:
			literal := string(currentChar)
			for !isWhiteSpace(lexer) {
				literal += string(Next(lexer))
			}
			addToken(lexer, Token{content: literal, tokenType: LiteralString})
		}
	}
}

func printLexer(lexer *Lexer) {
	for i := 0; i < len(lexer.tokens); i++ {
		fmt.Printf("%s, ", TokenToString(lexer.tokens[i]))
	}
	fmt.Println()
}

func TokenToString(token Token) string {
	return fmt.Sprintf("{%q, %s}", token.content, TokenType_str[token.tokenType])
}
