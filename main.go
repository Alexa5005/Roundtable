package main

import ("fmt"
				"net/http"
				"html/template"
	  		"database/sql"
				_ "github.com/go-sql-driver/mysql"
)

// type Mero struct {
// 	Id uint16
// 	Title, Organization, Description, Date1, Date2 string
// }
//
// // var lineMero = []Mero{}

func current(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/current.html", "templates/header.html", "templates/footer.html","templates/create.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "current", nil)
}

func success(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/success.html", "templates/header.html", "templates/footer.html","templates/create.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "success", nil)
}

<<<<<<< HEAD
func roundtable(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/roundtable.html", "templates/header.html", "templates/footer.html","templates/create.html")
=======
func current(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/current.html", "templates/header.html", "templates/footer.html", "templates/roundtable.html")
>>>>>>> 8c96117d7de81fbb62baf7455e4390e824e32e75

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

<<<<<<< HEAD
	t.ExecuteTemplate(w, "roundtable", nil)
=======
	t.ExecuteTemplate(w, "current", nil)
>>>>>>> 8c96117d7de81fbb62baf7455e4390e824e32e75
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

	lastName := r.FormValue("lastName")
	firstName := r.FormValue("firstName")
	middleName := r.FormValue("middleName")
	organization := r.FormValue("organization")
	position := r.FormValue("position")
	phone := r.FormValue("phone")
	email := r.FormValue("email")



	db, err := sql.Open("mysql", "kostenko:qwerty123@tcp(172.18.1.25:3306)/Roundtable")
	if err != nil {
		panic(err)
	}

	// Установка данных
	insert, err := db.Query(fmt.Sprintf("INSERT INTO Roundtable (lastName, firstName, middleName, organization, position, phone, email) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s')", lastName, firstName, middleName, organization, position, phone, email))
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	defer db.Close()

	http.Redirect(w, r, "/success", http.StatusSeeOther)
}

func handleFunc()  {
	http.Handle("/static", http.StripPrefix("/static/", http.FileServer(http.Dir("/static/"))))
	http.HandleFunc("/", current)
<<<<<<< HEAD
	http.HandleFunc("/current", current)
=======
>>>>>>> 8c96117d7de81fbb62baf7455e4390e824e32e75
	http.HandleFunc("/create", create)
	http.HandleFunc("/success", success)
	http.HandleFunc("/save_article", save_article)
	http.HandleFunc("/roundtable", roundtable)
	http.ListenAndServe("172.18.1.25:8080", nil)

}

func main() {
	handleFunc()
}
