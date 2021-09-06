package rnd

import (
	"math/rand"
	"time"
)

func RandString(n int) string {
	rand.Seed(time.Now().Unix())

	dict := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	blob := make([]rune, n)
	for i := range blob {
		blob[i] = dict[rand.Intn(len(dict))]
	}

	return string(blob)
}
