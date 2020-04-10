package flight24

import (
	"CrawlerTable/lembarsaham"
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"regexp"
	"time"
)

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Crawler(filename string, delay time.Duration, filecodereg string) {
	var row, rowair []string
	var rows [][]string
	var headings []string

	dataAirlineCodee, _ := readLines(filecodereg)

	for _, data := range dataAirlineCodee {
		rowair = nil
		doc := lembarsaham.GetBody("https://www.flightradar24.com/data/aircraft/" + data)

		doc.Find(".row.h-30.p-l-20.p-t-5").Each(func(index int, table *goquery.Selection) {
			table.Find("span.details").Each(func(indexs int, data *goquery.Selection) {
				space := regexp.MustCompile(`\s+`)
				s := space.ReplaceAllString(data.Text(), " ")
				rowair = append(rowair, s)
			})
		})
		fmt.Println(len(rowair))
		doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
			tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
				rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
					space := regexp.MustCompile(`\s+`)
					s := space.ReplaceAllString(tablecell.Text(), " ")
					row = append(row, s)
				})
				if len(row) > 3 {
					row[0] = row[5]
					RemoveIndex(row, 5)
					row[11] = rowair[0]
					row[12] = rowair[1]
					row[13] = rowair[2]
					row = append(row, rowair[3])
					row = append(row, rowair[4])
					row = append(row, rowair[5])
					row = append(row, rowair[6])
					row = append(row, rowair[7])
					row = append(row, rowair[8])
					rows = append(rows, row)
				}
				row = nil
			})
		})
		time.Sleep(delay * time.Second)
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
	headings = []string{"Flight#", "", "Flight Date", "From", "To", "", "STD", "ATD", "STA", "",
		"Arrival Status", "Aircraft", "Airline", "Operator", "Type Code", "Code1", "Code2", "Mode S", "Serial Number(MSN)", "Age"}
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
