package merger

// TSVConfig represents a list of configuration values in order to merge data efficiently.
type TSVConfig struct {
	// Seperator represents the text seperating values. 
	Seperator string `json:"seperator"`

	// OutputSeperator represents the text seperating values within the output. 
	OutputSeperator string `json:"output_seperator"`

	// KeyColumn is the column we are using to uniquely identify rows.
	KeyColumn string `json:"key_column"`

	// ColumnOperations maps column names to operation bytes. These operation bytes depict how this column should be merged.
	ColumnOperations map[string]string `json:"column_operations"`

	// Files are the files we are importing data from.
	Files []string `json:"files"`
}
