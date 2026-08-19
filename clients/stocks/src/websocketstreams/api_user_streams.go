/*
Stocks Trading WebSocket Streams

WebSocket stream definitions for Binance Stocks Trading. Base URL: wss://nbstream.binance.com/equity
*/

package binancestockswebsocketstreams

import (
	"github.com/binance/binance-connector-go/clients/stocks/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// UserStreamsAPIService UserStreamsAPI Service
type UserStreamsAPIService Service

type ApiOrderReportStreamRequest struct {
	ApiService *UserStreamsAPIService
	listenKey  *string
}

// User data listen key obtained from the Listen Key endpoint.
func (r ApiOrderReportStreamRequest) ListenKey(listenKey string) ApiOrderReportStreamRequest {
	r.listenKey = &listenKey
	return r
}

func (r ApiOrderReportStreamRequest) Execute() (*common.StreamHandler[models.OrderReportStreamResponse], error) {
	return r.ApiService.OrderReportStreamExecute(r)
}

/*
OrderReportStream Order Report Stream
/<listenKey>@orderReport

https://developers.binance.com/en/docs/catalog/advanced-trading-stocks-trading/api/ws-streams/user-streams#order-report-stream

@param listenKey User data listen key obtained from the Listen Key endpoint.
@return ApiOrderReportStreamRequest
*/
func (a *UserStreamsAPIService) OrderReportStream() ApiOrderReportStreamRequest {
	return ApiOrderReportStreamRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return OrderReportStreamResponse
func (a *UserStreamsAPIService) OrderReportStreamExecute(r ApiOrderReportStreamRequest) (*common.StreamHandler[models.OrderReportStreamResponse], error) {
	if r.listenKey == nil {
		return nil, common.ReportError("listenKey is required and must be specified")
	}

	localStream := common.WsStreamsPlaceholder(
		"/<listenKey>@orderReport"[1:],
		map[string]string{
			"listenKey": func() string {
				if r.listenKey == nil {
					return ""
				}
				return *r.listenKey
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.OrderReportStreamResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
