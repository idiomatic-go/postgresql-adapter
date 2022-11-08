package queryv1

import (
	v1 "github.com/idiomatic-go/core-types/corev1"
)

const (
	sloEntryInsert       = "INSERT INTO slo_entry (id,customer_id,category,traffic_type,traffic_protocol,processing_interval,window_interval,watch_percent,threshold_percent,threshold_value,threshold_minimum,rps_low_comparison,rps_high_comparison,locality_scope,disable_processing,disable_triage,name,application,route_name,filter_status_codes,status_codes,created_ts) VALUES"
	sloEntryDelete       = "DELETE FROM slo_entry WHERE id = $1"
	sloEntryDeleteByName = "DELETE FROM slo_entry WHERE id = $1 AND name = $2"

	Id              = "Id"
	CustomerId      = "customerId"
	Category        = "Category"
	TrafficType     = "TrafficType"
	TrafficProtocol = "TrafficProtocol"

	ProcessingInterval = "ProcessingWindow"
	WindowInterval     = "WindowInterval"
	WatchPercent       = "WatchPercent"
	ThresholdPercent   = "ThresholdPercent"
	ThresholdValue     = "ThresholdValue"
	ThresholdMinimum   = "ThresholdMinimum"

	RPSLowComparison  = "RPSLowComparison"
	RPSHighComparison = "RPSHighComparison"

	LocalityScope     = "LocalityScope"
	DisableProcessing = "DisableProcessing"
	DisableTriage     = "DisableTriage"

	FilterStatusCodes = "FilterStatusCodes"
	StatusCodes       = "StatusCodes"
	Name              = "Name"
	Application       = "Application"
	RouteName         = "RouteName"
)

var (
	sloEntryToColumnName = map[string]string{
		Id:              "id",
		CustomerId:      "customer_id",
		Category:        "category",
		TrafficType:     "traffic_type",
		TrafficProtocol: "traffic_protocol",

		ProcessingInterval: "processing_interval",
		WindowInterval:     "window_interval",
		WatchPercent:       "watch_percent",
		ThresholdPercent:   "threshold_percent",
		ThresholdValue:     "threshold_value",
		ThresholdMinimum:   "threshold_minimum",

		RPSLowComparison:  "rps_low_comparison",
		RPSHighComparison: "rps_high_comparison",

		LocalityScope:     "locality_scope",
		DisableProcessing: "disable_processing",
		DisableTriage:     "disable_triage",

		FilterStatusCodes: "filter_status_codes",
		StatusCodes:       "status_codes",
		Name:              "name",
		Application:       "application",
		RouteName:         "route_name",
	}
)

func writeSLOEntryInsert(entry *v1.SLOEntry) string {
	if entry == nil {
		return ""
	}
	return ""
}
