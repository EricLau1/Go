package main

import (
  "net/http"
  "html/template"
  "github.com/gorilla/mux"
  "encoding/base64"
  "io/ioutil"
  "fmt"
  "log"
  "./models"
)

var templates *template.Template

func main() {
  
  templates = template.Must(template.ParseGlob("templates/*.html"))
  
  router := mux.NewRouter()

  router.HandleFunc("/", handlerGet).Methods("GET")
  router.HandleFunc("/", handlerPost).Methods("POST")

  http.Handle("/", router)

  log.Fatal(http.ListenAndServe(":8080", nil))
  
}

func handlerGet(w http.ResponseWriter, r *http.Request) {
  images, err := models.GetImages()
  if err != nil {
    fmt.Println(err)
  }
  var i int = 0
  for i < len(images) {
    fmt.Printf("Image %d, Tamanho: %d\n", (i + 1), len(images[i]))
    i++
  }
  templates.ExecuteTemplate(w, "index.html", struct{
    Images []string
  }{
    Images: images,
  })
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  file, _, err := r.FormFile("image")
  if err != nil {
    log.Println(err)
    http.Error(w, "Error uploading file", http.StatusInternalServerError)
  }
  defer file.Close()
  byteSize, _ := ioutil.ReadAll(file)
  img := base64.StdEncoding.EncodeToString(byteSize)
  _, err = models.InsertImage(img) 
  if err != nil {
		log.Println(err)
    http.Error(w, "Erro ao salvar imagem.", http.StatusInternalServerError)
	}
  fmt.Println("Imagem salva no banco de dados!")
  http.Redirect(w, r, "/", 302)
}
