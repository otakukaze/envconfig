package envconfig

import (
	"reflect"
	"testing"
)

func TestExport(t *testing.T) {
	type s struct {
		Name string `env:"NAME"`
	}

	ss := &s{}

	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test export struct tags",
			args: args{i: ss},
			want: []string{"NAME"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Export(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Export() = %v, want %v", got, tt.want)
			}
		})
	}
}
