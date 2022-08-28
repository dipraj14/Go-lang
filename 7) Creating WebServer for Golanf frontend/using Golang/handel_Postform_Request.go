package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	fmt.Println("Welcome to Web verb Video")
	PerformPostformRequest()
}

func PerformPostformRequest() {
	const myUrl = "http://localhost:3000/postform"

	data := url.Values{}
	// form Data
	data.Add("firstname", "Dipraj")
	data.Add("lastname", "Daripa")
	data.Add("email", "diprajdaripabnk@gmail.com")

	response, err := http.PostForm(myUrl, data)

	checkErrorNil(err)
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	checkErrorNil(err)

	fmt.Println(string(content))
}

func checkErrorNil(err error) {
	if err != nil {
		panic(err)
	}
}
