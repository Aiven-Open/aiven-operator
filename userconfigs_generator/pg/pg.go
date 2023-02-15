// Code generated by user config generator. DO NOT EDIT.
// +kubebuilder:object:generate=true

package pgtestuserconfig

import "encoding/json"

func (ip *IpFilter) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	var s string
	err := json.Unmarshal(data, &s)
	if err == nil {
		ip.Network = s
		return nil
	}

	type this struct {
		Network     string  `json:"network"`
		Description *string `json:"description,omitempty" `
	}

	var t *this
	err = json.Unmarshal(data, &t)
	if err != nil {
		return err
	}
	ip.Network = t.Network
	ip.Description = t.Description
	return nil
}

// CIDR address block, either as a string, or in a dict with an optional description field
type IpFilter struct {
	// +kubebuilder:validation:MaxLength=1024
	// Description for IP filter list entry
	Description *string `groups:"create,update" json:"description,omitempty"`

	// +kubebuilder:validation:MaxLength=43
	// CIDR address block
	Network string `groups:"create,update" json:"network"`
}

// Migrate data from existing server
type Migration struct {
	// +kubebuilder:validation:MaxLength=63
	// Database name for bootstrapping the initial connection
	Dbname *string `groups:"create,update" json:"dbname,omitempty"`

	// +kubebuilder:validation:MaxLength=255
	// Hostname or IP address of the server where to migrate data from
	Host string `groups:"create,update" json:"host"`

	// +kubebuilder:validation:MaxLength=2048
	// Comma-separated list of databases, which should be ignored during migration (supported by MySQL only at the moment)
	IgnoreDbs *string `groups:"create,update" json:"ignore_dbs,omitempty"`

	// +kubebuilder:validation:Enum="dump";"replication"
	// The migration method to be used (currently supported only by Redis and MySQL service types)
	Method *string `groups:"create,update" json:"method,omitempty"`

	// +kubebuilder:validation:MaxLength=256
	// Password for authentication with the server where to migrate data from
	Password *string `groups:"create,update" json:"password,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// Port number of the server where to migrate data from
	Port int `groups:"create,update" json:"port"`

	// The server where to migrate data from is secured with SSL
	Ssl *bool `groups:"create,update" json:"ssl,omitempty"`

	// +kubebuilder:validation:MaxLength=256
	// User name for authentication with the server where to migrate data from
	Username *string `groups:"create,update" json:"username,omitempty"`
}

// postgresql.conf configuration values
type Pg struct {
	// Specifies a fraction of the table size to add to autovacuum_analyze_threshold when deciding whether to trigger an ANALYZE. The default is 0.2 (20% of table size)
	AutovacuumAnalyzeScaleFactor *float64 `groups:"create,update" json:"autovacuum_analyze_scale_factor,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=2147483647
	// Specifies the minimum number of inserted, updated or deleted tuples needed to trigger an  ANALYZE in any one table. The default is 50 tuples.
	AutovacuumAnalyzeThreshold *int `groups:"create,update" json:"autovacuum_analyze_threshold,omitempty"`

	// +kubebuilder:validation:Minimum=200000000
	// +kubebuilder:validation:Maximum=1500000000
	// Specifies the maximum age (in transactions) that a table's pg_class.relfrozenxid field can attain before a VACUUM operation is forced to prevent transaction ID wraparound within the table. Note that the system will launch autovacuum processes to prevent wraparound even when autovacuum is otherwise disabled. This parameter will cause the server to be restarted.
	AutovacuumFreezeMaxAge *int `groups:"create,update" json:"autovacuum_freeze_max_age,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=20
	// Specifies the maximum number of autovacuum processes (other than the autovacuum launcher) that may be running at any one time. The default is three. This parameter can only be set at server start.
	AutovacuumMaxWorkers *int `groups:"create,update" json:"autovacuum_max_workers,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=86400
	// Specifies the minimum delay between autovacuum runs on any given database. The delay is measured in seconds, and the default is one minute
	AutovacuumNaptime *int `groups:"create,update" json:"autovacuum_naptime,omitempty"`

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=100
	// Specifies the cost delay value that will be used in automatic VACUUM operations. If -1 is specified, the regular vacuum_cost_delay value will be used. The default value is 20 milliseconds
	AutovacuumVacuumCostDelay *int `groups:"create,update" json:"autovacuum_vacuum_cost_delay,omitempty"`

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=10000
	// Specifies the cost limit value that will be used in automatic VACUUM operations. If -1 is specified (which is the default), the regular vacuum_cost_limit value will be used.
	AutovacuumVacuumCostLimit *int `groups:"create,update" json:"autovacuum_vacuum_cost_limit,omitempty"`

	// Specifies a fraction of the table size to add to autovacuum_vacuum_threshold when deciding whether to trigger a VACUUM. The default is 0.2 (20% of table size)
	AutovacuumVacuumScaleFactor *float64 `groups:"create,update" json:"autovacuum_vacuum_scale_factor,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=2147483647
	// Specifies the minimum number of updated or deleted tuples needed to trigger a VACUUM in any one table. The default is 50 tuples
	AutovacuumVacuumThreshold *int `groups:"create,update" json:"autovacuum_vacuum_threshold,omitempty"`

	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=10000
	// Specifies the delay between activity rounds for the background writer in milliseconds. Default is 200.
	BgwriterDelay *int `groups:"create,update" json:"bgwriter_delay,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=2048
	// Whenever more than bgwriter_flush_after bytes have been written by the background writer, attempt to force the OS to issue these writes to the underlying storage. Specified in kilobytes, default is 512. Setting of 0 disables forced writeback.
	BgwriterFlushAfter *int `groups:"create,update" json:"bgwriter_flush_after,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1073741823
	// In each round, no more than this many buffers will be written by the background writer. Setting this to zero disables background writing. Default is 100.
	BgwriterLruMaxpages *int `groups:"create,update" json:"bgwriter_lru_maxpages,omitempty"`

	// The average recent need for new buffers is multiplied by bgwriter_lru_multiplier to arrive at an estimate of the number that will be needed during the next round, (up to bgwriter_lru_maxpages). 1.0 represents a “just in time” policy of writing exactly the number of buffers predicted to be needed. Larger values provide some cushion against spikes in demand, while smaller values intentionally leave writes to be done by server processes. The default is 2.0.
	BgwriterLruMultiplier *float64 `groups:"create,update" json:"bgwriter_lru_multiplier,omitempty"`

	// +kubebuilder:validation:Minimum=500
	// +kubebuilder:validation:Maximum=1800000
	// This is the amount of time, in milliseconds, to wait on a lock before checking to see if there is a deadlock condition.
	DeadlockTimeout *int `groups:"create,update" json:"deadlock_timeout,omitempty"`

	// +kubebuilder:validation:Enum="lz4";"pglz"
	// Specifies the default TOAST compression method for values of compressible columns (the default is lz4).
	DefaultToastCompression *string `groups:"create,update" json:"default_toast_compression,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=604800000
	// Time out sessions with open transactions after this number of milliseconds
	IdleInTransactionSessionTimeout *int `groups:"create,update" json:"idle_in_transaction_session_timeout,omitempty"`

	// Controls system-wide use of Just-in-Time Compilation (JIT).
	Jit *bool `groups:"create,update" json:"jit,omitempty"`

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=2147483647
	// Causes each action executed by autovacuum to be logged if it ran for at least the specified number of milliseconds. Setting this to zero logs all autovacuum actions. Minus-one (the default) disables logging autovacuum actions.
	LogAutovacuumMinDuration *int `groups:"create,update" json:"log_autovacuum_min_duration,omitempty"`

	// +kubebuilder:validation:Enum="TERSE";"DEFAULT";"VERBOSE"
	// Controls the amount of detail written in the server log for each message that is logged.
	LogErrorVerbosity *string `groups:"create,update" json:"log_error_verbosity,omitempty"`

	// +kubebuilder:validation:Enum="'pid=%p,user=%u,db=%d,app=%a,client=%h '";"'%t [%p]: [%l-1] user=%u,db=%d,app=%a,client=%h '";"'%m [%p] %q[user=%u,db=%d,app=%a] '"
	// Choose from one of the available log-formats. These can support popular log analyzers like pgbadger, pganalyze etc.
	LogLinePrefix *string `groups:"create,update" json:"log_line_prefix,omitempty"`

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=86400000
	// Log statements that take more than this number of milliseconds to run, -1 disables
	LogMinDurationStatement *int `groups:"create,update" json:"log_min_duration_statement,omitempty"`

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=2147483647
	// Log statements for each temporary file created larger than this number of kilobytes, -1 disables
	LogTempFiles *int `groups:"create,update" json:"log_temp_files,omitempty"`

	// +kubebuilder:validation:Minimum=1000
	// +kubebuilder:validation:Maximum=4096
	// PostgreSQL maximum number of files that can be open per process
	MaxFilesPerProcess *int `groups:"create,update" json:"max_files_per_process,omitempty"`

	// +kubebuilder:validation:Minimum=64
	// +kubebuilder:validation:Maximum=6400
	// PostgreSQL maximum locks per transaction
	MaxLocksPerTransaction *int `groups:"create,update" json:"max_locks_per_transaction,omitempty"`

	// +kubebuilder:validation:Minimum=4
	// +kubebuilder:validation:Maximum=64
	// PostgreSQL maximum logical replication workers (taken from the pool of max_parallel_workers)
	MaxLogicalReplicationWorkers *int `groups:"create,update" json:"max_logical_replication_workers,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=96
	// Sets the maximum number of workers that the system can support for parallel queries
	MaxParallelWorkers *int `groups:"create,update" json:"max_parallel_workers,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=96
	// Sets the maximum number of workers that can be started by a single Gather or Gather Merge node
	MaxParallelWorkersPerGather *int `groups:"create,update" json:"max_parallel_workers_per_gather,omitempty"`

	// +kubebuilder:validation:Minimum=64
	// +kubebuilder:validation:Maximum=5120
	// PostgreSQL maximum predicate locks per transaction
	MaxPredLocksPerTransaction *int `groups:"create,update" json:"max_pred_locks_per_transaction,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=10000
	// PostgreSQL maximum prepared transactions
	MaxPreparedTransactions *int `groups:"create,update" json:"max_prepared_transactions,omitempty"`

	// +kubebuilder:validation:Minimum=8
	// +kubebuilder:validation:Maximum=64
	// PostgreSQL maximum replication slots
	MaxReplicationSlots *int `groups:"create,update" json:"max_replication_slots,omitempty"`

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=2147483647
	// PostgreSQL maximum WAL size (MB) reserved for replication slots. Default is -1 (unlimited). wal_keep_size minimum WAL size setting takes precedence over this.
	MaxSlotWalKeepSize *int `groups:"create,update" json:"max_slot_wal_keep_size,omitempty"`

	// +kubebuilder:validation:Minimum=2097152
	// +kubebuilder:validation:Maximum=6291456
	// Maximum depth of the stack in bytes
	MaxStackDepth *int `groups:"create,update" json:"max_stack_depth,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=43200000
	// Max standby archive delay in milliseconds
	MaxStandbyArchiveDelay *int `groups:"create,update" json:"max_standby_archive_delay,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=43200000
	// Max standby streaming delay in milliseconds
	MaxStandbyStreamingDelay *int `groups:"create,update" json:"max_standby_streaming_delay,omitempty"`

	// +kubebuilder:validation:Minimum=20
	// +kubebuilder:validation:Maximum=64
	// PostgreSQL maximum WAL senders
	MaxWalSenders *int `groups:"create,update" json:"max_wal_senders,omitempty"`

	// +kubebuilder:validation:Minimum=8
	// +kubebuilder:validation:Maximum=96
	// Sets the maximum number of background processes that the system can support
	MaxWorkerProcesses *int `groups:"create,update" json:"max_worker_processes,omitempty"`

	// +kubebuilder:validation:Minimum=3600
	// +kubebuilder:validation:Maximum=604800
	// Sets the time interval to run pg_partman's scheduled tasks
	PgPartmanBgwInterval *int `groups:"create,update" json:"pg_partman_bgw.interval,omitempty"`

	// +kubebuilder:validation:MaxLength=64
	// +kubebuilder:validation:Pattern=`^[_A-Za-z0-9][-._A-Za-z0-9]{0,63}$`
	// Controls which role to use for pg_partman's scheduled background tasks.
	PgPartmanBgwRole *string `groups:"create,update" json:"pg_partman_bgw.role,omitempty"`

	// Enables or disables query plan monitoring
	PgStatMonitorPgsmEnableQueryPlan *bool `groups:"create,update" json:"pg_stat_monitor.pgsm_enable_query_plan,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=10
	// Sets the maximum number of buckets
	PgStatMonitorPgsmMaxBuckets *int `groups:"create,update" json:"pg_stat_monitor.pgsm_max_buckets,omitempty"`

	// +kubebuilder:validation:Enum="all";"top";"none"
	// Controls which statements are counted. Specify top to track top-level statements (those issued directly by clients), all to also track nested statements (such as statements invoked within functions), or none to disable statement statistics collection. The default value is top.
	PgStatStatementsTrack *string `groups:"create,update" json:"pg_stat_statements.track,omitempty"`

	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=2147483647
	// PostgreSQL temporary file limit in KiB, -1 for unlimited
	TempFileLimit *int `groups:"create,update" json:"temp_file_limit,omitempty"`

	// +kubebuilder:validation:MaxLength=64
	// PostgreSQL service timezone
	Timezone *string `groups:"create,update" json:"timezone,omitempty"`

	// +kubebuilder:validation:Minimum=1024
	// +kubebuilder:validation:Maximum=10240
	// Specifies the number of bytes reserved to track the currently executing command for each active session.
	TrackActivityQuerySize *int `groups:"create,update" json:"track_activity_query_size,omitempty"`

	// +kubebuilder:validation:Enum="off";"on"
	// Record commit time of transactions.
	TrackCommitTimestamp *string `groups:"create,update" json:"track_commit_timestamp,omitempty"`

	// +kubebuilder:validation:Enum="all";"pl";"none"
	// Enables tracking of function call counts and time used.
	TrackFunctions *string `groups:"create,update" json:"track_functions,omitempty"`

	// +kubebuilder:validation:Enum="off";"on"
	// Enables timing of database I/O calls. This parameter is off by default, because it will repeatedly query the operating system for the current time, which may cause significant overhead on some platforms.
	TrackIoTiming *string `groups:"create,update" json:"track_io_timing,omitempty"`

	// Terminate replication connections that are inactive for longer than this amount of time, in milliseconds. Setting this value to zero disables the timeout.
	WalSenderTimeout *int `groups:"create,update" json:"wal_sender_timeout,omitempty"`

	// +kubebuilder:validation:Minimum=10
	// +kubebuilder:validation:Maximum=200
	// WAL flush interval in milliseconds. Note that setting this value to lower than the default 200ms may negatively impact performance
	WalWriterDelay *int `groups:"create,update" json:"wal_writer_delay,omitempty"`
}

// PGBouncer connection pooling settings
type Pgbouncer struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=86400
	// If the automatically created database pools have been unused this many seconds, they are freed. If 0 then timeout is disabled. [seconds]
	AutodbIdleTimeout *int `groups:"create,update" json:"autodb_idle_timeout,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=2147483647
	// Do not allow more than this many server connections per database (regardless of user). Setting it to 0 means unlimited.
	AutodbMaxDbConnections *int `groups:"create,update" json:"autodb_max_db_connections,omitempty"`

	// +kubebuilder:validation:Enum="session";"transaction";"statement"
	// PGBouncer pool mode
	AutodbPoolMode *string `groups:"create,update" json:"autodb_pool_mode,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=10000
	// If non-zero then create automatically a pool of that size per user when a pool doesn't exist.
	AutodbPoolSize *int `groups:"create,update" json:"autodb_pool_size,omitempty"`

	// +kubebuilder:validation:MaxItems=32
	// List of parameters to ignore when given in startup packet
	IgnoreStartupParameters []string `groups:"create,update" json:"ignore_startup_parameters,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=10000
	// Add more server connections to pool if below this number. Improves behavior when usual load comes suddenly back after period of total inactivity. The value is effectively capped at the pool size.
	MinPoolSize *int `groups:"create,update" json:"min_pool_size,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=86400
	// If a server connection has been idle more than this many seconds it will be dropped. If 0 then timeout is disabled. [seconds]
	ServerIdleTimeout *int `groups:"create,update" json:"server_idle_timeout,omitempty"`

	// +kubebuilder:validation:Minimum=60
	// +kubebuilder:validation:Maximum=86400
	// The pooler will close an unused server connection that has been connected longer than this. [seconds]
	ServerLifetime *int `groups:"create,update" json:"server_lifetime,omitempty"`

	// Run server_reset_query (DISCARD ALL) in all pooling modes
	ServerResetQueryAlways *bool `groups:"create,update" json:"server_reset_query_always,omitempty"`
}

// PGLookout settings
type Pglookout struct {
	// +kubebuilder:validation:Minimum=10
	// Number of seconds of master unavailability before triggering database failover to standby
	MaxFailoverReplicationTimeLag *int `groups:"create,update" json:"max_failover_replication_time_lag,omitempty"`
}

// Allow access to selected service ports from private networks
type PrivateAccess struct {
	// Allow clients to connect to pg with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations
	Pg *bool `groups:"create,update" json:"pg,omitempty"`

	// Allow clients to connect to pgbouncer with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations
	Pgbouncer *bool `groups:"create,update" json:"pgbouncer,omitempty"`

	// Allow clients to connect to prometheus with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations
	Prometheus *bool `groups:"create,update" json:"prometheus,omitempty"`
}

// Allow access to selected service components through Privatelink
type PrivatelinkAccess struct {
	// Enable pg
	Pg *bool `groups:"create,update" json:"pg,omitempty"`

	// Enable pgbouncer
	Pgbouncer *bool `groups:"create,update" json:"pgbouncer,omitempty"`

	// Enable prometheus
	Prometheus *bool `groups:"create,update" json:"prometheus,omitempty"`
}

// Allow access to selected service ports from the public Internet
type PublicAccess struct {
	// Allow clients to connect to pg from the public internet for service nodes that are in a project VPC or another type of private network
	Pg *bool `groups:"create,update" json:"pg,omitempty"`

	// Allow clients to connect to pgbouncer from the public internet for service nodes that are in a project VPC or another type of private network
	Pgbouncer *bool `groups:"create,update" json:"pgbouncer,omitempty"`

	// Allow clients to connect to prometheus from the public internet for service nodes that are in a project VPC or another type of private network
	Prometheus *bool `groups:"create,update" json:"prometheus,omitempty"`
}

// TimescaleDB extension configuration values
type Timescaledb struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=4096
	// The number of background workers for timescaledb operations. You should configure this setting to the sum of your number of databases and the total number of concurrent background workers you want running at any given point in time.
	MaxBackgroundWorkers *int `groups:"create,update" json:"max_background_workers,omitempty"`
}
type PgTestUserConfig struct {
	// +kubebuilder:validation:MaxItems=1
	// Additional Cloud Regions for Backup Replication
	AdditionalBackupRegions []string `groups:"create,update" json:"additional_backup_regions,omitempty"`

	// +kubebuilder:validation:MinLength=8
	// +kubebuilder:validation:MaxLength=256
	// +kubebuilder:validation:Pattern=`^[a-zA-Z0-9-_]+$`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	// Custom password for admin user. Defaults to random string. This must be set only when a new service is being created.
	AdminPassword *string `groups:"create" json:"admin_password,omitempty"`

	// +kubebuilder:validation:MaxLength=64
	// +kubebuilder:validation:Pattern=`^[_A-Za-z0-9][-._A-Za-z0-9]{0,63}$`
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	// Custom username for admin user. This must be set only when a new service is being created.
	AdminUsername *string `groups:"create" json:"admin_username,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=23
	// The hour of day (in UTC) when backup for the service is started. New backup is only started if previous backup has already completed.
	BackupHour *int `groups:"create,update" json:"backup_hour,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=59
	// The minute of an hour when backup for the service is started. New backup is only started if previous backup has already completed.
	BackupMinute *int `groups:"create,update" json:"backup_minute,omitempty"`

	// Register AAAA DNS records for the service, and allow IPv6 packets to service ports
	EnableIpv6 *bool `groups:"create,update" json:"enable_ipv6,omitempty"`

	// +kubebuilder:validation:MaxItems=1024
	// Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'
	IpFilter []*IpFilter `groups:"create,update" json:"ip_filter,omitempty"`

	// Migrate data from existing server
	Migration *Migration `groups:"create,update" json:"migration,omitempty"`

	// postgresql.conf configuration values
	Pg *Pg `groups:"create,update" json:"pg,omitempty"`

	// Should the service which is being forked be a read replica (deprecated, use read_replica service integration instead).
	PgReadReplica *bool `groups:"create,update" json:"pg_read_replica,omitempty"`

	// +kubebuilder:validation:MaxLength=64
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	// Name of the PG Service from which to fork (deprecated, use service_to_fork_from). This has effect only when a new service is being created.
	PgServiceToForkFrom *string `groups:"create" json:"pg_service_to_fork_from,omitempty"`

	// Enable the pg_stat_monitor extension. Enabling this extension will cause the cluster to be restarted.When this extension is enabled, pg_stat_statements results for utility commands are unreliable
	PgStatMonitorEnable *bool `groups:"create,update" json:"pg_stat_monitor_enable,omitempty"`

	// +kubebuilder:validation:Enum="10";"11";"12";"13";"14"
	// PostgreSQL major version
	PgVersion *string `groups:"create,update" json:"pg_version,omitempty"`

	// PGBouncer connection pooling settings
	Pgbouncer *Pgbouncer `groups:"create,update" json:"pgbouncer,omitempty"`

	// PGLookout settings
	Pglookout *Pglookout `groups:"create,update" json:"pglookout,omitempty"`

	// Allow access to selected service ports from private networks
	PrivateAccess *PrivateAccess `groups:"create,update" json:"private_access,omitempty"`

	// Allow access to selected service components through Privatelink
	PrivatelinkAccess *PrivatelinkAccess `groups:"create,update" json:"privatelink_access,omitempty"`

	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	// Name of another project to fork a service from. This has effect only when a new service is being created.
	ProjectToForkFrom *string `groups:"create" json:"project_to_fork_from,omitempty"`

	// Allow access to selected service ports from the public Internet
	PublicAccess *PublicAccess `groups:"create,update" json:"public_access,omitempty"`

	// +kubebuilder:validation:MaxLength=32
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	// Recovery target time when forking a service. This has effect only when a new service is being created.
	RecoveryTargetTime *string `groups:"create" json:"recovery_target_time,omitempty"`

	// +kubebuilder:validation:MaxLength=64
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	// Name of another service to fork from. This has effect only when a new service is being created.
	ServiceToForkFrom *string `groups:"create" json:"service_to_fork_from,omitempty"`

	// Percentage of total RAM that the database server uses for shared memory buffers. Valid range is 20-60 (float), which corresponds to 20% - 60%. This setting adjusts the shared_buffers configuration value.
	SharedBuffersPercentage *float64 `groups:"create,update" json:"shared_buffers_percentage,omitempty"`

	// Use static public IP addresses
	StaticIps *bool `groups:"create,update" json:"static_ips,omitempty"`

	// +kubebuilder:validation:Enum="quorum";"off"
	// Synchronous replication type. Note that the service plan also needs to support synchronous replication.
	SynchronousReplication *string `groups:"create,update" json:"synchronous_replication,omitempty"`

	// TimescaleDB extension configuration values
	Timescaledb *Timescaledb `groups:"create,update" json:"timescaledb,omitempty"`

	// +kubebuilder:validation:Enum="aiven";"timescale"
	// Variant of the PostgreSQL service, may affect the features that are exposed by default
	Variant *string `groups:"create,update" json:"variant,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1024
	// Sets the maximum amount of memory to be used by a query operation (such as a sort or hash table) before writing to temporary disk files, in MB. Default is 1MB + 0.075% of total RAM (up to 32MB).
	WorkMem *int `groups:"create,update" json:"work_mem,omitempty"`
}
