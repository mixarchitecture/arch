package user

import (
	"github.com/mixarchitecture/arch/auth/delivery/http/dto"
)

type Mapper interface {
	ToRegisterRes(token string) *dto.RegisterResponse
	ToLoginRes(token string) *dto.LoginResponse
}

type mapper struct{}

func New() Mapper {
	return &mapper{}
}

func (m *mapper) ToRegisterRes(token string) *dto.RegisterResponse {
	return &dto.RegisterResponse{
		Token: token,
	}
}

func (m *mapper) ToLoginRes(token string) *dto.LoginResponse {
	return &dto.LoginResponse{
		Token: token,
	}
}
