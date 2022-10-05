package main

import "fmt"

func main() {
	var uName string = "Premal Upadhyay"
	fmt.Println(uName)
	fmt.Printf("Type of uName is : %T \n", uName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of isLoggedIn is : %T \n", isLoggedIn)

	/*
		Used 8 - bit unsigned integer here but we can use other as well

		uint8  - 8-bit unsigned integer
		uint16 - 16-bit unsigned integer
		uint32 - 32-bit unsigned integer
		uint64 - 64-bit unsigned integer
		uint   - Either 32 bit or 64 bit (unsigned)

		int8  - 8-bit  signed integer
		int16 - 16-bit signed integer
		int32 - 32-bit signed integer
		int64 - 64-bit signed integer
		int   - Either 32 bit or 64 bit (unsigned)
	*/
	var intVariable uint8 = 255
	fmt.Println(intVariable)
	fmt.Printf("Type of intVariable is : %T \n", intVariable)

	/*
		Used 32-bit float here we can use float 64 as well for more precision.
	*/
	var floatVariable float32 = 234.23456754321
	fmt.Println(floatVariable)
	fmt.Printf("Type of floatVariable is : %T \n", floatVariable)
}
