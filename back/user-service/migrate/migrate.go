package migrate

import (
	"github.com/alisherkenzhegaliyev/onai/initializers"
	"github.com/alisherkenzhegaliyev/onai/models"
)

func Migrate() {
	initializers.DB.AutoMigrate(&models.Profile{}, &models.Language{})
	initializers.DB.AutoMigrate(&models.Family{}, &models.Sibling{})
	initializers.DB.AutoMigrate(&models.Testing{}, &models.AcademicTest{})
  	initializers.DB.AutoMigrate(&models.Education{}, &models.College{}, &models.Course{}, &models.Honor{})
}
