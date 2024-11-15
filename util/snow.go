package utils

import (
	"errors"
	"sync"
	"time"
)

const (
	workerIDBits       = 8                           // 机器 ID 所占的位数
	sequenceBits       = 8                           // 序列号所占的位数
	maxWorkerID        = -1 ^ (-1 << workerIDBits)   // 最大机器 ID 1023
	maxSequence        = -1 ^ (-1 << sequenceBits)   // 最大序列号 4095
	workerIDShift      = sequenceBits                // 机器 ID 左移位数
	timestampShift     = sequenceBits + workerIDBits // 时间戳左移位数
	startEpoch     int = 1704067200000               // 时间起点（2021-01-01）
)

type Snowflake struct {
	mu        sync.Mutex
	timestamp int
	workerID  int
	sequence  int
}

func NewSnowflake(workerID int) (*Snowflake, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, errors.New("worker ID out of range")
	}
	return &Snowflake{
		workerID: workerID,
	}, nil
}

func (s *Snowflake) Generate() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := int(time.Now().UnixMilli())
	if now == s.timestamp {
		// 同一毫秒内
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			// 序列号用完，等待下一毫秒
			for now <= s.timestamp {
				now = int(time.Now().UnixMilli())
			}
		}
	} else {
		// 新的一毫秒，重置序列号
		s.sequence = 0
	}
	s.timestamp = now

	// 生成唯一 ID
	id := ((now - startEpoch) << timestampShift) |
		(s.workerID << workerIDShift) |
		s.sequence

	return id
}
