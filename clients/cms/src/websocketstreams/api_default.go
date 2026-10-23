/*
Binance Announcements

Binance announcements in English
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
}

func (r ApiComAnnouncementEnRequest) Execute() (*common.StreamHandler[models.ComAnnouncementEnResponse], error) {
	return r.ApiService.ComAnnouncementEnExecute(r)
}

/*
ComAnnouncementEn Method for ComAnnouncementEn
/com_announcement_en

https://developers.binance.com/en/docs/products/announcements/announcement

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
		map[string]string{},
	)
	ws := a.client.Ws

	id := []any{common.GenerateUUID()}
	resp, err := common.CreateStreamHandler[models.ComAnnouncementEnResponse](&common.StreamHandlerWrapper{
		WebsocketStreams: ws,
	}, localStream, id, false)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
