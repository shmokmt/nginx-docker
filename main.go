package main

import (
	"log"
	"net/http"
)

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			_, err := http.Get("http://localhost")
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

}
