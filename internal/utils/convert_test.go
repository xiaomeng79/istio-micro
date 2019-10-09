package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	oddsData = []struct {
		o1 float64
		o2 float64
		b  bool
	}{
		{o1: 3.65441100, o2: 3.2312541, b: false},
		{o1: 3.65441100, o2: 3.65331100, b: true},
		{o1: 3.65441100, o2: 2.65331100, b: false},
		{o1: 3.610000, o2: 3.610000, b: true},
	}
)

func TestOddsCompute(t *testing.T) {
	for _, v := range oddsData {
		assert.Equal(t, v.b, OddsCompute(v.o1, v.o2))
	}
}
