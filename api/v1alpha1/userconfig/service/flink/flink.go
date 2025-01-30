// Code generated by user config generator. DO NOT EDIT.
// +kubebuilder:object:generate=true

package flinkuserconfig

// CIDR address block, either as a string, or in a dict with an optional description field
type IpFilter struct {
	// +kubebuilder:validation:MaxLength=1024
	// Description for IP filter list entry
	Description *string `groups:"create,update" json:"description,omitempty"`

	// +kubebuilder:validation:MaxLength=43
	// CIDR address block
	Network string `groups:"create,update" json:"network"`
}

// Allow access to selected service components through Privatelink
type PrivatelinkAccess struct {
	// Enable flink
	Flink *bool `groups:"create,update" json:"flink,omitempty"`

	// Enable prometheus
	Prometheus *bool `groups:"create,update" json:"prometheus,omitempty"`
}

// Allow access to selected service ports from the public Internet
type PublicAccess struct {
	// Allow clients to connect to flink from the public internet for service nodes that are in a project VPC or another type of private network
	Flink *bool `groups:"create,update" json:"flink,omitempty"`
}
type FlinkUserConfig struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:deprecatedversion:warning="additional_backup_regions is deprecated"
	// Deprecated. Additional Cloud Regions for Backup Replication
	AdditionalBackupRegions []string `groups:"create,update" json:"additional_backup_regions,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	// Enable to upload Custom JARs for Flink applications
	CustomCode *bool `groups:"create" json:"custom_code,omitempty"`

	// +kubebuilder:validation:Enum="1.19";"1.20"
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	// Flink major version
	FlinkVersion *string `groups:"create" json:"flink_version,omitempty"`

	// +kubebuilder:validation:MaxItems=2048
	// Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'
	IpFilter []*IpFilter `groups:"create,update" json:"ip_filter,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1024
	// Task slots per node. For a 3 node plan, total number of task slots is 3x this value
	NumberOfTaskSlots *int `groups:"create,update" json:"number_of_task_slots,omitempty"`

	// +kubebuilder:validation:Minimum=5
	// +kubebuilder:validation:Maximum=60
	// Timeout in seconds used for all futures and blocking Pekko requests
	PekkoAskTimeoutS *int `groups:"create,update" json:"pekko_ask_timeout_s,omitempty"`

	// +kubebuilder:validation:Minimum=1048576
	// +kubebuilder:validation:Maximum=52428800
	// Maximum size in bytes for messages exchanged between the JobManager and the TaskManagers
	PekkoFramesizeB *int `groups:"create,update" json:"pekko_framesize_b,omitempty"`

	// Allow access to selected service components through Privatelink
	PrivatelinkAccess *PrivatelinkAccess `groups:"create,update" json:"privatelink_access,omitempty"`

	// Allow access to selected service ports from the public Internet
	PublicAccess *PublicAccess `groups:"create,update" json:"public_access,omitempty"`

	// Store logs for the service so that they are available in the HTTP API and console.
	ServiceLog *bool `groups:"create,update" json:"service_log,omitempty"`

	// Use static public IP addresses
	StaticIps *bool `groups:"create,update" json:"static_ips,omitempty"`
}
