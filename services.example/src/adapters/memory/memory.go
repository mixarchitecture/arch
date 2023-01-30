package memory

import (
	memory_example "github.com/mixarchitecture/arch/example/src/adapters/memory/example"
	"github.com/mixarchitecture/arch/example/src/domain/example"
)

type Memory interface {
	NewExample(exampleFactory example.Factory) example.Repository
}

type memory struct{}

func New() Memory {
	return &memory{}
}

func (m *memory) NewExample(exampleFactory example.Factory) example.Repository {
	return memory_example.NewExampleRepo(exampleFactory)
}
