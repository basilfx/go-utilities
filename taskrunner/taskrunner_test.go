package taskrunner_test

import (
	"context"
	"testing"

	"github.com/basilfx/go-utilities/taskrunner"

	"gopkg.in/stretchr/testify.v1/assert"
)

// Test_New tests the New method.
func Test_New(t *testing.T) {
	r := taskrunner.New()

	assert.NotNil(t, r)
}

// Test_Run tests the Run in combination with the Wait method.
func Test_Run(t *testing.T) {
	r := taskrunner.New()
	f := false

	r.Run("Name", func(ctx context.Context) {
		f = true
	})
	r.Wait()

	assert.Equal(t, true, f)
}
