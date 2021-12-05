package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/civiledcode/tsv-merger/merger"
	"github.com/civiledcode/tsv-merger/tsv"
)

func main() {
	// The first argument should be the config file name.
	args := os.Args
	

	configContents, _ := ioutil.ReadFile(args[1])
	config := merger.TSVConfig{}
	err := json.Unmarshal(configContents, &config)
	if err != nil {
		panic(err)
	}
	
	// If the output seperator isn't defined, we default to the reading seperator.
	if config.OutputSeperator == "" {
		config.OutputSeperator = config.Seperator
	}

	columnList := make([]string, len(config.ColumnOperations))
	
	i := 0
	for column, _ := range config.ColumnOperations {
		columnList[i] = column
		i++
	}

	merg := merger.NewMerger(config)

	for _, file := range config.Files {
		fmt.Printf("\x1b[1;32mReading file %v...\x1b[0m\n", file)
		parser := tsv.NewParser(config.Seperator, file)

		for {
			key, values, err := parser.Row(config.KeyColumn, columnList)
			
			if err != nil && err != bufio.ErrAdvanceTooFar {
				fmt.Println(err)
				break
			} else if err != nil {
				break
			}

			for columnID, value := range values {
				merg.Put(key, columnID, value)
			}
		}
	}

	if len(args) < 3 {
		err = merg.Output(columnList, "output.tsv")
	} else {
		err = merg.Output(columnList, args[2])
	}

	if err != nil {
		panic(err)
	}

	fmt.Println("\x1b[1;36mOK\x1b[0m\n")
}

