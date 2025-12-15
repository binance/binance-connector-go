# Migration Guide: Binance Spot Connector Modularization

With the transition to a modularized structure, the Binance Connector has been split into separate Go libraries, each focusing on a distinct product (e.g., Spot, Futures, etc.). This guide explains how to migrate from the monolithic `github.com/binance/binance-connector-go` package to the new `github.com/binance/binance-connector-go/clients/spot` library.

---

## Key Changes

1. **Package Name**:
   The modularised Margin Trading Connector has been moved to a new package:

   **Old:** `github.com/binance/binance-connector-go`
   **New:** `github.com/binance/binance-connector-go/clients/spot`

2. **Installation**:
   Uninstall the old package and install the new one:

   ```bash
    rm $(which binance-connector-go)
    go get github.com/binance/binance-connector-go/clients/spot
    ```

3. **Imports**:
   Update your import paths.

    **Old:**
    ```go
    import "github.com/binance/binance-connector-go"
    ```

    **New:**
    ```go
    import (
        "github.com/binance/binance-connector-go/clients/spot"
        "github.com/binance/binance-connector-go/common/common"
    )
    ```

4. **Configuration and Client Initialization**:  
   Update your client initialization code.

    **Old:**
    ```go
    client := binance.NewClient(apiKey, secretKey, baseURL)
    ```

    **New:**
    ```go
    configuration := common.NewConfigurationRestAPI(
        common.WithBasePath(common.SpotRestApiProdUrl),
    )
    client := spot.NewBinanceSpotClient(spot.WithRestAPI(configuration))
    ```

5. **Examples and Documentation**:  
   Updated examples can be found in the new repository folders:
    - REST API: [`examples/restapi/`](../examples/restapi/)
    - WebSocket API: [`examples/websocketapi/`](../examples/websocketapi/)
    - WebSocket Streams: [`examples/websocketstreams/`](../examples/websocketstreams/)

---

## Migration Steps

### 1. Replace the old package with the new one

Replace the old package from your project:

```bash
go mod edit -dropreplace github.com/binance/binance-connector-go
go get github.com/binance/binance-connector-go/clients/spot
go get github.com/binance/binance-connector-go/common
```

### 2. Update Import Paths

Change your import statements from:

```go
import "github.com/binance/binance-connector-go"
```

to:

```go
import (
    "github.com/binance/binance-connector-go/clients/spot"
    "github.com/binance/binance-connector-go/common/common"
)
```

### 3. Update Client Initialization

Update the client initialization code from:

```go
client := binance.NewClient(apiKey, secretKey, baseURL)
```

to:

```go
configuration := common.NewConfigurationRestAPI(
    common.WithBasePath(common.SpotRestApiProdUrl),
)
client := spot.NewBinanceSpotClient(spot.WithRestAPI(configuration))
```

### 4. Test and Verify

Run your application to ensure everything works as expected. Refer to the new documentation for any advanced features or configuration options.

---

## Additional Notes

- **Future Modular Packages**: Similar packages for other products (e.g., Wallet, Sub-Account) will follow this pattern.

For more details, refer to the updated [README](../README.md) and [Examples](../examples/).