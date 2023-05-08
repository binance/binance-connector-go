package binance_connector

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type stakingTestSuite struct {
	baseTestSuite
}

func TestStaking(t *testing.T) {
	suite.Run(t, new(stakingTestSuite))
}

func (s *stakingTestSuite) TestGetStakingProductList() {
	data := []byte(`
	[
		{
			"projectId": "Axs*90",
			"detail": {
				"asset":"AXS",
				"rewardAsset":"AXS",
				"duration":90,
				"renewable":true,
				"apy": "1.2069"
			},
			"quota": {
				"totalPersonalQuota":"2",
				"minimum":"0.001"
			}
		},
		{
			"projectId": "Fio*90",
			"detail": {
				"asset":"FIO",
				"rewardAsset":"FIO",
				"duration":90,
				"renewable":true,
				"apy":"1.0769"
			},
			"quota": {
				"totalPersonalQuota":"600",
				"minimum":"0.1"
			}
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetStakingProductListService().Product("STAKING").Do(context.Background())

	s.r().NoError(err)
	s.r().Equal(2, len(resp))
	s.r().Equal("Axs*90", resp[0].ProjectId)
	s.r().Equal("AXS", resp[0].Detail.Asset)
	s.r().Equal("AXS", resp[0].Detail.RewardAsset)
	s.r().Equal(int64(90), resp[0].Detail.Duration)
	s.r().Equal(true, resp[0].Detail.Renewable)
	s.r().Equal("1.2069", resp[0].Detail.Apy)
	s.r().Equal("2", resp[0].Quota.TotalPersonalQuota)
	s.r().Equal("0.001", resp[0].Quota.Minimum)
}

// PurchaseStakingProductService
func (s *stakingTestSuite) TestPurchaseStakingProduct() {
	data := []byte(`
	{
		"positionId":"12345",
		"success":true
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewPurchaseStakingProductService().Product("STAKING").ProductId("123").Amount(90.00).Do(context.Background())

	s.r().NoError(err)
	s.r().Equal("12345", resp.PositionId)
	s.r().Equal(true, resp.Success)
}

// RedeemStakingProductService
func (s *stakingTestSuite) TestRedeemStakingProduct() {
	data := []byte(`
	{
		"success":true
	}
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewRedeemStakingProductService().Product("STAKING").ProductId("123").Do(context.Background())

	s.r().NoError(err)
	s.r().Equal(true, resp.Success)
}

// GetStakingProductPositionService
func (s *stakingTestSuite) TestGetStakingProductPosition() {
	data := []byte(`
	[
		{
			"positionId":"123123",
			"projectId": "Axs*90",
			"asset":"AXS",
			"amount":"122.09202928",
			"purchaseTime": "1646182276000",
			"duration": "60",
			"accrualDays": "4",
			"rewardAsset":"AXS",
			"APY":"0.2032",
			"rewardAmt": "5.17181528",
			"extraRewardAsset":"BNB",
			"extraRewardAPY":"0.0203",
			"estExtraRewardAmt": "5.17181528",
			"nextInterestPay": "1.29295383",
			"nextInterestPayDate": "1646697600000",
			"payInterestPeriod": "1",
			"redeemAmountEarly": "2802.24068892",
			"interestEndDate": "1651449600000",
			"deliverDate": "1651536000000",
			"redeemPeriod": "1",
			"redeemingAmt":"232.2323",
			"partialAmtDeliverDate":"1651536000000",
			"canRedeemEarly": true,
			"renewable": true,
			"type":"AUTO",
			"status": "HOLDING"
		}
	]
	`)

	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetStakingProductPositionService().Product("STAKING").Do(context.Background())

	s.r().NoError(err)
	s.r().Equal(1, len(resp))
	s.r().Equal("123123", resp[0].PositionId)
	s.r().Equal("Axs*90", resp[0].ProjectId)
	s.r().Equal("AXS", resp[0].Asset)
	s.r().Equal("122.09202928", resp[0].Amount)
	s.r().Equal("1646182276000", resp[0].PurchaseTime)
	s.r().Equal("60", resp[0].Duration)
	s.r().Equal("4", resp[0].AccrualDays)
	s.r().Equal("AXS", resp[0].RewardAsset)
	s.r().Equal("0.2032", resp[0].APY)
	s.r().Equal("5.17181528", resp[0].RewardAmt)
	s.r().Equal("BNB", resp[0].ExtraRewardAsset)
	s.r().Equal("0.0203", resp[0].ExtraRewardAPY)
	s.r().Equal("5.17181528", resp[0].EstExtraRewardAmt)
	s.r().Equal("1.29295383", resp[0].NextInterestPay)
	s.r().Equal("1646697600000", resp[0].NextInterestPayDate)
	s.r().Equal("1", resp[0].PayInterestPeriod)
	s.r().Equal("2802.24068892", resp[0].RedeemAmountEarly)
	s.r().Equal("1651449600000", resp[0].InterestEndDate)
	s.r().Equal("1651536000000", resp[0].DeliverDate)
	s.r().Equal("1", resp[0].RedeemPeriod)
	s.r().Equal("232.2323", resp[0].RedeemingAmt)
	s.r().Equal("1651536000000", resp[0].PartialAmtDeliverDate)
	s.r().Equal(true, resp[0].CanRedeemEarly)
	s.r().Equal(true, resp[0].Renewable)
	s.r().Equal("AUTO", resp[0].Type)
	s.r().Equal("HOLDING", resp[0].Status)
}

// GetStakingHistoryService
func (s *stakingTestSuite) TestGetStakingHistory() {
	data := []byte(`
	[
		{
			"positionId":"123123",
			"time":1575018510000,
			"asset":"BNB",
			"project":"BSC",
			"amount":"21312.23223",
			"lockPeriod":"30",
			"deliverDate":"1575018510000",
			"type":"AUTO",
			"status":"success"
		}
	]
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewGetStakingHistoryService().Product("STAKING").TxnType("SUBSCRIPTION").Do(context.Background())

	s.r().NoError(err)
	s.r().Equal(1, len(resp))
	s.r().Equal("123123", resp[0].PositionId)
	s.r().Equal(uint64(1575018510000), resp[0].Time)
	s.r().Equal("BNB", resp[0].Asset)
	s.r().Equal("BSC", resp[0].Project)
	s.r().Equal("21312.23223", resp[0].Amount)
	s.r().Equal("30", resp[0].LockPeriod)
	s.r().Equal("1575018510000", resp[0].DeliverDate)
	s.r().Equal("AUTO", resp[0].Type)
	s.r().Equal("success", resp[0].Status)
}

// SetAutoStakingService
func (s *stakingTestSuite) TestSetAutoStaking() {
	data := []byte(`
	{
		"success": true
	}
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewSetAutoStakingService().Product("STAKING").PositionId("123").Renewable("true").Do(context.Background())

	s.r().NoError(err)
	s.r().Equal(true, resp.Success)
}

// PersonalLeftQuota
func (s *stakingTestSuite) TestPersonalLeftQuota() {
	data := []byte(`
	[
		{
			"leftPersonalQuota": "1000"
		}
	]
	`)
	s.mockDo(data, nil)
	defer s.assertDo()

	resp, err := s.client.NewPersonalLeftQuotaService().Product("STAKING").Do(context.Background())

	s.r().NoError(err)
	s.r().Equal(1, len(resp))
	s.r().Equal("1000", resp[0].LeftPersonalQuota)
}
