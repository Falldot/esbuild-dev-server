package api

func (s *DevServer) SendReload() {
	s.hub.SendReload()
}

func (s *DevServer) SendError(message string) {
	s.hub.SendError(message)
}
