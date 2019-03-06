package metrics

import (
	"net"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/vrischmann/go-metrics-influxdb"
)

type Metrics struct {
	opts Options
}

//NewMetrics creates a new Metrics
func NewMetrics(options ...Option) *Metrics {
	opts := applyOptions(options...)
	return &Metrics{opts: opts}
}

func (m *Metrics) WithPrefix(s string) string {
	return m.opts.Prefix + "." + s
}

func (m *Metrics) GetRegistry() metrics.Registry {
	return m.opts.Registry
}

func (m *Metrics) MemStats() {
	metrics.RegisterRuntimeMemStats(m.opts.Registry)
	go metrics.CaptureRuntimeMemStats(m.opts.Registry, 5*time.Second)
}

// Log reports metrics into logs.
//
// m.Log(30*time.Second, e.Logger)
//
func (m *Metrics) Log(freq time.Duration, l metrics.Logger) {
	go metrics.Log(m.opts.Registry, freq, l)
}

// Graphite reports metrics into graphite.
//
// 	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:2003")
//  m.Graphite(10e9, "metrics", addr)
//
func (m *Metrics) Graphite(freq time.Duration, prefix string, addr *net.TCPAddr) {
	go metrics.Graphite(m.opts.Registry, freq, prefix, addr)
}

// InfluxDB reports metrics into influxdb.
//
// 	m.InfluxDB(10e9, "http://127.0.0.1:8086","metrics", "test","test"})
//
func (m *Metrics) InfluxDB(freq time.Duration, url, database, username, password string) {
	go influxdb.InfluxDB(m.opts.Registry, freq, url, database, username, password)
}

// InfluxDBWithTags reports metrics into influxdb with tags.
// you can set node info into tags.
//
// 	m.InfluxDBWithTags(10e9, "http://127.0.0.1:8086","metrics", "test","test", map[string]string{"host":"127.0.0.1"})
//
func (m *Metrics) InfluxDBWithTags(freq time.Duration, url, database, username, password string, tags map[string]string) {
	go influxdb.InfluxDBWithTags(m.opts.Registry, freq, url, database, username, password, tags)
}
