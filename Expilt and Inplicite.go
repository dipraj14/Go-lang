package main

import "fmt"

func main() {
	var geek1 int = 845

	// explicit type conversion
	var geek2 float64 = float64(geek1)

	fmt.Println(geek2)

}

/* TODO:
Type conversion happens when we assign the value of one data type to another.
 Statically typed languages like C/C++, Java, provide the support for Implicit Type Conversion
 but Golang is different, as it doesn’t support the Automatic Type Conversion or Implicit
 Type Conversion even if the data types are compatible

/*

var geek1 int64 = 875

// it will give compile time error as we
// are performing an assignment between
// mixed types i.e. int64 as int type
var geek2 int = geek1

var geek3 int = 100

// it gives compile time error
// as this is invalid operation
// because types are mix i.e.
// int64 and int
var addition = geek1 + geek3
*/
