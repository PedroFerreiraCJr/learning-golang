package main

import (
	"fmt"
	"time"
)

func rotina(channel chan int) {
	fmt.Println("executando a função 'rotina'...")
	channel <- 1
	fmt.Println("o valor '1' foi enviado através do canal...")
	channel <- 2
	fmt.Println("o valor '2' foi enviado através do canal...")
	channel <- 3
	fmt.Println("o valor '3' foi enviado através do canal...")
	channel <- 4
	fmt.Println("o valor '4' foi enviado através do canal...")
	channel <- 5
	fmt.Println("o valor '5' foi enviado através do canal...")
	channel <- 6
	fmt.Println("o valor '6' foi enviado através do canal...")
}

func main() {
	channel := make(chan int, 3)
	go rotina(channel)

	fmt.Println("executando a Goroutine 'main'....")
	time.Sleep(time.Second)

	fmt.Printf("o valor '%d' foi lido do canal...\n", <-channel)
	time.Sleep(time.Second)
}
