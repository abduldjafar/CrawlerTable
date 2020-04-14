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
	"strconv"
	"strings"
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
		lines = append(lines, strings.Trim(scanner.Text(), " "))
	}
	return lines, scanner.Err()
}

func getfile(filename string) (*os.File, error) {
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
	return file, err
}

func Crawler(filename string, delay time.Duration, filecodereg string) {
	var row, rowair []string
	var rowslog [][]string
	var headings []string
	var temp2 []string

	file, err := getfile(filename)
	fileError, err2 := getfile("fligh24_zerotransaktion.csv")
	lembarsaham.CheckError("Error Mesages", err2)

	headingserror := []string{"Aircraft", "Airline", "Operator", "Type", "Code", "Code1", "Code2", "Mode S", "#Aircraft", "#CrawlerDate"}
	headings = []string{"Flight#", "Flight Date", "From", "To", "", "STD", "ATD", "STA", "",
		"Arrival Status", "Aircraft#", "Aircraft", "Airline", "Operator", "Type Code", "Code1", "Code2", "Mode S", "Serial Number(MSN)", "Age"}
	writer := csv.NewWriter(file)
	writererror := csv.NewWriter(fileError)
	defer writer.Flush()
	defer writererror.Flush()

	err = writer.Write(headings)
	err2 = writererror.Write(headingserror)
	lembarsaham.CheckError("Cannot write to file", err)

	dataAirlineCodee, _ := readLines(filecodereg)

	for _, data := range dataAirlineCodee {
		log.Println("Parsing https://www.flightradar24.com/data/aircraft/" + data)
		rowair = nil
		doc := lembarsaham.GetBody("https://www.flightradar24.com/data/aircraft/" + data)

		doc.Find(".row.h-30.p-l-20.p-t-5").Each(func(index int, table *goquery.Selection) {
			table.Find("span.details").Each(func(indexs int, data *goquery.Selection) {
				space := regexp.MustCompile(`\s+`)
				s := space.ReplaceAllString(data.Text(), " ")
				rowair = append(rowair, s)
			})
		})
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
					temp2 = row[11:]
					row = append(row[0:11], strings.ToUpper(data))
					row = append(row, temp2...)
					row[12] = rowair[0]
					RemoveIndex(row, 1)
					row = row[0:20]
					if len(data) > 0 {
						err := writer.Write(row)
						lembarsaham.CheckError("Cannot write to file", err)
					}
					rowslog = append(rowslog, row)

				}
				row = nil
			})
			log.Println("Success Get " + strconv.Itoa(len(rowslog)) + " rows")
			log.Println("=================================================================")
			if len(rowslog) == 0 {
				utc := time.Now().UTC()
				local := utc
				location, err := time.LoadLocation("Asia/Jakarta")
				if err == nil {
					local = local.In(location)
				}
				rowair[7] = data
				rowair[8] = strconv.Itoa(local.Year()) + "-" + local.Month().String() + "-" + strconv.Itoa(local.Day())
				err2 := writererror.Write(rowair)
				fmt.Println(rowair)
				lembarsaham.CheckError("Cannot write to file", err2)
			}
			rowslog = nil
		})
		time.Sleep(delay * time.Second)
	}
}
