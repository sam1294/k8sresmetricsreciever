package k8sresmetricsreciever

import (
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

type K8sResMetricsConfig struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`
	ResRef                                  string `mapstructure:"resRef"`
}
