package common

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ConfigurationRestAPI holds configuration settings for REST API interactions.
//
// @field apiKey The API key for authentication.
// @field apiSecret The API secret for authentication.
// @field BasePath The base URL for the API.
// @field Timeout The timeout duration for API requests.
// @field Proxy The proxy configuration.
// @field KeepAlive Indicates whether to keep connections alive.
// @field Compression Indicates whether to use compression.
// @field Retries The number of retry attempts for failed requests.
// @field Backoff The backoff duration between retries.
// @field TimeUnit The time unit used for rate limiting.
// @field privateKey The private key for signing requests.
// @field privateKeyPassphrase The passphrase for the private key.
// @field CustomHeaders A map of custom headers to include in requests.
// @field Signer The signer used for request signing.
// @field HTTPSAgent The HTTPS agent configuration.
type ConfigurationRestAPI struct {
	apiKey               string
	apiSecret            string
	BasePath             string
	Timeout              time.Duration
	Proxy                *ProxyConfig
	KeepAlive            bool
	Compression          bool
	Retries              int
	Backoff              int
	TimeUnit             TimeUnit
	privateKey           string
	privateKeyPassphrase string
	CustomHeaders        map[string]string
	Signer               Signer
	HTTPSAgent           HTTPSAgent
}

type ConfigurationRestAPIOption func(*ConfigurationRestAPI)

// ConfigurationWebsocketApi holds configuration settings for WebSocket API interactions.
//
// @field apiKey The API key for authentication.
// @field apiSecret The API secret for authentication.
// @field privateKey The private key for signing requests.
// @field privateKeyPassphrase The passphrase for the private key.
// @field BasePath The base URL for the WebSocket API.
// @field Timeout The timeout duration for WebSocket connections.
// @field ReconnectDelay The delay duration before attempting to reconnect.
// @field Compression Indicates whether to use compression.
// @field Proxy The proxy configuration.
// @field Mode The WebSocket connection mode.
// @field PoolSize The size of the connection pool.
// @field TimeUnit The time unit used for rate limiting.
// @field SessionReLogon Indicates whether to re-logon on session expiration.
// @field Signer The signer used for request signing.
// @field Agent The HTTPS agent configuration.
type ConfigurationWebsocketApi struct {
	apiKey               string
	apiSecret            string
	privateKey           string
	privateKeyPassphrase string
	BasePath             string
	Timeout              time.Duration
	ReconnectDelay       time.Duration
	Compression          bool
	Proxy                *ProxyConfig
	Mode                 WebsocketMode
	PoolSize             int
	TimeUnit             TimeUnit
	SessionReLogon       bool
	Signer               Signer
	Agent                HTTPSAgent
}

type ConfigurationWebsocketApiOption func(*ConfigurationWebsocketApi)

// ConfigurationWebsocketStreams holds configuration settings for WebSocket Streams interactions.
//
// @field BasePath The base URL for the WebSocket Streams.
// @field ReconnectDelay The delay duration before attempting to reconnect.
// @field Compression Indicates whether to use compression.
// @field Proxy The proxy configuration.
// @field Mode The WebSocket connection mode.
// @field PoolSize The size of the connection pool.
// @field TimeUnit The time unit used for rate limiting.
// @field Agent The HTTPS agent configuration.
type ConfigurationWebsocketStreams struct {
	BasePath       string
	ReconnectDelay time.Duration
	Compression    bool
	Proxy          *ProxyConfig
	Mode           WebsocketMode
	PoolSize       int
	TimeUnit       TimeUnit
	Agent          HTTPSAgent
}

type ConfigurationWebsocketStreamsOption func(*ConfigurationWebsocketStreams)

// NewConfigurationRestAPI creates a new ConfigurationRestAPI with the provided options.
//
// @param opts A variadic list of ConfigurationRestAPIOption functions to customize the configuration.
// @return A pointer to the newly created ConfigurationRestAPI.
func NewConfigurationRestAPI(opts ...ConfigurationRestAPIOption) *ConfigurationRestAPI {
	basePath := "https://api.binance.com"
	timeout := 5000 * time.Millisecond
	retries := 3
	backoff := 1000
	keepAlive := true
	compression := true

	cfg := &ConfigurationRestAPI{
		BasePath:    basePath,
		Timeout:     timeout,
		Retries:     retries,
		Backoff:     backoff,
		KeepAlive:   keepAlive,
		Compression: compression,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

func WithApiKey(v string) ConfigurationRestAPIOption {
	return func(c *ConfigurationRestAPI) { c.apiKey = v }
}

func (c *ConfigurationRestAPI) GetApiKey() string {
	return c.apiKey
}

func WithApiSecret(v string) ConfigurationRestAPIOption {
	return func(c *ConfigurationRestAPI) { c.apiSecret = v }
}

func (c *ConfigurationRestAPI) GetApiSecret() string {
	return c.apiSecret
}

func WithBasePath(v string) ConfigurationRestAPIOption {
	return func(c *ConfigurationRestAPI) { c.BasePath = v }
}

func WithTimeout(v time.Duration) ConfigurationRestAPIOption {
	return func(c *ConfigurationRestAPI) { c.Timeout = v }
}

func WithProxy(proxy ProxyConfig) ConfigurationRestAPIOption {
	return func(cfg *ConfigurationRestAPI) {
		cfg.Proxy = &proxy
	}
}

func WithKeepAlive(v bool) ConfigurationRestAPIOption {
	return func(c *ConfigurationRestAPI) { c.KeepAlive = v }
}

func WithCompression(v bool) ConfigurationRestAPIOption {
	return func(c *ConfigurationRestAPI) { c.Compression = v }
}

func WithRetries(v int) ConfigurationRestAPIOption {
	return func(c *ConfigurationRestAPI) { c.Retries = v }
}

func WithBackoff(v int) ConfigurationRestAPIOption {
	return func(c *ConfigurationRestAPI) { c.Backoff = v }
}

func WithTimeUnit(v TimeUnit) ConfigurationRestAPIOption {
	return func(c *ConfigurationRestAPI) { c.TimeUnit = v }
}

func WithPrivateKey(v string) ConfigurationRestAPIOption {
	return func(c *ConfigurationRestAPI) { c.privateKey = v }
}

func (c *ConfigurationRestAPI) GetPrivateKey() string {
	return c.privateKey
}

func WithPrivateKeyPassphrase(v string) ConfigurationRestAPIOption {
	return func(c *ConfigurationRestAPI) { c.privateKeyPassphrase = v }
}

func (c *ConfigurationRestAPI) GetPrivateKeyPassphrase() string {
	return c.privateKeyPassphrase
}

func WithCustomHeaders(v map[string]string) ConfigurationRestAPIOption {
	return func(c *ConfigurationRestAPI) { c.CustomHeaders = v }
}

func GetCustomHeaders(c *ConfigurationRestAPI) map[string]string {
	return c.CustomHeaders
}

func WithHTTPSAgent(v HTTPSAgent) ConfigurationRestAPIOption {
	return func(c *ConfigurationRestAPI) { c.HTTPSAgent = v }
}

// NewConfigurationWebsocketApi creates a new ConfigurationWebsocketApi with the provided options.
//
// @param opts A variadic list of ConfigurationWebsocketApiOption functions to customize the configuration.
// @return A pointer to the newly created ConfigurationWebsocketApi.
func NewConfigurationWebsocketApi(opts ...ConfigurationWebsocketApiOption) *ConfigurationWebsocketApi {
	stream_url := "wss://ws-api.binance.com:443/ws-api/v3"
	timeout := 5000 * time.Millisecond
	reconnect_delay := 5000 * time.Millisecond
	compression := true
	mode := SINGLE
	pool_size := 1
	sessionReLogon := true

	cfg := &ConfigurationWebsocketApi{
		BasePath:       stream_url,
		Timeout:        timeout,
		ReconnectDelay: reconnect_delay,
		Compression:    compression,
		Mode:           mode,
		PoolSize:       pool_size,
		SessionReLogon: sessionReLogon,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

func WithWsApiBasePath(v string) ConfigurationWebsocketApiOption {
	return func(c *ConfigurationWebsocketApi) { c.BasePath = v }
}

func (c *ConfigurationWebsocketApi) GetBasePath() string {
	return c.BasePath
}

func WithWsApiKey(v string) ConfigurationWebsocketApiOption {
	return func(c *ConfigurationWebsocketApi) { c.apiKey = v }
}

func (c *ConfigurationWebsocketApi) GetWsApiKey() string {
	return c.apiKey
}

func WithWsApiSecret(v string) ConfigurationWebsocketApiOption {
	return func(c *ConfigurationWebsocketApi) { c.apiSecret = v }
}

func (c *ConfigurationWebsocketApi) GetWsApiSecret() string {
	return c.apiSecret
}

func WithWsPrivateKey(v string) ConfigurationWebsocketApiOption {
	return func(c *ConfigurationWebsocketApi) { c.privateKey = v }
}

func (c *ConfigurationWebsocketApi) GetWsPrivateKey() string {
	return c.privateKey
}

func WithWsPrivateKeyPassphrase(v string) ConfigurationWebsocketApiOption {
	return func(c *ConfigurationWebsocketApi) { c.privateKeyPassphrase = v }
}

func (c *ConfigurationWebsocketApi) GetWsPrivateKeyPassphrase() string {
	return c.privateKeyPassphrase
}

func WithWsAPIBasePath(v string) ConfigurationWebsocketApiOption {
	return func(c *ConfigurationWebsocketApi) { c.BasePath = v }
}

func WithWsTimeout(v time.Duration) ConfigurationWebsocketApiOption {
	return func(c *ConfigurationWebsocketApi) { c.Timeout = v }
}

func WithWsReconnectDelay(v time.Duration) ConfigurationWebsocketApiOption {
	return func(c *ConfigurationWebsocketApi) { c.ReconnectDelay = v }
}

func (c *ConfigurationWebsocketApi) GetReconnectDelay() time.Duration {
	return c.ReconnectDelay
}

func WithWsCompression(v bool) ConfigurationWebsocketApiOption {
	return func(c *ConfigurationWebsocketApi) { c.Compression = v }
}

func (c *ConfigurationWebsocketApi) GetCompression() bool {
	return c.Compression
}

func WithWsProxy(proxy ProxyConfig) ConfigurationWebsocketApiOption {
	return func(cfg *ConfigurationWebsocketApi) {
		cfg.Proxy = &proxy
	}
}

func (c *ConfigurationWebsocketApi) GetProxy() *ProxyConfig {
	return c.Proxy
}

func WithWsMode(v WebsocketMode) ConfigurationWebsocketApiOption {
	return func(c *ConfigurationWebsocketApi) { c.Mode = v }
}

func WithWsPoolSize(v int) ConfigurationWebsocketApiOption {
	return func(c *ConfigurationWebsocketApi) { c.PoolSize = v }
}

func WithWsTimeUnit(v TimeUnit) ConfigurationWebsocketApiOption {
	return func(c *ConfigurationWebsocketApi) { c.TimeUnit = v }
}

func (c *ConfigurationWebsocketApi) GetTimeUnit() TimeUnit {
	return c.TimeUnit
}

func WithWsAgent(v HTTPSAgent) ConfigurationWebsocketApiOption {
	return func(c *ConfigurationWebsocketApi) { c.Agent = v }
}

func (c *ConfigurationWebsocketApi) GetAgent() HTTPSAgent {
	return c.Agent
}

func WithWsSessionReLogon(v bool) ConfigurationWebsocketApiOption {
	return func(c *ConfigurationWebsocketApi) { c.SessionReLogon = v }
}

// GetTLSConfig retrieves the TLS configuration from the HTTPS agent or proxy settings.
//
// @return A pointer to the tls.Config if available, otherwise nil.
func (c *ConfigurationWebsocketApi) GetTLSConfig() *tls.Config {
	if c.Agent != nil {
		if tlsConfig, ok := c.Agent.(*tls.Config); ok {
			return tlsConfig
		}
	}

	if c.Proxy != nil && c.Proxy.TLSConfig != nil {
		return c.Proxy.TLSConfig
	}
	return nil
}

// NewConfigurationWebsocketStreams creates a new ConfigurationWebsocketStreams with the provided options.
//
// @param opts A variadic list of ConfigurationWebsocketStreamsOption functions to customize the configuration.
// @return A pointer to the newly created ConfigurationWebsocketStreams.
func NewConfigurationWebsocketStreams(opts ...ConfigurationWebsocketStreamsOption) *ConfigurationWebsocketStreams {
	stream_url := "wss://stream.binance.com:9443/stream"
	reconnect_delay := 5000 * time.Millisecond
	compression := true
	mode := SINGLE
	pool_size := 1

	cfg := &ConfigurationWebsocketStreams{
		BasePath:       stream_url,
		ReconnectDelay: reconnect_delay,
		Compression:    compression,
		Mode:           mode,
		PoolSize:       pool_size,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

func WithWsStreamsBasePath(v string) ConfigurationWebsocketStreamsOption {
	return func(c *ConfigurationWebsocketStreams) { c.BasePath = v }
}

func (c *ConfigurationWebsocketStreams) GetBasePath() string {
	return c.BasePath
}

func WithWsStreamsReconnectDelay(v time.Duration) ConfigurationWebsocketStreamsOption {
	return func(c *ConfigurationWebsocketStreams) { c.ReconnectDelay = v }
}

func (c *ConfigurationWebsocketStreams) GetReconnectDelay() time.Duration {
	return c.ReconnectDelay
}

func WithWsStreamsCompression(v bool) ConfigurationWebsocketStreamsOption {
	return func(c *ConfigurationWebsocketStreams) { c.Compression = v }
}

func (c *ConfigurationWebsocketStreams) GetCompression() bool {
	return c.Compression
}

func WithWsStreamsProxy(proxy ProxyConfig) ConfigurationWebsocketStreamsOption {
	return func(c *ConfigurationWebsocketStreams) { c.Proxy = &proxy }
}

func (c *ConfigurationWebsocketStreams) GetProxy() *ProxyConfig {
	return c.Proxy
}

func WithWsStreamsMode(v WebsocketMode) ConfigurationWebsocketStreamsOption {
	return func(c *ConfigurationWebsocketStreams) { c.Mode = v }
}

func WithWsStreamsPoolSize(v int) ConfigurationWebsocketStreamsOption {
	return func(c *ConfigurationWebsocketStreams) { c.PoolSize = v }
}

func WithWsStreamsTimeUnit(v TimeUnit) ConfigurationWebsocketStreamsOption {
	return func(c *ConfigurationWebsocketStreams) { c.TimeUnit = v }
}

func (c *ConfigurationWebsocketStreams) GetTimeUnit() TimeUnit {
	return c.TimeUnit
}

func WithWsStreamsAgent(v HTTPSAgent) ConfigurationWebsocketStreamsOption {
	return func(c *ConfigurationWebsocketStreams) { c.Agent = v }
}

func (c *ConfigurationWebsocketStreams) GetAgent() HTTPSAgent {
	return c.Agent
}

// GetTLSConfig retrieves the TLS configuration from the HTTPS agent or proxy settings.
//
// @return A pointer to the tls.Config if available, otherwise nil.
func (c *ConfigurationWebsocketStreams) GetTLSConfig() *tls.Config {
	if c.Agent != nil {
		if tlsConfig, ok := c.Agent.(*tls.Config); ok {
			return tlsConfig
		}
	}
	return nil
}

// URL constructs the full URL for the proxy configuration.
//
// @return A pointer to the url.URL representing the proxy URL, or an error if the host is not set.
func (pc ProxyConfig) URL() (*url.URL, error) {
	if pc.Host == "" {
		return nil, fmt.Errorf("proxy host is required")
	}

	protocol := pc.Protocol
	if protocol == "" {
		protocol = "http" // default
	}

	u := &url.URL{
		Scheme: protocol,
		Host:   pc.Host,
	}

	if pc.Port > 0 {
		u.Host = fmt.Sprintf("%s:%d", pc.Host, pc.Port)
	}

	if pc.Auth.Username != "" {
		if pc.Auth.Password != "" {
			u.User = url.UserPassword(pc.Auth.Username, pc.Auth.Password)
		} else {
			u.User = url.User(pc.Auth.Username)
		}
	}

	return u, nil
}

// IsEmpty checks if the ProxyConfig is empty.
//
// @return True if all fields in the ProxyConfig are empty or zero, otherwise false.
func (pc ProxyConfig) IsEmpty() bool {
	return pc.Host == "" &&
		pc.Port == 0 &&
		pc.Protocol == "" &&
		pc.Auth.Username == "" &&
		pc.Auth.Password == ""
}

// BuildTransport constructs an http.RoundTripper based on the provided HTTPSAgent and ConfigurationRestAPI.
//
// @param agent The HTTPSAgent which can be nil, *tls.Config, or http.RoundTripper.
// @param cfg The ConfigurationRestAPI containing settings like Compression.
// @return An http.RoundTripper configured according to the provided agent and configuration.
func BuildTransport(agent HTTPSAgent, cfg *ConfigurationRestAPI) (http.RoundTripper, error) {
	switch a := agent.(type) {
	case nil:
		return &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: !cfg.Compression,
		}, nil
	case *tls.Config:
		return &http.Transport{
			TLSClientConfig:    a,
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: !cfg.Compression,
		}, nil
	case http.RoundTripper:
		return a, nil
	default:
		return nil, fmt.Errorf("unsupported HTTPSAgent type")
	}
}

// WsBuildTransport constructs an http.RoundTripper for WebSocket connections based on the provided agent and TLS configuration.
//
// @param agent The HTTPSAgent which can be a string (proxy URL) or other types.
// @param tlsConfig The TLS configuration to be used if applicable.
// @return An http.RoundTripper configured according to the provided agent and TLS configuration.
func WsBuildTransport(agent HTTPSAgent, tlsConfig *tls.Config) http.RoundTripper {
	if proxyURL, ok := agent.(string); ok {
		proxyURLParsed, err := url.Parse(proxyURL)
		if err != nil {
			return nil
		}

		return &http.Transport{
			Proxy:           http.ProxyURL(proxyURLParsed),
			TLSClientConfig: tlsConfig,
		}
	}

	return nil
}
