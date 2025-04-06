package propertybasedtests

import (
	"log"
	"testing"
	"testing/quick"
)

// the main lesson of this part is teaching us about Property based testing about the
// last function below!

// Instead of just using example based tests which might be limiting (what we normally do, but usaully enough)
// try to test the fixed properties you know that the input and output should respect!
// good addition to test suite and challenges your code and its limitations.

// using testing/quick -> helps quickly check many random values within the type

var tests = []struct{
	name string 
	arabic uint16
	roman string 
}{
	{
		name: "1 gets converted to I",
		arabic: 1,
		roman: "I",
	},
	{
		name: "2 gets converted to II",
		arabic: 2,
		roman: "II",
	},
	{
		name: "3 gets converted to III",
		arabic: 3,
		roman: "III",
	},
	{
		name: "4 gets converted to IV (can't repeat more than 3 times)",
		arabic: 4,
		roman: "IV",
	},
	{
		name: "5 gets converted to V",
		arabic: 5,
		roman: "V",
	},
	{
		name: "6 gets converted to VI",
		arabic: 6,
		roman: "VI",
	},
	{
		name: "7 gets converted to VII",
		arabic: 7,
		roman: "VII",
	},
	{
		name: "9 gets converted to IX",
		arabic: 9,
		roman: "IX",
	},
	{
		name: "10 gets converted to X",
		arabic: 10,
		roman: "X",
	},
	{"18 gets converted to XVIII", 18, "XVIII"},
	{"20 gets converted to XX", 20, "XX"},
	{"39 gets converted to XXXIX", 39, "XXXIX"},
	{arabic: 900, roman: "CM"},
	{arabic: 1000, roman: "M"},
	{arabic: 1984, roman: "MCMLXXXIV"},
	{arabic: 3999, roman: "MMMCMXCIX"},
	{arabic: 2014, roman: "MMXIV"},
	{arabic: 1006, roman: "MVI"},
	{arabic: 798, roman: "DCCXCVIII"},
}

func Test_ConvertToRoman(t *testing.T){
	for _, tt := range(tests){
		
			t.Run(tt.name, func(t *testing.T) {
				got := ConvertToRoman(tt.arabic)
				want := tt.roman
		
				if got != want {
					t.Errorf("got %q, want %q", got, want)
				}
			})
	}
}

func Test_ConvertToArabic(t *testing.T){
	for _, tt := range(tests){
		
			t.Run(tt.name, func(t *testing.T) {
				got := ConvertToArabic(tt.roman)
				want := tt.arabic
		
				if got != want {
					t.Errorf("got %q, want %q", got, want)
				}
			})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			log.Println(arabic)
			return true
		}
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	 }

	if err := quick.Check(assertion, &quick.Config{
			MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}