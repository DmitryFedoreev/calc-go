package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Словарь для преобразования римских чисел в арабские и наоборот
var romanToIntMap = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var intToRomanMap = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
}

func toInt(roman string) (int, error) {
	if val, ok := romanToIntMap[roman]; ok {
		return val, nil
	}
	return 0, errors.New("недопустимое римское число")
}

func toRoman(num int) (string, error) {
	if num < 1 {
		return "", errors.New("результат меньше единицы для римских чисел недопустим")
	}
	if val, ok := intToRomanMap[num]; ok {
		return val, nil
	}
	return "", errors.New("число выходит за пределы поддерживаемых значений")
}

func calculate(expression string) (string, error) {
	// Удаляем пробелы и разбиваем строку на операнды и оператор
	expression = strings.ReplaceAll(expression, " ", "")
	var operator string
	var operands []string

	if strings.Contains(expression, "+") {
		operands = strings.Split(expression, "+")
		operator = "+"
	} else if strings.Contains(expression, "-") {
		operands = strings.Split(expression, "-")
		operator = "-"
	} else if strings.Contains(expression, "*") {
		operands = strings.Split(expression, "*")
		operator = "*"
	} else if strings.Contains(expression, "/") {
		operands = strings.Split(expression, "/")
		operator = "/"
	} else {
		return "", errors.New("неверный формат математической операции")
	}

	if len(operands) != 2 {
		return "", errors.New("формат математической операции не удовлетворяет заданию")
	}

	operand1 := operands[0]
	operand2 := operands[1]

	var num1, num2 int
	var isRoman bool

	// Проверяем, являются ли операнды римскими или арабскими числами
	if n1, err := strconv.Atoi(operand1); err == nil {
		if n2, err := strconv.Atoi(operand2); err == nil {
			num1 = n1
			num2 = n2
		} else {
			return "", errors.New("используются одновременно разные системы счисления или неподходящие числа")
		}
	} else if n1, err := toInt(operand1); err == nil {
		if n2, err := toInt(operand2); err == nil {
			num1 = n1
			num2 = n2
			isRoman = true
		} else {
			return "", errors.New("используются одновременно разные системы счисления или неподходящие числа")
		}
	} else {
		return "", errors.New("используются одновременно разные системы счисления или неподходящие числа")
	}

	if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		return "", errors.New("числа должны быть в диапазоне от 1 до 10 включительно")
	}

	// Выполняем арифметическую операцию
	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return "", errors.New("деление на ноль недопустимо")
		}
		result = num1 / num2
	}

	// Возвращаем результат в нужной системе счисления
	if isRoman {
		return toRoman(result)
	}
	return strconv.Itoa(result), nil
}

func main() {
	fmt.Print("Введите математическую операцию: ")
	var expression string
	fmt.Scanln(&expression)

	result, err := calculate(expression)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}
