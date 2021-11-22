package merger

// MergerType represents a proper
type MergerType interface {
	// Output receives a file path and writes the values within itself to the file.
	Output(string) error

	// Put receives a key and increments the value at that key accordingly.
	Put(string, interface{})
}
