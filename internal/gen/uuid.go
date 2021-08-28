package gen

import (
	"fmt"

	"github.com/google/uuid"
)

type genUUID struct{}

func (g *genUUID) Generate(option GenOption) (interface{}, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate uuid: %v\n", err)
	}
	return u.String(), nil
}
