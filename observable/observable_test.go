package observable_test

import (
	"testing"

	"github.com/basilfx/go-utilities/observable"
	"gopkg.in/stretchr/testify.v1/assert"
)

// Test_New tests the New method.
func Test_New(t *testing.T) {
	o := observable.New()

	assert.NotNil(t, o)
}

// Test_NewWithValue tests the NewWithValue method.
func Test_NewWithValue(t *testing.T) {
	o := observable.NewWithValue([]byte{0xaa})

	assert.NotNil(t, o)

	assert.Equal(t, []byte{0xaa}, o.GetValue())
}
