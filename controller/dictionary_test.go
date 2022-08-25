package controller

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"testing"
)

func TestNewDictHandler(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		j := i + 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("start request: ", j)
			defer fmt.Println("done request: ", j)
			url := "http://localhost:8080/api/dictionary"
			if j == 2 {
				url = url + "?not_cache=true&level=n2"
			}
			res, err := http.Get(url)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Printf("%+v", res.Body)
		}()
	}
	wg.Wait()
	fmt.Println("end")
}
