package scheduler

import (
	"os"

	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main () {
	// Set up scheduler
	command := app.NewSchedulerCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}

}
