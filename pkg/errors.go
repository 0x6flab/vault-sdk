package pkg

import "errors"

var (
	ErrMissingBindDB      = errors.New("missing binddb parameter")
	ErrMissingBindPass    = errors.New("missing bindpass parameters")
	ErrMissingServiceName = errors.New("missing service name acccount")
)
