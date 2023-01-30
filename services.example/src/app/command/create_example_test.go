package command

import (
	"context"
	"testing"

	"github.com/mixarchitecture/arch/shared/events"
	"github.com/mixarchitecture/arch/shared/i18n"

	event_mocks "github.com/mixarchitecture/arch/shared/events/mocks"

	"github.com/mixarchitecture/arch/example/src/config"
	"github.com/mixarchitecture/arch/example/src/domain/example"
	example_mocks "github.com/mixarchitecture/arch/example/src/domain/example/mocks"
)

func Test_createExampleHandler_Handle(t *testing.T) {
	mock := preCreateTest()

	type fields struct {
		exampleRepo   example.Repository
		exampleTopics config.ExampleTopics
		publisher     events.Publisher
	}
	type args struct {
		ctx     context.Context
		command CreateExampleCommand
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should create example",
			fields: fields{
				exampleRepo:   mock.repo,
				exampleTopics: mock.topics,
				publisher:     mock.publisher,
			},
			args: args{
				ctx: mock.ctx,
				command: CreateExampleCommand{
					Field:   "field",
					Content: "content",
				},
			},
			wantErr: false,
		},
		{
			name: "should return error",
			fields: fields{
				exampleRepo:   mock.repo,
				exampleTopics: mock.topics,
				publisher:     mock.publisher,
			},
			args: args{
				ctx: mock.ctx,
				command: CreateExampleCommand{
					Field:   "field2",
					Content: "content2",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := createExampleHandler{
				exampleRepo:   tt.fields.exampleRepo,
				exampleTopics: tt.fields.exampleTopics,
				publisher:     tt.fields.publisher,
			}
			if _, err := h.Handle(tt.args.ctx, tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("createExampleHandler.Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type createTestMocks struct {
	repo      *example_mocks.Repository
	publisher *event_mocks.Publisher
	topics    config.ExampleTopics
	ctx       context.Context
}

func preCreateTest() createTestMocks {
	ctx := context.Background()

	repo := &example_mocks.Repository{}
	publisher := &event_mocks.Publisher{}

	topics := config.ExampleTopics{
		Created: "Example.Created",
	}

	publisher.On("Publish", topics.Created, &example.Example{
		Field:   "field",
		Content: "content",
		UUID:    "",
	}).Return(nil)

	repo.On("Create", context.Background(), &example.Example{
		Field:   "field",
		Content: "content",
		UUID:    "",
	}).Return(nil, nil)

	repo.On("Create", ctx, &example.Example{
		Field:   "field2",
		Content: "content2",
		UUID:    "",
	}).Return(nil, i18n.NewError("error"))

	return createTestMocks{
		repo:      repo,
		publisher: publisher,
		topics:    topics,
		ctx:       ctx,
	}
}
