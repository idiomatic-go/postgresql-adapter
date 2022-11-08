package pgxsql

func ScanInt16(pgsqlInt any) int16 {
	if pgsqlInt != nil {
		if i, ok := pgsqlInt.(int32); ok {
			return int16(i)
		}
	}
	return 0
}

func ScanInt32(pgsqlInt any) int32 {
	if pgsqlInt != nil {
		if i, ok := pgsqlInt.(int32); ok {
			return i
		}
	}
	return 0
}
