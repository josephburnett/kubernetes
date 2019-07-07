package podautoscaler

import (
	"time"

	autoscalingv2 "k8s.io/api/autoscaling/v2beta2"
)

// BEGIN INTERFACE

const (
	SkInterfaceVersion  = 1
	SkMetricCpu         = "cpu"
	SkMetricConcurrency = "concurrency"
)

type SkAutoscaler interface {
	Scale(int64) (int32, error)
	Stat(SkStat) error
}

type SkStat interface {
	Time() int64
	Metric() string
	Value() (int32, bool)
	AverageValue() (int32, bool)
	AverageUtilization() (int32, bool)
}

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

// END INTERFACE

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
