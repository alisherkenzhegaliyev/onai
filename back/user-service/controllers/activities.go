package controllers

import (
	"database/sql"
	"net/http"

	"github.com/alisherkenzhegaliyev/onai/initializers"
	"github.com/alisherkenzhegaliyev/onai/models"
	"github.com/gin-gonic/gin"
)

func ActivitiesRoutes(r *gin.Engine) {
	r.POST("/activities-section", CreateActivitiesSection)
	r.POST("/activities", CreateActivities)
	r.GET("/activities/:user_id", GetUserActivities)
	r.GET("/activity/:activity_id", GetActivityByID)
	r.PUT("/activities/:activity_id", UpdateActivity)
	r.DELETE("/activities/:activity_id", DeleteActivity) 
}

func CreateActivitiesSection(c *gin.Context) {
	input := models.ActivitiesSection{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingID uint
	err := initializers.DB.QueryRow(`SELECT id FROM activities_section WHERE user_id = $1`, input.UserID).Scan(&existingID)
	if err == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "User already has an activities section"})
		return
	} else if err != sql.ErrNoRows {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	var newID uint
	err = initializers.DB.QueryRow(`INSERT INTO activities_section (user_id) VALUES ($1) RETURNING id`, input.UserID).Scan(&newID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create activities_section"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"activities_section_id": newID})
}

func CreateActivities(c *gin.Context) {
	input := models.Activity{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := initializers.DB.Exec(`
		INSERT INTO activity (activities_id, activity_type, position_description, organization, activity_description, hours, weeks, intend) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`, input.ActivitiesID, input.ActivityType, input.PositionDescription, input.Organization, input.ActivityDescription, input.Hours, input.Weeks, input.Intend)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create activity"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Activity created!"})
}

func GetUserActivities(c *gin.Context) {
	userID := c.Param("user_id")

	query := `
		SELECT a.id, a.activities_id, a.activity_type, a.position_description, 
		       a.organization, a.activity_description, a.hours, a.weeks, a.intend
		FROM activity a
		JOIN activities_section s ON a.activities_id = s.id
		WHERE s.user_id = $1
	`

	rows, err := initializers.DB.Query(query, userID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve activities"})
		return
	}
	defer rows.Close()

	activities := []models.Activity{}

	for rows.Next() {
		activity := models.Activity{}
		if err := rows.Scan(&activity.ID, &activity.ActivitiesID, &activity.ActivityType, &activity.PositionDescription, &activity.Organization, &activity.ActivityDescription, &activity.Hours, &activity.Weeks, &activity.Intend); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error scanning activity data"})
			return
		}
		activities = append(activities, activity)
	}

	if err = rows.Err(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over activities"})
		return
	}

	c.IndentedJSON(http.StatusOK, activities)
}

func GetActivityByID(c *gin.Context) {
	activityID := c.Param("activity_id")

	query := `
		SELECT id, activities_id, activity_type, position_description, 
		       organization, activity_description, hours, weeks, intend
		FROM activity
		WHERE id = $1
	`

	activity := models.Activity{}

	err := initializers.DB.QueryRow(query, activityID).Scan(
		&activity.ID, &activity.ActivitiesID, &activity.ActivityType,
		&activity.PositionDescription, &activity.Organization, &activity.ActivityDescription,
		&activity.Hours, &activity.Weeks, &activity.Intend,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, activity)
}

func UpdateActivity(c *gin.Context) {
	activityID := c.Param("activity_id")

	var exists bool
	err := initializers.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM activity WHERE id = $1)", activityID).Scan(&exists)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if !exists {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}

	input := models.Activity{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		UPDATE activity 
		SET activity_type = $1, position_description = $2, organization = $3, 
		    activity_description = $4, hours = $5, weeks = $6, intend = $7
		WHERE id = $8
	`

	_, err = initializers.DB.Exec(query,
		input.ActivityType, input.PositionDescription, input.Organization,
		input.ActivityDescription, input.Hours, input.Weeks, input.Intend, activityID,
	)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update activity"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Activity updated successfully!"})
}

func DeleteActivity(c *gin.Context) {
	activityID := c.Param("activity_id")

	query := `DELETE FROM activity WHERE id = $1`

	result, err := initializers.DB.Exec(query, activityID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving delete result"})
		return
	}

	if rowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Activity deleted successfully"})
}