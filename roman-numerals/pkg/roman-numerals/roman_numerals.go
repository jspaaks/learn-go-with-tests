package roman_numerals

import "strings"

var allSymbols = []struct {
	roman  string
	arabic int
}{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func ConvertToRoman(arabic int) (roman string) {

	var result strings.Builder
	for _, symbol := range allSymbols {
		for arabic >= symbol.arabic {
			result.WriteString(symbol.roman)
			arabic -= symbol.arabic
		}
	}
	return result.String()
}
