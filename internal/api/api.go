package api

import (
	"log"
	"net/http"
	"os"

	"github.com/Falldot/esbuild-dev-server/internal/watcher"
	"github.com/Falldot/esbuild-dev-server/internal/ws"
)

type DevServer struct {
	Port      string
	Index     string
	StaticDir string
	WatchDir  string
	OnReload  func()

	IsStart bool

	hub *ws.Hub
}

func (s *DevServer) Start() error {
	log.Println("Starting dev server.")
	if s.Port == "" {
		s.Port = "8080"
	}
	s.Port = ":" + s.Port

	s.hub = ws.NewHub()
	go s.hub.Run()

	if s.StaticDir != "" {
		files, err := os.ReadDir(s.StaticDir)
		if err != nil {
			return err
		}
		for _, file := range files {
			if file.IsDir() {
				path := "/" + file.Name() + "/"
				http.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(s.StaticDir+path))))
			}
		}
	}

	if s.Index == "" {
		http.HandleFunc("/", s.ViewDefault)
	} else {
		http.HandleFunc("/", s.View)
	}
	http.HandleFunc("/reload", s.Reload)
	http.HandleFunc("/connect", s.Connect)
	http.HandleFunc("/error", s.Error)

	if s.WatchDir == "" {
		s.WatchDir = "src"
	}
	go watcher.Watch(s.WatchDir, s.OnReload)

	log.Println("Dev server started.")
	log.Println("http://localhost" + s.Port + "/")

	s.IsStart = true
	return http.ListenAndServe(s.Port, nil)
}
