package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swaggo/swag"
)

type Stuff struct {
	Value int `json:"value"`
}

var stuff Stuff

// @Summary Add Stuff
// @Description Add stuff with a value
// @ID add-stuff
// @Accept json
// @Produce json
// @Param stuff body Stuff true "Stuff object"
// @Success 200 {string} string "Added stuff with value: {value}"
// @Router /addStuff [post]
func addStuff(c *gin.Context) {
	var stuff Stuff
	if err := c.ShouldBindJSON(&stuff); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	fmt.Printf("Added stuff with value: %d\n", stuff.Value)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Added stuff with value: %d", stuff.Value)})
}

// @Summary Update Stuff
// @Description Update stuff with a value
// @ID update-stuff
// @Accept json
// @Produce json
// @Param stuff body Stuff true "Stuff object"
// @Success 200 {string} string "Updated stuff with value: {value}"
// @Router /updateStuff [put]
func updateStuff(c *gin.Context) {
	var stuff Stuff
	if err := c.ShouldBindJSON(&stuff); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	fmt.Printf("Updated stuff with value: %d\n", stuff.Value)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Updated stuff with value: %d", stuff.Value)})
}

// @Summary Delete Stuff
// @Description Delete the stuff
// @ID delete-stuff
// @Produce json
// @Success 200 {string} string "Deleted stuff"
// @Router /deleteStuff [delete]
func deleteStuff(c *gin.Context) {
	stuff = Stuff{} // Reset stuff to default value
	fmt.Println("Deleted stuff")
	c.JSON(http.StatusOK, gin.H{"message": "Deleted stuff"})
}

// @title Swagger Example API"
// @version 1.0
// @description This is an example API with Swagger documentation
// @host localhost:8080
// @BasePath /
func jonasApiTest() {
	router := gin.Default()

	// Swagger documentation routes
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/swagger.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, swag.GetSwagger())
	})

	router.POST("/addStuff", addStuff)
	router.PUT("/updateStuff", updateStuff)
	router.DELETE("/deleteStuff", deleteStuff)

	log.Fatal(router.Run(":8080"))
}
