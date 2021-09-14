package main

import (
	"fmt"
	"log"
	"net/http"
)

type Site struct{
	URL string
	Status int
}

type Result struct{
	URL string
	Status int
}

func Crawl(wId int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs{
		log.Printf("Worker ID: %d\n", wId)
		resp, err := http.Get(site.URL)
		
		if(err != nil){
			log.Println(err.Error())
		}
		results <- Result{
			URL: site.URL,
			Status: resp.StatusCode,
		}
	}
}

func main(){
	fmt.Println("worker pool in Go")

	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	for w := 1; w <= 3; w++{
		go Crawl(w, jobs, results)
	}

	urls := []string{
		"https://www.google.com/",
		"https://www.linkedin.com",
		"https://www.facebook.com",
	}

	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	for a := 1; a<=4; a++{
		result := <-results
		log.Println(result)
	}
}
