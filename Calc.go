package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Маппинг римских символов к их арабским значениям
var romanToArabic = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

// Функция для определения типа числа
func detectNumberType(input string) string {
	_, err := strconv.Atoi(input)
	if err == nil {
		return "arabic"
	}

	_, ok := romanToArabic[input]
	if ok {
		return "roman"
	}

	return "invalid"
}

// Функция для парсинга арабских и римских чисел
func parseNumber(input string) (int, error) {
	// Проверяем, является ли введенное значение арабским числом
	num, err := strconv.Atoi(input)
	if err == nil {
		return num, nil
	}

	// Проверяем, является ли введенное значение римским числом
	roman := strings.ToUpper(input)
	value, ok := romanToArabic[roman]
	if !ok {
		return 0, fmt.Errorf("некорректное арабское или римское число: %s", input)
	}

	return value, nil
}

// Функция для преобразования арабских чисел в римские
func arabicToRoman(num int) string {
	if num <= 0 || num > 1000 {
		return "Недопустимое значение"
	}

	var result strings.Builder

	for num >= 1000 {
		result.WriteString("M")
		num -= 1000
	}
	for num >= 900 {
		result.WriteString("CM")
		num -= 900
	}
	for num >= 500 {
		result.WriteString("D")
		num -= 500
	}
	for num >= 400 {
		result.WriteString("CD")
		num -= 400
	}
	for num >= 100 {
		result.WriteString("C")
		num -= 100
	}
	for num >= 90 {
		result.WriteString("XC")
		num -= 90
	}
	for num >= 50 {
		result.WriteString("L")
		num -= 50
	}
	for num >= 40 {
		result.WriteString("XL")
		num -= 40
	}
	for num >= 10 {
		result.WriteString("X")
		num -= 10
	}
	for num >= 9 {
		result.WriteString("IX")
		num -= 9
	}
	for num >= 5 {
		result.WriteString("V")
		num -= 5
	}
	for num >= 4 {
		result.WriteString("IV")
		num -= 4
	}
	for num >= 1 {
		result.WriteString("I")
		num -= 1
	}

	return result.String()
}

// Функция для парсинга операторов
func parseOperator(input string) (string, error) {
	switch input {
	case "+", "plus":
		return "+", nil
	case "-", "minus":
		return "-", nil
	case "*", "multiply":
		return "*", nil
	case "/", "divide":
		return "/", nil
	default:
		return "", fmt.Errorf("некорректный оператор: %s", input)
	}
}

// Функция для выполнения арифметических операций
func calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		result := a + b
		if result < 0 {
			return 0, fmt.Errorf("результат отрицательный")
		}
		return result, nil
	case "-":
		result := a - b
		if result < 0 {
			return 0, fmt.Errorf("результат отрицательный")
		}
		return result, nil
	case "*":
		result := a * b
		if result < 0 {
			return 0, fmt.Errorf("результат отрицательный")
		}
		return result, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		result := a / b
		if result < 0 {
			return 0, fmt.Errorf("результат отрицательный")
		}
		return result, nil
	default:
		return 0, fmt.Errorf("некорректная операция: %s", operator)
	}
}

func main() {
	fmt.Println("Добро пожаловать в калькулятор!")

	// Считываем строку с консоли
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Парсим введенное выражение
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		fmt.Println("Некорректное выражение. Используйте формат: число оператор число")
		return
	}

	a, err := parseNumber(parts[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	operator, err := parseOperator(parts[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	b, err := parseNumber(parts[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Проверяем, чтобы оба числа были либо арабскими, либо римскими
	if (a > 10 && b <= 10) || (a <= 10 && b > 10) {
		fmt.Println("Калькулятор умеет работать только с числами меньше 10 (X).")
		return
	}

	// Проверяем, чтобы оба числа были одного типа
	if detectNumberType(parts[0]) != detectNumberType(parts[2]) {
		fmt.Println("Нельзя складывать арабские и римские числа одновременно.")
		return
	}

	// Выполняем операцию
	result, err := calculate(a, b, operator)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Выводим результат
	if detectNumberType(parts[0]) == "roman" && detectNumberType(parts[2]) == "roman" {
		fmt.Printf("%s %s %s = %s\n", parts[0], operator, parts[2], arabicToRoman(result))
	} else {
		fmt.Printf("%d %s %d = %d\n", a, operator, b, result)
	}
}
