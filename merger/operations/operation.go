package operations

import "fmt"

type OperationCallback func(oldValue, newAddition interface{}) interface{}

var registeredOperations map[string]OperationCallback

func init() {
  registeredOperations = make(map[string]OperationCallback)
    RegisterOperation("+", additionCallback)
    RegisterOperation("-", subtractionCallback)
    RegisterOperation("_", constantCallback)
}

// RegisterOperation maps a new operation callback to an operation byte.
// If there's already an operation registered under that byte, then an ErrOperationOverwrite error will be returned.
func RegisterOperation(operationExpressor string, op OperationCallback) error {
  if registeredOperations[operationExpressor] != nil {
    return fmt.Errorf("%v: %w", operationExpressor, ErrOperationOverwrite)
  }

  registeredOperations[operationExpressor] = op
  return nil
}

// GetOperation retrieves the respective opertor callback for the expression byte.
// If the expression byte has no callback registered, we return an ErrOperationNotFound error.
func GetOperation(operationExpressor string) (op OperationCallback, e error) {
  op = registeredOperations[operationExpressor]

  if op == nil {
    return nil, fmt.Errorf("could not retrieve %s: %w", operationExpressor, ErrOperationNotFound)
  }

  return op, nil
}
