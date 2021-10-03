package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	done := make(chan bool) // bool - канал для сообщения о завершении
	until := time.After(5 * time.Second)
	go send(msg, done) // передача двух каналов в send
	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			done <- true
			time.Sleep(1 * time.Second)
			return
		}
	}
}

func send(ch chan<- string, done <-chan bool) { //ch используется для отправки, а done - для получения
	for {
		select {
		case <-done: // завершить работу после получения сообщения из канала done
			println("Done - Success")
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(1 * time.Second)
		}
	}

}
