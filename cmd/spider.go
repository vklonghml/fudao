package main

import (
	"fudao/cmd/options"
	"fudao/pkg/server"
	// "github.com/PuerkitoBio/goquery"
)

func main() {
	opts := options.GetOptions()
	server.StartServer(opts)

}
