package trial

import (
	"math/rand"

	"github.com/mayank-muthanna/Boolean-Bench/level"
)

type Trial struct {
	Level      int
	Seed       int64
	Expression string
	Answer     uint8
}

func Generate(levelNumber int, seed int64, config level.Config) Trial {
	lvl := level.Generate(levelNumber, config)

	rng := rand.New(rand.NewSource(seed))

	expression := level.GenerateExpression(lvl.Gates, rng)

	return Trial{
		Level:      levelNumber,
		Seed:       seed,
		Expression: level.ExpressionString(expression),
		Answer:     level.Evaluate(expression),
	}
}
