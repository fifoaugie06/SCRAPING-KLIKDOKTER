package controllers

import (
	"SCRAPING-INFORMATIKAUMM/structs"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (idb *InDB) GetAllNews(c *gin.Context) {
	var (
		result gin.H
	)

	res, err := http.Get("https://www.klikdokter.com/info-sehat/berita-kesehatan")
	if err != nil {
		result = gin.H{
			"status":  400,
			"message": "Bad Request",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		result = gin.H{
			"status":  400,
			"message": "Bad Request",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		result = gin.H{
			"status":  400,
			"message": "Bad Request",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	rows := make([]structs.News, 0)

	doc.Find(".streamline--articles--iridescent-series").Children().Each(func(i int, sel *goquery.Selection) {
		row := new(structs.News)
		row.Title = sel.Find("h4").Text()
		row.Link, _ = sel.Find("a").Attr("href")
		rows = append(rows, *row)
	})

	log.Println(rows)

	if len(rows) == 0 {
		result = gin.H{
			"status":  404,
			"message": "News is null",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	} else {
		result = gin.H{
			"status":     200,
			"message":    "Success Retrieving News Data",
			"data_count": len(rows),
			"data":       rows,
		}
	}

	c.JSON(http.StatusOK, result)
}
