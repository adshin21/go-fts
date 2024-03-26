package utils

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

type document struct {
	Title string `xml:"title"`
	Url   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

// load the documents and returns
func LoadDocuments(path string) ([]document, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	gzip, err := gzip.NewReader(file)

	if err != nil {
		return nil, err
	}
	defer gzip.Close()

	var documents []document
	decoder := xml.NewDecoder(gzip)

	dump := struct {
		Documents []document `xml:"doc"`
	}{}

	err = decoder.Decode(&dump)

	if err != nil {
		return nil, err
	}

	documents = dump.Documents
	for i := range documents {
		documents[i].ID = i
	}

	return documents, nil
}
