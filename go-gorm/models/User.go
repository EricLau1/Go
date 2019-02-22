package models


type User struct {
	ID int
	Nickname string
}

func TableUserExists() bool {
	db := Connect()
	defer db.Close()
	return db.HasTable(&User{})
}

func CreateTableUser() {
	db := Connect()
	defer db.Close()
	// tabela criada após a conexão
	// O parametro ID ja é considerado com primary key auto_increment
	db.CreateTable(&User{})
}

func DropTableUser() {
	db := Connect()
	defer db.Close()
	db.DropTableIfExists(&User{})
}

func NewUser(user User) int64 {
	db := Connect()
	defer db.Close()
	rs := db.Create(&user)
	return rs.RowsAffected
}

func FindAndUpdateUser(user User) int64 {
	db := Connect()
	defer db.Close()
	db.Find(&user)
	user.Nickname = "BatGirl"
	rs := db.Save(&user)
	return rs.RowsAffected
}

func TableDeleteUser(id int64) int64 {
	db := Connect()
	defer db.Close()
	rs := db.Table("users").Where("id = ?", id).Delete(&User{})
	return rs.RowsAffected
}

func DeleteUser(id int64) int64 {
	db := Connect()
	defer db.Close()
	rs := db.Where("id = ?", id).Delete(&User{})
	return rs.RowsAffected	
}

func DeleteAllUsers() int64 {
	db := Connect()
	defer db.Close()
	rs := db.Model(&User{}).Delete(&User{})
	return rs.RowsAffected
}