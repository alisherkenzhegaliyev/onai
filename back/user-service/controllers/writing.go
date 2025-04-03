package controllers

import (
	"database/sql"
	"net/http"

	"github.com/alisherkenzhegaliyev/onai/initializers"
	"github.com/alisherkenzhegaliyev/onai/models"
	"github.com/gin-gonic/gin"
)

func WritingRoutes(r *gin.Engine) {
	r.POST("/writing", CreateWriting)
	r.GET("/writing/:id", GetWritingByID)
	r.GET("/writing/user/:user_id", GetWritingByUserID)
	r.PUT("/writing/:id", UpdateWriting)
	r.DELETE("/writing/:id", DeleteWriting)
}

func CreateWriting(c *gin.Context) {
	var input models.Writing
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingID int
	err := initializers.DB.QueryRow("SELECT id FROM writing WHERE user_id = $1", input.UserID).Scan(&existingID)
	if err == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "User already has a writing record"})
		return
	} else if err != sql.ErrNoRows {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	query := `
		INSERT INTO writing (user_id, personal_essay, additional_information_essay)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err = initializers.DB.QueryRow(query, input.UserID, input.PersonalEssay, input.AdditionalInformationEssay).Scan(&input.ID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create writing"})
		return
	}

	c.IndentedJSON(http.StatusCreated, input)
}

func GetWritingByID(c *gin.Context) {
	writingID := c.Param("id")
	var writing models.Writing

	query := `
		SELECT id, user_id, personal_essay, additional_information_essay
		FROM writing
		WHERE id = $1
	`
	err := initializers.DB.QueryRow(query, writingID).Scan(
		&writing.ID,
		&writing.UserID,
		&writing.PersonalEssay,
		&writing.AdditionalInformationEssay,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Writing not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, writing)
}

func GetWritingByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	var writing models.Writing

	query := `
		SELECT id, user_id, personal_essay, additional_information_essay
		FROM writing
		WHERE user_id = $1
	`
	err := initializers.DB.QueryRow(query, userID).Scan(
		&writing.ID,
		&writing.UserID,
		&writing.PersonalEssay,
		&writing.AdditionalInformationEssay,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Writing not found for this user"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, writing)
}

func UpdateWriting(c *gin.Context) {
	writingID := c.Param("id")
	var input models.Writing

	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		UPDATE writing
		SET personal_essay = $1,
		    additional_information_essay = $2
		WHERE id = $3
	`
	result, err := initializers.DB.Exec(query, input.PersonalEssay, input.AdditionalInformationEssay, writingID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update writing"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Writing not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Writing updated successfully"})
}

func DeleteWriting(c *gin.Context) {
	writingID := c.Param("id")

	query := `DELETE FROM writing WHERE id = $1`
	result, err := initializers.DB.Exec(query, writingID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete writing"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Writing not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Writing deleted successfully"})
}
