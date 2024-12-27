// @author AlphaSnow

package xjwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type UserClaims struct {
	jwt.RegisteredClaims
	Role string `json:"role,omitempty"`
}

func (c UserClaims) GetRole() (string, error) {
	return c.Role, nil
}

func (c UserClaims) ToUserPayload() *UserPayload {
	py := &UserPayload{}
	if c.ID != "" {
		py.RequestID = c.ID
	}
	if c.Role != "" {
		py.UserModel = c.Role
	}
	if c.ExpiresAt != nil {
		py.ExpiresAt = c.ExpiresAt.Time
	}
	if c.Subject != "" {
		id, _ := strconv.ParseUint(c.Subject, 10, 0)
		py.UserID = uint(id)
	}
	return py
}

type UserPayload struct {
	ExpiresAt time.Time `json:"expired_at"`
	UserID    uint      `json:"user_id"`
	RequestID string    `json:"request_id"`
	UserModel string    `json:"user_model"`
}

func (j *UserPayload) ToClaims() UserClaims {
	return UserClaims{RegisteredClaims: jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("%d", j.UserID),
		ID:        j.RequestID,
		ExpiresAt: jwt.NewNumericDate(j.ExpiresAt),
	},
		Role: j.UserModel,
	}
}

type UserPayloadOption func(*UserPayload)

func WithPayloadExpiresAt(expiresAt time.Time) UserPayloadOption {
	return func(payload *UserPayload) {
		payload.ExpiresAt = expiresAt
	}
}
