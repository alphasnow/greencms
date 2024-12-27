// @author AlphaSnow

package xjwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserToken struct {
	secret []byte
	expire time.Duration
	model  string
}

func NewUserToken(secret string, expire int, model string) *UserToken {
	return &UserToken{secret: []byte(secret), expire: time.Duration(expire) * time.Second, model: model}
}

func (j *UserToken) GenerateID(id uint, opts ...UserPayloadOption) (string, error) {
	py := &UserPayload{UserID: id, ExpiresAt: time.Now().Add(j.expire), UserModel: j.model}
	return j.Generate(py, opts...)
}

func (j *UserToken) Generate(py *UserPayload, opts ...UserPayloadOption) (string, error) {
	if py.UserModel == "" {
		py.UserModel = j.model
	}
	if py.ExpiresAt.IsZero() {
		py.ExpiresAt = time.Now().Add(j.expire)
	}
	for _, opt := range opts {
		opt(py)
	}
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, py.ToClaims())
	return tk.SignedString(j.secret)
}

func (j *UserToken) ParseID(token string) (uint, error) {
	payload, err := j.Parse(token)
	if err != nil {
		return 0, err
	}
	return payload.UserID, nil
}

func (j *UserToken) Parse(token string) (*UserPayload, error) {
	claims := new(UserClaims)
	tk, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !tk.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	payload := claims.ToUserPayload()
	if payload.UserModel != j.model {
		return nil, fmt.Errorf("invalid model")
	}
	return payload, nil
}
