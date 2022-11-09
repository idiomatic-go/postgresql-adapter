package dml

const (
	GetCustomerFn = "SELECT * FROM GetCustomer($1)"
	//CustomerInsertStmt       = "INSERT INTO customer (nextval('customer_id'),org_activity,track_activity,created_ts) VALUES"

	CustomerInsertStmt = "INSERT INTO customer (id,org_activity,track_activity,created_ts) VALUES"
	CustomerDeleteStmt = "DELETE FROM customer WHERE id = $1"
	//EntryDeleteByNameStmt = "DELETE FROM slo WHERE id = $1 AND name = $2"

	OrgId         = "org_id"
	TrackActivity = "track_activity"
)
