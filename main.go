package main

import (
	"LangAssist/wordArch"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var m = map[string]wordArch.WordArch{
	"meal":   wordArch.WordArch{WordEn: "Meal", Word: "le repas", Lang: "Francias", Categ: "Food"},
	"soup":   wordArch.WordArch{WordEn: "Soup", Word: "la soupe", Lang: "Francias", Categ: "Food"},
	"salade": wordArch.WordArch{WordEn: "Salade", Word: "la salade", Lang: "Francias", Categ: "Food"},
}

func loadWordsOn(c *gin.Context) {
	payload := exportToArray(m)
	fmt.Println(payload)
	c.HTML(http.StatusOK, "home.html", gin.H{
		"wordArr": payload,
	})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/home", loadWordsOn)
	saveWordTable(m)
	lMap := loadWordTable()
	fmt.Println(lMap)
	router.Run("localhost:8080")

}
