package tsv

type Row struct {
	Key string

	// TODO: Make this a map instead.
	Value map[string]interface{}
}
