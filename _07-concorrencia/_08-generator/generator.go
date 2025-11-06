package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// esta é uma função generator, porque ela encapsula a responsabilidade de gerencia
// tanto o uso das goroutines quanto a criação do channel, que está sendo devolvido para
// que a função que à invocar possa fazer uso das informações enviadas através do channel;
func titulo(urls ...string) <-chan string {
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

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com", "https://www.youtube.com")

	fmt.Println("Primeiros colocados:", <-t1, "|", <-t2)
	fmt.Println("Segundos colocados:", <-t1, "|", <-t2)
}
