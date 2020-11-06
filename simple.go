package main

import (
    "log"
    "github.com/elastic/go-elasticsearch/v8"
)

func main() {
    es, _ := elasticsearch.NewDefaultClient()
    log.Println(elasticsearch.Version)
    log.Println(es.Info())
}
