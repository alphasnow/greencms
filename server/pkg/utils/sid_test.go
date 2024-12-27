package utils

import "testing"

func TestGenSid(t *testing.T) {
	sid, _ := GenSID()
	// 未设start time长度18位
	// 489571939870835099
	// 设置start time 2023 长度 16位
	// 4262640301900187

	t.Log("debug", sid)
}

func TestGenUUID(t *testing.T) {
	sid := GenUUID()

	t.Log("debug", sid)
}
