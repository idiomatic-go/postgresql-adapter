package dml

const (
	GetSLOEntryStmt       = "SELECT * FROM GetSLOEntry($1)"
	GetSLOEntryByNameStmt = "SELECT * FROM GetSLOEntryByName($1,$2)"
	GetSLOEntryBySegment  = "SELECT * FROM GetSLOEntryBySegment($1,$2)"

	// nextval('slo_entry_id')
	EntryInsertStmt       = "INSERT INTO slo (id,customer_id,category,traffic_type,traffic_protocol,processing_interval,window_interval,watch_percent,threshold_percent,threshold_value,threshold_minimum,rps_low_comparison,rps_high_comparison,locality_scope,disable_processing,disable_triage,name,application,route_name,filter_status_codes,status_codes) VALUES"
	EntryDeleteStmt       = "DELETE FROM slo WHERE id = $1"
	EntryDeleteByNameStmt = "DELETE FROM slo WHERE id = $1 AND name = $2"

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
)
