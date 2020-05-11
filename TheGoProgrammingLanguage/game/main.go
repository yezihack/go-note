package main


import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Help(args []string) int {
	fmt.Println(`
		Commands:
			login <username><level><exp>
			logout <username>
			quit(q)
			help(h)
	`)
	return 0
}
func Quit(args []string) int {
	p := 0
	for p < 100 {
		time.Sleep(time.Millisecond * 200)
		p += 10
	}
	fmt.Println("Quit")
	return 1
}
func Logout(args []string) int {
	fmt.Println("USAGE:logout <username>")
	return 0
}
func Login(args []string) int {
	if len(args) != 4 {
		fmt.Println("USAGE: login <username><level><exp>")
		return 0
	}
	p := 0
	for p < 100 {
		time.Sleep(time.Millisecond * 200)
		p += 10
	}
	fmt.Println("login...")
	return 0
}

func GetCommandHandlers() map[string]func(args []string) int {
	return map[string]func(args []string) int{
		"help":Help,
		"h":Help,
		"quit":Quit,
		"q":Quit,
		"login":Login,
		"logout":Logout,
	}
}
func main() {
	fmt.Println("Start Game")
	rd := bufio.NewReader(os.Stdin)
	Help(nil)
	handlers := GetCommandHandlers()
	for {
		fmt.Print("Command>")
		b, _, _ := rd.ReadLine()
		line := string(b)
		tokens := strings.Split(line, " ")
		if handler, ok := handlers[tokens[0]]; ok {
			ret := handler(tokens)
			if ret != 0 {
				break
			}
		} else {
			fmt.Println("Unknown command:", tokens[0])
			Help(nil)
		}
	}
}
