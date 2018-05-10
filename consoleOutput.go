package main

import "fmt"

func main() {
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."

	aNumber := 42
	isTrue := true 

	stringLength, _ := fmt.Println(str1, str2, str3)
	fmt.Println("the strings length: ", stringLength)
	fmt.Println("the aNumbers value: ", aNumber)
	fmt.Println("the isTrue value: ", isTrue)
	fmt.Printf("the aNumbers value as a float: %.2f\n", float64(aNumber))
	fmt.Printf("the data values of our variables %T %T %T %T and %T \n", str1, str2, str3, aNumber, isTrue)
}
