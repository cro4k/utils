/**
 * Created by 93201 on 2017/6/8.
 */
package utils

import (
	"crypto/rand"
	"math/big"
)

const (
	S_NUMBER       = "0123456789"
	S_LETTER_UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	S_LETTER_LOWER = "abcdefghijklmnopqrstuvwxyz"
	S_LETTER       = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	S_CHAR         = "!@#$%^&*()_+?:"
	S_DEFAULT      = S_LETTER + S_NUMBER
)

func RandomString(n int, seeds ...string) string {
	var seed = S_DEFAULT
	if len(seeds) > 0 {
		seed = seeds[0]
	}
	buffer := make([]byte, n)
	max := big.NewInt(int64(len(seed)))

	for i := 0; i < n; i++ {
		index, err := randomInt(max)
		if err != nil {
			return ""
		}
		buffer[i] = seed[index]
	}
	return string(buffer)
}

func randomInt(max *big.Int) (int, error) {
	r, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return int(r.Int64()), nil
}
