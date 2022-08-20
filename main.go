package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ablbum struct {
	Id     string  `json:id`
	Title  string  `json:title`
	Artist string  `json:artist`
	Price  float64 `json:price`
}

var albums = []ablbum{
	{Id: "1", Title: "Train of thought", Artist: "Dream Theater", Price: 9.90},
	{Id: "2", Title: "Nightfall In Middlewarth", Artist: "Blind Guardian", Price: 7.38},
	{Id: "3", Title: "Brave New World", Artist: "Iron Maide", Price: 6.99},
	{Id: "4", Title: "Holy Land", Artist: "Angra", Price: 3.75},
	{Id: "5", Title: "Nevermind", Artist: "Nirvana", Price: 2.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func addAlbum(c *gin.Context) {
	var newAlbum ablbum

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", addAlbum)

	router.Run("localhost:5001")
}
