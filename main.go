package main

import (
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
	config := merger.TSVConfig{Seperator: '	'}
	err := json.Unmarshal(configContents, &config)
	if err != nil {
		panic(err)
	}

	var m merger.MergerType

	if config.ArithmeticIncrementation {
		m = merger.IntegerMerger{Config: config, Content: make(map[string]int)}
	} else {
		m = merger.StringMerger{Config: config, Content: make(map[string]string)}
	}

	for _, file := range config.Files {
		parser := tsv.NewParser(config.Seperator, file)

		for {
			createdRow, err := parser.Row(config.KeyColumn, []string{config.IncrementalColumn})

			if err != nil {
				fmt.Println(err)
				break
			}

			m.Put(createdRow.Key, createdRow.Value[config.IncrementalColumn])
		}
	}

	err = m.Output("output.tsv")

	if err != nil {
		panic(err)
	}

	fmt.Println("\x1b[1;32mOK")
}

