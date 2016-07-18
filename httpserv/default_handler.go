package httpserv

import (
	"fmt"
	"io"
	"net/http"
)

type defaultHandler struct {
}

func (d defaultHandler) Handle(resp http.ResponseWriter, req *http.Request) {
	// set header
	resp.Header().Set("Content-Type", "text/plain")

	// compose body
	content := fmt.Sprintf("No handler for Request URI:%s", req.RequestURI)

	// set body
	io.WriteString(resp, content)
}
