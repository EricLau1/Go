package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"time"
)

func main() {

	Escrever("log.txt", fmt.Sprintf("%s\n", time.Now()))

	Sobrescrever("test.txt", "Hello world!\n")
	fmt.Println("O arquivo test.txt foi sobreescrito com sucesso!")
}

func Sobrescrever(file, message string) {

	// o ultimo parametro é um Octal referente a permissão do arquivo
	err := ioutil.WriteFile(file, []byte(message), 0644)

	if err != nil {
		log.Fatal(err)
	}

}

func Escrever(filename, text string) {

	if !Existe(filename) {
		
		Criar(filename)
	
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {

		fmt.Println(err)
		
		return
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}

	fmt.Printf("Novo conteudo adicionado ao arquivo %s\r\n", filename)

}

func Criar(filename string) {

	_, err := os.Create(filename)

	if err != nil {

		log.Fatal(err)

	}

	fmt.Printf("arquivo %s criado.\n", filename)
}

func Existe(filename string) bool {

	if _, err := os.Stat(filename); err == nil {
		
		return true
	  
	  } else if os.IsNotExist(err) {
		
		fmt.Println(err)
	
	  }

	  return false
}