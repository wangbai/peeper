package httpserv

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/wangbai/peeper/config"
)

func init() {
	config.Register("httpserv", &httpservConfig{})
}

type httpservConfig struct {
	Port uint32 `json:"port"`
}

const configFile = "httpserv.conf"

func (hc *httpservConfig) ParseAndBuild(dir string) {
	filePath := dir + "/" + configFile

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("When read ", filePath, " : ", err)
	}

	err = json.Unmarshal(file, hc)
	if err != nil {
		log.Fatal("When parse ", filePath, " : ", err)
	}

	s := NewServer()
	s.Port = hc.Port
}
