package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSolutions(t *testing.T) {
	assert.Equal(t, 1, CountSolutionsSlow(1))
	assert.Equal(t, 3, CountSolutionsSlow(2))
	assert.Equal(t, 4, CountSolutionsSlow(4))
	assert.Equal(t, 8, CountSolutionsSlow(6))
}
