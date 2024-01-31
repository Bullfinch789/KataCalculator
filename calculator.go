package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "VIII": 8,
	"VII": 7, "VI": 6, "V": 5, "IV": 4, "III": 3, "II": 2, "I": 1,
}
var arabicNumerals = [...]int{
	100, 90, 50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
}
var a, b *int
var operators = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"/": func() int { return *a / *b },
	"*": func() int { return *a * *b },
}
var data []string

const (
	panic1 = "Ошибка, строка не является математической операцией."
	panic2 = "Ошибка, формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	panic3 = "Ошибка, используются одновременно разные системы счисления."
	panic4 = "Ошибка, в римской системе нет отрицательных чисел."
	panic5 = "Ошибка, в римской системе нет числа 0."
	panic6 = "Калькулятор работает только с арабскими цифрами от I до X или римскими цифрами от 1 до 10"
)

func calc(s string) {
	var operator string
	var stringsFound int
	numbers := make([]int, 0)
	romans := make([]string, 0)
	romansToInt := make([]int, 0)
	for idx := range operators {
		for _, val := range s {
			if idx == string(val) {
				operator += idx
				data = strings.Split(s, operator)
			}
		}
	}
	switch {
	case 1 < len(operator):
		panic(panic2)
	case 1 > len(operator):
		panic(panic1)
	}
	for _, elem := range data {
		num, err := strconv.Atoi(elem)
		if err != nil {
			stringsFound++
			romans = append(romans, elem)
		} else {
			numbers = append(numbers, num)
		}
	}

	switch stringsFound {
	case 1:
		panic(panic3)
	case 0:
		errCheck := numbers[0] > 0 && numbers[0] < 11 &&
			numbers[1] > 0 && numbers[1] < 11
		if val, ok := operators[operator]; ok && errCheck == true {
			a, b = &numbers[0], &numbers[1]
			fmt.Println(val())
		} else {
			panic(panic6)
		}
	case 2:
		for _, elem := range romans {
			if val, ok := romanNumerals[elem]; ok && val > 0 && val < 11 {
				romansToInt = append(romansToInt, val)
			} else {
				panic(panic6)
			}
		}
		if val, ok := operators[operator]; ok {
			a, b = &romansToInt[0], &romansToInt[1]
			intToRoman(val())
		}
	}
}
func intToRoman(romanResult int) {
	var romanNum string
	if romanResult == 0 {
		panic(panic5)
	} else if romanResult < 0 {
		panic(panic4)
	}
	for romanResult > 0 {
		for _, elem := range arabicNumerals {
			for i := elem; i <= romanResult; {
				for index, value := range romanNumerals {
					if value == elem {
						romanNum += index
						romanResult -= elem
					}
				}
			}
		}
	}
	fmt.Println(romanNum)
}
func main() {
	fmt.Println("Калькулятор для тестового задания для Kata Academy")
	fmt.Println("При вводе используйте только:")
	fmt.Println("Арабские цифры от 1 до 10 и римские цифры от I до X")
	fmt.Println("Введите значение:")
	reader := bufio.NewReader(os.Stdin)
	for {
		cons, _ := reader.ReadString('\n')
		c := strings.TrimSpace(cons)
		calc(strings.ToUpper(strings.TrimSpace(c)))
	}
}
