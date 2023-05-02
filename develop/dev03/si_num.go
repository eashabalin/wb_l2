package dev03

import (
	"fmt"
	"regexp"
)

const pattern = ` *([+-]?[0-9]*[.,]?[0-9]+)(E|P|T|G|M|K|k|c|m|u|n|p|f|a) *`

type SINum struct {
	num    float64
	suffix SISuffix
}

func NewSINum(s string) (*SINum, error) {
	//var si SINum
	regex := regexp.MustCompile(pattern)
	match := regex.FindStringSubmatch(s)
	fmt.Println(match)
	if len(match[0]) == len(s) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
	return nil, nil
}
