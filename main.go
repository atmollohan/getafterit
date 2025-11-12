package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
	"github.com/atmollohan/getafterit/internal/utils"
)

func currentUser() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Println(err)
	}

	username := currentUser.Username

	fmt.Printf("Current user is: %s\n", username)
	return username
}

func getWorkingDirectory() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	return path
}

func printEnvVariables() string {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	joinedEnv := strings.Join(os.Environ(), "\n")
	return joinedEnv
}

func useUtils() {
	utils.UtilFunction()
}


func main() {
	command := flag.String("command", "user", "a string")
	flag.Parse()
	commandFlagMap:= map[string]func()string{
		"printenv": printEnvVariables,
		"pwd":  getWorkingDirectory,
		"whoami": currentUser,
	}
	if command != nil {
		if fn, ok := commandFlagMap[*command]; ok && fn != nil {
			fn()
		} else {
			fmt.Println("Unknown command")
		}
	} else {
		fmt.Println("No command provided")
	}
	// for testing utils package
	useUtils()
}
