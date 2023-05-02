package dev03

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

func IsSISuffix(s string) bool {
	all := []string{
		Exa.String(),
		Peta.String(),
		Tera.String(),
		Giga.String(),
		Mega.String(),
		Kilo.String(),
		centi.String(),
		milli.String(),
		micro.String(),
		nano.String(),
		pico.String(),
		femto.String(),
		atto.String(),
	}
	for _, v := range all {
		if v == s {
			return true
		}
	}
	return false
}
