package models

type Feedback struct{
	Id uint64
	User User
	Comment string
}

type Rating struct {
	Id uint64
	Feedback Feedback
	Likes int
	Dislikes int
	Reports int
}

func NewFeedback(feedback Feedback) (bool, error) {
	con := Connect()
	defer con.Close()
	tx, err :=	con.Begin()
	if err != nil {
		return false, err
	}
	sql := "insert into feedback (usr, comment) values ($1, $2) returning id"
	var insertId uint64
	{
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			return false, err
		}
		defer stmt.Close()
		err = stmt.QueryRow(feedback.User.Id, feedback.Comment).Scan(&insertId)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}
	sql = "insert into rating (feedback) values ($1)"
	{
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			return false, err
		}
		_, err = stmt.Exec(insertId)
		if err != nil {
			tx.Rollback()
			return false, err
		}
	}
	return true, tx.Commit()
}