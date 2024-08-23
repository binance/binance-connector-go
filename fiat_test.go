package binance_connector

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type fiatTestSuite struct {
	baseTestSuite
}

func TestFiat(t *testing.T) {
	suite.Run(t, new(fiatTestSuite))
}

func (s *fiatTestSuite) TestGetFiatDepositWithdrawHistory() {
	data := []byte(`{
		"code": "000000",
		"message": "success",
		"data": [
			{
				"orderNo":"7d76d611-0568-4f43-afb6-24cac7767365",
				"fiatCurrency": "BRL",
				"indicatedAmount": "10.00",  
				"amount": "10.00", 
				"totalFee": "0.00",
				"method": "BankAccount",
				"status": "Expired",
				"createTime": 1626144956000,
				"updateTime": 1626400907000   
			}
		],
		"total": 1,
		"success": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	transactionType := "0"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"transactionType": transactionType,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetFiatDepositWithdrawHistoryService().TransactionType(transactionType).Do(newContext())
	s.r().NoError(err)
	s.Equal("000000", res.Code)
	s.Equal("success", res.Message)
	s.r().Len(res.Data, 1)
	s.Equal("7d76d611-0568-4f43-afb6-24cac7767365", res.Data[0].OrderNo)
	s.Equal("BRL", res.Data[0].FiatCurrency)
	s.Equal("10.00", res.Data[0].IndicatedAmount)
	s.Equal("10.00", res.Data[0].Amount)
	s.Equal("0.00", res.Data[0].TotalFee)
	s.Equal("BankAccount", res.Data[0].Method)
	s.Equal("Expired", res.Data[0].Status)
	s.Equal(uint64(1626144956000), res.Data[0].CreateTime)
	s.Equal(uint64(1626400907000), res.Data[0].UpdateTime)
	s.Equal(int64(1), res.Total)
	s.True(res.Success)
}

func (s *fiatTestSuite) TestGetPaymentFiatHistory() {
	data := []byte(`{
		"code": "000000",
		"message": "success",
		"data": [
			{
				"orderNo": "353fca443f06466db0c4dc89f94f027a",
				"sourceAmount": "20.0",
				"fiatCurrency": "EUR",
				"obtainAmount": "4.462",
				"cryptoCurrency": "LUNA",
				"totalFee": "0.2",
				"price": "4.437472", 
				"status": "Failed",
				"paymentMethod": "Credit Card",
				"createTime": 1624529919000,
				"updateTime": 1624529919000 
			}
		],
		"total": 1,
		"success": true
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	transactionType := "0"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"transactionType": transactionType,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetFiatPaymentHistoryService().TransactionType(transactionType).Do(newContext())
	s.r().NoError(err)
	s.Equal("000000", res.Code)
	s.Equal("success", res.Message)
	s.r().Len(res.Data, 1)
	s.Equal("353fca443f06466db0c4dc89f94f027a", res.Data[0].OrderNo)
	s.Equal("20.0", res.Data[0].SourceAmount)
	s.Equal("EUR", res.Data[0].FiatCurrency)
	s.Equal("4.462", res.Data[0].ObtainAmount)
	s.Equal("LUNA", res.Data[0].CryptoCurrency)
	s.Equal("0.2", res.Data[0].TotalFee)
	s.Equal("4.437472", res.Data[0].Price)
	s.Equal("Failed", res.Data[0].Status)
	s.Equal("Credit Card", res.Data[0].PaymentMethod)
	s.Equal(uint64(1624529919000), res.Data[0].CreateTime)
	s.Equal(uint64(1624529919000), res.Data[0].UpdateTime)
	s.Equal(int64(1), res.Total)
	s.True(res.Success)
}
