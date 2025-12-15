package common

import (
	"crypto"
	"crypto/tls"
	"sync"
	"time"

	"golang.org/x/net/proxy"
)

type RateLimit struct {
	RateLimitType RateLimitType
	Interval      Interval
	IntervalNum   int32
	Count         int32
	RetryAfter    int32
}

type IntervalDetails struct {
	Interval    string
	IntervalNum int
}

type ProxyConfig struct {
	Host     string
	Port     int
	Protocol string
	Auth     struct {
		Username string
		Password string
	}
	TLSConfig *tls.Config
}

type HTTPSAgent interface{}

type Signer interface {
	Sign(data []byte) ([]byte, error)
}

type cryptoSigner struct {
	s crypto.Signer
}

type SendParams struct {
	Signed           bool
	WithAPIKey       bool
	WithSessionLogon bool
}

type SessionLogonRequest struct {
	Method string
	Params map[string]interface{}
	ID     string
	Priv   crypto.Signer
}

type ConfigurationWrapper struct {
	APIConfig     *ConfigurationWebsocketApi
	StreamsConfig *ConfigurationWebsocketStreams
}

type WebSocketConfig interface {
	GetAgent() HTTPSAgent
	GetBasePath() string
	GetReconnectDelay() time.Duration
	GetCompression() bool
	GetProxy() *ProxyConfig
	GetTLSConfig() *tls.Config
	GetTimeUnit() TimeUnit
}

type WebSocketConn interface {
	WriteMessage(messageType int, data []byte) error
	ReadMessage() (messageType int, p []byte, err error)
	Close() error
	SetReadLimit(limit int64)
}

type WebSocketConnection struct {
	Id                  string
	Connected           WebsocketStatus
	PendingMessages     sync.Map
	StreamCallbackMap   map[string][]func(map[string]interface{})
	Websocket           WebSocketConn
	SessionLogon        bool
	SessionLogonRequest *SessionLogonRequest
	StreamConnectionMap []string
	Done                chan struct{}
	ErrorChan           chan error
	mu                  sync.Mutex
}

type WebSocketCommon struct {
	Connections     []*WebSocketConnection
	ReconnectTasks  map[string]chan struct{}
	Mode            WebsocketMode
	PoolSize        int
	ProxyDialer     proxy.Dialer
	RoundRobinIndex int
}

type WebsocketAPI struct {
	Cfg                       *ConfigurationWebsocketApi
	GlobalStreamConnectionMap map[string][]*WebSocketConnection
	WsCommon                  *WebSocketCommon
}

type WebsocketStreams struct {
	Cfg                       *ConfigurationWebsocketStreams
	GlobalStreamConnectionMap map[string][]*WebSocketConnection
	WsCommon                  *WebSocketCommon
}

type StreamHandlerWrapper struct {
	WebsocketStreams *WebsocketStreams
	WebsocketAPI     *WebsocketAPI
}

type StreamHandler[T any] struct {
	Stream   string
	Wrapper  *StreamHandlerWrapper
	Callback func(T)
	ErrorCb  func(error)
}

type WebSocketError struct {
	Op      string
	ConnID  string
	Err     error
	Message string
	Code    int
}

type ResultWebsocket[T any] struct {
	Value *ResponseOrRaw[T]
	Err   error
}

type ResultUserData[T any, Ty any] struct {
	Value  *ResponseOrRaw[T]
	Stream *StreamHandler[Ty]
	Err    error
}

type httpProxyDialer struct {
	proxyAddr string
	auth      *proxy.Auth
}

type MappedNullable interface {
	ToMap() (map[string]interface{}, error)
}

type NullableTime struct {
	value *time.Time
	isSet bool
}
