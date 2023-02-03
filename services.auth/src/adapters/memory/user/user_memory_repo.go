package user

import "github.com/mixarchitecture/arch/auth/src/domain/user"

type repo struct {
	userFactory user.Factory
	users       map[string]user.User
}

func New(userFactory user.Factory) user.Repository {
	if userFactory.IsZero() {
		panic("exampleFactory is zero")
	}
	return &repo{
		userFactory: userFactory,
	}
}
