package unmarshaler

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type vendorPostResponse struct {
	RCode    string              `json:"rCode"`
	RMessage string              `json:"rMessage"`
	Data     vendorPostData      `json:"data"`
}

type vendorPostData struct {
	Result []vendorPostResult `json:"result"`
}

type vendorPostResult struct {
	ProductID    string `json:"productId"`
	ProductURL   string `json:"productUrl"`
	ProductImage string `json:"productImage"`
}

type VendorPost struct{}

func (s *VendorPost) UnmarshalResponse(ctx context.Context, body []byte) ([]PartnerResp, error) {
	var resp vendorPostResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		log.WithContext(ctx).Errorf("fail to unmarshal response body: %s", string(body))
		return nil, newInvalidFormatError(body)
	}

	res := make([]PartnerResp, 0, len(resp.Data.Result))
	for _, item := range resp.Data.Result {
		res = append(res, PartnerResp{
			ProductID:    item.ProductID,
			ProductURL:   item.ProductURL,
			ProductImage: item.ProductImage,
		})
	}

	if len(res) == 0 {
		return nil, ErrNoProducts
	}

	return res, nil
}
