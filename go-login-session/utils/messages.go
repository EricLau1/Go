package utils

type Message struct {
	Title string
	Body string
}

func Alert(title, body string) Message {

	return Message{Title: title, Body: body}

}