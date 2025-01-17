package main

import (
	"log"
	"math"
	"strconv"
)

func (n *Binary) Evaluate() float64 {
	switch n.operator.content {
	case "+":
		return n.left.Evaluate() + n.right.Evaluate()
	case "-":
		return n.left.Evaluate() - n.right.Evaluate()
	case "/":
		return n.left.Evaluate() / n.right.Evaluate()
	case "*":
		return n.left.Evaluate() * n.right.Evaluate()
	case "**":
		return math.Pow(n.left.Evaluate(), n.right.Evaluate())
	default:
		log.Fatalf("Error evaluating binary expr: %")
	}
	return -999
}

func (n *Unary) Evaluate() float64 {
	switch n.operator.content {
	case "-":
		return -n.expr.Evaluate()
	default:
		log.Fatalf("Error evaluating unary expr: %")
	}
	return -999
}

func (n *Grouping) Evaluate() float64 {
	return n.expr.Evaluate()
}

func (n *Literal) Evaluate() float64 {

	value, err := strconv.ParseFloat(n.value.content, 64)

	if err != nil {
		log.Fatalf("Error evaluating literal expr: %")
	}

	return value
}
