package binance_connector

type SignatureType int

const (
	SIGNATURE_HMAC_SHA256 SignatureType = iota
	SIGNATURE_RSA_SHA256
)
