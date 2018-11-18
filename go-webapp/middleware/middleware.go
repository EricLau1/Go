package middleware

import (
	"net/http"
	"webapp/sessions"
)

func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// iniciando a sessão
		session, _ := sessions.Store.Get(r, "session")

		// verifica se o parametro da sessão existe
		_, ok := session.Values["user_id"]

		if !ok || session.Values["user_id"] == 0 {

			http.Redirect(w, r, "/login", 302)
			return

		}

		// fmt.Println("successfully logged in.")

		handler.ServeHTTP(w, r)
	}

}
