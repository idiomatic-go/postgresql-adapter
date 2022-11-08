package queryv1

import (
	v1 "github.com/idiomatic-go/core-types/corev1"
)

const (
	sloEntryInsert       = "INSERT INTO slo_entry ( id,customer_id,name,application,route_name,category,traffic_type,traffic_protocol,processing_interval,window_interval,watch_percent,threshold_percent,threshold_value,threshold_minimum,locality_scope,rps_low_comparison,rps_high_comparison, filter_status_codes,status_codes,metric_name,metric_name_secondary,function_name,disable_processing,disable_triage,created_ts) VALUES"
	sloEntryDelete       = "DELETE FROM slo_entry WHERE id = $1"
	sloEntryDeleteByName = "DELETE FROM slo_entry WHERE name = $1"
)

func writeSLOEntryInsert(entry *v1.SLOEntry) string {
	if entry == nil {
		return ""
	}
	return ""
}
