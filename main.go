package main

import (
	"fmt"
	"math/rand"
)

func main() {
	config := Config{
		Bits:   6,
		AlphaG: 1,
		Trials: 25,
	}

	fmt.Println("BooleanBench")
	fmt.Println("====================")
	fmt.Println()

	for n := 1; n <= 10; n++ {
		level := GenerateLevel(n, config)

		fmt.Printf(
			"Level %d: %d bits, %d gates\n",
			level.Number,
			level.Bits,
			level.Gates,
		)
	}

	fmt.Println()
	fmt.Println("====================")
	fmt.Println("Expression generation")
	fmt.Println()

	rng := rand.New(rand.NewSource(12345))

	for n := 1; n <= 10; n++ {
		level := GenerateLevel(n, config)
		expr := GenerateExpression(level.Gates, rng)

		fmt.Printf("Level %d\n", n)
		fmt.Printf("  Expected gates: %d\n", level.Gates)
		fmt.Printf("  Actual gates:   %d\n", CountGates(expr))
		fmt.Printf("  Inputs:         %d\n", CountInputs(expr))
		fmt.Printf("  Expression:     %s\n", ExpressionString(expr))
		fmt.Printf("  Answer:         %06b\n", Evaluate(expr))
		fmt.Println()
	}
}
