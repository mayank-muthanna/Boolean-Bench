package main

import (
	"fmt"

	"github.com/mayank-muthanna/Boolean-Bench/level"
	"github.com/mayank-muthanna/Boolean-Bench/trial"
)

func main() {
	config := level.Config{
		Bits:   6,
		AlphaG: 1,
		Trials: 25,
	}

	for n := 1; n <= 10; n++ {
		t := trial.Generate(n, int64(n), config)

		fmt.Printf("Level %d\n", t.Level)
		fmt.Printf("Expression: %s\n", t.Expression)
		fmt.Printf("Answer: %06b\n\n", t.Answer)
	}
}
