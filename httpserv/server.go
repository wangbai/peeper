package httpserv

import (
	"fmt"
	"net/http"
	"strings"
)

func serve(resp http.ResponseWriter, req *http.Request) {
	//parse command
	commStr := strings.Split(req.RequestURI, "?")[0]
	commStr = strings.TrimLeft(commStr, "/")

	comm := matchCommand(commStr)
	comm.Handle(resp, req)
}

func Start(port uint32) {
	addr := fmt.Sprintf(":%d", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", serve)
	http.ListenAndServe(addr, mux)
}
