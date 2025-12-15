# **Migration Guide: Transition from Monolithic Binance Connector**

With the move towards modularization, Binance connectors are now split into smaller, product-specific libraries. This guide explains how to migrate from the monolithic `binance-connector-go` package to the new modular connectors.

## **Overview of Changes**

| Feature | Monolithic Connector | Modular Connector |
|---------|----------------------|------------------|
| Package Name | `github.com/binance/binance-connector-go` | `github.com/binance/binance-connector-go/clients/<product>` |
| API Coverage | Partial Binance APIs | Individual APIs (Spot, Wallet, Algo Trading, Mining, etc.) |
| Imports | Single package import | Separate package per product |
| Code Structure | One large client | Smaller, focused clients |

## **Migration Steps**

### **Step 1: Uninstall the Monolithic Connector**

If you were using the old connector, remove it from your project:

```bash
go mod edit -dropreplace github.com/binance/binance-connector-go
```

### **Step 2: Install the New Modular Connectors**

Install only the connector(s) you need. For example, to install the Spot Trading connector:

```bash
go get github.com/binance/binance-connector-go/clients/spot
```

To install multiple connectors:

```bash
go get github.com/binance/binance-connector-go/clients/spot github.com/binance/binance-connector-go/clients/margintrading github.com/binance/binance-connector-go/clients/wallet
```

### **Step 3: Update Imports**

Update your import paths:

**Old:**

```go
binance_connector "github.com/binance/binance-connector-go"
```

**New:**

```go
client "github.com/binance/binance-connector-go/clients/spot"
```

### **Step 4: Update Client Initialization**

The new structure introduces a more modular approach to client initialization.

**Old (Monolithic Connector):**

```go
import (
	"context"
	"fmt"

	binance_connector "github.com/binance/binance-connector-go"
)

client := binance_connector.NewClient(apiKey, secretKey, baseURL)
accountInformation, err := client.NewGetAccountService().Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(binance_connector.PrettyPrint(accountInformation))
```

**New (Modular Spot Connector):**

```python
import (
	"context"
	"encoding/json"
	"log"

	client "github.com/binance/binance-connector-go/clients/spot/src"
	"github.com/binance/common/common"
)

func main() {
	GetAccount()
}

func GetAccount() {
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.SpotRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceSpotClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.AccountAPI.GetAccount(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
}
```

### **Step 5: Check for API Differences**

Some function names or response structures may have changed. Refer to the modular connector's documentation for details.

## **Backward Compatibility**

- If a modular connector is **not yet available** for your use case, continue using the monolithic connector (`binance-connector-go`).
- The monolithic connector will remain available, but it is **recommended** to migrate when modular versions are released.

---

## **FAQs**

### **What if my product does not have a modular connector yet?**

You can continue using `binance-connector-go` until the modular version is released.

### **Will the monolithic connector still receive updates?**

Critical bug fixes will be provided, but feature updates will focus on the modular connectors.

### **Where can I find more examples?**

Check the modular connector's documentation for detailed examples.
