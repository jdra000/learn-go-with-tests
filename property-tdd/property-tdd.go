package main

import (
	"errors"
	"strings"
)

type validNumber uint16

func (v validNumber) LessThan() bool{
	if v <= 3999 {
		return true 
	}
	return false
}

type RomanNumeral struct {
	Value validNumber
	Symbol string
}
var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic validNumber) (string, error) {
	if arabic.LessThan() {
		var res strings.Builder
	
		for _, numeral := range allRomanNumerals{
			for arabic >= numeral.Value {
				res.WriteString(numeral.Symbol)
				arabic -= numeral.Value
				}
			}
		return res.String(), nil
	}
	return "", errors.New("Invalid number")
}
func ConvertToArabic(roman string) validNumber{
	var res validNumber

	for _, numeral := range allRomanNumerals{
		for strings.HasPrefix(roman, numeral.Symbol){
			res += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}
	return res
}
