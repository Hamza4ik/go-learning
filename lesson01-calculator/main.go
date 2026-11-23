package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Ошибка: на 0 делить нельзя!")
		}
		return a / b, nil
	case "%":
		if b == 0 {
			return 0, errors.New("Остаток от делении на ноль не определен")
		}
		return math.Mod(a, b), nil
	case "^":
		return math.Pow(a, b), nil
	default:
		return 0, fmt.Errorf("Ошибка: Неизвестная команда: %q", op)
	}
}

func readNumber(prompt string) (float64, error) {
	fmt.Println(prompt)
	var s string
	if _, err := fmt.Scan(&s); err != nil {
		return 0, err
	}
	return strconv.ParseFloat(s, 64)
}

var last float64
var hasLast bool

func main() {

	for {

		fmt.Println("*===========Калькулятор===========*")
		fmt.Println("Оператор (/ * - + %, ^ , c) или q для выхода")
		var op string
		if _, err := fmt.Scan(&op); err != nil || op == "q" {
			fmt.Println("Пока!")
			return
		}
		if op == "c" {
			if hasLast {
				fmt.Println("Последный результат: ", last)
			} else {
				fmt.Println("Результатов пока нет")
			}
			continue
		}

		a, err := readNumber("Первое число: ")
		if err != nil {
			fmt.Println("нужно ввести число")
			continue
		}
		b, err := readNumber("Второе число: ")
		if err != nil {
			fmt.Println("нужно ввести число!")
			continue
		}

		result, err := calculate(a, b, op)
		if err != nil {
			fmt.Println("Ошибка: ", err)
			continue
		}
		last = result
		hasLast = true

		fmt.Println("Результат: ", result)
	}
}
