package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func request() {
	resp, err := http.Get("http://localhost")
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	title := "server-2"
	if strings.Contains(string(result), "server-1") {
		title = "server-1"
	}
	log.Println(title)
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go request()
	}
	wg.Wait()
}
