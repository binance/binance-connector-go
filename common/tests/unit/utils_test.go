package commontests

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/binance/binance-connector-go/common/common"
)

// Mock type implementing MappedNullable
type mockMappedNullable struct {
	Data map[string]interface{}
	Err  error
}

func (m mockMappedNullable) ToMap() (map[string]interface{}, error) {
	return m.Data, m.Err
}

func TestParameterAddToHeaderOrQuery_BasicTypes(t *testing.T) {
	tests := []struct {
		name		   string
		obj			interface{}
		expectedValue  string
		expectedKey	string
		collectionType string
	}{
		{"int", 42, "42", "intKey", ""},
		{"uint", uint(99), "99", "uintKey", ""},
		{"float", 3.14, "3.14", "floatKey", ""},
		{"bool", true, "true", "boolKey", ""},
		{"string", "hello", "hello", "stringKey", ""},
		{"nil", nil, "null", "nilKey", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := url.Values{}
			common.ParameterAddToHeaderOrQuery(values, tt.expectedKey, tt.obj, "form", tt.collectionType)
			if got := values.Get(tt.expectedKey); got != tt.expectedValue {
				t.Errorf("Expected %q, got %q", tt.expectedValue, got)
			}
		})
	}
}

func TestParameterAddToHeaderOrQuery_StructMappedNullable(t *testing.T) {
	mock := mockMappedNullable{
		Data: map[string]interface{}{
			"field1": "value1",
			"field2": 123,
		},
	}

	values := url.Values{}
	common.ParameterAddToHeaderOrQuery(values, "obj", mock, "form", "")
	got := values.Get("obj")

	var result struct {
		Data map[string]interface{} `json:"Data"`
	}
	if err := json.Unmarshal([]byte(got), &result); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if result.Data["field1"] != "value1" {
		t.Errorf("Expected field1=value1, got %v", result.Data["field1"])
	}
	if result.Data["field2"] != float64(123) { // JSON numbers become float64
		t.Errorf("Expected field2=123, got %v", result.Data["field2"])
	}
}

func TestParameterAddToHeaderOrQuery_StructTime(t *testing.T) {
	now := time.Date(2023, 3, 10, 12, 0, 0, 0, time.UTC)
	values := url.Values{}
	common.ParameterAddToHeaderOrQuery(values, "timestamp", now, "form", "")

	expected := `"2023-03-10T12:00:00Z"`
	if values.Get("timestamp") != expected {
		t.Errorf("Expected %q, got %q", expected, values.Get("timestamp"))
	}
}

func TestParameterAddToHeaderOrQuery_Slice(t *testing.T) {
	values := url.Values{}
	arr := []int{1, 2, 3}
	common.ParameterAddToHeaderOrQuery(values, "numbers", arr, "form", "")

	expected := `[1,2,3]`
	got := values.Get("numbers")
	if got != expected {
		t.Errorf("Expected %q, got %q", expected, got)
	}
}

func TestParameterAddToHeaderOrQuery_SliceDeepObject(t *testing.T) {
	values := url.Values{}
	arr := []string{"a", "b"}
	common.ParameterAddToHeaderOrQuery(values, "letters", arr, "deepObject", "")

	expected := `["a","b"]`
	got := values.Get("letters")
	if got != expected {
		t.Errorf("Expected %q, got %q", expected, got)
	}
}

func TestParameterAddToHeaderOrQuery_Map(t *testing.T) {
	values := url.Values{}
	m := map[string]interface{}{
		"a": 1,
		"b": "two",
	}
	common.ParameterAddToHeaderOrQuery(values, "mapKey", m, "form", "")
	if values.Get("mapKey[a]") != "1" || values.Get("mapKey[b]") != "two" {
		t.Errorf("Unexpected map: %v", values)
	}
}

func TestParameterAddToHeaderOrQuery_PointerAndInterface(t *testing.T) {
	str := "hello"
	ptr := &str

	values := url.Values{}
	common.ParameterAddToHeaderOrQuery(values, "ptrKey", ptr, "form", "")
	if values.Get("ptrKey") != "hello" {
		t.Errorf("Expected ptrKey=hello, got %q", values.Get("ptrKey"))
	}

	var iface interface{} = 123
	values2 := url.Values{}
	common.ParameterAddToHeaderOrQuery(values2, "ifaceKey", iface, "form", "")
	if values2.Get("ifaceKey") != "123" {
		t.Errorf("Expected ifaceKey=123, got %q", values2.Get("ifaceKey"))
	}
}

func TestParameterAddToHeaderOrQuery_CSVCollection(t *testing.T) {
	values := url.Values{}
	common.ParameterAddToHeaderOrQuery(values, "list", "first", "form", "csv")
	common.ParameterAddToHeaderOrQuery(values, "list", "second", "form", "csv")

	if got := values.Get("list"); got != "first,second" {
		t.Errorf("Expected list=first,second got %q", got)
	}
}

func TestParameterAddToHeaderOrQuery_MapStringString(t *testing.T) {
	header := map[string]string{}
	common.ParameterAddToHeaderOrQuery(header, "X-Key", 42, "form", "")
	if header["X-Key"] != "42" {
		t.Errorf("Expected header[X-Key]=42, got %v", header["X-Key"])
	}
}

func TestParameterAddToHeaderOrQuery_InvalidTypes(t *testing.T) {
	values := url.Values{}
	ch := make(chan int)
	common.ParameterAddToHeaderOrQuery(values, "chan", ch, "form", "")
	if got := values.Get("chan"); got == "" {
		t.Errorf("Expected a type string, got empty")
	}
}

func TestParameterAddToHeaderOrQuery_NilPointer(t *testing.T) {
	var ptr *int
	values := url.Values{}
	common.ParameterAddToHeaderOrQuery(values, "nilPtr", ptr, "form", "")
	if values.Get("nilPtr") != "" {
		t.Errorf("Expected no value, got %v", values)
	}
}

func TestDecode_String(t *testing.T) {
	var s string
	err := common.Decode(&s, []byte("hello"), "text/plain")
	if err != nil {
		t.Fatal(err)
	}
	if s != "hello" {
		t.Errorf("Expected 'hello', got %q", s)
	}
}

func TestDecode_FilePointer(t *testing.T) {
	var f *os.File
	data := []byte("file content")
	err := common.Decode(&f, data, "application/octet-stream")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			t.Errorf("Failed to close file: %v", err)
		}
	}()
	content, _ := io.ReadAll(f)
	if string(content) != string(data) {
		t.Errorf("Expected %q, got %q", string(data), string(content))
	}
}

func TestDecode_ByteNil(t *testing.T) {
	var s string
	err := common.Decode(&s, []byte{}, "text/plain")
	if err != nil {
		t.Fatal(err)
	}
	if s != "" {
		t.Errorf("Expected empty string, got %q", s)
	}
}

type SampleJSON struct {
	Name string `json:"name"`
}

func TestDecode_JSON(t *testing.T) {
	obj := SampleJSON{}
	data := []byte(`{"name":"Alice"}`)
	err := common.Decode(&obj, data, "application/json")
	if err != nil {
		t.Fatal(err)
	}
	if obj.Name != "Alice" {
		t.Errorf("Expected Alice, got %s", obj.Name)
	}
}

type SampleXML struct {
	XMLName xml.Name `xml:"sample"`
	Value   string   `xml:"value"`
}

func TestDecode_XML(t *testing.T) {
	obj := SampleXML{}
	data := []byte(`<sample><value>42</value></sample>`)
	err := common.Decode(&obj, data, "application/xml")
	if err != nil {
		t.Fatal(err)
	}
	if obj.Value != "42" {
		t.Errorf("Expected 42, got %s", obj.Value)
	}
}

func TestDecode_UnknownContentType(t *testing.T) {
	err := common.Decode(&struct{}{}, []byte("data"), "unknown/type")
	if err == nil || !strings.Contains(err.Error(), "undefined response type") {
		t.Errorf("Expected undefined response type error, got %v", err)
	}
}

func TestParseIntervalDetails_ValidXMBXUsedWeight(t *testing.T) {
	h := "x-mbx-used-weight-5m"
	res := common.ParseIntervalDetails(h)
	if res == nil || res.Interval != "MINUTE" || res.IntervalNum != 5 {
		t.Errorf("Unexpected result: %+v", res)
	}
}

func TestParseIntervalDetails_ValidXMBXOrderCount(t *testing.T) {
	h := "x-mbx-order-count-10H"
	res := common.ParseIntervalDetails(h)
	if res == nil || res.Interval != "HOUR" || res.IntervalNum != 10 {
		t.Errorf("Unexpected result: %+v", res)
	}
}

func TestParseIntervalDetails_InvalidHeader(t *testing.T) {
	res := common.ParseIntervalDetails("invalid-header")
	if res != nil {
		t.Errorf("Expected nil, got %+v", res)
	}
}

func TestParseIntervalDetails_InvalidNumber(t *testing.T) {
	res := common.ParseIntervalDetails("x-mbx-used-weight-abcM")
	if res != nil {
		t.Errorf("Expected nil, got %+v", res)
	}
}

func TestParseIntervalDetails_InvalidLetter(t *testing.T) {
	res := common.ParseIntervalDetails("x-mbx-order-count-5Z")
	if res != nil {
		t.Errorf("Expected nil, got %+v", res)
	}
}

func TestParseRateLimitHeaders_RequestWeight(t *testing.T) {
	h := http.Header{}
	h.Set("X-MBX-USED-WEIGHT-1m", "100")
	rates, err := common.ParseRateLimitHeaders(h)
	if err != nil {
		t.Fatal(err)
	}
	if len(rates) != 1 {
		t.Fatalf("Expected 1 rate limit, got %d", len(rates))
	}
	r := rates[0]
	if r.RateLimitType != common.RequestWeight || r.Interval != "MINUTE" || r.IntervalNum != 1 || r.Count != 100 {
		t.Errorf("Unexpected rate limit: %+v", r)
	}
}

func TestParseRateLimitHeaders_OrderCount(t *testing.T) {
	h := http.Header{}
	h.Set("X-MBX-ORDER-COUNT-2H", "50")
	rates, err := common.ParseRateLimitHeaders(h)
	if err != nil {
		t.Fatal(err)
	}
	if len(rates) != 1 {
		t.Fatalf("Expected 1 rate limit, got %d", len(rates))
	}
	r := rates[0]
	if r.RateLimitType != common.Orders || r.Interval != "HOUR" || r.IntervalNum != 2 || r.Count != 50 {
		t.Errorf("Unexpected rate limit: %+v", r)
	}
}

func TestParseRateLimitHeaders_RetryAfter(t *testing.T) {
	h := http.Header{}
	h.Set("X-MBX-USED-WEIGHT-1s", "10")
	h.Set("Retry-After", "5")
	rates, err := common.ParseRateLimitHeaders(h)
	if err != nil {
		t.Fatal(err)
	}
	if rates[0].RetryAfter != 5 {
		t.Errorf("Expected RetryAfter=5, got %d", rates[0].RetryAfter)
	}
}

func TestParseRateLimitHeaders_InvalidCount(t *testing.T) {
	h := http.Header{}
	h.Set("X-MBX-USED-WEIGHT-1m", "abc") // invalid number
	rates, err := common.ParseRateLimitHeaders(h)
	if err != nil {
		t.Fatal(err)
	}
	if len(rates) != 0 {
		t.Errorf("Expected 0 rate limits, got %d", len(rates))
	}
}

func TestParseRateLimitHeaders_EmptyHeader(t *testing.T) {
	h := http.Header{}
	rates, err := common.ParseRateLimitHeaders(h)
	if err != nil {
		t.Fatal(err)
	}
	if len(rates) != 0 {
		t.Errorf("Expected 0 rate limits, got %d", len(rates))
	}
}

type SampleResponse struct {
	Message string `json:"message"`
}

var originalShouldRetry func(error, string, int, *http.Response) bool

func mockShouldRetry(err error, method string, remainingRetries int, resp *http.Response) bool {
	return true
}

func TestSendRequest_Success(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		if _, err := fmt.Fprintln(w, `{"message":"ok"}`); err != nil {
			t.Fatal(err)
		}
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	cfg := &common.ConfigurationRestAPI{}

	resp, err := common.SendRequest[SampleResponse](context.Background(), server.URL, "GET", url.Values{}, nil, cfg)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != 200 || resp.Data.Message != "ok" {
		t.Errorf("Unexpected response: %+v", resp)
	}
}

func TestSendRequest_RetryLogic(t *testing.T) {
	attempts := 0
	handler := func(w http.ResponseWriter, r *http.Request) {
		attempts++
		w.WriteHeader(500)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	cfg := &common.ConfigurationRestAPI{
		Retries: 2,
		Backoff: 0,
	}

	// Mock retry logic
	originalShouldRetry = common.ShouldRetryRequest
	common.ShouldRetryRequest = mockShouldRetry
	defer func() { common.ShouldRetryRequest = originalShouldRetry }()

	resp, err := common.SendRequest[SampleResponse](context.Background(), server.URL, "GET", url.Values{}, nil, cfg)
	if err == nil {
		t.Fatalf("Expected error for server failure")
	}

	if attempts != 3 {
		t.Errorf("Expected 3 attempts (1 initial + 2 retries), got %d", attempts)
	}

	if resp.Status != 0 {
		t.Errorf("Expected status 0 on failure, got %d", resp.Status)
	}
}

func TestSendRequest_HTTPErrorStatus(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	cfg := &common.ConfigurationRestAPI{}

	resp, err := common.SendRequest[SampleResponse](context.Background(), server.URL, "GET", url.Values{}, nil, cfg)
	if err == nil {
		t.Fatalf("Expected error for 404")
	}

	if resp.Status != 404 {
		t.Errorf("Expected status 404, got %d", resp.Status)
	}
}

func TestSendRequest_ContentEncodingGzip(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)

		var buf bytes.Buffer
		gz := gzip.NewWriter(&buf)
		_, _ = gz.Write([]byte(`{"message":"compressed"}`))
		if err := gz.Close(); err != nil {
			t.Fatal(err)
		}

		if _, err := w.Write(buf.Bytes()); err != nil {
			t.Fatal(err)
		}
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	cfg := &common.ConfigurationRestAPI{}

	resp, err := common.SendRequest[SampleResponse](context.Background(), server.URL, "GET", url.Values{}, nil, cfg)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Data.Message != "compressed" {
		t.Errorf("Expected 'compressed', got %q", resp.Data.Message)
	}
}

func TestPrepareRequest_BasicHeadersAndQuery(t *testing.T) {
	cfg := &common.ConfigurationRestAPI{
		ApiKey: "apikey123",
	}
	query := url.Values{}
	query.Set("param", "value")

	req, err := common.PrepareRequest(context.Background(), "https://example.com/test", "GET", map[string]string{}, query, nil, cfg)
	if err != nil {
		t.Fatal(err)
	}

	if req.Header.Get("X-MBX-APIKEY") != "apikey123" {
		t.Errorf("Expected API key header set")
	}

	if req.URL.Query().Get("param") != "value" {
		t.Errorf("Expected query param preserved")
	}
}

func TestPrepareRequest_CustomHeadersAndCompression(t *testing.T) {
	cfg := &common.ConfigurationRestAPI{
		CustomHeaders: map[string]string{"X-Custom": "abc"},
		Compression:   true,
	}
	req, err := common.PrepareRequest(context.Background(), "https://example.com", "GET", map[string]string{}, url.Values{}, nil, cfg)
	if err != nil {
		t.Fatal(err)
	}

	if req.Header.Get("X-Custom") != "abc" {
		t.Errorf("Expected custom header set")
	}

	if req.Header.Get("Accept-Encoding") != "gzip, deflate, br" {
		t.Errorf("Expected Accept-Encoding header for compression")
	}
}

func TestPrepareRequest_WithURLValuesBody(t *testing.T) {
	cfg := &common.ConfigurationRestAPI{}
	body := url.Values{}
	body.Set("key1", "value1")
	body.Set("key2", "value2")

	req, err := common.PrepareRequest(context.Background(), "https://example.com", "POST",
		map[string]string{}, url.Values{}, body, cfg)
	if err != nil {
		t.Fatal(err)
	}

	if req.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		t.Errorf("Expected Content-Type header to be set")
	}

	// Verify the body is included in paramsToSign
	// We can't directly check paramsToSign as it's not returned, but we can verify the behavior
	// by checking that the request would be properly signed if ApiSecret was set
}

func TestPrepareRequest_WithMapStringStringBody(t *testing.T) {
	cfg := &common.ConfigurationRestAPI{}
	body := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	req, err := common.PrepareRequest(context.Background(), "https://example.com", "POST",
		map[string]string{}, url.Values{}, body, cfg)
	if err != nil {
		t.Fatal(err)
	}

	if req.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		t.Errorf("Expected Content-Type header to be set")
	}
}

func TestPrepareRequest_WithMapStringInterfaceBody(t *testing.T) {
	cfg := &common.ConfigurationRestAPI{}
	body := map[string]interface{}{
		"key1": "value1",
		"key2": 123,
		"key3": true,
	}

	req, err := common.PrepareRequest(context.Background(), "https://example.com", "POST",
		map[string]string{}, url.Values{}, body, cfg)
	if err != nil {
		t.Fatal(err)
	}

	if req.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		t.Errorf("Expected Content-Type header to be set")
	}
}

func TestPrepareRequest_WithJSONBody(t *testing.T) {
	cfg := &common.ConfigurationRestAPI{}
	body := struct {
		Key1 string `json:"key1"`
		Key2 int	`json:"key2"`
	}{
		Key1: "value1",
		Key2: 123,
	}

	req, err := common.PrepareRequest(context.Background(), "https://example.com", "POST",
		map[string]string{}, url.Values{}, body, cfg)
	if err != nil {
		t.Fatal(err)
	}

	if req.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		t.Errorf("Expected Content-Type header to be set")
	}
}

func TestPrepareRequest_WithNilBody(t *testing.T) {
	cfg := &common.ConfigurationRestAPI{}

	req, err := common.PrepareRequest(context.Background(), "https://example.com", "GET",
		map[string]string{}, url.Values{}, nil, cfg)
	if err != nil {
		t.Fatal(err)
	}

	if req.Header.Get("Content-Type") != "" {
		t.Errorf("Expected no Content-Type header when body is nil")
	}
}

func TestPrepareRequest_WithQueryAndBody(t *testing.T) {
	cfg := &common.ConfigurationRestAPI{}
	query := url.Values{}
	query.Set("queryParam", "queryValue")
	body := url.Values{}
	body.Set("bodyParam", "bodyValue")

	req, err := common.PrepareRequest(context.Background(), "https://example.com", "POST",
		map[string]string{}, query, body, cfg)
	if err != nil {
		t.Fatal(err)
	}

	if req.URL.Query().Get("queryParam") != "queryValue" {
		t.Errorf("Expected query param preserved")
	}

	if req.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		t.Errorf("Expected Content-Type header to be set")
	}
}

func TestGenerateUUID(t *testing.T) {
	t.Run("returns non-empty string", func(t *testing.T) {
		uuid := common.GenerateUUID()
		if len(uuid) == 0 {
			t.Error("GenerateUUID() returned empty string")
		}
	})

	t.Run("returns valid UTF-8 string", func(t *testing.T) {
		uuid := common.GenerateUUID()
		if !utf8.ValidString(uuid) {
			t.Error("GenerateUUID() returned invalid UTF-8 string")
		}
	})

	t.Run("returns string of expected length", func(t *testing.T) {
		uuid := common.GenerateUUID()

		if len(uuid) != 36 {
			t.Errorf("GenerateUUID() returned string of length %d, expected 36", len(uuid))
		}
	})

	t.Run("returns different values on subsequent calls", func(t *testing.T) {
		first := common.GenerateUUID()
		second := common.GenerateUUID()
		if first == second {
			t.Error("GenerateUUID() returned same value on subsequent calls")
		}
	})
}

func TestGenerateIntUUID(t *testing.T) {
	t.Run("returns a value", func(t *testing.T) {
		uuid := common.GenerateIntUUID()
		_ = uuid
	})

	t.Run("returns different values on subsequent calls", func(t *testing.T) {
		first := common.GenerateIntUUID()
		second := common.GenerateIntUUID()
		if first == second {
			t.Error("GenerateIntUUID() returned same value on subsequent calls")
		}
	})

	t.Run("returns value within int32 range", func(t *testing.T) {
		uuid := common.GenerateIntUUID()

		if uuid < -2147483648 || uuid > 2147483647 {
			t.Errorf("GenerateIntUUID() returned value %d outside int32 range", uuid)
		}
	})
}
