package enum

type Gender byte

const (
	Male   Gender = 'M'
	Female Gender = 'F'
	Other  Gender = 'O'
)

func (g Gender) String() string {
	switch g {
	case Male:
		return "先生"
	case Female:
		return "女士"
	case Other:
		return "保密"
	}
	return "unknown"
}

func GetGender(g string) Gender {
	switch g {
	case "M":
		return Male
	case "F":
		return Female
	case "O":
		return Other
	}
	return Other
}
