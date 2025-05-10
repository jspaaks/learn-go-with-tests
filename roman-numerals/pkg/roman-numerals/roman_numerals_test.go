package roman_numerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

type testcase struct {
	arabic uint16
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

func TestProperties(t *testing.T) {
	config := quick.Config{
		MaxCount: 1000,
	}
	t.Run("doesnt have symbol occurring more than 3 times consecutively", func(t *testing.T) {
		assertion := func(arabic uint16) bool {
			if arabic == 0 || arabic > 3999 {
				return true
			}
			roman := []rune(ConvertArabicToRoman(arabic))
			prev := roman[0]
			nconsecutive := 1
			for _, this := range roman[1:] {
				if this == prev {
					nconsecutive++
				} else {
					nconsecutive = 1
				}
				if nconsecutive > 3 {
					return false
				}
				prev = this
			}
			return true
		}
		err := quick.Check(assertion, &config)
		if err != nil {
			t.Error("failed consecutively recurring symbol test", err)
		}
	})
	t.Run("subtractors", func(t *testing.T) {
		assertion := func(arabic uint16) bool {
			if arabic == 0 || arabic > 3999 {
				return true
			}
			var pairs = map[rune]uint16{
				'M': 1000,
				'D': 500,
				'C': 100,
				'L': 50,
				'X': 10,
				'V': 5,
				'I': 1,
			}
			roman := []rune(ConvertArabicToRoman(arabic))
			prev := roman[0]
			for _, this := range roman[1:] {
				symbolsAreInOrder := pairs[prev] >= pairs[this]
				previousSymbolsIsSubtractor := prev == 'I' || prev == 'X' || prev == 'C'
				if symbolsAreInOrder {
					prev = this
					continue
				}
				if previousSymbolsIsSubtractor {
					prev = this
					continue
				}
				return false
			}
			return true
		}
		err := quick.Check(assertion, &config)
		if err != nil {
			t.Error("failed subtractors test on input", err)
		}
	})
	t.Run("roundtrip conversion", func(t *testing.T) {
		assertion := func(arabic uint16) bool {
			if arabic == 0 || arabic > 3999 {
				return true
			}
			roman := ConvertArabicToRoman(arabic)
			return ConvertRomanToArabic(roman) == arabic
		}
		err := quick.Check(assertion, &config)
		if err != nil {
			t.Error("failed roundtrip test", err)
		}
	})
}
