package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, u := range urls {
		go func(url string) {
			response, _ := http.Get(url)
			html, _ := ioutil.ReadAll(response.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(u)
	}

	return c
}

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		Titulo("https://www.coder.com.br", "https://www.google.com"),
		Titulo("https://www.amazon.com", "https://www.youtube.com"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
