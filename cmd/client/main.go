package main

import (
	"fmt"

	"github.com/rachellezhu/chatroom/internal/chatroom"
)

func main() {
	fmt.Println("Starting client from cmd/client...")
	chatroom.StartClient()
}
