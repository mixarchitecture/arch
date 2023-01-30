package jwt

import (
	"errors"
	"time"
)

type UserClaim struct {
	UUID      string   `json:"uuid"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
	ExpiresIn int64    `json:"expiresIn"`
}

func (c *UserClaim) Valid() error {
	if c.IsExpired() {
		return errors.New("user is expired")
	}
	return nil
}

func (c *UserClaim) Expire() {
	c.ExpiresIn = time.Now().Unix() - 1
}

func (c *UserClaim) SetExpireIn(d time.Duration) {
	c.ExpiresIn = time.Now().Add(d).Unix()
}

func (c *UserClaim) IsExpired() bool {
	return c.ExpiresIn < time.Now().Unix()
}
