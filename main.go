package main

import (
	"LangAssist/wordArch"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func loadWordsOn(c *gin.Context) {
	m := loadWordTable()
	fmt.Println(m)
	payload := exportToArray(m)
	fmt.Println(payload)
	c.HTML(http.StatusOK, "home.html", gin.H{
		"wordArr": payload,
	})
}
func appendWrdOn(c *gin.Context) {
	wrdEn := c.PostForm("WrdEn")
	wrdFr := c.PostForm("WrdFr")
	categ := c.PostForm("Categ")
	lang := c.PostForm("Lang")

	myWords := loadWordTable()
	myWords[string(wrdEn)] = wordArch.WordArch{WordEn: wrdEn, Word: wrdFr, Lang: lang, Categ: categ}
	saveWordTable(myWords)
	c.Redirect(http.StatusFound, "/home")

}
func loadLessonOn(c *gin.Context) {
	c.HTML(http.StatusOK, "lesson.html", gin.H{})

}

func loadDeleteOn(c *gin.Context) {
	m := loadWordTable()
	fmt.Println(m)
	payload := exportToArray(m)
	fmt.Println(payload)
	c.HTML(http.StatusOK, "deleteWord.html", gin.H{
		"wordArr": payload,
	})

}

func deleteWordOn(c *gin.Context) {
	wrdEn := c.PostForm("WordEn")
	myWords := loadWordTable()
	var newMap = make(map[string]wordArch.WordArch)
	for index, element := range myWords {
		if index != wrdEn {
			newMap[index] = element
		}
	}
	saveWordTable(newMap)
	c.Redirect(http.StatusFound, "/deleteWrd")
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.LoadHTMLGlob("templates/*")
	router.GET("/home", loadWordsOn)
	router.GET("/lesson", loadLessonOn)
	router.GET("/deleteWrd", loadDeleteOn)
	router.POST("/deleteWrd", deleteWordOn)
	router.POST("/lesson", appendWrdOn)
	router.Run("localhost:8080")

}
