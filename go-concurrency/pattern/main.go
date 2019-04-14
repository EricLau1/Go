package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

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

func main() {

	t1 := title("https://ericlau1.github.io", "https://google.com")
	t2 := title("https://golang.org", "https://github.com")

	fmt.Println(<-t1, " | ", <-t2)
	fmt.Println(<-t1, " | ", <-t2)
}
