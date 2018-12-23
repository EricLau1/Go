package models

import (
	"errors"
	"../utils"
)

var (
	ErrUserNotFound = errors.New("Usuário não encontrado.")
	ErrEncryptPass = errors.New("Bcrypt Error.")
	ErrInvalidPass = errors.New("Senha inválida.")
)

type User struct {
	Id int
	Username string
	Password string
}

func NewUser(username, password string)(bool, error) {

	hash, err := utils.Hash(password)

	if err != nil {

		return false, ErrEncryptPass

	}

	con, err := Connect()

	if err != nil {

		return false, err

	}

	sql := "insert into users (username, password) values (?, ?)"

	stmt, err := con.Prepare( sql )

	if err != nil {

		return false, err

	}

	_, err = stmt.Exec(username, hash)

	if err != nil {

		return false, err

	}

	defer stmt.Close()
	defer con.Close()

	return true, nil

}

func GetUsers()([]User, error) {

	con, err := Connect()

	if err != nil {

		return nil, err

	}

	sql := "select * from users"

	rs, err := con.Query(sql)

	if err != nil {

		return nil, err

	}

	var users []User

	for rs.Next() {

		var user User

		err := rs.Scan(&user.Id, &user.Username, &user.Password)

		if err != nil {

			return nil, err
	
		}

		users = append(users, user)
	}

	defer rs.Close()
	defer con.Close()

	return users, nil

}

func SignIn(username, password string) (User, error) {

	return auth(username, password)

}

func GetUserByUsername(username string) (User, error) {

	con, err := Connect()

	if err != nil {

		return User{}, err

	}

	sql := "select * from users where username = ?"

	stmt, err := con.Prepare(sql)

	if err != nil {

		return User{}, err

	}

	rs, err := stmt.Query( username ) // recebe 1 ou mais argumentos

	if err != nil {

		return User{}, err

	}

	var user = User{Id: -1, Username: "", Password: ""}

	if rs.Next() {

		err := rs.Scan(&user.Id, &user.Username, &user.Password)

		if err != nil {

			return User{}, err
	
		}

	}

	if user.Id <= 0 {

		return user, ErrUserNotFound

	}

	defer rs.Close()
	defer stmt.Close()
	defer con.Close()

	return user, nil

}

func GetUserById(id int) (User, error) {

	con, err := Connect()

	if err != nil {

		return User{}, err

	}

	sql := "select * from users where id = ?"

	rs, err := con.Query(sql, id)

	if err != nil {

		return User{}, err

	}

	var user = User{Id: -1, Username: "", Password: ""}

	if rs.Next() {

		err := rs.Scan(&user.Id, &user.Username, &user.Password)

		if err != nil {

			return User{}, err
	
		}

	}

	if user.Id <= 0 {

		return user, ErrUserNotFound

	}

	defer rs.Close()
	defer con.Close()

	return user, nil

}

func auth(username, password string) (User, error) {

	user, err := GetUserByUsername(username)

	if err != nil {

		return User{}, err

	}

	_, err = utils.VerifyHash([]byte(user.Password), []byte(password) )

	if err != nil {

		return User{}, ErrInvalidPass

	}

	return user, nil

}

