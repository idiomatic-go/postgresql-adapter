package pgxsql

// DATABASE_URL=postgres://{user}:{password}@{hostname}:{port}/{database-name}
// psql -x "postgres://tsdbadmin@t9aggksc24.gspnhi29bv.tsdb.cloud.timescale.com:33251/tsdb?sslmode=require"
// Password for user tsdbadmin:
const (
	ConfigFileName = "postgresql.txt"
	DatabaseURLKey = "DATABASE_URL"
)
