package main

import (
	"flag"
	"fmt"
	"go-tuckshop-manager/services"
)

func main() {
	services.InitEnvProperties()
	user, email, phone := parseInputUserConfigFlags()
	fmt.Printf("Running the application for user - %s, email - %s, phone - %s\n", user, email, phone)
	args := flag.Args()
	fmt.Println(args)
	argCount := len(args)

	switch argCount {
	case 0:
		fileString, _ := services.ReadReadmeFile("README.md")
		fmt.Println(fileString)
	default:
		switch args[0] {
		case "get":
			if argCount == 1 {
				fmt.Println("hello")
			}
		case "order":
			fmt.Println("order")
		case "delete":
			if argCount != 2 {
				fmt.Println("Invalid number of arguments")
			}
			services.Delete(args[1])
		case "configure":
			if argCount != 3 {
				fmt.Println("Invalid number of arguments")
			}
			services.Configure(args[1], args[2])
		}
	}
}

func parseInputUserConfigFlags() (string, string, string) {
	envMap := services.GetEnvProperties()
	user := flag.String("user", envMap["user"], "The user value")
	email := flag.String("email", envMap["email"], "The email value")
	phone := flag.String("phone", envMap["phone"], "The phone number")
	flag.Parse()
	fmt.Printf("Running the application for user - %s, email - %s, phone - %s\n", *user, *email, *phone)
	return *user, *email, *phone
}
