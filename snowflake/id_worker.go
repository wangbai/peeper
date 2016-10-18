package snowflake

import (
	"fmt"
	"sync"
	"time"
)

const (
	nodeIdBits     = 10
	sequenceIdBits = 12

	nodeIdShift    = sequenceIdBits
	timestampShift = sequenceIdBits + nodeIdBits

	maxNodeId    uint64 = 1<<nodeIdBits - 1
	maxSequencId uint64 = 1<<sequenceIdBits - 1

	//use our own epoch
	epoch uint64 = 1476770314256
)

type IdWorker struct {
	sync.Mutex
	timestamp  uint64 //millisecond
	nodeId     uint64
	sequenceId uint64
}

func NewIdWorker(nid uint64) (*IdWorker, error) {
	if nid < 0 || nid > maxNodeId {
		return nil, fmt.Errorf("node id %d is not in the range between %d and %d", nid, 0, maxNodeId)
	}

	return &IdWorker{
		timestamp:  0,
		nodeId:     nid,
		sequenceId: 0,
	}, nil
}

func (w *IdWorker) NextId() uint64 {
	w.Lock()
	defer w.Unlock()

	//get millisecond
	now := uint64(time.Now().UnixNano() / 1000000)

	if w.timestamp == now {
		w.sequenceId = (w.sequenceId + 1) & maxSequencId

		//the squence used up
		if w.sequenceId == 0 {
			for now <= w.timestamp {
				now = uint64(time.Now().UnixNano() / 1000000)
			}
		}
	} else {
		w.sequenceId = 0
	}

	w.timestamp = now

	//compute unique id
	nextId := ((w.timestamp - epoch) << timestampShift) | (w.nodeId << nodeIdShift) | w.sequenceId

	return nextId
}
