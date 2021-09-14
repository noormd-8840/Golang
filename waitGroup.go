package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	// "time"
)

// func myFunc(wg *sync.WaitGroup){
// 	time.Sleep(time.Second*1)
// 	fmt.Println("Finished Executing Goroutine")
// 	wg.Done()
// }

var urls = []string{
	"https://www.google.com",
	"https://www.facebook.com",
	"https://www.linkedin.com",
}

func fetchStatus(w http.ResponseWriter, r *http.Request){
	var wg sync.WaitGroup
	for _,url := range urls{
		wg.Add(1)
		go func (url string)  {
			resp, err := http.Get(url)
			if err != nil{
				fmt.Fprintf(w, "%+v\n", err)
			}
			//fmt.Fprintf(w, "%+v\n", resp)
			fmt.Fprintf(w, "%+v\n", resp.Status)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func main(){
	// fmt.Println("Go WaitGroup")	
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go myFunc(&wg)
	// wg.Wait()//block until 0
	// fmt.Println("Finished Executing my go program")


	fmt.Println("Go WaitGroup Example")
	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe(":5500", nil))
}