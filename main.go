package main

import (
	"context"
	"flag"
	"fmt"
	"go-tuckshop-manager/repositories"
	"go-tuckshop-manager/services"
	"go-tuckshop-manager/services/cmdline"
)

func main() {
	services.InitEnvProperties()
	user, email, phone := parseInputUserConfigFlags()
	_, client := repositories.GetConnection()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	fmt.Printf("Running the application for user - %s, email - %s, phone - %s\n", user, email, phone)
	args := flag.Args()
	fmt.Println(args)
	argCount := len(args)

	switch argCount {
	case 0:
		fileString, _ := services.ReadReadmeFile("README.md")
		fmt.Println(fileString)
	default:
		command := args[0]
		switch command {
		case "shop":
			subCommand := args[1]
			switch subCommand {
			case "get":
				fmt.Println(cmdline.GetShop())
			case "refresh":
				cmdline.UpdateShop()
			}
		case "order":
			subCommand := args[1]
			switch subCommand {
			case "history":
			case "place":
			case "setup":
				cmdline.CreateOrderTemplateFlow()
			}
		case "configure":
			cmdline.ConfigurationFlow(args)
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
