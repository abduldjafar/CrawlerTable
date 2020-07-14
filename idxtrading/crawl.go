package idxtrading

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func Crawl(filename string){

	res, err := http.Get("https://www.idx.co.id/perusahaan-tercatat/profil-perusahaan-tercatat/" +
		"detail-profile-perusahaan-tercatat/?kodeEmiten=AALI#trading")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div.grid-container.block-double").Each(func(index int, html *goquery.Selection) {
		html.Find("dl").Each(func(index int, hasil *goquery.Selection){
			fmt.Println(hasil.Text())
		})
	})
}
