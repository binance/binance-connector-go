package binance_connector

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type walletTestSuite struct {
	baseTestSuite
}

func TestWallet(t *testing.T) {
	suite.Run(t, new(walletTestSuite))
}

func (s *walletTestSuite) TestCreateWithdraw() {
	data := []byte(`
	{
		"id":"7213fea8e94b4a5593d507237e5a555b"
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	coin := "USDT"
	withdrawOrderID := "testID"
	network := "ETH"
	address := "myaddress"
	addressTag := "xyz"
	amount := 0.01
	transactionFeeFlag := true
	name := "eth"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"coin":               coin,
			"withdrawOrderId":    withdrawOrderID,
			"network":            network,
			"address":            address,
			"addressTag":         addressTag,
			"amount":             amount,
			"transactionFeeFlag": transactionFeeFlag,
			"name":               name,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewWithdrawService().
		Coin(coin).
		WithdrawOrderId(withdrawOrderID).
		Network(network).
		Address(address).
		AddressTag(addressTag).
		Amount(amount).
		TransactionFeeFlag(transactionFeeFlag).
		Name(name).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal("7213fea8e94b4a5593d507237e5a555b", res.Id)
}

func (s *walletTestSuite) TestDustTransfer() {
	data := []byte(`{
		"totalServiceCharge":"0.02102542",
		"totalTransfered":"1.05127099",
		"transferResult":[
			{
				"amount":"0.03000000",
				"fromAsset":"ETH",
				"operateTime":1563368549307,
				"serviceChargeAmount":"0.00500000",
				"tranId":2970932918,
				"transferedAmount":"0.25000000"
			},
			{
				"amount":"0.09000000",
				"fromAsset":"LTC",
				"operateTime":1563368549404,
				"serviceChargeAmount":"0.01548000",
				"tranId":2970932918,
				"transferedAmount":"0.77400000"
			},
			{
				"amount":"248.61878453",
				"fromAsset":"TRX",
				"operateTime":1563368549489,
				"serviceChargeAmount":"0.00054542",
				"tranId":2970932918,
				"transferedAmount":"0.02727099"
			}
		]
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()
	asset := []string{"ETH", "LTC", "TRX"}
	s.assertReq(func(r *request) {
		e := newSignedRequest()
		for _, a := range asset {
			e.addParam("asset", a)
		}
		s.assertRequestEqual(e, r)
	})
	res, err := s.client.NewDustTransferService().Asset(asset).Do(newContext())
	s.r().NoError(err)
	e := &DustTransferResponse{
		TotalServiceCharge: "0.02102542",
		TotalTransfered:    "1.05127099",
		TransferResult: []*DustTransferResult{
			{
				Amount:              "0.03000000",
				FromAsset:           "ETH",
				OperateTime:         1563368549307,
				ServiceChargeAmount: "0.00500000",
				TranID:              2970932918,
				TransferedAmount:    "0.25000000",
			},
			{
				Amount:              "0.09000000",
				FromAsset:           "LTC",
				OperateTime:         1563368549404,
				ServiceChargeAmount: "0.01548000",
				TranID:              2970932918,
				TransferedAmount:    "0.77400000",
			},
			{
				Amount:              "248.61878453",
				FromAsset:           "TRX",
				OperateTime:         1563368549489,
				ServiceChargeAmount: "0.00054542",
				TranID:              2970932918,
				TransferedAmount:    "0.02727099",
			},
		},
	}
	s.assertTransferResponse(e, res)
}

func (s *walletTestSuite) assertTransferResponse(e, a *DustTransferResponse) {
	r := s.r()
	r.Equal(e.TotalServiceCharge, a.TotalServiceCharge, "TotalServiceCharge")
	r.Equal(e.TotalTransfered, a.TotalTransfered, "TotalTransfered")
	for i, etr := range e.TransferResult {
		r.Equal(etr.Amount, a.TransferResult[i].Amount, "Amount")
		r.Equal(etr.FromAsset, a.TransferResult[i].FromAsset, "FromAsset")
		r.Equal(etr.OperateTime, a.TransferResult[i].OperateTime, "OperateTime")
		r.Equal(etr.ServiceChargeAmount, a.TransferResult[i].ServiceChargeAmount, "ServiceChargeAmount")
		r.Equal(etr.TranID, a.TransferResult[i].TranID, "TranID")
		r.Equal(etr.TransferedAmount, a.TransferResult[i].TransferedAmount, "TransferedAmount")
	}
}

func (s *walletTestSuite) TestUserUniversalTransfer() {
	data := []byte(`
	{
		"tranId":13526853623
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	types := "MAIN_C2C"
	asset := "USDT"
	amount := 0.1
	fromSymbol := "USDT"
	toSymbol := "USDT"

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"type":       types,
			"asset":      asset,
			"amount":     amount,
			"fromSymbol": fromSymbol,
			"toSymbol":   toSymbol,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewUserUniversalTransferService().
		TransferType(types).
		Asset(asset).
		Amount(amount).
		FromSymbol(fromSymbol).
		ToSymbol(toSymbol).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal(int64(13526853623), res.TranId)
}

func (s *walletTestSuite) TestGetDepositAddress() {
	data := []byte(`
	{
		"address": "1HPn8Rx2y6nNSfagQBKy27GB99Vbzg89wv",
		"coin": "BTC",
		"tag": "",
		"url": "https://btc.com/1HPn8Rx2y6nNSfagQBKy27GB99Vbzg89wv"
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	coin := "BTC"
	network := "BTC"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"coin":    coin,
			"network": network,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewDepositAddressService().
		Coin(coin).
		Network(network).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal("1HPn8Rx2y6nNSfagQBKy27GB99Vbzg89wv", res.Address)
	r.Equal("", res.Tag)
	r.Equal("BTC", res.Coin)
	r.Equal("https://btc.com/1HPn8Rx2y6nNSfagQBKy27GB99Vbzg89wv", res.Url)
}

func (s *walletTestSuite) TestDepositAddressListService() {
	data := []byte(`[
		{
			"coin": "ETH",
			"address": "0xD316E95Fd9E8E237Cb11f8200Babbc5D8D177BA4",
			"tag":"",
			"isDefault": 0
		}
	]
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	coin := "BTC"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"coin": coin,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewDepositAddressListService().
		Coin(coin).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	s.Len(res, 1)
	r.Equal("ETH", res[0].Coin)
	r.Equal("0xD316E95Fd9E8E237Cb11f8200Babbc5D8D177BA4", res[0].Address)
	r.Equal("", res[0].Tag)
	r.Equal(int32(0), res[0].IsDefault)
}

func (s *walletTestSuite) TestAccountApiTradingStatus() {
	data := []byte(`
	{
		"data": {
			"isLocked": false,
			"plannedRecoverTime": 0,
			"triggerCondition": {
				"GCR": 90,
				"IFER": 90,
				"UFR": 90
			},
			"updateTime": 1613450271000
		}
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewAccountApiTradingStatusService().Do(context.Background())

	s.r().NoError(err)
	s.False(resp.Data.IsLocked)
	s.Equal(int64(0), resp.Data.PlannedRecoverTime)
	s.Equal(90, resp.Data.TriggerCondition.GCR)
	s.Equal(90, resp.Data.TriggerCondition.IFER)
	s.Equal(90, resp.Data.TriggerCondition.UFR)
	s.Equal(uint64(1613450271000), resp.Data.UpdateTime)
}

func (s *walletTestSuite) TestAccountStatusService() {
	data := []byte(`
	{
		"data": "active"
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewAccountStatusService().Do(context.Background())

	s.r().NoError(err)
	s.Equal("active", resp.Data)
}

func (s *walletTestSuite) TestAPIKeyPermission() {
	data := []byte(`
	{
		"ipRestrict":false, 
		"createTime":1698645219000,
		"enableReading":true, 
		"enableWithdrawals":false,
		"enableInternalTransfer":false,
		"enableMargin":false,
		"enableFutures":false,
		"permitsUniversalTransfer":false,
		"enableVanillaOptions":false,
		"enableFixApiTrade": false,
		"enableFixReadOnly": true,
		"enableSpotAndMarginTrading":false,
		"enablePortfolioMarginTrading":true
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewAPIKeyPermissionService().Do(context.Background())

	s.r().NoError(err)
	s.NotNil(resp)
	s.False(resp.IPRestrict)
	s.Equal(uint64(1698645219000), resp.CreateTime)
	s.True(resp.EnableReading)
	s.False(resp.EnableWithdrawals)
	s.False(resp.EnableInternalTransfer)
	s.False(resp.EnableMargin)
	s.False(resp.EnableFutures)
	s.False(resp.PermitsUniversalTransfer)
	s.False(resp.EnableVanillaOptions)
	s.False(resp.EnableFixApiTrade)
	s.True(resp.EnableFixReadOnly)
	s.False(resp.EnableSpotAndMarginTrading)
	s.True(resp.EnablePortfolioMarginTrading)
}

func (s *walletTestSuite) TestAssetDetail() {
	data := []byte(`
	{
		"details": [
			{
				"asset": "BTC",
				"assetFullName": "Bitcoin",
				"amountFree": "0.00001000",
				"toBTC": "0.00000910",
				"toBNB": "0.00000890",
				"toBNBOffExchange": "0.00000800",
				"exchange": "Binance"
			}
		],
		"totalTransferBtc": "0.00000910",
		"totalTransferBnb": "0.00000890",
		"dribbletPercentage": "0.10000000"
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewAssetDetailService().
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp.Details, 1)
	s.Equal("BTC", resp.Details[0].Asset)
	s.Equal("Bitcoin", resp.Details[0].AssetFullName)
	s.Equal("0.00001000", resp.Details[0].AmountFree)
	s.Equal("0.00000910", resp.Details[0].ToBTC)
	s.Equal("0.00000890", resp.Details[0].ToBNB)
	s.Equal("0.00000800", resp.Details[0].ToBNBOffExchange)
	s.Equal("Binance", resp.Details[0].Exchange)
	s.Equal("0.00000910", resp.TotalTransferBtc)
	s.Equal("0.00000890", resp.TotalTransferBnb)
	s.Equal("0.10000000", resp.DribbletPercentage)
}

func (s *walletTestSuite) TestAssetDetailV2() {
	data := []byte(`
	{
		"assetDetail": {
			"minWithdrawAmount": "0.001",
			"depositStatus": true,
			"withdrawFee": "0.0005",
			"withdrawStatus": true,
			"depositTip": "Please don't deposit any other assets except RUB"
		}
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewAssetDetailV2Service().
		Asset("BTC").
		Do(context.Background())

	s.r().NoError(err)
	s.NotNil(resp.AssetDetail)
	s.Equal("0.001", resp.AssetDetail.MinWithdrawAmount)
	s.True(resp.AssetDetail.DepositStatus)
	s.Equal("0.0005", resp.AssetDetail.WithdrawFee)
	s.True(resp.AssetDetail.WithdrawStatus)
	s.Equal("Please don't deposit any other assets except RUB", resp.AssetDetail.DepositTip)
}

func (s *walletTestSuite) TestWalletBalance() {
	data := []byte(`[
		{
			"activate": true,
			"balance": "0",
			"walletName": "Spot"
		}, 
		{
			"activate": true,
			"balance": "0",
			"walletName": "Funding"
		}, 
		{
			"activate": true,
			"balance": "0",
			"walletName": "Cross Margin"
		}, 
		{
			"activate": true,
			"balance": "0",
			"walletName": "Isolated Margin"
		}, 
		{
			"activate": true,
			"balance": "0.71842752",
			"walletName": "USDⓈ-M Futures"
		}, 
		{
			"activate": true,
			"balance": "0",
			"walletName": "COIN-M Futures"
		}, 
		{
			"activate": true,
			"balance": "0",
			"walletName": "Earn"
		}, 
		{
			"activate": false,
			"balance": "0",
			"walletName": "Options"
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewWalletBalanceService().Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 8)
	s.Equal(true, resp[0].Activate)
	s.Equal("0", resp[0].Balance)
	s.Equal("Spot", resp[0].WalletName)
	s.Equal(true, resp[1].Activate)
	s.Equal("0", resp[1].Balance)
	s.Equal("Funding", resp[1].WalletName)
	s.Equal(true, resp[2].Activate)
	s.Equal("0", resp[2].Balance)
	s.Equal("Cross Margin", resp[2].WalletName)
	s.Equal(true, resp[3].Activate)
	s.Equal("0", resp[3].Balance)
	s.Equal("Isolated Margin", resp[3].WalletName)
	s.Equal(true, resp[4].Activate)
	s.Equal("0.71842752", resp[4].Balance)
	s.Equal("USDⓈ-M Futures", resp[4].WalletName)
	s.Equal(true, resp[5].Activate)
	s.Equal("0", resp[5].Balance)
	s.Equal("COIN-M Futures", resp[5].WalletName)
	s.Equal(true, resp[6].Activate)
	s.Equal("0", resp[6].Balance)
	s.Equal("Earn", resp[6].WalletName)
	s.Equal(false, resp[7].Activate)
	s.Equal("0", resp[7].Balance)
	s.Equal("Options", resp[7].WalletName)
}

func (s *walletTestSuite) TestAssetDividendRecord() {
	data := []byte(`
	{
		"rows": [
			{
				"id": 123,
				"amount": "1.00000000",
				"asset": "BTC",
				"divTime": 1613450271000,
				"enInfo": "BTC distribution",
				"tranId": 456
			}
		],
		"total": 1
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewAssetDividendRecordService().
		Asset("BTC").
		StartTime(1613450271000).
		EndTime(1613536671000).
		Limit(500).
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp.Rows, 1)
	s.Equal(int64(123), resp.Rows[0].Id)
	s.Equal("1.00000000", resp.Rows[0].Amount)
	s.Equal("BTC", resp.Rows[0].Asset)
	s.Equal(uint64(1613450271000), resp.Rows[0].DivTime)
	s.Equal("BTC distribution", resp.Rows[0].EnInfo)
	s.Equal(int64(456), resp.Rows[0].TranId)
	s.Equal(int64(1), resp.Total)
}

func (s *walletTestSuite) TestBUSDConvertHistory() {
	data := []byte(`
	{
		"total": 1,
		"rows": [
			{
				"tranId":118263615991,
				"type":244,
				"time":1664442078000,
				"deductedAsset":"BUSD",
				"deductedAmount":"1",
				"targetAsset":"USDC",
				"targetAmount":"1",
				"status":"S",
				"accountType":"MAIN"
			}
		]
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewBUSDConvertHistoryService().
		Asset("BUSD").
		StartTime(1664442078000).
		EndTime(1613536671000).
		Current(1).
		Size(500).
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp.Rows, 1)
	s.Equal(int64(118263615991), resp.Rows[0].TranId)
	s.Equal(int32(244), resp.Rows[0].Type)
	s.Equal(uint64(1664442078000), resp.Rows[0].Time)
	s.Equal("BUSD", resp.Rows[0].DeductedAsset)
	s.Equal("1", resp.Rows[0].DeductedAmount)
	s.Equal("S", resp.Rows[0].Status)
	s.Equal("MAIN", resp.Rows[0].AccountType)
	s.Equal(int32(1), resp.Total)
}

func (s *walletTestSuite) TestCloudMiningPaymentHistory() {
	data := []byte(`
	{
		"rows": [
			{
				"createTime": 1613450271000,
				"tranId": 123,
				"type": 0,
				"asset": "BTC",
				"amount": "1.00000000",
				"status": "CONFIRMED"
			}
		],
		"total": 1
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewCloudMiningPaymentHistoryService().
		Asset("BTC").
		StartTime(1613450271000).
		EndTime(1613536671000).
		Current(1).
		Size(500).
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp.Rows, 1)
	s.Equal(uint64(1613450271000), resp.Rows[0].CreateTime)
	s.Equal(int64(123), resp.Rows[0].TranId)
	s.Equal(int32(0), resp.Rows[0].Type)
	s.Equal("BTC", resp.Rows[0].Asset)
	s.Equal("1.00000000", resp.Rows[0].Amount)
	s.Equal("CONFIRMED", resp.Rows[0].Status)
	s.Equal(int32(1), resp.Total)
}

func (s *walletTestSuite) TestUserDelegationHistory() {
	data := []byte(`
	{
		"total": 3316,
		"rows": [
			{
				"clientTranId": "293915932290879488",
				"transferType": "Undelegate",
				"asset": "ETH",
				"amount": "1",
				"time": 1695205406000
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewUserDelegationHistoryService().
		Email("email@email.com").
		StartTime(1695200406000).
		EndTime(1695210406000).
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp.Rows, 1)
	s.Equal("293915932290879488", resp.Rows[0].ClientTranId)
	s.Equal("Undelegate", resp.Rows[0].TransferType)
	s.Equal("ETH", resp.Rows[0].Asset)
	s.Equal("1", resp.Rows[0].Amount)
	s.Equal(uint64(1695205406000), resp.Rows[0].Time)
	s.Equal(int32(3316), resp.Total)
}

func (s *walletTestSuite) TestAccountInfo() {
	data := []byte(`
	{
		"vipLevel": 0,
		"isMarginEnabled": true,
		"isFutureEnabled": true,
		"isOptionsEnabled":true,
		"isPortfolioMarginRetailEnabled":true
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewAccountInfoService().Do(context.Background())
	s.r().NoError(err)
	s.Equal(0, resp.VipLevel)
	s.True(resp.IsMarginEnabled)
	s.True(resp.IsFutureEnabled)
	s.True(resp.IsOptionsEnabled)
	s.True(resp.IsPortfolioMarginRetailEnabled)
}

func (s *walletTestSuite) TestDepositHistory() {
	data := []byte(`
	[
		{
			"id": "769800519366885376",
			"amount": "0.001",
			"coin": "BNB",
			"network": "BNB",
			"status": 0,
			"address": "bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23",
			"addressTag": "101764890",
			"txId": "98A3EA560C6B3336D348B6C83F0F95ECE4F1F5919E94BD006E5BF3BF264FACFC",
			"insertTime": 1661493146000,
			"transferType": 0,
			"confirmTimes": "1/1",
			"unlockConfirm": 0,
			"walletType": 0
		}
	]`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewDepositHistoryService().
		Coin("BTC").
		Status(1).
		StartTime(1609459200000).
		EndTime(1609545600000).
		Offset(0).
		Limit(500).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("769800519366885376", resp[0].Id)
	s.Equal("0.001", resp[0].Amount)
	s.Equal("BNB", resp[0].Coin)
	s.Equal("BNB", resp[0].Network)
	s.Equal(0, resp[0].Status)
	s.Equal("bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23", resp[0].Address)
	s.Equal("101764890", resp[0].AddressTag)
	s.Equal("98A3EA560C6B3336D348B6C83F0F95ECE4F1F5919E94BD006E5BF3BF264FACFC", resp[0].TxId)
	s.Equal(uint64(1661493146000), resp[0].InsertTime)
	s.Equal(0, resp[0].TransferType)
	s.Equal("1/1", resp[0].ConfirmTimes)
	s.Equal(0, resp[0].UnlockConfirm)
	s.Equal(0, resp[0].WalletType)
}

func (s *walletTestSuite) TestDustLog() {
	data := []byte(`
	{
		"total": 1,
		"userAssetDribblets": [
			{
				"operateTime": 1613450271000,
				"totalTransferedAmount": "0.00100000",
				"totalServiceChargeAmount": "0.00000020",
				"transId": 123,
				"userAssetDribbletDetails": [
					{
						"transId": 123,
						"serviceChargeAmount": "0.00000010",
						"amount": "0.00050000",
						"operateTime": 1613450271000,
						"transferedAmount": "0.00049000",
						"fromAsset": "BTC"
					},
					{
						"transId": 123,
						"serviceChargeAmount": "0.00000010",
						"amount": "0.00050000",
						"operateTime": 1613450271000,
						"transferedAmount": "0.00049000",
						"fromAsset": "ETH"
					}
				]
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewDustLogService().
		StartTime(1613450271000).
		EndTime(1613536671000).
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp.UserAssetDribblets, 1)
	s.Equal(1613450271000, int(resp.UserAssetDribblets[0].OperateTime))
	s.Equal("0.00100000", resp.UserAssetDribblets[0].TotalTransferedAmount)
	s.Equal("0.00000020", resp.UserAssetDribblets[0].TotalServiceChargeAmount)
	s.Equal(int64(123), resp.UserAssetDribblets[0].TransId)
	s.Len(resp.UserAssetDribblets[0].UserAssetDribbletDetails, 2)
	s.Equal(int64(123), resp.UserAssetDribblets[0].UserAssetDribbletDetails[0].TransId)
	s.Equal("0.00000010", resp.UserAssetDribblets[0].UserAssetDribbletDetails[0].ServiceChargeAmount)
	s.Equal("0.00050000", resp.UserAssetDribblets[0].UserAssetDribbletDetails[0].Amount)
	s.Equal(uint64(1613450271000), uint64(resp.UserAssetDribblets[0].UserAssetDribbletDetails[0].OperateTime))
	s.Equal("0.00049000", resp.UserAssetDribblets[0].UserAssetDribbletDetails[0].TransferedAmount)
	s.Equal("BTC", resp.UserAssetDribblets[0].UserAssetDribbletDetails[0].FromAsset)
	s.Equal(int64(123), resp.UserAssetDribblets[0].UserAssetDribbletDetails[1].TransId)
	s.Equal("0.00000010", resp.UserAssetDribblets[0].UserAssetDribbletDetails[1].ServiceChargeAmount)
	s.Equal("0.00050000", resp.UserAssetDribblets[0].UserAssetDribbletDetails[1].Amount)
	s.Equal(1613450271000, int(resp.UserAssetDribblets[0].UserAssetDribbletDetails[1].OperateTime))
	s.Equal("0.00049000", resp.UserAssetDribblets[0].UserAssetDribbletDetails[1].TransferedAmount)
	s.Equal("ETH", resp.UserAssetDribblets[0].UserAssetDribbletDetails[1].FromAsset)
}

func (s *walletTestSuite) TestFundingWallet() {
	data := []byte(`
	[
		{
			"asset": "BTC",
			"free": "0.1",
			"locked": "0.2",
			"freeze": "0.3",
			"withdrawing": "0.4",
			"btcValuation": "0.5"
		}
	]`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewFundingWalletService().
		Asset("BTC").
		NeedBtcValuation("true").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("BTC", resp[0].Asset)
	s.Equal("0.1", resp[0].Free)
	s.Equal("0.2", resp[0].Locked)
	s.Equal("0.3", resp[0].Freeze)
	s.Equal("0.4", resp[0].Withdrawing)
	s.Equal("0.5", resp[0].BtcValuation)
}

func (s *walletTestSuite) TestGetAccountSnapshot() {
	data := []byte(`
	{
		"code":200,
		"msg":"success",
		"snapshotVos":[
			{
				"type":"SPOT",
				"updateTime":1557712656239,
				"data":{
					"balances":[
						{
							"asset":"BTC",
							"free":"0.00219821",
							"locked":"0.00000000"
						},
						{
							"asset":"LTC",
							"free":"0.00000000",
							"locked":"0.00000000"
						}
					],
					"totalAssetOfBtc":"0.00167717"
				}
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetAccountSnapshotService().
		MarketType("SPOT").
		StartTime(1557712656239).
		Limit(500).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(200, resp.Code)
	s.Equal("success", resp.Msg)
	s.Len(resp.SnapshotVos, 1)
	s.Equal("SPOT", resp.SnapshotVos[0].Type)
	s.Equal(uint64(1557712656239), resp.SnapshotVos[0].UpdateTime)
	s.Len(resp.SnapshotVos[0].Data.Balances, 2)
	s.Equal("BTC", resp.SnapshotVos[0].Data.Balances[0].Asset)
	s.Equal("0.00219821", resp.SnapshotVos[0].Data.Balances[0].Free)
	s.Equal("0.00000000", resp.SnapshotVos[0].Data.Balances[0].Locked)
	s.Equal("LTC", resp.SnapshotVos[0].Data.Balances[1].Asset)
	s.Equal("0.00000000", resp.SnapshotVos[0].Data.Balances[1].Free)
	s.Equal("0.00000000", resp.SnapshotVos[0].Data.Balances[1].Locked)
	s.Equal("0.00167717", resp.SnapshotVos[0].Data.TotalAssetOfBtc)
}

func (s *walletTestSuite) TestGetAllCoinsInfoService() {
	data := []byte(`
	[
		{
			"coin": "BTC",
			"depositAllEnable": true,
			"free": "0.00000000",
			"freeze": "0.00000000",
			"ipoable": "0.00000000",
			"ipoing": "0.00000000",
			"isLegalMoney": false,
			"locked": "0.00000000",
			"name": "Bitcoin",
			"networkList": [
				{
					"addressRegex": "^(bnb1)[0-9a-z]{38}$",
					"coin": "BTC",
					"depositDesc": "Deposit to address",
					"depositEnable": true,
					"isDefault": false,
					"memoRegex": "",
					"minConfirm": 1,
					"name": "BEP2",
					"network": "BNB",
					"specialTips": "",
					"unLockConfirm": 0,
					"withdrawDesc": "Withdrawal to address",
					"withdrawEnable": true,
					"withdrawFee": "0.00000120",
					"withdrawIntegerMultiple": "0.00000001",
					"withdrawMax": "0.00000000",
					"withdrawMin": "0.00200000",
					"sameAddress": false,
					"estimatedArrivalTime": 0,
					"busy": false,
					"contractAddressUrl": "",
					"contractAddress": ""
				}
			],
			"storage": "0.00000000",
			"trading": true,
			"withdrawAllEnable": true,
			"withdrawing": "0.00000000"
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetAllCoinsInfoService().Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 1)
	s.Equal("BTC", resp[0].Coin)
	s.True(resp[0].DepositAllEnable)
	s.Equal("0.00000000", resp[0].Free)
	s.Equal("0.00000000", resp[0].Freeze)
	s.Equal("0.00000000", resp[0].Ipoable)
	s.Equal("0.00000000", resp[0].Ipoing)
	s.False(resp[0].IsLegalMoney)
	s.Equal("0.00000000", resp[0].Locked)
	s.Equal("Bitcoin", resp[0].Name)
	s.Len(resp[0].NetworkList, 1)
	s.Equal("^(bnb1)[0-9a-z]{38}$", resp[0].NetworkList[0].AddressRegex)
	s.Equal("BTC", resp[0].NetworkList[0].Coin)
	s.Equal("Deposit to address", resp[0].NetworkList[0].DepositDesc)
	s.True(resp[0].NetworkList[0].DepositEnable)
	s.False(resp[0].NetworkList[0].IsDefault)
	s.Equal("", resp[0].NetworkList[0].MemoRegex)
	s.Equal(1, resp[0].NetworkList[0].MinConfirm)
	s.Equal("BEP2", resp[0].NetworkList[0].Name)
	s.Equal("BNB", resp[0].NetworkList[0].Network)
	s.Equal("", resp[0].NetworkList[0].SpecialTips)
	s.Equal(0, resp[0].NetworkList[0].UnLockConfirm)
	s.Equal("Withdrawal to address", resp[0].NetworkList[0].WithdrawDesc)
	s.True(resp[0].NetworkList[0].WithdrawEnable)
	s.Equal("0.00000120", resp[0].NetworkList[0].WithdrawFee)
	s.Equal("0.00000001", resp[0].NetworkList[0].WithdrawIntegerMultiple)
	s.Equal("0.00000000", resp[0].NetworkList[0].WithdrawMax)
	s.Equal("0.00200000", resp[0].NetworkList[0].WithdrawMin)
	s.False(resp[0].NetworkList[0].SameAddress)
	s.Equal(uint64(0), resp[0].NetworkList[0].EstimatedArrivalTime)
	s.False(resp[0].NetworkList[0].Busy)
	s.Equal("", resp[0].NetworkList[0].ContractAddressUrl)
	s.Equal("", resp[0].NetworkList[0].ContractAddress)
	s.Equal("0.00000000", resp[0].Storage)
	s.True(resp[0].Trading)
	s.True(resp[0].WithdrawAllEnable)
	s.Equal("0.00000000", resp[0].Withdrawing)
}

func (s *walletTestSuite) TestGetSymbolsDelistSchedule() {
	data := []byte(`[
		{
			"delistTime": 1686161202000,
			"symbols": [
				"ADAUSDT",
				"BNBUSDT"
			]
		},
		{
			"delistTime": 1686222232000,
			"symbols": [
				"ETHUSDT"
			]
		}
	]
  `)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetSymbolsDelistScheduleService().Do(context.Background())

	s.r().NoError(err)
	s.Len(resp, 2)
	s.Equal(uint64(1686161202000), resp[0].DelistTime)
	s.Equal("ADAUSDT", resp[0].Symbols[0])
	s.Equal("BNBUSDT", resp[0].Symbols[1])
	s.Equal(uint64(1686222232000), resp[1].DelistTime)
	s.Equal("ETHUSDT", resp[1].Symbols[0])
}

func (s *walletTestSuite) TestGetSystemStatus() {
	data := []byte(`
	{
		"status": 0,
		"msg": "normal"
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetSystemStatusService().Do(context.Background())

	s.r().NoError(err)
	s.Equal(0, resp.Status)
	s.Equal("normal", resp.Msg)
}

func (s *walletTestSuite) TestTradeFee() {
	data := []byte(`
	[
		{
			"symbol": "BTCUSDT",
			"makerCommission": "0.00050000",
			"takerCommission": "0.00050000"
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewTradeFeeService().
		Symbol("BTCUSDT").
		Do(context.Background())

	s.r().NoError(err)
	s.NotNil(resp)
	s.Equal("BTCUSDT", resp[0].Symbol)
	s.Equal("0.00050000", resp[0].MakerCommission)
	s.Equal("0.00050000", resp[0].TakerCommission)
}

func (s *walletTestSuite) TestUserAsset() {
	data := []byte(`
	[
		{
			"asset": "BTC",
			"free": "1.00000000",
			"locked": "0.00000000",
			"freeze": "0.00000000",
			"withdrawing": "0.00000000",
			"ipoable": "0.00000000",
			"btcValuation": "1000000.00000000"
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewUserAssetService().
		Asset("BTC").
		NeedBtcValuation(true).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("BTC", resp[0].Asset)
	s.Equal("1.00000000", resp[0].Free)
	s.Equal("0.00000000", resp[0].Locked)
	s.Equal("0.00000000", resp[0].Freeze)
	s.Equal("0.00000000", resp[0].Withdrawing)
	s.Equal("0.00000000", resp[0].Ipoable)
	s.Equal("1000000.00000000", resp[0].BtcValuation)
}

func (s *walletTestSuite) TestUserUniversalTransferHistory() {
	data := []byte(`
	{
		"rows": [
			{
				"asset": "BTC",
				"amount": "1.00000000",
				"type": "MAIN_UMFUTURE",
				"status": "CONFIRMED",
				"tranId": 123,
				"timestamp": 1613450271000
			}
		],
		"total": 1
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewUserUniversalTransferHistoryService().
		TransferType("MAIN_UMFUTURE").
		StartTime(1613450271000).
		EndTime(1613536671000).
		Current(1).
		Size(500).
		FromSymbol("BTC").
		ToSymbol("USDT").
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp.Rows, 1)
	s.Equal("BTC", resp.Rows[0].Asset)
	s.Equal("1.00000000", resp.Rows[0].Amount)
	s.Equal("MAIN_UMFUTURE", resp.Rows[0].Type)
	s.Equal("CONFIRMED", resp.Rows[0].Status)
	s.Equal(int64(123), resp.Rows[0].TranId)
	s.Equal(uint64(1613450271000), resp.Rows[0].Timestamp)
	s.Equal(int64(1), resp.Total)
}

func (s *walletTestSuite) TestWithdrawHistory() {
	data := []byte(`
	[
		{
			"id": "btc123",
			"amount": "1.00000000",
			"transactionFee": "0.00010000",
			"coin": "BTC",
			"status": 6,
			"address": "abc123",
			"txId": "def456",
			"applyTime": "1617233588000",
			"network": "btc",
			"transferType": 0,
			"withdrawOrderId": "ghi789",
			"info": "",
			"confirmNo": 2,
			"walletType": 1,
			"txKey": "jkl012"
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewWithdrawHistoryService().
		Coin("BTC").
		WithdrawOrderId("ghi789").
		Status(6).
		Offset(0).
		Limit(10).
		StartTime(1617233588000).
		EndTime(1617320000000).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("btc123", resp[0].Id)
	s.Equal("1.00000000", resp[0].Amount)
	s.Equal("0.00010000", resp[0].TransactionFee)
	s.Equal("BTC", resp[0].Coin)
	s.Equal(6, resp[0].Status)
	s.Equal("abc123", resp[0].Address)
	s.Equal("def456", resp[0].TxId)
	s.Equal("1617233588000", resp[0].ApplyTime)
	s.Equal("btc", resp[0].Network)
	s.Equal(0, resp[0].TransferType)
	s.Equal("ghi789", resp[0].WithdrawOrderId)
	s.Equal("", resp[0].Info)
	s.Equal(2, resp[0].ConfirmNo)
	s.Equal(1, resp[0].WalletType)
	s.Equal("jkl012", resp[0].TxKey)
}
