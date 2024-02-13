package k8sresmetric

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"gopkg.in/yaml.v3"
)

type ResourceCollector interface {
	RegisterMetric(m MetricsConfig) error
	LabelNames(metrics string) []string
	Update() error
	Values(metric string) (Result, error)
}

// Result returns all the values related to a metric.
type Result struct {
	// Array of all the metric Values.
	Vals []interface{}
	// Array of the labels associated with the metric Value.
	LabelValues [][]string
}

type Collector struct {
	ResourceCollector
	MetricConfigList []MetricsConfig
	lastScrapeTime   time.Time
	nextConsumer     consumer.Metrics
	sync.Mutex
}

var (
	// interval between each metric scrape in seconds.
	scrapeInterval float64 = 25
)

func NewResourceCollector(resType string) ResourceCollector {

	switch resType {
	case "kubernetes":
		return &kMetrics{metricsMap: make(map[string]*MetricsInfo)}
	}

	return nil
}

func (c *Collector) Collect() {

	c.Lock()
	defer c.Unlock()

	// Refresh data only after scrape interval.
	t := time.Now()
	elapsed := t.Sub(c.lastScrapeTime)
	if elapsed.Seconds() >= scrapeInterval {
		c.Update()
		c.lastScrapeTime = time.Now()
	}

	md := pmetric.NewMetrics()

	rs := md.ResourceMetrics().AppendEmpty()

	ms := rs.ScopeMetrics().AppendEmpty().Metrics()
	for _, m := range c.MetricConfigList {

		// Get the value and labels associated with the metric.
		r, err := c.Values(m.Name)
		if err != nil {
			log.Errorf("error resolving metric %v", err)
			continue
		}
		metric := ms.AppendEmpty()

		metric.SetName(m.Name)
		metric.SetDescription(m.Help)
		metric.SetUnit(m.Properties.Unit)
		metric.Sum().SetIsMonotonic(true)

		sum := metric.SetEmptySum()
		dps := sum.DataPoints()
		tim := pcommon.NewTimestampFromTime(t)
		// Range over the result.
		for _, val := range r.Vals {
			var v float64
			// convert the value based on the unit.
			v, err = ConvertUnit(m.Properties.Unit, val)
			if err != nil {
				v = 0.0
			}
			dp := dps.AppendEmpty()
			dp.SetTimestamp(tim)
			dp.SetIntValue(int64(v))

		}
	}
}

func SetCollectors(config string) error {
	exp := &ExporterConfig{}

	err := yaml.Unmarshal([]byte(config), exp)
	if err != nil {
		return err
	}
	resMap := make(map[string]*Collector)
	// Iterate over
	for _, propType := range exp.Objects() {

		// Create instance of the collector
		c := &Collector{ResourceCollector: &kMetrics{metricsMap: make(map[string]*MetricsInfo)}}

		resMap[propType] = c
	}

	// Iterate over all the metrics and register those
	// according to their property Type.
	for _, metric := range exp.Metrics {
		c, ok := resMap[metric.Properties.PropertyType]
		if !ok {
			continue
		}

		c.RegisterMetric(metric)
		c.MetricConfigList = append(c.MetricConfigList, metric)
	}

	return nil
}
