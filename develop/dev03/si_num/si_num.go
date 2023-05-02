package si_num

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

const pattern = ` *([+-]?[0-9]*[.,]?[0-9]+)(E|P|T|G|M|K|k|c|m|u|n|p|f|a) *`

type SINum struct {
	num    float64
	suffix SISuffix
}

func NewSINum(s string) (*SINum, error) {
	var si SINum
	regex := regexp.MustCompile(pattern)
	match := regex.FindStringSubmatch(s)
	if len(match) > 0 && len(match[0]) == len(s) {
		suffix, err := NewSISuffix(match[2])
		if err != nil {
			return nil, err
		}
		si.suffix = suffix
		num, err := strconv.ParseFloat(match[1], 64)
		if err != nil {
			return nil, err
		}
		si.num = num
		return &si, nil
	}
	return nil, errors.New("couldn't parse string to si num")
}

func (n SINum) String() string {
	return fmt.Sprintf("%v%v", n.num, n.suffix)
}

func (n SINum) LessThan(other SINum) bool {
	n1 := n.num * math.Pow(10, float64(n.suffix))
	n2 := other.num * math.Pow(10, float64(other.suffix))
	return n1 < n2
}
