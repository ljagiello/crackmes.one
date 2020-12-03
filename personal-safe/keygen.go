package main

import (
	"fmt"
	"math/rand"
	"time"
)

const str1Sum int = 262
const str2Sum int = 273
const minCharacterCode int = 48
const maxCharacterCode int = 122

// RandomCharacterCode returns random character code between min and max
func RandomCharacterCode(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}

// GenerateString returns random string with expected sum of characters
func GenerateString(sum int) string {
	var char1, char2, char3, char4 int
	var str string

	for {
		char1 = RandomCharacterCode(minCharacterCode, maxCharacterCode)
		char2 = RandomCharacterCode(minCharacterCode, maxCharacterCode)
		char3 = RandomCharacterCode(minCharacterCode, maxCharacterCode)
		char4 = RandomCharacterCode(minCharacterCode, maxCharacterCode)
		if char1+char2+char3+char4 == sum {
			break
		}
	}
	for _, v := range []int{char1, char2, char3, char4} {
		str = str + fmt.Sprintf("%c", v)
	}
	return str
}

func main() {
	str1 := GenerateString(str1Sum)
	str2 := GenerateString(str2Sum)
	fmt.Println(str1 + str2 + str1 + str2)
}
