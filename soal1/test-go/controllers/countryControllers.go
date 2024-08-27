package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/PashaAkbar/test-go/models"
	"github.com/gin-gonic/gin"
)


func GetCountries(c *gin.Context) {

	resp, err := http.Get("https://citcall.com/test/countries.json")
	if err != nil {
		log.Fatalf("Failed to call API: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	print(string(body))

	var countries []models.Country
	err = json.Unmarshal(body, &countries)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"countries": countries,
	})
}