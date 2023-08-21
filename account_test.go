package binance_connector

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"
)

type accountTestSuite struct {
	baseTestSuite
}

func TestAccount(t *testing.T) {
	suite.Run(t, new(accountTestSuite))
}

func (s *accountTestSuite) TestGetAccountInfo() {
	data := []byte(`{
			"makerCommission": 15,
			"takerCommission": 15,
			"buyerCommission": 0,
			"sellerCommission": 0,
			"canTrade": true,
			"canWithdraw": true,
			"canDeposit": true,
			"updateTime": 123456789,
			"accountType": "SPOT",
			"balances": [
					{
							"asset": "BTC",
							"free": "4723846.89208129",
							"locked": "0.00000000"
					},
					{
							"asset": "LTC",
							"free": "4763368.68006011",
							"locked": "0.00000000"
					}
			],
			"permissions": [
				"SPOT"
			]
  }`)
	s.mockDo(data, nil)
	defer s.assertDo()
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetAccountService().Do(newContext())
	s.r().NoError(err)
	e := &AccountResponse{
		MakerCommission:  15,
		TakerCommission:  15,
		BuyerCommission:  0,
		SellerCommission: 0,
		CanTrade:         true,
		CanWithdraw:      true,
		CanDeposit:       true,
		UpdateTime:       123456789,
		AccountType:      "SPOT",
		Balances: []Balance{
			{
				Asset:  "BTC",
				Free:   "4723846.89208129",
				Locked: "0.00000000",
			},
			{
				Asset:  "LTC",
				Free:   "4763368.68006011",
				Locked: "0.00000000",
			},
		},
		Permissions: []string{"SPOT"},
	}
	s.assertAccountEqual(e, res)
}

func (s *accountTestSuite) assertAccountEqual(e, a *AccountResponse) {
	r := s.r()
	r.Equal(e.MakerCommission, a.MakerCommission, "MakerCommission")
	r.Equal(e.TakerCommission, a.TakerCommission, "TakerCommission")
	r.Equal(e.BuyerCommission, a.BuyerCommission, "BuyerCommission")
	r.Equal(e.SellerCommission, a.SellerCommission, "SellerCommission")
	r.Equal(e.CanTrade, a.CanTrade, "CanTrade")
	r.Equal(e.CanWithdraw, a.CanWithdraw, "CanWithdraw")
	r.Equal(e.CanDeposit, a.CanDeposit, "CanDeposit")
	r.Len(a.Balances, len(e.Balances))
	for i := 0; i < len(a.Balances); i++ {
		r.Equal(e.Balances[i].Asset, a.Balances[i].Asset, "Asset")
		r.Equal(e.Balances[i].Free, a.Balances[i].Free, "Free")
		r.Equal(e.Balances[i].Locked, a.Balances[i].Locked, "Locked")
	}
}

func (s *accountTestSuite) TestGetMyTrades() {
	data := []byte(`[
        {
			"symbol": "BNBBTC",
            "id": 28457,
            "orderId": 12345,
			"orderListId": -1,
            "price": "4.00000100",
			"qty": "12.00000000",
			"quoteQty": "48.000012",
            "commission": "10.10000000",
            "commissionAsset": "BNB",
            "time": 1499865549590,
            "isBuyer": true,
            "isMaker": false,
            "isBestMatch": true
        }
    ]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BNBBTC"
	limit := 3
	fromID := int64(1)
	startTime := uint64(1499865549590)
	endTime := uint64(1499865549590)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"startTime": startTime,
			"endTime":   endTime,
			"limit":     limit,
			"fromId":    fromID,
		})
		s.assertRequestEqual(e, r)
	})

	trades, err := s.client.NewGetMyTradesService().Symbol(symbol).
		StartTime(startTime).EndTime(endTime).
		Limit(limit).FromId(fromID).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(trades, 1)
	e := &AccountTradeListResponse{
		Id:              28457,
		Symbol:          "BNBBTC",
		OrderId:         12345,
		OrderListId:     -1,
		Price:           "4.00000100",
		Quantity:        "12.00000000",
		QuoteQuantity:   "48.000012",
		Commission:      "10.10000000",
		CommissionAsset: "BNB",
		Time:            1499865549590,
		IsBuyer:         true,
		IsMaker:         false,
		IsBestMatch:     true,
	}
	s.assertMyTradesEqual(e, trades[0])
}

func (s *baseTestSuite) assertMyTradesEqual(e, a *AccountTradeListResponse) {
	r := s.r()
	r.Equal(e.Id, a.Id, "ID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderId, a.OrderId, "OrderID")
	r.Equal(e.OrderListId, a.OrderListId, "OrderListId")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.QuoteQuantity, a.QuoteQuantity, "QuoteQuantity")
	r.Equal(e.Commission, a.Commission, "Commission")
	r.Equal(e.CommissionAsset, a.CommissionAsset, "CommissionAsset")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.IsBuyer, a.IsBuyer, "IsBuyer")
	r.Equal(e.IsMaker, a.IsMaker, "IsMaker")
	r.Equal(e.IsBestMatch, a.IsBestMatch, "IsBestMatch")
}

func (s *accountTestSuite) TestNewOrder() {
	data := []byte(`{
		"symbol": "BTCUSDT",
		"orderId": 28,
		"orderListId": -1,
		"clientOrderId": "6gCrw2kRUAF9CvJDGP16IP",
		"transactTime": 1507725176595,
		"price": "0.00000000",
		"origQty": "10.00000000",
		"executedQty": "10.00000000",
		"cummulativeQuoteQty": "10.00000000",
		"status": "FILLED",
		"timeInForce": "GTC",
		"type": "MARKET",
		"side": "SELL",
		"workingTime": 1507725176595,
		"selfTradePreventionMode": "NONE"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSDT"
	side := "SELL"
	orderType := "MARKET"
	quantity := 10.0
	clientOrderId := "6gCrw2kRUAF9CvJDGP16IP"
	respType := "RESULT"

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":           symbol,
			"side":             side,
			"type":             orderType,
			"quantity":         quantity,
			"newClientOrderId": clientOrderId,
			"newOrderRespType": respType,
		})
		s.assertRequestEqual(e, r)
		inputQuantity := strconv.FormatFloat(quantity, 'f', -1, 64) // Convert the quantity from float to string
		apiQuantity := r.query.Get("quantity")
		s.Equal(inputQuantity, apiQuantity, "User-input quantity does not match API value") // Check value of quantity being sent is same as input
	})

	orderResp, err := s.client.NewCreateOrderService().Symbol(symbol).
		Side(side).Type(orderType).
		Quantity(quantity).
		NewClientOrderId(clientOrderId).
		NewOrderRespType(respType).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.NotNil(orderResp)

	expectedResp := &CreateOrderResponseRESULT{
		Symbol:                  "BTCUSDT",
		OrderId:                 28,
		OrderListId:             -1,
		ClientOrderId:           "6gCrw2kRUAF9CvJDGP16IP",
		TransactTime:            1507725176595,
		Price:                   "0.00000000",
		OrigQty:                 "10.00000000",
		ExecutedQty:             "10.00000000",
		CumulativeQuoteQty:      "10.00000000",
		Status:                  "FILLED",
		TimeInForce:             "GTC",
		Type:                    "MARKET",
		Side:                    "SELL",
		WorkingTime:             1507725176595,
		SelfTradePreventionMode: "NONE",
	}

	s.assertCreateOrderResponseEqual(expectedResp, orderResp.(*CreateOrderResponseRESULT))
}

func (s *baseTestSuite) assertCreateOrderResponseEqual(e, a *CreateOrderResponseRESULT) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderId, a.OrderId, "OrderId")
	r.Equal(e.OrderListId, a.OrderListId, "OrderListId")
	r.Equal(e.ClientOrderId, a.ClientOrderId, "ClientOrderId")
	r.Equal(e.TransactTime, a.TransactTime, "TransactTime")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.OrigQty, a.OrigQty, "OrigQty")
	r.Equal(e.ExecutedQty, a.ExecutedQty, "ExecutedQty")
	r.Equal(e.CumulativeQuoteQty, a.CumulativeQuoteQty, "CumulativeQuoteQty")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.WorkingTime, a.WorkingTime, "WorkingTime")
	r.Equal(e.SelfTradePreventionMode, a.SelfTradePreventionMode, "SelfTradePreventionMode")
}

func (s *accountTestSuite) TestCancelOrder() {
	data := []byte(`{
		"symbol": "BTCUSDT",
		"origClientOrderId": "my_order_id_123",
		"orderId": 123456789,
		"orderListId": 10,
		"clientOrderId": "cancel_order_id_456",
		"price": "0.001",
		"origQty": "1.00000000",
		"executedQty": "0.00000000",
		"cumulativeQuoteQty": "0.00000000",
		"status": "CANCELED",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "BUY",
		"selfTradePrevention": "DECREMENT_AND_CANCEL"
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest()
		e.setParam("symbol", "BTCUSDT")
		e.setParam("origClientOrderId", "my_order_id_123")
		e.setParam("newClientOrderId", "cancel_order_id_456")
		s.assertRequestEqual(e, r)
	})
	ctx := context.Background()
	res, err := s.client.NewCancelOrderService().Symbol("BTCUSDT").OrigClientOrderId("my_order_id_123").NewClientOrderId("cancel_order_id_456").Do(ctx)
	s.r().NoError(err)

	e := &CancelOrderResponse{
		Symbol:              "BTCUSDT",
		OrigClientOrderId:   "my_order_id_123",
		OrderId:             123456789,
		OrderListId:         10,
		ClientOrderId:       "cancel_order_id_456",
		Price:               "0.001",
		OrigQty:             "1.00000000",
		ExecutedQty:         "0.00000000",
		CumulativeQuoteQty:  "0.00000000",
		Status:              "CANCELED",
		TimeInForce:         "GTC",
		Type:                "LIMIT",
		Side:                "BUY",
		SelfTradePrevention: "DECREMENT_AND_CANCEL",
	}
	s.assertCancelOrderEqual(e, res)
}

func (s *accountTestSuite) assertCancelOrderEqual(e, a *CancelOrderResponse) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrigClientOrderId, a.OrigClientOrderId, "OrigClientOrderId")
	r.Equal(e.OrderId, a.OrderId, "OrderId")
	r.Equal(e.OrderListId, a.OrderListId, "OrderListId")
	r.Equal(e.ClientOrderId, a.ClientOrderId, "ClientOrderId")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.OrigQty, a.OrigQty, "OrigQty")
	r.Equal(e.ExecutedQty, a.ExecutedQty, "ExecutedQty")
	r.Equal(e.CumulativeQuoteQty, a.CumulativeQuoteQty, "CumulativeQuoteQty")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.TimeInForce, a.TimeInForce, "TimeInForce")
	r.Equal(e.Type, a.Type, "Type")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.SelfTradePrevention, a.SelfTradePrevention, "SelfTradePrevention")
}

func (s *accountTestSuite) TestListRateLimit() {
	data := []byte(`
	[
		{
			"rateLimitType": "ORDERS",
			"interval": "SECOND",
			"intervalNum": 10,
			"limit": 10000,
			"count": 0
		},
		{
			"rateLimitType": "RAW_REQUESTS",
			"interval": "MINUTE",
			"intervalNum": 5,
			"limit": 5000,
			"count": 100
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	limits, err := s.client.NewGetQueryCurrentOrderCountUsageService().Do(context.Background())
	s.r().NoError(err)
	rows := limits

	s.Len(rows, 2)
	s.assertRateLimitServiceEqual(&QueryCurrentOrderCountUsageResponse{
		RateLimitType: "ORDERS",
		Interval:      "SECOND",
		IntervalNum:   10,
		Limit:         10000,
		Count:         0,
	},
		rows[0])
	s.assertRateLimitServiceEqual(&QueryCurrentOrderCountUsageResponse{
		RateLimitType: "RAW_REQUESTS",
		Interval:      "MINUTE",
		IntervalNum:   5,
		Limit:         5000,
		Count:         100,
	},
		rows[1])
}

func (a *accountTestSuite) assertRateLimitServiceEqual(expected, other *QueryCurrentOrderCountUsageResponse) {
	r := a.r()

	r.EqualValues(expected, other)
}

func (s *accountTestSuite) TestGetOrder() {
	data := []byte(`{
		"symbol": "BTCUSDT",
		"orderId": 12345,
		"orderListId": -1,
		"clientOrderId": "abcde12345",
		"price": "100.00",
		"origQty": "10.00",
		"executedQty": "0.00",
		"cumulativeQuoteQty": "0.00",
		"status": "NEW",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "BUY",
		"stopPrice": "0.00",
		"icebergQty": "0.00",
		"time": 1499827319559,
		"updateTime": 1499827319559,
		"isWorking": true,
		"origQuoteOrderQty": "0.000000"
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	res, err := s.client.NewGetOrderService().Symbol("BTCUSDT").OrderId(12345).Do(context.Background())
	s.r().NoError(err)

	s.assertGetOrderResponseEqual(&GetOrderResponse{
		Symbol:             "BTCUSDT",
		OrderId:            12345,
		OrderListId:        -1,
		ClientOrderId:      "abcde12345",
		Price:              "100.00",
		OrigQty:            "10.00",
		ExecutedQty:        "0.00",
		CumulativeQuoteQty: "0.00",
		Status:             "NEW",
		TimeInForce:        "GTC",
		Type:               "LIMIT",
		Side:               "BUY",
		StopPrice:          "0.00",
		IcebergQty:         "0.00",
		Time:               1499827319559,
		UpdateTime:         1499827319559,
		IsWorking:          true,
		OrigQuoteOrderQty:  "0.000000",
	}, res)
}

func (s *accountTestSuite) assertGetOrderResponseEqual(expected, other *GetOrderResponse) {
	r := s.r()

	r.EqualValues(expected, other)
}

func (s *accountTestSuite) TestQueryOCOService() {
	data := []byte(`{
		"orderListId": 10,
		"contingencyType": "OCO",
		"listStatusType": "EXEC_STARTED",
		"listOrderStatus": "ALL_DONE",
		"listClientOrderId": "test-list-client-order-id",
		"transactionTime": 123456789,
		"symbol": "BTCUSDT",
		"orders": [
			{
				"symbol": "BTCUSDT",
				"orderId": 1001,
				"clientOrderId": "test-client-order-id-1"
			},
			{
				"symbol": "BTCUSDT",
				"orderId": 1002,
				"clientOrderId": "test-client-order-id-2"
			}
		]
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	expected := &OCOResponse{
		OrderListId:       10,
		ContingencyType:   "OCO",
		ListStatusType:    "EXEC_STARTED",
		ListOrderStatus:   "ALL_DONE",
		ListClientOrderId: "test-list-client-order-id",
		TransactionTime:   123456789,
		Symbol:            "BTCUSDT",
		Orders: []struct {
			Symbol        string `json:"symbol"`
			OrderId       int64  `json:"orderId"`
			ClientOrderId string `json:"clientOrderId"`
		}{
			{
				Symbol:        "BTCUSDT",
				OrderId:       1001,
				ClientOrderId: "test-client-order-id-1",
			},
			{
				Symbol:        "BTCUSDT",
				OrderId:       1002,
				ClientOrderId: "test-client-order-id-2",
			},
		},
	}

	oco, err := s.client.NewQueryOCOService().
		OrderListId(10).
		OrigClientOrderId("test-client-order-id").
		Do(context.Background())

	s.r().NoError(err)
	s.r().Equal(expected, oco)
}

func (s *accountTestSuite) TestQueryAllOCOService() {
	data := []byte(`
	[
		{
			"orderListId": 29,
			"contingencyType": "OCO",
			"listStatusType": "EXEC_STARTED",
			"listOrderStatus": "EXECUTING",
			"listClientOrderId": "amEEAXryFzFwYF1FeRpUoZ",
			"transactionTime": 1565245913483,
			"symbol": "LTCBTC",
			"orders": [
				{
					"symbol": "LTCBTC",
					"orderId": 4,
					"clientOrderId": "oD7aesZqjEGlZrbtRpy5zB"
				},
				{
					"symbol": "LTCBTC",
					"orderId": 5,
					"clientOrderId": "Jr1h6xirOxgeJOUuYQS7V3"
				}
			]
		},
		{
			"orderListId": 28,
			"contingencyType": "OCO",
			"listStatusType": "EXEC_STARTED",
			"listOrderStatus": "EXECUTING",
			"listClientOrderId": "hG7hFNxJV6cZy3Ze4AUT4d",
			"transactionTime": 1565245913407,
			"symbol": "LTCBTC",
			"orders": [
				{
					"symbol": "LTCBTC",
					"orderId": 2,
					"clientOrderId": "j6lFOfbmFMRjTYA7rRJ0LP"
				},
				{
					"symbol": "LTCBTC",
					"orderId": 3,
					"clientOrderId": "z0KCjOdditiLS5ekAFtK81"
				}
			]
		}
	]`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewQueryAllOCOService().
		FromId(123456).
		StartTime(1616976140402).
		Limit(100).
		Do(context.Background())

	s.r().NoError(err)
	s.r().Equal(2, len(resp))
	s.r().Equal(int64(29), resp[0].OrderListId)
	s.r().Equal(int64(28), resp[1].OrderListId)
	s.r().Equal("OCO", resp[0].ContingencyType)
	s.r().Equal("OCO", resp[1].ContingencyType)
	s.r().Equal("EXEC_STARTED", resp[0].ListStatusType)
	s.r().Equal("EXEC_STARTED", resp[1].ListStatusType)
	s.r().Equal("EXECUTING", resp[0].ListOrderStatus)
	s.r().Equal("EXECUTING", resp[1].ListOrderStatus)
	s.r().Equal("amEEAXryFzFwYF1FeRpUoZ", resp[0].ListClientOrderId)
	s.r().Equal("hG7hFNxJV6cZy3Ze4AUT4d", resp[1].ListClientOrderId)
	s.r().Equal(uint64(1565245913483), resp[0].TransactionTime)
	s.r().Equal(uint64(1565245913407), resp[1].TransactionTime)
	s.r().Equal("LTCBTC", resp[0].Symbol)
	s.r().Equal("LTCBTC", resp[1].Symbol)
	s.r().Equal(2, len(resp[0].Orders))
	s.r().Equal(2, len(resp[1].Orders))
	s.r().Equal("LTCBTC", resp[0].Orders[0].Symbol)
	s.r().Equal("LTCBTC", resp[0].Orders[1].Symbol)
	s.r().Equal("LTCBTC", resp[1].Orders[0].Symbol)
	s.r().Equal("LTCBTC", resp[1].Orders[1].Symbol)
	s.r().Equal(int64(4), resp[0].Orders[0].OrderId)
	s.r().Equal(int64(5), resp[0].Orders[1].OrderId)
	s.r().Equal(int64(2), resp[1].Orders[0].OrderId)
	s.r().Equal(int64(3), resp[1].Orders[1].OrderId)
	s.r().Equal("oD7aesZqjEGlZrbtRpy5zB", resp[0].Orders[0].ClientOrderId)
	s.r().Equal("Jr1h6xirOxgeJOUuYQS7V3", resp[0].Orders[1].ClientOrderId)
	s.r().Equal("j6lFOfbmFMRjTYA7rRJ0LP", resp[1].Orders[0].ClientOrderId)
	s.r().Equal("z0KCjOdditiLS5ekAFtK81", resp[1].Orders[1].ClientOrderId)
}

func (s *accountTestSuite) TestGetOpenOrders() {
	data := []byte(`
	[
		{
			"symbol": "BTCUSDT",
			"orderId": 10,
			"clientOrderId": "my-order-id",
			"price": "100.0",
			"origQty": "1.0",
			"executedQty": "0.0",
			"status": "NEW",
			"timeInForce": "GTC",
			"type": "LIMIT",
			"side": "BUY",
			"stopPrice": "0.0",
			"icebergQty": "0.0",
			"time": 1613450271000,
			"updateTime": 1613450271000,
			"isWorking": true,
			"origQuoteOrderQty": "100.0",
			"selfTradePreventionMode": "DECREMENT_AND_CANCEL"
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetOpenOrdersService().
		Symbol("BTCUSDT").
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 1)
	s.Equal("BTCUSDT", resp[0].Symbol)
	s.Equal(int64(10), resp[0].OrderId)
	s.Equal("my-order-id", resp[0].ClientOrderId)
	s.Equal("100.0", resp[0].Price)
	s.Equal("1.0", resp[0].OrigQty)
	s.Equal("0.0", resp[0].ExecutedQty)
	s.Equal("NEW", resp[0].Status)
	s.Equal("GTC", resp[0].TimeInForce)
	s.Equal("LIMIT", resp[0].Type)
	s.Equal("BUY", resp[0].Side)
	s.Equal("0.0", resp[0].StopPrice)
	s.Equal("0.0", resp[0].IcebergQty)
	s.Equal(uint64(1613450271000), resp[0].Time)
	s.Equal(uint64(1613450271000), resp[0].UpdateTime)
	s.True(resp[0].IsWorking)
	s.Equal("100.0", resp[0].OrigQuoteOrderQty)
	s.Equal("DECREMENT_AND_CANCEL", resp[0].SelfTradePreventionMode)
}

func (s *accountTestSuite) TestGetAllOrders() {
	data := []byte(`
	[
		{
			"symbol": "BTCUSDT",
			"orderId": 3127658,
			"orderListId": -1,
			"clientOrderId": "U3BbvrQe2upXyv9qgPUByL",
			"price": "25000.00000000",
			"origQty": "1.00000000",
			"executedQty": "0.00000000",
			"cumulativeQuoteQty": "0.00000000",
			"status": "NEW",
			"timeInForce": "GTC",
			"type": "LIMIT",
			"side": "BUY",
			"stopPrice": "0.00000000",
			"icebergQty": "0.00000000",
			"time": 1617167610255,
			"updateTime": 1617167610255,
			"isWorking": true,
			"origQuoteOrderQty": "25000.00000000",
			"workingTime": 0,
			"selfTradePreventionMode": "DECREMENT_AND_CANCEL",
			"preventedQuantity": "0.00000000",
			"preventedMatchId": 0
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetAllOrdersService().
		Symbol("BTCUSDT").
		OrderId(3127658).
		StartTime(1617167610000).
		EndTime(1617254010000).
		Limit(1000).
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 1)
	s.Equal("BTCUSDT", resp[0].Symbol)
	s.Equal(int64(3127658), resp[0].OrderId)
	s.Equal(int64(-1), resp[0].OrderListId)
	s.Equal("U3BbvrQe2upXyv9qgPUByL", resp[0].ClientOrderId)
	s.Equal("25000.00000000", resp[0].Price)
	s.Equal("1.00000000", resp[0].OrigQty)
	s.Equal("0.00000000", resp[0].ExecutedQty)
	s.Equal("0.00000000", resp[0].CumulativeQuoteQty)
	s.Equal("NEW", resp[0].Status)
	s.Equal("GTC", resp[0].TimeInForce)
	s.Equal("LIMIT", resp[0].Type)
	s.Equal("BUY", resp[0].Side)
	s.Equal("0.00000000", resp[0].StopPrice)
	s.Equal("0.00000000", resp[0].IcebergQty)
	s.Equal(uint64(1617167610255), resp[0].Time)
	s.Equal(uint64(1617167610255), resp[0].UpdateTime)
	s.True(resp[0].IsWorking)
	s.Equal("25000.00000000", resp[0].OrigQuoteOrderQty)
	s.Equal(uint64(0), resp[0].WorkingTime)
	s.Equal("DECREMENT_AND_CANCEL", resp[0].SelfTradePreventionMode)
	s.Equal("0.00000000", resp[0].PreventedQuantity)
	s.Equal(int64(0), resp[0].PreventedMatchId)
}

func (s *accountTestSuite) TestCancelOpenOrders() {
	data := []byte(`[
		{
			"symbol": "BTCUSDT",
			"origClientOrderId": "1uL7mdU6TlTzTqTddTfNhV",
			"orderId": 123456,
			"orderListId": -1,
			"clientOrderId": "pXLV6Hz6mprAcVYpLkd1KH",
			"price": "0.00000000",
			"origQty": "1.00000000",
			"executedQty": "0.00000000",
			"cumulativeQuoteQty": "0.00000000",
			"status": "CANCELED",
			"timeInForce": "GTC",
			"type": "LIMIT",
			"side": "BUY",
			"selfTradePrevention": "DECREASE_AND_CANCEL"
		},
		{
			"symbol": "ETHUSDT",
			"origClientOrderId": "2uL7mdU6TlTzTqTddTfNhV",
			"orderId": 123457,
			"orderListId": -1,
			"clientOrderId": "qXLV6Hz6mprAcVYpLkd1KH",
			"price": "0.00000000",
			"origQty": "2.00000000",
			"executedQty": "0.00000000",
			"cumulativeQuoteQty": "0.00000000",
			"status": "CANCELED",
			"timeInForce": "GTC",
			"type": "LIMIT",
			"side": "SELL",
			"selfTradePrevention": "DECREASE_AND_CANCEL"
		}
	]`)

	s.mockDo(data, nil)
	defer s.assertDo()

	res, err := s.client.NewCancelOpenOrdersService().
		Symbol("BTCUSDT").
		Do(context.Background())

	s.r().NoError(err)
	s.Len(res, 2)
	s.Equal("BTCUSDT", res[0].Symbol)
	s.Equal("1uL7mdU6TlTzTqTddTfNhV", res[0].OrigClientOrderId)
	s.Equal(int64(123456), res[0].OrderId)
	s.Equal(int64(-1), res[0].OrderListId)
	s.Equal("pXLV6Hz6mprAcVYpLkd1KH", res[0].ClientOrderId)
	s.Equal("0.00000000", res[0].Price)
	s.Equal("1.00000000", res[0].OrigQty)
	s.Equal("0.00000000", res[0].ExecutedQty)
	s.Equal("0.00000000", res[0].CumulativeQuoteQty)
	s.Equal("CANCELED", res[0].Status)
	s.Equal("GTC", res[0].TimeInForce)
	s.Equal("LIMIT", res[0].Type)
	s.Equal("BUY", res[0].Side)
	s.Equal("DECREASE_AND_CANCEL", res[0].SelfTradePrevention)

	s.Equal("ETHUSDT", res[1].Symbol)
	s.Equal("2uL7mdU6TlTzTqTddTfNhV", res[1].OrigClientOrderId)
	s.Equal(int64(123457), res[1].OrderId)
	s.Equal(int64(-1), res[1].OrderListId)
	s.Equal("qXLV6Hz6mprAcVYpLkd1KH", res[1].ClientOrderId)
	s.Equal("0.00000000", res[1].Price)
	s.Equal("2.00000000", res[1].OrigQty)
	s.Equal("0.00000000", res[1].ExecutedQty)
	s.Equal("0.00000000", res[1].CumulativeQuoteQty)
	s.Equal("CANCELED", res[1].Status)
	s.Equal("GTC", res[1].TimeInForce)
	s.Equal("LIMIT", res[1].Type)
	s.Equal("SELL", res[1].Side)
	s.Equal("DECREASE_AND_CANCEL", res[1].SelfTradePrevention)
}

func (s *accountTestSuite) TestQueryOpenOCOService() {
	data := []byte(`
	{
		"orderListId": 1,
		"contingencyType": "OCO",
		"listStatusType": "EXEC_STARTED",
		"listOrderStatus": "EXECUTING",
		"listClientOrderId": "test-list-client-order-id",
		"transactionTime": 1613450271000,
		"symbol": "BTCUSDT",
		"orders": [
			{
				"symbol": "BTCUSDT",
				"orderId": 123,
				"clientOrderId": "test-client-order-id-1"
			},
			{
				"symbol": "BTCUSDT",
				"orderId": 456,
				"clientOrderId": "test-client-order-id-2"
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewQueryOpenOCOService().Do(context.Background())

	s.r().NoError(err)
	s.NotNil(resp)
	s.Equal(int64(1), resp.OrderListId)
	s.Equal("OCO", resp.ContingencyType)
	s.Equal("EXEC_STARTED", resp.ListStatusType)
	s.Equal("EXECUTING", resp.ListOrderStatus)
	s.Equal("test-list-client-order-id", resp.ListClientOrderId)
	s.Equal(uint64(1613450271000), resp.TransactionTime)
	s.Equal("BTCUSDT", resp.Symbol)
	s.Len(resp.Orders, 2)
	s.Equal("BTCUSDT", resp.Orders[0].Symbol)
	s.Equal(int64(123), resp.Orders[0].OrderId)
	s.Equal("test-client-order-id-1", resp.Orders[0].ClientOrderId)
	s.Equal("BTCUSDT", resp.Orders[1].Symbol)
	s.Equal(int64(456), resp.Orders[1].OrderId)
	s.Equal("test-client-order-id-2", resp.Orders[1].ClientOrderId)
}

func (s *accountTestSuite) TestGetQueryPreventedMatches() {
	data := []byte(`{
		"preventedMatches": [
			{
				"symbol": "BTCUSDT",
				"preventedMatchId": 123,
				"takerOrderId": 456,
				"makerOrderId": 789,
				"tradeGroupId": 101112,
				"selfTradePreventionMode": "CANCEL_BOTH",
				"price": "65000.00",
				"makerPreventedQuantity": "0.00005",
				"transactTime": 1613450271000
			}
		]
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetQueryPreventedMatchesService().
		Symbol("BTCUSDT").
		PreventMatchId(123).
		OrderId(456).
		FromPreventedMatchId(789).
		Limit(500).
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp.PreventedMatches, 1)
	s.Equal("BTCUSDT", resp.PreventedMatches[0].Symbol)
	s.Equal(int64(123), resp.PreventedMatches[0].PreventedMatchId)
	s.Equal(int64(456), resp.PreventedMatches[0].TakerOrderId)
	s.Equal(int64(789), resp.PreventedMatches[0].MakerOrderId)
	s.Equal(int64(101112), resp.PreventedMatches[0].TradeGroupId)
	s.Equal("CANCEL_BOTH", resp.PreventedMatches[0].SelfTradePreventionMode)
	s.Equal("65000.00", resp.PreventedMatches[0].Price)
	s.Equal("0.00005", resp.PreventedMatches[0].MakerPreventedQuantity)
	s.Equal(uint64(1613450271000), resp.PreventedMatches[0].TransactTime)
}
