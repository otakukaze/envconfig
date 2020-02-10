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
		Port      int       `env:"PORT"`
		Port16    int16     `env:"PORT16"`
		Port32    int32     `env:"PORT32"`
		Port64    int64     `env:"PORT64"`
		F32       float32   `env:"F32"`
		F64       float64   `env:"F64"`
		Host      string    `env:"HOST"`
		Enabled   bool      `env:"ENABLED"`
		Keys      []string  `env:"KEYS"`
		Levels    []int     `env:"LEVELS"`
		SI16      []int16   `env:"SI16"`
		SI32      []int32   `env:"SI32"`
		SI64      []int64   `env:"SI64"`
		SF32      []float32 `env:"SF32"`
		SF64      []float64 `env:"SF64"`
		SBool     []bool    `env:"SBOOL"`
		NoEnv     string
		NoSupport map[string]string `env:"TIME"`
		InStruct  struct {
			Name string `env:"NAME"`
		}
	}
	os.Setenv("NAME", "name")
	os.Setenv("PORT", "10230")
	os.Setenv("PORT16", "10230")
	os.Setenv("PORT32", "10230")
	os.Setenv("PORT64", "10230")
	os.Setenv("F32", "1.23")
	os.Setenv("F64", "1.23")
	os.Setenv("HOST", "localhost")
	os.Setenv("ENABLED", "true")
	os.Setenv("KEYS", "key1,key2")
	os.Setenv("LEVELS", "1,2,3")
	os.Setenv("SI16", "123")
	os.Setenv("SI32", "123")
	os.Setenv("SI64", "123")
	os.Setenv("SF32", "1.23")
	os.Setenv("SF64", "1.23")
	os.Setenv("SBOOL", "false")

	test := &testStruct{}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "test env config parse (struct ptr)",
			args: args{
				i: test,
			},
		},
		{
			name: "test env config parse (ptr in ptr)",
			args: args{
				i: &test,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Parse(tt.args.i)
		})
	}
}
