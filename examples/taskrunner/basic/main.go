package main

import (
	"context"

	"github.com/basilfx/go-utilities/taskrunner"

	log "github.com/sirupsen/logrus"
)

func main() {
	taskRunner := taskrunner.New()

	taskRunner.Run("Task", func(ctx context.Context) {
		log.Infof("Hello!")
	})

	log.Infof("Waiting for tasks to complete.")

	taskRunner.Wait()
}
