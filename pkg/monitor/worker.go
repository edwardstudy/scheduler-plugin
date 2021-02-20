package monitor

// Worker receives metrics count and errors from other components
import (
	"net/http"
	"time"
)

type Monitor interface {
	// Get one item
	Pop() interface{}
	// Store one item
	Push(item interface{})
	// Send all remaining metrics/alerts to remote.
	Send()
}

// buffered channel size would block providers(plugin, provider) to send errors, monitor if it hit capacity.
// Use FIFO to store monitor/metrics.
type FakeWorker struct {
	queue       []string
	monitClient http.Client
}


func NewWorker(args ...string) Monitor {
	return &FakeWorker{}
}

// Pop returns the first item
func (w *FakeWorker) Pop() interface{} {
	return nil
}

// Push append item to the tail
func (w *FakeWorker) Push(item interface{}) {}

func (w *FakeWorker) Send() {
	// Invoke monitClient to push monitor/metrics
}

type Item struct {
	Type      string
	Timestamp time.Time
}
