package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func SendValue(c chan string) {
	fmt.Println("Executing Goroutine")
	time.Sleep(time.Second)
	c <- "Hello world"
	fmt.Println("Finished executing goroutine")
}

func call() {
	fmt.Println("Go channels tutorial")

	values := make(chan string, 2)
	defer close(values)

	go SendValue(values)
	go SendValue(values)

	value := <-values
	fmt.Println(value)

	time.Sleep(time.Second)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

func call1() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func myFunc(wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	fmt.Println("Finished executing goroutines")
	wg.Done()
}

func call2() {
	fmt.Println("Go waitgroup tutorial")
	var wg sync.WaitGroup
	wg.Add(1)
	go myFunc(&wg)
	wg.Wait() // blocks until 0
	fmt.Println("Finished executing my go program")
}

func fetchStatus(w http.ResponseWriter, r *http.Request) {
	var urls = []string{
		"https://google.com",
		"https://tutorialedge.net",
		"https://twitter.com",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w, "%+v\n", err)
			}
			fmt.Fprintf(w, "%+v\n", resp)
			wg.Done()
		}(url)
	}

	wg.Wait()
}

func call3() {
	fmt.Println("Go waitgroup tutorial")
	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	// call()
	// call1()
	// call2()
	call3()
}
