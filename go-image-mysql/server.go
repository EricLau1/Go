package main

import(
	"net/http"
	"html/template"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Image struct {
	Id int
	Name string
	Image string
}

var templates *template.Template

func main() {
	
	fmt.Println("Listening port 3000")

	templates = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", index)

	http.ListenAndServe(":3000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	images, err := GetImages()

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server error"))
		return
	}

	fmt.Println( len(images) )

	templates.ExecuteTemplate(w, "index.html", struct {
				Images []Image
			}{
				Images: images,
			})

}

func Connect() *sql.DB {

	const driver = "mysql"
	const user   = "root"
	const pass   = "@root"
	const dbname = "test"

	url := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", user, pass, dbname)
	
	con, err := sql.Open( driver, url )

	if err != nil {

		log.Fatal(err)

	}

	return con

}



func GetImages() ([]Image, error) {

	con := Connect()

	sql := "select * from images"

	rs, err := con.Query(sql)

	if err != nil {

		return nil, err

	}

	var images []Image

	for rs.Next() {

		var image Image

		err := rs.Scan( &image.Id, &image.Name, &image.Image )

		if err != nil {

			return nil, err

		}

		images = append(images, image)
	
	}

	defer con.Close()
	defer rs.Close()

	return images, nil
}
