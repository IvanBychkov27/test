//Запускает сервер который строит QR код по введенным симвалам

package main

import (
	"flag"
	"github.com/arl/statsviz"
	"html/template"
	"log"
	"net/http"
)

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET">
<input maxLength=1024 size=70 
name=s value="" title="Text to QR Encode">
<input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func QR(w http.ResponseWriter, req *http.Request) {
	err := templ.Execute(w, req.FormValue("s"))
	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	var err error
	err = statsviz.RegisterDefault() // см: 127.0.0.1:1718/debug/statsviz/      - https://github.com/arl/statsviz
	if err != nil {
		log.Fatal("statsviz.RegisterDefault:", err)
	}

	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	log.Println("listen...", *addr)
	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
