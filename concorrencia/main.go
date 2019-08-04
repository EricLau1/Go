package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	Target  = "https://www.dicionariodenomesproprios.com.br/nomes-masculinos/"
	Regex   = `<a class="lista-nome" href="[^"]+">([^<]+)</a>`
	Samples = `<a class="lista-nome" href="/alexandre/">Alexandre</a><a class="lista-nome" href="/arthur/">Arthur</a>`
)

func main() {
	//requestWithoutGoroutinesInAllPages()
	requestWithGoroutinesInAllPages()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getPage(n int) (string, error) {
	r, err := http.Get(fmt.Sprintf("%s%d", Target, n))
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func getNames(text string) ([]string, error) {
	re, err := regexp.Compile(Regex)
	if err != nil {
		return nil, err
	}

	matches := re.FindAllStringSubmatch(text, -1) // -1 indica que irá retornar todos os matches
	if len(matches) == 0 {
		return nil, errors.New(fmt.Sprintf("Don't matches with: %s", text))
	}
	names := make([]string, 0, 10)
	for _, m := range matches {
		names = append(names, strings.TrimSpace(m[1]))
	}

	return names, nil
}

func getNamesPerPage(n int) ([]string, error) {
	text, err := getPage(n)
	if err != nil {
		return nil, err
	}
	return getNames(text)
}

func requestWithoutGoroutinesInAllPages() {
	timer := time.Now()
	for i := 0; i < 100; i++ {
		page := i + 1
		names, err := getNamesPerPage(page)
		checkErr(err)
		for _, name := range names {
			fmt.Println(name)
		}
		fmt.Println("===================\nPage ", page)
	}
	fmt.Println("Loaded at: ", time.Since(timer))
}

func requestWithGoroutinesInAllPages() {
	timer := time.Now()

	pages := 100
	maxGoroutines := 16

	semaphore := make(chan struct{}, maxGoroutines) // canal com 10 espaços
	for i := 0; i <= pages; i++ {
		page := i + 1
		semaphore <- struct{}{}
		go func(p int) {
			defer func() { <-semaphore }() // liberando o canal
			names, err := getNamesPerPage(p)
			checkErr(err)
			for _, name := range names {
				fmt.Println(name)
			}
			fmt.Printf("\n\nPage: %d\n\n", p)
		}(page)
	}
	for i := 0; i < maxGoroutines; i++ {
		semaphore <- struct{}{}
	}
	fmt.Println("Loaded at: ", time.Since(timer))
}
