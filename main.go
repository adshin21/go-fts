package main

import (
	"flag"
	"log"
	"time"

	"github.com/adshin21/go-fts/constants"
	"github.com/adshin21/go-fts/utils"
)

func main() {

	var source, query string

	flag.StringVar(&source, "source", constants.DefaultSource, "Source path for xml zip")
	flag.StringVar(&query, "query", constants.DefaultQuery, "query")
	flag.Parse()

	start := time.Now()
	docs, err := utils.LoadDocuments(source)
	log.Printf("Total docs are %d, loaded in %v", len(docs), time.Since(start))

	if err != nil {
		log.Fatal(err)
	}

	start = time.Now()
	idx := make(utils.Index)

	// Add to cache once it is indexed
	idx.Add(docs)

	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	matchedIDs := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}

}
