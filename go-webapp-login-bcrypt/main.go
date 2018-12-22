package main

import(
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"html/template"
	"github.com/gorilla/mux"
	"./models"
)

var templates *template.Template

func load(w http.ResponseWriter, tmpl string, data interface{} ) {

	templates.ExecuteTemplate(w, tmpl, data)

}

func shell(url string) error {
	
	cmd := "xdg-open"
	args := []string{url}

	return exec.Command(cmd, args...).Start()

}

func internalError(w http.ResponseWriter) {

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error..."))

}

func main() {

	fmt.Println("Listening port 8080")

	templates = template.Must(template.ParseGlob("templates/*.html"))

	r := mux.NewRouter()

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/signup", signupGetHandler).Methods("GET")
	r.HandleFunc("/signup", signupPostHandler).Methods("POST")
	r.HandleFunc("/signin", signinGetHandler).Methods("GET")
	r.HandleFunc("/signin", signinPostHandler).Methods("POST")

	http.Handle("/", r)

	shell("http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func index(w http.ResponseWriter, r *http.Request) {

	users, err := models.GetUsers()
	
	if err != nil {

		internalError(w)
		return

	}

	load(w, "index.html", struct{
		Users []models.User
	}{
		Users: users,
	})

}

func signupGetHandler(w http.ResponseWriter, r *http.Request) {

	load(w, "signup.html", nil)

}

func signupPostHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	name := r.PostForm.Get("name")
	password := r.PostForm.Get("password")

	_, err := models.NewUser(name, password)

	if err != nil {

		switch(err) {

		case models.ErrHashPassword:
			w.Write([]byte("Bcrypt Error."))
			break
		default:
			internalError(w)

		}// end switch
		return
	}

	http.Redirect(w, r, "/", 301)

}

func signinGetHandler(w http.ResponseWriter, r *http.Request) {

	load(w, "signin.html", struct{ Message string }{ Message: "", })

}

func signinPostHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	name := r.PostForm.Get("name")
	password := r.PostForm.Get("password")

	fmt.Printf("POST: NAME: %s; PASSWORD: %s\r\n", name, password)

	user, err := models.Auth(name, password)

	if err != nil {

		switch(err) {

		case models.ErrUserNotFound:
			load(w, "signin.html", struct{
				Message string
			}{	
				Message: "Usuário não encontrado",
			})
			break
		
		case models.ErrHashPassword:
			w.Write([]byte("Bcrypt Error."))
			break

		default:
			internalError(w)

		}// end switch

		return
	}

	fmt.Println(user)

	w.Write([]byte("autenticado com sucesso!"))

}