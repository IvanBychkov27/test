// https://golangs.org/url-query-strings
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Обработчик главной страницы
func home(w http.ResponseWriter, r *http.Request) {
	// Проверяется, если текущий путь URL запроса точно совпадает с шаблоном "/".
	// Если нет, то вызывается функция http.NotFound() для возвращения клиенту ошибки 404.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Привет из Snippetbox"))
}

// Обработчик для отображения содержимого заметки
// Пример: http://127.0.0.1:4000/snippet?id=123
func showSnippet(w http.ResponseWriter, r *http.Request) {
	idURL := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idURL)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
	//w.Write([]byte("Отображение заметки..."))
}

// Обработчик для создания новой заметки
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Используем r.Method для проверки, использует ли запрос метод POST или нет. Обратите внимание,
	// что http.MethodPost является строкой и содержит текст "POST".
	if r.Method != http.MethodPost {
		// Используем метод Header().Set() для добавления заголовка 'Allow: POST' в
		// карту HTTP-заголовков. Первый параметр - название заголовка, а
		// второй параметр - значение заголовка.
		w.Header().Set("Allow", http.MethodPost)
		w.Header()["Date"] = nil

		// Вызываем метод w.WriteHeader() для возвращения статус-кода 405
		// и вызывается метод w.Write() для возвращения тела-ответа с текстом "Метод запрещен".
		//w.WriteHeader(405)
		//w.Write([]byte("GET-Метод запрещен!"))

		// Используем функцию http.Error() для отправки кода состояния 405 с соответствующим сообщением.
		http.Error(w, "Метод запрещен!", 405)
		return

		// https://golangs.org/customizing-http-headers
		// Управление HTTP заголовками в Go

		// Устанавливаем новый заголовок управления кешем. Если заголовок "Cache-Control" уже указан
		// то он будет переписан.
		//w.Header().Set("Cache-Control", "public, max-age=31536000")

		// Метод Add() добавляет новый заголовок "Cache-Control" и может
		// вызываться несколько раз.
		//w.Header().Add("Cache-Control", "public")
		//w.Header().Add("Cache-Control", "max-age=31536000")

		// Удаляем все значения из заголовка "Cache-Control".
		//w.Header().Del("Cache-Control")

		//Получаем первое значение из заголовка "Cache-Control".
		//w.Header().Get("Cache-Control")

	}

	w.Write([]byte("Форма для создания новой заметки..."))
}

func main() {
	// Регистрируем два новых обработчика и соответствующие URL-шаблоны в
	// маршрутизаторе servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Println("error: ", err.Error())
	}
}
