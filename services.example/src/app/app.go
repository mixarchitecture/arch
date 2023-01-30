package app

import (
	"github.com/mixarchitecture/arch/example/src/app/command"
	"github.com/mixarchitecture/arch/example/src/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateExample command.CreateExampleHandler
	UpdateExample command.UpdateExampleHandler
}

type Queries struct {
	ListExample query.ListExampleHandler
	GetExample  query.GetExampleHandler
}
