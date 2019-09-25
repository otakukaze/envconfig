package envconfig

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		i interface{}
	}
	type testStruct struct {
		Port    int      `env:"PORT"`
		Host    string   `env:"HOST"`
		Enabled bool     `env:"ENABLED"`
		Keys    []string `env:"KEYS"`
	}
	os.Setenv("PORT", "10230")
	os.Setenv("HOST", "localhost")
	os.Setenv("ENABLED", "true")

	tests := []struct {
		name string
		args args
	}{
		{
			name: "test env config parse",
			args: args{
				i: &testStruct{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Parse(tt.args.i)
		})
	}
}
