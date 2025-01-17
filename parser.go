package main

import "log"

type ASTNode interface {
	Evaluate() float64
}

type Binary struct {
	left     ASTNode
	right    ASTNode
	operator Token
}

type Unary struct {
	expr     ASTNode
	operator Token
}

type Absolute struct {
	expr ASTNode
}

type Grouping struct {
	expr ASTNode
}

type Literal struct {
	value Token
}

type Parser struct {
	tokens       []Token
	currentIndex int
}

func Consume(parser *Parser, tokenType TokenType) Token {
	token := Peek_parser(parser)
	if token.tokenType == tokenType {
		parser.currentIndex++
		return token
	} else {
		log.Fatalf("Error, couldn't consume token of type: %q", TokenType_str[tokenType])
		return Token{content: "", tokenType: tokenType}
	}
}

func ConsumeGroup(parser *Parser, types ...TokenType) Token {
	for i := 0; i < len(types); i++ {
		if types[i] == Peek_parser(parser).tokenType {
			token := Peek_parser(parser)
			parser.currentIndex++
			return token
		}
	}
	log.Fatalf("Error consuming token, expected types: %+v", types)
	return Token{}
}

func isAtEnd_parser(parser *Parser) bool {
	return parser.currentIndex >= len(parser.tokens)
}

func Peek_parser(parser *Parser) Token {
	if !isAtEnd_parser(parser) {
		return parser.tokens[parser.currentIndex]
	}
	// log.Fatalf("Error, couldn't peek since parser reached end of tokens!")
	return Token{content: "", tokenType: -1}
}

func matchAhead(parser *Parser, chars ...TokenType) bool {
	if parser.currentIndex < len(parser.tokens) {
		for i := 0; i < len(chars); i++ {
			if chars[i] == Peek_parser(parser).tokenType {
				return true
			}
		}
		return false
	} else {
		return false
	}
}

func Previous_parser(parser *Parser) Token {
	return parser.tokens[parser.currentIndex-1]
}

func match(parser *Parser, chars ...TokenType) bool {
	for i := 0; i < len(chars); i++ {
		if chars[i] == Peek_parser(parser).tokenType {
			parser.currentIndex++
			return true
		}
	}
	return false
}

func CreateParser(lexer *Lexer) *Parser {
	parser := Parser{
		tokens:       lexer.tokens,
		currentIndex: 0,
	}

	return &parser
}

func parseExpression(parser *Parser) ASTNode {
	return parseAddition(parser)
}

func parseAddition(parser *Parser) ASTNode {
	left := parseMultiplication(parser)

	for match(parser, Plus, Minus) {
		operator := Previous_parser(parser)
		right := parseMultiplication(parser)
		left = &Binary{
			left:     left,
			right:    right,
			operator: operator,
		}
	}

	return left
}

func parseMultiplication(parser *Parser) ASTNode {

	left := parsePower(parser)

	for match(parser, Slash, Star) {
		operator := Previous_parser(parser)
		right := parsePower(parser)
		left = &Binary{
			left:     left,
			right:    right,
			operator: operator,
		}
	}

	return left
}

func parsePower(parser *Parser) ASTNode {
	left := parseUnary(parser)

	for match(parser, StarStar) {
		operator := Previous_parser(parser)
		right := parseUnary(parser)
		left = &Binary{
			left:     left,
			right:    right,
			operator: operator,
		}
	}

	return left
}

func parseUnary(parser *Parser) ASTNode {
	if match(parser, Minus) {
		operator := Previous_parser(parser)
		operand := parseUnary(parser)
		return &Unary{
			expr:     operand,
			operator: operator,
		}
	}

	return parseAbsolute(parser)
}

func parseAbsolute(parser *Parser) ASTNode {
	var expr ASTNode

	if match(parser, Pipe) {
		expr = &Absolute{
			expr: parseExpression(parser),
		}
		Consume(parser, Pipe)
	} else {
		expr = parseGrouping(parser)
	}

	return expr
}

func parseGrouping(parser *Parser) ASTNode {
	var expr ASTNode

	if match(parser, LeftParen) {
		expr = &Grouping{
			expr: parseExpression(parser),
		}
		Consume(parser, RightParen)
	} else {
		expr = parseLiteral(parser)
	}

	return expr
}

func parseLiteral(parser *Parser) ASTNode {
	expr := Literal{
		value: Consume(parser, LiteralString),
	}

	return &expr
}
