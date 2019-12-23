package day1partB

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtraFuel(t *testing.T) {

	tests := []struct {
		initialFuel       int
		extraFuelRequired int
	}{
		{
			initialFuel:       2,
			extraFuelRequired: 0,
		},
		{
			initialFuel:       654,
			extraFuelRequired: 312,
		},
	}

	for _, st := range tests {
		t.Run("1", func(t *testing.T) {
			extra := ExtraFuel(st.initialFuel)
			assert.Equal(t, st.extraFuelRequired, extra)
		})
	}
}
