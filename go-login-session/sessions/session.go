package sessions

import (
	"net/http"
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("4rm1t4g3256"))

func AuthRequired(r *http.Request) (int, bool) {

	session, _ := Store.Get(r, "session")
	untypedUserId := session.Values["USERID"]

	currentUserId, ok := untypedUserId.(int)

	return currentUserId, ok

}