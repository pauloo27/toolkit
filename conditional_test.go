package toolkit_test

import (
	"testing"

	k "github.com/pauloo27/toolkit"
	"github.com/stretchr/testify/assert"
)

func TestIs(t *testing.T) {
	assert.Equal(t, k.Is(true, "hello", "bye"), "hello")
	assert.Equal(t, k.Is(false, "hello", "bye"), "bye")
}
