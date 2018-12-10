package main

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"bufio"
    "encoding/base64"
    "fmt"
    "io/ioutil"
	"os"
	"log"
)

func main() {

	nameFile := "cat.jpg"

    // Open file on disk.
    file, _ := os.Open( "images/" + nameFile )

    // Read entire JPG into byte slice.
    reader := bufio.NewReader(file)
    content, _ := ioutil.ReadAll(reader)

    // Encode as base64.
    image := base64.StdEncoding.EncodeToString(content)

	fmt.Println( len(image) )

	_, erro := NewImage( nameFile , image )

	if erro != nil {

		log.Fatal( erro )

	}

	fmt.Println("Imagem salva com sucesso!")

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

func NewImage(name, image string) (bool, error) {

	con := Connect()

	sql := "insert into images (name, image) values (? , ?)"

	stmt, err := con.Prepare(sql)

	if err != nil {
		
		return false, err

	}

	_, err = stmt.Exec(name, image)

	if err != nil {

		return false, err

	}

	defer con.Close()
	defer stmt.Close()

	return true, nil

}