package gen

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type genDefault struct {
	typ string
}

func (g *genDefault) Generate(option GenOption) (interface{}, error) {
	switch g.typ {
	case "int":
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(100), nil
	case "string":
		str, err := randomStr(10)
		if err != nil {
			return nil, err
		}
		return str, nil
	case "bool":
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(2)
		list := []bool{true, false}
		return list[num], nil
	default:
		return nil, fmt.Errorf("default generator not support typ: %s\n", g.typ)
	}
}

func randomStr(digit int32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
