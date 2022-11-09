package slo

const (
	GetSLOEntryFn = "GetSLOEntry"
	// nextval('slo_entry_id')
	EntryInsertStmt       = "INSERT INTO slo (id,customer_id,category,traffic_type,traffic_protocol,processing_interval,window_interval,watch_percent,threshold_percent,threshold_value,threshold_minimum,rps_low_comparison,rps_high_comparison,locality_scope,disable_processing,disable_triage,name,application,route_name,filter_status_codes,status_codes,created_ts) VALUES"
	EntryDeleteStmt       = "DELETE FROM slo WHERE id = $1"
	EntryDeleteByNameStmt = "DELETE FROM slo WHERE id = $1 AND name = $2"

	Id              = "id"
	CustomerId      = "customer_id"
	Category        = "category"
	TrafficType     = "traffic_type"
	TrafficProtocol = "traffic_protocol"

	ProcessingInterval = "processing_window"
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
	StatusCodes       = "statusCodes"
	Name              = "name"
	Application       = "application"
	RouteName         = "route_name"

	CreatedTS = "created_ts"
	ChangedTS = "changed_ts"
)
