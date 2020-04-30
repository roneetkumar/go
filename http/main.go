package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {

	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
		"https://instagram.com",
		"https://play.google.com",
		"https://dribbble.com",
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
		"https://instagram.com",
		"https://play.google.com",
		"https://dribbble.com",
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
		"https://instagram.com",
		"https://play.google.com",
		"https://dribbble.com",
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
		"https://instagram.com",
		"https://play.google.com",
		"https://dribbble.com",
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
		"https://instagram.com",
		"https://play.google.com",
		"https://dribbble.com",
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
		"https://instagram.com",
		"https://play.google.com",
		"https://dribbble.com",
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
		"https://instagram.com",
		"https://play.google.com",
		"https://dribbble.com",
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
		"https://instagram.com",
		"https://play.google.com",
		"https://dribbble.com",
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
		"https://instagram.com",
		"https://play.google.com",
		"https://dribbble.com",
	}

	for _, link := range links {
		wg.Add(1)
		go checkStat(link)
	}

	wg.Wait()

}

func checkStat(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		return
	}

	fmt.Println(link, "is up")

	defer wg.Done()
}
