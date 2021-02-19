package metricsprovider

type fakeProvider struct {
}

func NewFakeProvider() MetricsProvider {
	return &fakeProvider{}
}

func (f *fakeProvider) FetchHostMetrics(hostName string) []Metric {
	return []Metric{}
}

func (f *fakeProvider) FetchAllHostMetrics() map[string][]Metric {
	return map[string][]Metric{}
}
