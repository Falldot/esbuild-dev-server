package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Falldot/esbuild-dev-server/internal/server"
)

func main() {
	args := os.Args[1:]
	server.SetOptions(server.DevServerOptions{
		Port:      args[0],
		Index:     args[1],
		StaticDir: args[2],
		WatchDir:  args[3],
		OnReload: func() {
			fmt.Print("Reload")
		},
	})
	if err := server.StartDevServer(); err != nil {
		log.Fatalln(err)
	}
}
