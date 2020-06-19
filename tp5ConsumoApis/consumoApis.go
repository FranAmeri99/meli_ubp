package tp5_consumo_apis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type structura2 struct  {
	Name               string `json:"name"`
	CountryId          string `json:"country_id"`
	Currencies[] struct{
		Id     string
		Symbol string
	} `json:"currencies"`
}

func ConsumoApi(){
	run := gin.Default()
	resp, error := http.Get("https://api.mercadolibre.com/sites/MLA")
	if error != nil {
		log.Fatal(error)
	}
	var resultado structura2
	error = json.NewDecoder(resp.Body).Decode(&resultado)
	if error != nil {
		log.Fatal(error)
	}
	run.GET("api.mercadolibre.com/sites/MLA", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"resultado": resultado})
	})
	run.Run(":3000")
}

