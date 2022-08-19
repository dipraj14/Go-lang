package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev" // any url give

func main() {
	fmt.Println("LCO webSite web request")

	response, err := http.Get(url) // this is type of a pointer return meanse original data of this webSite ... you can manipulate it and changes occour in original data in website
	checkErrorNil(err)

	fmt.Printf("Response is of type: %T\n", response) // response was a pointer ... so its return [*http.Response]

	defer response.Body.Close() // caller's responsibility to close the Connection

	dataByte, err := ioutil.ReadAll(response.Body)

	checkErrorNil(err)

	content := string(dataByte)
	fmt.Println("the content inside the wesite was : ", content) // print the web Front end like if content was written in css or html then return taht one
}
func checkErrorNil(err error) {
	if err != nil {
		panic(err)
	}
}
