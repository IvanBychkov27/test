// Запуск эмуляции DSP VIDEO (listen 127.0.0.1:2999)

package main

import (
	"log"
	"net/http"
	"os"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

const (
	videoData = `
{
    "id":"e85ec70a-067f-4cff-afa1-36d37c60615e",
    "seatbid":[
        {
            "bid":[
                {
                    "id":"Pig40U8mEvM",
                    "impid":"1",
                    "price":1,
                    "adid":"361456.vp",
                    "nurl":"",
                    "adm":"\u003c?xml version='1.0' encoding='UTF-8'?\u003e\u003cVAST xmlns:xsi='https://www.w3.org/2001/XMLSchema-instance' version='3.0'\u003e\u003cAd id='571' sequence='1'\u003e\u003cWrapper allowMultipleAds='true' followAdditonalWrappers='true'\u003e\u003cAdSystem\u003eUnionTraff\u003c/AdSystem\u003e\u003cImpression\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=ECpeCuHTjIEQl9u1rxn3UoKVgZt_tq-EyBJSBoegwQfZLfhBdIWCsrv2Za6G9mys0cYiz8ICqjRSdX_e8YMnGwvB8Gdt5AFAeS9mFFQAxg_E4Z1pRTKBvKbRymLyviR3BCATHyQk8jfLBxT9QyzkJw&v=(visibility)]]\u003e\u003c/Impression\u003e\u003cVASTAdTagURI\u003e\u003c![CDATA[http://a.utraff.com/wrapper?pl=pkPvah8O1o0QZw3m_OeoNY8qAIyyq4g3tocHqcYbn8XgwQ-jwItf7nUTyRqQM3g_5h99wVkJnLu5kje88crWBzEPwhvhpPiQNh1hJE9STznvJVx2-7e6-RBbMtJuCcxBoRBzCmHGAgcLCCIHF__MDbLrt-m-ow1lyUTGEEoqeP9a2AUZ1SZnsY68VSU-DoJWSE4HfHACLJA0yfkeg9u3s6uH2ylcetTFkwof7AAmskmieXCFIWz8aQ5NkFpg3HqObgq5ehHIXNMIKOQImabPs-6CHZuD-ZNxTtxJtingxUcXwM-0rh6oULY7tfRtVKDSSZPanYklQ7auygxL9D-YFDN8rKKUXcuzo3NKeDI-CKt1WN0OeggfnFSXhyMaBiMuiFrN7NlhSOHqERdi00rl4_oFrOj_c_nyMqAcboJql4jjIYOBxBMGFrpl8BWgjBDDNZbI2RPD3lYpzbYucRNIIYeQ98FHeWvIV7pZdHaSpSrjX-N4ASslFYcO3saSypA4JCvGhRfIj6wt-U9w5dm7SJI5c6_iYvzgGLVwY_PPaGGrkczFo2uJsTOgwQskLNL6VlJgP5igDqbc2TZfiUlJUSNIF9SXNvr7YHDi0WiRsRmuHUbg2NoSurVXOiPqeZfDHuqYtHjjHc2_Lci6E1l3S83kx3svwr8IFipoAvbPH-n6uxuJSBdkQY2hgEC4JMLT2I9W5pUsRBpte_4AenIzKhfdfF3bLEVVzsI9nwe2R-7gYKY4OwEwia8q4DKgWyWmiVWuCVLTnPdyMoiPVN-7iXGasFGiTcNL0chQ41IJ_qRnAyIMTk6fTqxJY2Kz_lppDj0Nuy7f0ncRHMH0ytM-ZKatUexkTWQb2tHdO338fWWx0UlFeR5P3qGIA5LJq2LisR38fLbKquvixOdDCkRR0qmpY0g5IvcaEFydUUAXqpkwf9bgqqXlXiIOBmludx7OYKrm0cvgIySgGi3Z-AsMx42Rz9KyPOkrAa7XrjyhtrG6NdyC73t8VtcTPQRlQWajUidXxQPWZieVXQD-SP9WQ30G_inciN-WOD-bJxWT_YKHGgPS6VIulm0d6UIxwqVCbAcXMHxvWks2yQJwO-QURa_bSLQ53X-v-nudX_nxDHCv31QCrGycJjRdj5wEqqRoOt1k14qnAjE2r_lTgoT56UmTUv-_u5Fx3mzP-Pgo_Zjx5MNi3Y9RDsAML1EuWGanpsNXztxL0rC2bE9yAM1IQhXIuGX9-ZwGkNPZE58CUOg6qTPGozcd99nQCaaIRSB0-zZOcvInRrsK-73I9sIG-070dYPPEN0lqsDkycybgo99cWlHS4j7lxhsALMMm4WPXQaUaeP543yVyQLbHlsVFn8C9H8_haxNldJlptgwQj9kYc7bEuP211COvQNl6hMVqaE6l64PvNL-ajGHLi1534TRLczmEXwWw7jLSF6cBjC32KfIOaQURdeIAezDsSxxpejuxnH-6glchv0jBX4RPpyaa9apBsnJwQ3Ffogq6ns3dlWpq1oWDjJ4SDGDLR5I6XCurPnxQ16wB4x5DT-WEavuhZWeNjeDaI8Ck5eiiAo9jKHQYsL7h79OEjgU0-Li&be=2&t=4&pb=1]]\u003e\u003c/VASTAdTagURI\u003e\u003cError\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=PtuXr2-9vQKdeQ7bj35WsYi_gX8WAN8UN59obehWcCqgI24drAwH7E8YFhEsj5KB_LaVPeUn8sHA_I9G8NyyJiksdUfCA4Mzfa7kM2sXkDIv0CrxawCz_wjQdxSeRWR16cZ4JhvDoEFrAR6a4k8gdQ&ec=[ERRORCODE] ]]\u003e\u003c/Error\u003e\u003cCreatives\u003e\u003cCreative adId='571'\u003e\u003cLinear\u003e\u003cTrackingEvents\u003e\u003cTracking event='creativeView'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=vQVXSWCg4ctv5KJ3IofUGUqOunJqbs2ILz9qSjtmvzrmEzCL_eJGQFeaLmqsLktTX33b5m6sxMYzvhnekigjsjIiWm8LrhsFpK9Y-VebayqPWPEXkGZuea5xwUa3uKcdEbWy8SMCuvZMXGjjQWsgEA]]\u003e\u003c/Tracking\u003e\u003cTracking event='start'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=-oUcRmMsnxTqooqcMN5Vi7fDUyvCc7TrEv6ulrK7Pk2khxf3TI3tijsQk-iT_XIJmpFRdIbpf3VBrXGj7v8nKQn_fXqPuVjPb_jJRglMW4fIckK92WEZnEIyB4FiOo4TMHxJaOe1rXsyPWDmb2-sFg]]\u003e\u003c/Tracking\u003e\u003cTracking event='firstQuartile'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=NmRAiQzaxqUcsItaUjIqFwFxAsTUxbZifZX5gFd6OFsLTb-Gx6W4WDdIeshIvYYrQ3eYpf4wFXbbjYN-PCY_uC6AvnFr5H0snH4-7smDfyN_9AlnAxDHgV5yXD1tYyskFDTZQVqjfgFodt_f7rt2cg]]\u003e\u003c/Tracking\u003e\u003cTracking event='midpoint'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=F1D7Xd4Bcoxlm-kbXkbB53wt5PNRKVW2LxX2L17Wp5jKOo4X5-Bn_rxPxi1tafRuWYv5cdlB7eynWFT7eVXwTGWt7aKNiNGYGRmuidtMnzoHb-kxBbAlpoPQCPJer5nz2Pwk-iQoEA5dhvGI6OTd0A]]\u003e\u003c/Tracking\u003e\u003cTracking event='thirdQuartile'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=VS1FaGqmhPbvhKRDkoCd9g1wmMGdLkloXeEmSCOxaKcSs9ml6fIFaFpR5xJsQ1tmqrPK61QE3HPcXIh8_Btla4BXJic8eFXbYN8Xsvvbo7uBk82enCQzBGHnBaFndXhcrGq3giEvUs2sWa2d31idfg]]\u003e\u003c/Tracking\u003e\u003cTracking event='complete'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=g6u2_Z2lJKWtZW4Uc9OkJO_cpRB49NsZnhU3Sa5IcIyT17Hp3k3X9mZSvEt9zP5XdP5r5ROT7U_9Q-YTFTVSDgQr_flcTQOEZ15Vg6nrZPIr191KlA9iYtwj52EpuC5Nxi6Yyw4FNd2PnJAxcF1rQA]]\u003e\u003c/Tracking\u003e\u003cTracking event='mute'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=q1Z6-1fZDXMf0Xf-ahRRx8SkNMjQSgvXJEnHoRunDvawSj8m2JLFaJerG_f8kgfDY-nNlxlJxOPFVNd6M6lCdQnS4_UXIs0sCIzzUZnZDOICKvzR7fxRX3XHobtLvO5q6ba9BaDNI0eextNmnMtuLw]]\u003e\u003c/Tracking\u003e\u003cTracking event='unmute'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=ISiunIEDmpbq1YcL3pHY-dsqgaNyY1Ytiu0WKHDwx1JQquLHtpFgX899WPJmhhLl_ENSV-Po4Iaf0_mamAN4pqMr4PHAfyuh-NRLBtXPXs2BJt7h1ArVQHwHixYfduopS3HLI4i-efqoZSgsppRPig]]\u003e\u003c/Tracking\u003e\u003cTracking event='pause'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=jeThT-vcy8tPd_RVbqfmsEb3omHIprV_y7aCATKbaFmDwpplBax-mewGMfVP973QjLbNtcuWMwyWkuS9CIvKRUhlH0EtENah-6TM8d7BNb8oaZsZPAg641pi5veTexmWiAE20zvjD8GjtNsbfWRRVA]]\u003e\u003c/Tracking\u003e\u003cTracking event='resume'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=0G0thDgyV9wRTQGVO7Erkvam4B5xOS3yBz-tiILqNjdrUABYkrxAqoIVtxUI3k3bKilGvdQGHzfK_qRHlLDq-bmZa4GFUO_VJ56fYCwBJJJlHHGldmEoYyY9DY7XIN59OAAhF8bYfXKnEDjJuPDWrw]]\u003e\u003c/Tracking\u003e\u003cTracking event='fullscreen'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=E60QGhHLmkts0jfHhvP_0L9R7JLyQITV1_glv7YSDYuWuWTaWbrQsNOogsdrGggfEVzGYO4Voa19yjYqWK0p1hx_91CkEz0kOA4DYLrzRKsq7VcrJDK2VMtruzsMugJNZCu1zO8EqepJ710lthKrUw]]\u003e\u003c/Tracking\u003e\u003cTracking event='exitFullscreen'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=APhhWdIdvAYR5YDe2TEHaSxr-Y9ta-4IqYnyqV-kfIeQbQECw3ddBCUDG4HFUhqRb-aetjIdyycFmLQTpdKPKUWpttgekVTFW3dWpMN7iL8mHVbWTMauKJ_X1yPejp4493_T5-YN2BNyat8ScnRKkA]]\u003e\u003c/Tracking\u003e\u003cTracking event='close'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=y6ssyplu7WqAcqxvtZOTxH34wia9tdSEuSR0BcMdrjHivoGRUib-37mV0DUg6D5GWmuO36LikDdaLgmZgk7mnjM1Z3dH13ndSqRx4riiXCzWKMBjvLC8dhW_43hI-1T9UAXI7DWvbqIPHAohlue1Mw]]\u003e\u003c/Tracking\u003e\u003cTracking event='acceptInvitation'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=zZgUAxqY9iFGc4Fwday9xADTcuMQW1_Yfpkv709qqabsgueiGrsi7W2H5PurYxcld_3Y16I36YOi273fQtYh4FV_6L3HXh410YJwv3Ruuifn4YajHQ177wGgKuhG9-zzRHuZkD6JCbqXV1UrW5rx4A]]\u003e\u003c/Tracking\u003e\u003cTracking event='expand'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=8fNRBKsTS_CI6N80Q9oXJ5kEXR3Ty-PTIlBofum2V80DDfebIusXwyB8qbcOnx0JXb3ZBLLwxcvF5q_u-1YzGz29jPkOsvCvkEkQQyZ_efMqs21KwZ8-Rcsgzo9HeuKJf8i6zpEQkJgogbLjjGKfFg]]\u003e\u003c/Tracking\u003e\u003cTracking event='collapse'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=82WjrfdfXvuC82l8_sJY-gWohubCUjwkvpgoyx5DvncKv394UQcS-ZodgYPkXL2Xc5CGJck5-8_uY9-A_xTKGpLGDHJ_GPuYYt0cf4DBjJCNlER66OhLXRfY5W7OUmvF9Jd2KinxhCwnZF_g8ycD1w]]\u003e\u003c/Tracking\u003e\u003cTracking event='closeLinear'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=jfnrXHhW5Tlz4Nodfntf69lod-MSRAJlNXj_6hUD1KSMr7pyHW5lwC7AMp-SCXd9LBhNhsnhiQz1lGJ86mwnqU6TPaL3EisOEshPHqwPlvp5txH1ubIYw_WlnoHeiX7ADcwANRQMjxb7D0tDNE847g]]\u003e\u003c/Tracking\u003e\u003cTracking event='skip'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=rxcUh2s9JnhagZiEfhv5zNy7Qfoe0NPzdc-mhhDDpAB0K9oYbAyyMMRpTAGehuyDfmlHv7SxYNdOPInGITN7Q6cDHdM2a-nruPRWx1faMTJHY9GpsOv0kCwyyvrODfGkSxEAmjBOM_ubsM88qUzHUQ]]\u003e\u003c/Tracking\u003e\u003c/TrackingEvents\u003e\u003cVideoClicks\u003e\u003cClickTracking\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=2J1R7gWC5Om0QVVhagJ51PUNr4MsakBQz3GUHQLS6gONmLJwVz_b9c9VxNmUUEdL5vpUBsmTsY__nHvmbUkOZrD53u4y2zxtUnctw6mlBoq6LEc3hPIkzLxuMWQ7H11uH-0ysR-SjuuVHbIVAoS8lQ]]\u003e\u003c/ClickTracking\u003e\u003c/VideoClicks\u003e\u003c/Linear\u003e\u003c/Creative\u003e\u003c/Creatives\u003e\u003cExtensions\u003e\u003cExtension type='second2'\u003e\u003c![CDATA[http://a.utraff.com/tr?ev=FWH7NI6vxkvLRqal_d0G-33mdAEycg3Grj4BqfiPMNWIGQZxL-ScrNW80zQtp2syUTecz0KskDbueICxLEcUZfLFzmG8Pqd5XmoXZTQuB2qlWy00r_OiqcNzFTDhH-GOTf2c-SJvzfD7csdC47LxGw&v=(visibility)&vl=(volume)]]\u003e\u003c/Extension\u003e\u003cExtension domain='lookmeet.tv' priority='3'/\u003e\u003c/Extensions\u003e\u003c/Wrapper\u003e\u003c/Ad\u003e\u003c/VAST\u003e",
                    "adomain":[
                        "pgbonus.ru"
                    ],
                    "cid":"237477",
                    "crid":"361456.vp",
                    "exp":360,
                    "burl":""
                }
            ]
        }
    ],
    "bidid":"e85ec70a-067f-4cff-afa1-36d37c60615e",
    "cur":"RUB"
}
`
	dspInfo = `
[
 {
   "url": "http://dsp_video.test",
   "price": 0.02999,
   "icon": "dsp_icon.com",
   "image": "dsp_image",
   "description": "dsp_description",
   "title":"dsp_title",
   "nurl":"dsp_nurl:2999"
 }
]
`
)

func (app *Application) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	//res := videoData
	res := dspInfo

	rw.Header().Add("Content-Type", "application/json")
	_, err := rw.Write([]byte(res))
	if err != nil {
		log.Println("error write ", err.Error())
	}
	log.Println("call ", req.RemoteAddr)
}

func main() {
	app := New()
	log.Printf("emulation DSP Video - listen 127.0.0.1:2999")
	err := http.ListenAndServe("127.0.0.1:2999", app)
	if err != nil {
		log.Printf("error %v", err)
		os.Exit(1)
	}
}
