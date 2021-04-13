package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math/rand"
	"strconv"

	"log"
	"net/http"
)

/*
Работа с метриками
   1. создаем переменную
   2. регистрируем ее при запуске программы
   3. записываем данные в метрику

Для просмотра данных метрики нужно ввести в браузере: 127.0.0.1:2000/metrics

*/

var ( // создаем переменную метрики
	requestsMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		//Namespace:   "",
		//Subsystem:   "",
		Name: "requests",
		Help: "Requests Count",
		//ConstLabels: nil,
	}, []string{"server_name", "hour"}) // создаем lable
)

type H struct {
}

func (h *H) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	hour := rand.Intn(23)
	serverName := "server1"
	n := rand.Intn(100)
	if n < 30 {
		serverName = "server2"
	}

	requestsMetric.WithLabelValues(serverName, strconv.Itoa(hour)).Inc() //запись данных в метрику
}

func main() {
	prometheus.MustRegister(requestsMetric) // регистрируем метрику

	addr := "127.0.0.1:2000"

	mux := http.NewServeMux()
	mux.Handle("/", &H{})
	mux.Handle("/metrics", promhttp.Handler()) // просмотр метрик

	log.Printf("serve on %s", addr)

	http.ListenAndServe(addr, mux)
}
