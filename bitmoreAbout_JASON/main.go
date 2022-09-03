package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"Course Name"`
	Price    int    `json:"Cost"`
	Platform string `json:"Website"`
	Password string
	Tags     []string `json:"tags,omitempty"` // if any value is empty then use omitempty
}

func main() {
	fmt.Println("Discuss more About jason")
	// EncodeJSon()
	DecodeJson()
}

func EncodeJSon() {
	myCorses := []course{
		{"C++", 299, "takeUforword.org", "abc1234", []string{"Data Structure Algo", "Tries"}},
		{"java", 399, "udemy.org", "dipraj__14", []string{"Android Dev.", "Multi Thrading"}},
		{"Golang", 499, "lco.in", "dip_raj14", []string{"backend Devlopment", "Web Api"}},
		{"Angular", 199, "lco.in", "dip_raj15", nil},
	}

	// package this data as JSON data
	//finaljson, err := json.Marshal(myCorses)  // this is the correct way to store the JSON data but this was not clear
	// so instade of this Marshal we are going to use MarshalIndent ---> this clear and looks better

	//finaljson, err := json.MarshalIndent(myCorses, "dipraj", "\t")
	finaljson, err := json.MarshalIndent(myCorses, "", "\t") // its bassically convert the structure data inito json formet

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finaljson)
}

func DecodeJson() {
	jasonDataFromWeb := []byte(`
		{
			"Course Name": "C++",
			"Cost": 299,
			"Website": "takeUforword.org",
			"Password": "abc1234",
			"tags": ["Data Structure Algo","Tries"]
		}
	
	`)

	var myCourses course

	checkvalid := json.Valid(jasonDataFromWeb)

	if checkvalid {
		fmt.Println("Json Data Valid")
		json.Unmarshal(jasonDataFromWeb, &myCourses) // its basically Unmarshal the json data in course structure formet
		fmt.Printf("%#v\n", myCourses)
	} else {
		fmt.Println("Jason data not valid")
	}

	// some Case where ypu want to add data to key valuse pair

	var myData map[string]interface{} // i use interface because the first value / key its a string in JSON but second value / value its can be anythis int,float,string etc

	json.Unmarshal(jasonDataFromWeb, &myData)
	fmt.Printf("%#v\n", myData)

	for k, v := range myData {
		fmt.Printf("Key: %v AND Value %v And Type: %T\n", k, v, v)

	}
}
