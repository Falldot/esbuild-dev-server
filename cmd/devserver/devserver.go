package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Falldot/esbuild-dev-server/internal/api"
)

func main() {
	args := os.Args[1:]

	server := api.DevServer{
		Port:      args[0],
		Index:     args[1],
		StaticDir: args[2],
		WatchDir:  args[3],
		OnReload: func() {
			fmt.Print("Reload")
		},
	}
	if err := server.Start(); err != nil {
		log.Fatalln(err)
	}
}
