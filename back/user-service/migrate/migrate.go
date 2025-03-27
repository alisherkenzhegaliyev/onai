package migrate

import (
	"github.com/alisherkenzhegaliyev/onai/initializers"
	"github.com/alisherkenzhegaliyev/onai/models"
)

func Migrate() {
	initializers.DB.AutoMigrate(&models.Profile{}, &models.Language{})
	// other models
	// ...
}
