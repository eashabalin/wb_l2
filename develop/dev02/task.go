package dev02

import (
	"strconv"
	"unicode"
)

// "ap11bc2d5e"

func UnpackStr(str string) string {
	if len(str) == 0 {
		return ""
	}
	newStr := ""
	numStr := ""
	curr := ""
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		if unicode.IsLetter(runes[i]) {
			if curr != "" {
				num := 1
				var err error
				if numStr != "" {
					num, err = strconv.Atoi(numStr)
					if err != nil {
						return ""
					}
				}
				for j := 0; j < num; j++ {
					newStr += curr
				}
				curr = ""
				numStr = ""
			}
			curr = string(runes[i])
		}
		if unicode.IsDigit(runes[i]) {
			numStr += string(runes[i])
		}
	}

	num := 1
	var err error
	if numStr != "" {
		num, err = strconv.Atoi(numStr)
		if err != nil {
			return ""
		}
	}
	for j := 0; j < num; j++ {
		newStr += curr
	}

	return newStr
}
