package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	flag := 0
	// Убираем все пробелы, чтобы привести выражение к исходному слитному виду
	expression = strings.ReplaceAll(expression, " ", "")
	// Отделяем все операции пробелами
	expression = strings.ReplaceAll(expression, "+", " + ")
	expression = strings.ReplaceAll(expression, "-", " - ")
	expression = strings.ReplaceAll(expression, "*", " * ")
	expression = strings.ReplaceAll(expression, "/", " / ")
	expression = strings.ReplaceAll(expression, "(", " ( ")
	expression = strings.ReplaceAll(expression, ")", " ) ")

	// Разделяем выражение на части
	parts := strings.Fields(expression)
	// Проверяем, что выражение корректно
	if len(parts) < 3 || isOperand(parts[len(parts)-1]) || isOperand(parts[0]) {
		return 0, errors.New("некорректное выражение")
	}

	// Инициализируем стек для чисел и стек для операций
	numStack := make([]float64, 0)
	opStack := make([]string, 0)

	for i, part := range parts {
		if isOperand(part) && isOperand(parts[i-1]) {
			return 0, errors.New("некорректное выражение")
		}
		// Если часть - число, добавляем его в стек чисел
		if _, err := strconv.ParseFloat(part, 64); err == nil {
			num, _ := strconv.ParseFloat(part, 64)
			numStack = append(numStack, num)
		} else if isOperand(part) || part == "(" || part == ")" {
			// Если часть - операция, обрабатываем ее
			if part == "(" {
				opStack = append(opStack, part)
				flag++
			} else if part == ")" {
				flag--
				if flag < 0 {
					return 0, errors.New("некорректное выражение")
				}
				if flag+1 != 0 {
					// Выполняем операции в скобках
					for i := 1; i < len(opStack); i++ {
						if len(opStack) == 0 {
							break
						}
						op := opStack[len(opStack)-i]
						if op == "*" || op == "/" {
							opStack = append(opStack[:len(opStack)-i], opStack[len(opStack)-i+1:]...)
							num2 := numStack[len(numStack)-i]
							numStack = append(numStack[:len(numStack)-i], numStack[len(numStack)-i+1:]...)
							num1 := numStack[len(numStack)-i]
							numStack = append(numStack[:len(numStack)-i], numStack[len(numStack)-i+1:]...)
							switch op {
							case "*":
								numStack = append(numStack, num1*num2)
							case "/":
								numStack = append(numStack, num1/num2)
							}
						} else if op == "(" {
							break
						}
					}
					for i := 1; i < len(opStack); i++ {
						if len(opStack) == 0 {
							break
						}
						op := opStack[len(opStack)-i]
						if op == "+" || op == "-" {
							opStack = append(opStack[:len(opStack)-i], opStack[len(opStack)-i+1:]...)
							num2 := numStack[len(numStack)-i]
							numStack = append(numStack[:len(numStack)-i], numStack[len(numStack)-i+1:]...)
							num1 := numStack[len(numStack)-i]
							numStack = append(numStack[:len(numStack)-i], numStack[len(numStack)-i+1:]...)
							switch op {
							case "+":
								numStack = append(numStack, num1+num2)
							case "-":
								numStack = append(numStack, num1-num2)
							}
						} else if op == "(" {
							break
						}
					}
					// Удаляем скобку из стека операций
					opStack = opStack[:len(opStack)-1]
				}
			} else {
				// Если приоритет текущей операции выше или равен приоритету операции в стеке,
				// выполняем операцию в стеке и добавляем результат в стек чисел
				if flag == 0 {
					for i := 1; i < len(opStack); i++ {
						if len(opStack) == 0 {
							break
						}
						op := opStack[len(opStack)-i]
						if op == "*" || op == "/" {
							opStack = append(opStack[:len(opStack)-i], opStack[len(opStack)-i+1:]...)
							num2 := numStack[len(numStack)-i]
							numStack = append(numStack[:len(numStack)-i], numStack[len(numStack)-i+1:]...)
							num1 := numStack[len(numStack)-i]
							numStack = append(numStack[:len(numStack)-i], numStack[len(numStack)-i+1:]...)
							switch op {
							case "*":
								numStack = append(numStack, num1*num2)
							case "/":
								numStack = append(numStack, num1/num2)
							}
						} else {
							continue
						}
					}
					if !(part == "*" || part == "/") {
						for i := 1; i < len(opStack); i++ {
							if len(opStack) == 0 {
								break
							}
							op := opStack[len(opStack)-i]
							if op == "+" || op == "-" {
								opStack = append(opStack[:len(opStack)-i], opStack[len(opStack)-i+1:]...)
								num2 := numStack[len(numStack)-i]
								numStack = append(numStack[:len(numStack)-i], numStack[len(numStack)-i+1:]...)
								num1 := numStack[len(numStack)-i]
								numStack = append(numStack[:len(numStack)-i], numStack[len(numStack)-i+1:]...)
								switch op {
								case "+":
									numStack = append(numStack, num1+num2)
								case "-":
									numStack = append(numStack, num1-num2)
								}
							} else if op == "(" {
								break
							}
						}
					}
				}
				// Добавляем текущую операцию в стек операций
				opStack = append(opStack, part)
			}
		} else {
			return 0, errors.New("некорректное выражение")
		}
	}

	if flag != 0 {
		return 0, errors.New("некорректное выражение")
	}

	// Выполняем оставшиеся операции
	for {
		if len(opStack) == 0 {
			break
		}
		op := opStack[len(opStack)-1]
		opStack = opStack[:len(opStack)-1]
		num2 := numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-1]
		num1 := numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-1]
		switch op {
		case "+":
			numStack = append(numStack, num1+num2)
		case "-":
			numStack = append(numStack, num1-num2)
		case "*":
			numStack = append(numStack, num1*num2)
		case "/":
			numStack = append(numStack, num1/num2)
		}
	}

	// Возвращаем результат
	return numStack[0], nil
}

func isOperand(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

func main() {
	expression := "10.5-1"
	result, err := Calc(expression)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
