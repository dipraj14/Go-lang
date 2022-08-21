package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Println("Welcome to Web verb Video")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:3000/get"

	response, err := http.Get(myUrl)
	checkErrorNil(err)

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length is : ", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)

	checkErrorNil(err)

	fmt.Println(string(content))
}
func checkErrorNil(err error) {
	if err != nil {
		panic(err)
	}
}
