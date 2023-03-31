package binance_connector

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type marketTestSuite struct {
	baseTestSuite
}

func TestMarket(t *testing.T) {
	suite.Run(t, new(marketTestSuite))
}

func (s *marketTestSuite) TestPing() {
	data := []byte(`{}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})

	err := s.client.NewPingService().Do(newContext())
	s.r().NoError(err)
}

func (s *marketTestSuite) TestAggTradesList() {
	data := []byte(`[
        {
            "a": 26129,
            "p": "0.01633102",
            "q": "4.70443515",
            "f": 27781,
            "l": 27781,
            "T": 1498793709153,
            "m": true,
            "M": true
        }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	fromID := int(1)
	startTime := uint64(1498793709153)
	endTime := uint64(1498793709156)
	limit := 1
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"symbol":    symbol,
			"fromId":    fromID,
			"startTime": startTime,
			"endTime":   endTime,
			"limit":     limit,
		})
		s.assertRequestEqual(e, r)
	})

	aggTrades, err := s.client.NewAggTradesListService().Symbol(symbol).
		FromId(fromID).StartTime(startTime).EndTime(endTime).Limit(limit).
		Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(aggTrades, 1)
	e := &AggTradesListResponse{
		AggTradeId:   26129,
		Price:        "0.01633102",
		Qty:          "4.70443515",
		FirstTradeId: 27781,
		LastTradeId:  27781,
		Time:         1498793709153,
		IsBuyer:      true,
		IsBest:       true,
	}
	s.assertAggTradeEqual(e, aggTrades[0])
}

func (s *marketTestSuite) assertAggTradeEqual(e, a *AggTradesListResponse) {
	r := s.r()
	r.Equal(e.AggTradeId, a.AggTradeId, "AggTradeID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Qty, a.Qty, "Quantity")
	r.Equal(e.FirstTradeId, a.FirstTradeId, "FirstTradeID")
	r.Equal(e.LastTradeId, a.LastTradeId, "LastTradeID")
	r.Equal(e.Time, a.Time, "Timestamp")
	r.Equal(e.IsBuyer, a.IsBuyer, "IsBuyerMaker")
	r.Equal(e.IsBest, a.IsBest, "IsBestPriceMatch")
}

func (s *marketTestSuite) TestServerTime() {
	data := []byte(`{
        "serverTime": 1499827319559
    }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})

	serverTime, err := s.client.NewServerTimeService().Do(newContext())

	e1 := &ServerTimeResponse{
		ServerTime: 1499827319559,
	}
	s.r().NoError(err)
	s.assertServerTimeEqual(e1, serverTime)
}

func (s *marketTestSuite) assertServerTimeEqual(e, a *ServerTimeResponse) {
	r := s.r()
	r.Equal(e.ServerTime, a.ServerTime, "ServerTime")
}

func (s *marketTestSuite) TestExchangeInfo() {

}

func (s *marketTestSuite) TestListBookTickers() {
	data := []byte(`[
        {
            "symbol": "LTCBTC",
            "bidPrice": "4.00000000",
            "bidQty": "431.00000000",
            "askPrice": "4.00000200",
            "askQty": "9.00000000"
        },
        {
            "symbol": "ETHBTC",
            "bidPrice": "0.07946700",
            "bidQty": "9.00000000",
            "askPrice": "100000.00000000",
            "askQty": "1000.00000000"
        }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})

	tickers, err := s.client.NewTickerBookTickerService().Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(tickers, 2)
	e1 := &TickerBookTickerResponse{
		Symbol:   "LTCBTC",
		BidPrice: "4.00000000",
		BidQty:   "431.00000000",
		AskPrice: "4.00000200",
		AskQty:   "9.00000000",
	}
	e2 := &TickerBookTickerResponse{
		Symbol:   "ETHBTC",
		BidPrice: "0.07946700",
		BidQty:   "9.00000000",
		AskPrice: "100000.00000000",
		AskQty:   "1000.00000000",
	}
	s.assertBookTickerEqual(e1, tickers[0])
	s.assertBookTickerEqual(e2, tickers[1])
}

func (s *marketTestSuite) assertBookTickerEqual(e, a *TickerBookTickerResponse) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.BidPrice, a.BidPrice, "BidPrice")
	r.Equal(e.BidQty, a.BidQty, "BidQuantity")
	r.Equal(e.AskPrice, a.AskPrice, "AskPrice")
	r.Equal(e.AskQty, a.AskQty, "AskQuantity")
}

func (s *marketTestSuite) TestRecentTrades() {
	data := []byte(`[
        {
            "id": 28457,
            "price": "4.00000100",
            "qty": "12.00000000",
            "quoteQty": "48.000012",
            "time": 1499865549590,
            "isBuyerMaker": true,
            "isBestMatch": true
        }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	limit := 3
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"symbol": symbol,
			"limit":  limit,
		})
		s.assertRequestEqual(e, r)
	})

	trades, err := s.client.NewRecentTradesListService().Symbol(symbol).Limit(limit).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(trades, 1)
	e := &RecentTradesListResponse{
		Id:           28457,
		Price:        "4.00000100",
		Qty:          "12.00000000",
		QuoteQty:     "48.000012",
		Time:         1499865549590,
		IsBuyerMaker: true,
		IsBest:       true,
	}
	s.assertTradeEqual(e, trades[0])
}

func (s *marketTestSuite) assertTradeEqual(e, a *RecentTradesListResponse) {
	r := s.r()
	r.Equal(e.Id, a.Id, "ID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Qty, a.Qty, "Qty")
	r.Equal(e.QuoteQty, a.QuoteQty, "QuoteQty")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.IsBuyerMaker, a.IsBuyerMaker, "IsBuyerMaker")
	r.Equal(e.IsBest, a.IsBest, "IsBest")
}

func (s *marketTestSuite) TestHistoricalTrades() {
	data := []byte(`[
        {
            "id": 28457,
            "price": "4.00000100",
            "qty": "12.00000000",
            "quoteQty": "48.000012",
            "time": 1499865549590,
            "isBuyerMaker": true,
            "isBestMatch": true
        }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	limit := uint(3)
	fromId := int64(1)
	s.assertReq(func(r *request) {
		e := newRequest().setParams(params{
			"symbol": symbol,
			"limit":  limit,
			"fromId": fromId,
		})
		s.assertRequestEqual(e, r)
	})

	trades, err := s.client.NewHistoricalTradeLookupService().Symbol(symbol).
		Limit(limit).FromId(fromId).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(trades, 1)
	e := &RecentTradesListResponse{
		Id:           28457,
		Price:        "4.00000100",
		Qty:          "12.00000000",
		QuoteQty:     "48.000012",
		Time:         1499865549590,
		IsBuyerMaker: true,
		IsBest:       true,
	}
	s.assertTradeEqual(e, trades[0])
}

func (s *marketTestSuite) TestAveragePrice() {
	data := []byte(`{
		"mins": 5,
		"price": "9.35751834"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "LTCBTC"
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewAvgPriceService().Symbol(symbol).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &AvgPriceResponse{
		Mins:  5,
		Price: "9.35751834",
	}
	s.assertAvgPrice(e, res)
}

func (s *marketTestSuite) assertAvgPrice(e, a *AvgPriceResponse) {
	s.r().Equal(e.Mins, a.Mins, "Mins")
	s.r().Equal(e.Price, a.Price, "Price")
}

func (s *marketTestSuite) Test24hrTicker() {
	data := []byte(`{
		"symbol": "BNBBTC",
        "priceChange": "-94.99999800",
        "priceChangePercent": "-95.960",
        "weightedAvgPrice": "0.29628482",
        "prevClosePrice": "0.10002000",
		"lastPrice": "4.00000200",
		"lastQty": "200.00000000",
        "bidPrice": "4.00000000",
        "askPrice": "4.00000200",
        "openPrice": "99.00000000",
        "highPrice": "100.00000000",
        "lowPrice": "0.10000000",
        "volume": "8913.30000000",
        "openTime": 1499783499040,
        "closeTime": 1499869899040,
        "firstId": 28385,
        "lastId": 28460,
        "count": 76
    }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BNBBTC"
	s.assertReq(func(r *request) {
		e := newRequest().setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})
	stats, err := s.client.NewTicker24hrService().Symbol(symbol).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &Ticker24hrResponse{
		Symbol:             "BNBBTC",
		PriceChange:        "-94.99999800",
		PriceChangePercent: "-95.960",
		WeightedAvgPrice:   "0.29628482",
		PrevClosePrice:     "0.10002000",
		LastPrice:          "4.00000200",
		LastQty:            "200.00000000",
		BidPrice:           "4.00000000",
		AskPrice:           "4.00000200",
		OpenPrice:          "99.00000000",
		HighPrice:          "100.00000000",
		LowPrice:           "0.10000000",
		Volume:             "8913.30000000",
		OpenTime:           1499783499040,
		CloseTime:          1499869899040,
		FirstId:            28385,
		LastId:             28460,
		Count:              76,
	}
	s.assertPriceChangeStatsEqual(e, stats)
}

func (s *marketTestSuite) assertPriceChangeStatsEqual(e, a *Ticker24hrResponse) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.PriceChange, a.PriceChange, "PriceChange")
	r.Equal(e.PriceChangePercent, a.PriceChangePercent, "PriceChangePercent")
	r.Equal(e.WeightedAvgPrice, a.WeightedAvgPrice, "WeightedAvgPrice")
	r.Equal(e.PrevClosePrice, a.PrevClosePrice, "PrevClosePrice")
	r.Equal(e.LastPrice, a.LastPrice, "LastPrice")
	r.Equal(e.LastQty, a.LastQty, "LastQty")
	r.Equal(e.BidPrice, a.BidPrice, "BidPrice")
	r.Equal(e.AskPrice, a.AskPrice, "AskPrice")
	r.Equal(e.OpenPrice, a.OpenPrice, "OpenPrice")
	r.Equal(e.HighPrice, a.HighPrice, "HighPrice")
	r.Equal(e.LowPrice, a.LowPrice, "LowPrice")
	r.Equal(e.Volume, a.Volume, "Volume")
	r.Equal(e.OpenTime, a.OpenTime, "OpenTime")
	r.Equal(e.CloseTime, a.CloseTime, "CloseTime")
	r.Equal(e.FirstId, a.FirstId, "FristID")
	r.Equal(e.LastId, a.LastId, "LastID")
	r.Equal(e.Count, a.Count, "Count")
}
