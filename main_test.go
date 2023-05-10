package main_test

import (
	"os"
	"testing"

	"github.com/tsubasaogawa/lfm2ytm"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestStr2unixtime(t *testing.T) {
	tests := map[string]struct {
		date     string
		expected int64
	}{
		"jst": {date: "1990-06-14T00:00:00+09:00", expected: 645289200},
		"gmt": {date: "1990-06-14T00:00:00+00:00", expected: 645321600},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			u := main.Str2unixtime(tt.date)
			if tt.expected != u {
				t.Errorf("Expected: %d, Actual: %d", tt.expected, u)
			}
		})
	}
}
