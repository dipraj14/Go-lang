package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in goLang")

	content := "This needs to go in a file - git@github.com:dipraj14/Go-lang.git "

	file, err := os.Create("./myLocalGoFile.txt")

	checkErrorNil(err)

	length, err := io.WriteString(file, content) // write content into "myLocalGoFile.txt" file // return length of the the file
	checkErrorNil(err)

	fmt.Println("length is : ", length)
	defer file.Close()

	readFile("./myLocalGoFile.txt")

}
func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkErrorNil(err)

	fmt.Println("Text inside the file is in byte stream: ", dataByte) // this return the file in byte
	fmt.Println()
	fmt.Println("After type casting this byte stream the text was: ", string(dataByte))

}
func checkErrorNil(err error) {
	if err != nil {
		panic(err)
	}
}
