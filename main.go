package main

import (
	"fmt"
	"github.com/galuhpradipta/go-mongo/config"
	"github.com/galuhpradipta/go-mongo/src/modules/profile/model"
	"github.com/galuhpradipta/go-mongo/src/modules/profile/repository"
	"time"
)

func main() {

	db, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")

	// saveProfile(profileRepository)
	// updateProfile(profileRepository)
	// deleteProfile(profileRepository)
	// getProfile("U2", profileRepository)
	getProfiles(profileRepository)
}

func saveProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "U2"
	p.FirstName = "Rob"
	p.LastName = "Pike"
	p.Email = "Rob@gmail.com"
	p.Password = "123456"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Save(&p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile saved ..")
	}
}

func updateProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "U1"
	p.FirstName = "Galuh"
	p.LastName = "Pradipta"
	p.Email = "galuh@gmail.com"
	p.Password = "12345678"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Update("U1", &p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile updated ..")
	}
}

func deleteProfile(profileRepository repository.ProfileRepository) {
	err := profileRepository.Delete("U1")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile Deleted ..")
	}
}

func getProfile(id string, profileRepository repository.ProfileRepository) {
	profile, err := profileRepository.FindByID(id)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("=====Profile=====")
	fmt.Println("Firstname : ", profile.FirstName)
	fmt.Println("Lastname: ", profile.LastName)
	fmt.Println("Email : ", profile.Email)
}

func getProfiles(profileRepository repository.ProfileRepository) {
	profiles, err := profileRepository.FindAll()

	if err != nil {
		fmt.Println(err)
	}

	for _, profile := range profiles {
		fmt.Println("-----------------")
		fmt.Println("ID : ", profile.ID)
		fmt.Println("Firstname : ", profile.FirstName)
		fmt.Println("Lastname: ", profile.LastName)
		fmt.Println("Email : ", profile.Email)
	}
}
