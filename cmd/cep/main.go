package main

import (
	"fmt"
	"github.com/joaomannes/GoExpert-Cep/internal/ApiCep"
	"github.com/joaomannes/GoExpert-Cep/internal/Cep"
	"github.com/joaomannes/GoExpert-Cep/internal/ViaCep"
	"os"
	"time"
)

func main() {
	searchers := []Cep.Searcher{
		&ApiCep.ApiCepSearcher{},
		&ViaCep.ViaCepSearcher{},
	}
	cep := os.Args[1]
	ch := make(chan *Cep.CepResult)
	for _, searcher := range searchers {
		go searcher.Search(cep, ch)
	}
	select {
	case result := <-ch:
		fmt.Println(*result)
	case <-time.After(time.Second):
		println("Timeout")
	}
}
