/*
Эта программа использует сопрограмму для эхо-вывода в фоновом
режиме, пока в основном режиме работает таймер. Если
*/
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func echo(in io.Reader, out io.Writer) { // является обычной функцией
	io.Copy(out, in)
}

func main() {
	fmt.Println("Let's Print!")
	go echo(os.Stdin, os.Stdout) //вызов функции echo как go - подпрограммы
	time.Sleep(15 * time.Second) // 15s пауза
	fmt.Println("Timed out.")    // Вывод сообщения о завершении ожидания
	os.Exit(0)
}
