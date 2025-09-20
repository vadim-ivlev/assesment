package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(GetHTTPStatusCategory(200)) // Output: Success
	fmt.Println(GetHTTPStatusCategory(404)) // Output: Client Error
	fmt.Println(GetHTTPStatusCategory(500)) // Output: Server Error
}

func GetHTTPStatusCategory(s int) string {
	switch {
	case s >= 100 && s <= 199:
		return "Informational"
	case s >= 200 && s <= 299:
		return "Success"
	case s >= 300 && s <= 399:
		return "Redirection"
	case s >= 400 && s <= 499:
		return "Client Error"
	case s >= 500 && s <= 599:
		return "Server Error"
	default:
		return "Unknown"
	}
}
