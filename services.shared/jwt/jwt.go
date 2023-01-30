package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Jwt struct {
	secret   []byte
	Duration time.Duration
}

type Config struct {
	Secret   string
	Duration time.Duration
}

func New(config Config) *Jwt {
	return &Jwt{
		secret:   []byte(config.Secret),
		Duration: config.Duration,
	}
}

func (j *Jwt) Sign(p *UserClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, p)
	return token.SignedString(j.secret)
}

func (j *Jwt) Parse(t string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(t, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.secret, nil
	})
}

func (j *Jwt) Verify(t string) (bool, error) {
	token, err := j.Parse(t)
	if err != nil {
		return false, err
	}
	if _, ok := token.Claims.(*UserClaim); ok && token.Valid {
		return true, nil
	}
	return false, nil
}

func (j *Jwt) VerifyAndParse(t string) (*UserClaim, error) {
	token, err := j.Parse(t)
	if err != nil {
		return nil, err
	}
	if res, ok := token.Claims.(*UserClaim); ok && token.Valid {
		return res, nil
	}
	return nil, nil
}

func (j *Jwt) Refresh(t string) (string, error) {
	token, err := j.Parse(t)
	if err != nil {
		return "", err
	}
	if res, ok := token.Claims.(*UserClaim); ok && token.Valid {
		res.ExpiresIn = time.Now().Add(j.Duration).Unix()
		return j.Sign(res)
	}
	return "", nil
}

func (j *Jwt) Expire(t string) (string, error) {
	token, err := j.Parse(t)
	if err != nil {
		return "", err
	}
	if res, ok := token.Claims.(*UserClaim); ok && token.Valid {
		res.ExpiresIn = time.Now().Add(-time.Hour).Unix()
		return j.Sign(res)
	}
	return "", nil
}
