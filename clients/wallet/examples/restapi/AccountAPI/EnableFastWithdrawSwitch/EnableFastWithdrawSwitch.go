package main

import (
	"context"
	"log"

	client "github.com/binance/binance-connector-go/clients/wallet/src"
	"github.com/binance/binance-connector-go/common/common"
)

func main() {
	EnableFastWithdrawSwitch()
}

func EnableFastWithdrawSwitch() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.WalletRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceWalletClient(
		client.WithRestAPI(configuration),
	)
	_, err := apiClient.RestApi.AccountAPI.EnableFastWithdrawSwitch(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}
}
