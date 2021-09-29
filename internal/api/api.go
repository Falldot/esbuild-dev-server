package api

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/Falldot/esbuild-dev-server/internal/ws"
)

type DevServer struct {
	hub *ws.Hub
}

func New() *DevServer {
	hub := ws.NewHub()
	go hub.Run()
	return &DevServer{
		hub: hub,
	}
}

func (ds *DevServer) StartServer(port string, pathToIndex string) error {
	return http.ListenAndServe(port, nil)
}

func (ds *DevServer) Connect(w http.ResponseWriter, r *http.Request) {
	ws.NewClient(ds.hub, w, r)
}

func (ds *DevServer) Reload(w http.ResponseWriter, r *http.Request) {
	ds.SendReload()
}

func (ds *DevServer) Error(w http.ResponseWriter, r *http.Request) {
	message, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	ds.SendErrorBytes(message)
}

func (ds *DevServer) SendError(mes string) {
	message := bytes.TrimSpace([]byte(mes))
	ds.hub.Broadcast <- message
}

func (ds *DevServer) SendErrorBytes(mes []byte) {
	message := bytes.TrimSpace(mes)
	ds.hub.Broadcast <- message
}

func (ds *DevServer) SendReload() {
	message := bytes.TrimSpace([]byte("reload"))
	ds.hub.Broadcast <- message
}
