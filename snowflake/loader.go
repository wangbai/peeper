package snowflake

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/wangbai/peeper/server"
)

func init() {
	server.RegisterModule("snowflake", &Loader{})
}

type Loader struct {
	nodeId uint64 `json:"node_id"`
}

const configFile = "snowflake.conf"

func (l *Loader) ParseAndLoad(dir string) {
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

	nid := l.nodeId
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
