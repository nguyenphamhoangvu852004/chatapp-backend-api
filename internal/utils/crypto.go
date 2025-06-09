package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func GetHash(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	hashBytes := hash.Sum(nil)

	return hex.EncodeToString(hashBytes)
}

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}
