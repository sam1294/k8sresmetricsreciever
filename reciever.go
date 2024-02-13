package k8sresmetricsreciever

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	kresmetrics "quark.netapp.io/otel-controller/internal/k8sresourcemetrics"
	"quark.netapp.io/otel-controller/internal/metadata"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

type k8sresmetrics struct {
	config *K8sResMetricsConfig
}

func (r *k8sresmetrics) Start(context.Context, component.Host) error {

	rConfig, err := config.GetConfig()
	if err != nil {
		log.Errorf("error in getting KUBECONFIG %v", err)
		return nil
	}

	err = kresmetrics.SetClients(rConfig)
	if err != nil {
		log.Fatalf("Error building extra clientset: %s", err.Error())
	}

	// err = krmetrics.SetCollectors(supCfg.K8sResourceMetricYaml)
	// if err != nil {
	// 	log.Error("error setting resource to metrics collector", err.Error())
	// }

	return nil
}
func (r *k8sresmetrics) scrape(ctx context.Context) (pmetric.Metrics, error) {
	md := pmetric.NewMetrics()

	return md, nil
}

func newK8sResMetrics(ctx context.Context, params receiver.CreateSettings, cfg *K8sResMetricsConfig, consumer consumer.Metrics) (receiver.Metrics, error) {

	k8s := &k8sresmetrics{
		config: cfg,
	}

	scrp, err := scraperhelper.NewScraper(metadata.Type, k8s.scrape, scraperhelper.WithStart(k8s.Start))
	if err != nil {
		return nil, err
	}

	return scraperhelper.NewScraperControllerReceiver(&k8s.config.ScraperControllerSettings, params, consumer, scraperhelper.AddScraper(scrp))
}
