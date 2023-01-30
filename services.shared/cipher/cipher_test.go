package cipher

import (
	"reflect"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func Test_cipher_Hash(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestEncrypt",
			args: args{
				text: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hash(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("cipher.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.args.text) {
				t.Errorf("cipher.Encrypt() = %v, want %v", got, tt.args.text)
			}
		})
	}
}

func Test_cipher_Compare(t *testing.T) {
	type args struct {
		hash []byte
		text string
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte("test"), 9)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestCompare",
			args: args{
				hash: hash,
				text: "test",
			},
			wantErr: false,
		},
		{
			name: "TestCompareError",
			args: args{
				hash: hash,
				text: "test2",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Compare(tt.args.hash, tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("cipher.Compare() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
