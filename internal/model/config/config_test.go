package config

import (
	// "fmt"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestReadConfig(t *testing.T) {
	cfg := NewConfig()
	cfg.ReadConfig()
	tests := []struct {
		name string
		want *Config
	}{
		// Add test cases.
		{
			name: "Check for Struct",
			want: cfg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cfg; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfig() = %+v, want %+v\n", got, tt.want)
			}
		})
	}
}