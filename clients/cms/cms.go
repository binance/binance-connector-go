package cms

import (
	BinanceCmsWebsocketStreams "github.com/binance/binance-connector-go/clients/cms/src/websocketstreams"
	"github.com/binance/binance-connector-go/common/v2/common"
)

type BinanceCmsClient struct {
	WebsocketStreams *BinanceCmsWebsocketStreams.WebsocketStreamsClient
}

type Option func(*BinanceCmsClient)

func WithWebsocketStreams(cfg *common.ConfigurationWebsocketStreams) Option {
	return func(c *BinanceCmsClient) {
		c.WebsocketStreams = BinanceCmsWebsocketStreams.NewWebsocketStreamsClient(cfg)
	}
}

func NewBinanceCmsClient(opts ...Option) *BinanceCmsClient {
	client := &BinanceCmsClient{}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
