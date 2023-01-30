package memory_example

import "github.com/mixarchitecture/arch/example/src/domain/example"

type exampleRepo struct {
	examples       map[string]example.Example
	exampleFactory example.Factory
}

func NewExampleRepo(exampleFactory example.Factory) example.Repository {
	if exampleFactory.IsZero() {
		panic("exampleFactory is zero")
	}
	return &exampleRepo{
		examples:       make(map[string]example.Example),
		exampleFactory: exampleFactory,
	}
}
