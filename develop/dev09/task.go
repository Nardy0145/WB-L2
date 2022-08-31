package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Я реализую утилиту скачивающую index и href-pages с него.
// Воспользуюсь colly парсером.

func main() {
	var url string
	var links []string
	url = "https://teddit.net"
	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		if r.Request.URL.Path == "" {
			os.Mkdir("site", os.ModePerm)
			err := r.Save("site/index.html")
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		if r.Request.URL.String() == url+"/" {
			return
		}
		fmt.Println(r.FileName())
		r.Save("site/" + strings.Replace(r.FileName(), ".unknown", ".html", -1))
	})

	c.OnHTML("a[href]", func(html *colly.HTMLElement) {
		if html.Request.URL.String() == url {
			links = append(links, html.Attr("href"))
		}
	})

	c.Visit(url)
	for _, v := range links {
		c.Visit(url + v)
	}
}
