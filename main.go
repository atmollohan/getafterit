package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
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

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, ctr := range containers {
		fmt.Printf("%s %s\n", ctr.ID, ctr.Image)
	}
}
