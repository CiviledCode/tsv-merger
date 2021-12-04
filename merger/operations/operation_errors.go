package operations

import "errors"

var ErrOperationNotFound = errors.New("operation not found.")
var ErrOperationOverwrite = errors.New("operation already exists.")
