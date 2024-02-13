package k8sresmetric

import (
	"testing"

	"github.com/spyzhov/ajson"
	"github.com/stretchr/testify/assert"
)

var jsonStr1 string = `
{
	"value": 3,
	"label" : "test"
}
`
var jsonStr2 string = `
{
	"value": 4,
	"label": "test1"
}
`
var resConfig string = `
metrics:
- name: ntaprep_total_transfer_bytes
  help: Total number of protocol write ops
  type: counter
  properties:
    type: kubernetes
    object: NetAppVolumeReplication
    value: $.status.totalTransferBytes
    unit: bytes
    labels:
      sourceUUID: $.spec.sourceVolume.volumeUUID
      destinationUUID: $.spec.destinationVolume.volumeUUID
- name: quark_health_status_etcd
  help: Etcd health Status
  type: gauge
  properties:
    type: kubernetes
    object: Quark
    value: $.status.health.etcdCluster
    labels:
      name: $.metadata.name
      component: etcd
      cluster: $.spec.project.clusterName
- name: quark_health_status_nvc
  help: NetApp Volume Controller health Status
  type: gauge
  properties:
    type: kubernetes
    object: Quark
    value: $.status.health.nvc
    labels:
      name: $.metadata.name
      component: nvc
      cluster: $.spec.project.clusterName
- name: quark_health_status_operator
  help: Quark Operator health Status
  type: gauge
  properties:
    type: kubernetes
    object: Quark
    value: $.status.health.operator
    labels:
      name: $.metadata.name
      component: operator
      cluster: $.spec.project.clusterName
- name: quark_health_status_nodevol
  help: Node Vol health Status
  type: gauge
  properties:
    type: kubernetes
    object: Quark
    value: $.status.health.nodevol
    labels:
      name: $.metadata.name
      component: nodevol
      cluster: $.spec.project.clusterName
- name: quark_health_status_nodemonitor
  help: Node Monitor health Status
  type: gauge
  properties:
    type: kubernetes
    object: Quark
    value: $.status.health.nodeMonitor
    labels:
      name: $.metadata.name
      component: nodeMonitor
      cluster: $.spec.project.clusterName
- name: quark_health_status_nsc
  help: Netapp Snapshot controller health Status
  type: gauge
  properties:
    type: kubernetes
    object: Quark
    value: $.status.health.nsc
    labels:
      name: $.metadata.name
      component: nsc
      cluster: $.spec.project.clusterName
- name: quark_health_status_svcmesh
  help: ServiceMesh health Status
  type: gauge
  properties:
    type: kubernetes
    object: Quark
    value: $.status.health.svcmeshctrl
    labels:
      name: $.metadata.name
      component: servicemesh
      cluster: $.spec.project.clusterName
- name: quark_health_status_nhc
  help: Netapp HA Controller health Status
  type: gauge
  properties:
    type: kubernetes
    object: Quark
    value: $.status.health.nhc
    labels:
      name: $.metadata.name
      component: nhc
      cluster: $.spec.project.clusterName
# - name: ntap_quarketcd_size
#   help: Etcd cluster size. This is for testing purposes.
#   type: gauge
#   properties:
#     type: kubernetes
#     object: Quark
#     value: $.spec.controllerConfig.etcdClusterSize
#     labels:
#       cloud: $.spec.cloud
#       name: $.metadata.name
#       foo: bar
# - name: ntap_pod_foo
#   help: This is an example of fetching metric from list.
#   type: gauge
#   properties:
#     type: kubernetes
#     object: Pod
#     value: $.status.conditions[?(@.type=='Ready')].status
#     labels:
#       name: $.metadata.name
#       foo: bar
`

func TestKMetricsValue(t *testing.T) {
	var rNodes []*ajson.Node
	rNode1, err := ajson.Unmarshal([]byte(jsonStr1))
	assert.Nil(t, err)
	rNode2, err := ajson.Unmarshal([]byte(jsonStr2))
	assert.Nil(t, err)

	rNodes = append(rNodes, rNode1, rNode2)

	mInf := &MetricsInfo{
		Data:      rNodes,
		Path:      "$.value",
		LabelPath: []string{"$.label"},
		LabelKeys: []string{"foo"},
	}
	nMap := make(map[string]*MetricsInfo)
	nMap["metric"] = mInf
	km := &kMetrics{
		metricsMap: nMap,
	}

	r, err := km.Values("metric")
	assert.Equal(t, 1, len(km.LabelNames("metric")))
	assert.Nil(t, err)
	assert.Equal(t, 2, len(r.Vals))
}

func TestSet(t *testing.T) {
	err := SetCollectors(resConfig)
	assert.Nil(t, err)

}
