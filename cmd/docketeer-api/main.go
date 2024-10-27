package main

import (
	"dockeeter-API/internal/server"
	"fmt"
)

func main() {

	server.Initialize()
	fmt.Printf("Listen in port 3000")
}
