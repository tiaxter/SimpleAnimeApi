package main

import (
  AnimeController "main/controllers/anime"
  "main/utils/db"

	"github.com/gin-gonic/gin"
)

func main() {
  // Migrate the schemas
  db.MigrateSchemas()
  // Start the http server
  router := gin.Default()

  // Anime Routes
  animeRoute := router.Group("/anime")
  {
    animeRoute.GET("", AnimeController.All)
    animeRoute.GET("/:id", AnimeController.Get)
    animeRoute.POST("", AnimeController.Store)
  }

  router.Run()
}
