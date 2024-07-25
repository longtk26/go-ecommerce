package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	assert.Equal(t, 3, AddOne(1), "AddOne(1) should be 2")
}