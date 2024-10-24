package cmdline

import (
	"fmt"
	"go-tuckshop-manager/services"
	"os"
)

func ConfigurationFlow(args []string) {
	argCount := len(args)
	subCommand := args[1]
	switch subCommand {
	case "add":
		if argCount != 4 {
			fmt.Println("Invalid number of arguments")
			os.Exit(0)
		}
		key := args[2]
		value := args[3]
		configure(key, value)
	case "delete":
		if argCount != 3 {
			fmt.Println("Invalid number of arguments")
			os.Exit(0)
		}
		key := args[2]
		deleteEnv(key)
	}
}

func configure(key string, value string) {
	envMap := services.GetEnvProperties()
	isKeyPresent := false
	for k := range envMap {
		if k == key {
			envMap[k] = value
			isKeyPresent = true
		}
	}
	if !isKeyPresent {
		envMap[key] = value
	}
	services.UpdateEnvFile(envMap)
}

func deleteEnv(key string) {
	envMap := services.GetEnvProperties()
	delete(envMap, key)
	services.UpdateEnvFile(envMap)
}
