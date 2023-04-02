package database

import "task-5-vix-btpns-Ahmad_Yanto/models"

func MigrateDb() {
	db := ConnDb()
	db.AutoMigrate(&models.User{}, &models.Photo{})
}
