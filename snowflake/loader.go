package snowflake

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/wangbai/peeper/server"
)

func init() {
	server.RegisterModule("snowflake", &loader{})
}

type loader struct {
	NodeId uint64 `json:"node_id"`
}

const configFile = "snowflake.conf"

func (l *loader) ParseAndLoad(dir string) {
	filePath := dir + "/" + configFile

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("When read ", filePath, " : ", err)
        return;
	}

	err = json.Unmarshal(file, l)
	if err != nil {
		log.Fatal("When parse ", filePath, " : ", err)
	}

	nid := l.NodeId
	w, err := NewIdWorker(nid)
	if err != nil {
		log.Fatal(err)
	}

	hd := &handler{
		worker: w,
	}

	server.RegisterHandler("snowflake", hd)

	log.Printf("snowflake module has been started")
}
