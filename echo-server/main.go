package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen()
}

func listen() {
	listener, err := net.Listen("tcp", ":1026") //запуск нового сервера, обслуживающего порт 1026
	if err != nil {
		fmt.Println("Failed to open port on 1026")
		return
	}
	for { //прием новых клиентских запросов и обработка ошибок подключения
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			continue
		}
		go handle(conn) //при появлении запроса передать его в handle
	}
}

func handle(conn net.Conn) {
	reader := bufio.NewReader(conn) //попытка чтения строки из подключения
	data, err := reader.ReadBytes('\n')
	if err != nil { //в случае ошибки чтения вывести сообщение и закрыть подключение
		fmt.Println("Failed to read from socket.")
		conn.Close()
	}
	response(data, conn)
}

func response(data []byte, conn net.Conn) {
	defer func() {
		conn.Close()
	}()
	conn.Write(data)
}

// $ telnet localhost 1026