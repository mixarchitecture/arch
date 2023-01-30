package token

import (
	"errors"
	"time"

	"github.com/mixarchitecture/arch/shared/db/redis"
	"github.com/mixarchitecture/arch/shared/jwt"
)

type Service interface {
	Expire(token string) error
	Extend(token string) error
	Generate(u *jwt.UserClaim) (string, error)
	Parse(token string) (*jwt.UserClaim, error)
	Verify(token string) error
}

type SrvConfig struct {
	Secret     string
	Expiration int
	Redis      redis.Service
}

type tokenService struct {
	cnf   *SrvConfig
	jwt   *jwt.Jwt
	redis redis.Service
}

func NewService(cnf *SrvConfig) Service {
	j := jwt.New(jwt.Config{
		Secret:   cnf.Secret,
		Duration: time.Duration(cnf.Expiration) * time.Minute,
	})
	return &tokenService{
		cnf:   cnf,
		redis: cnf.Redis,
		jwt:   j,
	}
}

func (t *tokenService) Expire(token string) error {
	return t.redis.Del(token)
}

func (t *tokenService) Extend(token string) error {
	err := t.Verify(token)
	if err != nil {
		return err
	}
	return t.redis.SetEx(token, "1", t.getExpirationTime())
}

func (t *tokenService) Generate(u *jwt.UserClaim) (string, error) {
	u.SetExpireIn(t.getExpiration() * time.Minute)
	tkn, err := t.jwt.Sign(u)
	if err != nil {
		return "", err
	}
	err = t.setToken(tkn)
	return tkn, err
}

func (t *tokenService) Parse(token string) (*jwt.UserClaim, error) {
	err := t.Verify(token)
	if err != nil {
		return nil, err
	}
	return t.jwt.VerifyAndParse(token)
}

func (t *tokenService) Verify(token string) error {
	ex, err := t.redis.Exist(token)
	if err != nil {
		return err
	}
	if !ex {
		return errors.New("token is expired")
	}
	return nil
}

func (t *tokenService) getExpiration() time.Duration {
	return time.Duration(t.cnf.Expiration)
}

func (t *tokenService) getExpirationTime() int64 {
	return int64(t.cnf.Expiration)
}

func (t *tokenService) setToken(token string) error {
	return t.redis.SetEx(token, "1", t.getExpirationTime())
}
