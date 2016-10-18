package snowflake

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// json decode helper for response
type idJsonHelper struct {
	Id string `json:"id"`
}

type handler struct {
	worker *IdWorker
}

func (h *handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	//Get id
	id := h.worker.NextId()

	// compose json data
	jsonHelper := &idJsonHelper{
		Id: strconv.FormatUint(id, 10),
	}

	j, err := json.Marshal(jsonHelper)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(resp, "%s", err)
	}

	// set header
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	fmt.Fprintf(resp, "%s", j)
}
