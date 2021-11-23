package merger

import (
	"fmt"
	"os"
	"strconv"
)

// IntegerMerger uses integer arithmetic instructions to increment values.
// This implements the Merger interface.
type IntegerMerger struct {
	// Config allows us to access key column, seperators, and incremental columns.
	Config TSVConfig

	// TODO: Make this a double map so we can store multiple columns per key or hash the column and key together and store within the map.

	// Content is the content we are merging.
	Content map[string]int
}

// Put ...
func (i IntegerMerger) Put(key string, value interface{}) {
	val, err := strconv.ParseInt(value.(string), 10, 64)
	if err != nil {
		panic(err)
	}

	i.Content[key] += int(val)
}

// Output ...
func (i IntegerMerger) Output(file string) error {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	fmt.Fprintf(f, "%v%c%v\n", i.Config.KeyColumn, i.Config.Seperator, i.Config.IncrementalColumn)
	for key, value := range i.Content {
		fmt.Fprintf(f, "%v%c%v\n", key, i.Config.Seperator, value)	
	}

	return nil
}
