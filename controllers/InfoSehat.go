package controllers

import (
	"SCRAPING-KLIKDOKTER/structs"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInfoSehat(c *gin.Context) {
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

	rows := make([]structs.InfoSehat, 0)

	doc.Find(".streamline--articles--iridescent-series").Children().Each(func(i int, sel *goquery.Selection) {
		row := new(structs.InfoSehat)
		row.Title = sel.Find("h4").Text()
		row.Image, _ = sel.Find("img").Attr("src")
		row.LinkDetail, _ = sel.Find("a").Attr("href")
		rows = append(rows, *row)
	})

	if len(rows) == 0 {
		result = gin.H{
			"status":  404,
			"message": "InfoSehat is null",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	} else {
		result = gin.H{
			"Status":    200,
			"Message":   "Success Retrieving InfoSehat",
			"DataCount": len(rows),
			"Data":      rows,
		}
	}

	c.JSON(http.StatusOK, result)
}
