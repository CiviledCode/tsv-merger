package tsv

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Parser struct {
	reader *bufio.Reader

	// Seperator is the byte that values are seperated with.
	Seperator byte

	// Columns maps column names to index.
	Columns map[string]int
}

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

func (p *Parser) Row(key string, columns []string) (Row, error) {
	line, _, err := p.reader.ReadLine()
	if err != nil {
		return Row{}, err
	}

	l := strings.Split(string(line), string(p.Seperator))
	index := 0
	row := Row{Key: l[p.Columns[key] - 1], Value: make(map[string]interface{})}
	
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


