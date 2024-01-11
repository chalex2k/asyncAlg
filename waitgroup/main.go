package main

import (
	"fmt"
	"sync"
)

type httpPkg struct{}

func (httpPkg) Get(url string) {}

var http httpPkg

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.go.dev/",
	}
	fmt.Println("Начинаем выполнять запросы")
	for _, url := range urls {
		// Увеличиваем WaitGroup счетчик.
		wg.Add(1)
		// Запускаем  goroutine для выполения запроса по URL.
		go func(url string) {
			// Уменьшаем счетчик когда goroutine завершается.
			defer wg.Done()
			// Выполеняем запрос по URL.
			http.Get(url)
			fmt.Println(url)
		}(url)
	}
	// Ожидаем пока все HTTP запросы завершатся.
	wg.Wait()
	fmt.Println("Все запросы выполнены")
}
