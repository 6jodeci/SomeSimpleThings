package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("Can't divide by zero")

func main() {
	fmt.Println("Divide 1 by 0") //первое деление выполняет функция checkDivide, возвращающая ошибку
	_, err := checkDivide(1, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Println("Divide 2 by 0") //такое же деление с помощью функции divide
	divide(2, 0)
}

func checkDivide(a, b int) (int, error) { //возвращает ошибку при делении на 0
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return divide(a, b), nil
}

func divide(a, b int) int { //выполнение операции без проверок
	return a / b
}
