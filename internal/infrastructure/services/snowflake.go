package infraServices

import (
	"sync"
	"time"
)

const (
	// SfTimestampBits is used to define timestamp bits
	SfTimestampBits uint16 = 41
	// SfSequenceBits is used to define sequence bits
	SfSequenceBits uint16 = 12
	// SfWorkerBits is used to define worker bits
	SfWorkerBits uint16 = 10
	// SfBits is used to define bits
	SfBits uint16 = 64
	// SfEpoch is used to define epoch
	SfEpoch uint64 = 1704067200 // 2024-01-01 00:00:00
)

// Snowflake struct is used to define snowflake id
type Snowflake struct {
	timestamp uint64
	worker    uint16
	sequence  uint16
}

// Uint64 function is used to convert snowflake id to uint64
func (s *Snowflake) Uint64() uint64 {
	return (s.timestamp-SfEpoch)<<(SfBits-SfTimestampBits) | uint64(s.worker)<<SfSequenceBits | uint64(s.sequence)
}

// Timestamp function is used to get timestamp from snowflake id
func (s *Snowflake) Timestamp() time.Time {
	return time.Unix(int64(s.timestamp), 0)
}

// SfService struct is used to define snowflake service
type SfService struct {
	lock     sync.Mutex
	worker   uint16
	sequence uint16
	lastTs   uint64
}

// NewSfService function is used to create a new snowflake service
func NewSfService(worker uint16) *SfService {
	return &SfService{
		worker:   worker,
		sequence: 0,
		lastTs:   0,
	}
}

// waitNextMillis function is used to wait until next millisecond
func (s *SfService) waitNextMillis() uint64 {
	timestamp := uint64(time.Now().Unix())
	for timestamp <= s.lastTs {
		timestamp = uint64(time.Now().Unix())
	}
	return timestamp
}

// Create function is used to create a new snowflake id
func (s *SfService) Create() *Snowflake {
	s.lock.Lock()
	defer s.lock.Unlock()
	timestamp := uint64(time.Now().Unix())
	if timestamp == s.lastTs {
		s.sequence = (s.sequence + 1) & ((1 << SfSequenceBits) - 1)
		if s.sequence == 0 {
			for timestamp <= s.lastTs {
				timestamp = s.waitNextMillis()
			}
		}
	} else {
		s.sequence = 0
	}
	s.lastTs = timestamp
	return &Snowflake{
		timestamp: s.lastTs,
		sequence:  s.sequence,
		worker:    s.worker,
	}
}

// Extract function is used to extract snowflake id
func (s *SfService) Extract(sf uint64) *Snowflake {
	worker := uint16((sf >> uint64(SfSequenceBits)) & ((1 << uint64(SfWorkerBits)) - 1))
	timestamp := (sf >> (SfBits - SfTimestampBits)) + SfEpoch
	sequence := uint16(sf & ((1 << SfSequenceBits) - 1))
	return &Snowflake{
		sequence:  sequence,
		worker:    worker,
		timestamp: timestamp,
	}
}

// New function is used to create a new snowflake id and return it as uint64
func (s *SfService) New() uint64 {
	return s.Create().Uint64()
}
