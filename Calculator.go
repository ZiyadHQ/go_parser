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
	case "%":
		return math.Mod(n.left.Evaluate(), n.right.Evaluate())
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

func (n *Absolute) Evaluate() float64 {
	return math.Abs(n.expr.Evaluate())
}

func (n *Grouping) Evaluate() float64 {
	return n.expr.Evaluate()
}

func (n *FunctionNode) Evaluate() float64 {
	switch n.function.content {
	case "Floor":
		return math.Floor(n.expr.Evaluate())
	case "Ceil":
		return math.Ceil(n.expr.Evaluate())
	case "Sin":
		return math.Sin(n.expr.Evaluate())
	case "Cos":
		return math.Cos(n.expr.Evaluate())
	case "Tan":
		return math.Tan(n.expr.Evaluate())
	case "Sqrt":
		return math.Sqrt(n.expr.Evaluate())
	}
	log.Fatalf("Error, unrecognized function: %q", n.function.content)
	return -0
}

func (n *Literal) Evaluate() float64 {

	value, err := strconv.ParseFloat(n.value.content, 64)

	if err != nil {
		log.Fatalf("Error evaluating literal expr: %")
	}

	return value
}
