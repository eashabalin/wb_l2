package si_num

import "errors"

// SISuffix is an int type. Its value represents the power of 10. For example: kilo = 3 (10^3), nano = 9 (10^-9).
type SISuffix int

const (
	Exa   SISuffix = 18
	Peta  SISuffix = 15
	Tera  SISuffix = 12
	Giga  SISuffix = 9
	Mega  SISuffix = 6
	Kilo  SISuffix = 3
	centi SISuffix = -2
	milli SISuffix = -3
	micro SISuffix = -6
	nano  SISuffix = -9
	pico  SISuffix = -12
	femto SISuffix = -15
	atto  SISuffix = -18
)

func (s SISuffix) String() string {
	switch s {
	case 18:
		return "E"
	case 15:
		return "P"
	case 12:
		return "T"
	case 9:
		return "G"
	case 6:
		return "M"
	case 3:
		return "k"
	case -2:
		return "c"
	case -3:
		return "m"
	case -6:
		return "u"
	case -9:
		return "n"
	case -12:
		return "p"
	case -15:
		return "f"
	case -18:
		return "a"
	}
	return ""
}

func NewSISuffix(s string) (SISuffix, error) {
	switch s {
	case "E":
		return Exa, nil
	case "P":
		return Peta, nil
	case "T":
		return Tera, nil
	case "G":
		return Giga, nil
	case "M":
		return Mega, nil
	case "K":
		return Kilo, nil
	case "k":
		return Kilo, nil
	case "c":
		return centi, nil
	case "m":
		return milli, nil
	case "u":
		return micro, nil
	case "n":
		return nano, nil
	case "p":
		return pico, nil
	case "f":
		return femto, nil
	case "a":
		return atto, nil
	default:
		return 0, errors.New("string is not a si suffix")
	}
}
