package pgxsql

func ScanInt16(pgxInt any) int16 {
	if pgxInt != nil {
		if i, ok := pgxInt.(int32); ok {
			return int16(i)
		}
	}
	return 0
}

func ScanInt32(pgxInt any) int32 {
	if pgxInt != nil {
		if i, ok := pgxInt.(int32); ok {
			return i
		}
	}
	return 0
}
