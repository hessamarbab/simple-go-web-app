package seeders

import (
	"log"
	"os"
	"sika-hessam/database"
	"sika-hessam/dto"
	"sika-hessam/models"
	"sika-hessam/parser"
)

func SeedUsers() {
	users := readFile()
	ch := make(chan dto.User, 10000)
	for i := 0; i < 10; i++ {
		go insertToDB(ch)
	}
	for _, usr := range users {
		ch <- usr
	}
	log.Println("seed Done!")
}

func insertToDB(ch chan dto.User) {
	for {
		usr := <-ch
		var user models.User
		user.Fill(usr)
		dbClient := database.GetDB()
		err := dbClient.Create(&user).Error
		if err != nil {
			log.Fatal("Error creating users")
		}
	}
}

func readFile() []dto.User {
	dat, err := os.ReadFile("./users_data.json")
	if err != nil {
		log.Fatal("Error read file  ./users_data.json")
	}
	users, err := parser.UsersParse(dat)
	if err != nil {
		log.Fatal(err)
	}
	return users
}
