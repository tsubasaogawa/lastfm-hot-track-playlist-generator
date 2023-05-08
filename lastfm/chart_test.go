package lastfm

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestFetch(t *testing.T) {
	var c WeeklyTrackChart
	ENDPOINT_BASE = "http://xxx" // Overwrite endpoint

	t.Run("Fail with invalid endpoint", func(t *testing.T) {
		err := c.Fetch("testuser", "testkey", 123, 456, 1)
		if err == nil {
			t.Errorf("Expected: error with no such host, Actual: no error")
		}
	})
}
