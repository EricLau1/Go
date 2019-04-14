package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func main() {

	winner := getFastSite([3]string{
		"https://golang.org", "https://google.com", "https://github.com",
	})

	fmt.Println(winner)
}

func getFastSite(urls [3]string) string {
	c1 := title(urls[0])
	c2 := title(urls[1])
	c3 := title(urls[2])

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam..."
		//default:
		//	return "Sem resposta..."
	}
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
