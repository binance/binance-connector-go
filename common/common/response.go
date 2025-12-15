package common

import (
	"net/http"
)

// RestApiResponse represents a generic REST API response.
//
// @param T The type of the data contained in the response.
// @field Data The data returned by the API.
// @field Status The HTTP status code of the response.
// @field Headers The HTTP headers of the response.
// @field RateLimits A list of rate limit information associated with the response.
type RestApiResponse[T any] struct {
	Data       T
	Status     int
	Headers    http.Header
	RateLimits []RateLimit
}

// ResponseOrRaw represents a response that can be either a typed structure or a raw map.
//
// @param T The type of the typed response.
// @field Typed A pointer to the typed response structure.
// @field Raw A map representing the raw response data.
type ResponseOrRaw[T any] struct {
	Typed *T
	Raw   map[string]interface{}
}
