// Code generated by user config generator. DO NOT EDIT.
// +kubebuilder:object:generate=true

package mysqluserconfig

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

// mysql.conf configuration values
type Mysql struct {
	// +kubebuilder:validation:Minimum=2
	// +kubebuilder:validation:Maximum=3600
	// The number of seconds that the mysqld server waits for a connect packet before responding with Bad handshake
	ConnectTimeout *int `groups:"create,update" json:"connect_timeout,omitempty"`

	// +kubebuilder:validation:MinLength=2
	// +kubebuilder:validation:MaxLength=100
	// Default server time zone as an offset from UTC (from -12:00 to +12:00), a time zone name, or 'SYSTEM' to use the MySQL server default.
	DefaultTimeZone *string `groups:"create,update" json:"default_time_zone,omitempty"`

	// +kubebuilder:validation:Minimum=4
	// The maximum permitted result length in bytes for the GROUP_CONCAT() function.
	GroupConcatMaxLen *int `groups:"create,update" json:"group_concat_max_len,omitempty"`

	// +kubebuilder:validation:Minimum=900
	// +kubebuilder:validation:Maximum=31536000
	// The time, in seconds, before cached statistics expire
	InformationSchemaStatsExpiry *int `groups:"create,update" json:"information_schema_stats_expiry,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=50
	// Maximum size for the InnoDB change buffer, as a percentage of the total size of the buffer pool. Default is 25
	InnodbChangeBufferMaxSize *int `groups:"create,update" json:"innodb_change_buffer_max_size,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=2
	// Specifies whether flushing a page from the InnoDB buffer pool also flushes other dirty pages in the same extent (default is 1): 0 - dirty pages in the same extent are not flushed,  1 - flush contiguous dirty pages in the same extent,  2 - flush dirty pages in the same extent
	InnodbFlushNeighbors *int `groups:"create,update" json:"innodb_flush_neighbors,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=16
	// Minimum length of words that are stored in an InnoDB FULLTEXT index. Changing this parameter will lead to a restart of the MySQL service.
	InnodbFtMinTokenSize *int `groups:"create,update" json:"innodb_ft_min_token_size,omitempty"`

	// +kubebuilder:validation:MaxLength=1024
	// +kubebuilder:validation:Pattern=`^.+/.+$`
	// This option is used to specify your own InnoDB FULLTEXT index stopword list for all InnoDB tables.
	InnodbFtServerStopwordTable *string `groups:"create,update" json:"innodb_ft_server_stopword_table,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3600
	// The length of time in seconds an InnoDB transaction waits for a row lock before giving up.
	InnodbLockWaitTimeout *int `groups:"create,update" json:"innodb_lock_wait_timeout,omitempty"`

	// +kubebuilder:validation:Minimum=1048576
	// +kubebuilder:validation:Maximum=4294967295
	// The size in bytes of the buffer that InnoDB uses to write to the log files on disk.
	InnodbLogBufferSize *int `groups:"create,update" json:"innodb_log_buffer_size,omitempty"`

	// +kubebuilder:validation:Minimum=65536
	// +kubebuilder:validation:Maximum=1099511627776
	// The upper limit in bytes on the size of the temporary log files used during online DDL operations for InnoDB tables.
	InnodbOnlineAlterLogMaxSize *int `groups:"create,update" json:"innodb_online_alter_log_max_size,omitempty"`

	// When enabled, information about all deadlocks in InnoDB user transactions is recorded in the error log. Disabled by default.
	InnodbPrintAllDeadlocks *bool `groups:"create,update" json:"innodb_print_all_deadlocks,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=64
	// The number of I/O threads for read operations in InnoDB. Default is 4. Changing this parameter will lead to a restart of the MySQL service.
	InnodbReadIoThreads *int `groups:"create,update" json:"innodb_read_io_threads,omitempty"`

	// When enabled a transaction timeout causes InnoDB to abort and roll back the entire transaction. Changing this parameter will lead to a restart of the MySQL service.
	InnodbRollbackOnTimeout *bool `groups:"create,update" json:"innodb_rollback_on_timeout,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1000
	// Defines the maximum number of threads permitted inside of InnoDB. Default is 0 (infinite concurrency - no limit)
	InnodbThreadConcurrency *int `groups:"create,update" json:"innodb_thread_concurrency,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=64
	// The number of I/O threads for write operations in InnoDB. Default is 4. Changing this parameter will lead to a restart of the MySQL service.
	InnodbWriteIoThreads *int `groups:"create,update" json:"innodb_write_io_threads,omitempty"`

	// +kubebuilder:validation:Minimum=30
	// +kubebuilder:validation:Maximum=604800
	// The number of seconds the server waits for activity on an interactive connection before closing it.
	InteractiveTimeout *int `groups:"create,update" json:"interactive_timeout,omitempty"`

	// +kubebuilder:validation:Enum="TempTable";"MEMORY"
	// The storage engine for in-memory internal temporary tables.
	InternalTmpMemStorageEngine *string `groups:"create,update" json:"internal_tmp_mem_storage_engine,omitempty"`

	// The slow_query_logs work as SQL statements that take more than long_query_time seconds to execute. Default is 10s
	LongQueryTime *float64 `groups:"create,update" json:"long_query_time,omitempty"`

	// +kubebuilder:validation:Minimum=102400
	// +kubebuilder:validation:Maximum=1073741824
	// Size of the largest message in bytes that can be received by the server. Default is 67108864 (64M)
	MaxAllowedPacket *int `groups:"create,update" json:"max_allowed_packet,omitempty"`

	// +kubebuilder:validation:Minimum=1048576
	// +kubebuilder:validation:Maximum=1073741824
	// Limits the size of internal in-memory tables. Also set tmp_table_size. Default is 16777216 (16M)
	MaxHeapTableSize *int `groups:"create,update" json:"max_heap_table_size,omitempty"`

	// +kubebuilder:validation:Minimum=1024
	// +kubebuilder:validation:Maximum=1048576
	// Start sizes of connection buffer and result buffer. Default is 16384 (16K). Changing this parameter will lead to a restart of the MySQL service.
	NetBufferLength *int `groups:"create,update" json:"net_buffer_length,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3600
	// The number of seconds to wait for more data from a connection before aborting the read.
	NetReadTimeout *int `groups:"create,update" json:"net_read_timeout,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3600
	// The number of seconds to wait for a block to be written to a connection before aborting the write.
	NetWriteTimeout *int `groups:"create,update" json:"net_write_timeout,omitempty"`

	// Slow query log enables capturing of slow queries. Setting slow_query_log to false also truncates the mysql.slow_log table. Default is off
	SlowQueryLog *bool `groups:"create,update" json:"slow_query_log,omitempty"`

	// +kubebuilder:validation:Minimum=32768
	// +kubebuilder:validation:Maximum=1073741824
	// Sort buffer size in bytes for ORDER BY optimization. Default is 262144 (256K)
	SortBufferSize *int `groups:"create,update" json:"sort_buffer_size,omitempty"`

	// +kubebuilder:validation:MaxLength=1024
	// +kubebuilder:validation:Pattern=`^[A-Z_]*(,[A-Z_]+)*$`
	// Global SQL mode. Set to empty to use MySQL server defaults. When creating a new service and not setting this field Aiven default SQL mode (strict, SQL standard compliant) will be assigned.
	SqlMode *string `groups:"create,update" json:"sql_mode,omitempty"`

	// Require primary key to be defined for new tables or old tables modified with ALTER TABLE and fail if missing. It is recommended to always have primary keys because various functionality may break if any large table is missing them.
	SqlRequirePrimaryKey *bool `groups:"create,update" json:"sql_require_primary_key,omitempty"`

	// +kubebuilder:validation:Minimum=1048576
	// +kubebuilder:validation:Maximum=1073741824
	// Limits the size of internal in-memory tables. Also set max_heap_table_size. Default is 16777216 (16M)
	TmpTableSize *int `groups:"create,update" json:"tmp_table_size,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2147483
	// The number of seconds the server waits for activity on a noninteractive connection before closing it.
	WaitTimeout *int `groups:"create,update" json:"wait_timeout,omitempty"`
}

// Allow access to selected service ports from private networks
type PrivateAccess struct {
	// Allow clients to connect to mysql with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations
	Mysql *bool `groups:"create,update" json:"mysql,omitempty"`

	// Allow clients to connect to mysqlx with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations
	Mysqlx *bool `groups:"create,update" json:"mysqlx,omitempty"`

	// Allow clients to connect to prometheus with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations
	Prometheus *bool `groups:"create,update" json:"prometheus,omitempty"`
}

// Allow access to selected service components through Privatelink
type PrivatelinkAccess struct {
	// Enable mysql
	Mysql *bool `groups:"create,update" json:"mysql,omitempty"`

	// Enable mysqlx
	Mysqlx *bool `groups:"create,update" json:"mysqlx,omitempty"`

	// Enable prometheus
	Prometheus *bool `groups:"create,update" json:"prometheus,omitempty"`
}

// Allow access to selected service ports from the public Internet
type PublicAccess struct {
	// Allow clients to connect to mysql from the public internet for service nodes that are in a project VPC or another type of private network
	Mysql *bool `groups:"create,update" json:"mysql,omitempty"`

	// Allow clients to connect to mysqlx from the public internet for service nodes that are in a project VPC or another type of private network
	Mysqlx *bool `groups:"create,update" json:"mysqlx,omitempty"`

	// Allow clients to connect to prometheus from the public internet for service nodes that are in a project VPC or another type of private network
	Prometheus *bool `groups:"create,update" json:"prometheus,omitempty"`
}
type MysqlUserConfig struct {
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

	// +kubebuilder:validation:Minimum=600
	// +kubebuilder:validation:Maximum=86400
	// The minimum amount of time in seconds to keep binlog entries before deletion. This may be extended for services that require binlog entries for longer than the default for example if using the MySQL Debezium Kafka connector.
	BinlogRetentionPeriod *int `groups:"create,update" json:"binlog_retention_period,omitempty"`

	// +kubebuilder:validation:MaxItems=1024
	// Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'
	IpFilter []*IpFilter `groups:"create,update" json:"ip_filter,omitempty"`

	// Migrate data from existing server
	Migration *Migration `groups:"create,update" json:"migration,omitempty"`

	// mysql.conf configuration values
	Mysql *Mysql `groups:"create,update" json:"mysql,omitempty"`

	// +kubebuilder:validation:Enum="8"
	// MySQL major version
	MysqlVersion *string `groups:"create,update" json:"mysql_version,omitempty"`

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

	// Use static public IP addresses
	StaticIps *bool `groups:"create,update" json:"static_ips,omitempty"`
}
