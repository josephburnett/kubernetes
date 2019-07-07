package podautoscaler

import (
	"time"

	autoscalingv2 "k8s.io/api/autoscaling/v2beta2"
)

// BEGIN INTERFACE

const (
	SkInterfaceVersion = 1

	SkMetricCpu         = "cpu"
	SkMetricConcurrency = "concurrency"

	SkStatePending     = "pending"
	SkStateRunning     = "running"
	SkStateReady       = "ready"
	SkStateTerminating = "terminating"
)

type SkPlugin interface {
	NewAutoscaler(SkEnvironment, string) SkAutoscaler
}

type SkEnvironment interface {
	Pods() []SkPod
}

type SkPod interface {
	Name() string
	State() string
	LastTransistion() int64
	CpuRequest() int32
}

type SkAutoscaler interface {
	Scale(int64) (int32, error)
	Stat(SkStat) error
}

type SkStat interface {
	Time() int64
	PodName() string
	Metric() string
	Value() int32
}

// END INTERFACE

func NewSkAutoscaler(hpa string) SkAutoscaler {
	// TODO: make a bunch of fake stuff
	// TODO: parse hpa string as hpa object
	return &kubernetesAutoscaler{
		controller: NewHorizontalController(
			evtNamespacer,
			scaleNamespacer,
			hpaNamespacer,
			mapper,
			metricsClient,
			hpaInformer,
			podInformer,
			resyncPeriod,
			downscaleStabilisationWindow,
			tolerance,
			cpuInitializationPeriod,
			delayOfInitialReadinessStatus,
		),
		hpa: (*autoscalingv2.HorizontalPodAutoscaler)(nil),
	}
}

type kubernetesAutoscaler struct {
	controller *HorizontalController
	hpa        *autoscalingv2.HorizontalPodAutoscaler
}

var _ SkAutoscaler = (*kubernetesAutoscaler)(nil)

func (ka *kubernetesAutoscaler) Scale(time.Time) (int32, error) {
	// TODO: reconcile hpa
}

func (ka *kubernetesAutoscaler) Record(Stat) error {
	// TODO: record to fake metrics client
}
