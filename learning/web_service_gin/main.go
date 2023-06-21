package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Album Model
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// creating and populating array of Album
var albums = []album{
	{
		ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{
		ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{
		ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbum(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	//call BindJSON to bind the recieved JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	//initializing go router
	router := gin.Default()

	//Endpoints
	router.GET("albums", getAlbums)
	router.GET("/album/:id", getAlbum)
	router.POST("/albums", postAlbums)

	//start server
	router.Run("localhost:8000")
}
