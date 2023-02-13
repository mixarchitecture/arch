package mapper

import "github.com/mixarchitecture/arch/auth/src/delivery/http/mapper/user"

type Mapper struct {
	User user.Mapper
}

func New() *Mapper {
	return &Mapper{
		User: user.New(),
	}
}
