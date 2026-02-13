package main

import (
	"context"
	"log"

	client "github.com/binance/binance-connector-go/clients/wallet"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	DisableFastWithdrawSwitch()
}

func DisableFastWithdrawSwitch() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.WalletRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceWalletClient(
		client.WithRestAPI(configuration),
	)
	_, err := apiClient.RestApi.AccountAPI.DisableFastWithdrawSwitch(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}
}
