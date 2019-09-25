package envconfig

import (
	"os"
	"testing"
)

func Test_parseString(t *testing.T) {
	type args struct {
		envName string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name:  "parse string test 1",
			args:  args{envName: "HOST"},
			want:  "host",
			want1: true,
		},
		{
			name:  "parse string test 2",
			args:  args{envName: "HOST2"},
			want:  "",
			want1: false,
		},
	}
	os.Setenv("HOST", "host")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseString(tt.args.envName)
			if got != tt.want {
				t.Errorf("parseString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_parseInt(t *testing.T) {
	type args struct {
		envName string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name:  "test parse int 1",
			args:  args{envName: "PORT"},
			want:  123,
			want1: true,
		},
		{
			name:  "test parse int 2",
			args:  args{envName: "PORT2"},
			want:  0,
			want1: false,
		},
	}
	os.Setenv("PORT", "123")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseInt(tt.args.envName)
			if got != tt.want {
				t.Errorf("parseInt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseInt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_parseInt16(t *testing.T) {
	type args struct {
		envName string
	}
	tests := []struct {
		name  string
		args  args
		want  int16
		want1 bool
	}{
		{
			name:  "test parse int16 1",
			args:  args{envName: "PORT"},
			want:  123,
			want1: true,
		},
		{
			name:  "test parse int16 2",
			args:  args{envName: "PORT2"},
			want:  0,
			want1: false,
		},
	}
	os.Setenv("PORT", "123")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseInt16(tt.args.envName)
			if got != tt.want {
				t.Errorf("parseInt16() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseInt16() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_parseInt32(t *testing.T) {
	type args struct {
		envName string
	}
	tests := []struct {
		name  string
		args  args
		want  int32
		want1 bool
	}{
		{
			name:  "test parse int32 1",
			args:  args{envName: "PORT"},
			want:  123,
			want1: true,
		},
		{
			name:  "test parse int32 2",
			args:  args{envName: "PORT2"},
			want:  0,
			want1: false,
		},
	}
	os.Setenv("PORT", "123")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseInt32(tt.args.envName)
			if got != tt.want {
				t.Errorf("parseInt32() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseInt32() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_parseInt64(t *testing.T) {
	type args struct {
		envName string
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 bool
	}{
		{
			name:  "test parse int64 1",
			args:  args{envName: "PORT"},
			want:  123,
			want1: true,
		},
		{
			name:  "test parse int64 2",
			args:  args{envName: "PORT2"},
			want:  0,
			want1: false,
		},
	}
	os.Setenv("PORT", "123")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseInt64(tt.args.envName)
			if got != tt.want {
				t.Errorf("parseInt64() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseInt64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_parseBool(t *testing.T) {
	type args struct {
		envName string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 bool
	}{
		{
			name:  "test parse bool 1",
			args:  args{envName: "ENABLED"},
			want:  false,
			want1: true,
		},
		{
			name:  "test parse bool 2",
			args:  args{envName: "DISABLED"},
			want:  false,
			want1: false,
		},
	}
	os.Setenv("ENABLED", "false")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseBool(tt.args.envName)
			if got != tt.want {
				t.Errorf("parseBool() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseBool() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
