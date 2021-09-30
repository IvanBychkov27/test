// https://kodazm.ru/articles/go/https-i-go/

// для генерирование сертификата и приватного ключа с помощью OpenSSL достаточно одной команды:
// openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
// в поле "Common Name (e.g. server FQDN or YOUR name)" - Тут имя сервера (например 127.0.0.1:8081)
// После этого, вы обнаружите два файла "cert.pem" и "key.pem" в той папке, где вы запускали OpenSSL команду.
// Учтите, что эти файлы называются самоподписанным сертификатом. Это значит, что вы можете использовать эти файлы,
// но браузер будет определять соединение как не безопасное.
// Это означает, что сертификат на сервере не подписан доверенным центром сертификации.
// В Firefox нужно кликнуть "I Understand the Risks" и после этого браузер перейдет на сайт.

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет!")
}

func main() {
	fmt.Println("start server: http://127.0.0.1:8080 and https://127.0.0.1:8081")

	http.HandleFunc("/", handler)
	// Запуск HTTPS сервера в отдельной go-рутине
	go http.ListenAndServeTLS("127.0.0.1:8081", "cmd/dsphttps/cert.pem", "cmd/dsphttps/key.pem", nil)
	// Запуск HTTP сервера
	http.ListenAndServe("127.0.0.1:8080", nil)

	fmt.Println("Done...")
}
