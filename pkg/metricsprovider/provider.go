package metricsprovider

type MetricsProvider interface {
	FetchHostMetrics(hostName string) []Metric
	FetchAllHostMetrics() map[string][]Metric
}

type Metric struct {
	Name  string     `json:"name"`  // Name of metric at the provider
	Type  MetricType `json:"type"`  // Type of metric like CPU or Memory
	Value float64    `json:"value"` // Value of metric
}

type MetricType string

const (
	CPUType    MetricType = "CPU"
	MemoryType MetricType = "Memory"
)
