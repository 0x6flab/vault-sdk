package pkg

import "errors"

var (
	ErrMissingBindDB          = errors.New("missing binddb parameter")
	ErrMissingBindPass        = errors.New("missing bindpass parameters")
	ErrMissingServiceName     = errors.New("missing service name acccount")
	ErrMissingAccessKey       = errors.New("missing access key")
	ErrMissingSecretKey       = errors.New("missing secret key")
	ErrMissingName            = errors.New("missing name")
	ErrMissingLease           = errors.New("missing lease")
	ErrMissingLeaseMax        = errors.New("missing lease max")
	ErrMissingCredentialsType = errors.New("missing credentials type")
	ErrInvalidAWSCredType     = errors.New("invalid aws credentials type it is either sts or creds")
)
