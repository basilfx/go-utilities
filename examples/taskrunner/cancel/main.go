package main

import (
	"context"
	"time"

	"github.com/basilfx/go-utilities/taskrunner"

	log "github.com/sirupsen/logrus"
)

func main() {
	taskRunner := taskrunner.New()

	taskRunner.Run("TaskA", func(ctx context.Context) {
		select {
		case <-ctx.Done():
			log.Infof("Context cancelled by other task.")
		case <-time.After(5 * time.Second):
			log.Infof("Will not be printed.")
		}
	})
	taskRunner.RunWithCancel("TaskB", func(ctx context.Context) {
		time.Sleep(3 * time.Second)
	})

	log.Infof("Waiting for tasks to complete.")

	taskRunner.Wait()
}
