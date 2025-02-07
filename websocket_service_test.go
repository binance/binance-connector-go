package binance_connector

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type websocketTestSuite struct {
	baseTestSuite
	origWsServe func(*WsConfig, WsHandler, ErrHandler) (chan struct{}, chan struct{}, error)
	serveCount  int
}

func TestWebsocketService(t *testing.T) {
	suite.Run(t, new(websocketTestSuite))
}

func (s *websocketTestSuite) SetupTest() {
	s.origWsServe = wsServe
}

func (s *websocketTestSuite) TearDownTest() {
	wsServe = s.origWsServe
	s.serveCount = 0
}

func (s *websocketTestSuite) mockWsServe(data []byte, err error) {
	wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, innerErr error) {
		s.serveCount++
		doneCh = make(chan struct{})
		stopCh = make(chan struct{})
		go func() {
			<-stopCh
			close(doneCh)
		}()
		handler(data)
		if err != nil {
			errHandler(err)
		}
		return doneCh, stopCh, nil
	}
}

func (s *websocketTestSuite) assertWsServe(count ...int) {
	e := 1
	if len(count) > 0 {
		e = count[0]
	}
	s.r().Equal(e, s.serveCount)
}

func (s *websocketTestSuite) assertUserDataEvent(e, a *WsUserDataEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.AccountUpdateTime, a.AccountUpdateTime, "AccountUpdateTime")
	for i, e := range e.AccountUpdate.WsAccountUpdates {
		a := a.AccountUpdate.WsAccountUpdates[i]
		s.assertAccountUpdate(&e, &a)
	}
	s.assertOrderUpdate(&e.OrderUpdate, &a.OrderUpdate)
	s.assertBalanceUpdate(&e.BalanceUpdate, &a.BalanceUpdate)
}

func (s *websocketTestSuite) testWsUserDataServe(data []byte, expectedEvent *WsUserDataEvent) {
	websocketStreamClient := NewWebsocketStreamClient(false, "wss://stream.testnet.binance.vision")

	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := websocketStreamClient.WsUserDataServe("listenKey", func(event *WsUserDataEvent) {
		s.assertUserDataEvent(expectedEvent, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketTestSuite) assertAccountUpdate(e, a *WsAccountUpdate) {
	r := s.r()
	r.Equal(e.Asset, a.Asset)
	r.Equal(e.Free, a.Free)
	r.Equal(e.Locked, a.Locked)
}

func (s *websocketTestSuite) assertOrderUpdate(e, a *WsOrderUpdate) {
	r := s.r()
	r.Equal(e.TransactionTime, a.TransactionTime, "TransactionTime")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Volume, a.Volume, "Volume")
	r.Equal(e.QuoteVolume, a.QuoteVolume, "QuoteVolume")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.IsMaker, a.IsMaker, "IsMaker")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.CreateTime, a.CreateTime, "CreateTime")
	r.Equal(e.Id, a.Id, "Id")
	r.Equal(e.StopPrice, a.StopPrice, "StopPrice")
	r.Equal(e.TradeId, a.TradeId, "TradeId")
	r.Equal(e.ExecutionType, a.ExecutionType, "ExecutionType")
	r.Equal(e.FeeAsset, a.FeeAsset, "FeeAsset")
	r.Equal(e.FeeCost, a.FeeCost, "FeeCost")
	r.Equal(e.FilledQuoteVolume, a.FilledQuoteVolume, "FilledQuoteVolume")
	r.Equal(e.FilledVolume, a.FilledVolume, "FilledVolume")
	r.Equal(e.IceBergVolume, a.IceBergVolume, "IceBergVolume")
	r.Equal(e.IsInOrderBook, a.IsInOrderBook, "IsInOrderBook")
	r.Equal(e.LatestPrice, a.LatestPrice, "LatestPrice")
	r.Equal(e.OrderListId, a.OrderListId, "LatestQuoteVolume")
	r.Equal(e.LatestQuoteVolume, a.LatestQuoteVolume, "LatestQuoteVolume")
	r.Equal(e.LatestVolume, a.LatestVolume, "OrigCustomOrderId")
	r.Equal(e.OrigCustomOrderId, a.OrigCustomOrderId, "OrigCustomOrderId")
	r.Equal(e.RejectReason, a.RejectReason, "RejectReason")
}

func (s *websocketTestSuite) assertBalanceUpdate(e, a *WsBalanceUpdate) {
	r := s.r()
	r.Equal(e.Asset, a.Asset)
	r.Equal(e.Change, a.Change)
}

func (s *websocketTestSuite) TestWsUserDataServeAccountUpdate() {
	data := []byte(`{
	   "e":"outboundAccountPosition",
	   "E":1629771130464,
	   "u":1629771130463,
	   "B":[
	      {
	         "a":"LTC",
	         "f":"503.70000000",
	         "l":"0.00000000"
	      }
	   ]
	}`)
	expectedEvent := &WsUserDataEvent{
		Event:             "outboundAccountPosition",
		Time:              1629771130464,
		AccountUpdateTime: 1629771130463,
		AccountUpdate: WsAccountUpdateList{
			[]WsAccountUpdate{
				{
					"LTC",
					"503.70000000",
					"0.00000000",
				},
			},
		},
	}
	s.testWsUserDataServe(data, expectedEvent)
}

func (s *websocketTestSuite) TestWsTradeServe() {
	websocketStreamClient := NewWebsocketStreamClient(false, "wss://stream.testnet.binance.vision")

	data := []byte(`{
        "e": "trade",
        "E": 123456789,
        "s": "BNBBTC",
        "t": 12345,
        "p": "0.001",
        "q": "100",
        "b": 88,
        "a": 50,
        "T": 123456785,
        "m": true,
        "M": true
    }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := websocketStreamClient.WsTradeServe("BNBBTC", func(event *WsTradeEvent) {
		e := &WsTradeEvent{
			Event:         "trade",
			Time:          123456789,
			Symbol:        "BNBBTC",
			TradeID:       12345,
			Price:         "0.001",
			Quantity:      "100",
			BuyerOrderId:  88,
			SellerOrderId: 50,
			TradeTime:     123456785,
			IsBuyerMaker:  true,
		}
		s.assertWsTradeEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketTestSuite) assertWsTradeEventEqual(e, a *WsTradeEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.TradeID, a.TradeID, "TradeID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.BuyerOrderId, a.BuyerOrderId, "BuyerOrderID")
	r.Equal(e.SellerOrderId, a.SellerOrderId, "SellerOrderID")
	r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
	r.Equal(e.IsBuyerMaker, a.IsBuyerMaker, "IsBuyerMaker")
}

func (s *websocketTestSuite) TestWsAllMarketMiniTickersStatServe() {
	websocketStreamClient := NewWebsocketStreamClient(false, "wss://stream.testnet.binance.vision")
	data := []byte(`[{
  		"e": "24hrMiniTicker",
    	"E": 1523658017154,
    	"s": "BNBBTC",
   	 	"c": "0.00175640",
    	"o": "0.00161200",
    	"h": "0.00176000",
    	"l": "0.00159370",
    	"v": "3479863.89000000",
    	"q": "5725.90587704"
	},{
  		"e": "24hrMiniTicker",
    	"E": 1523658017133,
    	"s": "BNBETH",
    	"c": "0.02827000",
    	"o": "0.02628100",
    	"h": "0.02830300",
    	"l": "0.02469400",
    	"v": "456266.78000000",
    	"q": "11873.11095682"
	}]`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := websocketStreamClient.WsAllMarketMiniTickersStatServe(func(event WsAllMarketMiniTickersStatEvent) {
		e := WsAllMarketMiniTickersStatEvent{
			&WsMarketMiniStatEvent{
				Event:       "24hrMiniTicker",
				Time:        1523658017154,
				Symbol:      "BNBBTC",
				LastPrice:   "0.00175640",
				OpenPrice:   "0.00161200",
				HighPrice:   "0.00176000",
				LowPrice:    "0.00159370",
				BaseVolume:  "3479863.89000000",
				QuoteVolume: "5725.90587704",
			},
			&WsMarketMiniStatEvent{
				Event:       "24hrMiniTicker",
				Time:        1523658017133,
				Symbol:      "BNBETH",
				LastPrice:   "0.02827000",
				OpenPrice:   "0.02628100",
				HighPrice:   "0.02830300",
				LowPrice:    "0.02469400",
				BaseVolume:  "456266.78000000",
				QuoteVolume: "11873.11095682",
			},
		}
		s.assertWsAllMarketMiniTickersStatEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketTestSuite) assertWsAllMarketMiniTickersStatEventEqual(e, a WsAllMarketMiniTickersStatEvent) {
	for i := range e {
		s.assertWsMarketMiniTickersStatEventEqual(e[i], a[i])
	}
}

func (s *websocketTestSuite) assertWsMarketMiniTickersStatEventEqual(e, a *WsMarketMiniStatEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.LastPrice, a.LastPrice, "LastPrice")
	r.Equal(e.OpenPrice, a.OpenPrice, "OpenPrice")
	r.Equal(e.HighPrice, a.HighPrice, "HighPrice")
	r.Equal(e.LowPrice, a.LowPrice, "LowPrice")
	r.Equal(e.BaseVolume, a.BaseVolume, "BaseVolume")
	r.Equal(e.QuoteVolume, a.QuoteVolume, "QuoteVolume")
}

func (s *websocketTestSuite) TestWsMarketMiniTickersStatServe() {
	websocketStreamClient := NewWebsocketStreamClient(false, "wss://stream.testnet.binance.vision")
	data := []byte(`{
  		"e": "24hrMiniTicker",
    	"E": 1523658017154,
    	"s": "BNBBTC",
   	 	"c": "0.00175640",
    	"o": "0.00161200",
    	"h": "0.00176000",
    	"l": "0.00159370",
    	"v": "3479863.89000000",
    	"q": "5725.90587704"
	}`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := websocketStreamClient.WsMarketMiniTickersStatServe("BNBBTC", func(event WsMarketMiniTickerStatEvent) {
		e := &WsMarketMiniStatEvent{
			Event:       "24hrMiniTicker",
			Time:        1523658017154,
			Symbol:      "BNBBTC",
			LastPrice:   "0.00175640",
			OpenPrice:   "0.00161200",
			HighPrice:   "0.00176000",
			LowPrice:    "0.00159370",
			BaseVolume:  "3479863.89000000",
			QuoteVolume: "5725.90587704",
		}
		s.assertWsMarketMiniTickersStatEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketTestSuite) TestBookTickerServe() {
	websocketStreamClient := NewWebsocketStreamClient(false, "wss://stream.testnet.binance.vision")
	data := []byte(`{
  		"u":17242169,
  		"s":"BNBUSDT",
  		"b":"9548.1",
  		"B":"52",
  		"a":"9548.5",
  		"A":"11"
	  }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := websocketStreamClient.WsBookTickerServe("BNBUSDT", func(event *WsBookTickerEvent) {
		e := &WsBookTickerEvent{
			UpdateID:     17242169,
			Symbol:       "BNBUSDT",
			BestBidPrice: "9548.1",
			BestBidQty:   "52",
			BestAskPrice: "9548.5",
			BestAskQty:   "11",
		}
		s.assertWsBookTickerEvent(e, event)
	},
		func(err error) {
			s.r().EqualError(err, fakeErrMsg)
		})

	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketTestSuite) assertWsBookTickerEvent(e, a *WsBookTickerEvent) {
	r := s.r()
	r.Equal(e.UpdateID, a.UpdateID, "UpdateID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.BestBidPrice, a.BestBidPrice, "BestBidPrice")
	r.Equal(e.BestBidQty, a.BestBidQty, "BestBidQty")
	r.Equal(e.BestAskPrice, a.BestAskPrice, "BestAskPrice")
	r.Equal(e.BestAskQty, a.BestAskQty, "BestAskQty")
}

func (s *websocketTestSuite) TestDepthServe() {
	websocketStreamClient := NewWebsocketStreamClient(false, "wss://stream.testnet.binance.vision")
	data := []byte(`{
        "e": "depthUpdate",
        "E": 1499404630606,
        "s": "ETHBTC",
        "u": 7913455,
        "U": 7913452,
        "b": [
            [
                "0.10376590",
                "59.15767010",
                []
            ]
        ],
        "a": [
            [
                "0.10376586",
                "159.15767010",
                []
            ],
            [
                "0.10383109",
                "345.86845230",
                []
            ],
            [
                "0.10490700",
                "0.00000000",
                []
            ]
        ]
    }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := websocketStreamClient.WsDepthServe("ETHBTC", func(event *WsDepthEvent) {
		e := &WsDepthEvent{
			Event:         "depthUpdate",
			Time:          1499404630606,
			Symbol:        "ETHBTC",
			LastUpdateID:  7913455,
			FirstUpdateID: 7913452,
			Bids: []Bid{
				{
					Price:    "0.10376590",
					Quantity: "59.15767010",
				},
			},
			Asks: []Ask{
				{
					Price:    "0.10376586",
					Quantity: "159.15767010",
				},
				{
					Price:    "0.10383109",
					Quantity: "345.86845230",
				},
				{
					Price:    "0.10490700",
					Quantity: "0.00000000",
				},
			},
		}
		s.assertWsDepthEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketTestSuite) assertWsDepthEventEqual(e, a *WsDepthEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.LastUpdateID, a.LastUpdateID, "UpdateID")
	r.Equal(e.FirstUpdateID, a.FirstUpdateID, "FirstUpdateID")
	for i := 0; i < len(e.Bids); i++ {
		r.Equal(e.Bids[i].Price, a.Bids[i].Price, "Price")
		r.Equal(e.Bids[i].Quantity, a.Bids[i].Quantity, "Quantity")
	}
	for i := 0; i < len(e.Asks); i++ {
		r.Equal(e.Asks[i].Price, a.Asks[i].Price, "Price")
		r.Equal(e.Asks[i].Quantity, a.Asks[i].Quantity, "Quantity")
	}
}

func (s *websocketTestSuite) TestKlineServe() {
	websocketStreamClient := NewWebsocketStreamClient(false, "wss://stream.testnet.binance.vision")
	data := []byte(`{
        "e": "kline",
        "E": 1499404907056,
        "s": "ETHBTC",
        "k": {
            "t": 1499404860000,
            "T": 1499404919999,
            "s": "ETHBTC",
            "i": "1m",
            "f": 77462,
            "L": 77465,
            "o": "0.10278577",
            "c": "0.10278645",
            "h": "0.10278712",
            "l": "0.10278518",
            "v": "17.47929838",
            "n": 4,
            "x": false,
            "q": "1.79662878",
            "V": "2.34879839",
            "Q": "0.24142166",
            "B": "13279784.01349473"
        }
    }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := websocketStreamClient.WsKlineServe("ETHBTC", "1m", func(event *WsKlineEvent) {
		e := &WsKlineEvent{
			Event:  "kline",
			Time:   1499404907056,
			Symbol: "ETHBTC",
			Kline: WsKline{
				StartTime:            1499404860000,
				EndTime:              1499404919999,
				Symbol:               "ETHBTC",
				Interval:             "1m",
				FirstTradeID:         77462,
				LastTradeID:          77465,
				Open:                 "0.10278577",
				Close:                "0.10278645",
				High:                 "0.10278712",
				Low:                  "0.10278518",
				Volume:               "17.47929838",
				TradeNum:             4,
				IsFinal:              false,
				QuoteVolume:          "1.79662878",
				ActiveBuyVolume:      "2.34879839",
				ActiveBuyQuoteVolume: "0.24142166",
			},
		}
		s.assertWsKlineEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketTestSuite) TestCombinedKlineServe() {
	websocketStreamClient := NewWebsocketStreamClient(false, "wss://stream.testnet.binance.vision")
	data := []byte(`{
        "e": "kline",
        "E": 1499404907056,
        "s": "ETHBTC",
        "k": {
            "t": 1499404860000,
            "T": 1499404919999,
            "s": "ETHBTC",
            "i": "1m",
            "f": 77462,
            "L": 77465,
            "o": "0.10278577",
            "c": "0.10278645",
            "h": "0.10278712",
            "l": "0.10278518",
            "v": "17.47929838",
            "n": 4,
            "x": false,
            "q": "1.79662878",
            "V": "2.34879839",
            "Q": "0.24142166",
            "B": "13279784.01349473"
        }
    }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()
	symbolInterval := []CombinedKlineSymbolInterval{
		{
			Symbol:   "ETHBTC",
			Interval: "1m",
		},
		{
			Symbol:   "BTCUSDT",
			Interval: "1m",
		},
	}

	doneC, stopC, err := websocketStreamClient.WsCombinedKlineServe(symbolInterval, func(event *WsKlineEvent) {
		e := &WsKlineEvent{
			Event:  "kline",
			Time:   1499404907056,
			Symbol: "ETHBTC",
			Kline: WsKline{
				StartTime:            1499404860000,
				EndTime:              1499404919999,
				Symbol:               "ETHBTC",
				Interval:             "1m",
				FirstTradeID:         77462,
				LastTradeID:          77465,
				Open:                 "0.10278577",
				Close:                "0.10278645",
				High:                 "0.10278712",
				Low:                  "0.10278518",
				Volume:               "17.47929838",
				TradeNum:             4,
				IsFinal:              false,
				QuoteVolume:          "1.79662878",
				ActiveBuyVolume:      "2.34879839",
				ActiveBuyQuoteVolume: "0.24142166",
			},
		}
		s.assertWsKlineEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketTestSuite) assertWsKlineEventEqual(e, a *WsKlineEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	ek, ak := e.Kline, a.Kline
	r.Equal(ek.StartTime, ak.StartTime, "StartTime")
	r.Equal(ek.EndTime, ak.EndTime, "EndTime")
	r.Equal(ek.Symbol, ak.Symbol, "Symbol")
	r.Equal(ek.Interval, ak.Interval, "Interval")
	r.Equal(ek.FirstTradeID, ak.FirstTradeID, "FirstTradeID")
	r.Equal(ek.LastTradeID, ak.LastTradeID, "LastTradeID")
	r.Equal(ek.Open, ak.Open, "Open")
	r.Equal(ek.Close, ak.Close, "Close")
	r.Equal(ek.High, ak.High, "High")
	r.Equal(ek.Low, ak.Low, "Low")
	r.Equal(ek.Volume, ak.Volume, "Volume")
	r.Equal(ek.TradeNum, ak.TradeNum, "TradeNum")
	r.Equal(ek.IsFinal, ak.IsFinal, "IsFinal")
	r.Equal(ek.QuoteVolume, ak.QuoteVolume, "QuoteVolume")
	r.Equal(ek.ActiveBuyVolume, ak.ActiveBuyVolume, "ActiveBuyVolume")
	r.Equal(ek.ActiveBuyQuoteVolume, ak.ActiveBuyQuoteVolume, "ActiveBuyQuoteVolume")
}

func (s *websocketTestSuite) TestWsAggTradeServe() {
	websocketStreamClient := NewWebsocketStreamClient(false, "wss://stream.testnet.binance.vision")
	data := []byte(`{
        "e": "aggTrade",
        "E": 1499405254326,
        "s": "ETHBTC",
        "a": 70232,
        "p": "0.10281118",
        "q": "8.15632997",
        "f": 77489,
        "l": 77489,
        "T": 1499405254324,
        "m": false,
        "M": true
    }`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := websocketStreamClient.WsAggTradeServe("ETHBTC", func(event *WsAggTradeEvent) {
		e := &WsAggTradeEvent{
			Event:                 "aggTrade",
			Time:                  1499405254326,
			Symbol:                "ETHBTC",
			AggTradeID:            70232,
			Price:                 "0.10281118",
			Quantity:              "8.15632997",
			FirstBreakdownTradeID: 77489,
			LastBreakdownTradeID:  77489,
			TradeTime:             1499405254324,
			IsBuyerMaker:          false,
		}
		s.assertWsAggTradeEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketTestSuite) assertWsAggTradeEventEqual(e, a *WsAggTradeEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.AggTradeID, a.AggTradeID, "AggTradeID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.FirstBreakdownTradeID, a.FirstBreakdownTradeID, "FirstBreakdownTradeID")
	r.Equal(e.LastBreakdownTradeID, a.LastBreakdownTradeID, "LastBreakdownTradeID")
	r.Equal(e.TradeTime, a.TradeTime, "TradeTime")
	r.Equal(e.IsBuyerMaker, a.IsBuyerMaker, "IsBuyerMaker")
}

func (s *websocketTestSuite) TestWsAllMarketTickersStatServe() {
	websocketStreamClient := NewWebsocketStreamClient(false, "wss://stream.testnet.binance.vision")
	data := []byte(`[{
  		"e": "24hrTicker",
  		"E": 123456789,
  		"s": "BNBBTC",
  		"p": "0.0015",
  		"P": "250.00",
  		"w": "0.0018",
  		"x": "0.0009",
  		"c": "0.0025",
  		"Q": "10",
  		"b": "0.0024",
  		"B": "10",
  		"a": "0.0026",
  		"A": "100",
  		"o": "0.0010",
  		"h": "0.0026",
  		"l": "0.0010",
  		"v": "10000",
  		"q": "18",
 		"O": 0,
  		"C": 86400000,
  		"F": 0,
  		"L": 18150,
  		"n": 18151
	},{
  		"e": "24hrTicker",
  		"E": 123456789,
  		"s": "ETHBTC",
  		"p": "0.0015",
  		"P": "250.00",
  		"w": "0.0018",
  		"x": "0.0009",
  		"c": "0.0025",
  		"Q": "10",
  		"b": "0.0024",
  		"B": "10",
  		"a": "0.0026",
  		"A": "100",
  		"o": "0.0010",
  		"h": "0.0026",
  		"l": "0.0010",
  		"v": "10000",
  		"q": "18",
 		"O": 0,
  		"C": 86400000,
  		"F": 0,
  		"L": 18150,
  		"n": 18151
	}]`)
	fakeErrMsg := "fake error"
	s.mockWsServe(data, errors.New(fakeErrMsg))
	defer s.assertWsServe()

	doneC, stopC, err := websocketStreamClient.WsAllMarketTickersStatServe(func(event WsAllMarketTickersStatEvent) {
		e := WsAllMarketTickersStatEvent{
			&WsMarketTickerStatEvent{
				Event:              "24hrTicker",
				Time:               123456789,
				Symbol:             "BNBBTC",
				PriceChange:        "0.0015",
				PriceChangePercent: "250.00",
				WeightedAvgPrice:   "0.0018",
				PrevClosePrice:     "0.0009",
				LastPrice:          "0.0025",
				CloseQty:           "10",
				BidPrice:           "0.0024",
				BidQty:             "10",
				AskPrice:           "0.0026",
				AskQty:             "100",
				OpenPrice:          "0.0010",
				HighPrice:          "0.0026",
				LowPrice:           "0.0010",
				BaseVolume:         "10000",
				QuoteVolume:        "18",
				OpenTime:           0,
				CloseTime:          86400000,
				FirstID:            0,
				LastID:             18150,
				Count:              18151,
			},
			&WsMarketTickerStatEvent{
				Event:              "24hrTicker",
				Time:               123456789,
				Symbol:             "ETHBTC",
				PriceChange:        "0.0015",
				PriceChangePercent: "250.00",
				WeightedAvgPrice:   "0.0018",
				PrevClosePrice:     "0.0009",
				LastPrice:          "0.0025",
				CloseQty:           "10",
				BidPrice:           "0.0024",
				BidQty:             "10",
				AskPrice:           "0.0026",
				AskQty:             "100",
				OpenPrice:          "0.0010",
				HighPrice:          "0.0026",
				LowPrice:           "0.0010",
				BaseVolume:         "10000",
				QuoteVolume:        "18",
				OpenTime:           0,
				CloseTime:          86400000,
				FirstID:            0,
				LastID:             18150,
				Count:              18151,
			},
		}
		s.assertWsAllMarketTickersStatEventEqual(e, event)
	}, func(err error) {
		s.r().EqualError(err, fakeErrMsg)
	})
	s.r().NoError(err)
	stopC <- struct{}{}
	<-doneC
}

func (s *websocketTestSuite) assertWsAllMarketTickersStatEventEqual(e, a WsAllMarketTickersStatEvent) {
	for i := range e {
		s.assertWsMarketTickerStatEventEqual(e[i], a[i])
	}
}

func (s *websocketTestSuite) assertWsMarketTickerStatEventEqual(e, a *WsMarketTickerStatEvent) {
	r := s.r()
	r.Equal(e.Event, a.Event, "Event")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.PriceChange, a.PriceChange, "PriceChange")
	r.Equal(e.PriceChangePercent, a.PriceChangePercent, "PriceChangePercent")
	r.Equal(e.WeightedAvgPrice, a.WeightedAvgPrice, "WeightedAvgPrice")
	r.Equal(e.PrevClosePrice, a.PrevClosePrice, "PrevClosePrice")
	r.Equal(e.LastPrice, a.LastPrice, "LastPrice")
	r.Equal(e.CloseQty, a.CloseQty, "CloseQty")
	r.Equal(e.BidPrice, a.BidPrice, "BidPrice")
	r.Equal(e.BidQty, a.BidQty, "BidQty")
	r.Equal(e.AskPrice, a.AskPrice, "AskPrice")
	r.Equal(e.AskQty, a.AskQty, "AskQty")
	r.Equal(e.OpenPrice, a.OpenPrice, "OpenPrice")
	r.Equal(e.HighPrice, a.HighPrice, "HighPrice")
	r.Equal(e.LowPrice, a.LowPrice, "LowPrice")
	r.Equal(e.BaseVolume, a.BaseVolume, "BaseVolume")
	r.Equal(e.QuoteVolume, a.QuoteVolume, "QuoteVolume")
	r.Equal(e.OpenTime, a.OpenTime, "OpenTime")
	r.Equal(e.CloseTime, a.CloseTime, "CloseTime")
	r.Equal(e.FirstID, a.FirstID, "FirstID")
	r.Equal(e.LastID, a.LastID, "LastID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
}
