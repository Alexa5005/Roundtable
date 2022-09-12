package main

import ("fmt"
				"net/http"
				"html/template"
	  		"database/sql"
				_ "github.com/go-sql-driver/mysql"
)

type Mero struct {
	Id uint16
	Title, Organization, Description, Date1, Date2 string
}

var lineMero = []Mero{}

func index(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html","templates/create.html","templates/MenuOrg.html","templates/line.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/MenuCreate.html","templates/header.html", "templates/footer.html", "templates/create.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}


func save_article(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Println(err)
		return
	}
	// title := r.FormValue("title")
	// organization := r.FormValue("organization")
	// description := r.FormValue("description")
	// date1 := r.FormValue("date1")
	// date2 := r.FormValue("date2")

	db, err := sql.Open("mysql", "kostenko:qwerty123@tcp(172.18.1.25:3306)/Mero")
	if err != nil {
		panic(err)
	}

	// Установка данных
	// insert, err := db.Query(fmt.Sprintf("INSERT INTO CreateEvents (title, organization, description, date1, date2) VALUES ('%s', '%s', '%s', '%s', '%s')", title, organization, description, date1, date2))
	// if err != nil {
	// 	panic(err)
	// }
	// defer insert.Close()

	defer db.Close()

	// //Создание соединения с базой данных
	// 	db, err := sql.Open(typeBase, adressBase)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	//
	// 	// Установка данных
	// 	insert, err := db.Query(fmt.Sprintf("INSERT INTO" + NameTable +"(" + arg1 + ", " + arg2 + ", " + arg3 + ", " + arg4 + ", " + arg5 + ") VALUES ('%s', '%s', '%s', '%s', '%s')", arg6, arg7, arg8, arg9, arg10))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer insert.Close() //остановка записи в базу
	//
	// 	defer db.Close() //Остановка соединения с базой



	http.Redirect(w, r, "/badge", http.StatusSeeOther)
}

func handleFunc()  {
	http.Handle("/static", http.StripPrefix("/static/", http.FileServer(http.Dir("/static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", save_article)
	http.ListenAndServe("172.18.1.25:8080", nil)
}

func main() {
	handleFunc()
}
