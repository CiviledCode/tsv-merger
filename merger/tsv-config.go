package merger

// TSVConfig represents a list of configuration values in order to merge data efficiently.
type TSVConfig struct {
	// Seperator represents the character seperating values. 
	Seperator byte `json:"seperator"`

	// KeyColumn is the column we are using to uniquely identify rows.
	KeyColumn string `json:"key_column"`

	// IncrementalColumn is the column we are adding into the key within the map holding values.
	IncrementalColumn string `json:"incremental_column"`

	// ArithmeticIncrementation depicts if we should be adding rather than appending.
	ArithmeticIncrementation bool `json:"is_arithmetic"`

	// Files are the files we are importing data from.
	Files []string `json:"files"`
}
