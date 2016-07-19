package httpserv

import (
	"fmt"
	"net/http"
	"strings"
)

type server struct {
	Port uint32
}

var s *server

func NewServer() *server {
	if s == nil {
		s = &server{}
	}

	return s
}

func (s *server) Start() {
	addr := fmt.Sprintf(":%d", s.Port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", serve)
	http.ListenAndServe(addr, mux)
}

func serve(resp http.ResponseWriter, req *http.Request) {
	//parse command
	commStr := strings.Split(req.RequestURI, "?")[0]
	commStr = strings.TrimLeft(commStr, "/")

	comm := matchCommand(commStr)
	comm.Handle(resp, req)
}
