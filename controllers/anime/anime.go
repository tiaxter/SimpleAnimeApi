package controllers

import (
  "main/utils/db"
  "main/models"
  AnimeFormRequests "main/form-requests/anime"

	"github.com/gin-gonic/gin"
	"net/http"
)

func All(c *gin.Context) {
  animes := []models.Anime{}

  DB := db.GetDbConnection()
  DB.Find(&animes)
  
	c.SecureJSON(http.StatusOK, animes)
}

func Get(c *gin.Context) {
  // Get the param
  id := c.Param("id")
  // Create empty Anime Model Instance
  anime := models.Anime{}
  // Get DB Connection
  DB := db.GetDbConnection()
  // Get db value and save it inside the Anime Model
  result := DB.Where("id = ?", id).First(&anime)
  // If there was an error then abort
  if result.Error != nil {
    c.AbortWithStatus(http.StatusNotFound)
    return
  }
  // Return the found data
	c.SecureJSON(http.StatusOK, anime)
}

func Store(c *gin.Context) {
  // Get the data from the request
  var json AnimeFormRequests.StoreRequest
  if err := c.ShouldBindJSON(&json); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  // Connect to the DB
  DB := db.GetDbConnection()
  // Create the model instance with data from the request
  anime := &models.Anime{Name: json.Name, Description: json.Description, CoverImage: json.CoverImage, Status: json.Status, PublishedYear: json.PublishedYear}
  // Save the model inside the DB
  err := DB.Create(anime)
  // If there was any error abort
  if err.Error != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
  // Return the inserted data
	c.SecureJSON(http.StatusCreated, anime)

}