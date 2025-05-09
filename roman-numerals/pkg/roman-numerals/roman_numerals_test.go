package roman_numerals

import "testing"

func TestRomanNumerals(t *testing.T) {
	testcases := []struct {
		desc   string
		arabic int
		roman  string
	}{
		{
			desc:   "1 gets converted to I",
			arabic: 1,
			roman:  "I",
		},
		{
			desc:   "2 gets converted to II",
			arabic: 2,
			roman:  "II",
		},
		{
			desc:   "3 gets converted to III",
			arabic: 3,
			roman:  "III",
		},
		{
			desc:   "4 gets converted to IV",
			arabic: 4,
			roman:  "IV",
		},
		{
			desc:   "5 gets converted to V",
			arabic: 5,
			roman:  "V",
		},
		{
			desc:   "6 gets converted to VI",
			arabic: 6,
			roman:  "VI",
		},
		{
			desc:   "7 gets converted to VII",
			arabic: 7,
			roman:  "VII",
		},
		{
			desc:   "8 gets converted to VIII",
			arabic: 8,
			roman:  "VIII",
		},
		{
			desc:   "9 gets converted to IX",
			arabic: 9,
			roman:  "IX",
		},
		{
			desc:   "10 gets converted to X",
			arabic: 10,
			roman:  "X",
		},
		{
			desc:   "11 gets converted to XI",
			arabic: 11,
			roman:  "XI",
		},
		{
			desc:   "14 gets converted to XIV",
			arabic: 14,
			roman:  "XIV",
		},
		{
			desc:   "18 gets converted to XVIII",
			arabic: 18,
			roman:  "XVIII",
		},
		{
			desc:   "20 gets converted to XX",
			arabic: 20,
			roman:  "XX",
		},
		{
			desc:   "39 gets converted to XXXIX",
			arabic: 39,
			roman:  "XXXIX",
		},
		{
			desc:   "40 gets converted to XL",
			arabic: 40,
			roman:  "XL",
		},
		{
			desc:   "47 gets converted to XLVII",
			arabic: 47,
			roman:  "XLVII",
		},
		{
			desc:   "49 gets converted to XLIX",
			arabic: 49,
			roman:  "XLIX",
		},
		{
			desc:   "50 gets converted to L",
			arabic: 50,
			roman:  "L",
		},
		{
			desc:   "51 gets converted to LI",
			arabic: 51,
			roman:  "LI",
		},
		{
			desc:   "59 gets converted to LIX",
			arabic: 59,
			roman:  "LIX",
		},
		{
			desc:   "88 gets converted to LXXXVIII",
			arabic: 88,
			roman:  "LXXXVIII",
		},
		{
			desc:   "89 gets converted to LXXXIX",
			arabic: 89,
			roman:  "LXXXIX",
		},
		{
			desc:   "90 gets converted to XC",
			arabic: 90,
			roman:  "XC",
		},
		{
			desc:   "94 gets converted to XCIV",
			arabic: 94,
			roman:  "XCIV",
		},
		{
			desc:   "98 gets converted to XCVIII",
			arabic: 98,
			roman:  "XCVIII",
		},
		{
			desc:   "99 gets converted to XCIX",
			arabic: 99,
			roman:  "XCIX",
		},
		{
			desc:   "100 gets converted to C",
			arabic: 100,
			roman:  "C",
		},
		{
			desc:   "101 gets converted to CI",
			arabic: 101,
			roman:  "CI",
		},
		{
			desc:   "105 gets converted to CV",
			arabic: 105,
			roman:  "CV",
		},
		{
			desc:   "108 gets converted to CVIII",
			arabic: 108,
			roman:  "CVIII",
		},
		{
			desc:   "109 gets converted to CIX",
			arabic: 109,
			roman:  "CIX",
		},
		{
			desc:   "110 gets converted to CX",
			arabic: 110,
			roman:  "CX",
		},
		{
			desc:   "114 gets converted to CXIV",
			arabic: 114,
			roman:  "CXIV",
		},
		{
			desc:   "118 gets converted to CXVIII",
			arabic: 118,
			roman:  "CXVIII",
		},
		{
			desc:   "119 gets converted to CXIX",
			arabic: 119,
			roman:  "CXIX",
		},
		{
			desc:   "400 gets converted to CD",
			arabic: 400,
			roman:  "CD",
		},
		{
			desc:   "499 gets converted to CDXCIX",
			arabic: 499,
			roman:  "CDXCIX",
		},
		{
			desc:   "500 gets converted to D",
			arabic: 500,
			roman:  "D",
		},
		{
			desc:   "798 gets converted to DCCXCVIII",
			arabic: 798,
			roman:  "DCCXCVIII",
		},
		{
			desc:   "900 gets converted to CM",
			arabic: 900,
			roman:  "CM",
		},
		{
			desc:   "1000 gets converted to M",
			arabic: 1000,
			roman:  "M",
		},
		{
			desc:   "1006 gets converted to MVI",
			arabic: 1006,
			roman:  "MVI",
		},
		{
			desc:   "1666 gets converted to MDCLXVI",
			arabic: 1666,
			roman:  "MDCLXVI",
		},
		{
			desc:   "1984 gets converted to MCMLXXXIV",
			arabic: 1984,
			roman:  "MCMLXXXIV",
		},
		{
			desc:   "1999 gets converted to MCMXCIX",
			arabic: 1999,
			roman:  "MCMXCIX",
		},
		{
			desc:   "2014 gets converted to MMXIV",
			arabic: 2014,
			roman:  "MMXIV",
		},
		{
			desc:   "3999 gets converted to MMMCMXCIX",
			arabic: 3999,
			roman:  "MMMCMXCIX",
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			actual := ConvertToRoman(testcase.arabic)
			if actual != testcase.roman {
				t.Errorf("Got %q instead of expected %q when passing %d", actual, testcase.roman, testcase.arabic)
			}
		})
	}
}
