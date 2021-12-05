package merger

import (
	"fmt"
	"os"
	"strings"

	"github.com/civiledcode/tsv-merger/merger/operations"
)

type Merger struct {
	// Content is the final column values.
	Content map[string]map[string]interface{}

	// Config depicts the way we merge columns by loading our configuration file.
	Config TSVConfig
}

func NewMerger(config TSVConfig) *Merger {
	return &Merger{Config: config, Content: make(map[string]map[string]interface{})}
}

// Put hashes the rowID and columnID and stores the value within a content map using this hash as a key.
func (m *Merger) Put(rowID, columnID string, value interface{}) {
	operation, err := operations.GetOperation(m.Config.ColumnOperations[columnID])	
	if err != nil {
		panic(err)
	}

	if m.Content[rowID] == nil {
		m.Content[rowID] = make(map[string]interface{})
	}

	oldValue := m.Content[rowID][columnID]
	if oldValue == nil {
		m.Content[rowID][columnID] = value
		return
	}

	merged := operation(oldValue, value)
	
	m.Content[rowID][columnID] = merged
}

// Output receives the columns we are outputting to the file and the file name and writes the contents of the merger in TSV format.
// The order of outputColumns dictates the column order in the output file.
func (m *Merger) Output(outputColumns []string, file string) error {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	
	var sb strings.Builder
	sb.WriteString(outputColumns[0])

	for i := 1; i < len(outputColumns); i++ {
		sb.WriteString(m.Config.OutputSeperator + outputColumns[i])
	}
	sb.WriteString("\n")
	_, err = f.WriteString(sb.String())
	if err != nil {
		return err
	}

	sb.Reset()

	for row, columns := range m.Content {
		sb.WriteString(row)
		for _, column := range outputColumns {
			sb.WriteString(fmt.Sprintf("%v%v", m.Config.OutputSeperator, columns[column]))			
		}
		sb.WriteString("\n")

		f.WriteString(sb.String())
		sb.Reset()
	}

	return f.Close()
}

func unhashPosition(hash string) (string, string) {
	s := strings.Split(hash, ":")

	return s[0], s[1] 
}

// hashPosition completes the simplest possible hash on the row and column IDs.
func hashPosition(rowID, columnID string) string {
	return rowID + ":" + columnID
}
