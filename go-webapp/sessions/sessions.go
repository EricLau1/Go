package sessions

import (
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("t0p-S3cr3t"))

func SessionOptions(domain string, path string, duration int, httpOnly bool) {

	// definindo as opções da sessão
	Store.Options = &sessions.Options{
		Domain:   domain,   // localhost
		Path:     path,     // (/)
		MaxAge:   duration, // duração: 1 hora
		HttpOnly: httpOnly,
	}

}
