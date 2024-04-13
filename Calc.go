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

// Функция для парсинга арабских и римских чисел
func parseNumber(input string) (int, error) {
	// Пытаемся спарсить как арабское число
	num, err := strconv.Atoi(input)
	if err == nil {
		if num < 1 || num > 10 {
			return 0, fmt.Errorf("число должно быть от 1 до 10")
		}
		return num, nil
	}

	// Пытаемся спарсить как римское число
	roman := strings.ToUpper(input)
	value, ok := romanToArabic[roman]
	if !ok {
		return 0, fmt.Errorf("некорректное арабское или римское число: %s", input)
	}

	return value, nil
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
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return a / b, nil
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

	// Выполняем операцию
	result, err := calculate(a, b, operator)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Выводим результат
	fmt.Printf("%d %s %d = %d\n", a, operator, b, result)
}
