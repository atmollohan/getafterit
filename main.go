package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

func CurrentUser(name string) string {
	currentUser, err := user.Current()
	if err != nil {
		log.Println(err)
	}

	username := currentUser.Username

	fmt.Printf("Current user is: %s\n", username)
	return username
}

func Getwd() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	return path
}

func main() {
	CurrentUser("Andrew")
	Getwd()
}