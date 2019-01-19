package main

import (
	"log"
	"os"
)

func main() {
	Log2()
}
func Log2() {
	logFile, err := os.OpenFile("log2.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Panic(err)
	}
	log2 := log.New(logFile, "[LOG2]", log.LstdFlags)
	log2.Println("222")

	log3 := log.New(logFile, "警告信息:", log.Ldate|log.Lmicroseconds|log.Llongfile|log.Ltime)
	log3.Println("33333")
}
func SimpleLog() {
	filename := "log.txt"
	logFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	log3 := log.New(logFile, "[INFO]", log.Lshortfile)
	log3.Println("三大王日志")
}
