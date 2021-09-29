package server

import (
	"log"
	"net/http"
	"os"

	"github.com/Falldot/esbuild-dev-server/internal/api"
	"github.com/Falldot/esbuild-dev-server/internal/watcher"
	"github.com/Falldot/esbuild-dev-server/internal/ws"
)

type DevServerOptions struct {
	Port      string
	Index     string
	StaticDir string
	WatchDir  string
	OnReload  func()
}

var (
	server  *api.DevServer
	config  DevServerOptions
	IsStart bool = false
)

func SetOptions(c DevServerOptions) {
	config = c
}

func SetError(message string) {
	server.SendError(message)
}

func Reload() {
	server.SendReload()
}

func StartDevServer() error {
	log.Println("Starting reload server.")

	server = api.New()

	files, err := os.ReadDir(config.StaticDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			path := "/" + file.Name() + "/"
			http.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(config.StaticDir+path))))
		}
	}

	tmpl, err := ws.AddHotReloadScript(config.Index, config.Port)
	if err != nil {
		return err
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/reload", server.Reload)

	http.HandleFunc("/connect", server.Connect)

	http.HandleFunc("/error", server.Error)

	go watcher.Watch(config.WatchDir, config.OnReload)

	log.Println("Reload server started.")
	log.Println("Reload server listening at", config.Port)

	IsStart = true

	return server.StartServer(config.Port, config.Index)
}
