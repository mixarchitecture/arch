package query

import (
	"context"
	"reflect"
	"testing"

	"github.com/mixarchitecture/arch/shared/i18n"

	"github.com/mixarchitecture/arch/example/src/domain/example"
	example_mocks "github.com/mixarchitecture/arch/example/src/domain/example/mocks"
)

func Test_getExampleHandler_Handle(t *testing.T) {
	mock := preGetTest()

	type fields struct {
		exampleRepo example.Repository
	}
	type args struct {
		ctx   context.Context
		query GetExampleQuery
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    GetExampleResult
		wantErr bool
	}{
		{
			name: "should get example",
			fields: fields{
				exampleRepo: mock.repo,
			},
			args: args{
				ctx:   mock.ctx,
				query: GetExampleQuery{Field: "id"},
			},
			want:    GetExampleResult{},
			wantErr: false,
		},
		{
			name: "should return error",
			fields: fields{
				exampleRepo: mock.repo,
			},
			args: args{
				ctx:   mock.ctx,
				query: GetExampleQuery{Field: "id2"},
			},
			want:    GetExampleResult{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := getExampleHandler{
				exampleRepo: tt.fields.exampleRepo,
			}
			got, err := h.Handle(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("getExampleHandler.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getExampleHandler.Handle() = %v, want %v", got, tt.want)
			}
		})
	}
}

type getTestMocks struct {
	repo *example_mocks.Repository
	ctx  context.Context
}

func preGetTest() getTestMocks {
	ctx := context.Background()

	repo := &example_mocks.Repository{}

	repo.On("Get", ctx, "id").Return(&example.Example{}, nil)

	repo.On("Get", ctx, "id2").Return(&example.Example{}, i18n.NewError("error"))

	return getTestMocks{
		repo: repo,
		ctx:  ctx,
	}
}
