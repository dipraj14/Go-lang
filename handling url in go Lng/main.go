package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling URLs in goLang")
	fmt.Println(myUrl)
	// parsing the url

	result, _ := url.Parse(myUrl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query() // qparams hold the all the rawquery like a key, value pair --->
	// "coursename" -> [reactjs]
	// "paymentid"  -> [ghbj456ghb]
	fmt.Printf("The type of query params are : %T\n", qparams) // url.Values --> basiclly its a key,value pair

	fmt.Println(qparams["coursename"])
	// fmt.Println(qparams["paymentid"])

	for _, value := range qparams { // key value pair
		fmt.Println("param is : ", value)
	}

	// construct a URl

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "User=Dipraj",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
