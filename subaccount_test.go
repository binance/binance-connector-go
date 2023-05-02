package binance_connector

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type subAccountTestSuite struct {
	baseTestSuite
}

func TestSubAccount(t *testing.T) {
	suite.Run(t, new(subAccountTestSuite))
}

func (s *subAccountTestSuite) TestSubaccountDepositAddressService() {
	data := []byte(`
	{
		"address":"TDunhSa7jkTNuKrusUTU1MUHtqXoBPKETV",
		"coin":"USDT",
		"tag":"a_tag",
		"url":"https://tronscan.org/#/address/TDunhSa7jkTNuKrusUTU1MUHtqXoBPKETV"
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	email := "testsub@gmail.com"
	coin := "a_coin"
	network := "a_network"

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"email":   email,
			"coin":    coin,
			"network": network,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetSubAccountDepositAddressService().
		Email(email).
		Coin(coin).
		Network(network).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal("TDunhSa7jkTNuKrusUTU1MUHtqXoBPKETV", res.Address, "Address")
	r.Equal("USDT", res.Coin, "Coin")
	r.Equal("a_tag", res.Tag, "Tag")
	r.Equal("https://tronscan.org/#/address/TDunhSa7jkTNuKrusUTU1MUHtqXoBPKETV", res.Url, "URL")
}

func (s *subAccountTestSuite) TestInternalUniversalTransferHistory() {
	data := []byte(`
	{
		"result": [
			{
				"tranId": 92275823339,
				"fromEmail": "sub1@gmail.com",
				"toEmail": "sub2@gmail.com",
				"asset": "USDT",
				"amount": "100.0",
				"createTimeStamp": 1640317374000,
				"fromAccountType": "USDT_FUTURE",
				"toAccountType": "SPOT",
				"status": "SUCCESS",
				"clientTranId": "testID"
			}
		],
		"totalCount": 1
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	fromEmail := "sub1@gmail.com"
	toEmail := "sub2@gmail.com"
	clientTranId := "testID"
	endTime := time.Now().UnixNano() / 1000 / 1000
	startTime := endTime - 3600*1000
	page := 1
	limit := 10
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"fromEmail":    fromEmail,
			"toEmail":      toEmail,
			"startTime":    startTime,
			"endTime":      endTime,
			"clientTranId": clientTranId,
			"page":         1,
			"limit":        10,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewQueryUniversalTransferHistoryService().
		FromEmail(fromEmail).
		ToEmail(toEmail).
		StartTime(uint64(startTime)).
		EndTime(uint64(endTime)).
		Page(page).
		Limit(limit).
		ClientTranId(clientTranId).
		Do(newContext())

	r := s.r()
	r.NoError(err)

	s.assertInternalUniversalTransferEqual(&InternalUniversalTransfer{
		FromEmail:       fromEmail,
		ToEmail:         toEmail,
		FromAccountType: "USDT_FUTURE",
		ToAccountType:   "SPOT",
		TranId:          92275823339,
		ClientTranId:    clientTranId,
		Asset:           "USDT",
		Amount:          "100.0",
		Status:          "SUCCESS",
		CreateTimeStamp: 1640317374000,
	}, res.Result[0])
	r.Equal(1, res.TotalCount, "TotalCount")
}

func (s *subAccountTestSuite) assertInternalUniversalTransferEqual(e, a *InternalUniversalTransfer) {
	r := s.r()
	r.Equal(e.FromEmail, a.FromEmail, "FromEmail")
	r.Equal(e.ToEmail, a.ToEmail, "ToEmail")
	r.Equal(e.FromAccountType, a.FromAccountType, "FromAccountType")
	r.Equal(e.ToAccountType, a.ToAccountType, "ToAccountType")
	r.Equal(e.TranId, a.TranId, "TranId")
	r.Equal(e.ClientTranId, a.ClientTranId, "ClientTranId")
	r.Equal(e.Asset, a.Asset, "Asset")
	r.Equal(e.Amount, a.Amount, "Amount")
	r.Equal(e.Status, a.Status, "Status")
	r.Equal(e.CreateTimeStamp, a.CreateTimeStamp, "CreateTimeStamp")
}

func (s *subAccountTestSuite) TestCreateSubAccount() {
	expectedEmail := "test@example.com"
	data := []byte(fmt.Sprintf(`{"email": "%s"}`, expectedEmail))

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewCreateSubAccountService().
		SubAccountString("Test Subaccount").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(expectedEmail, resp.Email)
}

func (s *subAccountTestSuite) TestDeleteIPListForSubAccountAPIKey() {
	ipAddress := "192.168.0.1"
	respData := []byte(`
	{
		"ipRestrict": "NONE",
		"ipList": [
			{
				"ip": "192.168.0.2"
			},
			{
				"ip": "192.168.0.3"
			}
		],
		"updateTime": 1617205250105,
		"apiKey": "a1b2c3d4"
	}
	`)
	s.mockDo(respData, nil)
	defer s.assertDo()

	resp, err := s.client.NewDeleteIPListForSubAccountAPIKeyService().
		Email("subaccount@example.com").
		SubAccountApiKey("a1b2c3d4").
		IpAddress(ipAddress).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("NONE", resp.IpRestrict)
	s.Len(resp.IpList, 2)
	s.Equal("192.168.0.2", resp.IpList[0].Ip)
	s.Equal("192.168.0.3", resp.IpList[1].Ip)
	s.Equal(uint64(1617205250105), resp.UpdateTime)
	s.Equal("a1b2c3d4", resp.ApiKey)
}

func (s *subAccountTestSuite) TestDepositAssetsIntoTheManagedSubAccount() {
	data := []byte(`{"tranId":123}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	toEmail := "test@example.com"
	asset := "BTC"
	amount := 1.0

	resp, err := s.client.NewDepositAssetsIntoManagedSubAccountService().
		ToEmail(toEmail).
		Asset(asset).
		Amount(amount).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(int64(123), resp.TranId)
}

func (s *subAccountTestSuite) TestEnableFuturesForSubAccount() {
	data := []byte(`{"email":"test@test.com", "isFuturesEnabled": true}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewEnableFuturesForSubAccountService().
		Email("test@test.com").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("test@test.com", resp.Email)
	s.True(resp.IsFuturesEnabled)
}

func (s *marginTestSuite) TestEnableLeverageTokenForSubAccount() {
	data := []byte(`{
		"email": "example@example.com",
		"enableBlvt": true
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewEnableLeverageTokenForSubAccountService().
		Email("example@example.com").
		EnableBlvt(true).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("example@example.com", resp.Email)
	s.True(resp.EnableBlvt)
}

func (s *marginTestSuite) TestEnableMarginForSubAccount() {
	email := "test@example.com"
	expectedResp := &EnableMarginForSubAccountResp{
		Email:           email,
		IsMarginEnabled: true,
	}

	data, _ := json.Marshal(expectedResp)
	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewEnableMarginForSubAccountService().
		Email(email).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(expectedResp.Email, resp.Email)
	s.Equal(expectedResp.IsMarginEnabled, resp.IsMarginEnabled)
}

func (s *subAccountTestSuite) TestFuturesTransferForSubAccount() {
	data := []byte(`{"txnId":12345}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	email := "example@binance.com"
	asset := "BTC"
	amount := 1.0
	transferType := 2

	resp, err := s.client.NewFuturesTransferForSubAccountService().
		Email(email).
		Asset(asset).
		Amount(amount).
		TransferType(transferType).
		Do(context.Background())

	s.r().NoError(err)
	s.NotNil(resp)
	s.Equal(12345, resp.TxnId)
}

func (s *subAccountTestSuite) TestGetDetailOnSubAccountFuturesAccount() {
	data := []byte(`{
		"email": "user@example.com",
		"asset": "BNB",
		"assets": [
		  {
			"asset": "BTC",
			"initialMargin": "0.001",
			"maintenanceMargin": "0.001",
			"marginBalance": "0.001",
			"maxWithdrawAmount": "0.001",
			"openOrderInitialMargin": "0.001",
			"positionInitialMargin": "0.001",
			"unrealizedProfit": "0.001",
			"walletBalance": "0.001"
		  }
		],
		"canDeposit": true,
		"canTrade": true,
		"canWithdraw": true,
		"feeTier": 1,
		"maxWithdrawAmount": "0.001",
		"totalInitialMargin": "0.001",
		"totalMaintenanceMargin": "0.001",
		"totalMarginBalance": "0.001",
		"totalOpenOrderInitialMargin": "0.001",
		"totalPositionInitialMargin": "0.001",
		"totalUnrealizedProfit": "0.001",
		"totalWalletBalance": "0.001",
		"updateTime": 1613450271000
	  }`)
	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetDetailOnSubAccountFuturesAccountService().
		Email("user@example.com").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("user@example.com", resp.Email)
	s.Equal("BNB", resp.Asset)
	s.Len(resp.Assets, 1)
	s.Equal("BTC", resp.Assets[0].Asset)
	s.Equal("0.001", resp.Assets[0].InitialMargin)
	s.Equal("0.001", resp.Assets[0].MaintenanceMargin)
	s.Equal("0.001", resp.Assets[0].MarginBalance)
	s.Equal("0.001", resp.Assets[0].MaxWithdrawAmount)
	s.Equal("0.001", resp.Assets[0].OpenOrderInitialMargin)
	s.Equal("0.001", resp.Assets[0].PositionInitialMargin)
	s.Equal("0.001", resp.Assets[0].UnrealizedProfit)
	s.Equal("0.001", resp.Assets[0].WalletBalance)
	s.True(resp.CanDeposit)
	s.True(resp.CanTrade)
	s.True(resp.CanWithdraw)
	s.Equal(1, resp.FeeTier)
	s.Equal("0.001", resp.MaxWithdrawAmount)
	s.Equal("0.001", resp.TotalInitialMargin)
	s.Equal("0.001", resp.TotalMaintenanceMargin)
	s.Equal("0.001", resp.TotalMarginBalance)
	s.Equal("0.001", resp.TotalOpenOrderInitialMargin)
	s.Equal("0.001", resp.TotalPositionInitialMargin)
	s.Equal("0.001", resp.TotalUnrealizedProfit)
	s.Equal("0.001", resp.TotalWalletBalance)
	s.Equal(uint64(1613450271000), resp.UpdateTime)
}

func (s *subAccountTestSuite) TestGetDetailOnSubAccountFuturesAccountV2() {
	data := []byte(`
	{
		"futureAccountResp": {
		"email": "abc@test.com",
		"assets":[
			{
				"asset": "USDT",
				"initialMargin": "0.00000000",
				"maintenanceMargin": "0.00000000",
				"marginBalance": "0.88308000",
				"maxWithdrawAmount": "0.88308000",
				"openOrderInitialMargin": "0.00000000",
				"positionInitialMargin": "0.00000000",
				"unrealizedProfit": "0.00000000",
				"walletBalance": "0.88308000"
			 }
		],
		"canDeposit": true,
		"canTrade": true,
		"canWithdraw": true,
		"feeTier": 2,
		"maxWithdrawAmount": "0.88308000",
		"totalInitialMargin": "0.00000000",
		"totalMaintenanceMargin": "0.00000000",
		"totalMarginBalance": "0.88308000",
		"totalOpenOrderInitialMargin": "0.00000000",
		"totalPositionInitialMargin": "0.00000000",
		"totalUnrealizedProfit": "0.00000000",
		"totalWalletBalance": "0.88308000",
		"updateTime": 1576756674610
	 }
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetDetailOnSubAccountFuturesAccountV2Service().
		Email("test@example.com").
		FuturesType(1).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("abc@test.com", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.Email)
	s.Len(resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.Assets, 1)
	s.Equal("USDT", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.Assets[0].Asset)
	s.Equal("0.00000000", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.Assets[0].InitialMargin)
	s.Equal("0.00000000", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.Assets[0].MaintenanceMargin)
	s.Equal("0.88308000", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.Assets[0].MarginBalance)
	s.Equal("0.88308000", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.Assets[0].MaxWithdrawAmount)
	s.Equal("0.00000000", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.Assets[0].OpenOrderInitialMargin)
	s.Equal("0.00000000", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.Assets[0].PositionInitialMargin)
	s.Equal("0.00000000", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.Assets[0].UnrealizedProfit)
	s.Equal("0.88308000", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.Assets[0].WalletBalance)
	s.True(resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.CanDeposit)
	s.True(resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.CanTrade)
	s.True(resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.CanWithdraw)
	s.Equal(2, resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.FeeTier)
	s.Equal("0.88308000", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.MaxWithdrawAmount)
	s.Equal("0.00000000", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.TotalInitialMargin)
	s.Equal("0.00000000", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.TotalMaintenanceMargin)
	s.Equal("0.88308000", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.TotalMarginBalance)
	s.Equal("0.00000000", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.TotalOpenOrderInitialMargin)
	s.Equal("0.00000000", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.TotalPositionInitialMargin)
	s.Equal("0.00000000", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.TotalUnrealizedProfit)
	s.Equal("0.88308000", resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.TotalWalletBalance)
	s.Equal(uint64(1576756674610), resp.(*GetDetailOnSubAccountFuturesAccountV2USDTResp).FutureAccountResp.UpdateTime)
}

func (s *subAccountTestSuite) TestGetDetailOnSubAccountMarginAccount() {
	data := []byte(`
	{
		"email": "test@example.com",
		"marginLevel": "1.00000000",
		"totalAssetOfBtc": "0.10000000",
		"totalLiabilityOfBtc": "0.00000000",
		"totalNetAssetOfBtc": "0.10000000",
		"marginTradeCoeffVo": {
			"forceLiquidationRate": "1.00000000",
			"marginCallBar": "1.00000000",
			"normalBar": "1.00000000"
		},
		"marginUserAssetVoList": [
			{
				"asset": "BTC",
				"borrowed": "0.00000000",
				"free": "0.10000000",
				"interest": "0.00000000",
				"locked": "0.00000000",
				"netAsset": "0.10000000"
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetDetailOnSubAccountMarginAccountService().
		Email("test@example.com").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("test@example.com", resp.Email)
	s.Equal("1.00000000", resp.MarginLevel)
	s.Equal("0.10000000", resp.TotalAssetOfBtc)
	s.Equal("0.00000000", resp.TotalLiabilityOfBtc)
	s.Equal("0.10000000", resp.TotalNetAssetOfBtc)
	s.Equal("1.00000000", resp.MarginTradeCoeffVo.ForceLiquidationRate)
	s.Equal("1.00000000", resp.MarginTradeCoeffVo.MarginCallBar)
	s.Equal("1.00000000", resp.MarginTradeCoeffVo.NormalBar)
	s.Len(resp.MarginUserAssetVoList, 1)
	s.Equal("BTC", resp.MarginUserAssetVoList[0].Asset)
	s.Equal("0.00000000", resp.MarginUserAssetVoList[0].Borrowed)
	s.Equal("0.10000000", resp.MarginUserAssetVoList[0].Free)
	s.Equal("0.00000000", resp.MarginUserAssetVoList[0].Interest)
	s.Equal("0.00000000", resp.MarginUserAssetVoList[0].Locked)
	s.Equal("0.10000000", resp.MarginUserAssetVoList[0].NetAsset)
}

func (s *subAccountTestSuite) TestGetFuturesPositionRiskOfSubAccount() {
	data := []byte(`{
		"entryPrice": "36621.32911",
		"leverage": "1",
		"maxNotional": "2000000",
		"liquidationPrice": "0",
		"markPrice": "36772.64828",
		"positionAmount": "1.00000000",
		"symbol": "BTCUSDT",
		"unrealizedProfit": "11.99924692"
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetFuturesPositionRiskOfSubAccountService().
		Email("example@example.com").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("36621.32911", resp.EntryPrice)
	s.Equal("1", resp.Leverage)
	s.Equal("2000000", resp.MaxNotional)
	s.Equal("0", resp.LiquidationPrice)
	s.Equal("36772.64828", resp.MarkPrice)
	s.Equal("1.00000000", resp.PositionAmount)
	s.Equal("BTCUSDT", resp.Symbol)
	s.Equal("11.99924692", resp.UnrealizedProfit)
}

func (s *subAccountTestSuite) TestGetFuturesPositionRiskOfSubAccountV2() {
	data := []byte(`
	{
		"futurePositionRiskVos": [
		   {
			  "entryPrice": "9975.12000",
			  "leverage": "50",
			  "maxNotional": "1000000",
			  "liquidationPrice": "7963.54",
			  "markPrice": "9973.50770517",
			  "positionAmount": "0.010",
			  "symbol": "BTCUSDT",
			  "unrealizedProfit": "-0.01612295"
		   }
		 ]
	  }`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetFuturesPositionRiskOfSubAccountV2Service().
		Email("test@example.com").
		FuturesType(1).
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp.(*GetFuturesPositionRiskOfSubAccountV2USDTResp).FuturePositionRiskVos, 1)
	s.Equal("9975.12000", resp.(*GetFuturesPositionRiskOfSubAccountV2USDTResp).FuturePositionRiskVos[0].EntryPrice)
	s.Equal("50", resp.(*GetFuturesPositionRiskOfSubAccountV2USDTResp).FuturePositionRiskVos[0].Leverage)
	s.Equal("1000000", resp.(*GetFuturesPositionRiskOfSubAccountV2USDTResp).FuturePositionRiskVos[0].MaxNotional)
	s.Equal("7963.54", resp.(*GetFuturesPositionRiskOfSubAccountV2USDTResp).FuturePositionRiskVos[0].LiquidationPrice)
	s.Equal("9973.50770517", resp.(*GetFuturesPositionRiskOfSubAccountV2USDTResp).FuturePositionRiskVos[0].MarkPrice)
	s.Equal("0.010", resp.(*GetFuturesPositionRiskOfSubAccountV2USDTResp).FuturePositionRiskVos[0].PositionAmount)
	s.Equal("BTCUSDT", resp.(*GetFuturesPositionRiskOfSubAccountV2USDTResp).FuturePositionRiskVos[0].Symbol)
	s.Equal("-0.01612295", resp.(*GetFuturesPositionRiskOfSubAccountV2USDTResp).FuturePositionRiskVos[0].UnrealizedProfit)
}

func (s *subAccountTestSuite) TestGetIPRestrictionForSubAccountAPIKey() {
	data := []byte(`
	{
		"ipRestrict": "NONE",
		"ipList": [
			{
				"ip": "0.0.0.0/0"
			}
		],
		"updateTime": 1617137789000,
		"apiKey": "some-api-key"
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetIPRestrictionForSubAccountAPIKeyService().
		Email("subaccount@example.com").
		SubAccountApiKey("some-api-key").
		Do(context.Background())

	s.r().NoError(err)
	s.NotNil(resp)
	s.Equal("NONE", resp.IpRestrict)
	s.Len(resp.IpList, 1)
	s.Equal("0.0.0.0/0", resp.IpList[0].Ip)
	s.Equal(uint64(1617137789000), resp.UpdateTime)
	s.Equal("some-api-key", resp.ApiKey)
}

func (s *subAccountTestSuite) TestGetSubAccountDepositHistory() {
	data := []byte(`
	{
		"depositList": [
			{
				"id": 123,
				"amount": "1.00000000",
				"coin": "BTC",
				"network": "",
				"status": 1,
				"address": "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
				"addressTag": "",
				"txId": "tx123",
				"insertTime": 1613450271000,
				"transferType": 0,
				"confirmTimes": "0/1",
				"unlockConfirm": 2,
				"walletType": 1
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetSubAccountDepositHistoryService().
		Email("test@test.com").
		Coin("BTC").
		Status(1).
		StartTime(1613450271000).
		EndTime(1613536671000).
		Limit(500).
		Offset(0).
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp.DepositList, 1)
	s.Equal(int64(123), resp.DepositList[0].Id)
	s.Equal("1.00000000", resp.DepositList[0].Amount)
	s.Equal("BTC", resp.DepositList[0].Coin)
	s.Equal("", resp.DepositList[0].Network)
	s.Equal(int64(1), resp.DepositList[0].Status)
	s.Equal("1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", resp.DepositList[0].Address)
	s.Equal("", resp.DepositList[0].AddressTag)
	s.Equal("tx123", resp.DepositList[0].TxId)
	s.Equal(uint64(1613450271000), resp.DepositList[0].InsertTime)
	s.Equal(int64(0), resp.DepositList[0].TransferType)
	s.Equal("0/1", resp.DepositList[0].ConfirmTimes)
	s.Equal(int64(2), resp.DepositList[0].UnlockConfirm)
	s.Equal(1, resp.DepositList[0].WalletType)
}

func (s *subAccountTestSuite) TestGetSubAccountStatus() {
	data := []byte(`
	{
		"email": "user@example.com",
		"isSubUserEnabled": true,
		"isUserActive": true,
		"insertTime": 1621829948472,
		"isMarginEnabled": true,
		"isFuturesEnabled": false,
		"mobile": 1234567890
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetSubAccountStatusService().
		Email("user@example.com").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("user@example.com", resp.Email)
	s.True(resp.IsSubUserEnabled)
	s.True(resp.IsUserActive)
	s.Equal(uint64(1621829948472), resp.InsertTime)
	s.True(resp.IsMarginEnabled)
	s.False(resp.IsFuturesEnabled)
	s.Equal(int64(1234567890), resp.Mobile)
}

func (s *subAccountTestSuite) TestGetSummaryOfSubAccountFuturesAccount() {
	data := []byte(`
	{
		"totalInitialMargin": "1.00000000",
		"totalMaintenanceMargin": "1.00000000",
		"totalMarginBalance": "1.00000000",
		"totalOpenOrderInitialMargin": "1.00000000",
		"totalPositionInitialMargin": "1.00000000",
		"totalUnrealizedProfit": "1.00000000",
		"totalWalletBalance": "1.00000000",
		"asset": "BTC",
		"subAccountList": [
			{
				"email": "example@example.com",
				"totalInitialMargin": "1.00000000",
				"totalMaintenanceMargin": "1.00000000",
				"totalMarginBalance": "1.00000000",
				"totalOpenOrderInitialMargin": "1.00000000",
				"totalPositionInitialMargin": "1.00000000",
				"totalUnrealizedProfit": "1.00000000",
				"totalWalletBalance": "1.00000000",
				"asset": "BTC"
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetSummaryOfSubAccountFuturesAccountService().Do(context.Background())

	s.r().NoError(err)
	s.Equal("1.00000000", resp.TotalInitialMargin)
	s.Equal("1.00000000", resp.TotalMaintenanceMargin)
	s.Equal("1.00000000", resp.TotalMarginBalance)
	s.Equal("1.00000000", resp.TotalOpenOrderInitialMargin)
	s.Equal("1.00000000", resp.TotalPositionInitialMargin)
	s.Equal("1.00000000", resp.TotalUnrealizedProfit)
	s.Equal("1.00000000", resp.TotalWalletBalance)
	s.Equal("BTC", resp.Asset)
	s.Len(resp.SubAccountList, 1)
	s.Equal("example@example.com", resp.SubAccountList[0].Email)
	s.Equal("1.00000000", resp.SubAccountList[0].TotalInitialMargin)
	s.Equal("1.00000000", resp.SubAccountList[0].TotalMaintenanceMargin)
	s.Equal("1.00000000", resp.SubAccountList[0].TotalMarginBalance)
	s.Equal("1.00000000", resp.SubAccountList[0].TotalOpenOrderInitialMargin)
	s.Equal("1.00000000", resp.SubAccountList[0].TotalPositionInitialMargin)
	s.Equal("1.00000000", resp.SubAccountList[0].TotalUnrealizedProfit)
	s.Equal("1.00000000", resp.SubAccountList[0].TotalWalletBalance)
	s.Equal("BTC", resp.SubAccountList[0].Asset)
}

func (s *subAccountTestSuite) TestGetSummaryOfSubAccountFuturesAccountV2() {
	data := []byte(`
	{
		"futureAccountSummaryResp": {
			"totalInitialMargin": "100.00000000",
			"totalMaintenanceMargin": "50.00000000",
			"totalMarginBalance": "150.00000000",
			"totalOpenOrderInitialMargin": "25.00000000",
			"totalPositionInitialMargin": "75.00000000",
			"totalUnrealizedProfit": "10.00000000",
			"totalWalletBalance": "160.00000000",
			"asset": "BTC",
			"subAccountList": [
				{
					"email": "test1@test.com",
					"totalInitialMargin": "50.00000000",
					"totalMaintenanceMargin": "25.00000000",
					"totalMarginBalance": "75.00000000",
					"totalOpenOrderInitialMargin": "15.00000000",
					"totalPositionInitialMargin": "45.00000000",
					"totalUnrealizedProfit": "5.00000000",
					"totalWalletBalance": "80.00000000",
					"asset": "BTC"
				},
				{
					"email": "test2@test.com",
					"totalInitialMargin": "50.00000000",
					"totalMaintenanceMargin": "25.00000000",
					"totalMarginBalance": "75.00000000",
					"totalOpenOrderInitialMargin": "10.00000000",
					"totalPositionInitialMargin": "30.00000000",
					"totalUnrealizedProfit": "5.00000000",
					"totalWalletBalance": "80.00000000",
					"asset": "BTC"
				}
			]
		}
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetSummaryOfSubAccountFuturesAccountV2Service().
		FuturesType(1).
		Page(1).
		Limit(100).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("100.00000000", resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.TotalInitialMargin)
	s.Equal("50.00000000", resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.TotalMaintenanceMargin)
	s.Equal("150.00000000", resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.TotalMarginBalance)
	s.Equal("25.00000000", resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.TotalOpenOrderInitialMargin)
	s.Equal("75.00000000", resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.TotalPositionInitialMargin)
	s.Equal("10.00000000", resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.TotalUnrealizedProfit)
	s.Equal("160.00000000", resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.TotalWalletBalance)
	s.Equal("BTC", resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.Asset)
	s.Len(resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.SubAccountList, 2)
	s.Equal("test1@test.com", resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.SubAccountList[0].Email)
	s.Equal("50.00000000", resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.SubAccountList[0].TotalInitialMargin)
	s.Equal("25.00000000", resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.SubAccountList[0].TotalMaintenanceMargin)
	s.Equal("75.00000000", resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.SubAccountList[0].TotalMarginBalance)
	s.Equal("15.00000000", resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.SubAccountList[0].TotalOpenOrderInitialMargin)
	s.Equal("45.00000000", resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.SubAccountList[0].TotalPositionInitialMargin)
	s.Equal("5.00000000", resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.SubAccountList[0].TotalUnrealizedProfit)
	s.Equal("80.00000000", resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.SubAccountList[0].TotalWalletBalance)
	s.Equal("BTC", resp.(*GetSummaryOfSubAccountFuturesAccountV2USDTResp).FutureAccountSummaryResp.SubAccountList[0].Asset)
}

func (s *marginTestSuite) TestGetSummaryOfSubAccountMarginAccount() {
	data := []byte(`
	{
		"totalAssetOfBtc": "0.10000000",
		"totalLiabilityOfBtc": "0.05000000",
		"totalNetAssetOfBtc": "0.05000000",
		"subAccountList": [
			{
				"email": "test@example.com",
				"totalAssetOfBtc": "0.05000000",
				"totalLiabilityOfBtc": "0.00000000",
				"totalNetAssetOfBtc": "0.05000000"
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetSummaryOfSubAccountMarginAccountService().
		Do(context.Background())

	s.r().NoError(err)
	s.NotNil(resp)
	s.Equal("0.10000000", resp.TotalAssetOfBtc)
	s.Equal("0.05000000", resp.TotalLiabilityOfBtc)
	s.Equal("0.05000000", resp.TotalNetAssetOfBtc)
	s.Len(resp.SubAccountList, 1)
	s.Equal("test@example.com", resp.SubAccountList[0].Email)
	s.Equal("0.05000000", resp.SubAccountList[0].TotalAssetOfBtc)
	s.Equal("0.00000000", resp.SubAccountList[0].TotalLiabilityOfBtc)
	s.Equal("0.05000000", resp.SubAccountList[0].TotalNetAssetOfBtc)
}

func (s *subAccountTestSuite) TestMarginTransferForSubAccount() {
	data := []byte(`
	{
		"txnId": 12345
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewMarginTransferForSubAccountService().
		Email("test@example.com").
		Asset("BTC").
		Amount(1.23).
		TransferType(1).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(12345, resp.TxnId)
}

func (s *subAccountTestSuite) TestQueryManagedSubAccountAssetDetails() {
	data := []byte(`
	{
		"assetDetail": [
			{
				"coin": "BTC",
				"name": "Bitcoin",
				"totalBalance": "1.00000000",
				"availableBalance": "0.50000000",
				"inOrder": "0.50000000",
				"btcValue": "50000.00"
			},
			{
				"coin": "ETH",
				"name": "Ethereum",
				"totalBalance": "10.00000000",
				"availableBalance": "5.00000000",
				"inOrder": "5.00000000",
				"btcValue": "3000.00"
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewQueryManagedSubAccountAssetDetailsService().
		Email("test@example.com").
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp.AssetDetail, 2)
	s.Equal("BTC", resp.AssetDetail[0].Coin)
	s.Equal("Bitcoin", resp.AssetDetail[0].Name)
	s.Equal("1.00000000", resp.AssetDetail[0].TotalBalance)
	s.Equal("0.50000000", resp.AssetDetail[0].AvailableBalance)
	s.Equal("0.50000000", resp.AssetDetail[0].InOrder)
	s.Equal("50000.00", resp.AssetDetail[0].BtcValue)
	s.Equal("ETH", resp.AssetDetail[1].Coin)
	s.Equal("Ethereum", resp.AssetDetail[1].Name)
	s.Equal("10.00000000", resp.AssetDetail[1].TotalBalance)
	s.Equal("5.00000000", resp.AssetDetail[1].AvailableBalance)
	s.Equal("5.00000000", resp.AssetDetail[1].InOrder)
	s.Equal("3000.00", resp.AssetDetail[1].BtcValue)
}

func (s *subAccountTestSuite) TestQueryManagedSubAccountList() {
	data := []byte(`
	{
		"total": 1,
		"managerSubUserInfoVoList": [
			{
				"rootUserId": 123456,
				"managersubUserId": 789012,
				"bindParentUserId": 345678,
				"email": "test@example.com",
				"insertTimestamp": 1613450271000,
				"bindParentEmail": "parent@example.com",
				"isSubUserEnabled": true,
				"isUserActive": true,
				"isMarginEnabled": true,
				"isFutureEnabled": true,
				"isSignedLVTRiskAgreement": true
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewQueryManagedSubAccountList().
		Email("test@example.com").
		Page(1).
		Limit(500).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(1, resp.Total)
	s.Len(resp.ManagerSubUserInfoVoList, 1)
	s.Equal(123456, resp.ManagerSubUserInfoVoList[0].RootUserId)
	s.Equal(789012, resp.ManagerSubUserInfoVoList[0].ManagersubUserId)
	s.Equal(345678, resp.ManagerSubUserInfoVoList[0].BindParentUserId)
	s.Equal("test@example.com", resp.ManagerSubUserInfoVoList[0].Email)
	s.Equal(uint64(1613450271000), resp.ManagerSubUserInfoVoList[0].InsertTimestamp)
	s.Equal("parent@example.com", resp.ManagerSubUserInfoVoList[0].BindParentEmail)
	s.True(resp.ManagerSubUserInfoVoList[0].IsSubUserEnabled)
	s.True(resp.ManagerSubUserInfoVoList[0].IsUserActive)
	s.True(resp.ManagerSubUserInfoVoList[0].IsMarginEnabled)
	s.True(resp.ManagerSubUserInfoVoList[0].IsFutureEnabled)
	s.True(resp.ManagerSubUserInfoVoList[0].IsSignedLVTRiskAgreement)
}

func (s *accountTestSuite) TestQueryManagedSubAccountMarginAssetDetails() {
	data := []byte(`
	{
		"marginLevel": "999.00000000",
		"totalAssetOfBtc": "1.00000000",
		"totalLiabilityOfBtc": "0.00000000",
		"totalNetAssetOfBtc": "1.00000000",
		"userAssets": [
			{
				"asset": "BTC",
				"borrowed": "0.00000000",
				"free": "1.00000000",
				"interest": "0.00000000",
				"locked": "0.00000000",
				"netAsset": "1.00000000"
			},
			{
				"asset": "ETH",
				"borrowed": "0.00000000",
				"free": "1.00000000",
				"interest": "0.00000000",
				"locked": "0.00000000",
				"netAsset": "1.00000000"
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewQueryManagedSubAccountMarginAssetDetailsService().
		Email("test@example.com").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("999.00000000", resp.MarginLevel)
	s.Equal("1.00000000", resp.TotalAssetOfBtc)
	s.Equal("0.00000000", resp.TotalLiabilityOfBtc)
	s.Equal("1.00000000", resp.TotalNetAssetOfBtc)
	s.Len(resp.UserAssets, 2)
	s.Equal("BTC", resp.UserAssets[0].Asset)
	s.Equal("0.00000000", resp.UserAssets[0].Borrowed)
	s.Equal("1.00000000", resp.UserAssets[0].Free)
	s.Equal("0.00000000", resp.UserAssets[0].Interest)
	s.Equal("0.00000000", resp.UserAssets[0].Locked)
	s.Equal("1.00000000", resp.UserAssets[0].NetAsset)
	s.Equal("ETH", resp.UserAssets[1].Asset)
	s.Equal("0.00000000", resp.UserAssets[1].Borrowed)
	s.Equal("1.00000000", resp.UserAssets[1].Free)
	s.Equal("0.00000000", resp.UserAssets[1].Interest)
	s.Equal("0.00000000", resp.UserAssets[1].Locked)
	s.Equal("1.00000000", resp.UserAssets[1].NetAsset)
}

func (s *subAccountTestSuite) TestQueryManagedSubAccountSnapshotService() {
	data := []byte(`
	{
		"code": 200,
		"msg": "success",
		"snapshotVos": [
			{
				"data": [
					{
						"balances": [
							{
								"asset": "BTC",
								"free": "0.1",
								"locked": "0.2"
							},
							{
								"asset": "ETH",
								"free": "1.1",
								"locked": "0.0"
							}
						],
						"totalAssetOfBtc": "0.025"
					}
				],
				"type": "SPOT",
				"updateTime": 1619140921388
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewQueryManagedSubAccountSnapshotService().
		Email("example@test.com").
		SubType("SPOT").
		StartTime(1619137321000).
		EndTime(1619144521000).
		Limit(10).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(200, resp.Code)
	s.Equal("success", resp.Msg)
	s.Len(resp.SnapshotVos, 1)
	s.Len(resp.SnapshotVos[0].Data, 1)
	s.Len(resp.SnapshotVos[0].Data[0].Balances, 2)
	s.Equal("BTC", resp.SnapshotVos[0].Data[0].Balances[0].Asset)
	s.Equal("0.1", resp.SnapshotVos[0].Data[0].Balances[0].Free)
	s.Equal("0.2", resp.SnapshotVos[0].Data[0].Balances[0].Locked)
	s.Equal("ETH", resp.SnapshotVos[0].Data[0].Balances[1].Asset)
	s.Equal("1.1", resp.SnapshotVos[0].Data[0].Balances[1].Free)
	s.Equal("0.0", resp.SnapshotVos[0].Data[0].Balances[1].Locked)
	s.Equal("0.025", resp.SnapshotVos[0].Data[0].TotalAssetOfBtc)
	s.Equal("SPOT", resp.SnapshotVos[0].Type)
	s.Equal(uint64(1619140921388), resp.SnapshotVos[0].UpdateTime)
}

func (s *subAccountTestSuite) TestQueryManagedSubAccountTransferLogService() {
	data := []byte(`
	{
		"managerSubTransferHistoryVos": [
			{
				"fromEmail": "foo@bar.com",
				"fromAccountType": "SPOT",
				"toEmail": "baz@qux.com",
				"toAccountType": "MARGIN",
				"asset": "BTC",
				"amount": 100000000,
				"scheduledData": 0,
				"createTime": 1613450271000,
				"status": "SUCCESS"
			}
		],
		"count": 1
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewQueryManagedSubAccountTransferLogService().
		Email("foo@bar.com").
		StartTime(1613450271000).
		EndTime(1613536671000).
		Page(1).
		Limit(500).
		Transfers("IN").
		TransferFunctionAccountType("MARGIN").
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp.ManagerSubTransferHistoryVos, 1)
	s.Equal("foo@bar.com", resp.ManagerSubTransferHistoryVos[0].FromEmail)
	s.Equal("SPOT", resp.ManagerSubTransferHistoryVos[0].FromAccountType)
	s.Equal("baz@qux.com", resp.ManagerSubTransferHistoryVos[0].ToEmail)
	s.Equal("MARGIN", resp.ManagerSubTransferHistoryVos[0].ToAccountType)
	s.Equal("BTC", resp.ManagerSubTransferHistoryVos[0].Asset)
	s.Equal(100000000, resp.ManagerSubTransferHistoryVos[0].Amount)
	s.Equal(0, resp.ManagerSubTransferHistoryVos[0].ScheduledData)
	s.Equal(uint64(1613450271000), resp.ManagerSubTransferHistoryVos[0].CreateTime)
	s.Equal("SUCCESS", resp.ManagerSubTransferHistoryVos[0].Status)
	s.Equal(1, resp.Count)
}

func (s *subAccountTestSuite) TestQueryManagedSubAccountTransferLogForTradingTeam() {
	data := []byte(`
	{
		"managerSubTransferHistoryVos": [
			{
				"fromEmail": "test1@test.com",
				"fromAccountType": "SPOT",
				"toEmail": "test2@test.com",
				"toAccountType": "MARGIN",
				"asset": "BTC",
				"amount": "1.00000000",
				"scheduledData": 1613450271000,
				"createTime": 1613450271000,
				"status": "CONFIRMED"
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewQueryManagedSubAccountTransferLogForTradingTeamService().
		Email("test@test.com").
		StartTime(1613450271000).
		EndTime(1613536671000).
		Page(1).
		Limit(500).
		Transfers("BETWEEN_PARENT_TRADING_ACCOUNTS").
		TransferFunctionAccountType("MARGIN").
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp.ManagerSubTransferHistoryVos, 1)
	s.Equal("test1@test.com", resp.ManagerSubTransferHistoryVos[0].FromEmail)
	s.Equal("SPOT", resp.ManagerSubTransferHistoryVos[0].FromAccountType)
	s.Equal("test2@test.com", resp.ManagerSubTransferHistoryVos[0].ToEmail)
	s.Equal("MARGIN", resp.ManagerSubTransferHistoryVos[0].ToAccountType)
	s.Equal("BTC", resp.ManagerSubTransferHistoryVos[0].Asset)
	s.Equal("1.00000000", resp.ManagerSubTransferHistoryVos[0].Amount)
	s.Equal(int64(1613450271000), resp.ManagerSubTransferHistoryVos[0].ScheduledData)
	s.Equal(uint64(1613450271000), resp.ManagerSubTransferHistoryVos[0].CreateTime)
	s.Equal("CONFIRMED", resp.ManagerSubTransferHistoryVos[0].Status)
}

func (s *subAccountTestSuite) TestQuerySubAccountAssets() {
	data := []byte(`
	{
		"balances": [
			{
				"asset": "BTC",
				"free": "0.1",
				"locked": "0.2"
			},
			{
				"asset": "ETH",
				"free": "1.0",
				"locked": "2.0"
			}
		]
	}`)

	expectedBalances := []struct {
		Asset  string
		Free   string
		Locked string
	}{
		{
			Asset:  "BTC",
			Free:   "0.1",
			Locked: "0.2",
		},
		{
			Asset:  "ETH",
			Free:   "1.0",
			Locked: "2.0",
		},
	}

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewQuerySubAccountAssetsService().
		Email("test@example.com").
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp.Balances, 2)
	for i, balance := range resp.Balances {
		s.Equal(expectedBalances[i].Asset, balance.Asset)
		s.Equal(expectedBalances[i].Free, balance.Free)
		s.Equal(expectedBalances[i].Locked, balance.Locked)
	}
}

func (s *subAccountTestSuite) TestQuerySubAccountAssetsForMasterAccount() {
	data := []byte(`
	{
		"balances": [
			{
				"asset": "BTC",
				"free": "0.005",
				"locked": "0.005"
			},
			{
				"asset": "ETH",
				"free": "0.01",
				"locked": "0.01"
			}
		]
	}
	`)

	expectedResp := &QuerySubAccountAssetsForMasterAccountResp{
		Balances: []struct {
			Asset  string `json:"asset"`
			Free   string `json:"free"`
			Locked string `json:"locked"`
		}{
			{
				Asset:  "BTC",
				Free:   "0.005",
				Locked: "0.005",
			},
			{
				Asset:  "ETH",
				Free:   "0.01",
				Locked: "0.01",
			},
		},
	}

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewQuerySubAccountAssetsForMasterAccountService().
		Email("example@email.com").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(expectedResp, resp)
}

func (s *subAccountTestSuite) TestQuerySubAccountSpotAssetsSummary() {
	data := []byte(`
	{
		"totalCount": 2,
		"masterAccountTotalAsset": "1.2345",
		"spotSubUserAssetBtcVoList": [
			{
				"email": "user1@example.com",
				"toAsset": "BTC"
			},
			{
				"email": "user2@example.com",
				"toAsset": "ETH"
			}
		]
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewQuerySubAccountSpotAssetsSummaryService().
		Email("example@example.com").
		Page(1).
		Size(50).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(int64(2), resp.TotalCount)
	s.Equal("1.2345", resp.MasterAccountTotalAsset)
	s.Len(resp.SpotSubUserAssetBtcVoList, 2)
	s.Equal("user1@example.com", resp.SpotSubUserAssetBtcVoList[0].Email)
	s.Equal("BTC", resp.SpotSubUserAssetBtcVoList[0].ToAsset)
	s.Equal("user2@example.com", resp.SpotSubUserAssetBtcVoList[1].Email)
	s.Equal("ETH", resp.SpotSubUserAssetBtcVoList[1].ToAsset)
}

func (s *subAccountTestSuite) TestQuerySubAccountTransactionTatistics() {
	data := []byte(`{
		"recent30BtcTotal": "0.00000000",
		"recent30BtcFuturesTotal": "0.00000000",
		"recent30BtcMarginTotal": "0.00000000",
		"recent30BusdTotal": "0.00000000",
		"recent30BusdFuturesTotal": "0.00000000",
		"recent30BusdMarginTotal": "0.00000000",
		"tradeInfoVos": [
		  {
			"userId": 123,
			"btc": 1,
			"btcFutures": 2,
			"btcMargin": 3,
			"busd": 4,
			"busdFutures": 5,
			"busdMargin": 6,
			"date": 1617235200000
		  }
		]
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	email := "test@test.com"
	resp, err := s.client.NewQuerySubAccountTransactionTatistics().
		Email(email).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("0.00000000", resp.Recent30BtcTotal)
	s.Equal("0.00000000", resp.Recent30BtcFuturesTotal)
	s.Equal("0.00000000", resp.Recent30BtcMarginTotal)
	s.Equal("0.00000000", resp.Recent30BusdTotal)
	s.Equal("0.00000000", resp.Recent30BusdFuturesTotal)
	s.Equal("0.00000000", resp.Recent30BusdMarginTotal)
	s.Len(resp.TradeInfoVos, 1)
	s.Equal(int64(123), resp.TradeInfoVos[0].UserId)
	s.Equal(1, resp.TradeInfoVos[0].Btc)
	s.Equal(2, resp.TradeInfoVos[0].BtcFutures)
	s.Equal(3, resp.TradeInfoVos[0].BtcMargin)
	s.Equal(4, resp.TradeInfoVos[0].Busd)
	s.Equal(5, resp.TradeInfoVos[0].BusdFutures)
	s.Equal(6, resp.TradeInfoVos[0].BusdMargin)
	s.Equal(int64(1617235200000), resp.TradeInfoVos[0].Date)
}

func (s *subAccountTestSuite) TestSubAccountFuturesAssetTransfer() {
	data := []byte(`{"success": true, "txnId": "123"}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewSubAccountFuturesAssetTransferService().
		FromEmail("from@test.com").
		ToEmail("to@test.com").
		FuturesType(1).
		Asset("BTC").
		Amount(1.23).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(true, resp.Success)
	s.Equal("123", resp.TxnId)
}

func (s *subAccountTestSuite) TestQuerySubAccountFuturesAssetTransferHistory() {
	data := []byte(`
	{
		"success": true,
		"futuresType": 1,
		"transfers": [
			{
				"from": "xxx",
				"to": "yyy",
				"asset": "BTC",
				"qty": "0.5",
				"tranId": 123,
				"time": 1613450271000
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewQuerySubAccountFuturesAssetTransferHistoryService().
		Email("test@example.com").
		FuturesType(1).
		StartTime(1613450271000).
		EndTime(1613536671000).
		Page(1).
		Limit(500).
		Do(context.Background())

	s.r().NoError(err)
	s.True(resp.Success)
	s.Equal(int64(1), resp.FuturesType)
	s.Len(resp.Transfers, 1)
	s.Equal("xxx", resp.Transfers[0].From)
	s.Equal("yyy", resp.Transfers[0].To)
	s.Equal("BTC", resp.Transfers[0].Asset)
	s.Equal("0.5", resp.Transfers[0].Qty)
	s.Equal(int64(123), resp.Transfers[0].TranId)
	s.Equal(uint64(1613450271000), resp.Transfers[0].Time)
}

func (s *subAccountTestSuite) TestQuerySubAccountList() {
	data := []byte(`
	{
		"subAccounts": [
			{
				"email": "user@example.com",
				"isFreeze": true,
				"createTime": 1613450271000,
				"isManagedSubAccount": true,
				"isAssetManagementSubAccount": false
			},
			{
				"email": "user2@example.com",
				"isFreeze": false,
				"createTime": 1613450271000,
				"isManagedSubAccount": false,
				"isAssetManagementSubAccount": true
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewQuerySubAccountListService().
		Email("user@example.com").
		IsFreeze("true").
		Page(1).
		Limit(500).
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp.SubAccounts, 2)
	s.Equal("user@example.com", resp.SubAccounts[0].Email)
	s.Equal(true, resp.SubAccounts[0].IsFreeze)
	s.Equal(uint64(1613450271000), resp.SubAccounts[0].CreateTime)
	s.Equal(true, resp.SubAccounts[0].IsManagedSubAccount)
	s.Equal(false, resp.SubAccounts[0].IsAssetManagementSubAccount)
	s.Equal("user2@example.com", resp.SubAccounts[1].Email)
	s.Equal(false, resp.SubAccounts[1].IsFreeze)
	s.Equal(uint64(1613450271000), resp.SubAccounts[1].CreateTime)
	s.Equal(false, resp.SubAccounts[1].IsManagedSubAccount)
	s.Equal(true, resp.SubAccounts[1].IsAssetManagementSubAccount)
}

func (s *subAccountTestSuite) TestQuerySubAccountSpotAssetTransferHistory() {
	data := []byte(`
	{
		"rows": [
			{
				"from": "email1@example.com",
				"to": "email2@example.com",
				"asset": "BTC",
				"qty": "1.00000000",
				"status": "CONFIRMED",
				"tranId": 123456789,
				"time": 1613450271000
			}
		]
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewQuerySubAccountSpotAssetTransferHistoryService().
		FromEmail("email1@example.com").
		ToEmail("email2@example.com").
		StartTime(1613450271000).
		EndTime(1613536671000).
		Page(1).
		Limit(500).
		Do(context.Background())

	s.r().NoError(err)
	s.Len(resp.Rows, 1)
	s.Equal("email1@example.com", resp.Rows[0].From)
	s.Equal("email2@example.com", resp.Rows[0].To)
	s.Equal("BTC", resp.Rows[0].Asset)
	s.Equal("1.00000000", resp.Rows[0].Qty)
	s.Equal("CONFIRMED", resp.Rows[0].Status)
	s.Equal(int64(123456789), resp.Rows[0].TranId)
	s.Equal(uint64(1613450271000), resp.Rows[0].Time)
}

func (s *subAccountTestSuite) TestTransferToMaster() {
	data := []byte(`{"txnId": 123}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewTransferToMasterService().
		Asset("BTC").
		Amount(1.0).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(123, resp.TxnId)
}

func (s *subAccountTestSuite) TestDo() {
	data := []byte(`
	{
		"tranId": 123456789,
		"clientTranId": "123abc"
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewUniversalTransferService().
		FromAccountType("SPOT").
		ToAccountType("MARGIN").
		FromEmail("example1@gmail.com").
		ToEmail("example2@gmail.com").
		ClientTranId("123abc").
		Symbol("BTCUSDT").
		Asset("BTC").
		Amount(1.0).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(123456789, resp.TranId)
	s.Equal("123abc", resp.ClientTranId)
}

func (s *subAccountTestSuite) TestUpdateIPRestrictionForSubAccountAPIKey() {
	data := []byte(`{
		"status": "SUCCESS",
		"ipList": [
			{
				"ip": "192.168.1.1"
			},
			{
				"ip": "192.168.1.2"
			}
		],
		"updateTime": 1617383620000,
		"apiKey": "h-1234567890"
	}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewUpdateIPRestrictionForSubAccountAPIKeyService().
		Email("subaccount@test.com").
		SubAccountApiKey("h-1234567890").
		Status("ENABLED").
		IpAddress("192.168.1.1").
		Do(context.Background())

	s.r().NoError(err)
	s.Equal("SUCCESS", resp.Status)
	s.Len(resp.IpList, 2)
	s.Equal("192.168.1.1", resp.IpList[0].Ip)
	s.Equal("192.168.1.2", resp.IpList[1].Ip)
	s.Equal(uint64(1617383620000), resp.UpdateTime)
	s.Equal("h-1234567890", resp.ApiKey)
}

func (s *subAccountTestSuite) TestWithdrawAssetsFromTheManagedSubAccount() {
	data := []byte(`{"tranId":123}`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewWithdrawAssetsFromTheManagedSubAccountService().
		FromEmail("user@example.com").
		Asset("BTC").
		Amount(1.0).
		TransferDate(1617220000).
		Do(context.Background())

	s.r().NoError(err)
	s.Equal(int64(123), resp.TranId)
}

func (s *subAccountTestSuite) TestManagedSubaccountDepositAddressService() {
	data := []byte(`
	{
		"coin": "USDT",
		"address": "0x206c22d833bb0bb2102da6b7c7d4c3eb14bcf73d",
		"tag": "",
		"url": "https://etherscan.io/address/0x206c22d833bb0bb2102da6b7c7d4c3eb14bcf73d"
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	email := "testsub@gmail.com"
	coin := "a_coin"
	network := "a_network"

	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"email":   email,
			"coin":    coin,
			"network": network,
		})
		s.assertRequestEqual(e, r)
	})

	res, err := s.client.NewGetManagedSubAccountDepositAddressService().
		Email(email).
		Coin(coin).
		Network(network).
		Do(newContext())

	r := s.r()
	r.NoError(err)
	r.Equal("0x206c22d833bb0bb2102da6b7c7d4c3eb14bcf73d", res.Address, "Address")
	r.Equal("USDT", res.Coin, "Coin")
	r.Equal("", res.Tag, "Tag")
	r.Equal("https://etherscan.io/address/0x206c22d833bb0bb2102da6b7c7d4c3eb14bcf73d", res.Url, "URL")
}
