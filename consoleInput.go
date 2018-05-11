package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Type something that you want the world to know about: ")
    value, _ := reader.ReadString('\n')
    fmt.Println("User wants to say: ",value)

    fmt.Print("Enter a number : ")
    value , _ = reader.ReadString('\n')
    num, err := strconv.ParseFloat(strings.TrimSpace(value), 64)

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("the value of your number is: ", num)
    }
	// var firstName string
	// var lastName string

	// fmt.Print("Hello please enter your first name: ")
	// fmt.Scanln(&firstName)
	// fmt.Print("Hello please enter your last name: ")
	// fmt.Scanln(&lastName)

    // fmt.Println("Hello ", firstName, lastName, "heres a warm welcome from go")
}
