package gen

import (
	"fmt"
	"math/rand"
	"time"
)

type genIPv4 struct{}

func (g *genIPv4) Generate(option GenOption) (interface{}, error) {
	rand.Seed(time.Now().UnixNano())
	num := func() int { return rand.Intn(256) }
	return fmt.Sprintf("%d.%d.%d.%d", num(), num(), num(), num()), nil
}
