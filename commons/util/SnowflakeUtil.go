package util

import (
	"errors"
	"sync"
	"time"
)

const (
	epoch           int64 = 1526285084373
	numWorkerBits         = 10
	numSequenceBits       = 12
	MaxWorkId             = -1 ^ (-1 << numWorkerBits)
	MaxSequence           = -1 ^ (-1 << numSequenceBits)
)

type SnowFlake struct {
	lastTimestamp uint64
	sequence      uint32
	workerId      uint32
	lock          sync.Mutex
}

func (sf *SnowFlake) pack() uint64 {
	uuid := (sf.lastTimestamp << (numWorkerBits + numSequenceBits)) | (uint64(sf.workerId) << numSequenceBits) | (uint64(sf.sequence))
	return uuid
}

// New returns a new snowflake node that can be used to generate snowflake
func New(workerId uint32) (*SnowFlake, error) {
	if workerId < 0 || workerId > MaxWorkId {
		return nil, errors.New("invalid worker Id")
	}
	return &SnowFlake{workerId: workerId}, nil
}

// Generate Next creates and returns a unique snowflake ID
func (sf *SnowFlake) Generate() (uint64, error) {
	sf.lock.Lock()
	defer sf.lock.Unlock()

	ts := timestamp()
	if ts == sf.lastTimestamp {
		sf.sequence = (sf.sequence + 1) & MaxSequence
		if sf.sequence == 0 {
			ts = sf.waitNextMilli(ts)
		}
	} else {
		sf.sequence = 0
	}

	if ts < sf.lastTimestamp {
		return 0, errors.New("invalid system clock")
	}

	sf.lastTimestamp = ts
	return sf.pack(), nil
}

// waitNextMilli if that microsecond is full
// wait for the next microsecond
func (sf *SnowFlake) waitNextMilli(ts uint64) uint64 {
	for ts == sf.lastTimestamp {
		time.Sleep(100 * time.Microsecond)
		ts = timestamp()
	}
	return ts
}

// timestamp
func timestamp() uint64 {
	return uint64(time.Now().UnixNano()/int64(1000000) - epoch)
}
