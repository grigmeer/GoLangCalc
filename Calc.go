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

	panic(fmt.Sprintf("некорректный тип числа: %s", input))
}

// Функция для парсинга арабских и римских чисел
func parseNumber(input string) int {
	// Проверяем, является ли введенное значение арабским числом
	num, err := strconv.Atoi(input)
	if err == nil {
		return num
	}

	// Проверяем, является ли введенное значение римским числом
	roman := strings.ToUpper(input)
	value, ok := romanToArabic[roman]
	if !ok {
		panic(fmt.Sprintf("некорректное арабское или римское число: %s", input))
	}

	return value
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
func parseOperator(input string) string {
	switch input {
	case "+", "plus":
		return "+"
	case "-", "minus":
		return "-"
	case "*", "multiply":
		return "*"
	case "/", "divide":
		return "/"
	default:
		panic(fmt.Sprintf("некорректный оператор: %s", input))
	}
}

// Функция для выполнения арифметических операций
func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		result := a + b
		if result < 0 {
			panic("результат отрицательный")
		}
		return result
	case "-":
		result := a - b
		if result < 0 {
			panic("результат отрицательный")
		}
		return result
	case "*":
		result := a * b
		if result < 0 {
			panic("результат отрицательный")
		}
		return result
	case "/":
		if b == 0 {
			panic("деление на ноль")
		}
		result := a / b
		if result < 0 {
			panic("результат отрицательный")
		}
		return result
	default:
		panic(fmt.Sprintf("некорректная операция: %s", operator))
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ошибка:", r)
		}
	}()

	fmt.Println("Добро пожаловать в калькулятор!")

	// Считываем строку с консоли
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Парсим введенное выражение
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		panic("некорректное выражение. Используйте формат: число оператор число")
	}

	a := parseNumber(parts[0])
	operator := parseOperator(parts[1])
	b := parseNumber(parts[2])

	// Проверяем, чтобы оба числа были либо арабскими, либо римскими
	if (a > 10 && b <= 10) || (a <= 10 && b > 10) {
		panic("калькулятор умеет работать только с числами меньше 10 (X)")
	}

	// Проверяем, чтобы оба числа были одного типа
	if detectNumberType(parts[0]) != detectNumberType(parts[2]) {
		panic("нельзя складывать арабские и римские числа одновременно")
	}

	// Выполняем операцию
	result := calculate(a, b, operator)

	// Выводим результат
	if detectNumberType(parts[0]) == "roman" && detectNumberType(parts[2]) == "roman" {
		fmt.Printf("%s %s %s = %s\n", parts[0], operator, parts[2], arabicToRoman(result))
	} else {
		fmt.Printf("%d %s %d = %d\n", a, operator, b, result)
	}
}
