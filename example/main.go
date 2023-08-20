package main

import (
	"fmt"
	"log"
	"os"

	"github.com/felipeornelis/todoist-go-client"
	"github.com/joho/godotenv"
)

var todoistAuthToken string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	todoistAuthToken = os.Getenv("TODOIST_AUTH_TOKEN")
}

func main() {
	t := todoist.New(todoistAuthToken)

	args := todoist.AddTaskArgs{
		Content: "Ornelis Corp. development tool",
	}
	task, err := t.AddTask(args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(task)
}
