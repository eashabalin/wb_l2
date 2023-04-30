package dev02

import (
	"strconv"
	"unicode"
)

// "ap11bc2d\5e"

func UnpackStr(str string) string {
	if len(str) == 0 {
		return ""
	}
	newStr := ""
	numStr := ""
	curr := ""
	runes := []rune(str)
	escape := false
	for i := 0; i < len(runes); i++ {
		if runes[i] == '\\' && !escape {
			escape = true
			continue
		}
		if unicode.IsLetter(runes[i]) || escape == true {
			if escape {
				escape = false
			}
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
			continue
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
