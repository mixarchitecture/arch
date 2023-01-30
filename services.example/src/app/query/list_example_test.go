package query

import (
	"context"
	"reflect"
	"testing"

	"github.com/mixarchitecture/arch/shared/i18n"

	"github.com/mixarchitecture/arch/example/src/domain/example"
	example_mocks "github.com/mixarchitecture/arch/example/src/domain/example/mocks"
)

func Test_listExampleHandler_Handle(t *testing.T) {
	mock := preListTest()
	type fields struct {
		exampleRepo example.Repository
	}
	type args struct {
		ctx   context.Context
		query ListExampleQuery
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ListExampleResult
		wantErr bool
	}{
		{
			name: "should list example",
			fields: fields{
				exampleRepo: mock.repo,
			},
			args: args{
				ctx:   mock.ctx,
				query: ListExampleQuery{Limit: 10, Offset: 0},
			},
			want: ListExampleResult{
				Examples: []*example.Example{},
				Total:    0,
			},
			wantErr: false,
		},
		{
			name: "should return error",
			fields: fields{
				exampleRepo: mock.repo,
			},
			args: args{
				ctx:   mock.ctx,
				query: ListExampleQuery{Limit: 10, Offset: 2},
			},
			want:    ListExampleResult{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := listExampleHandler{
				exampleRepo: tt.fields.exampleRepo,
			}
			got, err := h.Handle(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("listExampleHandler.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listExampleHandler.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}

type listTestMocks struct {
	repo *example_mocks.Repository
	ctx  context.Context
}

func preListTest() listTestMocks {
	ctx := context.Background()

	repo := &example_mocks.Repository{}

	repo.On("List", ctx, 10, 0).Return([]*example.Example{}, 0, nil)

	repo.On("List", ctx, 10, 2).Return([]*example.Example{}, 0, i18n.NewError("test"))

	return listTestMocks{
		repo: repo,
		ctx:  ctx,
	}
}
