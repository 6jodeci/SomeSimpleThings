/*
Код имитирует отправку сообщений. Функция-заглушка SendRequest генерирует случайный ответ. Ответ может
быть признаком успешной отправки или одной из двух ошибок:
ErrTimeout и ErrRejected
*/
package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var ErrTimeout = errors.New("the request timed out")      //экземляр ошибки повышения времени ожидания
var ErrRejection = errors.New("the request was rejected") //экземляр ошибки отказа

var random = rand.New(rand.NewSource(36)) //генерация случайных чисел с помощью фиксированного источника

func main() {
	response, err := SendRequest("Hello") // вызов функции заглушки
	for err == ErrTimeout {               // обработка превышения времени ожидания повторением попытки
		fmt.Println("Timeout, Retrying.")
		response, err = SendRequest("Hello")
	}
	if err != nil { //обработка любой другой ошибки как сбоя
		fmt.Println(err)
	} else {
		fmt.Println(response) //если нет ошибок вывод результата
	}
}

func SendRequest(req string) (string, error) { //фуккция-заглушка имитурующая отправку сообщений
	switch random.Int() % 3 {                  // имитировать отправку сообщения слкчайно выбирая разное поведение
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejection
	default:
		return "", ErrTimeout
	}
}
