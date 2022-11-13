package dml

const (
	Id         = "id"
	CustomerId = "customer_id"
	CreatedTS  = "created_ts"
	ChangedTS  = "changed_ts"

	ChangedTimestampFn = "now()"
)

const (
	GetSLOEntryStmt          = "SELECT * FROM GetSLOEntry($1)"
	GetSLOEntryByNameStmt    = "SELECT * FROM GetSLOEntryByName($1,$2)"
	GetSLOEntryBySegmentStmt = "SELECT * FROM GetSLOEntryBySegment($1,$2)"
	GetSLOEntryLogStmt       = "SELECT * FROM GetSLOEntryLog($1,$2)"

	InsertSLOEntryStmt       = "INSERT INTO slo_entry (id,customer_id,category,traffic_type,traffic_protocol,processing_interval,window_interval,watch_percent,threshold_percent,threshold_value,threshold_minimum,rps_low_comparison,rps_high_comparison,locality_scope,disable_processing,disable_triage,name,application,route_name,filter_status_codes,status_codes) VALUES"
	UpdateSLOEntryStmt       = "UPDATE slo_entry"
	DeleteSLOEntryByNameStmt = "DELETE FROM slo_entry WHERE customerId = $1 AND name = $2"

	SLOEntryNextValFn = "nextval('slo_entry_Id')"

	Category        = "category"
	TrafficType     = "traffic_type"
	TrafficProtocol = "traffic_protocol"

	ProcessingInterval = "processing_interval"
	WindowInterval     = "window_interval"
	WatchPercent       = "watch_percent"
	ThresholdPercent   = "threshold_percent"
	ThresholdValue     = "threshold_value"
	ThresholdMinimum   = "threshold_minimum"

	RPSLowComparison  = "rps_low_comparison"
	RPSHighComparison = "rps_high_comparison"

	LocalityScope     = "locality_scope"
	DisableProcessing = "disable_processing"
	DisableTriage     = "disable_triage"

	FilterStatusCodes = "filter_status_codes"
	StatusCodes       = "status_codes"
	Name              = "name"
	Application       = "application"
	RouteName         = "route_name"
)

const (
	GetCustomerStmt      = "SELECT * FROM GetCustomer($1)"
	GetCustomerByOrgStmt = "SELECT * FROM GetCustomerByOrg($1,$2)"

	InsertCustomerStmt = "INSERT INTO customer (id,org_activity,track_activity) VALUES"
	UpdateCustomerStmt = "UPDATE customer $1 WHERE id = $2 AND name = $3"
	DeleteCustomerStmt = "DELETE FROM customer WHERE id = $1"

	InsertCustomerLocalityStmt = "INSERT INTO customer_locality (id,customer_id,region,zone,sub_zone) VALUES"
	DeleteCustomerLocalityStmt = "DELETE FROM customer_locality WHERE id = $1"
	//UpdateCustomerStmt = "UPDATE customer $1 WHERE id = $2 AND name = $3"

	InsertCustomerMetricStmt = "INSERT INTO customer_metric (id,customer_id,application,route_name,region,zone,sub_zone,name,value) VALUES"
	UpdateCustomerMetricStmt = "UPDATE customer_metric $1 WHERE id = $2 AND name = $3"
	//DeleteCustomerMetricStmt = "DELETE FROM customer_locality WHERE id = $1"

	SequenceCustomerId         = "customer_id"
	SequenceCustomerLocalityId = "customer_locality_id"
	SequenceCustomerMetricId   = "customer_metric_id"

	OrgId         = "org_id"
	TrackActivity = "track_activity"
)
