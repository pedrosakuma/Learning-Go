package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     int     "json:id"
	Title  string  "json:title"
	Artist string  "json:artist"
	Price  float64 "json:price"
}

var albums = []album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	initTelemetry()

	router.Use(Telemetry())

	router.GET("/", getAlbuns)
	router.GET("/:id", getAlbunsById)

	router.Run("[::]:8080")
}

func getAlbuns(c *gin.Context) {
	time.Sleep(time.Minute * 5)
	return
	// c.IndentedJSON(http.StatusOK, albums)
}

func getAlbunsById(c *gin.Context) {
	id := c.Param("id")

	idc, err := strconv.Atoi(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id must be int"})

		return
	}

	founded := bynaRySearchAlmbum(idc)

	if founded == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	} else {
		c.IndentedJSON(http.StatusOK, albums[founded])
	}
}

func bynaRySearchAlmbum(id int) int {
	length := len(albums)
	mid := length / 2

	if id > length || id < 1 {
		return -1
	}

	for i := 0; i < length; i++ {
		if albums[mid].ID == id {
			return mid
		} else if albums[mid].ID > id {
			mid = mid / 2
		} else {
			mid = mid + (length / 2)
		}
	}

	return -1
}
