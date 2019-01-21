package models

import(
  "fmt"
  "log"
  "database/sql"
  _ "github.com/lib/pq"
)

const(
  DRIVER = "postgres"
  USER   = "postgres"
  PASS   = "@root"
  DBNAME = "test"
)

func Connect() *sql.DB {
  URL := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", USER, PASS, DBNAME)
  con, err := sql.Open(DRIVER, URL)
  if err != nil {
    log.Fatal(err)
  }
  err = con.Ping()
  if err != nil {
    log.Fatal(err)
  }
  return con
}

func InsertImage(image string) (bool, error) {
  con := Connect()
  defer con.Close()
  sql := "insert into images (image) values ($1)"
  stmt, err := con.Prepare(sql)
  if err != nil {
    return false, err
  }
  defer stmt.Close()
  _, err = stmt.Exec(image)
  if err != nil {
    return false, err
  }
  return true, nil
}

func GetImages() ([]string, error) {
  con := Connect()
  defer con.Close()
  sql := "select image from images"
  rs, err := con.Query(sql) 
  if err != nil {
    return nil, err
  }
  defer rs.Close()
  var images []string
  for rs.Next() {
    var image string
    err := rs.Scan(&image)
    if err != nil {
      return nil, err
    }
    images = append(images, image)
  }
  return images, nil
}
