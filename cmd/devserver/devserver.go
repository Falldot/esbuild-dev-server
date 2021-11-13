package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Falldot/esbuild-dev-server/internal/api"
)

func main() {

	port := flag.String("p", "", "local server start port")
	idnex := flag.String("i", "", "path to index html file")
	staticDir := flag.String("s", "", "path to static files (js, css, img ...)")
	watchDir := flag.String("w", "", "path to working directory")
	flag.Parse()

	server := api.DevServer{
		Port:      *port,
		Index:     *idnex,
		StaticDir: *staticDir,
		WatchDir:  *watchDir,
		OnReload: func() {
			fmt.Print("Reload")
		},
	}
	if err := server.Start(); err != nil {
		log.Fatalln(err)
	}
}
