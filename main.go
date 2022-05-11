package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.webstone.kr"),
	)

	// categories := make([]Category, 0, 260)
	id := 1
	bible := make(map[string]string)
	
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {	
		link := e.Attr("href")
		korTitle := e.Attr("title")
		engTitle := e.Text
		w1 := strings.Split(link, "?")
		if len(w1) == 2 {
			q := strings.Split(w1[1], "=")
			if len(q) == 2 && len(korTitle) > 1 {
				if _, ok := bible[engTitle]; !ok {
					log.Printf("%2d. %-15s %s\n", id, e.Text, q[1])
					bible[e.Text] = q[1]
					id++
				}	
			}
		}
	})

	// start scraping
	c.Visit("http://www.webstone.kr/ap_shop/um_webstone/content_view.php?category=D16168185609")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	// Dump json to stdout
	// enc.Encode(bible)
}