package models

import(
	"errors"
	"../utils"
	"fmt"
)

var (
	ErrUserNotFound = errors.New("Usuário não encontrado")
	ErrHashPassword = errors.New("Erro ao criptografar senha.")
)

type User struct {
	Id int
	Name string
	Password string
}

func NewUser(name , password string) (bool, error) {

	con, err := Connect();

	if err != nil {

		return false, err

	}

	hash, err := utils.Hash(password)
	
	if err != nil {

		return false, ErrHashPassword

	}

	sql := "insert into users (name, password) values (?, ?)"

	stmt, err := con.Prepare(sql)

	if err != nil {

		return false, err

	}

	_, err = stmt.Exec( name, hash )

	if err != nil {

		return false, err

	}

	defer stmt.Close()
	defer con.Close()

	return true, nil

}


func GetUsers() ([]User, error) {

	con, err := Connect();

	if err != nil {

		return nil, err

	}

	sql := "select * from users"

	rs, err := con.Query( sql )

	if err != nil {

		return nil, err

	}

	var users []User

	for rs.Next() {

		var user User

		err := rs.Scan( &user.Id, &user.Name, &user.Password )

		if err != nil {

			return nil, err

		}
 
		users = append(users, user)

	}

	defer rs.Close()
	defer con.Close()

	return users, nil
}

func Auth(name, password string) (User, error) {

	con, err := Connect()

	if err != nil {

		return User{}, err

	}

	sql := "select * from users where name = ? limit 1"

	stmt, err := con.Prepare( sql )

	if err != nil {

		return User{}, err

	}

	rs, err := stmt.Query(name) // aceita 1 ou mais argumentos

	if err != nil {

		return User{}, err

	}

	var user = User{Id:-1, Name: "", Password:""}

	if rs.Next() {

		err := rs.Scan(&user.Id, &user.Name, &user.Password)

		if err != nil {

			return User{}, nil

		}

	}

	if user.Id <= 0 {

		return User{}, ErrUserNotFound

	}

	err = utils.Verify( []byte(user.Password) , []byte(password) )

	if err != nil {

		return User{}, ErrHashPassword

	}


	defer stmt.Close()
	defer rs.Close()
	defer con.Close()

	return user, nil

}