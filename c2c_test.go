package binance_connector

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type c2cTestSuite struct {
	baseTestSuite
}

func TestC2C(t *testing.T) {
	suite.Run(t, new(c2cTestSuite))
}

func (s *c2cTestSuite) TestGetC2CTradeHistory() {
	data := []byte(`
	{
		"code": "000000",
		"message": "success",
		"data": [
			{
				"orderNumber":"20219644646554779648",
				"advNo": "11218246497340923904",
				"tradeType": "SELL",  
				"asset": "BUSD", 
				"fiat": "CNY",  
				"fiatSymbol": "￥",
				"amount": "5000.00000000",
				"totalPrice": "33400.00000000",
				"unitPrice": "6.68",
				"orderStatus": "COMPLETED",
				"createTime": 1619361369000,
				"commission": "0",
				"counterPartNickName": "ab***",
				"advertisementRole": "TAKER"        
			}
		],
		"total": 1,
		"success": true 
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	startTimestamp := uint64(1619361369000)
	endTimestamp := uint64(1619361369000)
	page := 3
	tradeType := "SELL"
	recvWindow := uint64(10000)
	timestamp := uint64(1619361369000)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"startTimestamp": startTimestamp,
			"endTimestamp":   endTimestamp,
			"page":           page,
			"tradeType":      tradeType,
			"recvWindow":     recvWindow,
			"timestamp":      timestamp,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetC2CTradeHistoryService().
		StartTimestamp(startTimestamp).
		EndTimestamp(endTimestamp).
		Page(page).
		TradeType(tradeType).
		RecvWindow(recvWindow).
		Timestamp(timestamp).
		Do(newContext())

	s.r().NoError(err)
	s.Equal("000000", res.Code)
	s.Equal("success", res.Message)
	s.r().Len(res.Data, 1)
	s.Equal("20219644646554779648", res.Data[0].OrderNumber)
	s.Equal("11218246497340923904", res.Data[0].AdvNo)
	s.Equal("SELL", res.Data[0].TradeType)
	s.Equal("BUSD", res.Data[0].Asset)
	s.Equal("CNY", res.Data[0].Fiat)
	s.Equal("￥", res.Data[0].FiatSymbol)
	s.Equal("5000.00000000", res.Data[0].Amount)
	s.Equal("33400.00000000", res.Data[0].TotalPrice)
	s.Equal("6.68", res.Data[0].UnitPrice)
	s.Equal("COMPLETED", res.Data[0].OrderStatus)
	s.Equal(uint64(1619361369000), res.Data[0].CreateTime)
	s.Equal("0", res.Data[0].Commission)
	s.Equal("ab***", res.Data[0].CounterPartNickName)
	s.Equal("TAKER", res.Data[0].AdvertisementRole)
	s.Equal(int64(1), res.Total)
	s.True(res.Success)

}
