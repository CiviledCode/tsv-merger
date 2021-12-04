package operations

// constantCallback requires one initial value and keeps that value the same until output.
// This uses the first instance of the value out of any of the files.
func constantCallback(oldValue, _ interface{}) interface{} {
	return oldValue
}
