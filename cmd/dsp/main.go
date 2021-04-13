// Запуск эмуляции DSP (listen 127.0.0.1:2999)

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
