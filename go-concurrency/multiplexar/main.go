package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func encaminhar(origin <-chan string, target chan string) {
	for {
		target <- <-origin
	}
}

func juntar(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(input1, c)
	go encaminhar(input2, c)
	return c
}

func main() {
	c := juntar(
		title("https://ericlau1.github.io", "https://google.com"),
		title("https://golang.org", "https://github.com"),
	)

	fmt.Println(<-c, " | ", <-c)
	fmt.Println(<-c, " | ", <-c)
}

func title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			body, _ := ioutil.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(body))[1]
		}(url)
	}
	return c
}
