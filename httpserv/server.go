package httpserv

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Resource interface {
	ContentType() string
	HandleCommand(string) string
}

func serveResource(r Resource) func(http.ResponseWriter, *http.Request) {
	if r == nil {
		log.Fatal("There must be a discovery to serve")
	}

	return func(resp http.ResponseWriter, req *http.Request) {
		//parse command
		comm := strings.Split(req.RequestURI, "?")[0]
		comm = strings.TrimLeft(comm, "/")

		// set header
		resp.Header().Set("Content-Type", r.ContentType())
		// set Boday
		io.WriteString(resp, r.HandleCommand(comm))
	}
}

func Start(port uint32, resource Resource) {
	addr := fmt.Sprintf(":%d", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", serveResource(resource))

	http.ListenAndServe(addr, mux)
}
