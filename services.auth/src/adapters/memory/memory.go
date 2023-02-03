package memory

import (
	memory_user "github.com/mixarchitecture/arch/auth/src/adapters/memory/user"
	"github.com/mixarchitecture/arch/auth/src/domain/user"
)

type Memory interface {
	NewUser(userFactory user.Factory) user.Repository
}

type memory struct{}

func New() Memory {
	return &memory{}
}

func (m *memory) NewUser(userFactory user.Factory) user.Repository {
	return memory_user.New(userFactory)
}
