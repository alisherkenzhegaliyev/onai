package main

import (
	"github.com/alisherkenzhegaliyev/onai/controllers"
	"github.com/alisherkenzhegaliyev/onai/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	controllers.ActivitiesRoutes(r)
	controllers.EducationRoutes(r)
	// others...

	r.Run()
}
