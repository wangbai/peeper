package server

import (
	"fmt"
	"io"
	"net/http"
)

type defaultHandler struct {
}

func (d defaultHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// set header
	resp.Header().Set("Content-Type", "text/plain")
    resp.WriteHeader(http.StatusNotFound)
    
	// compose body
	content := fmt.Sprintf("No handler for Request URI:%s", req.RequestURI)

	// set body
	io.WriteString(resp, content)
}
