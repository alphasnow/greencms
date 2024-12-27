package utils

import (
	"github.com/sony/sonyflake"
	"time"
)

//type Sid struct {
//	sf *sonyflake.Sonyflake
//}
//
//func NewSid() *Sid {
//	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
//	if sf == nil {
//		panic("sonyflake not created")
//	}
//	return &Sid{sf}
//}
//func (s Sid) GenString() (string, ecode) {
//	id, err := s.sf.NextID()
//	if err != nil {
//		return "", err
//	}
//	return intToBase62(int(id)), nil
//}
//func (s Sid) GenUint64() (uint64, ecode) {
//	return s.sf.NextID()
//}

var sf = sonyflake.NewSonyflake(sonyflake.Settings{
	StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	MachineID: func() (uint16, error) {
		return 0, nil
	},
})

// GenSID
// 用于user id
func GenSID() (uint64, error) {
	// 486045662836703233
	// 18位
	return sf.NextID()
}

//const (
//	base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
//)
//
//func intToBase62(n int) string {
//	if n == 0 {
//		return string(base62[0])
//	}
//
//	var result []byte
//	for n > 0 {
//		result = append(result, base62[n%62])
//		n /= 62
//	}
//
//	// 反转字符串
//	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
//		result[i], result[j] = result[j], result[i]
//	}
//
//	return string(result)
//}
