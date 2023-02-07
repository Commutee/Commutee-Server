package main

import (
	"commutee/handler"
	"commutee/server"
	"fmt"
)

func main() {
	config, err := handler.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	server.Initserver(config)
	// kkk
}
