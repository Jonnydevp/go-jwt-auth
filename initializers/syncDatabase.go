package initializers

import "github.com/go-jwt-auth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})

}
