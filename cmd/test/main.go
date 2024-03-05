package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateUniqueKey() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"

	seed := int64(time.Now().UnixNano())
	src := rand.NewSource(seed)
	r := rand.New(src)

	key := make([]byte, 16)
	for i := 0; i < 16; i++ {
		key[i] = charset[r.Intn(len(charset))]
	}

	return string(key)
}

func main() {
	fmt.Println(generateUniqueKey())
}
