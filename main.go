package main

import (
    "github.com/wangbai/peeper/httpserv"
)

func main() {
    httpserv.Start(11111, &httpserv.DiscoveryResource{nil})
}
