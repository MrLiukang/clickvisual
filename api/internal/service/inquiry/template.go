package inquiry

var clickhouseTableDataORM = map[int]string{
	TableTypeApp: `create table %s
(
	_timestamp_ DateTime,
	_trace_time_ DateTime64(9, 'Asia/Shanghai'),
	_source_ String,
	_cluster_ String,
	_log_agent_ String,
	_namespace_ String,
	_node_name_ String,
	_node_ip_ String,
	_container_name_ String,
	_pod_name_ String,
	_raw_log_ String
)
engine = MergeTree PARTITION BY toYYYYMMDD(_timestamp_)
ORDER BY _timestamp_
TTL toDateTime(_timestamp_) + INTERVAL %d DAY 
SETTINGS index_granularity = 8192;`,
	TableTypeEgo: `create table %s
(
	_timestamp_ DateTime,
	_trace_time_ DateTime64(9, 'Asia/Shanghai'),
	_source_ String,
	_cluster_ String,
	_log_agent_ String,
	_namespace_ String,
	_node_name_ String,
	_node_ip_ String,
	_container_name_ String,
	_pod_name_ String,
	_raw_log_ String,
	lv Nullable(String),
	msg Nullable(String),
	compName Nullable(String),
	cost Nullable(Float64),
	method Nullable(String),
	code Nullable(Int64),
	ucode Nullable(Int64),
	peerName Nullable(String),
	peerIp Nullable(String),
	type Nullable(String),
	tid Nullable(String)
)
engine = MergeTree PARTITION BY toYYYYMMDD(_timestamp_)
ORDER BY _timestamp_ 
TTL toDateTime(_timestamp_) + INTERVAL %d DAY
SETTINGS index_granularity = 8192;`,
	TableTypeIngress: `create table %s
(
	_timestamp_ DateTime,
	_trace_time_ DateTime64(9, 'Asia/Shanghai'),
	_cluster_ String,
	_log_agent_ String,
	_namespace_ String,
	_node_name_ String,
	_node_ip_ String,
	_container_name_ String,
	_pod_name_ String,
	client_ip String,
	x_forward_for String,
	method String,
	url String,
	version String,
	status UInt16,
	body_bytes_sent UInt32,
	http_referer String,
	http_user_agent String,
	request_length UInt32,
	request_time Float32,
	proxy_upstream_name String,
	upstream_addr String,
	upstream_response_length String,
	upstream_response_time String,
	upstream_status String,
	req_id String,
	host String
)
engine = MergeTree PARTITION BY toYYYYMMDD(_timestamp_)
ORDER BY _timestamp_
TTL toDateTime(_timestamp_) + INTERVAL %d DAY
SETTINGS index_granularity = 8192;`,
}

var clickhouseTableStreamORM = map[int]string{
	TableTypeApp: `create table %s
(
	_source_ String,
	_time_ String,
	_pod_name_ String,
	_namespace_ String,
	_node_name_ String,
	_container_name_ String,
	_cluster_ String,
	_log_agent_ String,
	_node_ip_ String,
	log String
)
engine = Kafka SETTINGS kafka_broker_list = '%s', kafka_topic_list = '%s', kafka_group_name = '%s', kafka_format = 'JSONEachRow', kafka_num_consumers = 1;`,
	TableTypeEgo: `create table  %s
(
	_source_ String,
	_time_ String,
	_pod_name_ String,
	_namespace_ String,
	_node_name_ String,
	_container_name_ String,
	_cluster_ String,
	_log_agent_ String,
	_node_ip_ String,
	log String
)
engine = Kafka SETTINGS kafka_broker_list = '%s', kafka_topic_list = '%s', kafka_group_name = '%s', kafka_format = 'JSONEachRow', kafka_num_consumers = 1;`,
	TableTypeIngress: `create table if not exists %s
(
	_pod_name_ String,
	_namespace_ String,
	_node_name_ String,
	_container_name_ String,
	_cluster_ String,
	_log_agent_ String,
	_node_ip_ String,
	_source_ String,
	_time_ String,
	client_ip String,
	x_forward_for String,
	method String,
	url String,
	version String,
	status UInt16,
	body_bytes_sent UInt32,
	http_referer String,
	http_user_agent String,
	request_length UInt32,
	request_time Float32,
	proxy_upstream_name String,
	upstream_addr String,
	upstream_response_length String,
	upstream_response_time String,
	upstream_status String,
	req_id String,
	host String
)
engine = Kafka SETTINGS kafka_broker_list = '%s', kafka_topic_list = '%s', kafka_group_name = '%s', kafka_format = 'JSONEachRow', kafka_num_consumers = 1;`,
}

var clickhouseViewORM = map[int]string{
	TableTypeApp: `CREATE MATERIALIZED VIEW %s TO %s AS
SELECT
    %s,
    _source_,
    _cluster_,
    _log_agent_,
    _namespace_,
    _node_name_,
    _node_ip_,
    _container_name_,
    _pod_name_,
	log AS _raw_log_%s
	FROM %s where %s;`,
	TableTypeEgo: `CREATE MATERIALIZED VIEW %s TO %s AS
SELECT
    %s,
	_source_,
	_pod_name_,
	_namespace_,
	_node_name_,
	_container_name_,
	_cluster_,
	_log_agent_,
	_node_ip_,
	visitParamExtractString(log, 'lv') AS lv,
	visitParamExtractString(log, 'msg') AS msg,
	visitParamExtractString(log, 'compName') AS compName,
	visitParamExtractFloat(log, 'cost') AS cost,
	visitParamExtractString(log, 'method') AS method,
	visitParamExtractInt(log, 'code') AS code,
	visitParamExtractInt(log, 'ucode') AS ucode,
	visitParamExtractString(log, 'peerName') AS peerName,
	visitParamExtractString(log, 'peerIp') AS peerIp,
	visitParamExtractString(log, 'type') AS type,
	visitParamExtractString(log, 'tid') AS tid,
	log AS _raw_log_%s
	FROM %s where %s;`,
	TableTypeIngress: `CREATE MATERIALIZED VIEW %s TO %s AS
SELECT
    %s,
	_pod_name_,
	_namespace_,
	_node_name_,
	_container_name_,
	_cluster_,
	_log_agent_,
	_node_ip_,
	_source_,
	client_ip,
	x_forward_for,
	method,
	url,
	version,
	status,
	body_bytes_sent,
	http_referer,
	http_user_agent,
	request_length,
	request_time,
	proxy_upstream_name,
	upstream_addr,
	upstream_response_length,
	upstream_response_time,
	upstream_status,
	req_id,
	host%s
	FROM %s where %s;`,
}
