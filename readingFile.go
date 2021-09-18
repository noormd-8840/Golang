package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"bufio"
	"os"
)

func main(){
	//Open the File
	//Method 1: Using Ioutil
	b, err := ioutil.ReadFile("sample.txt")
	if err != nil {
	 log.Print(err)
	}
	fmt.Println(string(b))

	//Method 2: Using os.Open
	f, err := os.Open("sample.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()

	//Parse/Process
	//Create a Scanner
	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil{
		log.Fatal(err)
	}
}
