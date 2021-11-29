package main

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	auth "github.com/tmazitov/instabot/auth"
	"github.com/tmazitov/instabot/post"
	selenium "sourcegraph.com/sourcegraph/go-selenium"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	ExampleUploadPost()
}

func ExampleUploadPost() error {

	var webDriver selenium.WebDriver
	var err error
	caps := selenium.Capabilities(map[string]interface{}{
		"browserName":     "chrome",
		"mobileEmulation": "{deviceName:Nexus 5}",
	})
	if webDriver, err = selenium.NewRemote(caps, "http://localhost:9515"); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return err
	}
	// GO to page
	err = webDriver.Get("https://www.instagram.com/")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return err
	}

	time.Sleep(time.Second * 3)

	// AUTH
	err = auth.SetAuth(webDriver, time.Second*3)
	if err != nil {
		fmt.Printf("Failed to auth account: %s\n", err)
		return err
	}

	time.Sleep(time.Second * 3)

	// NEW POST
	err = post.CreateNewPost(webDriver, time.Second*3)
	if err != nil {
		fmt.Printf("Failed to create new post: %s\n", err)
		return err
	}

	return nil
}
