/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package binancesimpleearnrestapi

import (
	"context"
	"net/url"
	"runtime"

	"github.com/binance/binance-connector-go/common/common"
)

// RestAPIClient manages communication with the Binance Simple Earn REST API API v1.2.0
type RestAPIClient struct {
	cfg *common.ConfigurationRestAPI

	// API Services
	BfusdAPI          *BfusdAPIService
	FlexibleLockedAPI *FlexibleLockedAPIService
	RwusdAPI          *RwusdAPIService
}

type Service struct {
	client *RestAPIClient
}

// NewRestAPIClient creates a new Binance Binance Simple Earn REST API REST API client
//
// @param cfg *common.ConfigurationRestAPI - The configuration for the REST API client
// @return *RestAPIClient - The newly created REST API client
func NewRestAPIClient(cfg *common.ConfigurationRestAPI) *RestAPIClient {
	customHeaders := common.GetCustomHeaders(cfg)

	if customHeaders == nil {
		customHeaders = make(map[string]string)
	}
	customHeaders["User-Agent"] = "binance-simpleearn/1.2.0 (Go/" + runtime.Version() + "; " + runtime.GOOS + "; " + runtime.GOARCH + ")"
	cfg.CustomHeaders = customHeaders
	c := &RestAPIClient{cfg: cfg}

	// API Services
	c.BfusdAPI = &BfusdAPIService{client: c}
	c.FlexibleLockedAPI = &FlexibleLockedAPIService{client: c}
	c.RwusdAPI = &RwusdAPIService{client: c}

	return c
}

// SendRequest sends an API request and returns the API response
//
// # T is a generic type parameter representing the expected response type
//
// @param ctx context.Context - The context for the request
// @param path string - The API endpoint path
// @param method string - The HTTP method (GET, POST, etc.)
// @param queryParams url.Values - The query parameters for the request
// @param bodyParams interface{} - The body parameters for the request
// @param config *common.ConfigurationRestAPI - The configuration for the REST API client
// @return *common.RestApiResponse[T] - The API response containing the typed data
// @return error - An error if the request fails
func SendRequest[T any](ctx context.Context, path string, method string, queryParams url.Values, bodyParams interface{}, config *common.ConfigurationRestAPI, signed bool) (*common.RestApiResponse[T], error) {
	resp, err := common.SendRequest[T](ctx, path, method, queryParams, bodyParams, config, signed)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// GetConfig returns the configuration of the REST API client
//
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
// @return *common.ConfigurationRestAPI - The configuration of the REST API client
func (c *RestAPIClient) GetConfig() *common.ConfigurationRestAPI {
	return c.cfg
}
