package main

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"reflect"
	"bufio"
)

func main() {


	fmt.Println("Opção 1:")
	readByOsOpen("data.txt")
	
	fmt.Println("\nOpção 2:")
	readByIoUtilReadFile("data.txt")

	fmt.Println("\nLendo linhas:")
	readByBufioNewReader("data.txt")

}

func readByOsOpen(nameFile string) {
	file, err := os.Open(nameFile)
	if err != nil {
		fmt.Println("Oops, ocorreu um erro!", err)
		return
	}
	defer file.Close()
	content, _ := ioutil.ReadAll(file)
	fmt.Println(string(content))
	fmt.Printf("Type => %s\n", reflect.TypeOf(content))
}

func readByIoUtilReadFile(nameFile string) {
	file2, err := ioutil.ReadFile(nameFile)
	if err != nil {
		fmt.Println("Oops, ocorreu um erro!", err)
		return
	}
	fmt.Println(string(file2))
	fmt.Printf("Type => %s\n", reflect.TypeOf(file2))
}

func readByBufioNewReader(nameFile string) {
	file, err := os.Open(nameFile)
	if err != nil {
		fmt.Println("Oops, ocorreu um erro!", err)
		return
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		fmt.Print(line)
		if err == io.EOF {
			fmt.Println("")
			break
		}
	}
}