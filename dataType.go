// Numbers
// Booleans
// Strings

// Data Type

// Description

// int8	8-bit signed integer
// int16	16-bit signed integer
// int32	32-bit signed integer
// int64	64-bit signed integer
// uint8	8-bit unsigned integer
// uint16	16-bit unsigned integer
// uint32	32-bit unsigned integer
// uint64	64-bit unsigned integer
// int	Both int and uint contain same size, either 32 or 64 bit.
// uint	Both int and uint contain same size, either 32 or 64 bit.
// rune	It is a synonym of int32 and also represent Unicode code points.
// byte	It is a synonym of uint8.
// uintptr	It is an unsigned integer type. Its width is not defined, but its can hold all the bits of a pointer value.

/*
package main

import "fmt"

func main() {

	// Using 8-bit unsigned int
	var X uint8 = 225
	fmt.Println(X, X-3) // 225 222

	// Using 16-bit signed int
	var Y int16 = 32767   // this is the max range of
	fmt.Println(Y+2, Y-2) // so Y + 2 is going to OverFlow and Go to the INT_MIN
}
*/

// Floating-Point Numbers: In Go language, floating-point numbers are divided into two categories as shown in the below table:
// Data Type
// Description

// float32	32-bit IEEE 754 floating-point number
// float64	64-bit IEEE 754 floating-point number

/*
package main

import "fmt"

func main() {
	a := 20.45
	b := 34.89

	// Subtraction of two
	// floating-point number
	c := b - a

	// Display the result
	fmt.Printf("Result is: %f", c)

	// Display the type of c variable
	fmt.Printf("\nThe type of c is : %T", c)

	fmt.Printf("\nNumber: %.2f", 14.9870)   // print the 2  decimal point
	fmt.Printf("\nPadding %9q", "Dipraj")   // padding 9 spaces from left
	fmt.Printf("\n %-9q Padding", "Dipraj") // padding 9 spaces from right
	fmt.printf("Number: %07d ", 45) // fill 5 zeros in front of 45 because = 5 + 45 is 2 dig = 7

	// /t for tab and /T for check the type of the varible
}
*/

// NOt Requirment

// Data Type
// Description

// complex64	Complex numbers which contain float32 as a real and imaginary component.
// complex128	Complex numbers which contain float64 as a real and imaginary component.

//Booleans
// The boolean data type represents only one bit of information either true or false. The values of type boolean are not converted implicitly or explicitly to any other type.
package main

import "fmt"

func main() {

	// variables

	var temp bool
	fmt.Println(temp) //bydefault its false and any integer varible by Default 0

	str1 := "GeeksforGeeks"
	str2 := "geeksForgeeks"
	str3 := "GeeksforGeeks"
	result1 := str1 == str2
	result2 := str1 == str3

	// Display the result
	fmt.Println(result1)
	fmt.Println(result2)

	// Display the type of
	// result1 and result2
	fmt.Printf("The type of result1 is %T and "+
		"the type of result2 is %T",
		result1, result2)

	fmt.Print("\n")

	var tempString string = fmt.Sprintf("The type of result1 is %T and "+
		"the type of result2 is %T",
		result1, result2) //  the whole things we going to use as a string using Sprint

	fmt.Println(tempString)

}
