package roman_numerals

import (
	"fmt"
	"testing"
)

type testcase struct {
	arabic int
	roman  string
}

var testcases = []testcase{
	{arabic: 1, roman: "I"},
	{arabic: 2, roman: "II"},
	{arabic: 3, roman: "III"},
	{arabic: 4, roman: "IV"},
	{arabic: 5, roman: "V"},
	{arabic: 6, roman: "VI"},
	{arabic: 7, roman: "VII"},
	{arabic: 8, roman: "VIII"},
	{arabic: 9, roman: "IX"},
	{arabic: 10, roman: "X"},
	{arabic: 11, roman: "XI"},
	{arabic: 14, roman: "XIV"},
	{arabic: 18, roman: "XVIII"},
	{arabic: 20, roman: "XX"},
	{arabic: 39, roman: "XXXIX"},
	{arabic: 40, roman: "XL"},
	{arabic: 47, roman: "XLVII"},
	{arabic: 49, roman: "XLIX"},
	{arabic: 50, roman: "L"},
	{arabic: 51, roman: "LI"},
	{arabic: 59, roman: "LIX"},
	{arabic: 88, roman: "LXXXVIII"},
	{arabic: 89, roman: "LXXXIX"},
	{arabic: 90, roman: "XC"},
	{arabic: 94, roman: "XCIV"},
	{arabic: 98, roman: "XCVIII"},
	{arabic: 99, roman: "XCIX"},
	{arabic: 100, roman: "C"},
	{arabic: 101, roman: "CI"},
	{arabic: 105, roman: "CV"},
	{arabic: 108, roman: "CVIII"},
	{arabic: 109, roman: "CIX"},
	{arabic: 110, roman: "CX"},
	{arabic: 114, roman: "CXIV"},
	{arabic: 118, roman: "CXVIII"},
	{arabic: 119, roman: "CXIX"},
	{arabic: 400, roman: "CD"},
	{arabic: 499, roman: "CDXCIX"},
	{arabic: 500, roman: "D"},
	{arabic: 798, roman: "DCCXCVIII"},
	{arabic: 900, roman: "CM"},
	{arabic: 1000, roman: "M"},
	{arabic: 1006, roman: "MVI"},
	{arabic: 1666, roman: "MDCLXVI"},
	{arabic: 1984, roman: "MCMLXXXIV"},
	{arabic: 1999, roman: "MCMXCIX"},
	{arabic: 2014, roman: "MMXIV"},
	{arabic: 3999, roman: "MMMCMXCIX"},
}

func descArabicToRoman(testcase testcase) string {
	return fmt.Sprintf("%d gets converted to %s", testcase.arabic, testcase.roman)
}
func descRomanToArabic(testcase testcase) string {
	return fmt.Sprintf("%s gets converted to %d", testcase.roman, testcase.arabic)
}

func TestConvertArabicToRoman(t *testing.T) {
	for _, testcase := range testcases {
		desc := descArabicToRoman(testcase)
		t.Run(desc, func(t *testing.T) {
			actual := ConvertArabicToRoman(testcase.arabic)
			if actual != testcase.roman {
				t.Errorf("Got %q instead of expected %q when passing %d", actual, testcase.roman, testcase.arabic)
			}
		})
	}
}

func TestConvertRomanToArabic(t *testing.T) {
	for _, testcase := range testcases {
		desc := descRomanToArabic(testcase)
		t.Run(desc, func(t *testing.T) {
			actual := ConvertRomanToArabic(testcase.roman)
			if actual != testcase.arabic {
				t.Errorf("Got %d instead of expected %d when passing %q", actual, testcase.arabic, testcase.roman)
			}
		})
	}
}
