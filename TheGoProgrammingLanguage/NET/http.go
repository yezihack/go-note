package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	resp, err := http.Get("https://www.sgfoot.com")
	checkError(err)
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	req, err := http.NewRequest("GET", "https://www.sgfoot.com", nil)
	checkError(err)
	req.Header.Add("User-Agent", "GoBook")
	//
	//http.Client{
	//	CheckRedirect:http.RedirectHandler()
	//}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s\n", err.Error())
		os.Exit(1)
	}
}