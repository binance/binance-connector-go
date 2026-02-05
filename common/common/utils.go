package common

import (
	"bufio"
	"bytes"
	"compress/flate"
	"compress/gzip"
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	randpkg "math/rand/v2"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/andybalholm/brotli"
	"github.com/google/uuid"
	"github.com/youmark/pkcs8"
	"golang.org/x/net/proxy"
)

var (
	JsonCheck       = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?json)`)
	XmlCheck        = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?xml)`)
	queryParamSplit = regexp.MustCompile(`(^|&)([^&]+)`)
	queryDescape    = strings.NewReplacer("%5B", "[", "%5D", "]")
)

// Get returns the value of the NullableTime.
//
// @return A pointer to the time.Time value, or nil if not set.
func (v NullableTime) Get() *time.Time {
	return v.value
}

// Set assigns a value to the NullableTime and marks it as set.
//
// @param val A pointer to the time.Time value to set.
func (v *NullableTime) Set(val *time.Time) {
	v.value = val
	v.isSet = true
}

// IsSet checks if the NullableTime has been set.
func (v NullableTime) IsSet() bool {
	return v.isSet
}

// Unset clears the value of the NullableTime and marks it as not set.
func (v *NullableTime) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTime creates a new NullableTime with the given value.
//
// @param val A pointer to the time.Time value to set.
// @return A pointer to the newly created NullableTime.
func NewNullableTime(val *time.Time) *NullableTime {
	return &NullableTime{value: val, isSet: true}
}

// MarshalJSON serializes the NullableTime to JSON.
//
// @return The JSON representation of the NullableTime or an error if serialization fails.
func (v NullableTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes JSON data into the NullableTime.
//
// @param src The JSON data as a byte slice.
// @return An error if deserialization fails.
func (v *NullableTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// IsNil checks if the given interface is nil.
//
// @param i The interface to check.
// @return True if the interface is nil, otherwise false.
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	case reflect.Array:
		return reflect.ValueOf(i).IsZero()
	}
	return false
}

// NewStrictDecoder creates a JSON decoder that disallows unknown fields.
//
// @param data The JSON data as a byte slice.
// @return A JSON decoder configured to disallow unknown fields.
func NewStrictDecoder(data []byte) *json.Decoder {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	return dec
}

// ReportError formats an error message.
//
// @param format The format string.
// @param a The arguments to be formatted.
// @return An error with the formatted message.
func ReportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// SelectHeaderContentType selects a content type from the available options.
//
// @param contentTypes A slice of content type strings.
// @return The selected content type string.
func SelectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0]
}

// contains checks if a string slice contains a specific string, case-insensitive.
//
// @param haystack The slice of strings to search within.
// @param needle The string to search for.
// @return True if the slice contains the string, false otherwise.
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}
	return false
}

// ParameterAddToHeaderOrQuery adds a parameter to header or query parameters.
//
// @param headerOrQueryParams The header or query parameters to which the parameter will be added.
// @param keyPrefix The key under which the parameter will be added.
// @param obj The parameter value to be added.
// @param style The style of the parameter (e.g., "deepObject").
// @param collectionType The collection type (e.g., "csv") for handling multiple values.
func ParameterAddToHeaderOrQuery(headerOrQueryParams interface{}, keyPrefix string, obj interface{}, style string, collectionType string) {
	var v = reflect.ValueOf(obj)
	if v == reflect.ValueOf(nil) {
		addValue(headerOrQueryParams, keyPrefix, "null", collectionType)
		return
	}

	var value string
	switch v.Kind() {
	case reflect.Invalid:
		value = "invalid"

	case reflect.Struct, reflect.Slice:
		jsonValue, err := json.Marshal(obj)
		if err != nil {
			return
		}
		value = string(jsonValue)

	case reflect.Map:
		var indValue = reflect.ValueOf(obj)
		if indValue == reflect.ValueOf(nil) {
			return
		}
		iter := indValue.MapRange()
		for iter.Next() {
			k, v := iter.Key(), iter.Value()
			ParameterAddToHeaderOrQuery(headerOrQueryParams, fmt.Sprintf("%s[%s]", keyPrefix, k.String()), v.Interface(), style, collectionType)
		}
		return

	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			return // Don't add nil pointers/interfaces to parameters
		}
		ParameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, v.Elem().Interface(), style, collectionType)
		return

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		value = strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		value = strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		value = strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Bool:
		value = strconv.FormatBool(v.Bool())
	case reflect.String:
		value = v.String()
	default:
		value = v.Type().String() + " value"
	}

	addValue(headerOrQueryParams, keyPrefix, value, collectionType)
}

// addValue adds a key-value pair to the header or query parameters.
//
// @param headerOrQueryParams The header or query parameters to which the value will be added.
// @param keyPrefix The key under which the value will be added.
// @param value The value to be added.
// @param collectionType The collection type (e.g., "csv") for handling multiple values.
func addValue(headerOrQueryParams interface{}, keyPrefix, value, collectionType string) {
	switch valuesMap := headerOrQueryParams.(type) {
	case url.Values:
		if collectionType == "csv" && valuesMap.Get(keyPrefix) != "" {
			valuesMap.Set(keyPrefix, valuesMap.Get(keyPrefix)+","+value)
		} else {
			valuesMap.Add(keyPrefix, value)
		}
	case map[string]string:
		valuesMap[keyPrefix] = value
	case map[string]interface{}:
		valuesMap[keyPrefix] = value
	}
}

// Decode decodes response body into the given interface v based on content type.
//
// @param v The interface to decode the response body into.
// @param b The response body as a byte slice.
// @param contentType The content type of the response body.
// @return An error if decoding fails.
func Decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if _, ok := v.(*os.File); ok {
		f, err := os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return err
		}
		_, err = f.Write(b)
		if err != nil {
			return err
		}
		_, err = f.Seek(0, io.SeekStart)
		return err
	}
	if f, ok := v.(**os.File); ok {
		*f, err = os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return err
		}
		_, err = (*f).Write(b)
		if err != nil {
			return err
		}
		_, err = (*f).Seek(0, io.SeekStart)
		return err
	}
	if XmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if JsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// ParseIntervalDetails extracts interval details from a header key.
//
// @param headerKey The header key containing interval information.
// @return An IntervalDetails struct with the extracted interval and number, or nil if parsing fails.
func ParseIntervalDetails(headerKey string) *IntervalDetails {
	r := regexp.MustCompile(`(?i)^(x-mbx-used-weight-(\d+)([smhd])|x-mbx-order-count-(\d+)([smhd]))`)
	match := r.FindStringSubmatch(headerKey)

	if match == nil {
		return nil
	}

	var numStr, letter string

	if match[2] != "" {
		numStr = match[2]
		letter = match[3]
	} else {
		numStr = match[4]
		letter = match[5]
	}

	intervalNum, err := strconv.Atoi(numStr)
	if err != nil {
		return nil
	}

	intervalMap := map[string]string{
		"S": "SECOND",
		"M": "MINUTE",
		"H": "HOUR",
		"D": "DAY",
	}

	interval, ok := intervalMap[strings.ToUpper(letter)]
	if !ok {
		return nil
	}

	return &IntervalDetails{
		Interval:    interval,
		IntervalNum: intervalNum,
	}
}

// ParseRateLimitHeaders parses rate limit information from HTTP headers.
//
// @param header The HTTP headers containing rate limit information.
// @return A slice of RateLimit structs representing the parsed rate limits or an error if parsing fails.
func ParseRateLimitHeaders(header http.Header) ([]RateLimit, error) {
	rateLimits := []RateLimit{}

	for key, values := range header {
		normalizedKey := strings.ToLower(key)
		if len(values) == 0 {
			continue
		}

		countVal, err := strconv.Atoi(values[0])
		if err != nil {
			continue
		}

		if strings.HasPrefix(normalizedKey, "x-mbx-used-weight") {
			details := ParseIntervalDetails(normalizedKey)
			if details != nil {
				rateLimits = append(rateLimits, RateLimit{
					RateLimitType: RequestWeight,
					Interval:      Interval(details.Interval),
					IntervalNum:   int32(details.IntervalNum),
					Count:         int32(countVal),
				})
			}
		} else if strings.HasPrefix(normalizedKey, "x-mbx-order-count-") {
			details := ParseIntervalDetails(normalizedKey)
			if details != nil {
				rateLimits = append(rateLimits, RateLimit{
					RateLimitType: Orders,
					Interval:      Interval(details.Interval),
					IntervalNum:   int32(details.IntervalNum),
					Count:         int32(countVal),
				})
			}
		}
	}

	if header.Get("Retry-After") != "" {
		retryStr := header.Get("Retry-After")
		retryInt, err := strconv.Atoi(retryStr)
		if err != nil {
			return nil, err
		}
		retryAfterSeconds := int32(retryInt)
		for i := range rateLimits {
			rateLimits[i].RetryAfter = retryAfterSeconds
		}
	}

	return rateLimits, nil
}

// SleepContext a sleep that can be cancelled by context.
// It uses a timer and select to either wait for the specified duration
// or return immediately if the context is cancelled. This is useful
// for operations that need to respect context cancellation while
// waiting, such as retry logic or timeouts in concurrent operations.
func SleepContext(ctx context.Context, duration time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if duration <= 0 {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return nil
		}
	}

	timer := time.NewTimer(duration)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

// SendRequest sends an HTTP request and handles retries, response decoding, and error handling.
//
// @param ctx The context for the request.
// @param path The URL path for the request.
// @param method The HTTP method (GET, POST, etc.).
// @param queryParams The query parameters for the request.
// @param cfg The configuration containing API keys, secrets, and other settings.
// @param signed A boolean indicating whether the request requires signing.
// @return The response from the REST API or an error if the request fails.
func SendRequest[T any](ctx context.Context, path string, method string, queryParams url.Values, bodyParams interface{}, cfg *ConfigurationRestAPI, signed bool) (*RestApiResponse[T], error) {
	var (
		localVarHeaderParams     = make(map[string]string)
		localVarHTTPContentTypes = []string{}
	)

	localVarHTTPContentType := SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	req, err := PrepareRequest(ctx, path, method, localVarHeaderParams, queryParams, bodyParams, cfg, signed)
	if err != nil {
		return &RestApiResponse[T]{}, err
	}
	retries := cfg.Retries
	if retries <= 0 {
		retries = 0
	}

	backoff := cfg.Backoff
	if backoff <= 0 {
		backoff = 1000
	}
	var lastErr error

	httpClient := SetupProxy(cfg)
	for attempt := 0; attempt <= retries; attempt++ {
		resp, err := httpClient.Do(req)
		if err != nil {
			lastErr = err
			if attempt < retries && ShouldRetryRequest(err, method, retries-attempt, resp) {
				if err := SleepContext(ctx, time.Duration(backoff*(attempt+1))*time.Millisecond); err != nil {
					return &RestApiResponse[T]{}, err
				}
				continue
			}
			return &RestApiResponse[T]{}, NewNetworkError(fmt.Sprintf("Network error: %v", err))
		}

		var bodyReader = resp.Body
		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			gz, err := gzip.NewReader(resp.Body)
			if err != nil {
				return &RestApiResponse[T]{}, err
			}
			defer func() {
				if err := gz.Close(); err != nil {
					fmt.Printf("Failed to close gzip reader: %v", err)
				}
			}()
			bodyReader = gz
		case "deflate":
			bodyReader = flate.NewReader(resp.Body)
			defer func() {
				if err := bodyReader.Close(); err != nil {
					fmt.Printf("Failed to close bodyReader: %v", err)
				}
			}()
		case "br":
			bodyReader = io.NopCloser(brotli.NewReader(resp.Body))
		}

		localVarBody, err := io.ReadAll(bodyReader)
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Failed to close response body: %v", err)
			return &RestApiResponse[T]{}, err
		}
		if err != nil {
			return &RestApiResponse[T]{Status: resp.StatusCode, Headers: resp.Header, RateLimits: []RateLimit{}}, err
		}

		if resp.StatusCode >= 500 && resp.StatusCode <= 504 {
			if attempt < retries {
				if err := SleepContext(ctx, time.Duration(backoff*(attempt+1))*time.Millisecond); err != nil {
					return &RestApiResponse[T]{}, err
				}
				continue
			}
			return &RestApiResponse[T]{}, fmt.Errorf("request failed after %d retries: received status %d", retries, resp.StatusCode)
		}

		if resp.StatusCode >= 300 {
			errMsg := fmt.Sprintf("HTTP error %d: %s, %s", resp.StatusCode, http.StatusText(resp.StatusCode), string(localVarBody))
			switch resp.StatusCode {
			case 400:
				return &RestApiResponse[T]{Status: resp.StatusCode, Headers: resp.Header, RateLimits: []RateLimit{}}, NewBadRequestError(errMsg)
			case 401:
				return &RestApiResponse[T]{Status: resp.StatusCode, Headers: resp.Header, RateLimits: []RateLimit{}}, NewUnauthorizedError(errMsg)
			case 403:
				return &RestApiResponse[T]{Status: resp.StatusCode, Headers: resp.Header, RateLimits: []RateLimit{}}, NewForbiddenError(errMsg)
			case 404:
				return &RestApiResponse[T]{Status: resp.StatusCode, Headers: resp.Header, RateLimits: []RateLimit{}}, NewNotFoundError(errMsg)
			case 418:
				return &RestApiResponse[T]{Status: resp.StatusCode, Headers: resp.Header, RateLimits: []RateLimit{}}, NewRateLimitBanError(errMsg)
			case 429:
				return &RestApiResponse[T]{Status: resp.StatusCode, Headers: resp.Header, RateLimits: []RateLimit{}}, NewTooManyRequestsError(errMsg)
			default:
				if resp.StatusCode >= 500 && resp.StatusCode < 600 {
					return &RestApiResponse[T]{Status: resp.StatusCode, Headers: resp.Header, RateLimits: []RateLimit{}}, NewServerError(fmt.Sprintf("Server error: %d", resp.StatusCode), resp.StatusCode)
				}
				return &RestApiResponse[T]{Status: resp.StatusCode, Headers: resp.Header, RateLimits: []RateLimit{}}, NewConnectorClientError(fmt.Sprintf("Error connecting to the API (%d)", resp.StatusCode))
			}
		}
		var data T
		err = Decode(&data, localVarBody, resp.Header.Get("Content-Type"))
		if err != nil {
			return &RestApiResponse[T]{Status: resp.StatusCode, Headers: resp.Header, RateLimits: []RateLimit{}}, err
		}

		rateLimits, err := ParseRateLimitHeaders(resp.Header)
		if err != nil {
			return &RestApiResponse[T]{Status: resp.StatusCode, Headers: resp.Header, RateLimits: []RateLimit{}}, err
		}

		return &RestApiResponse[T]{Data: data, Status: resp.StatusCode, Headers: resp.Header, RateLimits: rateLimits}, nil
	}

	return &RestApiResponse[T]{}, fmt.Errorf("request failed after %d retries: %v", retries, lastErr)
}

// PrepareRequest prepares an HTTP request with the given parameters.
//
// @param ctx The context for the request.
// @param path The URL path for the request.
// @param method The HTTP method (GET, POST, etc.).
// @param headerParams A map of header parameters to include in the request.
// @param queryParams The query parameters for the request.
// @param c The configuration containing API keys, secrets, and other settings.
// @param signed A boolean indicating whether the request requires signing.
// @return The prepared HTTP request or an error if preparation fails.
func PrepareRequest(
	ctx context.Context,
	path string, method string,
	headerParams map[string]string,
	queryParams url.Values,
	bodyParams interface{},
	c *ConfigurationRestAPI,
	signed bool) (localVarRequest *http.Request, err error) {

	reqURL, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	query := reqURL.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	reqURL.RawQuery = queryParamSplit.ReplaceAllStringFunc(query.Encode(), func(s string) string {
		pieces := strings.Split(s, "=")
		pieces[0] = queryDescape.Replace(pieces[0])
		if len(pieces) > 1 {
			pieces[1] = strings.TrimPrefix(pieces[1], "INTERVAL_")
		}
		return strings.Join(pieces, "=")
	})

	paramsToSign := reqURL.RawQuery

	if bodyParams != nil {
		var bodyString string
		switch v := bodyParams.(type) {
		case url.Values:
			bodyString = v.Encode()
		case map[string]string:
			values := url.Values{}
			for k, val := range v {
				values.Add(k, val)
			}
			bodyString = values.Encode()
		case map[string]interface{}:
			values := url.Values{}
			for k, val := range v {
				values.Add(k, fmt.Sprintf("%v", val))
			}
			bodyString = values.Encode()
		default:
			jsonBody, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			bodyString = string(jsonBody)
		}

		if bodyString != "" {
			if paramsToSign != "" {
				paramsToSign += "&"
			}
			paramsToSign += bodyString
		}
	}

	if signed && c != nil {
		if c.ApiSecret != "" {
			paramsToSign += "&timestamp=" + strconv.FormatInt(time.Now().UnixMilli(), 10)

			signer := hmac.New(sha256.New, []byte(c.ApiSecret))
			signer.Write([]byte(paramsToSign))
			signature := signer.Sum(nil)
			if err != nil {
				return nil, err
			}
			paramsToSign += "&signature=" + fmt.Sprintf("%x", signature)
		}

		if c.PrivateKey != "" {
			paramsToSign += "&timestamp=" + strconv.FormatInt(time.Now().UnixMilli(), 10)

			if c.Signer == nil {
				key, err := LoadPrivateKey(c.PrivateKey, c.PrivateKeyPassphrase)
				if err != nil {
					panic(err)
				}
				c.Signer = &cryptoSigner{s: key}
			}

			signature, _ := c.Signer.Sign([]byte(paramsToSign))
			if err != nil {
				panic(err)
			}
			paramsToSign += "&signature=" + base64.StdEncoding.EncodeToString(signature)
		}
	}
	reqURL.RawQuery = paramsToSign

	req, err := http.NewRequest(method, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	for h, v := range headerParams {
		req.Header.Set(h, v)
	}

	if c != nil && c.ApiKey != "" {
		req.Header.Set("X-MBX-APIKEY", c.ApiKey)
	}

	if c != nil {
		if c.TimeUnit != "" {
			req.Header.Add("X-MBX-TIME-UNIT", string(c.TimeUnit))
		}
		if c.Compression {
			req.Header.Add("Accept-Encoding", "gzip, deflate, br")
		}
		if !c.KeepAlive {
			req.Header.Add("Connection", "close")
		}
		for h, v := range c.CustomHeaders {
			req.Header.Set(h, v)
		}
	}

	if bodyParams != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// SetupProxy sets up the HTTP client with proxy settings if provided in the configuration.
//
// @param cfg The configuration containing proxy settings.
// @return An HTTP client configured with the proxy settings if provided, otherwise a default HTTP client.
func SetupProxy(cfg *ConfigurationRestAPI) *http.Client {
	if cfg.Proxy != nil && !cfg.Proxy.IsEmpty() {
		proxyURL, err := cfg.Proxy.URL()

		if err != nil {
			panic(err)
		}

		return &http.Client{
			Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL), TLSClientConfig: cfg.Proxy.TLSConfig},
		}
	}

	if cfg.HTTPSAgent != nil {
		transport := BuildTransport(cfg.HTTPSAgent, cfg)
		return &http.Client{
			Transport: transport,
			Timeout:   cfg.Timeout,
		}
	}
	return &http.Client{
		Timeout: cfg.Timeout,
	}
}

// Sign signs a message using the provided crypto.Signer.
//
// @param data The message to be signed.
// @return The signature as a byte slice and an error if any occurs during signing.
func (s *cryptoSigner) Sign(data []byte) ([]byte, error) {
	return SignMessage(s.s, data)
}

// NormalizePrivateKeyInput normalizes the private key input by checking if it's a file path or a PEM string.
//
// @param input The private key input, which can be a file path or a PEM string.
// @return The normalized private key as a string and an error if any occurs during reading the file.
func NormalizePrivateKeyInput(input string) (string, error) {
	if _, err := os.Stat(input); err == nil {
		data, readErr := os.ReadFile(input)
		if readErr != nil {
			return "", readErr
		}
		return string(data), nil
	}

	return input, nil
}

// ParsePrivateKey parses a PEM encoded private key, which may be encrypted.
//
// @param pemData The PEM encoded private key data.
// @param password The password used to decrypt the private key, if it is encrypted.
// @return A crypto.Signer representing the private key and an error if parsing fails.
func ParsePrivateKey(pemData string, password []byte) (crypto.Signer, error) {
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	var key interface{}
	var err error

	switch block.Type {
	case "ENCRYPTED PRIVATE KEY":
		key, err = pkcs8.ParsePKCS8PrivateKey(block.Bytes, password)
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt PKCS#8 private key: %w", err)
		}

	case "PRIVATE KEY":
		key, err = x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PKCS#8 private key: %w", err)
		}

	case "RSA PRIVATE KEY":
		key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse PKCS#1 private key: %w", err)
		}

	case "EC PRIVATE KEY":
		key, err = x509.ParseECPrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse EC private key: %w", err)
		}

	default:
		return nil, fmt.Errorf("unsupported key type: %s", block.Type)
	}

	if err != nil {
		return nil, err
	}

	switch k := key.(type) {
	case ed25519.PrivateKey:
		return k, nil
	case *rsa.PrivateKey:
		return k, nil
	case *ecdsa.PrivateKey:
		return k, nil
	default:
		return nil, fmt.Errorf("unsupported private key type %T", k)
	}
}

// SignMessage signs a message using the provided private key.
//
// @param priv The private key used for signing.
// @param message The message to be signed.
// @return The signature as a byte slice and an error if any occurs during signing.
func SignMessage(priv crypto.Signer, message []byte) ([]byte, error) {
	switch k := priv.(type) {
	case ed25519.PrivateKey:
		return ed25519.Sign(k, message), nil

	case *rsa.PrivateKey:
		hash := sha256.Sum256(message)
		return rsa.SignPKCS1v15(rand.Reader, k, crypto.SHA256, hash[:])

	default:
		return nil, fmt.Errorf("unsupported key type %T for signing", k)
	}
}

// ShouldRetryRequest determines whether a request should be retried based on the error, HTTP method, retries left, and response.
//
// @param err The error encountered during the request.
// @param method The HTTP method used for the request.
// @param retriesLeft The number of retries left.
// @param resp The HTTP response received (if any).
// @return A boolean indicating whether the request should be retried.
var ShouldRetryRequest = func(err error, method string, retriesLeft int, resp *http.Response) bool {
	if retriesLeft <= 0 {
		return false
	}

	retriableMethods := map[string]bool{
		"GET":    true,
		"DELETE": true,
	}
	if _, ok := retriableMethods[strings.ToUpper(method)]; !ok {
		return false
	}

	if resp == nil {
		return true
	}

	retriableStatusCodes := map[int]bool{
		500: true,
		502: true,
		503: true,
		504: true,
	}

	return retriableStatusCodes[resp.StatusCode]
}

// PtrBool is a helper routine that returns a pointer to given boolean value.
//
// @param v The bool value to be pointed to.
// @return A pointer to the provided bool value.
func PtrBool(v bool) *bool { return &v }

// PtrInt is a helper routine that returns a pointer to given integer value.
//
// @param v The int value to be pointed to.
// @return A pointer to the provided int value.
func PtrInt(v int) *int { return &v }

// PtrInt32 is a helper routine that returns a pointer to given integer value.
//
// @param v The int32 value to be pointed to.
// @return A pointer to the provided int32 value.
func PtrInt32(v int32) *int32 { return &v }

// PtrInt64 is a helper routine that returns a pointer to given integer value.
//
// @param v The int64 value to be pointed to.
// @return A pointer to the provided int64 value.
func PtrInt64(v int64) *int64 { return &v }

// PtrFloat32 is a helper routine that returns a pointer to given float value.
//
// @param v The float32 value to be pointed to.
// @return A pointer to the provided float32 value.
func PtrFloat32(v float32) *float32 { return &v }

// PtrFloat64 is a helper routine that returns a pointer to given float value.
//
// @param v The float64 value to be pointed to.
// @return A pointer to the provided float64 value.
func PtrFloat64(v float64) *float64 { return &v }

// PtrString is a helper routine that returns a pointer to given string value.
//
// @param v The string value to be pointed to.
// @return A pointer to the provided string value.
func PtrString(v string) *string { return &v }

// PtrTime is helper routine that returns a pointer to given Time value.
//
// @param v The time.Time value to be pointed to.
// @return A pointer to the provided time.Time value.
func PtrTime(v time.Time) *time.Time { return &v }

// RandomString generates a random hexadecimal string of length 32 (16 bytes).
//
// @return A random hexadecimal string and an error if any occurs during generation.
func RandomString() (string, error) {
	bytes := make([]byte, 16) // 16 random bytes
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// GetTimestamp returns the current timestamp in milliseconds.
//
// @return The current timestamp in milliseconds as an int64.
func GetTimestamp() int64 {
	return time.Now().UnixMilli()
}

// Urlencode encodes a map of parameters into a URL-encoded query string.
//
// @param params A map of parameter names to their corresponding values.
// @return A URL-encoded query string representing the parameters.
func Urlencode(params map[string]any) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	values := url.Values{}
	for _, k := range keys {
		values.Set(k, fmt.Sprintf("%v", params[k]))
	}
	return values.Encode()
}

// FilterArrays filters elements from a slice based on a provided condition function.
//
// @param slice The input slice of type T.
// @param keep A function that determines whether to keep an element (returns true) or discard it (returns false).
// @return A new slice containing only the elements for which the keep function returned true.
func FilterArrays[T any](slice []T, keep func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if keep(v) {
			result = append(result, v)
		}
	}
	return result
}

// WsStreamsPlaceholder replaces placeholders in a WebSocket stream string with actual parameter values.
//
// @param stream The WebSocket stream string containing placeholders.
// @param params A map of parameter names to their corresponding values.
// @return The WebSocket stream string with placeholders replaced by actual values.
func WsStreamsPlaceholder(stream string, params map[string]string) string {
	cleanParams := make(map[string]string)
	for k, v := range params {
		if v != "" {
			cleanParams[k] = v
		}
	}

	normalizedVariables := make(map[string]string)
	for key, value := range cleanParams {
		nf := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(key, "-", ""), "_", ""))
		normalizedVariables[nf] = value
	}

	re := regexp.MustCompile(`([@_]?)<([^>]+)>`)

	result := re.ReplaceAllStringFunc(stream, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		prefix := submatches[1]
		rawVar := submatches[2]

		nf := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(rawVar, "-", ""), "_", ""))

		value, ok := normalizedVariables[nf]
		if !ok || value == "" {
			return ""
		}

		normalizedValue := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(value, "-", ""), "_", ""))
		if strings.HasPrefix(normalizedValue, nf) {
			value = value[len(value)-len(normalizedValue)+len(nf):]
			value = strings.TrimLeft(value, "_")
		}

		switch nf {
		case "symbol", "windowsize":
			if prefix == "_" {
				return "_" + strings.ToLower(value)
			}
			return strings.ToLower(value)
		case "updatespeed":
			return "@" + value
		default:
			return prefix + value
		}
	})

	return result
}

// GenerateUUID is a function that generates a new UUID string.
//
// @return A string representing the generated UUID.
var GenerateUUID = func() string {
	return uuid.New().String()
}

// GenerateIntUUID is a function that generates a new random int32 UUID.
//
// @return An int32 representing the generated UUID.
var GenerateIntUUID = func() int32 {
	return randpkg.Int32()
}

// Pretty converts a value to its pretty-printed JSON string representation.
//
// @param v The value to be pretty-printed.
// @return A string containing the pretty-printed JSON representation of the value.
func Pretty(v any) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}

// httpProxyDialer implements the proxy.Dialer interface for HTTP proxies.
//
// @param proxyAddr The address of the HTTP proxy.
// @return An instance of httpProxyDialer and an error if processing fails, otherwise nil.
func createHTTPProxyDialer(proxyURL *url.URL) (proxy.Dialer, error) {
	auth := &proxy.Auth{}
	if proxyURL.User != nil {
		auth.User = proxyURL.User.Username()
		if password, ok := proxyURL.User.Password(); ok {
			auth.Password = password
		}
	}

	return &httpProxyDialer{
		proxyAddr: proxyURL.Host,
		auth:      auth,
	}, nil
}

// Dial connects to the address addr via the HTTP proxy.
//
// @param network The network type (e.g., "tcp").
// @param addr The target address to connect to.
// @return A net.Conn representing the established connection, or an error if the connection fails.
func (d *httpProxyDialer) Dial(network, addr string) (net.Conn, error) {
	conn, err := net.Dial(network, d.proxyAddr)
	if err != nil {
		return nil, err
	}

	req := &http.Request{
		Method: "CONNECT",
		URL:    &url.URL{Host: addr},
		Host:   addr,
		Header: make(http.Header),
	}

	if d.auth != nil && d.auth.User != "" {
		req.Header.Set("Proxy-Authorization",
			"Basic "+base64.StdEncoding.EncodeToString(
				[]byte(d.auth.User+":"+d.auth.Password)))
	}

	if err := req.Write(conn); err != nil {
		if err := conn.Close(); err != nil {
			fmt.Printf("Failed to close connection: %v", err)
		}
		return nil, err
	}

	resp, err := http.ReadResponse(bufio.NewReader(conn), req)
	if err != nil {
		if err := conn.Close(); err != nil {
			fmt.Printf("Failed to close connection: %v", err)
		}
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		if err := conn.Close(); err != nil {
			fmt.Printf("Failed to close connection: %v", err)
		}
		return nil, fmt.Errorf("proxy error: %s", resp.Status)
	}

	return conn, nil
}

// processParams processes the input parameters for a WebSocket message.
//
// @param params The input parameters as a map.
// @param withAPIKey A boolean indicating whether to include the API key.
// @param skipAuth A boolean indicating whether to skip authentication.
// @param apiKey The API key to include if withAPIKey is true and skipAuth is false.
// @return A new map of processed parameters or an error if processing fails.
func processParams(
	params map[string]any,
	withAPIKey bool,
	skipAuth bool,
	apiKey string,
) (map[string]any, error) {
	newParams := make(map[string]any)

	for k, v := range params {
		if k == "id" {
			continue
		}

		if s, ok := v.(string); ok {
			v = strings.TrimPrefix(s, "INTERVAL_")
		}

		valType := reflect.TypeOf(v).Kind()
		if valType == reflect.Slice || valType == reflect.Map {
			bytes, err := json.Marshal(v)
			if err != nil {
				return nil, &WebSocketError{
					Op:      "marshal_params",
					Err:     errors.New("failed to marshal params in SendMessage"),
					Message: "failed to marshal params for key " + k,
				}
			}
			newParams[k] = string(bytes)
		} else {
			if k == "type_" {
				newParams["type"] = v
				continue
			} else {
				newParams[k] = v
			}
		}
	}

	if withAPIKey && !skipAuth {
		newParams["apiKey"] = apiKey
	}

	return newParams, nil
}

// LoadPrivateKey loads a private key from a PEM-formatted string or file path,
//
// @param privateKey The PEM-formatted private key string or file path.
// @param passphrase The passphrase used to decrypt the private key, if it is encrypted.
// @return A crypto.Signer representing the loaded private key, or an error if loading fails.
func LoadPrivateKey(privateKey string, passphrase string) (crypto.Signer, error) {
	pemStr, err := NormalizePrivateKeyInput(privateKey)
	if err != nil {
		return nil, err
	}

	priv, err := ParsePrivateKey(pemStr, []byte(passphrase))
	if err != nil {
		return nil, err
	}
	return priv, nil
}
