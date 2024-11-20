package binance_connector

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type marginTestSuite struct {
	baseTestSuite
}

func TestMargin(t *testing.T) {
	suite.Run(t, new(marginTestSuite))
}

func (s *marginTestSuite) TestGetMaxBorrowable() {
	data := []byte(`{
		"amount": "1.69248805"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset": "BNBBTC",
		})
		s.assertRequestEqual(e, r)
	})

	borrowable, err := s.client.NewMarginAccountQueryMaxBorrowService().
		Asset("BNBBTC").Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &MarginAccountQueryMaxBorrowResponse{
		Amount: "1.69248805",
	}
	s.assertMaxBorrowableEqual(e, borrowable)
}

func (s *marginTestSuite) assertMaxBorrowableEqual(e, a *MarginAccountQueryMaxBorrowResponse) {
	s.r().Equal(e.Amount, a.Amount, "Amount")
}

func (s *marginTestSuite) TestGetMaxTransferable() {
	data := []byte(`{
		"amount": "3.59498107"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"asset": "BNBBTC",
		})
		s.assertRequestEqual(e, r)
	})

	transferable, err := s.client.NewMarginAccountQueryMaxTransferOutAmountService().
		Asset("BNBBTC").Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &MarginAccountQueryMaxTransferOutAmountResponse{
		Amount: "3.59498107",
	}
	s.assertMaxTransferableEqual(e, transferable)
}

func (s *marginTestSuite) assertMaxTransferableEqual(e, a *MarginAccountQueryMaxTransferOutAmountResponse) {
	s.r().Equal(e.Amount, a.Amount, "Amount")
}

func (s *marginTestSuite) TestMarginAccountAdjustCrossMaxLeverage() {
	data := []byte(`{
		"success": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	res, err := s.client.NewMarginAccountAdjustCrossMaxLeverageService().MaxLeverage(5).Do(context.Background())
	s.r().NoError(err)
	s.True(res.Success)
}

func (s *marginTestSuite) TestGetMarginPriceIndex() {
	data := []byte(`{
		"calcTime": 1562046418000,
		"price": "0.00333930",
		"symbol": "BNBBTC"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BNBBTC"
	s.assertReq(func(r *request) {
		e := newRequest()
		e.setParam("symbol", symbol)
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewQueryMarginPriceIndexService().Symbol(symbol).Do(newContext())
	s.r().NoError(err)
	e := &QueryMarginPriceIndexResponse{
		CalcTime: 1562046418000,
		Symbol:   symbol,
		Price:    "0.00333930",
	}
	s.assertMarginPriceIndexEqual(e, res)
}

func (s *marginTestSuite) TestQueryMarginAvailableInventory() {
	data := []byte(`{
		"assets": {
			"MATIC": "100000000",
			"STPT": "100000000",
			"TVK": "100000000",
			"SHIB": "97409653"
		},
		"updateTime": 1699272487
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	res, err := s.client.NewQueryMarginAvailableInventoryService().Do(context.Background())
	s.r().NoError(err)
	s.Len(res.Assets, 4)
	s.Equal("100000000", res.Assets["MATIC"])
	s.Equal("100000000", res.Assets["STPT"])
	s.Equal("100000000", res.Assets["TVK"])
	s.Equal("97409653", res.Assets["SHIB"])
	s.Equal(int64(1699272487), res.UpdateTime)
}

func (s *marginTestSuite) TestQueryLiabilityCoinLeverageBracket() {
	data := []byte(`[
	{
		"assetNames":[
			"SHIB",
			"FDUSD",
			"BTC",
			"ETH",
			"USDC"
		],          
		"rank":1,
		"brackets":[
			{
				"leverage":10,
				"maxDebt":1000000.00000000,
				"maintenanceMarginRate":0.02000000,
				"initialMarginRate":0.1112,
				"fastNum":0
			},
			{
				"leverage":3,
				"maxDebt":4000000.00000000,
				"maintenanceMarginRate":0.07000000,
				"initialMarginRate":0.5000,
				"fastNum":60000.0000000000000000
			}
		]
	}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	res, err := s.client.NewQueryLiabilityCoinLeverageBracketService().Do(context.Background())
	s.r().NoError(err)
	s.Len(res, 1)
	s.Len(res[0].Brackets, 2)
	s.Equal(1, res[0].Rank)
	s.Len(res[0].AssetNames, 5)
	s.Equal("SHIB", res[0].AssetNames[0])
	s.Equal("FDUSD", res[0].AssetNames[1])
	s.Equal("BTC", res[0].AssetNames[2])
	s.Equal("ETH", res[0].AssetNames[3])
	s.Equal("USDC", res[0].AssetNames[4])
	s.Equal(10, res[0].Brackets[0].Leverage)
	s.Equal(1000000.00000000, res[0].Brackets[0].MaxDebt)
	s.Equal(0.02000000, res[0].Brackets[0].MaintenanceMarginRate)
	s.Equal(0.1112, res[0].Brackets[0].InitialMarginRate)
	s.Equal(0.0, res[0].Brackets[0].FastNum)
	s.Equal(3, res[0].Brackets[1].Leverage)
	s.Equal(4000000.00000000, res[0].Brackets[1].MaxDebt)
	s.Equal(0.07000000, res[0].Brackets[1].MaintenanceMarginRate)
	s.Equal(0.5000, res[0].Brackets[1].InitialMarginRate)
	s.Equal(60000.0000000000000000, res[0].Brackets[1].FastNum)
}

func (s *marginTestSuite) assertMarginPriceIndexEqual(e, a *QueryMarginPriceIndexResponse) {
	r := s.r()
	r.Equal(e.CalcTime, a.CalcTime, "CalcTime")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Price, a.Price, "Price")
}

func (s *marginTestSuite) TestGetIsolatedMarginAllPairs() {
	data := []byte(`[{
        "base": "BNB",
        "isBuyAllowed": true,
        "isMarginTrade": true,
        "isSellAllowed": true,
        "quote": "BTC",
        "symbol": "BNBBTC"
    },
    {
        "base": "TRX",
        "isBuyAllowed": true,
        "isMarginTrade": true,
        "isSellAllowed": true,
        "quote": "BTC",
        "symbol": "TRXBTC"
    }]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	s.assertReq(func(r *request) {
		e := newRequest()
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewAllIsolatedMarginSymbolService().
		Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(res, 2)
	e := []*MarginIsolatedSymbolResponse{
		{
			Symbol:        "BNBBTC",
			Base:          "BNB",
			Quote:         "BTC",
			IsMarginTrade: true,
			IsBuyAllowed:  true,
			IsSellAllowed: true,
		}, {
			Symbol:        "TRXBTC",
			Base:          "TRX",
			Quote:         "BTC",
			IsMarginTrade: true,
			IsBuyAllowed:  true,
			IsSellAllowed: true,
		},
	}

	for i := 0; i < len(res); i++ {
		s.assertIsolatedMarginAllPairsEqual(e[i], res[i])
	}
}

func (s *marginTestSuite) assertIsolatedMarginAllPairsEqual(e, a *MarginIsolatedSymbolResponse) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Base, a.Base, "Base")
	r.Equal(e.Quote, a.Quote, "Quote")
	r.Equal(e.IsMarginTrade, a.IsMarginTrade, "IsMarginTrade")
	r.Equal(e.IsBuyAllowed, a.IsBuyAllowed, "IsBuyAllowed")
	r.Equal(e.IsSellAllowed, a.IsSellAllowed, "IsSellAllowed")
}

func (s *marginTestSuite) TestGetAllMarginAssets() {
	data := []byte(`
	[
		{
			"assetFullName": "Bitcoin",
			"assetName": "BTC",
			"isBorrowable": true,
			"isMortgageable": true,
			"minLoanAmt": "0.00010000",
			"maxLoanAmt": "100000.00000000",
			"minMortgageAmt": "0.00010000",
			"maxMortgageAmt": "100000.00000000",
			"asset": "BTC"
		},
		{
			"assetFullName": "Ethereum",
			"assetName": "ETH",
			"isBorrowable": true,
			"isMortgageable": true,
			"minLoanAmt": "0.01000000",
			"maxLoanAmt": "10000.00000000",
			"minMortgageAmt": "0.01000000",
			"maxMortgageAmt": "10000.00000000",
			"asset": "ETH"
		}
	]`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetAllMarginAssetsService().Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 2)
	s.Equal("Bitcoin", resp[0].AssetFullName)
	s.Equal("BTC", resp[0].AssetName)
	s.True(resp[0].IsBorrowable)
	s.True(resp[0].IsMortgageable)
	s.Equal("0.00010000", resp[0].MinLoanAmt)
	s.Equal("100000.00000000", resp[0].MaxLoanAmt)
	s.Equal("0.00010000", resp[0].MinMortgageAmt)
	s.Equal("100000.00000000", resp[0].MaxMortgageAmt)
	s.Equal("BTC", resp[0].Asset)
}

func (s *marginTestSuite) TestForceLiquidationRecordService() {
	data := []byte(`{
		"rows": [
			{
				"avgPrice": "0.00333900",
				"executedQty": "0.07000000",
				"orderId": 1311524,
				"price": "0.00333700",
				"qty": "0.07000000",
				"side": "SELL",
				"symbol": "BNBUSDT",
				"timeInForce": "GTC",
				"isIsolated": false,
				"updatedTime": 1620661623199
			},
			{
				"avgPrice": "0.00343000",
				"executedQty": "0.06000000",
				"orderId": 123344,
				"price": "0.00343200",
				"qty": "0.06000000",
				"side": "BUY",
				"symbol": "BNBUSDT",
				"timeInForce": "GTC",
				"isIsolated": false,
				"updatedTime": 1620661623167
			}
		],
		"total": 2
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	record, err := s.client.NewForceLiquidationRecordService().Do(context.Background())
	s.r().NoError(err)

	expectedRecord := &ForceLiquidationRecordResponse{
		Rows: []struct {
			AvgPrice    string `json:"avgPrice"`
			ExecutedQty string `json:"executedQty"`
			OrderId     int    `json:"orderId"`
			Price       string `json:"price"`
			Qty         string `json:"qty"`
			Side        string `json:"side"`
			Symbol      string `json:"symbol"`
			TimeInForce string `json:"timeInForce"`
			IsIsolated  bool   `json:"isIsolated"`
			UpdatedTime uint64 `json:"updatedTime"`
		}{
			{
				AvgPrice:    "0.00333900",
				ExecutedQty: "0.07000000",
				OrderId:     1311524,
				Price:       "0.00333700",
				Qty:         "0.07000000",
				Side:        "SELL",
				Symbol:      "BNBUSDT",
				TimeInForce: "GTC",
				IsIsolated:  false,
				UpdatedTime: 1620661623199,
			},
			{
				AvgPrice:    "0.00343000",
				ExecutedQty: "0.06000000",
				OrderId:     123344,
				Price:       "0.00343200",
				Qty:         "0.06000000",
				Side:        "BUY",
				Symbol:      "BNBUSDT",
				TimeInForce: "GTC",
				IsIsolated:  false,
				UpdatedTime: 1620661623167,
			},
		},
		Total: 2,
	}

	s.Len(record.Rows, 2)
	s.EqualValues(expectedRecord, record)
}

func (s *marginTestSuite) TestInterestHistory() {
	data := []byte(`
	{
		"rows": [
		  	{       
				"txId": 123,
				"interestAccuredTime": 1613450271000,
				"asset": "BTC",
				"rawAsset": "BTC",
				"principal": "1.00000000",
				"interest": "0.00000005",
				"interestRate": "0.00000100",
				"type": "NORMAL",
				"isolatedSymbol": ""
			}
		],
		"total": 1
	}	  
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewInterestHistoryService().
		Asset("BTC").
		IsolatedSymbol("BTCUSDT").
		StartTime(1613450271000).
		EndTime(1613536671000).
		Current(1).
		Size(500).
		Archived("true").
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp.Rows, 1)
	s.Equal(uint64(1613450271000), resp.Rows[0].InterestAccuredTime)
	s.Equal("BTC", resp.Rows[0].Asset)
	s.Equal("BTC", resp.Rows[0].RawAsset)
	s.Equal("1.00000000", resp.Rows[0].Principal)
	s.Equal("0.00000005", resp.Rows[0].Interest)
	s.Equal("0.00000100", resp.Rows[0].InterestRate)
	s.Equal("NORMAL", resp.Rows[0].Type)
	s.Equal("", resp.Rows[0].IsolatedSymbol)
	s.Equal(1, resp.Total)
}

func (s *marginTestSuite) TestMarginAccountAllOrderService() {
	data := []byte(`
	[
		{
			"clientOrderId": "example-client-order-id",
			"cummulativeQuoteQty": "100.00000000",
			"executedQty": "1.00000000",
			"icebergQty": "0.00000000",
			"isWorking": false,
			"orderId": 123,
			"origQty": "1.00000000",
			"price": "100.00000000",
			"side": "BUY",
			"status": "FILLED",
			"stopPrice": "0.00000000",
			"symbol": "BTCUSDT",
			"isIsolated": true,
			"time": 1613450271000,
			"updateTime": 1613450271000
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginAccountAllOrderService().
		Symbol("BTCUSDT").
		IsIsolated("TRUE").
		OrderId(123).
		StartTime(1613450271000).
		EndTime(1613450271000).
		Limit(500).
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 1)
	s.Equal("example-client-order-id", resp[0].ClientOrderId)
	s.Equal("100.00000000", resp[0].CummulativeQuoteQty)
	s.Equal("1.00000000", resp[0].ExecutedQty)
	s.Equal("0.00000000", resp[0].IcebergQty)
	s.False(resp[0].IsWorking)
	s.Equal(123, resp[0].OrderId)
	s.Equal("1.00000000", resp[0].OrigQty)
	s.Equal("100.00000000", resp[0].Price)
	s.Equal("BUY", resp[0].Side)
	s.Equal("FILLED", resp[0].Status)
	s.Equal("0.00000000", resp[0].StopPrice)
	s.Equal("BTCUSDT", resp[0].Symbol)
	s.True(resp[0].IsIsolated)
	s.Equal(uint64(1613450271000), resp[0].Time)
	s.Equal(uint64(1613450271000), resp[0].UpdateTime)
}

func (s *marginTestSuite) TestMarginAccountCancelAllOrders() {
	data := []byte(`{"symbol":"BTCUSDT","isIsolated":false,"origClientOrderId":"a0e0a6f7-8b38-44d5-9dc5-87e3b9bd2d31","orderId":123456789,"orderListId":-1,"clientOrderId":"my_order_id","price":"10000.00000000","origQty":"1.00000000","executedQty":"0.00000000","cummulativeQuoteQty":"0.00000000","status":"CANCELED","timeInForce":"GTC","type":"LIMIT","side":"BUY"}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginAccountCancelAllOrdersService().
		Symbol("BTCUSDT").
		IsIsolated("false").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("BTCUSDT", resp.Symbol)
	s.False(resp.IsIsolated)
	s.Equal("a0e0a6f7-8b38-44d5-9dc5-87e3b9bd2d31", resp.OrigClientOrderId)
	s.Equal(123456789, resp.OrderId)
	s.Equal(-1, resp.OrderListId)
	s.Equal("my_order_id", resp.ClientOrderId)
	s.Equal("10000.00000000", resp.Price)
	s.Equal("1.00000000", resp.OrigQty)
	s.Equal("0.00000000", resp.ExecutedQty)
	s.Equal("0.00000000", resp.CummulativeQuoteQty)
	s.Equal("CANCELED", resp.Status)
	s.Equal("GTC", resp.TimeInForce)
	s.Equal("LIMIT", resp.Type)
	s.Equal("BUY", resp.Side)
}

func (s *marginTestSuite) TestMarginAccountCancelOCO() {
	data := []byte(`
	{
		"orderListId": 400000005,
		"contingencyType": "OCO",
		"listStatusType": "ALL_DONE",
		"listOrderStatus": "ALL_DONE",
		"listClientOrderId": "twxFWSEZrlzfOeNHDGwOCl",
		"transactionTime": 1629187528733,
		"symbol": "BNBUSDT",
		"isIsolated": false,
		"orders": [
			{
				"symbol": "BNBUSDT",
				"orderId": 2,
				"clientOrderId": "EkxGkowgJHXEtlMpoxqtXl"
			},
			{
				"symbol": "BNBUSDT",
				"orderId": 3,
				"clientOrderId": "QnTwYPMEnzQKfUwJSaFgWo"
			}
		],
		"orderReports": [
			{
				"symbol": "BNBUSDT",
				"origClientOrderId": "my6d5q5u5SxS2fVbYJ1imM",
				"orderId": 2,
				"orderListId": 400000005,
				"clientOrderId": "EkxGkowgJHXEtlMpoxqtXl",
				"price": "3600.00000000",
				"origQty": "0.00500000",
				"executedQty": "0.00000000",
				"cummulativeQuoteQty": "0.00000000",
				"status": "CANCELED",
				"timeInForce": "GTC",
				"type": "STOP_LOSS_LIMIT",
				"side": "BUY",
				"stopPrice": "3700.00000000"
			},
			{
				"symbol": "BNBUSDT",
				"origClientOrderId": "hEgRiLFF2Cv6Gdcy6lSnfE",
				"orderId": 3,
				"orderListId": 400000005,
				"clientOrderId": "QnTwYPMEnzQKfUwJSaFgWo",
				"price": "3400.00000000",
				"origQty": "0.00500000",
				"executedQty": "0.00000000",
				"cummulativeQuoteQty": "0.00000000",
				"status": "CANCELED",
				"timeInForce": "GTC",
				"type": "LIMIT_MAKER",
				"side": "SELL"
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginAccountCancelOCOService().
		Symbol("BNBUSDT").
		OrderListId(400000005).
		ListClientOrderId("twxFWSEZrlzfOeNHDGwOCl").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(400000005, resp.OrderListId)
	s.Equal("OCO", resp.ContingencyType)
	s.Equal("ALL_DONE", resp.ListStatusType)
	s.Equal("ALL_DONE", resp.ListOrderStatus)
	s.Equal("twxFWSEZrlzfOeNHDGwOCl", resp.ListClientOrderId)
	s.Equal(uint64(1629187528733), resp.TransactionTime)
	s.Equal("BNBUSDT", resp.Symbol)
	s.False(resp.IsIsolated)
	s.Len(resp.Orders, 2)
	s.Len(resp.OrderReports, 2)
}

func (s *marginTestSuite) TestMarginAccountCancelOrder() {
	data := []byte(`{
		"symbol": "BTCUSDT",
		"isIsolated": false,
		"orderId": 12345,
		"origClientOrderId": "my_order_id",
		"clientOrderId": "new_order_id",
		"price": "100.00000000",
		"origQty": "1.00000000",
		"executedQty": "0.00000000",
		"cummulativeQuoteQty": "0.00000000",
		"status": "CANCELED",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"side": "BUY"
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginAccountCancelOrderService().
		Symbol("BTCUSDT").
		OrderId(12345).
		OrigClientOrderId("my_order_id").
		NewClientOrderId("new_order_id").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("BTCUSDT", resp.Symbol)
	s.False(resp.IsIsolated)
	s.Equal(12345, resp.OrderId)
	s.Equal("my_order_id", resp.OrigClientOrderId)
	s.Equal("new_order_id", resp.ClientOrderId)
	s.Equal("100.00000000", resp.Price)
	s.Equal("1.00000000", resp.OrigQty)
	s.Equal("0.00000000", resp.ExecutedQty)
	s.Equal("0.00000000", resp.CummulativeQuoteQty)
	s.Equal("CANCELED", resp.Status)
	s.Equal("GTC", resp.TimeInForce)
	s.Equal("LIMIT", resp.Type)
	s.Equal("BUY", resp.Side)
}

func (s *marginTestSuite) TestMarginAccountNewOCO() {
	data := []byte(`
	{
		"orderListId": 0,
		"contingencyType": "OCO",
		"listStatusType": "EXEC_STARTED",
		"listOrderStatus": "EXECUTING",
		"listClientOrderId": "xVpDCr1DvuaVwRz8QULNjF",
		"transactionTime": 1618776703825,
		"symbol": "BTCUSDT",
		"orders": [
			{
				"symbol": "BTCUSDT",
				"orderId": 678101264,
				"clientOrderId": "NCPLVlIz4vQ7YDHdkb1V5E"
			},
			{
				"symbol": "BTCUSDT",
				"orderId": 678101265,
				"clientOrderId": "OXQ2tX0TUJlDAM7JeKVcZd"
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginAccountNewOCOService().
		Symbol("BTCUSDT").
		IsIsolated("false").
		ListClientOrderId("xVpDCr1DvuaVwRz8QULNjF").
		Side("SELL").
		Quantity(1).
		LimitClientOrderId("limit-client-order-id").
		Price(40000).
		LimitIcebergQty(0.5).
		StopClientOrderId("stop-client-order-id").
		StopPrice(50000).
		StopLimitPrice(45000).
		StopIcebergQty(0.5).
		StopLimitTimeInForce("GTC").
		NewOrderRespType("FULL").
		SideEffectType("NO_SIDE_EFFECT").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(int(0), resp.OrderListId)
	s.Equal("OCO", resp.ContingencyType)
	s.Equal("EXEC_STARTED", resp.ListStatusType)
	s.Equal("EXECUTING", resp.ListOrderStatus)
	s.Equal("xVpDCr1DvuaVwRz8QULNjF", resp.ListClientOrderId)
	s.Equal(uint64(1618776703825), resp.TransactionTime)
	s.Equal("BTCUSDT", resp.Symbol)
	s.Len(resp.Orders, 2)
	s.Equal("BTCUSDT", resp.Orders[0].Symbol)
	s.Equal(int(678101264), resp.Orders[0].OrderId)
	s.Equal("NCPLVlIz4vQ7YDHdkb1V5E", resp.Orders[0].ClientOrderId)
	s.Equal("BTCUSDT", resp.Orders[1].Symbol)
	s.Equal(int(678101265), resp.Orders[1].OrderId)
	s.Equal("OXQ2tX0TUJlDAM7JeKVcZd", resp.Orders[1].ClientOrderId)
}

func (s *marginTestSuite) TestMarginAccountNewOrder() {
	data := []byte(`{
		"symbol": "BTCUSDT",
		"orderId": 28,
		"clientOrderId": "6gCrw2kRUAF9CvJDGP16IP",
		"transactTime": 1507725176595,
		"price": "1.00000000",
		"origQty": "10.00000000",
		"executedQty": "10.00000000",
		"cummulativeQuoteQty": "10.00000000",
		"status": "FILLED",
		"timeInForce": "GTC",
		"type": "MARKET",
		"isIsolated": true,
		"side": "SELL"
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginAccountNewOrderService().
		Symbol("BTCUSDT").
		Side("SELL").
		Quantity(10.0).
		Price(1.00).
		NewClientOrderId("6gCrw2kRUAF9CvJDGP16IP").
		TimeInForce("GTC").
		OrderType("MARKET").
		NewOrderRespType("RESULT").
		Do(context.Background())

	s.r().NoError(err)
	a := resp.(*MarginAccountNewOrderResponseRESULT)
	s.Equal("BTCUSDT", a.Symbol)
	s.Equal(int64(28), a.OrderId)
	s.Equal("6gCrw2kRUAF9CvJDGP16IP", a.ClientOrderId)
	s.Equal(uint64(1507725176595), a.TransactTime)
	s.Equal("1.00000000", a.Price)
	s.Equal("10.00000000", a.OrigQty)
	s.Equal("10.00000000", a.ExecutedQty)
	s.Equal("10.00000000", a.CummulativeQuoteQty)
	s.Equal("FILLED", a.Status)
	s.Equal("GTC", a.TimeInForce)
	s.Equal("MARKET", a.Type)
	s.True(a.IsIsolated)
	s.Equal("SELL", a.Side)
}

func (s *marginTestSuite) TestMarginAccountOpenOrder() {
	data := []byte(`
	[
		{
			"clientOrderId": "abc123",
			"cummulativeQuoteQty": "1.00000000",
			"executedQty": "1.00000000",
			"icebergQty": "0.00000000",
			"isWorking": false,
			"orderId": 123,
			"origQty": "1.00000000",
			"price": "10000.00000000",
			"side": "BUY",
			"status": "FILLED",
			"stopPrice": "0.00000000",
			"symbol": "BTCUSDT",
			"isIsolated": false,
			"time": 1619549055345,
			"timeInForce": "GTC",
			"type": "LIMIT",
			"updateTime": 1619549055345
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginAccountOpenOrderService().
		Symbol("BTCUSDT").
		IsIsolated("FALSE").
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 1)
	s.Equal("abc123", resp[0].ClientOrderId)
	s.Equal("1.00000000", resp[0].CummulativeQuoteQty)
	s.Equal("1.00000000", resp[0].ExecutedQty)
	s.Equal("0.00000000", resp[0].IcebergQty)
	s.Equal(false, resp[0].IsWorking)
	s.Equal(123, resp[0].OrderId)
	s.Equal("1.00000000", resp[0].OrigQty)
	s.Equal("10000.00000000", resp[0].Price)
	s.Equal("BUY", resp[0].Side)
	s.Equal("FILLED", resp[0].Status)
	s.Equal("0.00000000", resp[0].StopPrice)
	s.Equal("BTCUSDT", resp[0].Symbol)
	s.Equal(false, resp[0].IsIsolated)
	s.Equal(uint64(1619549055345), resp[0].Time)
	s.Equal("GTC", resp[0].TimeInForce)
	s.Equal("LIMIT", resp[0].OrderType)
	s.Equal(uint64(1619549055345), resp[0].UpdateTime)
}

func (s *marginTestSuite) TestMarginAccountOrder() {
	data := []byte(`
	{
		"clientOrderId": "myclientorderid",
		"cummulativeQuoteQty": "1.00000000",
		"executedQty": "1.00000000",
		"icebergQty": "0.00000000",
		"isWorking": false,
		"orderId": 12345,
		"origQty": "1.00000000",
		"price": "10000.00000000",
		"side": "BUY",
		"status": "FILLED",
		"stopPrice": "0.00000000",
		"symbol": "BTCUSDT",
		"isIsolated": true,
		"time": 1613450271000,
		"timeInForce": "GTC",
		"type": "LIMIT",
		"updateTime": 1613450271000
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginAccountOrderService().
		Symbol("BTCUSDT").
		IsIsolated("BTC").
		OrderId(12345).
		OrigClientOrderId("myclientorderid").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("myclientorderid", resp.ClientOrderId)
	s.Equal("1.00000000", resp.CummulativeQuoteQty)
	s.Equal("1.00000000", resp.ExecutedQty)
	s.Equal("0.00000000", resp.IcebergQty)
	s.False(resp.IsWorking)
	s.Equal(12345, resp.OrderId)
	s.Equal("1.00000000", resp.OrigQty)
	s.Equal("10000.00000000", resp.Price)
	s.Equal("BUY", resp.Side)
	s.Equal("FILLED", resp.Status)
	s.Equal("0.00000000", resp.StopPrice)
	s.Equal("BTCUSDT", resp.Symbol)
	s.True(resp.IsIsolated)
	s.Equal(uint64(1613450271000), resp.Time)
	s.Equal("GTC", resp.TimeInForce)
	s.Equal("LIMIT", resp.OrderType)
	s.Equal(uint64(1613450271000), resp.UpdateTime)
}

func (s *marginTestSuite) TestMarginAccountQueryOCO() {
	data := []byte(`
	{
		"orderListId": 0,
		"contingencyType": "OCO",
		"listStatusType": "EXEC_STARTED",
		"listOrderStatus": "EXECUTING",
		"listClientOrderId": "c3a3c32efeb04d40b42f6e9f2a2b094d",
		"transactionTime": 1618856676234,
		"symbol": "BTCUSDT",
		"isIsolated": false,
		"orders": [
			{
				"symbol": "BTCUSDT",
				"orderId": 2950247107,
				"clientOrderId": "O5GLM3Bkg3xy49ATtF8rKX"
			},
			{
				"symbol": "BTCUSDT",
				"orderId": 2950247108,
				"clientOrderId": "jKk3m21pJb8wU6TlvF6zAb"
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginAccountQueryOCOService().
		IsIsolated("false").
		Symbol("BTCUSDT").
		OrderListId(0).
		OrigClientOrderId("").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(0, resp.OrderListId)
	s.Equal("OCO", resp.ContingencyType)
	s.Equal("EXEC_STARTED", resp.ListStatusType)
	s.Equal("EXECUTING", resp.ListOrderStatus)
	s.Equal("c3a3c32efeb04d40b42f6e9f2a2b094d", resp.ListClientOrderId)
	s.Equal(uint64(1618856676234), resp.TransactionTime)
	s.Equal("BTCUSDT", resp.Symbol)
	s.False(resp.IsIsolated)
	s.Len(resp.Orders, 2)
	s.Equal("BTCUSDT", resp.Orders[0].Symbol)
	s.Equal(2950247107, resp.Orders[0].OrderId)
	s.Equal("O5GLM3Bkg3xy49ATtF8rKX", resp.Orders[0].ClientOrderId)
	s.Equal("BTCUSDT", resp.Orders[1].Symbol)
	s.Equal(2950247108, resp.Orders[1].OrderId)
	s.Equal("jKk3m21pJb8wU6TlvF6zAb", resp.Orders[1].ClientOrderId)
}

func (s *marginTestSuite) TestMarginAccountQueryOpenOCO() {
	data := []byte(`
	[
		{
			"orderListId": 1613450271000,
			"contingencyType": "OCO",
			"listStatusType": "EXEC_STARTED",
			"listOrderStatus": "EXECUTING",
			"listClientOrderId": "JYVppVf2SgZdG8",
			"transactionTime": 1613450271000,
			"symbol": "BTCUSDT",
			"isIsolated": true,
			"orders": [
				{
					"symbol": "BTCUSDT",
					"orderId": 123,
					"clientOrderId": "abc123"
				},
				{
					"symbol": "BTCUSDT",
					"orderId": 456,
					"clientOrderId": "def456"
				}
			]
		}
	]`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginAccountQueryOpenOCOService().
		IsIsolated("true").
		Symbol("BTCUSDT").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(1613450271000, resp[0].OrderListId)
	s.Equal("OCO", resp[0].ContingencyType)
	s.Equal("EXEC_STARTED", resp[0].ListStatusType)
	s.Equal("EXECUTING", resp[0].ListOrderStatus)
	s.Equal("JYVppVf2SgZdG8", resp[0].ListClientOrderId)
	s.Equal(uint64(1613450271000), resp[0].TransactionTime)
	s.Equal("BTCUSDT", resp[0].Symbol)
	s.True(resp[0].IsIsolated)
	s.Len(resp[0].Orders, 2)
	s.Equal("BTCUSDT", resp[0].Orders[0].Symbol)
	s.Equal(123, resp[0].Orders[0].OrderId)
	s.Equal("abc123", resp[0].Orders[0].ClientOrderId)
	s.Equal("BTCUSDT", resp[0].Orders[1].Symbol)
	s.Equal(456, resp[0].Orders[1].OrderId)
	s.Equal("def456", resp[0].Orders[1].ClientOrderId)
}

func (s *marginTestSuite) TestMarginAccountSummary() {
	data := []byte(`
	{
		"normalBar": "10",
		"marginCallBar": "20",
		"forceLiquidationBar": "30"
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginAccountSummaryService().Do(context.Background())

	s.r().NoError(err)
	s.Equal("10", resp.NormalBar)
	s.Equal("20", resp.MarginCallBar)
	s.Equal("30", resp.ForceLiquidationBar)
}

func (s *marginTestSuite) TestMarginBnbBurnStatus() {
	data := []byte(`
	{
		"spotBNBBurn": true,
		"interestBNBBurn": true
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginBnbBurnStatusService().Do(context.Background())

	s.r().NoError(err)
	s.True(resp.SpotBNBBurn)
	s.True(resp.InterestBNBBurn)
}

func (s *marginTestSuite) TestMarginCrossCollateralRatio() {
	data := []byte(`
	[
		{
			"collaterals": [
				{
					"minUsdValue": "0",
					"maxUsdValue": "10000",
					"discountRate": "0.9000"
				},
				{
					"minUsdValue": "10000",
					"maxUsdValue": "20000",
					"discountRate": "0.8500"
				}
			],
			"assetNames": [
				{
					"asset": "BTC"
				},
				{
					"asset": "ETH"
				}
			]
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginCrossCollateralRatioService().Do(context.Background())

	s.r().NoError(err)
	s.Len(resp[0].Collaterals, 2)
	s.Equal("0", resp[0].Collaterals[0].MinUsdValue)
	s.Equal("10000", resp[0].Collaterals[0].MaxUsdValue)
	s.Equal("0.9000", resp[0].Collaterals[0].DiscountRate)
	s.Equal("10000", resp[0].Collaterals[1].MinUsdValue)
	s.Equal("20000", resp[0].Collaterals[1].MaxUsdValue)
	s.Equal("0.8500", resp[0].Collaterals[1].DiscountRate)
	s.Len(resp[0].AssetNames, 2)
	s.Equal("BTC", resp[0].AssetNames[0].Asset)
	s.Equal("ETH", resp[0].AssetNames[1].Asset)
}

func (s *marginTestSuite) TestMarginCrossMarginFee() {
	data := []byte(`[
		{
			"vipLevel": 0,
			"coin": "BTC",
			"transferIn": true,
			"transferOut": true,
			"dailyInterest": "0.00001000",
			"yearlyInterest": "0.36500000",
			"borrowLimit": "0.00000000",
			"marginablePairs": {
				"pair": "BTCBUSD"
			}
		}
	]`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginCrossMarginFeeService().
		Coin("BTC").
		VipLevel(0).
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 1)
	s.Equal(0, resp[0].VIPLevel)
	s.Equal("BTC", resp[0].Coin)
	s.True(resp[0].TransferIn)
	s.True(resp[0].Borrowable)
	s.Equal("0.00001000", resp[0].DailyInterest)
	s.Equal("0.36500000", resp[0].YearlyInterest)
	s.Equal("0.00000000", resp[0].BorrowLimit)
	s.Equal("BTCBUSD", resp[0].MarginablePairs.Pair)
}

func (s *marginTestSuite) TestMarginCurrentOrderCount() {
	data := []byte(`
	[
		{
			"rateLimitType": "REQUEST_WEIGHT",
			"interval": "MINUTE",
			"intervalNum": 1,
			"limit": 1200,
			"count": 0
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginCurrentOrderCountService().
		Symbol("BTCUSDT").
		IsIsolated("TRUE").
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 1)
	s.Equal("REQUEST_WEIGHT", resp[0].RateLimitType)
	s.Equal("MINUTE", resp[0].Interval)
	s.Equal(1, resp[0].IntervalNum)
	s.Equal(1200, resp[0].Limit)
	s.Equal(0, resp[0].Count)
}

func (s *marginTestSuite) TestMarginInterestRateHistory() {
	data := []byte(`
	[
		{
			"asset": "BTC",
			"vipLevel": 0,
			"timestamp": 1616697600000,
			"dailyInterestRate": "0.00025"
		},
		{
			"asset": "BNB",
			"vipLevel": 0,
			"timestamp": 1616697600000,
			"dailyInterestRate": "0.001"
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginInterestRateHistoryService().
		Asset("BTC").
		VipLevel(0).
		StartTime(1616697600000).
		EndTime(1616784000000).
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 2)
	s.Equal("BTC", resp[0].Asset)
	s.Equal("0.00025", resp[0].DailyInterestRate)
	s.Equal(uint64(1616697600000), resp[0].Timestamp)
	s.Equal(0, resp[0].VIPLevel)

	s.Equal("BNB", resp[1].Asset)
	s.Equal("0.001", resp[1].DailyInterestRate)
	s.Equal(uint64(1616697600000), resp[1].Timestamp)
	s.Equal(0, resp[1].VIPLevel)
}

func (s *marginTestSuite) TestMarginIsolatedAccountDisable() {
	data := []byte(`{"success": true}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginIsolatedAccountDisableService().
		Symbol("BTCUSDT").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(true, resp.Success)
}

func (s *marginTestSuite) TestMarginIsolatedAccountEnable() {
	data := []byte(`{"success": true}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginIsolatedAccountEnableService().
		Symbol("BTCUSDT").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(true, resp.Success)
}

func (s *marginTestSuite) TestMarginIsolatedAccountInfo() {
	data := []byte(`
	{
		"assets": [
			{
				"baseAsset": {
					"asset": "BNB",
					"borrowEnabled": true,
					"free": "0.00000000",
					"interest": "0.00000000",
					"locked": "0.00000000",
					"netAsset": "0.00000000",
					"netAssetOfBtc": "0.00000000",
					"repayEnabled": true,
					"totalAsset": "0.00000000"
				},
				"quoteAsset": {
					"asset": "USDT",
					"borrowEnabled": true,
					"free": "10000.00000000",
					"interest": "0.00000000",
					"locked": "0.00000000",
					"netAsset": "10000.00000000",
					"netAssetOfBtc": "10000.00000000",
					"repayEnabled": true,
					"totalAsset": "10000.00000000"
				},
				"symbol": "BNBUSDT",
				"isolatedCreated": true,
				"enabled": true,
				"marginLevel": "0.00000",
				"marginLevelStatus": "EXCESSIVE",
				"marginRatio": "0.00000",
				"indexPrice": "48.81190000",
				"liquidatePrice": "0.00000000",
				"liquidateRate": "0.00000000",
				"tradeEnabled": true
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginIsolatedAccountInfoService().Symbols("LTCBTC, BTCUSDT").Do(context.Background())
	s.r().NoError(err)
	a := resp.(*MarginIsolatedAccountInfoResponseSymbols)
	s.Len(a.Assets, 1)

	baseAsset := a.Assets[0].BaseAsset
	s.Equal("BNB", baseAsset.Asset)
	s.True(baseAsset.BorrowEnabled)
	s.Equal("0.00000000", baseAsset.Free)
	s.Equal("0.00000000", baseAsset.Interest)
	s.Equal("0.00000000", baseAsset.Locked)
	s.Equal("0.00000000", baseAsset.NetAsset)
	s.Equal("0.00000000", baseAsset.NetAssetOfBtc)
	s.True(baseAsset.RepayEnabled)
	s.Equal("0.00000000", baseAsset.TotalAsset)

	quoteAsset := a.Assets[0].QuoteAsset
	s.Equal("USDT", quoteAsset.Asset)
	s.True(quoteAsset.BorrowEnabled)
	s.Equal("10000.00000000", quoteAsset.Free)
	s.Equal("0.00000000", quoteAsset.Interest)
	s.Equal("0.00000000", quoteAsset.Locked)
	s.Equal("10000.00000000", quoteAsset.NetAsset)
	s.Equal("10000.00000000", quoteAsset.NetAssetOfBtc)
	s.True(quoteAsset.RepayEnabled)
	s.Equal("10000.00000000", quoteAsset.TotalAsset)

	s.Equal("BNBUSDT", a.Assets[0].Symbol)
	s.True(a.Assets[0].IsolatedCreated)
	s.True(a.Assets[0].Enabled)
	s.Equal("0.00000", a.Assets[0].MarginLevel)
	s.Equal("EXCESSIVE", a.Assets[0].MarginLevelStatus)
	s.Equal("0.00000", a.Assets[0].MarginRatio)
	s.Equal("48.81190000", a.Assets[0].IndexPrice)
	s.Equal("0.00000000", a.Assets[0].LiquidatePrice)
	s.Equal("0.00000000", a.Assets[0].LiquidateRate)
	s.True(a.Assets[0].TradeEnabled)
}

func (s *marginTestSuite) TestMarginIsolatedAccountLimit() {
	data := []byte(`
	{
		"enabledAccount": 5,
		"maxAccount": 10
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginIsolatedAccountLimitService().Do(context.Background())

	s.r().NoError(err)
	s.Equal(5, resp.EnabledAccount)
	s.Equal(10, resp.MaxAccount)
}

func (s *marginTestSuite) TestMarginIsolatedMarginFee() {
	data := []byte(`
	[
		{
			"vipLevel": 0,
			"symbol": "BTCUSDT",
			"leverage": "3",
			"data": {
				"coin": "BTC",
				"dailyInterest": "0.0015",
				"borrowLimit": "1.00000000"
			}
		},
		{
			"vipLevel": 1,
			"symbol": "BTCUSDT",
			"leverage": "5",
			"data": {
				"coin": "BTC",
				"dailyInterest": "0.0014",
				"borrowLimit": "2.00000000"
			}
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginIsolatedMarginFeeService().
		Symbol("BTCUSDT").
		VipLevel(0).
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 2)

	s.Equal(0, resp[0].VIPLevel)
	s.Equal("BTCUSDT", resp[0].Symbol)
	s.Equal("3", resp[0].Leverage)
	s.Equal("BTC", resp[0].Data.Coin)
	s.Equal("0.0015", resp[0].Data.DailyInterest)
	s.Equal("1.00000000", resp[0].Data.BorrowLimit)

	s.Equal(1, resp[1].VIPLevel)
	s.Equal("BTCUSDT", resp[1].Symbol)
	s.Equal("5", resp[1].Leverage)
	s.Equal("BTC", resp[1].Data.Coin)
	s.Equal("0.0014", resp[1].Data.DailyInterest)
	s.Equal("2.00000000", resp[1].Data.BorrowLimit)
}

func (s *marginTestSuite) TestMarginIsolatedMarginTier() {
	data := []byte(`
	[
		{
			"symbol": "BTCUSDT",
			"tier": 1,
			"effectiveMultiple": "5",
			"initialRiskRatio": "150",
			"liquidationRiskRatio": "110",
			"baseAssetMaxBorrowable": "0.10000000",
			"quoteAssetMaxBorrowable": "50.00000000"
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginIsolatedMarginTierService().
		Symbol("BTCUSDT").
		Tier(1).
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 1)
	s.Equal("BTCUSDT", resp[0].Symbol)
	s.Equal(1, resp[0].Tier)
	s.Equal("5", resp[0].EffectiveMultiple)
	s.Equal("150", resp[0].InitialRiskRatio)
	s.Equal("110", resp[0].LiquidationRiskRatio)
	s.Equal("0.10000000", resp[0].BaseAssetMaxBorrowable)
	s.Equal("50.00000000", resp[0].QuoteAssetMaxBorrowable)
}

func (s *marginTestSuite) TestMarginSmallLiabilityExchange() {
	data := []byte(`[{}]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginSmallLiabilityExchangeService().
		AssetNames("BTC").
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 1)
}

func (s *marginTestSuite) TestMarginSmallLiabilityExchangeCoinList() {
	data := []byte(`
	[
		{
			"asset": "BTC",
			"interest": "0.00000005",
			"principal": "1.00000000",
			"liabilityOfBUSD": "0.00000000"
		},
		{
			"asset": "ETH",
			"interest": "0.00000000",
			"principal": "0.00000000",
			"liabilityOfBUSD": "0.00000000"
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginSmallLiabilityExchangeCoinListService().Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 2)
	s.Equal("BTC", resp[0].Asset)
	s.Equal("0.00000005", resp[0].Interest)
	s.Equal("1.00000000", resp[0].Principal)
	s.Equal("0.00000000", resp[0].LiabilityOfBUSD)
	s.Equal("ETH", resp[1].Asset)
	s.Equal("0.00000000", resp[1].Interest)
	s.Equal("0.00000000", resp[1].Principal)
	s.Equal("0.00000000", resp[1].LiabilityOfBUSD)
}

func (s *marginTestSuite) TestMarginManualLiquidation() {
	data := []byte(`[
		{
		"asset": "ETH",
		"interest": "0.00083334",
		"principal": "0.001",
		"liabilityAsset": "USDT",
		"liabilityQty": 0.3552
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginManualLiquidationService().MarginType("MARGIN").Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 1)
	s.Equal("ETH", resp[0].Asset)
	s.Equal("0.00083334", resp[0].Interest)
	s.Equal("0.001", resp[0].Principal)
	s.Equal("USDT", resp[0].LiabilityAsset)
	s.Equal(0.3552, resp[0].LiabilityQty)
}

func (s *marginTestSuite) TestMarginAccountNewOTO() {
	data := []byte(`{
		"orderListId": 13551,
		"contingencyType": "OTO",
		"listStatusType": "EXEC_STARTED",
		"listOrderStatus": "EXECUTING",
		"listClientOrderId": "JDuOrsu0Ge8GTyvx8J7VTD",
		"transactionTime": 1725521998054,
		"symbol": "BTCUSDT",
		"isIsolated": false,
		"orders": [
			{
				"symbol": "BTCUSDT",
				"orderId": 29896699,
				"clientOrderId": "y8RB6tQEMuHUXybqbtzTxk"
			},
			{
				"symbol": "BTCUSDT",
				"orderId": 29896700,
				"clientOrderId": "dKQEdh5HhXb7Lpp85jz1dQ"
			}
		],
		"orderReports": [
			{
				"symbol": "BTCUSDT",
				"orderId": 29896699,
				"orderListId": 13551,
				"clientOrderId": "y8RB6tQEMuHUXybqbtzTxk",
				"transactTime": 1725521998054,
				"price": "80000.00000000",
				"origQty": "0.02000000",
				"executedQty": "0",
				"cummulativeQuoteQty": "0",
				"status": "NEW",
				"timeInForce": "GTC",
				"type": "LIMIT",
				"side": "SELL",
				"selfTradePreventionMode": "NONE"
			},
			{
				"symbol": "BTCUSDT",
				"orderId": 29896700,
				"orderListId": 13551,
				"clientOrderId": "dKQEdh5HhXb7Lpp85jz1dQ",
				"transactTime": 1725521998054,
				"price": "50000.00000000",
				"origQty": "0.02000000",
				"executedQty": "0",
				"cummulativeQuoteQty": "0",
				"status": "PENDING_NEW",
				"timeInForce": "GTC",
				"type": "LIMIT",
				"side": "BUY",
				"selfTradePreventionMode": "NONE"
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginAccountNewOTOService().
		Symbol("BTCUSDT").WorkingType("LIMIT").WorkingSide("SELL").WorkingPrice(80000.0).WorkingQuantity(0.02).
		PendingType("LIMIT").PendingSide("BUY").PendingQuantity(0.02).WorkingTimeInForce("GTC").
		PendingPrice(50000.0).PendingTimeInForce("GTC").Do(context.Background())

	s.r().NoError(err)
	s.Equal(int64(13551), resp.OrderListId)
	s.Equal("OTO", resp.ContingencyType)
	s.Equal("EXEC_STARTED", resp.ListStatusType)
	s.Equal("EXECUTING", resp.ListOrderStatus)
	s.Equal("JDuOrsu0Ge8GTyvx8J7VTD", resp.ListClientOrderId)
	s.Equal(int64(1725521998054), resp.TransactionTime)
	s.Equal("BTCUSDT", resp.Symbol)
	s.False(resp.IsIsolated)
	s.Len(resp.Orders, 2)
	s.Equal("BTCUSDT", resp.Orders[0].Symbol)
	s.Equal(int64(29896699), resp.Orders[0].OrderId)
	s.Equal("y8RB6tQEMuHUXybqbtzTxk", resp.Orders[0].ClientOrderId)
	s.Equal("BTCUSDT", resp.Orders[1].Symbol)
	s.Equal(int64(29896700), resp.Orders[1].OrderId)
	s.Equal("dKQEdh5HhXb7Lpp85jz1dQ", resp.Orders[1].ClientOrderId)
	s.Len(resp.OrderReports, 2)
	s.Equal("BTCUSDT", resp.OrderReports[0].Symbol)
	s.Equal(int64(29896699), resp.OrderReports[0].OrderId)
	s.Equal(int64(13551), resp.OrderReports[0].OrderListId)
	s.Equal("y8RB6tQEMuHUXybqbtzTxk", resp.OrderReports[0].ClientOrderId)
	s.Equal(int64(1725521998054), resp.OrderReports[0].TransactTime)
	s.Equal("80000.00000000", resp.OrderReports[0].Price)
	s.Equal("0.02000000", resp.OrderReports[0].OrigQty)
	s.Equal("0", resp.OrderReports[0].ExecutedQty)
	s.Equal("0", resp.OrderReports[0].CummulativeQuoteQty)
	s.Equal("NEW", resp.OrderReports[0].Status)
	s.Equal("GTC", resp.OrderReports[0].TimeInForce)
	s.Equal("LIMIT", resp.OrderReports[0].Type)
	s.Equal("SELL", resp.OrderReports[0].Side)
	s.Equal("NONE", resp.OrderReports[0].SelfTradePreventionMode)
	s.Equal("BTCUSDT", resp.OrderReports[1].Symbol)
	s.Equal(int64(29896700), resp.OrderReports[1].OrderId)
	s.Equal(int64(13551), resp.OrderReports[1].OrderListId)
	s.Equal("dKQEdh5HhXb7Lpp85jz1dQ", resp.OrderReports[1].ClientOrderId)
	s.Equal(int64(1725521998054), resp.OrderReports[1].TransactTime)
	s.Equal("50000.00000000", resp.OrderReports[1].Price)
	s.Equal("0.02000000", resp.OrderReports[1].OrigQty)
	s.Equal("0", resp.OrderReports[1].ExecutedQty)
	s.Equal("0", resp.OrderReports[1].CummulativeQuoteQty)
	s.Equal("PENDING_NEW", resp.OrderReports[1].Status)
	s.Equal("GTC", resp.OrderReports[1].TimeInForce)
	s.Equal("LIMIT", resp.OrderReports[1].Type)
	s.Equal("BUY", resp.OrderReports[1].Side)
	s.Equal("NONE", resp.OrderReports[1].SelfTradePreventionMode)
}

func (s *marginTestSuite) TestMarginAccountNewOTOCO() {
	data := []byte(`{
		"orderListId": 13509,
		"contingencyType": "OTO",
		"listStatusType": "EXEC_STARTED",
		"listOrderStatus": "EXECUTING",
		"listClientOrderId": "u2AUo48LLef5qVenRtwJZy",
		"transactionTime": 1725521881300,
		"symbol": "BNBUSDT",
		"isIsolated": false,
		"orders": [
			{
				"symbol": "BNBUSDT",
				"orderId": 28282534,
				"clientOrderId": "IfYDxvrZI4kiyqYpRH13iI"
			},
			{
				"symbol": "BNBUSDT",
				"orderId": 28282535,
				"clientOrderId": "0HCSsPRxVfW8BkTUy9z4np"
			},
			{
				"symbol": "BNBUSDT",
				"orderId": 28282536,
				"clientOrderId": "dypsgdxWnLY75kwT930cbD"
			}
		],
		"orderReports": [
			{
				"symbol": "BNBUSDT",
				"orderId": 28282534,
				"orderListId": 13509,
				"clientOrderId": "IfYDxvrZI4kiyqYpRH13iI",
				"transactTime": 1725521881300,
				"price": "300.00000000",
				"origQty": "1.00000000",
				"executedQty": "0",
				"cummulativeQuoteQty": "0",
				"status": "NEW",
				"timeInForce": "GTC",
				"type": "LIMIT",
				"side": "BUY",
				"selfTradePreventionMode": "NONE"
			},
			{
				"symbol": "BNBUSDT",
				"orderId": 28282535,
				"orderListId": 13509,
				"clientOrderId": "0HCSsPRxVfW8BkTUy9z4np",
				"transactTime": 1725521881300,
				"price": "0E-8",
				"origQty": "1.00000000",
				"executedQty": "0",
				"cummulativeQuoteQty": "0",
				"status": "PENDING_NEW",
				"timeInForce": "GTC",
				"type": "STOP_LOSS",
				"side": "SELL",
				"stopPrice": "299.00000000",
				"selfTradePreventionMode": "NONE"
			},
			{
				"symbol": "BNBUSDT",
				"orderId": 28282536,
				"orderListId": 13509,
				"clientOrderId": "dypsgdxWnLY75kwT930cbD",
				"transactTime": 1725521881300,
				"price": "301.00000000",
				"origQty": "1.00000000",
				"executedQty": "0",
				"cummulativeQuoteQty": "0",
				"status": "PENDING_NEW",
				"timeInForce": "GTC",
				"type": "LIMIT_MAKER",
				"side": "SELL",
				"selfTradePreventionMode": "NONE"
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginAccountNewOTOCOService().
		Symbol("BNBUSDT").WorkingType("LIMIT").WorkingSide("BUY").WorkingPrice(300).WorkingQuantity(1).
		PendingSide("SELL").PendingQuantity(1).PendingAboveType("LIMIT_MAKER").WorkingTimeInForce("GTC").
		PendingAbovePrice(301.0).PendingBelowType("STOP_LOSS").PendingBelowPrice(299.0).Do(context.Background())

	s.r().NoError(err)
	s.Equal(int64(13509), resp.OrderListId)
	s.Equal("OTO", resp.ContingencyType)
	s.Equal("EXEC_STARTED", resp.ListStatusType)
	s.Equal("EXECUTING", resp.ListOrderStatus)
	s.Equal("u2AUo48LLef5qVenRtwJZy", resp.ListClientOrderId)
	s.Equal(int64(1725521881300), resp.TransactionTime)
	s.Equal("BNBUSDT", resp.Symbol)
	s.False(resp.IsIsolated)
	s.Len(resp.Orders, 3)
	s.Equal("BNBUSDT", resp.Orders[0].Symbol)
	s.Equal(int64(28282534), resp.Orders[0].OrderId)
	s.Equal("IfYDxvrZI4kiyqYpRH13iI", resp.Orders[0].ClientOrderId)
	s.Equal("BNBUSDT", resp.Orders[1].Symbol)
	s.Equal(int64(28282535), resp.Orders[1].OrderId)
	s.Equal("0HCSsPRxVfW8BkTUy9z4np", resp.Orders[1].ClientOrderId)
	s.Equal("BNBUSDT", resp.Orders[2].Symbol)
	s.Equal(int64(28282536), resp.Orders[2].OrderId)
	s.Equal("dypsgdxWnLY75kwT930cbD", resp.Orders[2].ClientOrderId)
	s.Len(resp.OrderReports, 3)
	s.Equal("BNBUSDT", resp.OrderReports[0].Symbol)
	s.Equal(int64(28282534), resp.OrderReports[0].OrderId)
	s.Equal(int64(13509), resp.OrderReports[0].OrderListId)
	s.Equal("IfYDxvrZI4kiyqYpRH13iI", resp.OrderReports[0].ClientOrderId)
	s.Equal(int64(1725521881300), resp.OrderReports[0].TransactTime)
	s.Equal("300.00000000", resp.OrderReports[0].Price)
	s.Equal("1.00000000", resp.OrderReports[0].OrigQty)
	s.Equal("0", resp.OrderReports[0].ExecutedQty)
	s.Equal("0", resp.OrderReports[0].CummulativeQuoteQty)
	s.Equal("NEW", resp.OrderReports[0].Status)
	s.Equal("GTC", resp.OrderReports[0].TimeInForce)
	s.Equal("LIMIT", resp.OrderReports[0].Type)
	s.Equal("BUY", resp.OrderReports[0].Side)
	s.Equal("NONE", resp.OrderReports[0].SelfTradePreventionMode)
	s.Equal("BNBUSDT", resp.OrderReports[1].Symbol)
	s.Equal(int64(28282535), resp.OrderReports[1].OrderId)
	s.Equal(int64(13509), resp.OrderReports[1].OrderListId)
	s.Equal("0HCSsPRxVfW8BkTUy9z4np", resp.OrderReports[1].ClientOrderId)
	s.Equal(int64(1725521881300), resp.OrderReports[1].TransactTime)
	s.Equal("0E-8", resp.OrderReports[1].Price)
	s.Equal("1.00000000", resp.OrderReports[1].OrigQty)
	s.Equal("0", resp.OrderReports[1].ExecutedQty)
	s.Equal("0", resp.OrderReports[1].CummulativeQuoteQty)
	s.Equal("PENDING_NEW", resp.OrderReports[1].Status)
	s.Equal("GTC", resp.OrderReports[1].TimeInForce)
	s.Equal("STOP_LOSS", resp.OrderReports[1].Type)
	s.Equal("SELL", resp.OrderReports[1].Side)
	s.Equal("NONE", resp.OrderReports[1].SelfTradePreventionMode)
	s.Equal("299.00000000", resp.OrderReports[1].StopPrice)
	s.Equal("BNBUSDT", resp.OrderReports[2].Symbol)
	s.Equal(int64(28282536), resp.OrderReports[2].OrderId)
	s.Equal(int64(13509), resp.OrderReports[2].OrderListId)
	s.Equal("dypsgdxWnLY75kwT930cbD", resp.OrderReports[2].ClientOrderId)
	s.Equal(int64(1725521881300), resp.OrderReports[2].TransactTime)
	s.Equal("301.00000000", resp.OrderReports[2].Price)
	s.Equal("1.00000000", resp.OrderReports[2].OrigQty)
	s.Equal("0", resp.OrderReports[2].ExecutedQty)
	s.Equal("0", resp.OrderReports[2].CummulativeQuoteQty)
	s.Equal("PENDING_NEW", resp.OrderReports[2].Status)
	s.Equal("GTC", resp.OrderReports[2].TimeInForce)
	s.Equal("LIMIT_MAKER", resp.OrderReports[2].Type)
	s.Equal("SELL", resp.OrderReports[2].Side)
	s.Equal("NONE", resp.OrderReports[2].SelfTradePreventionMode)
}

func (s *marginTestSuite) TestMarginSmallLiabilityExchangeHistory() {
	data := []byte(`[
	  {
		"total": 1,
		"rows": [
		  {
			"asset": "BTC",
			"amount": "1.00000000",
			"targetAsset": "USDT",
			"targetAmount": "50000.00000000",
			"bizType": "REPAY",
			"timestamp": 1613450271000
		  }
		]
	  }
	]`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginSmallLiabilityExchangeHistoryService().
		Current(1).
		Size(500).
		StartTime(1613450271000).
		EndTime(1613536671000).
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 1)
	s.Equal(1, resp[0].Total)
	s.Len(resp[0].Rows, 1)
	s.Equal("BTC", resp[0].Rows[0].Asset)
	s.Equal("1.00000000", resp[0].Rows[0].Amount)
	s.Equal("USDT", resp[0].Rows[0].TargetAsset)
	s.Equal("50000.00000000", resp[0].Rows[0].TargetAmount)
	s.Equal("REPAY", resp[0].Rows[0].BizType)
	s.Equal(uint64(1613450271000), resp[0].Rows[0].Timestamp)
}

func (s *marginTestSuite) TestMarginToggleBnbBurn() {
	data := []byte(`{"spotBNBBurn":true,"interestBNBBurn":false}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginToggleBnbBurnService().
		SpotBNBBurn("true").
		InterestBNBBurn("false").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(true, resp.SpotBNBBurn)
	s.Equal(false, resp.InterestBNBBurn)
}

func (s *marginTestSuite) TestMarginIsolatedCapitalFlow() {
	data := []byte(`[
		{ 
			"id": 123456,
			"tranId": 123123,
			"timestamp": 1691116657000,
			"asset": "USDT",
			"symbol": "BTCUSDT",
			"type": "BORROW",
			"amount": "101"
		}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginIsolatedCapitalFlowService().Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 1)
	s.Equal(int(123456), resp[0].Id)
	s.Equal(int(123123), resp[0].TranId)
	s.Equal(uint64(1691116657000), resp[0].Timestamp)
	s.Equal("USDT", resp[0].Asset)
	s.Equal("BTCUSDT", resp[0].Symbol)
	s.Equal("BORROW", resp[0].Type)
	s.Equal("101", resp[0].Amount)
}
