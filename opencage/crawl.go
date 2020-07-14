package opencage

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func Crawl(filename string, datast string, key string) {
	data, err := os.Open(datast)
	if err != nil {
		log.Println(err.Error())
	}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {

		tempat := "stasiun%20" + strings.Trim(scanner.Text(), " ")
		tempat = strings.ReplaceAll(tempat, " ", "%20")
		url := "https://api.opencagedata.com/geocode/v1/json?key=" + key + "&q=" + tempat

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println(err.Error())
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err.Error())
		}
		defer resp.Body.Close()
		log.Println("response " + resp.Status)
		if resp.Status == "400 Bad Request" {
			log.Println(tempat + " Tidak ada Data")
		}
		body, _ := ioutil.ReadAll(resp.Body)

		var result map[string]interface{}
		json.Unmarshal(body, &result)

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

		jsondata, _ := json.Marshal(result)
		if jsondata != nil {
			_, err = file.Write(jsondata)
			if err != nil {
				log.Println(err.Error())
			}
			_, err = file.Write([]byte("\n"))
		}

	}
}
