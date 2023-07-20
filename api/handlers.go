package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemResponse struct {
	URL string `json:"url"`
}

type ExternalItemResponse struct {
	Activity      string  `json:"activity"`
	Type          string  `json:"type"`
	Participants  int     `json:"participants"`
	Price         float64 `json:"price"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
	Accessibility float64 `json:"accessibility"`
}

func getActivityWithCustomName(ime string) (*ExternalItemResponse, error) {
	resp, err := http.Get("https://www.boredapi.com/api/activity/")

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var externalResponse ExternalItemResponse
	if err := json.NewDecoder(resp.Body).Decode(&externalResponse); err != nil {
		return nil, err
	}
	externalResponse.Activity += " " + ime

	return &externalResponse, nil
}

func getItemsHandler(c *gin.Context) {
	ime := c.Param("ime")
	itemResponse, err := getActivityWithCustomName(ime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching activity"})
		return
	}
	c.JSON(http.StatusOK, itemResponse.Activity)
}

func InitializeHandlers(router *gin.Engine) {
	router.GET("/items/:ime", func(c *gin.Context) {
		getItemsHandler(c)
	})
}
