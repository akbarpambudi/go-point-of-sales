package errors

func NewAuthorizationError(messageKey string, message string) error {
	return NewPOSError(ErrorTypeAuthorizationError, messageKey, message)
}

func NewAuthenticationError(messageKey string, message string) error {
	return NewPOSError(ErrorTypeAuthenticationError, messageKey, message)
}

func NewInternalError(messageKey string, message string) error {
	return NewPOSError(ErrorTypeInternalError, messageKey, message)
}

func NewIllegalInputError(messageKey string, message string) error {
	return NewPOSError(ErrorTypeIllegalInputError, messageKey, message)
}

func NewResourceNotFoundError(messageKey string, message string) error {
	return NewPOSError(ErrorTypeResourceNotFoundError, messageKey, message)
}
