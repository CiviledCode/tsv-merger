package tsv

// Row contains values according to a row within a TSV file.
type Row struct {
	// Key is the identifying value for that row.
	// This is commonly like a name or ID.
	Key string

	// Value represents other values stored within columns inside of the row.
	Value map[string]string
}
