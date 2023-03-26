package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/siddhant-sar/url-shortner/models"
)

var urlObjects []models.UrlObject
func InitDB() {
	urlObjects = make([]models.UrlObject, 0)
}

func NewUrlObject(context *gin.Context) {
	var urlObject models.UrlObject

	if err := context.ShouldBindJSON(&urlObject); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	urlObject.ShortUrlKey = "xia2"
	urlObject.Ttl = time.Now().UnixMilli()
	urlObjects = append(urlObjects, urlObject)
}

func RedirectUrl(context *gin.Context) {
	
}