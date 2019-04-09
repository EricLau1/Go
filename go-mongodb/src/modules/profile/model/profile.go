package model

import (
	"time"
)

// execute o comando => go fmt src/modules/profile/models/profile.go

type Profile struct {
	ID        string    `bson:"id"`
	FirstName string    `bson:"first_name"`
	LastName  string    `bson:"last_name"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type Profiles []Profile
