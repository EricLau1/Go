package repository

import (
	"go-mongodb/src/modules/profile/model"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type profileRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewProfileRepositoryMongo(db *mgo.Database, collection string) *profileRepositoryMongo {
	return &profileRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *profileRepositoryMongo) Save(profile *model.Profile) error {
	return r.db.C(r.collection).Insert(profile)
}

func (r *profileRepositoryMongo) Update(id string, profile *model.Profile) error {
	profile.UpdatedAt = time.Now()
	return r.db.C(r.collection).Update(bson.M{"id": id}, profile)
}

func (r *profileRepositoryMongo) Delete(id string) error {
	return r.db.C(r.collection).Remove(bson.M{"id": id})
}

func (r *profileRepositoryMongo) FindByID(id string) (*model.Profile, error) {
	profile := model.Profile{}
	err := r.db.C(r.collection).Find(bson.M{"id": id}).One(&profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *profileRepositoryMongo) FindAll() (model.Profiles, error) {
	profiles := model.Profiles{}
	err := r.db.C(r.collection).Find(bson.M{}).All(&profiles)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}
