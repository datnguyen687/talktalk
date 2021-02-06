package utils

import "math/rand"

// GenerateActivationCode ...
func GenerateActivationCode(length int) string {
	if length <= 0 {
		return ""
	}
	r := []rune("1234567890QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm")
	l := len(r)
	result := make([]rune, length)
	for i := 0; i < length; i++ {
		result[i] = r[rand.Int()%l]

	}

	return string(result)
}
