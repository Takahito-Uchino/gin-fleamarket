package main

import (
	"github.com/Takahito-Uchino/gin-fleamarket/infra"
	"github.com/Takahito-Uchino/gin-fleamarket/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}, &models.User{}); err != nil {
		panic("Failed to migrate database")
	}
}
