package main

import (
	"fmt"
	"net/http"	
	"time"
	"os"
	"io"
	"strings"
	"bufio"
	"path"
)

const (
	DELAY = 3
	REPEAT = 3
)

func main() {

	for {
		menu()
		run(read())
	}

}

func menu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func read() int {
	var input int
	fmt.Print("Escolha: ")
	fmt.Scanf("%d", &input)
	return input
}

func readByFile(nameFile string) ([]string, error) {
	file, err := os.Open(path.Join(path.Dir(""), "sites.txt"))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var sites []string
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break;
		} 
	}
	return sites, nil
}

func run(input int) {
	// em Go o comando break é opcional no Switch
	switch input {
	case 1:
		monitoring()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Bye Bye...")
		os.Exit(0)
	default:
		fmt.Println("Comando Inválido")
	}
} 

func monitoring() {
	fmt.Println("\nMonitorando...")
	sites, err := readByFile("sites.txt")
	if err != nil {
		fmt.Println("Oops, ocorreu um erro!", err)
		return
	}
	fmt.Println(len(sites))
	bigger := getBigger(sites)
	for i := 0; i < REPEAT; i++ {
		for _, site := range sites {
			bar(len(bigger), "=")
			info(site)
		}
		if i < REPEAT - 1 {
			bar(len(bigger), "=")
			fmt.Println("Monitorando...")
			time.Sleep(DELAY * time.Second)
		}
	}
	bar(len(bigger), "=")
	fmt.Println("\nFinished.\n")
}

func info(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Oops, ocorreu um erro!", err)
		return
	}
	fmt.Println(site)
	fmt.Println("Status: ", resp.StatusCode)
}

func getBigger(arr []string) string {
	var bigger = arr[0]
	for i := 1; i < len(arr); i++ {
		if(len(arr[i]) > len(bigger)) {
			bigger = arr[i]
		} 
	}
	return bigger
}

func bar(max int, symbol string) {
	for i := 0; i < max; i++ {
		fmt.Print(symbol)
	}
	fmt.Println("")
}