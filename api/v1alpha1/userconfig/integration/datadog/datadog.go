// Code generated by user config generator. DO NOT EDIT.
// +kubebuilder:object:generate=true

package datadoguserconfig

// Datadog tag defined by user
type DatadogTags struct {
	// +kubebuilder:validation:MaxLength=1024
	// Optional tag explanation
	Comment *string `groups:"create,update" json:"comment,omitempty"`

	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=200
	// Tag format and usage are described here: https://docs.datadoghq.com/getting_started/tagging. Tags with prefix 'aiven-' are reserved for Aiven.
	Tag string `groups:"create,update" json:"tag"`
}

// Datadog Opensearch Options
type Opensearch struct {
	// Enable Datadog Opensearch Cluster Monitoring
	ClusterStatsEnabled *bool `groups:"create,update" json:"cluster_stats_enabled,omitempty"`

	// Enable Datadog Opensearch Index Monitoring
	IndexStatsEnabled *bool `groups:"create,update" json:"index_stats_enabled,omitempty"`

	// Enable Datadog Opensearch Pending Task Monitoring
	PendingTaskStatsEnabled *bool `groups:"create,update" json:"pending_task_stats_enabled,omitempty"`

	// Enable Datadog Opensearch Primary Shard Monitoring
	PshardStatsEnabled *bool `groups:"create,update" json:"pshard_stats_enabled,omitempty"`
}

// Datadog Redis Options
type Redis struct {
	// Enable command_stats option in the agent's configuration
	CommandStatsEnabled *bool `groups:"create,update" json:"command_stats_enabled,omitempty"`
}
type DatadogUserConfig struct {
	// Enable Datadog Database Monitoring
	DatadogDbmEnabled *bool `groups:"create,update" json:"datadog_dbm_enabled,omitempty"`

	// Enable Datadog PgBouncer Metric Tracking
	DatadogPgbouncerEnabled *bool `groups:"create,update" json:"datadog_pgbouncer_enabled,omitempty"`

	// +kubebuilder:validation:MaxItems=32
	// Custom tags provided by user
	DatadogTags []*DatadogTags `groups:"create,update" json:"datadog_tags,omitempty"`

	// +kubebuilder:validation:MaxItems=1024
	// List of custom metrics
	ExcludeConsumerGroups []string `groups:"create,update" json:"exclude_consumer_groups,omitempty"`

	// +kubebuilder:validation:MaxItems=1024
	// List of topics to exclude
	ExcludeTopics []string `groups:"create,update" json:"exclude_topics,omitempty"`

	// +kubebuilder:validation:MaxItems=1024
	// List of custom metrics
	IncludeConsumerGroups []string `groups:"create,update" json:"include_consumer_groups,omitempty"`

	// +kubebuilder:validation:MaxItems=1024
	// List of topics to include
	IncludeTopics []string `groups:"create,update" json:"include_topics,omitempty"`

	// +kubebuilder:validation:MaxItems=1024
	// List of custom metrics
	KafkaCustomMetrics []string `groups:"create,update" json:"kafka_custom_metrics,omitempty"`

	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=100000
	// Maximum number of JMX metrics to send
	MaxJmxMetrics *int `groups:"create,update" json:"max_jmx_metrics,omitempty"`

	// +kubebuilder:validation:MaxItems=1024
	// List of custom metrics
	MirrormakerCustomMetrics []string `groups:"create,update" json:"mirrormaker_custom_metrics,omitempty"`

	// Datadog Opensearch Options
	Opensearch *Opensearch `groups:"create,update" json:"opensearch,omitempty"`

	// Datadog Redis Options
	Redis *Redis `groups:"create,update" json:"redis,omitempty"`
}
