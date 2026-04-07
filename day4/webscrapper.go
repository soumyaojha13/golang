package main

import (
	"fmt"
	"net/http"
	"sync"
)

func fetch(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", url)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Success:", url, resp.Status)
}

func main() {
	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}

	wg.Wait()
}
