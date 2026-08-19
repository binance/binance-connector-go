package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/stocks"
	"github.com/binance/binance-connector-go/clients/stocks/src/restapi/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	TokenizedConvertStatus()
}

func TokenizedConvertStatus() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.StocksRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceStocksClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.TokenizedAPI.TokenizedConvertStatus(context.Background()).IssuerRequestId("mint-20260505-8f3b9e1a2d3c4b5a").ConvertType(models.TokenizedConvertStatusConvertTypeParameterMint).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	rateLimitsValue, _ := json.MarshalIndent(resp.RateLimits, "", "  ")
	log.Printf("Rate limits: %s\n", string(rateLimitsValue))

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
