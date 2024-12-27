package utils

import "github.com/google/uuid"

// GenUUID
// 用于request id
func GenUUID() string {
	// e0b200e2-f139-4ccf-a4fb-d7b25a31a214
	return uuid.NewString()
}
