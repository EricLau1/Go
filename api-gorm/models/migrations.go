package models

func AutoMigrations() {
	db := Connect()
	defer db.Close()
	db.DropTableIfExists(&Feedback{}, &User{})
	db.Debug().AutoMigrate(&User{}, &Feedback{})
	db.Model(&Feedback{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade")
	// o primeiro parametro referencia a chave estrangeira,
	// o segundo    ||         ||      a chave primaria da tabela da chave estrangeira
}