package tsv

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ErrColumnNotFound error = errors.New("column not found")

// Parser represents a text-seperated value (TSV) file parser.
type Parser struct {
	reader *bufio.Reader

	// Seperator is the byte that values are seperated with.
	Seperator string

	// Columns maps column names to index.
	Columns map[string]int
}

// NewParser creates a new TSV parser.
func NewParser(seperator string, file string) *Parser {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	
	buf := bufio.NewReader(f)


	colContent, _, err := buf.ReadLine()
	if err != nil {
		panic(err)
	}

	cols := strings.Split(string(colContent), seperator)
	
	parser := &Parser{Seperator: seperator, reader: buf, Columns: make(map[string]int)}

	for index, col := range cols {
		parser.Columns[col] = index + 1
		fmt.Printf("found column '%v'\n", col)
	}

	return parser
}

// Row receives the column that will be used as a string along with a list of columns we should be grabbing content from. 
// If a column received doesn't exist within the parser, this function returns ErrColumnNotFound.
// The values stored within values are properly type casted using interpret().
// Values are seperated using the seperator configuration value.
func (p *Parser) Row(k string, columns []string) (key string, values map[string]interface{}, err error) {
	line, _, err := p.reader.ReadLine()
	if err != nil {
		return key, values, err
	}

	l := strings.Split(string(line), string(p.Seperator))
	index := 0
	key = l[p.Columns[k] - 1]
	values = make(map[string]interface{})
	
	for _, column := range columns {
		index = p.Columns[column]
		if index == 0 {
			// Cast not found column error.
			return key, nil, fmt.Errorf("'%v': %w", column, ErrColumnNotFound)
		}
		
		values[column] = interpret(l[index - 1])
	}

	return key, values, nil
}

// interpret receives a value as a string and attempts to parse it into
// integers, floats, and booleans. If the value pareses with no error, then that type is selected.
// The order for parsing is int, float, then bool.
func interpret(value string) interface{} {
	vi, err := strconv.Atoi(value)
	if err == nil {
		return vi
	}

	vf, err := strconv.ParseFloat(value, 64)
	if err == nil {
		return vf
	}

	vb, err := strconv.ParseBool(value)
	if err == nil {
		return vb
	}

	return value
}

