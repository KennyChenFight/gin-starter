package business

const (
	// common errors
	Unknown            = 1000
	InvalidBodyParse   = 1001
	NotFound           = 1002
	Forbidden          = 1003
	ServiceUnavailable = 1004
	MethodNowAllowed   = 1005
	PathNotFound       = 1006

	// resource errors
	MemberNotFound = 1100

	// auth errors
	TokenRequired = 1200
	TokenInvalid  = 1201
	TokenExpired  = 1202
)
