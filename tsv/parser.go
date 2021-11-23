package tsv

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Parser represents a text-seperated value (TSV) file parser.
type Parser struct {
	reader *bufio.Reader

	// Seperator is the byte that values are seperated with.
	Seperator byte

	// Columns maps column names to index.
	Columns map[string]int
}

// NewParser creates a new TSV parser.
func NewParser(seperator byte, file string) *Parser {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	
	buf := bufio.NewReader(f)
	
	colContent, _, err := buf.ReadLine()
	if err != nil {
		panic(err)
	}

	cols := strings.Split(string(colContent), string(seperator))
	
	parser := &Parser{Seperator: seperator, reader: buf, Columns: make(map[string]int)}

	for index, col := range cols {
		parser.Columns[col] = index + 1
	}

	return parser
}

// Row receives a key column and a list of columns to add within the row struct and retrieves those values for the next line within the file.
// If the column does not exist within this file, we return an error.
func (p *Parser) Row(key string, columns []string) (Row, error) {
	line, _, err := p.reader.ReadLine()
	if err != nil {
		return Row{}, err
	}

	l := strings.Split(string(line), string(p.Seperator))
	index := 0
	row := Row{Key: l[p.Columns[key] - 1], Value: make(map[string]string)}
	
	for _, column := range columns {
		index = p.Columns[column]
		if index == 0 {
			// Cast not found column error.
			return row, fmt.Errorf("column %v not found.")
		}

		row.Value[column] = l[index - 1]
	}

	return row, nil
}


