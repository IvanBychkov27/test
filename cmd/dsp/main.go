// Запуск эмуляции DSP (listen 127.0.0.1:2999)
// Запускаем биндер
// Дергаем в Postmane: http://127.0.0.1:3001/dsptest/?token=3caa525b-9e20-413c-92a4-e4538542a3c4
/*  POST
    Body
{
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
		"body" : "",
		"request" : {
			"ip" : "",
			"ua" : "",
			"sid" : ""
			}
}

*/

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

func (app *Application) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	res := `
[
 { 
   "url": "http://dsp01.test",
   "price": 0.02999,
   "icon": "dsp_icon.com",
   "image": "dsp_image",
   "description": "dsp_description",
   "title":"dsp_title",
   "nurl":"dsp_nurl:2999"
 }
]
`
	rw.Header().Add("Content-Type", "application/json")
	_, err := rw.Write([]byte(res))
	if err != nil {
		log.Println("error write ", err.Error())
	}
	log.Println("call ", req.RemoteAddr)
}

func main() {
	app := New()
	log.Printf("emulation DSP - listen 127.0.0.1:2999")
	err := http.ListenAndServe("127.0.0.1:2999", app)
	if err != nil {
		log.Printf("error %v", err)
		os.Exit(1)
	}
}
