package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("Welcome to Web verb Video")
	PerformPostRequest()
}

func PerformPostRequest() {
	const myUrl = "http://localhost:3000/post"

	// Note: Post request accept the JSON data

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"courseName" : "Let't go with GoLang",
			"price" : 0,
			"paltform" : "learnCodeOnline.in"
		}
	`)
	response, err := http.Post(myUrl, "application/json", requestBody)
	// checkErrorNil(err)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	//checkErrorNil(err)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}

// func checkErrorNil(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
