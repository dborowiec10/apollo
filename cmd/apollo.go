package main

import (
	"log"
	"time"

	"github.com/amirgamil/apollo/pkg/apollo"
	"github.com/amirgamil/apollo/pkg/apollo/backend"
)

func main() {
	backend.InitializeFilesAndData()
	// two days in miliseconds
	// once every three days, takes all the records, pulls from the data sources,
	ticker := time.NewTicker(3 * 24 * time.Hour)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				log.Println("Refreshing inverted index at: ", t)
				backend.RefreshInvertedIndex()
			}
		}
	}()
	//server and the pipeline should run on concurrent threads, called regularly, for now manually do it
	//start the server on a concurrent thread so that when we need to refresh the inverted index, this happens on
	//different threads
	// backend.RefreshInvertedIndex()
	apollo.Start()

}
