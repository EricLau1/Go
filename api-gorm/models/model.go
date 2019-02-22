package models

import "github.com/jinzhu/gorm"

const (
	USERS     = "users"
	FEEDBACKS = "feedbacks"
)

func GetAll(table string) interface{} {
	db := Connect()
	defer db.Close()
	switch table {
	case USERS:
		return db.Order("id asc").Find(&[]User{}).Value
	case FEEDBACKS:
		//http://doc.gorm.io/associations.html#has-many
		var feedbacks []Feedback
		db.Order("id asc").Find(&feedbacks)
		for i, _ := range feedbacks {
			db.Model(feedbacks[i]).Related(&feedbacks[i].User)
		}
		return feedbacks
	}
	return nil
}

func GetById(table, id string) interface{} {
	db := Connect()
	defer db.Close()
	switch table {
	case USERS:
		return db.Where("id = ?", id).First(&User{}).Value
	case FEEDBACKS:
		return db.Where("id = ?", id).First(&Feedback{}).Value
	}
	return nil
}

func Delete(table, id string) (int64, error) {
	db := Connect()
	defer db.Close()
	var rs *gorm.DB
	switch table {
	case USERS:
		rs =  db.Where("id = ?", id).Delete(&User{})
		break
	case FEEDBACKS:
		rs = db.Where("id = ?", id).Delete(&Feedback{})
		break
	default:
		return 0, nil
	}
	return rs.RowsAffected, rs.Error
}