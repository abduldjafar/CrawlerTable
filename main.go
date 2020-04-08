package main

import (
	"CrawlerTable/bareksa"
	"CrawlerTable/financeyahoo"
	"CrawlerTable/lembarsaham"
	"CrawlerTable/tradingview"
	"flag"
)

func main() {
	// code here ...
	crawl := flag.String("crawl", "finyahoo", "crawl dari situs website.\nWebsite yang dapat di crawl:\n"+
		"https://finance.yahoo.com/screener/predefined/ms_basic_materials?count=25&offset=125 (finyahoo)\n"+
		"https://www.bareksa.com/id/saham/sector (bareksa)\n" +
		"https://lembarsaham.com/daftar-emiten/9-sektor-bei (lembarsaham)\n" +
		"https://id.tradingview.com/markets/stocks-indonesia/sectorandindustry-sector/ (tradingview)\n")
	count := flag.Int("count", 100, "parameter tambahan ketika crawl=finyahoo.\nBerguna untuk jumlah baris yang akan ditampilkam")
	fileName := flag.String("filename", "result.csv", "untuk menyimpan file hasil crawling")
	offset := flag.Int("offset", 100, "parameter tambahan ketika crawl=finyahoo.\nBerguna untuk offset untum menampilkan data")
	flag.Parse()

	if *crawl == "finyahoo" {
		for i := 0; i <= *offset; i += *offset {
			financeyahoo.Crawl(*fileName, *count, i)
		}
	} else if *crawl == "bareksa" {
		bareksa.Crawl(*fileName)
	} else if *crawl == "lembarsaham" {
		lembarsaham.Crawl(*fileName)
	}else if *crawl == "tradingview"{
		tradingview.Crawl(*fileName)
	}
}
