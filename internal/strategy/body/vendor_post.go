package body

import (
	"fmt"
	"rec-vendor-api/internal/strategy/utils"
	"strings"
)

type VendorPost struct{}

type vendorPostBody struct {
	App       vendorPostApp       `json:"app"`
	Device    vendorPostDevice    `json:"device"`
	Imp       vendorPostImp       `json:"imp"`
	Affiliate vendorPostAffiliate `json:"affiliate"`
}

type vendorPostApp struct {
	BundleID string `json:"bundleId"`
}

type vendorPostDevice struct {
	ID  string `json:"id"`
	Lmt int    `json:"lmt"`
}

type vendorPostImp struct {
	ImageSize string `json:"imageSize"`
}

type vendorPostAffiliate struct {
	SubID string `json:"subId"`
}

func (s *VendorPost) GenerateBody(params Params) any {
	clickIDBase64 := utils.EncodeClickID(params.ClickID)
	body := vendorPostBody{
		App: vendorPostApp{
			BundleID: params.BundleID,
		},
		Device: vendorPostDevice{
			ID:  strings.ToLower(params.UserID),
			Lmt: 0,
		},
		Imp: vendorPostImp{
			ImageSize: fmt.Sprintf("%dx%d", params.ImgWidth, params.ImgHeight),
		},
		Affiliate: vendorPostAffiliate{
			SubID: params.SubID,
		},
	}

	// Store clickIDBase64 for potential tracking use
	_ = clickIDBase64

	return body
}
