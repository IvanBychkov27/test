// https://itproger.com/course/golang/4

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name      string
	Age       uint16
	Money     int16
	AvgGrades float64
	Happiness float64
	Hobbies   []string
}

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User Name is: %s. He is %d and he has Money equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func homePage(w http.ResponseWriter, r *http.Request) {
	bob := &User{"Bob", 25, -50, 4.2, 0.8, []string{"Football", "Skate", "Dance"}}
	//bob.setNewName("Alex")
	//fmt.Fprintf(w, bob.getAllInfo())
	//fmt.Fprintf(w, "<B>Main text</B>")

	tmpl, _ := template.ParseFiles("cmd/website/exemple/templates/home_page.html")
	tmpl.Execute(w, bob)

}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page")
}

func handleRequest() {
	fmt.Println("start server... 127.0.0.1:8080")

	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts/", contactsPage)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}

/*

 {{ range .Hobbies }}  // - это цикл
 {{ end }}


 {{ if eq .Name "Bob" }}     // - это равно
    <span>Ну, привет!</span>
    {{ else }}
    <span>А ты кто?</span>
    {{ end }}


    {{ if ne .Name "Bob" }}  // - это не равно
    <span>Ну, привет!</span>
    {{ else }}
    <span>А ты кто?</span>
    {{ end }}

    {{ if gt .Age 20 }}     // - это больше
    <span>ok!</span>
    {{ else }}
    <span>not ok</span>
    {{ end }}


    {{ if lt .Age 20 }}     // - это меньше
    <span>ok!</span>
    {{ else }}
    <span>not ok</span>
    {{ end }}


<h1>Главная страница</h1>

    {{ if eq .Name "Bob" }}
    <span>Ну, привет!</span>
    {{ else }}
    <span>А ты кто?</span>
    {{ end }}

    <ul>
        {{ range .Hobbies }}
            <li><b> {{ . }} </b></li>
        {{ end }}
    </ul>



    <p>Пользователь: {{ .Name }} </p>
    <p>Возраст: {{ .Age }} </p>
    <p>Деньги: {{ .Money }} </p>




*/
