package k8sresmetricsreciever

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

func TestNewReceiver(t *testing.T) {
	k8s := &K8sResMetricsConfig{
		ScraperControllerSettings: scraperhelper.ScraperControllerSettings{
			CollectionInterval: 1 * time.Second,
		},
	}

	consumer := consumertest.NewNop()

	rm, err := newK8sResMetrics(context.Background(), receivertest.NewNopCreateSettings(), k8s, consumer)
	assert.Nil(t, err)
	assert.NotNil(t, rm)
}

func TestReceiverStart(t *testing.T) {
	k8s := &K8sResMetricsConfig{
		ScraperControllerSettings: scraperhelper.ScraperControllerSettings{
			CollectionInterval: 100 * time.Millisecond,
		},
	}

	consumer := consumertest.NewNop()

	rm, err := newK8sResMetrics(context.Background(), receivertest.NewNopCreateSettings(), k8s, consumer)
	assert.Nil(t, err)
	assert.NotNil(t, rm)

	err = rm.Start(context.Background(), componenttest.NewNopHost())

	time.Sleep(2 * time.Minute)
}
