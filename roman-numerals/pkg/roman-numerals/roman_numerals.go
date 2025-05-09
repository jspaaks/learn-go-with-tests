package roman_numerals

import "strings"

var pairs = []struct {
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

func ConvertArabicToRoman(arabic int) (roman string) {
	var result strings.Builder
	for _, pair := range pairs {
		for arabic >= pair.arabic {
			result.WriteString(pair.roman)
			arabic -= pair.arabic
		}
	}
	return result.String()
}


func ConvertRomanToArabic(roman string) (arabic int) {
	for _, pair := range pairs {
		for strings.HasPrefix(roman, pair.roman) {
			roman = strings.TrimPrefix(roman, pair.roman)
			arabic += pair.arabic
		}
	}
	return arabic
}