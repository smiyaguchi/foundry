package gen

import (
	"math/rand"
	"time"

	"github.com/smiyaguchi/foundry/internal/spec"
)

type genSeed struct {
	Num int
}

func (g *genSeed) Generate(spec spec.Field) (interface{}, error) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(g.Num), nil
}
