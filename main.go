package main

import (
	"CrawlerTable/bareksa"
	"CrawlerTable/financeyahoo"
	"CrawlerTable/flight24"
	"CrawlerTable/idxtrading"
	"CrawlerTable/lembarsaham"
	"CrawlerTable/opencage"
	"CrawlerTable/tradingview"
	"CrawlerTable/wikiperusahaan"
	"CrawlerTable/wikistasiun"
	"flag"
	"fmt"
	"strconv"
	"time"
)

func main() {
	// code here ...
	crawl := flag.String("crawl", "finyahoo", "crawl dari situs website.\nWebsite yang dapat di crawl:\n"+
		"https://finance.yahoo.com/screener/predefined/ms_basic_materials?count=25&offset=125 (finyahoo)\n"+
		"https://www.bareksa.com/id/saham/sector (bareksa)\n"+
		"https://lembarsaham.com/daftar-emiten/9-sektor-bei (lembarsaham)\n"+
		"https://id.tradingview.com/markets/stocks-indonesia/sectorandindustry-sector/ (tradingview)\n"+
		"https://www.flightradar24.com/ (flight24) \n")
	count := flag.Int("count", 100, "parameter tambahan ketika crawl=finyahoo.\nBerguna untuk jumlah baris yang akan ditampilkam")
	fileName := flag.String("filename", "result.csv", "untuk menyimpan file hasil crawling")
	offset := flag.Int("offset", 100, "parameter tambahan ketika crawl=finyahoo.\nBerguna untuk offset untum menampilkan data")
	delay := flag.Int("delay", 2, "delay saat crawling perhalaman (lembarsaham,financeyahoo,tradingview)\n")
	aircode := flag.String("aircodef", "/data/aircodef.txt", "setingan untuk list file code penerbangan")
	key := flag.String("key", "", "key for opencage")
	datast := flag.String("datast", "data", "list data stasiun")

	flag.Parse()

	if *crawl == "finyahoo" {
		for i := 0; i <= *offset; i += *offset {
			financeyahoo.Crawl(*fileName, *count, i)
			fmt.Println("Delay " + strconv.Itoa(*delay) + " detik")
			time.Sleep(time.Duration(*delay) * time.Second)

		}
	} else if *crawl == "bareksa" {
		bareksa.Crawl(*fileName)
	} else if *crawl == "lembarsaham" {
		lembarsaham.Crawl(*fileName, time.Duration(*delay))
	} else if *crawl == "tradingview" {
		tradingview.Crawl(*fileName, time.Duration(*delay))
	} else if *crawl == "flight24" {
		flight24.Crawler(*fileName, time.Duration(*delay), *aircode)
	} else if *crawl == "stasiun" {
		wikistasiun.Crawl(*fileName)
	} else if *crawl == "opencage" {
		opencage.Crawl(*fileName, *datast, *key)
	} else if *crawl == "wikiperusahaan" {
		wikiperusahaan.Crawl(*fileName)
	}else if *crawl == "idx" {
		idxtrading.Crawl(*fileName)
	}
}
