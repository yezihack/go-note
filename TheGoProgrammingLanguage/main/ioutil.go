package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var (
	inFile  = flag.String("i", "a.txt", "file name")
	outFile  = flag.String("o", "b.txt", "file name")
)
func main() {
	flag.Parse()
	writeFile()
}
func writeFile() {
	f, err := os.Create(*outFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	lst := []int{1,2,48,10}
	for _, v := range lst {
		_, err = f.WriteString(strconv.Itoa(v) + "\n")
		if err != nil {
			return
		}
	}
}
func readLine() {
	f, err := os.Open(*inFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	br := bufio.NewReader(f)
	values := make([]int, 0)
	for {
		b, isPre, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("err", err)
			return
		}
		if isPre {
			fmt.Println("to long")
			return
		}
		i, err := strconv.Atoi(string(b))
		if err != nil {
			log.Fatal(err)
			return
		}
		values = append(values, i)
	}
	fmt.Println(values)
}
func readAll() {
	flag.Parse()
	b, err := ioutil.ReadFile(*inFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(b))
}
