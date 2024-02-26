package sf

import (
	"sync"
	"time"
)

const (
	SF_TIMESTAMP_BITS int16 = 41
	SF_SEQUENCE_BITS  int16 = 12
	SF_WORKER_BITS    int16 = 10
	SF_BITS           int16 = 64
	SF_EPOCH          int64 = 1704067200 // 2024-01-01 00:00:00
)

type Snowflake struct {
	lock      sync.Mutex
	timestamp int64
	worker    int16
	sequence  int16
}

var sfIns *Snowflake
var sfOnce sync.Once

func New() int64 {
	sfOnce.Do(func() {
		sfIns = &Snowflake{
			timestamp: -1,
			sequence:  0,
			worker:    0,
		}
	})
	return sfIns.Next()
}

func (s *Snowflake) Next() int64 {
	s.lock.Lock()
	defer s.lock.Unlock()
	timestamp := time.Now().Unix()
	if timestamp == s.timestamp {
		s.sequence = (s.sequence + 1) & ((1 << SF_SEQUENCE_BITS) - 1)
		if s.sequence == 0 {
			for timestamp <= s.timestamp {
				timestamp = s.waitNextMillis()
			}
		}
	} else {
		s.sequence = 0
	}
	s.timestamp = timestamp
	id := ((s.timestamp - SF_EPOCH) << (SF_BITS - SF_TIMESTAMP_BITS)) | (int64(s.worker) << SF_WORKER_BITS) | int64(s.sequence)
	return id
}

func (s *Snowflake) waitNextMillis() int64 {
	t := time.Now().Unix()
	for t <= s.timestamp {
		t = time.Now().Unix()
	}
	return t
}

func Parse(id int64) (int64, int16, int16) {
	sequence := id & ((1 << int64(SF_SEQUENCE_BITS)) - 1)
	worker := (id >> int64(SF_SEQUENCE_BITS)) & ((1 << int64(SF_WORKER_BITS)) - 1)
	timestamp := (id >> int64(SF_BITS-SF_TIMESTAMP_BITS)) + SF_EPOCH
	return timestamp, int16(worker), int16(sequence)
}
