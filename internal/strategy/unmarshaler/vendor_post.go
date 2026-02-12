package unmarshaler

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type vendorPostResp struct {
	RCode    string             `json:"rCode"`
	RMessage string             `json:"rMessage"`
	Data     vendorPostDataResp `json:"data"`
}

type vendorPostDataResp struct {
	Result []vendorPostProduct `json:"result"`
}

type vendorPostProduct struct {
	ProductID    int    `json:"productId"`
	ProductURL   string `json:"productUrl"`
	ProductImage string `json:"productImage"`
}

type VendorPost struct{}

func (s *VendorPost) UnmarshalResponse(ctx context.Context, body []byte) ([]PartnerResp, error) {
	rResp := &vendorPostResp{}
	if err := json.Unmarshal(body, rResp); err != nil {
		log.WithContext(ctx).Errorf("fail to unmarshal response body: %s", string(body))
		return nil, newInvalidFormatError(body)
	}

	// Check response code if needed
	if rResp.RCode != "0" && rResp.RCode != "" {
		return nil, fmt.Errorf("resp code invalid. code: %s, msg: %s", rResp.RCode, rResp.RMessage)
	}

	res := make([]PartnerResp, 0, len(rResp.Data.Result))
	for _, item := range rResp.Data.Result {
		res = append(res, PartnerResp{
			ProductID:    strconv.Itoa(item.ProductID),
			ProductURL:   item.ProductURL,
			ProductImage: item.ProductImage,
		})
	}

	if len(res) == 1 && res[0].ProductID == "0" {
		return nil, ErrInvalidProductID
	}

	if len(res) == 0 {
		return nil, ErrNoProducts
	}

	return res, nil
}
