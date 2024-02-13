package k8sresmetricsreciever

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"quark.netapp.io/otel-controller/internal/metadata"
)

func TestConfig(t *testing.T) {
	t.Parallel()

	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))

	assert.Nil(t, err)
	assert.NotNil(t, cm)

	factory := NewFactory()
	cc := factory.CreateDefaultConfig()

	conf, err := cm.Sub(component.NewIDWithName(metadata.Type, "").String())
	assert.Nil(t, err)

	err = component.UnmarshalConfig(conf, cc)
	assert.Nil(t, err)

	res, ok := cc.(*K8sResMetricsConfig)
	assert.True(t, ok)
	assert.Equal(t, "testLoc", res.ResRef)
}
