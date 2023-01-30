package example

import (
	"reflect"
	"testing"

	"github.com/mixarchitecture/arch/shared/i18n"
)

func TestNewFactory(t *testing.T) {
	type args struct {
		fc FactoryConfig
	}
	tests := []struct {
		name    string
		args    args
		want    Factory
		wantErr bool
	}{
		{
			name: "valid factory config",
			args: args{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			want: Factory{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFactory(tt.args.fc)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFactory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustNewFactory(t *testing.T) {
	type args struct {
		fc FactoryConfig
	}
	tests := []struct {
		name  string
		args  args
		want  Factory
		panic bool
	}{
		{
			name: "valid factory config",
			args: args{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			want: Factory{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			panic: false,
		},
		{
			name: "invalid factory config",
			args: args{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 0,
				},
			},
			want: Factory{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 0,
				},
			},
			panic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.panic {
						t.Errorf("MustNewFactory() panic = %v, want %v", r, tt.panic)
					}
				}
			}()
			if got := MustNewFactory(tt.args.fc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustNewFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactory_Config(t *testing.T) {
	type fields struct {
		fc FactoryConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   FactoryConfig
	}{
		{
			name: "valid factory config",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			want: FactoryConfig{
				MinFieldLength:   1,
				MaxFieldLength:   10,
				MinContentLength: 1,
				MaxContentLength: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Factory{
				fc: tt.fields.fc,
			}
			if got := f.Config(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factory.Config() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactory_IsZero(t *testing.T) {
	type fields struct {
		fc FactoryConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid factory config",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			want: false,
		},
		{
			name: "invalid factory config",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   0,
					MaxFieldLength:   0,
					MinContentLength: 0,
					MaxContentLength: 0,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Factory{
				fc: tt.fields.fc,
			}
			if got := f.IsZero(); got != tt.want {
				t.Errorf("Factory.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactory_NewExample(t *testing.T) {
	type fields struct {
		fc FactoryConfig
	}
	type args struct {
		field   string
		content string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Example
		want1  *i18n.I18nError
	}{
		{
			name: "valid example",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			args: args{
				field:   "field",
				content: "content",
			},
			want: &Example{
				UUID:    "",
				Field:   "field",
				Content: "content",
			},
			want1: nil,
		},
		{
			name: "invalid example with empty field",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			args: args{
				field:   "",
				content: "content",
			},
			want:  nil,
			want1: i18n.NewError("error_example_field_too_short", i18n.P{"Param": 1}),
		},
		{
			name: "invalid example with too long field",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			args: args{
				field:   "12345678901",
				content: "content",
			},
			want:  nil,
			want1: i18n.NewError("error_example_field_too_long", i18n.P{"Param": 10}),
		},
		{
			name: "invalid example with empty content",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			args: args{
				field:   "field",
				content: "",
			},
			want:  nil,
			want1: i18n.NewError("error_example_content_too_short", i18n.P{"Param": 1}),
		},
		{
			name: "invalid example with too long content",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 10,
				},
			},
			args: args{
				field:   "field",
				content: "this is a too long content",
			},
			want:  nil,
			want1: i18n.NewError("error_example_content_too_long", i18n.P{"Param": 10}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Factory{
				fc: tt.fields.fc,
			}
			got, got1 := f.NewExample(tt.args.field, tt.args.content)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factory.NewExample() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Factory.NewExample() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFactory_Unmarshal(t *testing.T) {
	type fields struct {
		fc FactoryConfig
	}
	type args struct {
		uuid    string
		field   string
		content string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Example
		want1  *i18n.I18nError
	}{
		{
			name: "valid example",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			args: args{
				uuid:    "uuid",
				field:   "field",
				content: "content",
			},
			want: &Example{
				UUID:    "uuid",
				Field:   "field",
				Content: "content",
			},
			want1: nil,
		},
		{
			name: "invalid example with too short field",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			args: args{
				uuid:    "uuid",
				field:   "",
				content: "content",
			},
			want:  nil,
			want1: i18n.NewError("error_example_field_too_short", i18n.P{"Param": 1}),
		},
		{
			name: "invalid example with too long field",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 10,
				},
			},
			args: args{
				uuid:    "uuid",
				field:   "this is a too long field",
				content: "content",
			},
			want:  nil,
			want1: i18n.NewError("error_example_field_too_long", i18n.P{"Param": 10}),
		},
		{
			name: "invalid example with too short content",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			args: args{
				uuid:    "uuid",
				field:   "field",
				content: "",
			},
			want:  nil,
			want1: i18n.NewError("error_example_content_too_short", i18n.P{"Param": 1}),
		},
		{
			name: "invalid example with too long content",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 10,
				},
			},
			args: args{
				uuid:    "uuid",
				field:   "field",
				content: "this is a too long content",
			},
			want:  nil,
			want1: i18n.NewError("error_example_content_too_long", i18n.P{"Param": 10}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Factory{
				fc: tt.fields.fc,
			}
			got, got1 := f.Unmarshal(tt.args.uuid, tt.args.field, tt.args.content)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factory.Unmarshal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Factory.Unmarshal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFactory_Validate(t *testing.T) {
	type fields struct {
		fc FactoryConfig
	}
	type args struct {
		e *Example
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *i18n.I18nError
	}{
		{
			name: "valid example",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			args: args{
				e: &Example{
					UUID:    "uuid",
					Field:   "field",
					Content: "content",
				},
			},
			want: nil,
		},
		{
			name: "invalid example with too short field",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			args: args{
				e: &Example{
					UUID:    "uuid",
					Field:   "",
					Content: "content",
				},
			},
			want: i18n.NewError("error_example_field_too_short", i18n.P{"Param": 1}),
		},
		{
			name: "invalid example with too long field",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 10,
				},
			},
			args: args{
				e: &Example{
					UUID:    "uuid",
					Field:   "this is a too long field",
					Content: "content",
				},
			},
			want: i18n.NewError("error_example_field_too_long", i18n.P{"Param": 10}),
		},
		{
			name: "invalid example with too short content",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			args: args{
				e: &Example{
					UUID:    "uuid",
					Field:   "field",
					Content: "",
				},
			},
			want: i18n.NewError("error_example_content_too_short", i18n.P{"Param": 1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Factory{
				fc: tt.fields.fc,
			}
			if got := f.Validate(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factory.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactory_validateField(t *testing.T) {
	type fields struct {
		fc FactoryConfig
	}
	type args struct {
		field string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *i18n.I18nError
	}{
		{
			name: "valid field",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength: 1,
					MaxFieldLength: 10,
				},
			},
			args: args{
				field: "field",
			},
			want: nil,
		},
		{
			name: "invalid field with too short field",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength: 1,
					MaxFieldLength: 10,
				},
			},
			args: args{
				field: "",
			},
			want: i18n.NewError("error_example_field_too_short", i18n.P{"Param": 1}),
		},
		{
			name: "invalid field with too long field",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength: 1,
					MaxFieldLength: 10,
				},
			},
			args: args{
				field: "this is a too long field",
			},
			want: i18n.NewError("error_example_field_too_long", i18n.P{"Param": 10}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Factory{
				fc: tt.fields.fc,
			}
			if got := f.validateField(tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factory.validateField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactory_validateContent(t *testing.T) {
	type fields struct {
		fc FactoryConfig
	}
	type args struct {
		content string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *i18n.I18nError
	}{
		{
			name: "valid content",
			fields: fields{
				fc: FactoryConfig{
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			args: args{
				content: "content",
			},
			want: nil,
		},
		{
			name: "invalid content with too short content",
			fields: fields{
				fc: FactoryConfig{
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			args: args{
				content: "",
			},
			want: i18n.NewError("error_example_content_too_short", i18n.P{"Param": 1}),
		},
		{
			name: "invalid content with too long content",
			fields: fields{
				fc: FactoryConfig{
					MinContentLength: 1,
					MaxContentLength: 10,
				},
			},
			args: args{
				content: "This is a long content",
			},
			want: i18n.NewError("error_example_content_too_long", i18n.P{"Param": 10}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Factory{
				fc: tt.fields.fc,
			}
			if got := f.validateContent(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factory.validateContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactory_NewNotFoundError(t *testing.T) {
	type fields struct {
		fc FactoryConfig
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *i18n.I18nError
	}{
		{
			name: "valid key",
			fields: fields{
				fc: FactoryConfig{
					MinFieldLength:   1,
					MaxFieldLength:   10,
					MinContentLength: 1,
					MaxContentLength: 100,
				},
			},
			args: args{
				key: "key",
			},
			want: i18n.NewError("error_example_not_found", i18n.P{"Param": "key"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Factory{
				fc: tt.fields.fc,
			}
			if got := f.NewNotFoundError(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factory.NewNotFoundError() = %v, want %v", got, tt.want)
			}
		})
	}
}
