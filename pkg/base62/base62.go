package base62

import (
	"errors"
	"strings"
)

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base     = uint64(len(alphabet))
)

// Encode converts a uint64 ID into a Base62 string
func Encode(number uint64) string {
	if number == 0 {
		return string(alphabet[0])
	}

	var builder strings.Builder
	for number > 0 {
		remainder := number % base
		builder.WriteByte(alphabet[remainder])
		number = number / base
	}

	// Reverse the encoded characters
	bytes := []byte(builder.String())
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	return string(bytes)
}

// Decode converts a Base62 string back into a uint64 ID
func Decode(encoded string) (uint64, error) {
	var number uint64
	for _, char := range encoded {
		index := strings.IndexRune(alphabet, char)
		if index == -1 {
			return 0, errors.New("invalid character in base62 string")
		}
		number = number*base + uint64(index)
	}
	return number, nil
}
