# Compression Configuration

```go
package main

import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/nft"
	"github.com/binance/binance-connector-go/common/v2/common"
)

func main() {
	GetNFTAsset()
}

func GetNFTAsset() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.NFTRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
		common.WithCompression(false),
	)
	apiClient := client.NewBinanceNFTClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.NFTAPI.GetNFTAsset(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```
