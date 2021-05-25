// Запуск генерации в папке с файлом api.proto
// protoc -I=. api.proto --go_out=plugins=grpc:.
package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"os"
	"test/cmd/proto/api"
)

func main() {
	var addr string
	flag.StringVar(&addr, "a", "127.0.0.1:2100", "listen address")
	flag.Parse()

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("error dial, %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	client := api.NewCachejobClient(conn)

	data := api.BinderRequest{
		Name: "002",
	}

	PingCachejob(client, &data)

	GetDSP(client, &data)
	GetTokens(client, &data)
	//GetDPL(client, &data)
	//GetDefaultDPL(client, &data)
	//GetCampaign(client, &data)
	//GetUser(client, &data)
	//GetEndpoint(client, &data)

}

func GetTokens(client api.CachejobClient, data *api.BinderRequest) {
	result, err := client.GetTokens(context.Background(), data)
	if err != nil {
		log.Printf("error GetTokens, %v", err)
		os.Exit(1)
	}

	tokens := result.Data
	log.Println("tokens: ", tokens)
}

func PingCachejob(client api.CachejobClient, data *api.BinderRequest) {
	result, err := client.PingCachejob(context.Background(), data)
	if err != nil {
		log.Printf("error PingCachejob, %v", err)
		os.Exit(1)
	}
	log.Printf("\nresult: PingCachejob\ndata: %s\n\n", result.Data)
}

func GetDSP(client api.CachejobClient, data *api.BinderRequest) {
	result, err := client.GetDSP(context.Background(), data)
	if err != nil {
		log.Printf("error GetDSP, %v", err)
		os.Exit(1)
	}
	//log.Printf("\nresult: GetDSP\ndata1: %s\ndata2: %s\n\n", result.Data1, result.Data2)

	dsps1 := sync(result.Data1)
	dsps2 := sync(result.Data2)

	log.Printf("\nresult: GetDSP\ndata1: %s\n data2: %s\n\n", dsps1, dsps2)
}

func GetDPL(client api.CachejobClient, data *api.BinderRequest) {
	result, err := client.GetDPL(context.Background(), data)
	if err != nil {
		log.Printf("error GetDPL, %v", err)
		os.Exit(1)
	}
	log.Printf("\nresult: GetDPL\ndata1: %s\ndata2: %s\n\n", result.Data1, result.Data2)
}

func GetDefaultDPL(client api.CachejobClient, data *api.BinderRequest) {
	result, err := client.GetDefaultDPL(context.Background(), data)
	if err != nil {
		log.Printf("error GetDefaultDPL, %v", err)
		os.Exit(1)
	}
	log.Printf("\nresult: GetDefaultDPL\ndata: %s\n\n", result.Data)
}

func GetCampaign(client api.CachejobClient, data *api.BinderRequest) {
	result, err := client.GetCampaign(context.Background(), data)
	if err != nil {
		log.Printf("error GetCampaign, %v", err)
		os.Exit(1)
	}
	log.Printf("\nresult: GetCampaign\ndata: %s\n\n", result.Data)
}

func GetUser(client api.CachejobClient, data *api.BinderRequest) {
	result, err := client.GetUser(context.Background(), data)
	if err != nil {
		log.Printf("error GetUser, %v", err)
		os.Exit(1)
	}
	log.Printf("\nresult: GetUser\ndata: %s\n\n", result.Data)
}

func GetEndpoint(client api.CachejobClient, data *api.BinderRequest) {
	result, err := client.GetEndpoint(context.Background(), data)
	if err != nil {
		log.Printf("error GetEndpoint, %v", err)
		os.Exit(1)
	}
	log.Printf("\nresult: GetEndpoint\ndata: %s\n\n", result.Data)
}

func sync(data []*api.DSP) (dsps []*DSP) {
	for _, d := range data {
		dsp := NewDSP()

		dsp.ID = int(d.ID)
		dsp.NetworkID = int(d.NetworkID)
		dsp.URL = d.URL
		dsp.PaymentTarget = int(d.PaymentTarget)
		dsp.PaymentType = int(d.PaymentType)
		dsp.Rate = int(d.Rate)
		dsp.Method = int(d.Method)
		dsp.RequestType = int(d.RequestType)
		dsp.ResponseType = int(d.ResponseType)
		dsp.AdvertiserID = int(d.AdvertiserID)
		dsp.CustomResponseSettings = d.CustomResponseSettings
		dsp.TrafficType = int(d.TrafficType)
		dsp.AdvertiserRate = int(d.AdvertiserRate)
		dsp.Tags = d.Tags
		dsp.RulesToUsersAllow = d.RulesToUsersAllow
		dsp.RulesToUsersAllowType = int(d.RulesToUsersAllowType)
		dsp.RulesToUsersDeny = d.RulesToUsersDeny
		dsp.RulesToUsersDenyType = int(d.RulesToUsersDenyType)
		dsp.RulesToEndpointsAllow = d.RulesToEndpointsAllow
		dsp.RulesToEndpointsAllowType = int(d.RulesToEndpointsAllowType)
		dsp.RulesToEndpointsDeny = d.RulesToEndpointsDeny
		dsp.RulesToEndpointsDenyType = int(d.RulesToEndpointsDenyType)
		dsp.AdvertiserTags = d.AdvertiserTags
		dsp.AdvertiserRulesToUsersAllow = d.AdvertiserRulesToUsersAllow
		dsp.AdvertiserRulesToUsersAllowType = int(d.AdvertiserRulesToUsersAllowType)
		dsp.AdvertiserRulesToUsersDeny = d.AdvertiserRulesToUsersDeny
		dsp.AdvertiserRulesToUsersDenyType = int(d.AdvertiserRulesToUsersDenyType)
		dsp.AdvertiserRulesToEndpointsAllow = d.AdvertiserRulesToEndpointsAllow
		dsp.AdvertiserRulesToEndpointsAllowType = int(d.AdvertiserRulesToEndpointsAllowType)
		dsp.AdvertiserRulesToEndpointsDeny = d.AdvertiserRulesToEndpointsDeny
		dsp.AdvertiserRulesToEndpointsDenyType = int(d.AdvertiserRulesToEndpointsDenyType)
		dsp.CallNURL = d.CallNURL
		dsp.SubscriberAgeMin = int(d.SubscriberAgeMin)
		dsp.SubscriberAgeMax = int(d.SubscriberAgeMax)
		dsp.RedirectType = int(d.RedirectType)
		dsp.GeoFilterType = int(d.GeoFilterType)
		dsp.GeoFilterMapString = d.GeoFilterMapString
		dsp.NURLEvent = d.NURLEvent
		dsp.RequestFrequency = int(d.RequestFrequency)
		dsp.FastFilterSourceIDStr = d.FastFilterSourceIDStr
		dsp.FastFilterSourceIDEncodedStr = d.FastFilterSourceIDEncodedStr
		dsp.CheckIPMismatch = d.CheckIPMismatch
		dsp.CheckUAMismatch = d.CheckUAMismatch
		dsp.StopOnDClick = d.StopOnDClick
		dsp.FilterPlatform = int(d.FilterPlatform)
		dsp.ExchangeRate = d.ExchangeRate
		dsp.DatacenterID = int(d.DatacenterID)
		dsp.FastFilterSourceIDFilterType = int(d.FastFilterSourceIDFilterType)
		dsp.DplID = int(d.DplID)
		dsp.RequestOpenRTBExtBidRequest = d.RequestOpenRTBExtBidRequest
		dsp.RequestOpenRTBExtImp = d.RequestOpenRTBExtImp
		dsp.SendRefererHeader = d.SendRefererHeader
		dsp.CircuitBreakerLimit = int(d.CircuitBreakerLimit)
		dsp.RandomStarValues = d.RandomStarValues
		dsp.OverrideAllReferrers = d.OverrideAllReferrers
		dsp.OverrideReferrerListStr = d.OverrideReferrerListStr
		dsp.FFUserFilterType = int(d.FFUserFilterType)
		dsp.FFEndpointFilterType = int(d.FFEndpointFilterType)
		dsp.FFAdvertiserToPublisherFilterType = int(d.FFAdvertiserToPublisherFilterType)
		dsp.FFAdvertiserToEndpointFilterType = int(d.FFAdvertiserToEndpointFilterType)
		dsp.EscapeOpenRTBNativeRequest = d.EscapeOpenRTBNativeRequest
		dsp.RateLimitMax = d.RateLimitMax
		dsp.Strictly = d.Strictly
		dsp.MaxSourcesCount = int(d.MaxSourcesCount)
		dsp.Prepay = d.Prepay
		dsp.ProcessDoubleEvent = d.ProcessDoubleEvent
		dsp.TrafficQuality = int(d.TrafficQuality)
		dsp.ImpClickStrategy = int(d.ImpClickStrategy)
		dsp.OsDesktopJSONString = d.OsDesktopJSONString
		dsp.OsMobileJSONString = d.OsMobileJSONString
		dsp.UaDesktopJSONString = d.UaDesktopJSONString
		dsp.UaMobileJSONString = d.UaMobileJSONString
		dsp.DesktopOSFilterType = int(d.DesktopOSFilterType)
		dsp.DesktopUAFilterType = int(d.DesktopUAFilterType)
		dsp.MobileOSFilterType = int(d.MobileOSFilterType)
		dsp.MobileUAFilterType = int(d.MobileUAFilterType)
		dsp.CallWinURL = d.CallWinURL

		dsps = append(dsps, dsp)
	}
	return dsps
}

func NewDSP() *DSP {
	return &DSP{
		Tags:                            make([]int64, 0),
		RulesToUsersAllow:               make([]int64, 0),
		RulesToUsersDeny:                make([]int64, 0),
		RulesToEndpointsAllow:           make([]int64, 0),
		RulesToEndpointsDeny:            make([]int64, 0),
		AdvertiserTags:                  make([]int64, 0),
		AdvertiserRulesToUsersAllow:     make([]int64, 0),
		AdvertiserRulesToUsersDeny:      make([]int64, 0),
		AdvertiserRulesToEndpointsAllow: make([]int64, 0),
		AdvertiserRulesToEndpointsDeny:  make([]int64, 0),
	}
}

type DSP struct {
	ID                                  int
	NetworkID                           int
	URL                                 string
	PaymentTarget                       int
	PaymentType                         int
	Rate                                int
	Method                              int
	RequestType                         int
	ResponseType                        int
	AdvertiserID                        int
	CustomResponseSettings              string
	TrafficType                         int
	AdvertiserRate                      int
	Tags                                []int64
	RulesToUsersAllow                   []int64
	RulesToUsersAllowType               int
	RulesToUsersDeny                    []int64
	RulesToUsersDenyType                int
	RulesToEndpointsAllow               []int64
	RulesToEndpointsAllowType           int
	RulesToEndpointsDeny                []int64
	RulesToEndpointsDenyType            int
	AdvertiserTags                      []int64
	AdvertiserRulesToUsersAllow         []int64
	AdvertiserRulesToUsersAllowType     int
	AdvertiserRulesToUsersDeny          []int64
	AdvertiserRulesToUsersDenyType      int
	AdvertiserRulesToEndpointsAllow     []int64
	AdvertiserRulesToEndpointsAllowType int
	AdvertiserRulesToEndpointsDeny      []int64
	AdvertiserRulesToEndpointsDenyType  int
	CallNURL                            bool
	SubscriberAgeMin                    int
	SubscriberAgeMax                    int
	RedirectType                        int
	GeoFilterType                       int
	GeoFilterMapString                  string
	NURLEvent                           string
	RequestFrequency                    int
	FastFilterSourceIDStr               string
	FastFilterSourceIDEncodedStr        string
	CheckIPMismatch                     bool
	CheckUAMismatch                     bool
	StopOnDClick                        bool
	FilterPlatform                      int
	ExchangeRate                        float64
	DatacenterID                        int
	FastFilterSourceIDFilterType        int
	DplID                               int
	RequestOpenRTBExtBidRequest         string
	RequestOpenRTBExtImp                string
	SendRefererHeader                   bool
	CircuitBreakerLimit                 int
	RandomStarValues                    string
	OverrideAllReferrers                bool
	OverrideReferrerListStr             string
	FFUserFilterType                    int
	FFEndpointFilterType                int
	FFAdvertiserToPublisherFilterType   int
	FFAdvertiserToEndpointFilterType    int
	EscapeOpenRTBNativeRequest          bool
	RateLimitMax                        int64
	Strictly                            bool
	MaxSourcesCount                     int
	Prepay                              bool
	ProcessDoubleEvent                  bool
	TrafficQuality                      int
	ImpClickStrategy                    int
	OsDesktopJSONString                 string
	OsMobileJSONString                  string
	UaDesktopJSONString                 string
	UaMobileJSONString                  string
	DesktopOSFilterType                 int
	DesktopUAFilterType                 int
	MobileOSFilterType                  int
	MobileUAFilterType                  int
	CallWinURL                          bool
}
