package customer

const (
	GetCustomerFn = "GetCustomer"
	//CustomerInsertStmt       = "INSERT INTO customer (nextval('customer_id'),org_activity,track_activity,created_ts) VALUES"

	CustomerInsertStmt = "INSERT INTO customer (id,org_activity,track_activity,created_ts) VALUES"
	CustomerDeleteStmt = "DELETE FROM customer WHERE id = $1"
	//EntryDeleteByNameStmt = "DELETE FROM slo WHERE id = $1 AND name = $2"

	Id            = "id"
	OrgId         = "org_id"
	TrackActivity = "track_activity"
	CreatedTS     = "created_ts"
	ChangedTS     = "changed_ts"
)
