package k8sresmetricsreciever

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"quark.netapp.io/otel-controller/internal/metadata"
)

// NewFactory creates a factory for k8sresmetrics receiver.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, metadata.MetricsStability))
}

func createDefaultConfig() component.Config {
	return &K8sResMetricsConfig{}
}

func createMetricsReceiver(
	ctx context.Context,
	params receiver.CreateSettings,
	cfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	k8sCfg := cfg.(*K8sResMetricsConfig)

	return newK8sResMetrics(ctx, params, k8sCfg, consumer)
}
