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

func sert(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/sert.html", "templates/header.html", "templates/footer.html","templates/MenuCreate.html","templates/MenuOrg.html","templates/line.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "sert", nil)
}

func FormReg(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/FormReg.html", "templates/header.html", "templates/footer.html","templates/create.html","templates/MenuOrg.html","templates/MenuCreate.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "FormReg", nil)
}

func create(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/MenuCreate.html","templates/header.html", "templates/footer.html", "templates/create.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

func auth(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/header.html", "templates/footer.html", "templates/auth.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "auth", nil)
}

func reg(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/auth.html", "templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "reg", nil)
}

func recovery(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/reg.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "recovery", nil)
}

func konstrBadge(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html","templates/MenuCreate.html","templates/MenuOrg.html","templates/line.html", "templates/reg.html", "templates/konstrBadge.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "konstrBadge", nil)
}

func konstrSert(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html","templates/MenuCreate.html","templates/MenuOrg.html","templates/line.html", "templates/reg.html", "templates/konstrSert.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "konstrSert", nil)
}

func konstrFormReg(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html","templates/MenuCreate.html","templates/MenuOrg.html","templates/line.html", "templates/konstrFormReg.html", "templates/FormReg.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "konstrFormReg", nil)
}


func lineM(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/index.html", "templates/create.html", "templates/footer.html","templates/create.html","templates/MenuOrg.html","templates/line.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/WIT")
	if err != nil {
		panic(err)
	}

	res, err := db.Query("SELECT id, title, organization, description, date1, date2 FROM CreateEvents")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var post Mero
		err = res.Scan(&post.Id, &post.Title, &post.Organization, &post.Description, &post.Date1, &post.Date2)
		if err != nil {
			panic(err)
		}

		lineMero = append(lineMero, post)
	}
	defer db.Close()

	t.ExecuteTemplate(w, "lineM", lineMero)
}

func badge(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/MenuCreate.html", "templates/footer.html",  "templates/badge.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "badge", nil)
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
	http.HandleFunc("/badge", badge)
	http.HandleFunc("/auth", auth)
	http.HandleFunc("/lineM", lineM)
	http.HandleFunc("/reg", reg)
	http.HandleFunc("/recovery", recovery)
	http.HandleFunc("/konstrBadge", konstrBadge)
	http.HandleFunc("/konstrSert", konstrSert)
	http.HandleFunc("/sert", sert)
	http.HandleFunc("/FormReg", FormReg)
	http.HandleFunc("/konstrFormReg", konstrFormReg)
	http.ListenAndServe(":3306", nil)
}

func main() {
	handleFunc()
}
