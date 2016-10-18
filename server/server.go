package server

import (
	"fmt"
	"net/http"
	"strings"
)

type server struct {
	port int
}

var s *server

func NewServer(p int) *server {
	if s == nil {
		s = &server{
            port: p, 
        }
	}

	return s
}

func (s *server) Start() {
	addr := fmt.Sprintf(":%d", s.port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", serve)
	http.ListenAndServe(addr, mux)
}

func serve(resp http.ResponseWriter, req *http.Request) {
	//parse command
	commStr := strings.Split(req.RequestURI, "?")[0]
	commStr = strings.TrimLeft(commStr, "/")

	comm := matchHandler(commStr)
	comm.ServeHTTP(resp, req)
}
