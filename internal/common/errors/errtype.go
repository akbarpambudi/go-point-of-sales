package errors

const (
	ErrorTypeInternalError         ErrorType = "internal-error"
	ErrorTypeIllegalInputError     ErrorType = "illegal-input-error"
	ErrorTypeAuthenticationError   ErrorType = "authentication-error"
	ErrorTypeAuthorizationError    ErrorType = "authorization-error"
	ErrorTypeResourceNotFoundError ErrorType = "resource-not-found-error"
)
