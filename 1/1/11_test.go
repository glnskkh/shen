package shen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRevFact(t *testing.T) {
	assert.InDelta(t, 1.0, RevFact(0), 0.001)
	assert.InDelta(t, 1.0+1.0, RevFact(1), 0.001)
	assert.InDelta(t, 1.0+1.0+0.5, RevFact(2), 0.001)
	assert.InDelta(t, 1.0+1.0+0.5+(1/6.0), RevFact(3), 0.001)
}

func RevFact(n int) float64 {
	d := 1.0
	sum := 1.0

	// Inv. sum = 1/0! + ... 1/(k - 1)!
	for k := 1; k <= n; k++ {
		d /= float64(k)
		sum += d
	}

	return sum
}
