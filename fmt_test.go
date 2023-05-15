package toolkit_test

import (
	"fmt"
	"testing"

	k "github.com/pauloo27/toolkit"
	"github.com/stretchr/testify/assert"
)

func TestFmt(t *testing.T) {
	assert.Equal(t, k.Fmt("%s: %d", "flamengo", 1981), "flamengo: 1981")
	assert.Equal(t, k.Fmt("%s: %d", "flamengo", 1981), fmt.Sprintf("%s: %d", "flamengo", 1981))
}
