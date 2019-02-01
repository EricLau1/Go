package models

type User struct {
	Id uint64
	Name string
}

func NewUser(user User) (bool, error) {
	con := Connect()
	defer con.Close()
	sql := "insert into users (name) values ($1)"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetUserById(id uint64) (User, error) {
	con := Connect()
	defer con.Close()
	sql := "select * from users where id = $1"
	rs, err := con.Query(sql, id)
	if err != nil {
		return User{}, err
	}
	defer rs.Close()
	var user User
	if rs.Next() {
		err := rs.Scan(&user.Id, &user.Name)
		if err != nil {
			return User{}, err
		}
	}
	return user, nil
}