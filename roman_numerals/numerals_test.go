package main

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic int
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

func TestRomanNumerals(t *testing.T) {
	for _, tests := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", tests.Arabic, tests.Roman), func(t *testing.T) {
			got := ConvertToRoman(tests.Arabic)
			if got != tests.Roman {
				t.Errorf("got %q want %q", got, tests.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, tests := range cases[:10] {
		t.Run(fmt.Sprintf("%q gets converted to %d", tests.Roman, tests.Arabic), func(t *testing.T) {
			got := ConvertToArabic(tests.Roman)
			if got != tests.Arabic {
				t.Errorf("got %d want %d", got, tests.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic int) bool {
		if arabic > 3999 {
			return true
		} else if arabic < 0 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		FromRoman := ConvertToArabic(roman)
		return FromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed check", err)
	}
}
