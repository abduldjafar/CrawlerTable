package flight24

import (
	"CrawlerTable/lembarsaham"
	"encoding/csv"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"regexp"
	"time"
)

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func Crawler(filename string, delay time.Duration) {
	var row []string
	var rows [][]string
	var headings []string
	dataAirlineCodee := []string{"n773ck", "EW-511TQ", "N773CK", "TC-LJS", "A7-BFT", "D-AALN"}
	for _, data := range dataAirlineCodee {
		doc := lembarsaham.GetBody("https://www.flightradar24.com/data/aircraft/" + data)
		doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
			tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
				rowhtml.Find("th").Each(func(indexth int, tableheading *goquery.Selection) {

				})
				rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
					space := regexp.MustCompile(`\s+`)
					s := space.ReplaceAllString(tablecell.Text(), " ")
					row = append(row, s)
				})
				if len(row) > 3 {
					row[0] = row[5]
					RemoveIndex(row, 5)
					RemoveIndex(row, 13)
					RemoveIndex(row, 12)
					RemoveIndex(row, 11)
					rows = append(rows, row)
				}
				row = nil
			})
		})
	}

	file, err := os.Open(filename)
	//file, err := os.OpenFile(filename,os.O_APPEND|os.O_RDWR,os.ModeAppend)
	if err != nil {
		file, err = os.Create(filename)
		if err != nil {
			log.Fatal(err.Error())
		}
	} else {
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_RDWR, os.ModeAppend)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	defer file.Close()
	headings = []string{"NO PENERBANGAN", "", "TANGGAL PENERBANGAN", "DARI", "MENUJU", "", "STD", "ATD", "STA", "", "STATUS KEDATANGAN", "", "", ""}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(headings)
	lembarsaham.CheckError("Cannot write to file", err)

	for _, data := range rows {
		if len(data) > 0 {
			err := writer.Write(data)
			lembarsaham.CheckError("Cannot write to file", err)
		}
	}
}
