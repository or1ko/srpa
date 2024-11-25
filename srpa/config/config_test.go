package config

import (
	"testing"
)

func TestSave(t *testing.T) {
	config := Config{ReverseUrl: "http://example.com"}
	config.Save("config.yaml")
}
