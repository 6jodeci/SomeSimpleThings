package main

import (
	"log"
	"os"
)

func main()  {
	logfile, _ := os.Create("./logging/logs/log.txt") // создание файла журнала
	defer logfile.Close() // закрытие в дефере

	logger := log.New(logfile, "example", log.LstdFlags|log.Lshortfile) // создание регистра

	logger.Println("This is a regular message") // отсылка сообщений
	logger.Fatalln("This is a fatal error")
	logger.Println("This is the end of function") // не будет выполнена
}
