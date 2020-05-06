package main

import (
	"fmt"
	"net/http"
	"time"
)

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

	c := make(chan string, len(links))
	t := time.Now()

	// for _, link := range links {
	// 	checkStat(link)
	// }

	for _, link := range links {
		go checkStatWithGo(link, c)
	}

	for i := range links {
		fmt.Println(i+1, <-c)
	}

	fmt.Println("end")
	fmt.Println(time.Now().Sub(t))
}

func checkStat(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		return
	}
	fmt.Println(link, "is up")
}

func checkStatWithGo(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		return
	}
	c <- link + " is up"
}
