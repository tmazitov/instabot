package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/winterssy/instabot"
	"golang.org/x/net/proxy"
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

	// Init context
	ctx, cancel := context.WithCancel(context.Background())
	proxyUrl := "127.0.0.1:9150"

	/*
		Proxy activation
	*/

	// Init Dialer that makes SOCKS5 connection
	dialer, err := proxy.SOCKS5("tcp", proxyUrl, nil, proxy.Direct)
	if err != nil {
		panic(err)
	}

	// Init conn to proxy
	dialContext := func(ctx context.Context, network, address string) (net.Conn, error) {
		return dialer.Dial(network, address)
	}

	// Set dial ctx
	bot.Client.Transport = &http.Transport{DialContext: dialContext, DisableKeepAlives: true}
	defer cancel()

	/*
		Main part
	*/

	// Account auth
	_, err = bot.Login(ctx, true)
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
