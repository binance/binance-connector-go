/*
Prediction Trading Streams

Binance Web3 Prediction pushes 14 types of per-user business events in real time via the Binance SApi WebSocket channel, covering trading (market / limit), reward claim, fund transfer and market resolution, plus order book (bid/ask) quotes broadcast per topic. Business events are routed by exact `(topic, userId)` match; order book topics are broadcast to every subscriber. Delivery is at-most-once, ordering is not guaranteed, and offline messages are not cached — pull the latest state via the REST API after reconnecting and deduplicate on `pushId` (business events) or `marketId` + `updateTimestampMs` (order book).
*/

package binancecmswebsocketstreams

import (
	"github.com/binance/binance-connector-go/clients/cms/src/websocketstreams/models"
	"github.com/binance/binance-connector-go/common/v2/common"
)

// DefaultAPIService DefaultAPI Service
type DefaultAPIService Service

type ApiComAnnouncementEnRequest struct {
	ApiService *DefaultAPIService
	id         *string
}

// Unique WebSocket request ID.
func (r ApiComAnnouncementEnRequest) Id(id string) ApiComAnnouncementEnRequest {
	r.id = &id
	return r
}

func (r ApiComAnnouncementEnRequest) Execute() (*common.StreamHandler[models.ComAnnouncementEnResponse], error) {
	return r.ApiService.ComAnnouncementEnExecute(r)
}

/*
ComAnnouncementEn Method for ComAnnouncementEn
/com_announcement_en

https://developers.binance.com/en/docs/products/announcements/announcement

@param id Unique WebSocket request ID.
@return ApiComAnnouncementEnRequest
*/
func (a *DefaultAPIService) ComAnnouncementEn() ApiComAnnouncementEnRequest {
	return ApiComAnnouncementEnRequest{
		ApiService: a,
	}
}

// Execute executes the request
//
//	@return ComAnnouncementEnResponse
func (a *DefaultAPIService) ComAnnouncementEnExecute(r ApiComAnnouncementEnRequest) (*common.StreamHandler[models.ComAnnouncementEnResponse], error) {

	localStream := common.WsStreamsPlaceholder(
		"/com_announcement_en"[1:],
		map[string]string{
			"id": func() string {
				if r.id == nil {
					return ""
				}
				return *r.id
			}(),
		},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	if r.id != nil {
		id = []any{*r.id}
	}
	resp, err := common.CreateStreamHandler[models.ComAnnouncementEnResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
