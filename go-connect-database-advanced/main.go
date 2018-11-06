package main

	/* 
		Importar pacote
		
		> go get github.com/go-sql-driver/mysql

	*/
import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const driver = "mysql"
const dbname = "test"

/* Função de Conexão */
func connect() *sql.DB {

	url := "root:@tcp(127.0.0.1:3306)/" + dbname

	con, err := sql.Open(driver, url)

	if err != nil {
		fmt.Println("Erro na conexão do banco de dados")
		os.Exit(1)
	}

	return con

}

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

const table = "users"

func main() {

	var user = User{
		Id:       1,
		Name:     "Jane Doe",
		Email:    "jane@email",
		Password: "123",
	}

	if create(user) {

		fmt.Println("saved successfully!")

		fmt.Println(find(user.Id))

	}

	user.Name = "Jane Doe updated"
	user.Email = "jane@email updated"
	user.Password = "123 updated"

	if update(user) {

		fmt.Println("updated successfully!")

		fmt.Println(find(user.Id))

	}

	//fmt.Println(getAll())

	rowsDeleted := delete(user.Id)

	if rowsDeleted > 0 {

		fmt.Println("deleted rows: ", rowsDeleted)

	}

}

func create(user User) bool {

	con := connect()

	sql := "insert into " + table + "( name, email, password ) values ( ?, ?, ? )"

	stmt, prepareError := con.Prepare(sql)

	var saved bool = true

	if prepareError != nil {

		//panic(prepareError.Error())
		saved = false

	}

	_, execError := stmt.Exec(user.Name, user.Email, user.Password)

	if execError != nil {

		//panic(execError.Error())
		saved = false

	}

	defer stmt.Close()
	defer con.Close()

	return saved

}

func update(user User) bool {

	con := connect()

	sql := "update " + table + " set name = ?, email = ?, password = ? where id = ?"

	stmt, prepareError := con.Prepare(sql)

	var updated bool = true

	if prepareError != nil {

		//panic(prepareError.Error())
		updated = false

	}

	_, execError := stmt.Exec(user.Name, user.Email, user.Password, user.Id)

	if execError != nil {

		//panic(prepareError.Error())
		updated = false

	}

	defer stmt.Close()
	defer con.Close()

	return updated

}

func getAll() []User {

	con := connect()

	sql := "select id, name, email, password from " + table

	rs, err := con.Query(sql)

	if err != nil {

		//panic(err.Error())
		return nil
	}

	var users []User

	for rs.Next() {

		var user User

		err := rs.Scan(&user.Id, &user.Name, &user.Email, &user.Password)

		if err != nil {

			//panic(err.Error())
			return nil

		}

		fmt.Println(user)
		users = append(users, user)

	}

	defer con.Close()
	defer rs.Close()

	return users
}

func find(id int) User {

	con := connect()

	sql := "select id, name, email, password from " + table + " where id = ? limit 1"

	rs, queryError := con.Query(sql, id)

	var user User

	if queryError != nil {

		//	panic( queryError.Error() )
		user.Id = 0
	}

	if rs.Next() {

		scanError := rs.Scan(&user.Id, &user.Name, &user.Email, &user.Password)

		if scanError != nil {

			//panic( scanError.Error() )
			user.Id = 0
		}

	}

	defer con.Close()
	defer rs.Close()

	return user
}

func delete(id int) int64 {

	con := connect()

	sql := "delete from " + table + " where id = ?"

	stmt, prepareError := con.Prepare(sql)

	if prepareError != nil {

		panic(prepareError.Error())
		return 0

	}

	rs, execError := stmt.Exec(id)

	if execError != nil {

		//panic( execError.Error() )
		return 0

	}

	rows, rowsError := rs.RowsAffected()

	if rowsError != nil {

		// panic( rowsError.Error() )
		return 0

	}

	defer con.Close()
	defer stmt.Close()

	return rows

}
