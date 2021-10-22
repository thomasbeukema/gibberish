package main

import (
	"fmt"
	"internal/api"
	"internal/config"
)

func main() {
	key := config.LoadConfig().ApiKey
	fmt.Println(api.Send(key, "This is a test"))
}
