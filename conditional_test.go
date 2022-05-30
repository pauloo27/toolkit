package toolkit_test

import (
	"testing"

	"github.com/Pauloo27/toolkit"
	"github.com/stretchr/testify/assert"
)

func TestConditional(t *testing.T) {
	assert.Equal(t, toolkit.Conditional(true, "hello", "bye"), "hello")
	assert.Equal(t, toolkit.Conditional(false, "hello", "bye"), "bye")
}
