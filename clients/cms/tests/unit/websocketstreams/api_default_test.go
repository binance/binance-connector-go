package binancecmswebsocketstreams

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"time"

	client "github.com/binance/binance-connector-go/clients/cms"
	"github.com/binance/binance-connector-go/clients/cms/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
	tests "github.com/binance/binance-connector-go/common/v2/tests"
	"github.com/stretchr/testify/require"
)

func Test_binancecmswebsocketstreams_DefaultAPIService(t *testing.T) {

	t.Run("Test DefaultAPI ComAnnouncementEn Success", func(t *testing.T) {
		conn, mockWS, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceCmsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		mockedJSON := `{"catalogId":161,"catalogName":"Delisting","publishDate":1753257631403,"title":"Notice of...","body":"This is...","disclaimer":"Trade on-the-go..."}`
		mockWS.QueueMessage([]byte(mockedJSON))

		resp, err := mockClient.WebsocketStreams.DefaultAPI.ComAnnouncementEn().Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotEmpty(t, mockWS.MessagesWritten)

		require.IsType(t, &common.StreamHandler[models.ComAnnouncementEnResponse]{}, resp)
		require.NotNil(t, resp.On)
		require.NotNil(t, resp.Unsubscribe)

		called := 0
		var got models.ComAnnouncementEnResponse
		mockCallback := func(msg models.ComAnnouncementEnResponse) {
			called++
			got = msg
		}

		resp.On("message", mockCallback)

		if handlers, ok := conn.StreamCallbackMap[resp.Stream]; ok {
			for _, handler := range handlers {
				var rawData interface{}
				if err := json.Unmarshal([]byte(mockedJSON), &rawData); err != nil {
					t.Fatalf("Failed to unmarshal test data: %v", err)
				}

				var msgData map[string]interface{}
				switch v := rawData.(type) {
				case map[string]interface{}:
					msgData = v
				case []interface{}:
					msgData = map[string]interface{}{"data": v}
				default:
					t.Fatalf("Unexpected JSON type: %T", v)
				}

				handlerValue := reflect.ValueOf(handler)
				handlerValue.Call([]reflect.Value{reflect.ValueOf(msgData)})
			}
		}

		require.Equal(t, 1, called, "callback should be called once")
		require.NotNil(t, got)
	})

	t.Run("Test DefaultAPIService ComAnnouncementEn Server Error", func(t *testing.T) {
		conn, _, cleanup := tests.SetupMockClient("123")
		defer cleanup()

		conn.Listen()
		conn.ErrorChan = make(chan error, 1)
		cfg := common.NewConfigurationWebsocketStreams()
		mockClient := client.NewBinanceCmsClient(
			client.WithWebsocketStreams(cfg),
		)
		mockClient.WebsocketStreams.Ws.WsCommon.Connections = []*common.WebSocketConnection{conn}

		resp, err := mockClient.WebsocketStreams.DefaultAPI.ComAnnouncementEn().Execute()
		require.NotNil(t, resp)
		require.NoError(t, err)

		done := make(chan struct{})
		var gotErr error

		resp.OnError(func(err error) {
			gotErr = err
			close(done)
		})

		conn.ErrorChan <- errors.New("Invalid JSON: expected value at line 1 column 30")

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("timeout waiting for error callback")
		}

		require.EqualError(t, gotErr, "Invalid JSON: expected value at line 1 column 30")
	})
}
