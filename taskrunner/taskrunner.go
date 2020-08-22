package taskrunner

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"
)

// TaskRunner data structure.
type TaskRunner struct {
	waitGroup sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
}

// New returns a reference to a new TaskRunner instance.
func New() *TaskRunner {
	t := TaskRunner{}

	t.ctx, t.cancel = context.WithCancel(context.Background())

	return &t
}

// Cancel will instruct the TaskRunner to cancel the tasks.
func (t *TaskRunner) Cancel() {
	t.cancel()
}

// Wait will wait for all tasks to have completed.
func (t *TaskRunner) Wait() {
	t.waitGroup.Wait()
}

// Run a new task. The task will run in a goroutine, and receive a context.
func (t *TaskRunner) Run(name string, fn func(context.Context)) {
	t.waitGroup.Add(1)

	go func() {
		defer func() {
			t.waitGroup.Done()
			log.Debugf("Task '%s' completed.", name)
		}()
		log.Debugf("Starting task '%s'.", name)
		fn(t.ctx)
	}()
}
