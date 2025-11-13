package routes

import (
	"net/http"
	"urlshortener/models"
	"urlshortener/monitor"

	"github.com/gin-gonic/gin"
)

func NewUrl(context *gin.Context) {
	var url models.Url
	err := context.ShouldBindJSON(&url)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = url.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save url " + err.Error()})
		return
	}
	// shortUrl := fmt.Sprintf("http://%s/%d", context.Request.Host, url.Id)
	monitor.UrlsCount.Inc()
	context.JSON(http.StatusCreated, gin.H{"message": "url created", "id": url.Id, "long_url": url.URL})
}

func GetUrl(context *gin.Context) {
	id := context.Param("id")
	url, err := models.GetUrlById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "could not get url " + err.Error()})
		return
	}
	context.Redirect(http.StatusFound, url.URL)
	//context.JSON(http.StatusOK, url)
}
