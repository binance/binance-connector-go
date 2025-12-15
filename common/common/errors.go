package common

type BaseError struct {
	Name    string
	Message string
}

func (e *BaseError) Error() string {
	return e.Message
}

// ConnectorClientError represents a general error in the Connector client.
type ConnectorClientError struct {
	BaseError
}

func NewConnectorClientError(msg string) *ConnectorClientError {
	if msg == "" {
		msg = "An unexpected error occurred."
	}
	return &ConnectorClientError{
		BaseError: BaseError{
			Name:    "ConnectorClientError",
			Message: msg,
		},
	}
}

// RequiredError represents an error when a required parameter is missing or undefined.
type RequiredError struct {
	BaseError
	Field string
}

func NewRequiredError(field, msg string) *RequiredError {
	if msg == "" {
		msg = "Required parameter " + field + " was null or undefined."
	}
	return &RequiredError{
		BaseError: BaseError{
			Name:    "RequiredError",
			Message: msg,
		},
		Field: field,
	}
}

// UnauthorizedError represents an error when a client is unauthorized.
type UnauthorizedError struct {
	BaseError
}

func NewUnauthorizedError(msg string) *UnauthorizedError {
	if msg == "" {
		msg = "Unauthorized access. Authentication required."
	}
	return &UnauthorizedError{
		BaseError: BaseError{
			Name:    "UnauthorizedError",
			Message: msg,
		},
	}
}

// ForbiddenError represents an error when a resource is forbidden.
type ForbiddenError struct {
	BaseError
}

func NewForbiddenError(msg string) *ForbiddenError {
	if msg == "" {
		msg = "Access to the requested resource is forbidden."
	}
	return &ForbiddenError{
		BaseError: BaseError{
			Name:    "ForbiddenError",
			Message: msg,
		},
	}
}

// TooManyRequestsError represents an error when a client is rate-limited.
type TooManyRequestsError struct {
	BaseError
}

func NewTooManyRequestsError(msg string) *TooManyRequestsError {
	if msg == "" {
		msg = "Too many requests. You are being rate-limited."
	}
	return &TooManyRequestsError{
		BaseError: BaseError{
			Name:    "TooManyRequestsError",
			Message: msg,
		},
	}
}

// RateLimitBanError represents an error when a client IP has been banned.
type RateLimitBanError struct {
	BaseError
}

func NewRateLimitBanError(msg string) *RateLimitBanError {
	if msg == "" {
		msg = "The IP address has been banned for exceeding rate limits."
	}
	return &RateLimitBanError{
		BaseError: BaseError{
			Name:    "RateLimitBanError",
			Message: msg,
		},
	}
}

// ServerError represents an error that occurs when there is an internal server error.
type ServerError struct {
	BaseError
}

func NewServerError(msg string, statusCode int) *ServerError {
	if msg == "" {
		msg = "An internal server error occurred."
	}
	return &ServerError{
		BaseError: BaseError{
			Name:    "ServerError",
			Message: msg,
		},
	}
}

// NetworkError represents an error that occurs when a network error occurs.
type NetworkError struct {
	BaseError
}

func NewNetworkError(msg string) *NetworkError {
	if msg == "" {
		msg = "An internal server error occurred."
	}
	return &NetworkError{
		BaseError: BaseError{
			Name:    "NetworkError",
			Message: msg,
		},
	}
}

// NotFoundError represents an error that occurs when the requested resource was not found.
type NotFoundError struct {
	BaseError
}

func NewNotFoundError(msg string) *NotFoundError {
	if msg == "" {
		msg = "The requested resource was not found."
	}
	return &NotFoundError{
		BaseError: BaseError{
			Name:    "NotFoundError",
			Message: msg,
		},
	}
}

// BadRequestError represents an error that occurs when a request is invalid or cannot be otherwise served.
type BadRequestError struct {
	BaseError
}

func NewBadRequestError(msg string) *BadRequestError {
	if msg == "" {
		msg = "The request was invalid or cannot be otherwise served."
	}
	return &BadRequestError{
		BaseError: BaseError{
			Name:    "BadRequestError",
			Message: msg,
		},
	}
}
