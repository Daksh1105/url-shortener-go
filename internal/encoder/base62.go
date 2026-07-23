package encoder

import "strings"

var base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Base62Encode(id uint64) string {
    if id == 0 {
        return "0"
    }
    var result string
    for id > 0 {
        remainder := id % 62
        result = string(base62Chars[remainder]) + result
        id = id / 62
    }
    return result
}

func Base62Decode(str string) uint64 {
    var result uint64
    for i := 0; i < len(str); i++ {
        char := str[i]
        index := strings.Index(base62Chars, string(char))
        result = result*62 + uint64(index)
    }
    return result
}