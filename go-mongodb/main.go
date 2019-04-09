package main

import (
	"fmt"
	"go-mongodb/config"
	"go-mongodb/src/console"
	"go-mongodb/src/modules/profile/model"
	"go-mongodb/src/modules/profile/repository"
	"log"
	"time"
)

func main() {
	fmt.Println("Go MongoDB Tutorial")

	db, err := config.GetMongoDB()

	if err != nil {
		log.Fatal(err)
	}

	db.C("profiles").DropCollection()
	profileRepository := repository.NewProfileRepositoryMongo(db, "profiles")
	saveProfile(profileRepository)
	//updateProfile(profileRepository)
	//deleteProfile(profileRepository)
	fmt.Println("=========== FIND PROFILE ===========")
	getProfile(profileRepository)
	fmt.Println("=========== LIST PROFILES ===========")
	getProfiles(profileRepository)
}

var profiles = model.Profiles{
	model.Profile{ID: "U1", FirstName: "Bruce", LastName: "Wayne", Email: "brucewayne@email.com", Password: "123456", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	model.Profile{ID: "U2", FirstName: "Barbara", LastName: "Gordon", Email: "barbaragordon@email.com", Password: "123456", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	model.Profile{ID: "U3", FirstName: "Gwen", LastName: "Stacy", Email: "gwenstacy@email.com", Password: "123456", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	model.Profile{ID: "U4", FirstName: "Peter", LastName: "Parker", Email: "peterparker@email.com", Password: "123456", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func saveProfile(profileRepository repository.ProfileRepository) {

	for _, p := range profiles {
		err := profileRepository.Save(&p)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("Saved Successfully!")
	}
}

func updateProfile(profileRepository repository.ProfileRepository) {
	p := model.Profile{ID: "U1", FirstName: "Jane", LastName: "Doe", Email: "janedoe@email.com", Password: "updated-password", CreatedAt: time.Now()}
	err := profileRepository.Update(p.ID, &p)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Updated Successfully!")
}

func deleteProfile(profileRepository repository.ProfileRepository) {
	err := profileRepository.Delete("U1")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Deleted Successfully!")
}

func getProfile(profileRepository repository.ProfileRepository) {
	profile, err := profileRepository.FindByID("U3")
	if err != nil {
		log.Println(err)
		return
	}
	console.Pretty(profile)
}

func getProfiles(profileRepository repository.ProfileRepository) {
	profiles, err := profileRepository.FindAll()
	if err != nil {
		log.Println(err)
		return
	}
	console.Pretty(profiles)
}
