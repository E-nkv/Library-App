package errs

import "fmt"

var (
	ErrNotFound = fmt.Errorf("entity not found")
)
