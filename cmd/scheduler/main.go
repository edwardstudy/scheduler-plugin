package scheduler

import (
	"os"

	"github.com/edwardstudy/scheduler-plugin/pkg/capacityscheduling"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	// Set up scheduler with plugin factory
	// 1. plugin factory
	command := app.NewSchedulerCommand(
		app.WithPlugin(capacityscheduling.Name, capacityscheduling.New),
	)

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}

}
