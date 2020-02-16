package main

import (
    "github.com/tatsuki5820iso/HelloGo/core"
)

// TODO: Set this in .env file
const PORT = ":8000"

func main() {
    server := core.NewServer()

    server.Echo.Logger.Fatal(server.Echo.Start(PORT))
}