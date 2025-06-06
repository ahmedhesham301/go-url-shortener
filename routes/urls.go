package routes

import (
	"math/rand"
	"net/http"
	"strconv"
	"urlshortener/models"

	"github.com/gin-gonic/gin"
)

func NewUrl(context *gin.Context) {
	var url models.Url
	err := context.ShouldBindJSON(&url)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	url.UrlHash = strconv.Itoa(rand.Int())

	err = url.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save url " + err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "url created"})
}

func GetUrl(context *gin.Context) {
	hash := context.Param("hash")
	url, err := models.GetUrlByHash(hash)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not get url " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, url)
}
