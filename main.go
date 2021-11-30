package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/winterssy/instabot"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Get username
	username, exists := os.LookupEnv("LOGIN")
	if !exists {
		return
	}

	// Get password
	password, exists := os.LookupEnv("PASSWORD")
	if !exists {
		return
	}

	// Bot init
	bot := instabot.New(username, password)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Account auth
	_, err := bot.Login(ctx, true)
	if err != nil {
		panic(err)
	}

	// Create new post
	data, err := bot.PostPhoto(ctx, "example.jpg", "Golang please work!", false)
	if err != nil {
		panic(err)
	}

	// Json response
	fmt.Println(data)
}
