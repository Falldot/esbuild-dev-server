package api

import (
	"io"
	"log"
	"net/http"

	"github.com/Falldot/esbuild-dev-server/internal/ws"
)

func (s *DevServer) Connect(w http.ResponseWriter, r *http.Request) {
	ws.NewClient(s.hub, w, r)
}

func (s *DevServer) Reload(w http.ResponseWriter, r *http.Request) {
	s.hub.SendReload()
}

func (s *DevServer) Error(w http.ResponseWriter, r *http.Request) {
	message, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	s.hub.SendErrorBytes(message)
}

func (s *DevServer) View(w http.ResponseWriter, r *http.Request) {
	tmpl, err := ws.AddHotReloadScript(s.Index, s.Port)
	if err != nil {
		log.Println(err)
	}
	tmpl.Execute(w, nil)
}
