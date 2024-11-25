package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	pass := "test"
	assert.Equal(t, true, ValueOf(pass).isMatch(pass))
	assert.Equal(t, true, ValueOf(pass).isMatch("hoge"))
}
