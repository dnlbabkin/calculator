package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToInt = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var intToRoman = []struct {
	value   int
	numeral string
}{
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например, 3 + 4 или VI * II):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Некорректный ввод. Ожидается формат 'a + b'.")
	}

	aStr, op, bStr := parts[0], parts[1], parts[2]

	isRoman := isRomanNumeral(aStr) && isRomanNumeral(bStr)
	isArabic := isArabicNumeral(aStr) && isArabicNumeral(bStr)

	if !isRoman && !isArabic {
		panic("Некорректный ввод. Используйте только арабские или только римские цифры.")
	}

	var a, b int
	var err error

	if isRoman {
		a = romanToArabic(aStr)
		b = romanToArabic(bStr)
	} else {
		a, err = strconv.Atoi(aStr)
		if err != nil {
			panic("Некорректный ввод числа.")
		}
		b, err = strconv.Atoi(bStr)
		if err != nil {
			panic("Некорректный ввод числа.")
		}
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Числа должны быть в диапазоне от 1 до 10 включительно.")
	}

	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("Некорректная операция. Используйте +, -, * или /.")
	}

	if isRoman {
		if result < 1 {
			panic("Результат работы с римскими числами не может быть меньше 1.")
		}
		fmt.Println(arabicToRoman(result))
	} else {
		fmt.Println(result)
	}
}

func isRomanNumeral(s string) bool {
	_, exists := romanToInt[s]
	return exists
}

func isArabicNumeral(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func romanToArabic(s string) int {
	result := 0
	for len(s) > 0 {
		for _, pair := range intToRoman {
			if strings.HasPrefix(s, pair.numeral) {
				result += pair.value
				s = strings.TrimPrefix(s, pair.numeral)
				break
			}
		}
	}
	return result
}

func arabicToRoman(num int) string {
	result := ""
	for _, pair := range intToRoman {
		for num >= pair.value {
			result += pair.numeral
			num -= pair.value
		}
	}
	return result
}
