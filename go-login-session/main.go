package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"./utils"
	"./models"
	"./sessions"
)

func main() {

	fmt.Println("Listening port 8080")

	utils.LoadTemplates("templates/*.html")

	r := mux.NewRouter()

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/signup", signupGetHandler).Methods("GET")
	r.HandleFunc("/signup", signupPostHandler).Methods("POST")
	r.HandleFunc("/signin", signinGetHandler).Methods("GET")
	r.HandleFunc("/signin", signinPostHandler).Methods("POST")
	r.HandleFunc("/admin", admin).Methods("GET")
	r.HandleFunc("/logout", logout).Methods("GET")

	fileServer := http.FileServer(http.Dir("./assets"))

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	http.Handle("/", r)

	utils.Shell("http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", r))

}

func index(w http.ResponseWriter, r *http.Request) {

	users, err := models.GetUsers()

	if err != nil {

		utils.InternalServerError(w)
		return

	}

	utils.ExecuteTemplate(w, "index.html", struct{ Users []models.User }{ Users: users, })

}

func signupGetHandler(w http.ResponseWriter, r *http.Request) {

	utils.ExecuteTemplate(w, "signup.html", nil)

}

func signupPostHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	
	_, err := models.NewUser(username, password)

	if err != nil {

		utils.InternalServerError(w)
		return

	}

	http.Redirect(w, r, "/", 302)

}

func signinGetHandler(w http.ResponseWriter, r *http.Request) {

	_, ok := sessions.AuthRequired(r)

	if  ok {
		
		http.Redirect(w, r, "/admin", 302)
		return

	}

	utils.ExecuteTemplate(w, "signin.html",  utils.Alert("", "") )

}

func signinPostHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	user, err := models.SignIn(username, password)
	
	if err != nil {

		switch(err) {

		case models.ErrUserNotFound:
			utils.ExecuteTemplate(w, "signin.html", utils.Alert("Error", "Usuário não encontrado."))
			break
		case models.ErrInvalidPass:
			utils.ExecuteTemplate(w, "signin.html", utils.Alert("Error", "Senha inválida."))
			break		
		default:
			utils.InternalServerError(w)
		}// end switch
		return
	}

	session, _ := sessions.Store.Get(r, "session")
	session.Values["AUTHENTICATE"] = true
	session.Values["USERID"] = user.Id
	session.Save(r, w)
	
	http.Redirect(w, r, "/admin", 302)

}

func admin(w http.ResponseWriter, r *http.Request) {

	session, _ := sessions.Store.Get(r, "session")
	untypedUserId := session.Values["USERID"]

	currentUserId, _ := untypedUserId.(int)

	user, err := models.GetUserById( currentUserId )

	if err != nil {

		//utils.InternalServerError(w)
		http.Redirect(w, r, "/", 302)
		return

	}

	utils.ExecuteTemplate(w, "admin.html", struct{ User models.User }{ User: user, })

}

func logout(w http.ResponseWriter, r *http.Request) {

	session, _ := sessions.Store.Get(r, "session")

	delete(session.Values, "USERID")
	session.Values["AUTHENTICATE"] = false

	session.Save(r, w)

	http.Redirect(w, r, "/", 302)

}