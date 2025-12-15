

```go
package main

import (
	"context"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	client "github.com/binance/binance-connector-go/clients/algo"
	"github.com/binance/binance-connector-go/common/common"
)

// Replace this with the base64(SHA256(subjectPublicKeyInfo)) value you pin to.
const PINNED_PUBLIC_KEY = "YOUR-PINNED-PUBLIC-KEY" // e.g. "AbC...==" (base64 of sha256(pubkey_der))

func main() {
	if err := QueryHistoricalAlgoOrdersSpotAlgo(); err != nil {
		log.Fatalf("failed: %v", err)
	}
}

func QueryHistoricalAlgoOrdersSpotAlgo() error {
	// Build TLS config with certificate pinning
	tlsCfg, err := tlsConfigWithPinning(PINNED_PUBLIC_KEY)
	if err != nil {
		return fmt.Errorf("build tls config: %w", err)
	}

	// Create an HTTP transport that uses our TLS config
	transport := &http.Transport{
		TLSClientConfig:     tlsCfg,
		IdleConnTimeout:     90 * time.Second,
		DisableKeepAlives:   false,
		MaxIdleConnsPerHost: 10,
	}

	// Optionally wrap transport with logging, retrying, etc.
	// For the Binance client, pass the transport as the HTTP agent:
	configuration := common.NewConfigurationRestAPI(
		common.WithBasePath(common.AlgoRestApiProdUrl),
		common.WithApiKey("Your API Key"),
		common.WithApiSecret("Your API Secret"),
	)
	apiClient := client.NewBinanceAlgoClient(
		client.WithRestAPI(configuration),
	)
	resp, err := apiClient.RestApi.SpotAlgoAPI.QueryHistoricalAlgoOrdersSpotAlgo(context.Background()).Execute()
	if err != nil {
		log.Println(err)
		return
	}

	dataValue, _ := json.MarshalIndent(resp.Data, "", "  ")
	log.Printf("Response: %s\n", string(dataValue))
	return nil
}

// tlsConfigWithPinning returns a tls.Config that verifies the chain using system roots
// and also performs public-key pinning: the base64(SHA256(pubkey_der)) must match pinnedKey.
func tlsConfigWithPinning(pinnedKey string) (*tls.Config, error) {
	roots, err := x509.SystemCertPool()
	if err != nil {
		// On some platforms this can return nil, try to create a new pool.
		roots = x509.NewCertPool()
	}
	if roots == nil {
		roots = x509.NewCertPool()
	}

	cfg := &tls.Config{
		// We'll perform verification ourselves in VerifyPeerCertificate so set InsecureSkipVerify.
		InsecureSkipVerify: true,
		VerifyPeerCertificate: func(rawCerts [][]byte, _ [][]*x509.Certificate) error {
			// rawCerts are DER-encoded certs; first is leaf
			if len(rawCerts) == 0 {
				return errors.New("no server certificates presented")
			}

			// Parse certificates
			certs := make([]*x509.Certificate, len(rawCerts))
			for i, asn1Data := range rawCerts {
				cert, err := x509.ParseCertificate(asn1Data)
				if err != nil {
					return fmt.Errorf("failed to parse certificate[%d]: %w", i, err)
				}
				certs[i] = cert
			}

			// Verify chain using system roots
			opts := x509.VerifyOptions{
				Roots:         roots,
				Intermediates: x509.NewCertPool(),
				CurrentTime:   time.Now(),
				// DNSName will be checked by Verify, but we don't have the server name here.
				// If you need strict hostname matching, use cfg.ServerName or call Verify with DNSName set.
			}
			for i, cert := range certs {
				if i == 0 {
					continue // leaf skipping adding to intermediates
				}
				opts.Intermediates.AddCert(cert)
			}

			if _, err := certs[0].Verify(opts); err != nil {
				return fmt.Errorf("certificate chain verification failed: %w", err)
			}

			// Extract public key from leaf cert and marshal into PKIX (SubjectPublicKeyInfo) DER
			pubKeyDER, err := x509.MarshalPKIXPublicKey(certs[0].PublicKey)
			if err != nil {
				return fmt.Errorf("marshal public key: %w", err)
			}

			// Compute base64(SHA256(pubkey_der))
			sum := sha256.Sum256(pubKeyDER)
			hashB64 := base64.StdEncoding.EncodeToString(sum[:])

			if hashB64 != pinnedKey {
				return fmt.Errorf("public key pinning mismatch: expected %s got %s", pinnedKey, hashB64)
			}

			// All checks passed
			return nil
		},
	}

	return cfg, nil
}
```