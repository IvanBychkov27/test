// https://golangs.org/json

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {

	//jsonMarshal()

	//jsonUnmarshal()

	//jsonUnmarshalMAP()

	startUnmarshal()
}

// =====================

/*
json для проекта collector
{
"subscription": {
				"endpoint" : "https://updates.push.services.mozilla.com/wpush/v2/gAAAAABgdcO2M-ZPMHQNI5WDGRmrlnD5y9aLedXnAwU82hcP_JikOfIzFvzHOsWN9DghBO-uDhObwvgMGvp2E_WpJVue2RfJYuKm9U7xaInqRUpvgHXijpxoIYI7nnOTx0Hy_4uoibCzp84SBqzTSBhgsciM-0D5qiDfeSItwgLSzs1d_EVQaoE",
				"keys"     : {
								"auth"   : "B3MoFQ7nwCMOV2q7xl6qQg",
								"p256dh" : "BOxJMNmne1-BxOqjhoCHEAZ7VLaZzUJpwA4LJclnRqYxzWH-Wr7-rkhHqvpcPMUwxM2uYMUsgWx7rq5iEovKrGY"
							}
				}
}
*/

func startUnmarshal() {
	data := []byte(`
{
	"key": "body key string"
}`)

	resp, err := getRequest(data)
	if err != nil {
		fmt.Println("error unmarshal: ", err.Error())
		return
	}

	fmt.Println(resp)
}

type DateRequest struct {
	Key string `json:"key"`
}

func getRequest(body []byte) (DateRequest, error) {
	var d DateRequest
	err := json.Unmarshal(body, &d)
	if err != nil {
		return d, err
	}

	return d, nil
}

// ===== jsonUnmarshalMAP =========================================

func jsonUnmarshalMAP() {
	data := []byte(`
{
	"Ru": {},
    "ua": {},
	"YY": {}
}`)
	var d map[string]struct{}
	err := json.Unmarshal(data, &d)
	chk(err)
	fmt.Println(d)

	updateLangTarget(d)

	//for k := range d {
	//	d[strings.ToLower(k)] = struct{}{}
	//}

	fmt.Println(d)
	key := "ru"

	_, ok := d[strings.ToLower(key)]
	fmt.Println("ok  = ", ok)

}

func updateLangTarget(a map[string]struct{}) {
	for k := range a {
		a[strings.ToLower(k)] = struct{}{}
	}
}

// ===== jsonUnmarshal =========================================

type CustomResponseSettings struct {
	ItemsRoot        string `json:"items_root,omitempty"`
	Price            string `json:"price,omitempty"`
	URL              string `json:"url,omitempty"`
	Title            string `json:"title,omitempty"`
	Description      string `json:"description,omitempty"`
	IconURL          string `json:"icon_url,omitempty"`
	ImageURL         string `json:"image_url,omitempty"`
	OriginalIconURL  string `json:"original_icon_url,omitempty"`
	OriginalImageURL string `json:"original_image_url,omitempty"`
	NURL             string `json:"nurl,omitempty"`
}

type DSP struct {
	ID           int `json:"id"`
	AdvertiserID int `json:"advertiser_id"`
	NetworkID    int `json:"network_id"`
	//Method                            HTTPMethod                                 `json:"method"`
	URL            string `json:"url"`
	Rate           int    `json:"rate"`
	AdvertiserRate int    `json:"advertiser_rate"`
	//PaymentType                       types.PaymentType                          `json:"payment_type"`
	//PaymentTarget                     types.PaymentTarget                        `json:"payment_target"`
	//RequestType                       types.RequestType                          `json:"request_type"`
	//ResponseType                      types.ResponceType                         `json:"response_type"`
	//TrafficType                       types.TrafficType                          `json:"traffic_type"`
	CustomResponseSettings CustomResponseSettings `json:"custom_response_settings"`
	//TagRules                          map[tags.RulesTarget]tags.Rules            `json:"tag_rules"`
	Tags []int64 `json:"tags"`
	//AdvertiserTagRules                map[tags.RulesTarget]tags.Rules            `json:"advertiser_tag_rules"`
	AdvertiserTags   []int64 `json:"advertiser_tags"`
	CallNURL         bool    `json:"call_nurl"`
	NURLEvent        string  `json:"nurl_event"`
	SubscriberAgeMin int     `json:"subscriber_age_min"`
	SubscriberAgeMax int     `json:"subscriber_age_max"`
	//RedirectType                      types.RedirectType                         `json:"redirect_type"`
	Geo map[string]map[string]map[string]interface{} `json:"geo"`
	//GeoFilterType                     types.FilterType                           `json:"geo_filter_type"`
	RequestOpenRTBExtBidRequest string `json:"request_openrtb_ext_bidrequest"`
	RequestOpenRTBExtImp        string `json:"request_openrtb_ext_imp"`
	RequestFrequency            int    `json:"request_frequency"`
	//FastFilterSourceIDFilterType      types.FilterType                           `json:"fast_filter_sourceid_filter_type"`
	CheckIPMismatch bool `json:"check_ip_mismatch"`
	CheckUAMismatch bool `json:"check_ua_mismatch"`
	StopOnDClick    bool `json:"stop_on_double_click"`
	//FilterPlatform                    types.PlatformFilter                       `json:"filter_platform"`
	ExchangeRate               float64          `json:"exchange_rate"`
	DatacenterID               int              `json:"data_center_id"`
	DemandPriceLimit           map[string]int64 `json:"demand_price_limit"`
	SendRefererHeader          bool             `json:"send_referer_header"`
	CircuitBreakerLimit        int              `json:"throttle_limit"`
	RandomStarValues           []string
	OverrideAllReferrers       bool
	OverrideReferrerList       []string
	EscapeOpenRTBNativeRequest bool
	//FFUserFilterType                  types.FilterType
	//FFEndpointFilterType              types.FilterType
	//FFAdvertiserToPublisherFilterType types.FilterType
	//FFAdvertiserToEndpointFilterType  types.FilterType
	RateLimitCount     int64
	RateLimitMax       int64
	RateLimitChan      chan struct{}
	Strictly           bool `json:"strictly"`
	MaxSourcesCount    int
	Prepay             bool
	ProcessDoubleEvent bool
	TrafficQuality     int
}

func jsonUnmarshal() {
	var dsp DSP
	body := all_body()

	err := json.Unmarshal(body, &dsp)
	chk(err)

	//fmt.Println(string(body))
	fmt.Println(dsp)
}

func all_body() []byte {
	// Запуск dsptest
	// http://127.0.0.1:3001/dsptest/?token=3caa525b-9e20-413c-92a4-e4538542a3c4
	// Body:  - ввести только json
	data := []byte(`{
	"dsp" : {
			"custom_response_settings": {
                "url": "url",
                "nurl": "nurl",
                "price": "price",
                "title": "title",
                "icon_url": "icon",
                "image_url": "image",
                "items_root": null,
                "description": "description",
                "original_icon_url": null,
                "original_image_url": null
                },
			"escape_openrtb_native_request": false,
			"exchange_rate": 1,
            "id": 777,
            "macros_random_values": "null",
            "name": "rtrt - Copy",
            "nurl_event": 2,
            "request_method": 1,
            "request_openrtb_bidrequest_ext": "null",
            "request_openrtb_imp_ext": "null",
            "request_settings_template_item": "null",
            "request_settings_template_request": "null",
            "request_type": 3,
            "response_type": 2,
            "url": "http://127.0.0.1:2999"
		},
		"body" : "[{ \"url\": \"http://dsp01.test\",\"price\": 0.028,\"icon\": \"dsp_icon.com\",\"image\": \"dsp_image\",\"description\": \"dsp_description\",\"title\":\"dsp_title\",\"nurl\":\"dsp_nurl:2999\"}]",
		"request" : {
			"ip" : "",
			"ua" : "",
			"sid" : ""
			}
}`)

	//data := []byte(`{
	//   "ac_tive": true,
	//   "url": "test.ru",
	//   "u" : {"Id":1, "Name":"Iv", "Telefon":"22-33-55", "Place":{"City":"Br", "Street":"Sov", "House":102, "Apartment":221}},
	//   "id": 777,
	//   "custom_response_settings": {"description": "null1", "icon_url": "null2", "original_icon_url": "null3", "image_url": "null4", "original_image_url": "null5"}
	//}`)
	//
	//data := []byte(`{
	//	"active": false,
	//	"auto_rate": 100,
	//	"balance": 0,
	//	"block_empty_referer": false,
	//	"call_nurl": false,
	//	"check_ip_mismatch": false,
	//	"check_ua_mismatch": false,
	//	"circuit_breaker_limit": 0,
	//	"comment": "null",
	//	"custom_response_settings": {"description": "null", "icon_url": "null", "original_icon_url": "null", "image_url": "null", "original_image_url": "null"},
	//	"daily_budget": 0,
	//	"data_center_id": "null",
	//	"demand_price_limit_id": "null",
	//	"dsp_group_id": 1,
	//	"endpoint_allow_tag_rule_type": 0,
	//	"endpoint_deny_tag_rule_type": 0,
	//	"escape_openrtb_native_request": false,
	//	"exchange_rate": 1,
	//	"fast_filter_endpoint": [],
	//	"fast_filter_endpoint_filter_type": 0,
	//	"fast_filter_sourceid": "null",
	//	"fast_filter_sourceid_encoded": "null",
	//	"fast_filter_sourceid_filter_type": 0,
	//	"fast_filter_user": [],
	//	"fast_filter_user_filter_type": 0,
	//	"filter_platform": 0,
	//	"freeze_rules": "",
	//	"geo_filter_map": [],
	//	"geo_filter_map_type": 0,
	//	"id": 29,
	//	"is_prepay": false,
	//	"macros_random_values": "null",
	//	"max_auto_rate": 100,
	//	"max_sources_count": 0,
	//	"min_balance": 0,
	//	"name": "rtrt - Copy",
	//	"nurl_event": 2,
	//	"override_all_referrers": false,
	//	"override_referrer_list": "",
	//	"payment_target": 1,
	//	"payment_type": 0,
	//	"qps_limit": 0,
	//	"rate": 100,
	//	"rate_buffer": 0,
	//	"rate_total": 167,
	//	"rate_user": 167,
	//	"redirect_type": 0,
	//	"request_frequency": 100,
	//	"request_method": 1,
	//	"request_openrtb_bidrequest_ext": "null",
	//	"request_openrtb_imp_ext": "null",
	//	"request_settings_template_item": "null",
	//	"request_settings_template_request": "null",
	//	"request_type": 0,
	//	"response_type": 0,
	//	"send_referer_header": false,
	//	"spending_strategy": 0,
	//	"status": 1,
	//	"stop_on_double_click": false,
	//	"stopped_by_balance_at": "null",
	//	"strictly": false,
	//	"subscriber_age_max": 0,
	//	"subscriber_age_min": 0,
	//	"tags_endpoint_allow": [],
	//	"tags_endpoint_deny": [],
	//	"tags_self": [],
	//	"tags_user_allow": [],
	//	"tags_user_deny": [],
	//	"timeout": 400,
	//	"traffic_quality": 0,
	//	"traffic_type": 0,
	//	"url": "http://ya.ru",
	//	"user_allow_tag_rule_type": 0,
	//	"user_deny_tag_rule_type": 0,
	//	"user_id": 4,
	//	"write_demand": false
	//}`)

	// 1 active: false
	// 2 auto_rate: 100
	// 3 balance: 0
	// 4 block_empty_referer: false
	// 5 call_nurl: false
	// 6 check_ip_mismatch: false
	// 7 check_ua_mismatch: false
	// 8 circuit_breaker_limit: 0
	// 9 comment: null
	// 10 custom_response_settings: {description: null, icon_url: null, original_icon_url: null, image_url: null, original_image_url: null,…}
	// 11 daily_budget: 0
	// 12 data_center_id: null
	// 13 demand_price_limit_id: null
	// 14 dsp_group_id: 1
	// 15 endpoint_allow_tag_rule_type: 0
	// 16 endpoint_deny_tag_rule_type: 0
	// 17 escape_openrtb_native_request: false
	// 18 exchange_rate: 1
	// 19 fast_filter_endpoint: []
	// 20 fast_filter_endpoint_filter_type: 0
	// 21 fast_filter_sourceid: null
	// 22 fast_filter_sourceid_encoded: null
	// 23 fast_filter_sourceid_filter_type: 0
	// 24 fast_filter_user: []
	// 25 fast_filter_user_filter_type: 0
	// 26 filter_platform: 0
	// 27 freeze_rules: ""
	// 28 geo_filter_map: []
	// 29 geo_filter_map_type: 0
	// 30 id: 29
	// 40 is_prepay: false
	// 41 macros_random_values: null
	// 42 max_auto_rate: 100
	// 43 max_sources_count: 0
	// 44 min_balance: 0
	// 45 name: "rtrt - Copy"
	// 46 nurl_event: 2
	// 47 override_all_referrers: false
	// 48 override_referrer_list: ""
	// 49 payment_target: 1
	// 50 payment_type: 0
	// 51 qps_limit: 0
	// 52 rate: 100
	// 53 rate_buffer: 0
	// 54 rate_total: 167
	// 55 rate_user: 167
	// 56 redirect_type: 0
	// 57 request_frequency: 100
	// 58 request_method: 1
	// 59 request_openrtb_bidrequest_ext: null
	// 60 request_openrtb_imp_ext: null
	// 61 request_settings_template_item: null
	// 62 request_settings_template_request: null
	// 63 request_type: 0
	// 64 response_type: 0
	// 65 send_referer_header: false
	// 66 spending_strategy: 0
	// 67 status: 1
	// 68 stop_on_double_click: false
	// 69 stopped_by_balance_at: null
	// 70 strictly: false
	// 71 subscriber_age_max: 0
	// 72 subscriber_age_min: 0
	// 73 tags_endpoint_allow: []
	// 74 tags_endpoint_deny: []
	// 75 tags_self: []
	// 76 tags_user_allow: []
	// 77 tags_user_deny: []
	// 78 timeout: 400
	// 79 traffic_quality: 0
	// 80 traffic_type: 0
	// 81 url: "http://ya.ru"
	// 82 user_allow_tag_rule_type: 0
	// 83 user_deny_tag_rule_type: 0
	// 84 user_id: 4
	// 85 write_demand: false

	return data
}

// ======= json ============================================
//
//func j() {
//
//	"dsp" : {
//			"active": false,
//			"auto_rate": 100,
//			"balance": 0,
//			"block_empty_referer": false,
//			"call_nurl": false,
//			"check_ip_mismatch": false,
//			"check_ua_mismatch": false,
//			"circuit_breaker_limit": 0,
//			"comment": "null",
//			"custom_response_settings": {
//			          "description": "null",
//			          "icon_url": "null",
//			          "original_icon_url": "null",
//			          "image_url": "null",
//			          "original_image_url": "null"
//		         },
//			"daily_budget": 0,
//			"data_center_id": "null",
//			"demand_price_limit_id": "null",
//			"dsp_group_id": 1,
//			"endpoint_allow_tag_rule_type": 0,
//			"endpoint_deny_tag_rule_type": 0,
//			"escape_openrtb_native_request": false,
//			"exchange_rate": 1,
//			"fast_filter_endpoint": [],
//            "fast_filter_endpoint_filter_type": 0,
//            "fast_filter_sourceid": "null",
//            "fast_filter_sourceid_encoded": "null",
//            "fast_filter_sourceid_filter_type": 0,
//            "fast_filter_user": [],
//            "fast_filter_user_filter_type": 0,
//            "filter_platform": 0,
//            "freeze_rules": "",
//            "geo_filter_map": [],
//            "geo_filter_map_type": 0,
//            "id": 29,
//            "is_prepay": false,
//            "macros_random_values": "null",
//            "max_auto_rate": 100,
//            "max_sources_count": 0,
//            "min_balance": 0,
//            "name": "rtrt - Copy",
//            "nurl_event": 2,
//            "override_all_referrers": false,
//            "override_referrer_list": "",
//            "payment_target": 1,
//            "payment_type": 0,
//            "qps_limit": 0,
//            "rate": 100,
//            "rate_buffer": 0,
//            "rate_total": 167,
//            "rate_user": 167,
//            "redirect_type": 0,
//            "request_frequency": 100,
//            "request_method": 1,
//            "request_openrtb_bidrequest_ext": "null",
//            "request_openrtb_imp_ext": "null",
//            "request_settings_template_item": "null",
//            "request_settings_template_request": "null",
//            "request_type": 0,
//            "response_type": 0,
//            "send_referer_header": false,
//            "spending_strategy": 0,
//            "status": 1,
//            "stop_on_double_click": false,
//            "stopped_by_balance_at": "null",
//            "strictly": false,
//            "subscriber_age_max": 0,
//            "subscriber_age_min": 0,
//            "tags_endpoint_allow": [],
//            "tags_endpoint_deny": [],
//            "tags_self": [],
//            "tags_user_allow": [],
//            "tags_user_deny": [],
//            "timeout": 400,
//            "traffic_quality": 0,
//            "traffic_type": 0,
//            "url": "http://127.0.0.1:2999",
//            "user_allow_tag_rule_type": 0,
//            "user_deny_tag_rule_type": 0,
//            "user_id": 4,
//            "write_demand": false
//		},
//		"body" : "",
//		"request" : {
//			"ip" : "",
//			"ua" : "",
//			"sid" : ""
//			}
//}

// ===== jsonMarshal =========================================
type User struct {
	Id      int
	Name    string
	Telefon string
	Place   Address
}

type Address struct {
	City      string
	Street    string
	House     int
	Apartment int
}

func jsonMarshal() {
	a1 := Address{"Bryansk", "Sovet", 4, 82}
	u1 := User{1, "Iv", "9102306511", a1}

	json_data, err := json.Marshal(u1)
	chk(err)
	fmt.Println(string(json_data))

	fmt.Println()

	a := []Address{
		{"Bryansk", "22 sezd", 14, 93},
		{"Bryansk", "sov", 4, 82},
		{"Bryansk", "sov", 4, 82},
	}

	users := []User{
		{2, "Andrey", "9208647777", a[0]},
		{3, "Tanya", "9102379977", a[1]},
		{4, "Olya", "9107432216", a[2]},
	}

	json_data2, err := json.Marshal(users)
	chk(err)
	fmt.Println(string(json_data2))

}

func chk(err error) {
	if err != nil {
		fmt.Println("error: ", err)
	}
}
