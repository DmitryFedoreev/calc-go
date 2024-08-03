package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var romanToIntMap = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var intToRomanMap = []string{
	"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
	"XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
}

func romanToInt(roman string) (int, error) {
	if value, exists := romanToIntMap[roman]; exists {
		return value, nil
	}
	return 0, errors.New("invalid Roman numeral")
}

func intToRoman(num int) (string, error) {
	if num <= 0 || num >= len(intToRomanMap) {
		return "", errors.New("resulting Roman numeral out of range")
	}
	return intToRomanMap[num], nil
}

func isRomanNumeral(input string) bool {
	_, exists := romanToIntMap[input]
	return exists
}

func performOperation(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	default:
		return 0, errors.New("invalid operator")
	}
}

func calculate(input string) (string, error) {
	// Remove spaces
	input = strings.ReplaceAll(input, " ", "")

	// Check if input matches the pattern "number operator number"
	r := regexp.MustCompile(`^(\d+|[IVXLCDM]+)([\+\-\*/])(\d+|[IVXLCDM]+)$`)
	matches := r.FindStringSubmatch(input)
	if matches == nil {
		return "", errors.New("invalid input format")
	}

	aStr, op, bStr := matches[1], matches[2], matches[3]

	var a, b int
	var err error
	romanMode := false

	// Determine if the input is in Roman numerals or Arabic numbers
	if isRomanNumeral(aStr) && isRomanNumeral(bStr) {
		romanMode = true
		a, err = romanToInt(aStr)
		if err != nil {
			return "", err
		}
		b, err = romanToInt(bStr)
		if err != nil {
			return "", err
		}
	} else if !isRomanNumeral(aStr) && !isRomanNumeral(bStr) {
		a, err = strconv.Atoi(aStr)
		if err != nil {
			return "", err
		}
		b, err = strconv.Atoi(bStr)
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("mixed numeral systems are not allowed")
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		return "", errors.New("numbers must be between 1 and 10 inclusive")
	}

	result, err := performOperation(a, b, op)
	if err != nil {
		return "", err
	}

	if romanMode {
		if result < 1 {
			return "", errors.New("resulting Roman numeral must be positive")
		}
		return intToRoman(result)
	}
	return strconv.Itoa(result), nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run calc.go <expression>")
		return
	}

	input := os.Args[1]
	result, err := calculate(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(result)
}
