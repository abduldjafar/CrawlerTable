package financeyahoo

import (
	"encoding/csv"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func Crawl(filename string, count int, offset int) {
	var headings, row []string
	var rows [][]string
	var file *os.File
	counts := strconv.Itoa(count)
	offsets := strconv.Itoa(offset)

	res, err := http.Get("https://finance.yahoo.com/screener/predefined/ms_basic_materials?count=" + counts + "&offset=" + offsets)
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

	// Find each table
	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("th").Each(func(indexth int, tableheading *goquery.Selection) {
				headings = append(headings, tableheading.Text())
			})
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				row = append(row, strings.Trim(strings.Trim(tablecell.Text(), "\n\n"), "\t"))
			})
			rows = append(rows, row)
			row = nil
		})
	})
	file, err = os.Open(filename)
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

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if offset == 0 {
		err = writer.Write(headings)
		checkError("Cannot write to file", err)
	}
	for _, data := range rows {
		if len(data) > 0 {
			err := writer.Write(data)
			checkError("Cannot write to file", err)
		}
	}
}
