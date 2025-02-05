package binance_connector

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type copyTradingTestSuite struct {
	baseTestSuite
}

func TestCopyTrading(t *testing.T) {
	suite.Run(t, new(copyTradingTestSuite))
}

func (s *copyTradingTestSuite) TestGetFuturesLeadTraderStatus() {
	data := []byte(`
	{
		"isLeadTrader": false,
		"time": 1499827319559
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	res, err := s.client.NewGetFuturesLeadTraderStatusService().
		Do(context.Background())

	s.r().NoError(err)

	s.assetGetFuturesLeadTraderStatusEqual(&GetFuturesLeadTraderStatusResponse{
		IsLeadTrader: false,
		Time:         1499827319559,
	}, res)
}

func (s *copyTradingTestSuite) assetGetFuturesLeadTraderStatusEqual(expected, other *GetFuturesLeadTraderStatusResponse) {
	r := s.r()

	r.EqualValues(expected, other)

}

func (s *copyTradingTestSuite) TestGetFuturesLeadTradingSymbolWhitelist() {
	data := []byte(`
	{
		"code": "000000",
		"message": "success",
		"data": [
			{
				"symbol": "BTCUSDT",
				"baseAsset": "BTC",
				"quoteAsset": "USDT"
			},
			{
				"symbol": "ETHUSDT",
				"baseAsset": "ETH",
				"quoteAsset": "USDT"
			}
		],
		"success": true
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	res, err := s.client.NewGetFuturesLeadTradingSymbolWhitelistService().
		Do(context.Background())

	s.r().NoError(err)

	s.assetGetFuturesLeadTradingSymbolWhitelistEqual(&GetFuturesLeadTradingSymbolWhitelistResponse{
		Code:    "000000",
		Message: "success",
		Data: []struct {
			Symbol     string `json:"symbol"`
			BaseAsset  string `json:"baseAsset"`
			QuoteAsset string `json:"quoteAsset"`
		}{
			{
				Symbol:     "BTCUSDT",
				BaseAsset:  "BTC",
				QuoteAsset: "USDT",
			},
			{
				Symbol:     "ETHUSDT",
				BaseAsset:  "ETH",
				QuoteAsset: "USDT",
			},
		},
	}, res)
}

func (s *copyTradingTestSuite) assetGetFuturesLeadTradingSymbolWhitelistEqual(expected, other *GetFuturesLeadTradingSymbolWhitelistResponse) {
	r := s.r()

	r.EqualValues(expected, other)

}
