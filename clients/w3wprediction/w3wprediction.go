package w3wprediction

import (
	BinanceW3wPredictionRestApi "github.com/binance/binance-connector-go/clients/w3wprediction/src/restapi"
	BinanceW3wPredictionWebsocketStreams "github.com/binance/binance-connector-go/clients/w3wprediction/src/websocketstreams"
	"github.com/binance/binance-connector-go/common/v2/common"
)

type BinanceW3wPredictionClient struct {
	RestApi          *BinanceW3wPredictionRestApi.RestAPIClient
	WebsocketStreams *BinanceW3wPredictionWebsocketStreams.WebsocketStreamsClient
}

type Option func(*BinanceW3wPredictionClient)

func WithRestAPI(cfg *common.ConfigurationRestAPI) Option {
	return func(c *BinanceW3wPredictionClient) {
		c.RestApi = BinanceW3wPredictionRestApi.NewRestAPIClient(cfg)
	}
}

func WithWebsocketStreams(cfg *common.ConfigurationWebsocketStreams) Option {
	return func(c *BinanceW3wPredictionClient) {
		c.WebsocketStreams = BinanceW3wPredictionWebsocketStreams.NewWebsocketStreamsClient(cfg)
	}
}

func NewBinanceW3wPredictionClient(opts ...Option) *BinanceW3wPredictionClient {
	client := &BinanceW3wPredictionClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
