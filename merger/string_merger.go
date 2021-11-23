package merger

// StringMerger uses string appending for values to allow for piecing long strings of data together.
// This implements the Merger interface.
type StringMerger struct {
	// Config gives us configuration values for key column, incremental column, and seperator.
	Config TSVConfig

	// TODO: Make this a double map so we can store multiple columns per key or hash the column and key together and store within the map.
	
	// Content holds the string values for each row.
	Content map[string]string
}

// Put ...
func (s StringMerger) Put(key string, value interface{}) {
	s.Content[key] += value.(string)
}

// Output ...
func (s StringMerger) Output(file string) error {
	
	// TODO: Fill this out.

	return nil
}
