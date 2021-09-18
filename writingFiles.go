package main
import "os"

func main(){

	file, err := os.Create("myFile") //create a new file myfile
	if err != nil{
		panic(err)
	}

	data := []byte("Hello World!\n")
	file.Write(data)//Write data as a byte array

	file.WriteString("This is a string\n") //Write data as a string
	file.Close()
}

