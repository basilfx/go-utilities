package main

import (
	"context"
	"errors"

	"github.com/basilfx/go-utilities/taskrunner"

	log "github.com/sirupsen/logrus"
)

func main() {
	taskRunner := taskrunner.New()

	taskRunner.Run("TaskA", func(ctx context.Context) {
		select {
		case <-ctx.Done():
			log.Infof("Context cancelled by other task.")
		}
	})
	taskRunner.RunWithError("TaskB", func(ctx context.Context) error {
		return errors.New("cancel")
	})

	log.Infof("Waiting for tasks to complete.")

	taskRunner.Wait()
}
