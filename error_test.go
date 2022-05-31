package merrwrap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrapError_Current(t *testing.T) {
	assert := assert.New(t)

	testErr := fmt.Errorf("[err] test 1")

	emptyErr := Error(nil)
	existErr := Error(testErr)

	tests := map[string]struct {
		err    *WrapError
		output error
	}{
		"empty": {err: emptyErr, output: nil},
		"exist": {err: existErr, output: testErr},
	}

	fmt.Println(emptyErr)
	fmt.Println(existErr)

	for _, t := range tests {
		current := t.err.Current()
		assert.Equal(t.output, current)
	}
}
