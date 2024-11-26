package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func downloadCSV(wg *sync.WaitGroup, urls []string, ch chan []byte) {
	defer wg.Done()
	defer close(ch)

	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			log.Println("cannot download CSV: ", err)
			continue
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			log.Println("cannot read content: ", err)
			continue
		}
		resp.Body.Close()
		ch <- b
	}
}

func insertRecords(records []string) {
	fmt.Println("Inserting records:", records)
}

func main() {
	urls := []string{
		"http://my-server.com/data01.csv",
		"http://my-server.com/data02.csv",
		"http://my-server.com/data03.csv",
	}

	ch := make(chan []byte)

	var wg sync.WaitGroup
	wg.Add(1)
	go downloadCSV(&wg, urls, ch)

	for b := range ch {
		r := csv.NewReader(bytes.NewReader(b))
		for {
			records, err := r.Read()
			if err != nil {
				log.Fatal(err)
			}
			insertRecords(records)
		}
	}

	wg.Wait()
}
