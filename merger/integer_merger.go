package merger

import (
	"fmt"
	"os"
	"strconv"
)

// IntegerMerger ...
type IntegerMerger struct {
	Config TSVConfig
	
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
