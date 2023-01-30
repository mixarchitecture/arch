package example_mapper

import (
	"github.com/mixarchitecture/arch/example/src/app/command"
	"github.com/mixarchitecture/arch/example/src/app/query"
	"github.com/mixarchitecture/arch/example/src/delivery/http/dtos"
	"github.com/mixarchitecture/arch/example/src/domain/example"
)

type Mapper interface {
	ReqToGetExample(d *dtos.GetExampleRequest) query.GetExampleQuery
	GetExampleQueryToRes(d *query.GetExampleResult) dtos.GetExampleResponse
	ReqToListExample(d *dtos.ListExampleRequest) query.ListExampleQuery
	ListExampleQueryToRes(d *query.ListExampleResult) dtos.ListExampleResponse
	ListExampleQueryToResItems(d []*example.Example) []example.Example
	ListExampleQueryToResItem(d *example.Example) example.Example
	ReqToCreateExample(d *dtos.CreateExampleRequest) command.CreateExampleCommand
	ReqToUpdateExample(d *dtos.UpdateExampleRequest) command.UpdateExampleCommand
}

type mapper struct{}

func New() Mapper {
	return &mapper{}
}

func (m *mapper) ReqToGetExample(d *dtos.GetExampleRequest) query.GetExampleQuery {
	return query.GetExampleQuery{
		Field: d.Field,
	}
}

func (m *mapper) GetExampleQueryToRes(d *query.GetExampleResult) dtos.GetExampleResponse {
	return dtos.GetExampleResponse{
		Field:   d.Field,
		UUID:    d.UUID,
		Content: d.Content,
	}
}

func (m *mapper) ReqToListExample(d *dtos.ListExampleRequest) query.ListExampleQuery {
	return query.ListExampleQuery{
		Offset: *d.Offset,
		Limit:  *d.Limit,
	}
}

func (m *mapper) ListExampleQueryToRes(d *query.ListExampleResult) dtos.ListExampleResponse {
	return dtos.ListExampleResponse{
		Total: d.Total,
		Items: m.ListExampleQueryToResItems(d.Examples),
	}
}

func (m *mapper) ListExampleQueryToResItems(d []*example.Example) []example.Example {
	items := make([]example.Example, len(d))
	for i, item := range d {
		items[i] = m.ListExampleQueryToResItem(item)
	}
	return items
}

func (m *mapper) ListExampleQueryToResItem(d *example.Example) example.Example {
	return example.Example{
		Field:   d.Field,
		Content: d.Content,
		UUID:    d.UUID,
	}
}

func (m *mapper) ReqToCreateExample(d *dtos.CreateExampleRequest) command.CreateExampleCommand {
	return command.CreateExampleCommand{
		Field:   d.Field,
		Content: d.Content,
	}
}

func (m *mapper) ReqToUpdateExample(d *dtos.UpdateExampleRequest) command.UpdateExampleCommand {
	return command.UpdateExampleCommand{
		Field:   d.Field,
		Content: d.Content,
	}
}
