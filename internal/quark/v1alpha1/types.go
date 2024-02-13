/*
 *  Copyright (c) 2019 NetApp
 *  All rights reserved
 */

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppVolume is a specification for a NetApp Volume resource
type NetAppVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetAppVolumeSpec   `json:"spec"`
	Status NetAppVolumeStatus `json:"status"`
}

type DataProtectionType bool

const (
	DataProtectionEnabled  DataProtectionType = true
	DataProtectionDisabled DataProtectionType = false
)

// NetAppVolumeSpec is the spec for a NetAppVolume resource
type NetAppVolumeSpec struct {
	Size                  resource.Quantity  `json:"size"`
	DisplayName           string             `json:"displayName"`
	VolumePath            string             `json:"volumePath"`
	ServiceLevel          string             `json:"serviceLevel"`
	QuarkVersion          string             `json:"quarkVersion"`
	Protocol              string             `json:"protocol"`
	ExportPolicy          string             `json:"exportPolicy,omitempty"`
	SnapshotPolicy        string             `json:"snapshotPolicy"`
	CloudSnapshot         string             `json:"cloudSnapshot,omitempty"`
	CloneParentVolumeUUID string             `json:"cloneParentVolumeUUID,omitempty"`
	CloneParentSnapshot   string             `json:"cloneParentSnapshot,omitempty"`
	RestoreCacheSize      resource.Quantity  `json:"restoreCacheSize,omitempty"`
	SMBConfiguration      SMBConfiguration   `json:"smbConfiguration,omitempty"`
	SMBPasswordResetCount int                `json:"smbPasswordResetCount,omitempty"`
	NoSnapDir             bool               `json:"noSnapDir,omitempty"`
	ProvisionedSize       resource.Quantity  `json:"provisionedSize"`
	DataProtection        DataProtectionType `json:"dataProtection"`
}
type SMBConfiguration struct {
	Netbios                     string   `json:"netbios"`
	Domain                      string   `json:"domain"`
	OrganizationUnit            string   `json:"organizationUnit,omitempty"`
	DNSIPs                      []string `json:"dnsIps"`
	EncryptedADCommunication    bool     `json:"encryptedADCommunication"`
	PreferredADSite             string   `json:"preferredADSite,omitempty"`
	EncryptionEnabled           bool     `json:"encryptionEnabled"`
	HideShare                   bool     `json:"hideShare"`
	ShareAccessBasedEnumeration bool     `json:"accessBasedEnumeration"`
}

type NetAppVolumeConditionType string

// Valid conditions for a NetAppVolume
const (
	NetAppVolumeOnline                    NetAppVolumeConditionType = "Online"
	NetAppVolumeExportPolicyError         NetAppVolumeConditionType = "ExportPolicyError"
	NetAppVolumeSnapshotPolicyError       NetAppVolumeConditionType = "SnapshotPolicyError"
	NetAppVolumeCloudSnapshotPreparing    NetAppVolumeConditionType = "CloudSnapshotPreparing"
	NetAppVolumeCloudSnapshotRestoring    NetAppVolumeConditionType = "CloudSnapshotRestoring"
	NetAppVolumeCloudSnapshotRestored     NetAppVolumeConditionType = "CloudSnapshotRestored"
	NetAppVolumeCloudSnapshotRestoreError NetAppVolumeConditionType = "CloudSnapshotRestoreError"
	NetAppVolumeResized                   NetAppVolumeConditionType = "VolumeResized"
	NetAppVolumeResizeError               NetAppVolumeConditionType = "VolumeResizeError"
	NetAppVolumeSMBConfigurationError     NetAppVolumeConditionType = "SMBConfigurationError"
	NetAppVolumeSMBPasswordResetError     NetAppVolumeConditionType = "SMBPasswordResetError"
	NetAppVolumeUpgradeDMAPHaltRequested  NetAppVolumeConditionType = "HaltRequested"
	NetAppVolumeUpgradeDMAPHaltSucceeded  NetAppVolumeConditionType = "HaltSucceeded"
	NetAppVolumeUpgradeFailed             NetAppVolumeConditionType = "UpgradeFailed"
	NetAppVolumeCreateError               NetAppVolumeConditionType = "CreateFailed"
)

// NetAppVolumeCondition contains the condition information for a NetAppVolume
type NetAppVolumeCondition struct {
	// Type of NetAppVolume condition.
	Type NetAppVolumeConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// Last time we got an update on a given condition.
	// +optional
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime,omitempty"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
	// Timestamp of when the volume came online for the first time
	FirstVolumeOnlineTime metav1.Time `json:"firstVolumeOnlineTime,omitempty"`
}

type NetAppStoragePoolConditionType string

// Valid conditions for a NetAppStoragePool
const (
	NetAppStoragePoolOnline                      NetAppStoragePoolConditionType = "Online"
	NetAppStoragePoolResizeError                 NetAppStoragePoolConditionType = "StoragePoolResizeError"
	NetAppStoragePoolSMBConfigurationError       NetAppStoragePoolConditionType = "SMBConfigurationError"
	NetAppStoragePoolSMBPasswordResetError       NetAppStoragePoolConditionType = "SMBPasswordResetError"
	NetAppStoragePoolAddressAllocationError      NetAppStoragePoolConditionType = "AddressAllocationError"
	NetAppStoragePoolUpgradeDMAPHaltRequested    NetAppStoragePoolConditionType = "HaltRequested"
	NetAppStoragePoolUpgradeDMAPHaltSucceeded    NetAppStoragePoolConditionType = "HaltSucceeded"
	NetAppStoragePoolUpgradeFailed               NetAppStoragePoolConditionType = "UpgradeFailed"
	NetAppStoragePoolSMBServerConfigurationError NetAppStoragePoolConditionType = "SMBServerConfigurationError"
	NetAppStoragePoolADDNSConfigurationError     NetAppStoragePoolConditionType = "ADDNSConfigurationError"
)

// NetAppVolumeCondition contains the condition information for a NetAppVolume
type NetAppStoragePoolCondition struct {
	// Type of NetAppStoragePool condition.
	Type NetAppStoragePoolConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// Last time we got an update on a given condition.
	// +optional
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime,omitempty"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
	// Timestamp of when the StoragePool came online for the first time
	FirstStoragePoolOnlineTime metav1.Time `json:"firstStoragePoolOnlineTime,omitempty"`
}

// NetAppVolumeStatus is the status for a NetAppVolume resource
type NetAppVolumeStatus struct {
	ExportPolicy          string                  `json:"exportPolicy,omitempty"`
	SnapshotPolicy        string                  `json:"snapshotPolicy"`
	SMBConfiguration      SMBConfiguration        `json:"smbConfiguration,omitempty"`
	SMBPasswordResetCount int                     `json:"smbPasswordResetCount,omitempty"`
	NvramPvc              string                  `json:"nvramPvcs"`
	NvramSize             int64                   `json:"nvramSize"`
	DataPvcs              []string                `json:"dataPvcs"`
	QuarkDeployment       string                  `json:"quarkDeployment"`
	StorageDeployment     string                  `json:"storageDeployment"`
	Size                  resource.Quantity       `json:"size"`
	ServiceLevel          string                  `json:"serviceLevel"`
	ExportAddress         string                  `json:"exportAddress"`
	RestoreCacheSize      resource.Quantity       `json:"restoreCacheSize,omitempty"`
	RestorePercent        int64                   `json:"restorePercent"`
	QuarkUUID             string                  `json:"QuarkUUID"`
	QuarkNvramID          string                  `json:"QuarkNvramID"`
	QuarkVersion          string                  `json:"QuarkVersion"`
	VolumeUUID            string                  `json:"VolumeUUID"`
	Conditions            []NetAppVolumeCondition `json:"conditions"`
	NoSnapDir             bool                    `json:"noSnapDir"`
	ProvisionedSize       resource.Quantity       `json:"provisionedSize"`
	QubeVersion           string                  `json:"qubeVersion"`
	Zone                  string                  `json:"zone,omitempty"`
	VolumePath            string                  `json:"volumePath"`
	DataProtection        DataProtectionType      `json:"dataProtection"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppVolumeList is a list of NetAppVolume resources
type NetAppVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NetAppVolume `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppExportPolicy is the specification for a Netapp volume export policy
type NetAppExportPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetAppExportPolicySpec   `json:"spec"`
	Status            NetAppExportPolicyStatus `json:"status"`
}

// NetAppExportPolicySpec is the spec for a Netapp volume export policy
type NetAppExportPolicySpec struct {
	Rules []ExportRule `json:"rules"`
}

// NetAppExportPolicyStatus is the status for a Netapp volume export policy
type NetAppExportPolicyStatus struct {
	Rules []ExportRule `json:"rules"`
}

type ExportRule struct {
	RuleIndex    int    `json:"ruleIndex"`
	Clients      string `json:"clients"`
	AccessMode   string `json:"accessMode"`
	SecurityType string `json:"securityType"`
	Protocol     string `json:"protocol"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppVolumeExportPolicyList is a list of NetAppExportPolicy resources
type NetAppExportPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NetAppExportPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppSnapshotPolicy is the specification for a Netapp volume snapshot policy
type NetAppSnapshotPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetAppSnapshotPolicySpec   `json:"spec"`
	Status            NetAppSnapshotPolicyStatus `json:"status"`
}

// NetAppSnapshotPolicySpec is the spec for a Netapp volume snapshot policy
type NetAppSnapshotPolicySpec struct {
	Enabled   bool           `json:"enabled"`
	Schedules []CronSchedule `json:"schedules"`
}

// NetAppSnapshotPolicyStatus is the status for a Netapp volume snapshot policy
type NetAppSnapshotPolicyStatus struct {
	Enabled   bool           `json:"enabled"`
	Schedules []CronSchedule `json:"schedules"`
}

type CronSchedule struct {
	Prefix          string  `json:"prefix"`
	Months          []int64 `json:"months"`
	DaysOfMonth     []int64 `json:"daysOfMonth"`
	DaysOfWeek      []int64 `json:"daysOfWeek"`
	Hours           []int64 `json:"hours"`
	Minutes         []int64 `json:"minutes"`
	SnapshotsToKeep int     `json:"snapshotsToKeep"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppVolumeSnapshotPolicyList is a list of NetAppSnapshotPolicy resources
type NetAppSnapshotPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NetAppSnapshotPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppServiceLevel is the specification for a Netapp volume service level
type NetAppServiceLevel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetAppServiceLevelSpec   `json:"spec"`
	Status            NetAppServiceLevelStatus `json:"status"`
}

// NetAppServiceLevelSpec is the spec for a Netapp volume service level
type NetAppServiceLevelSpec struct {
	DataVolumeStorageClass  string            `json:"dataVolumeStorageClass"`
	NvramVolumeStorageClass string            `json:"nvramVolumeStorageClass"`
	QuarkVolumeStorageClass string            `json:"quarkVolumeStorageClass"`
	SchedulerName           string            `json:"schedulerName"`
	IOPSPerTiB              int               `json:"iopsPerTiB"`
	MiBsPerTiB              int               `json:"MiBsPerTiB"`
	CapacityLimit           resource.Quantity `json:"capacityLimit"`
	DiskIOPSPerGiB          int               `json:"diskIopsPerGiB"`
	TpAmpFactor             float32           `json:"tpAmpFactor"`
	IOPSLimit               int               `json:"iopsLimit"`
	TputLimit               int               `json:"tputLimit"`
	QuarkPodBuckets         []QuarkPodBucket  `json:"quarkPodBuckets"`
	Regional                bool              `json:"regional"`
	RightSize               bool              `json:"rightSize"`
	QuarkVersion            string            `json:"quarkVersion"`
}

type NetAppServiceLevelStatus struct {
	QuarkPodBuckets []QuarkPodBucket `json:"quarkPodBuckets"`
	Regional        bool             `json:"regional"`
	RightSize       bool             `json:"rightSize"`
	QuarkVersion    string           `json:"quarkVersion"`
}

type QuarkPodBucket struct {
	Name               string              `json:"name"`
	NodeSelector       string              `json:"nodeSelector,omitempty"`
	NvramSize          resource.Quantity   `json:"nvramSize"`
	MinVolumeSize      resource.Quantity   `json:"minVolumeSize"`
	MaxVolumeSize      resource.Quantity   `json:"maxVolumeSize"`
	DmsCPUs            resource.Quantity   `json:"dmsCPUs"`
	ContainerResources []ContainerResource `json:"containerResources,omitempty"`
}

type ContainerResource struct {
	ContainerName string            `json:"containerName"`
	MemRequest    resource.Quantity `json:"memRequest"`
	CPUsRequest   resource.Quantity `json:"cpusRequest"`
	MemLimit      resource.Quantity `json:"memLimit"`
	CPUsLimit     resource.Quantity `json:"cpusLimit"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppVolumeServiceLevelList is a list of NetAppServiceLevel resources
type NetAppServiceLevelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NetAppServiceLevel `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppVolumeSnapshot is a specification for a NetApp Volume Snapshot resource
type NetAppVolumeSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetAppVolumeSnapshotSpec   `json:"spec"`
	Status NetAppVolumeSnapshotStatus `json:"status"`
}

// NetAppVolumeSnapshotSpec is the spec for a NetAppVolumeSnapshot resource
type NetAppVolumeSnapshotSpec struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

type NetAppVolumeSnapshotConditionType string

// Valid conditions for a NetAppVolumeSnapshot
const (
	SnapshotCreating    NetAppVolumeSnapshotConditionType = "Creating"
	SnapshotReady       NetAppVolumeSnapshotConditionType = "Ready"
	SnapshotCreateError NetAppVolumeSnapshotConditionType = "CreateError"
)

// NetAppVolumeSnapshotCondition contains the condition information for a NetAppVolumeSnapshot
type NetAppVolumeSnapshotCondition struct {
	// Type of NetAppVolumeSnapshot condition.
	Type NetAppVolumeSnapshotConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// Last time we got an update on a given condition.
	// +optional
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime,omitempty"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

// NetAppVolumeSnapshotStatus is the status for a NetAppVolumeSnapshot resource
type NetAppVolumeSnapshotStatus struct {
	Conditions []NetAppVolumeSnapshotCondition `json:"conditions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppVolumeSnapshotList is a list of NetAppVolumeSnapshot resources
type NetAppVolumeSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NetAppVolumeSnapshot `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppVolumeSnapshotsList is a specification for a NetApp Volume Snapshot List resource
type NetAppVolumeSnapshotsList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetAppVolumeSnapshotsListSpec   `json:"spec"`
	Status NetAppVolumeSnapshotsListStatus `json:"status"`
}

// NetAppVolumeSnapshotsListSpec is the spec for a NetAppVolumeSnapshotsList resource
type NetAppVolumeSnapshotsListSpec struct {
	DeletedSnapshots []string `json:"deletedSnapshots"`
	RestoredSnapshot string   `json:"restoredSnapshot"`
}

type NetAppVolumeSnapshotsListConditionType string

// Valid conditions for a NetAppVolumeSnapshotsList
const (
	SnapshotsListReady   NetAppVolumeSnapshotsListConditionType = "Ready"
	SnapshotRestoreError NetAppVolumeSnapshotsListConditionType = "RestoreError"
)

// NetAppVolumeSnapshotsListCondition contains the condition information for a NetAppVolumeSnapshotsList
type NetAppVolumeSnapshotsListCondition struct {
	// Type of NetAppVolumeSnapshotList condition.
	Type NetAppVolumeSnapshotsListConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// Last time we got an update on a given condition.
	// +optional
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime,omitempty"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

// NetAppVolumeSnapshotStatus is the status for a NetAppVolumeSnapshot resource
type NetAppVolumeSnapshotInfo struct {
	Name         string            `json:"name"`
	Size         resource.Quantity `json:"size"`
	CreationTime metav1.Time       `json:"creationTime,omitempty"`
}

// NetAppVolumeSnapshotsListStatus is the status for a NetAppVolumeSnapshotsList resource
type NetAppVolumeSnapshotsListStatus struct {
	DeletedSnapshots []string                             `json:"deletedSnapshots"`
	RestoredSnapshot string                               `json:"restoredSnapshot"`
	Snapshots        []NetAppVolumeSnapshotInfo           `json:"snapshots"`
	Conditions       []NetAppVolumeSnapshotsListCondition `json:"conditions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppVolumeSnapshotsListList is a list of NetAppVolumeSnapshotsList resources
type NetAppVolumeSnapshotsListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NetAppVolumeSnapshotsList `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppVolumeCloudSnapshot is a specification for a NetApp Volume Snapshot resource
type NetAppVolumeCloudSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetAppVolumeCloudSnapshotSpec   `json:"spec"`
	Status NetAppVolumeCloudSnapshotStatus `json:"status"`
}

// NetAppVolumeCloudSnapshotSpec is the spec for a NetAppVolumeCloudSnapshot resource
type NetAppVolumeCloudSnapshotSpec struct {
	ProviderType NetAppVolumeCloudSnapshotProviderType `json:"providerType,omitempty"`
	DisplayName  string                                `json:"displayName,omitempty"`
	Imported     bool                                  `json:"imported,omitempty"`
	Container    string                                `json:"container"`
	Server       string                                `json:"server"`
	Port         int32                                 `json:"port"`
	SecretRef    v1.SecretReference                    `json:"secretRef"`
}

type NetAppVolumeCloudSnapshotProviderType string

// Valid provider types for cloud
const (
	AzureCloud  NetAppVolumeCloudSnapshotProviderType = "Azure_Cloud"
	GoogleCloud NetAppVolumeCloudSnapshotProviderType = "GoogleCloud"
	AwsS3       NetAppVolumeCloudSnapshotProviderType = "AWS_S3"
	SGWS        NetAppVolumeCloudSnapshotProviderType = "SGWS"
	ONTAPS3     NetAppVolumeCloudSnapshotProviderType = "ONTAPS3"
)

type NetAppVolumeCloudSnapshotConditionType string

// Valid conditions for a NetAppVolumeCloudSnapshot
const (
	CloudSnapshotReady    NetAppVolumeCloudSnapshotConditionType = "CloudSnapshotReady"
	CloudSnapshotCreating NetAppVolumeCloudSnapshotConditionType = "CloudSnapshotCreating"
	CloudSnapshotDeleting NetAppVolumeCloudSnapshotConditionType = "CloudSnapshotDeleting"
	CloudSnapshotError    NetAppVolumeCloudSnapshotConditionType = "CloudSnapshotError"
)

// NetAppVolumeCloudSnapshotCondition contains the condition information for a NetAppVolumeCloudSnapshot
type NetAppVolumeCloudSnapshotCondition struct {
	// Type of NetAppVolumeCloudSnapshot condition.
	Type NetAppVolumeCloudSnapshotConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// Last time we got an update on a given condition.
	// +optional
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime,omitempty"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

// NetAppVolumeCloudSnapshotStatus is the status for a NetAppVolumeCloudSnapshot resource
type NetAppVolumeCloudSnapshotStatus struct {
	Conditions        []NetAppVolumeCloudSnapshotCondition `json:"conditions"`
	EndpointUUID      string                               `json:"endpointUUID,omitempty"`
	BytesTransferred  int64                                `json:"bytesTransferred"`
	LogicalSize       int64                                `json:"logicalSize,omitempty"`
	CompletionPercent int64                                `json:"completionPercent"`
	CreationTime      metav1.Time                          `json:"creationTime,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppVolumeSnapshotList is a list of NetAppVolumeSnapshot resources
type NetAppVolumeCloudSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NetAppVolumeCloudSnapshot `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppVolumeUsage is a specification for a NetApp Volume Usage resource
type NetAppVolumeUsage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetAppVolumeUsageSpec   `json:"spec"`
	Status NetAppVolumeUsageStatus `json:"status"`
}

// NetAppVolumeSpec is the spec for a NetAppVolumeUsage resource
type NetAppVolumeUsageSpec struct {
	Name string `json:"name"`
}

// NetAppVolumeUsageStatus is the status for a NetAppVolumeUsage resource
type NetAppVolumeUsageStatus struct {
	LogicalUsed   int64 `json:"logicalUsed"`
	PhysicalUsed  int64 `json:"physicalUsed"`
	AllocatedSize int64 `json:"allocatedSize"`
	SnapshotUsed  int64 `json:"snapshotUsed"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppVolumeUsageList is a list of NetAppVolumeUsage resources
type NetAppVolumeUsageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NetAppVolumeUsage `json:"items"`
}

// QuarkSpec defines the desired state of Quark CVS deployment
type QuarkSpec struct {
	Version     string   `json:"version"`
	Namespace   string   `json:"namespace"`
	Cloud       string   `json:"cloud"`
	DarkContent string   `json:"darkContent"`
	Nodevol     string   `json:"nodevol"`
	Secrets     []Secret `json:"secrets,omitempty"`

	Regional           bool             `json:"regional,omitempty"`
	StorageClass       StorageClass     `json:"storageClass" yaml:"storageClass"`
	ServiceLevels      []ServiceLevel   `json:"serviceLevels,omitempty" yaml:"serviceLevels,omitempty"`
	Volume             Volume           `json:"volume,omitempty"`
	ILB                ILB              `json:"ilb"`
	Images             Images           `json:"images"`
	Project            Project          `json:"project"`
	Logging            Logging          `json:"logging"`
	ControllerConfig   ControllerConfig `json:"controllerConfig,omitempty" yaml:"controllerConfig,omitempty"`
	DisableDiskPacking string           `json:"disableDiskPacking"`
	Debug              bool             `json:"debug,omitempty"`
	EnableStoragePool  bool             `json:"enableStoragePool,omitempty"`
	EnableMultiVolume  bool             `json:"enableMultiVolume,omitempty"`
	EnableCcpd         bool             `json:"enableCcpd,omitempty"`
	// Enables the service mesh required for cross-region-replication networking to work
	EnableServiceMesh bool `json:"enableServiceMesh,omitempty"`
}

type QuarkReadyStatus string

const (
	QuarkReadyStatusReady    QuarkReadyStatus = "Ready"
	QuarkReadyStatusNotReady QuarkReadyStatus = "NotReady"
)

// QuarkHealth shows the current health of tracked components
type QuarkHealth struct {
	NVC                    QuarkReadyStatus `json:"nvc"`
	NHC                    QuarkReadyStatus `json:"nhc,omitempty"`
	NodeMonitor            QuarkReadyStatus `json:"nodeMonitor,omitempty"`
	SvcMeshCtrl            QuarkReadyStatus `json:"svcmeshctrl"`
	NSC                    QuarkReadyStatus `json:"nsc"`
	CSC                    QuarkReadyStatus `json:"csc"`
	ADC                    QuarkReadyStatus `json:"adc"`
	IPD                    QuarkReadyStatus `json:"ipd"`
	Operator               QuarkReadyStatus `json:"operator"`
	EtcdCluster            QuarkReadyStatus `json:"etcdCluster"`
	Nodevol                QuarkReadyStatus `json:"nodevol"`
	KubeStateMetrics       QuarkReadyStatus `json:"kubeStateMetrics"`
	PrometheusServer       QuarkReadyStatus `json:"prometheusServer"`
	PrometheusNodeExporter QuarkReadyStatus `json:"prometheusNodeExporter"`
}

func NewQuarkHealth() QuarkHealth {
	return QuarkHealth{
		NVC:         QuarkReadyStatusNotReady,
		NHC:         QuarkReadyStatusNotReady,
		NSC:         QuarkReadyStatusNotReady,
		CSC:         QuarkReadyStatusNotReady,
		ADC:         QuarkReadyStatusNotReady,
		SvcMeshCtrl: QuarkReadyStatusNotReady,
	}
}

// ControllerConfig defines controllerConfig section of Quark Spec
type ControllerConfig struct {
	AdcGetInterval         string `json:"adcGetInterval" yaml:"adcGetInterval"`
	Variant                string `json:"variant"`
	EtcdRepository         string `json:"etcdRepository" yaml:"etcdRepository"`
	EtcdVersion            string `json:"etcdVersion,omitempty" yaml:"etcdVersion,omitempty"`
	EtcdClusterSize        int    `json:"etcdClusterSize,omitempty" yaml:"etcdClusterSize,omitempty"`
	HealthCheckPollSeconds uint   `json:"healthCheckPollSeconds,omitempty" yaml:"healthCheckPollSeconds,omitempty"`
}

// Logging defines logging section of Quark Spec
type Logging struct {
	Hosts     string `json:"hosts,omitempty"`
	Port      string `json:"port,omitempty"`
	RemoteURL string `json:"remoteURL,omitempty"`
	Output    string `json:"output"`
	TLS       string `json:"TLS,omitempty"`
}

// Secret defines a secret to be created by operator
type Secret struct {
	Name      string `json:"name"`
	Key       string `json:"key"`
	Type      string `json:"type,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Data      []byte `json:"data"`
}

// Volume defines the volume section of Quark Spec
type Volume struct {
	QuarkVolumeSize string `json:"quarkVolumeSize,omitempty"`
	NvramSize       string `json:"nvramSize,omitempty"`
	MinPVCSize      string `json:"minPVCSize,omitempty"`
	MaxPVCSize      string `json:"maxPVCSize,omitempty"`
	MinDataPVCs     string `json:"minDataPVCs,omitempty"`
}

// ILB defines the ilb section of Quark Spec
type ILB struct {
	Enable       string `json:"enable"`
	Subnet       string `json:"subnet,omitempty"`
	Protocol     string `json:"protocol,omitempty"`
	GlobalAccess string `json:"globalAccess"`
}

// ServiceLevel defines an available service level for a Quark volume.
type ServiceLevel struct {
	Name         string `json:"name,omitempty"`
	Regional     bool   `json:"regional,omitempty"`
	RightSize    bool   `json:"rightSize,omitempty" yaml:"rightSize,omitempty"`
	NvramSC      string `json:"nvram,omitempty" yaml:"nvram,omitempty"`
	DataSC       string `json:"data,omitempty" yaml:"data,omitempty"`
	QuarkSC      string `json:"quark,omitempty" yaml:"quark,omitempty"`
	PrometheusSC string `json:"prometheus,omitempty" yaml:"prometheus,omitempty"`
}

// SCProvisioner define the various storage provisioners available
type SCProvisioner struct {
	Provisioner  string            `json:"provisioner,omitempty"`
	StorageClass string            `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`
	Name         string            `json:"name,omitempty"`
	Regional     bool              `json:"regional,omitempty"`
	Parameters   map[string]string `json:"parameters,omitempty"`
}

// StorageClass defines the storageClass section of Quark Spec
type StorageClass struct {
	Provisioners []SCProvisioner `json:"provisioners,omitempty"`
	Etcd         string          `json:"etcd,omitempty"`
	Nodevol      string          `json:"nodevol,omitempty"`
	Prometheus   string          `json:"prometheus,omitempty"`
	// These fields are mostly deprecated; however, they're used by NVC Client
	// to override default settings for the Standard service level storage
	// classes. This was necessary to preserve compatibility with nvcclient cli.
	// TODO: Deprecate these fully and set the corresponding values based on the
	// variant.
	Nvram string `json:"nvram,omitempty"`
	Data  string `json:"data,omitempty"`
	Quark string `json:"quark,omitempty"`
}

// Images defines the images section of Quark Spec
type Images struct {
	PullSecret string `json:"pullSecret,omitempty" yaml:"pullSecret,omitempty"`
	PullPolicy string `json:"pullPolicy" yaml:"pullPolicy"`
	// Going with the approach of supporting a single service mesh solution at a time.
	// Moving forward, if/when it is determined we will support multiple svc mesh
	// solutions (likely, and based on discussion with a wider audience), we should
	// replace SvcMeshCtrl<string> with its own type, and use it in images
	// Eg:
	//	type ServiceMeshImages struct {
	//		Anthos	string `json:"anthos,omitempty"`
	//		Linkerd string `json:"linkerd,omitempty"`
	//		Istio	string `json:"istio,omitempty"`
	// 	}
	// And entry below would be
	// SvcMeshImages		  ServiceMeshImages `json:"svcmeshimages" yaml:"svcmeshctrl",omitempty"`
	// Needless to say, this approach would need to be supported by our make/build system as well.
	SvcMeshCtrl               string `json:"svcmeshctrl" yaml:"svcmeshctrl,omitempty"`
	NVC                       string `json:"nvc" yaml:"nvc,omitempty"`
	DMAP                      string `json:"dmap" yaml:"dmap,omitempty"`
	Podrick                   string `json:"podrick" yaml:"podrick,omitempty"`
	Secd                      string `json:"secd,omitempty" yaml:"secd,omitempty"`
	Etcd                      string `json:"etcd,omitempty" yaml:"etcd,omitempty"`
	Supportability            string `json:"supportability" yaml:"supportability,omitempty"`
	Storage                   string `json:"storage,omitempty" yaml:"storage,omitempty"`
	Logging                   string `json:"logging" yaml:"logging,omitempty"`
	QuarkLogging              string `json:"quarkLogging" yaml:"quarkLogging,omitempty"`
	NSC                       string `json:"nsc" yaml:"nsc,omitempty"`
	Startup                   string `json:"startup" yaml:"startup,omitempty"`
	CSC                       string `json:"csc" yaml:"csc,omitempty"`
	ADC                       string `json:"adc" yaml:"adc,omitempty"`
	NHC                       string `json:"nhc" yaml:"nhc,omitempty"`
	CCPD                      string `json:"ccpd" yaml:"ccpd,omitempty"`
	NodeMonitor               string `json:"nodeMonitor,omitempty" yaml:"nodeMonitor,omitempty"`
	EtcdOperator              string `json:"etcdOperator,omitempty" yaml:"etcdOperator,omitempty"`
	EtcdInitContainer         string `json:"etcdInitContainer,omitempty" yaml:"etcdInitContainer,omitempty"`
	KubeMetrics               string `json:"kubeMetrics" yaml:"kubeMetrics,omitempty"`
	PrometheusExporter        string `json:"prometheusExporter" yaml:"prometheusExporter,omitempty"`
	PrometheusConfigMapReload string `json:"prometheusConfigMapReload" yaml:"prometheusConfigMapReload,omitempty"`
	Prometheus                string `json:"prometheus" yaml:"prometheus,omitempty"`
	Busybox                   string `json:"busybox" yaml:"busybox,omitempty"`
	HASpare                   string `json:"haSpare" yaml:"haSpare,omitempty"`
	InitContainer             string `json:"initcontainer" yaml:"initcontainer,omitempty"`
	NodevolWorker             string `json:"nodevolWorker" yaml:"nodevolWorker,omitempty"`
	NodevolController         string `json:"nodevolController" yaml:"nodevolController,omitempty"`
}

// Project defines the project section of Quark Spec
type Project struct {
	GCSBucketName               string `json:"gcsBucketName"`
	TenantProjectNumber         string `json:"tenantProjectNumber"`
	ClusterName                 string `json:"clusterName"`
	CustomerProjectNumber       string `json:"customerProjectNumber"`
	CustomerVPCName             string `json:"customerVPCName"`
	SharedCluster               string `json:"sharedCluster,omitempty"`
	AzureContainerName          string `json:"azureContainerName,omitempty"`
	AzureCustomerResourceGroup  string `json:"azureCustomerResourceGroup,omitempty"`
	AzureCustomerSubscriptionId string `json:"azureCustomerSubscriptionId,omitempty"`
}

// QuarkStatus defines the observed state of Quark CR
type QuarkStatus struct {
	Version string      `json:"version"`
	Images  Images      `json:"images"`
	Health  QuarkHealth `json:"health,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Quark is the instance of quark
type Quark struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   QuarkSpec   `json:"spec,omitempty"`
	Status QuarkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// QuarkList contains a list of Quark
type QuarkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Quark `json:"items"`
}

/* NodevolSet CRD related definitions */

// NodevolSetStatus defines the observed state of NodevolSet CR
type NodevolSetStatus struct {
	// Latest version where all the nodevol pods are running
	Version string `json:"version"`

	// Present Data
	// Number of nodevol pods in running state
	Ready int32 `json:"running"`
	// Number of nodevol pods deleted and recovery not yet started
	Deleted int32 `json:"deleted"`
	// Number of nodevol pods in recovery
	Recovering int32 `json:"recovering"`
	// Number of nodevol pods updating to latest version
	Updating int32 `json:"updating"`
	// Number of running nodevols
	AliveNodes int32 `json:"aliveNodes"`
	// NUmber of nodevols pending the recovery
	Orphan int32 `json:"orphanNodes"`
	// Total number of nodevols including orphan nodevols
	TotalNodes int32 `json:"totalNodes"`

	// Historical Data
	// Total number of Jobs failed from the creation time of nodevolSet
	TotalFailedJobs int32 `json:"totalFailedJobs"`
	// Total number of nodevol volumes recoverd from the creation time of nodeolSet
	TotalRecovered int32 `json:"totalRecovered"`

	// Time stamp for last update
	LastUpdatedTime metav1.Time `json:"lastUpdateTime,omitempty" protobuf:"bytes,4,opt,name=lastUpdateTime"`
}

// NodevolImages defines the images section of NodevolSet Spec
type NodevolImages struct {
	PullSecret string `json:"pullSecret,omitempty"`
	PullPolicy string `json:"pullPolicy"`

	Init           string `json:"init"`
	VolCleanup     string `json:"volCleanup"`
	Supportability string `json:"supportability"`
	Startup        string `json:"startup"`
	Dummy          string `json:"dummy"`
}

// NodevolVolume defines the volume section of nodevolSet spec
type NodevolVolume struct {
	NodevolSize   int64  `json:"nodevolSize"`
	StorageClass  string `json:"storageClass"`
	HostMountpath string `json:"hostMountpath"`
}

// NodevolSetSpec defines the desired state of NodevolSet
type NodevolSetSpec struct {
	Version   string `json:"version"`
	Namespace string `json:"namespace"`
	Cloud     string `json:"cloud"`
	Regional  bool   `json:"regional,omitempty"`

	Volume NodevolVolume `json:"volume"`
	Images NodevolImages `json:"images"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodevolSet is the instance of nodevolSet
type NodevolSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodevolSetSpec   `json:"spec,omitempty"`
	Status NodevolSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodevolSetList contains a list of NodevolSet
type NodevolSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodevolSet `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StoragePool is a specification for a Storage pool resource
type NetAppStoragePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetAppStoragePoolSpec   `json:"spec"`
	Status NetAppStoragePoolStatus `json:"status"`
}

// NetAppStoragePoolSpec is the spec for a NetAppStorgePool resource
type NetAppStoragePoolSpec struct {
	Size                   resource.Quantity      `json:"size"`
	ProvisionedSize        resource.Quantity      `json:"provisionedSize"`
	DisplayName            string                 `json:"displayName"`
	ServiceLevel           string                 `json:"serviceLevel"`
	QuarkVersion           string                 `json:"quarkVersion"`
	ADDNSIPs               []string               `json:"adDnsIps"`
	SMBServerConfiguration SMBServerConfiguration `json:"smbServerConfiguration,omitempty"`
	SMBADConfiguration     SMBADConfiguration     `json:"smbADConfiguration,omitempty"`
	SMBPasswordResetCount  int                    `json:"smbPasswordResetCount,omitempty"`
	Zone                   string                 `json:"zone,omitempty"`
	AvailableZones         []string               `json:"availableZones,omitempty"`
	LastZoneMigrationCause string                 `json:"lastZoneMigrationCause,omitempty"`
	// Using a pointer here to avoid the field being displayed as null in zonal clusters
	LastZoneMigrationTime *metav1.Time `json:"lastZoneMigrationTime,omitempty"`
	// ZoneMigrationCount is used by Podrick to determine when it needs to
	// update LastZoneMigrationCause and LastZoneMigrationTime
	ZoneMigrationCount int64 `json:"zoneMigrationCount,omitempty"`
	GlobalAccess       bool  `json:"globalAccess,omitempty"`
	// True if global access is being aborted before the services have been
	// set up.
	GlobalAccessAborting bool `json:"globalAccessAborting,omitempty"`
}

type SMBServerConfiguration struct {
	Netbios          string `json:"netbios"`
	Domain           string `json:"domain"`
	OrganizationUnit string `json:"organizationUnit,omitempty"`
	PreferredADSite  string `json:"preferredADSite,omitempty"`
}

type SMBADConfiguration struct {
	EncryptedADCommunication bool `json:"encryptedADCommunication"`
}

// NetAppStoragePoolStatus is the status for a NetAppStorgePool resource
type NetAppStoragePoolStatus struct {
	Size                   resource.Quantity      `json:"size"`
	ProvisionedSize        resource.Quantity      `json:"provisionedSize"`
	DisplayName            string                 `json:"displayName"`
	ServiceLevel           string                 `json:"serviceLevel"`
	QuarkVersion           string                 `json:"quarkVersion"`
	ADDNSIPs               []string               `json:"adDnsIps"`
	SMBServerConfiguration SMBServerConfiguration `json:"smbServerConfiguration,omitempty"`
	SMBADConfiguration     SMBADConfiguration     `json:"smbADConfiguration,omitempty"`
	SMBPasswordResetCount  int                    `json:"smbPasswordResetCount,omitempty"`
	ExportAddress          string                 `json:"exportAddress"`
	DataPvcs               []string               `json:"dataPvcs"`
	NvramPvc               string                 `json:"nvramPvcs"`
	NvramSize              int64                  `json:"nvramSize"`
	StorageDeployment      string                 `json:"storageDeployment"`
	QuarkUUID              string                 `json:"QuarkUUID"`
	QuarkNvramID           string                 `json:"QuarkNvramID"`
	QuarkDeployment        string                 `json:"quarkDeployment"`
	QubeVersion            string                 `json:"qubeVersion"`
	Zone                   string                 `json:"zone,omitempty"`
	AvailableZones         []string               `json:"availableZones,omitempty"`
	LastZoneMigrationCause string                 `json:"lastZoneMigrationCause,omitempty"`
	// Using a pointer here to avoid the field being displayed as null in zonal clusters
	LastZoneMigrationTime *metav1.Time `json:"lastZoneMigrationTime,omitempty"`
	// Should always be <= Spec.ZoneMigrationCount. See the comment on the
	// corresponding spec field.
	ZoneMigrationCount          int64                        `json:"zoneMigrationCount,omitempty"`
	Conditions                  []NetAppStoragePoolCondition `json:"conditions"`
	GlobalAccess                bool                         `json:"globalAccess,omitempty"`
	ADClientIP                  string                       `json:"adClientIP,omitempty"`
	IlbConnectionTrackingPolicy string                       `json:"ilbConnectionTrackingPolicy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppStoragePoolList is a list of NetAppStoragePool resources
type NetAppStoragePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NetAppStoragePool `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterPeering is a specification for a Storage pool resource
type NetAppClusterPeering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetAppClusterPeeringSpec   `json:"spec"`
	Status NetAppClusterPeeringStatus `json:"status"`
}

// NetAppClusterPeeringSpec is the spec for a NetAppClusterPeering resource
type NetAppClusterPeeringSpec struct {
	RemoteEndpoint string `json:"remoteEndpoint"`
}

type NetAppClusterPeeringStateType string

// Valid ClusterPeering states
const (
	ClusterPeeringStatePending NetAppClusterPeeringStateType = "Pending"
	ClusterPeeringStatePeered  NetAppClusterPeeringStateType = "Peered"
)

type NetAppClusterPeeringConditionType string

// Valid conditions for a NetAppClusterPeering
const (
	ServiceMeshEnabled  NetAppClusterPeeringConditionType = "ServiceMeshEnabled"
	ClusterPeeringReady NetAppClusterPeeringConditionType = "ClusterPeeringReady"
	ClusterPeeringError NetAppClusterPeeringConditionType = "ClusterPeeringError"
)

// NetAppClusterPeeringCondition contains the condition information for a NetAppClusterPeering
type NetAppClusterPeeringCondition struct {
	// Type of NetAppClusterPeering condition.
	Type NetAppClusterPeeringConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// Last time we got an update on a given condition.
	// +optional
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime,omitempty"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

// NetAppClusterPeeringStatus is the status for a NetAppClusterPeering resource
type NetAppClusterPeeringStatus struct {
	State      NetAppClusterPeeringStateType   `json:"state,omitempty"`
	Conditions []NetAppClusterPeeringCondition `json:"conditions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppClusterPeeringList is a list of NetAppClusterPeering resources
type NetAppClusterPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NetAppClusterPeering `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterPeering is a specification for a Storage pool resource
type NetAppVolumeReplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetAppVolumeReplicationSpec   `json:"spec"`
	Status NetAppVolumeReplicationStatus `json:"status"`
}

type NetAppVolumeReplicationVolume struct {
	VolumeUUID  string `json:"volumeUUID"`
	SVMUUID     string `json:"svmUUID,omitempty"`
	ClusterUUID string `json:"clusterUUID,omitempty"`
}

type NetAppVolumeReplicationMirrorStateType string

// Valid mirror states
const (
	MirrorStateUninitialized NetAppVolumeReplicationMirrorStateType = "Uninitialized"
	MirrorStateMirrored      NetAppVolumeReplicationMirrorStateType = "Mirrored"
	MirrorStateBroken        NetAppVolumeReplicationMirrorStateType = "Broken"
)

type NetAppVolumeReplicationPolicyType string

// Valid replication policies
const (
	ReplicationPolicyMirrorAllSnapshots NetAppVolumeReplicationPolicyType = "MirrorAllSnapshots"
	ReplicationPolicyMirrorLatest       NetAppVolumeReplicationPolicyType = "MirrorLatest"
	ReplicationPolicyMirrorAndVault     NetAppVolumeReplicationPolicyType = "MirrorAndVault"
)

type NetAppVolumeReplicationScheduleType string

// Valid replication policies
const (
	ReplicationSchedule10Minutely NetAppVolumeReplicationScheduleType = "10minutely"
	ReplicationScheduleHourly     NetAppVolumeReplicationScheduleType = "hourly"
	ReplicationScheduleDaily      NetAppVolumeReplicationScheduleType = "daily"
)

// NetAppVolumeReplicationSpec is the spec for a NetAppVolumeReplication resource
type NetAppVolumeReplicationSpec struct {
	DisplayName         string                                 `json:"displayName"`
	DestinationVolume   NetAppVolumeReplicationVolume          `json:"destinationVolume"`
	SourceVolume        NetAppVolumeReplicationVolume          `json:"sourceVolume"`
	MirrorState         NetAppVolumeReplicationMirrorStateType `json:"mirrorState"`
	ReplicationPolicy   NetAppVolumeReplicationPolicyType      `json:"replicationPolicy"`
	ReplicationSchedule NetAppVolumeReplicationScheduleType    `json:"replicationSchedule"`
	TransferSnapshot    string                                 `json:"transferSnapshot,omitempty"`
}

type NetAppVolumeReplicationStatusType string

// Valid replication statuses
const (
	ReplicationStatusIdle         NetAppVolumeReplicationStatusType = "Idle"
	ReplicationStatusTransferring NetAppVolumeReplicationStatusType = "Transferring"
)

type NetAppVolumeReplicationConditionType string

// Valid conditions for a NetAppVolumeReplication
const (
	VolumeReplicationHealthy          NetAppVolumeReplicationConditionType = "Healthy"
	VolumeReplicationTransferError    NetAppVolumeReplicationConditionType = "TransferError"
	VolumeReplicationTransferComplete NetAppVolumeReplicationConditionType = "TransferComplete"
	VolumeReplicationTransferAborted  NetAppVolumeReplicationConditionType = "TransferAborted"
)

// NetAppVolumeReplicationCondition contains the condition information for a NetAppVolumeReplication
type NetAppVolumeReplicationCondition struct {
	// Type of NetAppVolumeReplication condition.
	Type NetAppVolumeReplicationConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// Last time we got an update on a given condition.
	// +optional
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime,omitempty"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

// NetAppVolumeReplicationStatus is the status for a NetAppVolumeReplication resource
type NetAppVolumeReplicationStatus struct {
	VolumeReplicationUUID string                                 `json:"volumeReplicationUUID,omitempty"`
	TransferSnapshot      string                                 `json:"transferSnapshot,omitempty"`
	NewestSnapshotTime    metav1.Time                            `json:"newestSnapshotTime,omitempty"`
	MirrorState           NetAppVolumeReplicationMirrorStateType `json:"mirrorState,omitempty"`
	ReplicationPolicy     NetAppVolumeReplicationPolicyType      `json:"replicationPolicy,omitempty"`
	ReplicationSchedule   NetAppVolumeReplicationScheduleType    `json:"replicationSchedule"`
	ReplicationStatus     NetAppVolumeReplicationStatusType      `json:"replicationStatus,omitempty"`
	LastTransferDuration  int64                                  `json:"lastTransferDuration,omitempty"`
	LastTransferEndTime   metav1.Time                            `json:"lastTransferEndTime,omitempty"`
	LastTransferSize      int64                                  `json:"lastTransferSize,omitempty"`
	TotalTransferBytes    int64                                  `json:"totalTransferBytes,omitempty"`
	TotalTransferTimeSecs int64                                  `json:"totalTransferTimeSecs,omitempty"`
	Conditions            []NetAppVolumeReplicationCondition     `json:"conditions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetAppVolumeReplicationList is a list of NetAppVolumeReplication resources
type NetAppVolumeReplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NetAppVolumeReplication `json:"items"`
}
