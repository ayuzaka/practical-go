package chapter13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Add(a, b int) (int, error) {
	return a + b, nil
}

func TestByTestify(t *testing.T) {
	result, err := Add(1, 2)

	assert.Nil(t, err)
	assert.Equal(t, 3, result)
}
