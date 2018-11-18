package routes

import (
	"net/http"

	"../functions"

	"../middleware"
	"../models"
	"../sessions"
	"../utils"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	// esta pagina precisa ser autenticada na sessão para abrir
	r.HandleFunc("/", middleware.AuthRequired(indexGetHandler)).Methods("GET")

	r.HandleFunc("/", middleware.AuthRequired(indexPostHandler)).Methods("POST")

	r.HandleFunc("/login", loginGetHandler).Methods("GET")

	r.HandleFunc("/login", loginPostHandler).Methods("POST")

	r.HandleFunc("/logout", logoutGetHandler).Methods("GET")

	r.HandleFunc("/cadastro", registerGetHandler).Methods("GET")

	r.HandleFunc("/cadastro", registerPostHandler).Methods("POST")

	r.HandleFunc("/{username}", middleware.AuthRequired(userGetHandler)).Methods("GET")

	// mapeando arquivos estáticos (CSS/JS) dentro do diretorio [static]
	fileServer := http.FileServer(http.Dir("./static/"))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	return r

}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {

	updates, err := models.GetAllUpdates()

	if err != nil {

		utils.InternalServerError(w)
		return
	}

	utils.ExecuteTemplate(w, "index.html", struct {
		Title       string
		Updates     []*models.Update
		DisplayForm bool
	}{
		Title:       "All Updates",
		Updates:     updates,
		DisplayForm: true,
	})

}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := sessions.Store.Get(r, "session")
	untypedUserId := session.Values["user_id"]

	userId, ok := untypedUserId.(int64)

	if !ok {

		utils.InternalServerError(w)
		return

	}

	r.ParseForm()

	// retorna o valor do campo do formulario POST com name = comment
	body := r.PostForm.Get("update")

	err := models.PostUpdate(userId, body)

	if err != nil {

		utils.InternalServerError(w)
		return
	}

	// redireciona para a pagina raiz (index.html)
	http.Redirect(w, r, "/", 302)

	// 302 é o código ou status que informa sobre o redirecionamento de uma página ou documento web.

}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {

	utils.ExecuteTemplate(w, "login.html", nil)

}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	user, err := models.AuthenticateUser(username, password)

	if err != nil {

		switch err {

		case models.ErrUserNotFound:
			utils.ExecuteTemplate(w, "login.html", "usuário não encontrado.")
			break

		case models.ErrInvalidLogin:
			utils.ExecuteTemplate(w, "login.html", "Login inválido. Verifique se a senha está correta.")
			break

		default:
			utils.InternalServerError(w)

		}
		return
	}

	userId, err := user.GetUserId()

	if err != nil {

		utils.InternalServerError(w)
		return

	}

	// iniciando a SESSÃO
	session, _ := sessions.Store.Get(r, "session")

	// adicionando chave e valor na sessão
	session.Values["user_id"] = userId

	functions.Log(username + " logado")

	// salvando os dados na sessão
	session.Save(r, w)

	// redireciona para o index
	http.Redirect(w, r, "/", 302)
}

func logoutGetHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := sessions.Store.Get(r, "session")

	// excluindo a sessão
	delete(session.Values, "user_id")

	session.Save(r, w)

	functions.Log("logout!")

	http.Redirect(w, r, "/", 302)

}

func registerGetHandler(w http.ResponseWriter, r *http.Request) {

	utils.ExecuteTemplate(w, "cadastro.html", nil)

}

func registerPostHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	err := models.RegisterUser(username, password)

	if err == models.ErrUsernameTaken {

		utils.ExecuteTemplate(w, "cadastro.html", "username já existe!")
		return

	} else if err != nil {

		utils.InternalServerError(w)
		return

	}

	http.Redirect(w, r, "/login", 302)

}

func userGetHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := sessions.Store.Get(r, "session")
	untypedUserId := session.Values["user_id"]
	currentUserId, ok := untypedUserId.(int64)

	if !ok {

		utils.InternalServerError(w)
		return

	}

	// pega as variáveis passadas pela Url
	vars := mux.Vars(r)

	username := vars["username"]

	user, err := models.GetUserByUsername(username)

	if err != nil {

		utils.InternalServerError(w)
		return

	}

	userId, err := user.GetUserId()

	if err != nil {

		utils.InternalServerError(w)
		return

	}

	updates, err := models.GetUpdates(userId)

	if err != nil {

		utils.InternalServerError(w)
		return

	}

	utils.ExecuteTemplate(w, "index.html", struct {
		Title       string
		Updates     []*models.Update
		DisplayForm bool
	}{
		Title:   username,
		Updates: updates,
		// se o usuario logado na sessão for igual ao usuário digitado na url
		DisplayForm: currentUserId == userId,
	})

}
