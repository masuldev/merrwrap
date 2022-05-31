package merrwrap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrapError_Super(t *testing.T) {
	assert := assert.New(t)

	testErr := fmt.Errorf("[super] test 1")

	nilErr := Error(nil)
	exampleErr := Error(testErr)

	tests := map[string]struct {
		err    *WrapError
		output error
	}{
		"nil":     {err: nilErr, output: nil},
		"example": {err: exampleErr, output: testErr},
	}

	fmt.Println(nilErr)
	fmt.Println(exampleErr)

	for _, t := range tests {
		current := t.err.Super()
		assert.Equal(t.output, current)
	}
}

func TestWrapError_Origin(t *testing.T) {
	assert := assert.New(t)

	originErr := fmt.Errorf("[origin] test 1")
	superErr := fmt.Errorf("[origin] test 2")

	nilErr := Error(nil)
	exampleErr := Error(originErr)

	tests := map[string]struct {
		err    *WrapError
		output error
	}{
		"nil":     {err: nilErr, output: nil},
		"example": {err: exampleErr.Wrap(superErr), output: exampleErr},
	}

	for _, t := range tests {
		origin := t.err.Origin()
		assert.Equal(t.output, origin)
	}
}
