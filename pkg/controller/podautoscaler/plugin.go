package podautoscaler

import (
	"time"

	autoscalingv2 "k8s.io/api/autoscaling/v2beta2"
)

// BEGIN INTERFACE

var SkInterfaceVersion = 1

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

// END INTERFACE

type KubernetesAutoscaler struct {
	controller *HorizontalController
	hpa        *autoscalingv2.HorizontalPodAutoscaler
}

func (ka *KubernetesAutoscaler) Scale(time.Time) (int32, error) {
	// TODO: reconcile hpa
}

func (ka *KubernetesAutoscaler) Record(Stat) error {
	// TODO: record to fake metrics client
}
