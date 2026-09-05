package main

import (
	"fmt"
	"math/rand"
)

const (
	AND = iota
	OR
	XOR
	NOT
	OPERAND
)

const maxValue = 0b111111

type Config struct {
	Bits   int
	AlphaG int
	Trials int
}

type Level struct {
	Number int
	Bits   int
	Gates  int
}

type Expr struct {
	Type  int
	Value uint8
	Left  *Expr
	Right *Expr
}

func GenerateLevel(n int, config Config) Level {
	return Level{
		Number: n,
		Bits:   config.Bits,
		Gates:  1 + (n-1)/config.AlphaG,
	}
}

func GenerateExpression(gates int, rng *rand.Rand) *Expr {
	if gates == 0 {
		return &Expr{
			Type:  OPERAND,
			Value: uint8(rng.Intn(maxValue + 1)),
		}
	}

	gate := rng.Intn(4)

	if gate == NOT {
		return &Expr{
			Type: gate,
			Left: GenerateExpression(gates-1, rng),
		}
	}

	remaining := gates - 1
	leftGates := rng.Intn(remaining + 1)
	rightGates := remaining - leftGates

	return &Expr{
		Type:  gate,
		Left:  GenerateExpression(leftGates, rng),
		Right: GenerateExpression(rightGates, rng),
	}
}

func Evaluate(expr *Expr) uint8 {
	switch expr.Type {
	case OPERAND:
		return expr.Value
	case AND:
		return Evaluate(expr.Left) & Evaluate(expr.Right)
	case OR:
		return Evaluate(expr.Left) | Evaluate(expr.Right)
	case XOR:
		return Evaluate(expr.Left) ^ Evaluate(expr.Right)
	case NOT:
		return ^Evaluate(expr.Left) & maxValue
	default:
		panic("unknown expression type")
	}
}

func ExpressionString(expr *Expr) string {
	switch expr.Type {
	case OPERAND:
		return fmt.Sprintf("%06b", expr.Value)
	case NOT:
		return "NOT(" + ExpressionString(expr.Left) + ")"
	case AND:
		return "(" + ExpressionString(expr.Left) + " AND " + ExpressionString(expr.Right) + ")"
	case OR:
		return "(" + ExpressionString(expr.Left) + " OR " + ExpressionString(expr.Right) + ")"
	case XOR:
		return "(" + ExpressionString(expr.Left) + " XOR " + ExpressionString(expr.Right) + ")"
	default:
		return "UNKNOWN"
	}
}

func CountGates(expr *Expr) int {
	if expr == nil || expr.Type == OPERAND {
		return 0
	}

	if expr.Type == NOT {
		return 1 + CountGates(expr.Left)
	}

	return 1 + CountGates(expr.Left) + CountGates(expr.Right)
}

func CountInputs(expr *Expr) int {
	if expr == nil {
		return 0
	}

	if expr.Type == OPERAND {
		return 1
	}

	if expr.Type == NOT {
		return CountInputs(expr.Left)
	}

	return CountInputs(expr.Left) + CountInputs(expr.Right)
}
