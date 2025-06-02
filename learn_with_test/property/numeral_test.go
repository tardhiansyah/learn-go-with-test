package property

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestConvertToRoman(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)

			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)

			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestConvertToArabicRecursively(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabicRecursively(test.Roman)

			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

// TestPropertiesOfConversion checks that converting to Roman and back to Arabic
func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		log.Println("Testing Arabic:", arabic)

		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		Values: func(values []reflect.Value, rand *rand.Rand) {
			values[0] = reflect.ValueOf(uint16(rand.Intn(4000)))
		},
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}

// TestPropertiesCannotHaveMoreThanThreeConsecutiveSymbols checks that no Roman numeral has more than three consecutive symbols
func TestPropertiesCannotHaveMoreThanThreeConsecutiveSymbols(t *testing.T) {
	assertion := func(roman string) bool {
		log.Println("Testing Roman:", roman)

		for _, numeral := range allRomanNumerals {
			if len(numeral.Symbol) > 1 && strings.Count(roman, numeral.Symbol) > 3 {
				return false
			}
		}
		return true
	}

	if err := quick.Check(assertion, &quick.Config{
		Values: func(values []reflect.Value, r *rand.Rand) {
			validChars := []rune("IVXLCDM")
			length := rand.Intn(15) + 1
			result := make([]rune, length)
			for i := range result {
				result[i] = validChars[rand.Intn(len(validChars))]
			}

			values[0] = reflect.ValueOf(string(result))
		},

		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}

func BenchmarkConvertToRoman(b *testing.B) {
	for b.Loop() {
		for _, test := range cases {
			_ = ConvertToRoman(test.Arabic)
		}
	}
}

func BenchmarkConvertToArabicRecursively(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range cases {
			_ = ConvertToArabicRecursively(test.Roman)
		}
	}
}
