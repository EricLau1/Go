package models

type User struct {
	Id        uint32     `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string     `gorm:"type:varchar(20);unique_index" json:"nickname"`
	Feedbacks []Feedback `gorm:"ForeignKey:UserId" json:"feedbacks"` 
}

func NewUser(user User) (int64, error) {
	db := Connect()
	defer db.Close()
	rs := db.Create(&user)
	return rs.RowsAffected, rs.Error
}

func UpdateUser(user User) (int64, error) {
	db := Connect()
	defer db.Close()
	rs := db.Where("id = ?", user.Id).Find(&User{}).Update("nickname", user.Nickname)
	return rs.RowsAffected, rs.Error
}