package gen

import (
	"math/rand"
	"time"
)

type genSeed struct{}

func (g *genSeed) Generate(option GenOption) (interface{}, error) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10), nil
}
