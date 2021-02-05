package capacityscheduling

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	corelisters "k8s.io/client-go/listers/core/v1"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

const Name = "CapacityScheduling"

var _ framework.ScorePlugin = &CapacityScheduling{}

type CapacityScheduling struct {
	podLister corelisters.PodLister
}

// New registry CapacityScheduling plugin with configuration and scheduler framework tools.
func New(_ *runtime.Unknown, handle framework.FrameworkHandle) (framework.Plugin, error) {
	// podLister can get pod from scheduler cache
	podLister := handle.SharedInformerFactory().Core().V1().Pods().Lister()

	return &CapacityScheduling{
		podLister: podLister,
	}, nil
}

// Name return plugin name.
func (cs *CapacityScheduling) Name() string {
	return Name
}

// Score
func (cs *CapacityScheduling) Score(ctx context.Context, state *framework.CycleState, p *corev1.Pod, nodeName string) (int64, *framework.Status) {
	// Status:
	// - framework.Success
	// - framework.Error: internal errors
	// - framework.Unschedulable: a pod unschedulable -> preemption
	// - framework.UnschedulableAndUnresolvable: a pod unschedulable -> skip preemption
	//   pkg/scheduler/core/generic_scheduler.go:#L1201-L1203
	// - framework.Wait: should wait for premit
	// - framework.Skip: should skip for binding, then use default Binding action
	return 0, framework.NewStatus(framework.Error, "a message for place holding")
}

// ScoreExtensions return nil to disable ScoreExtensions.
func (cs *CapacityScheduling) ScoreExtensions() framework.ScoreExtensions {
	return nil
}
