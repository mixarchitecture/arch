package rpc

import (
	"github.com/mixarchitecture/arch/shared/genproto/example"

	"github.com/mixarchitecture/arch/example/src/app"
)

type Server struct {
	app app.Application
	example.ExampleServiceServer
}

func New(app app.Application) example.ExampleServiceServer {
	return Server{
		app: app,
	}
}
