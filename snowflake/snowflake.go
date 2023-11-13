/*
	struct
	total 64bit. 0~63
	0 符号位，固定0 代表正数 1bit
	1~41	时间戳	41bit
	42~52	机器号	10bit
	52~63	序列号  12bit
*/

package snowflake

import (
	"log"
	"sync"
	"time"
)

type Snowflake struct {
	sync.Mutex
	timestamp    int64
	workerid     int64 // machine id
	datacenterid int64
	sequence     int64
}

const (
	epoch             = int64(1699854639523) // 2023-11-13 13:50:39.523 起始时间戳
	timestampBits     = uint(41)             // 时间错占用位数
	datacenteridBits  = uint(3)
	workeridBits      = uint(7)
	sequenceBits      = uint(12)
	timestampMax      = int64(-1 ^ (-1 << timestampBits))
	datacenteridMax   = int64(-1 ^ (-1 << datacenteridBits))
	workeridMax       = int64(-1 ^ (-1 << workeridBits))
	sequenceMask      = int64(-1 ^ (-1 << sequenceBits))
	workeridShift     = sequenceBits
	datacenteridShift = workeridBits + sequenceBits
	timestampShift    = datacenteridBits + workeridBits + sequenceBits
)

// NextID
//  1. 获取当前的毫秒时间戳；
//  2. 用当前的毫秒时间戳和上次保存的时间戳进行比较；
//     a. 如果和上次保存的时间戳相等，那么对序列号 sequence 加一；
//     b. 如果不相等，那么直接设置 sequence 为 0 即可；
//  3. 然后通过或运算拼接雪花算法需要返回的 int64 返回值。
func (s *Snowflake) NextID() int64 {
	s.Lock()
	now := time.Now().UnixMilli()
	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & sequenceMask
		// 超过sequence最大值时等待下一毫秒
		// 如果当前序列超出12bit长度，则需要等待下一毫秒
		// 下一毫秒将使用sequence:0
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		s.sequence = 0
	}

	t := now - epoch
	if t > timestampMax {
		log.Printf("epoch must be between 0 and %d", timestampMax-1)
		return 0
	}
	s.timestamp = now

	r := int64((t)<<timestampShift | (s.datacenterid << datacenteridShift) | (s.workerid << workeridShift) | s.sequence)
	s.Unlock()
	return r
}
