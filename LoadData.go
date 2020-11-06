package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var es, _ = elasticsearch.NewDefaultClient()

func LoadData() {
	var spacecrafts []map[string]interface{}
	pageNumber := 0
	for {
		response, _ := http.Get("http://stapi.co/api/v1/rest/spacecraft/search?pageSize=100&pageNumber=" + strconv.Itoa(pageNumber))
		body, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		var result map[string]interface{}
		json.Unmarshal(body, &result)

		pageInterface := result["page"]
		page, _ := pageInterface.(map[string]interface{})
		totalPagesInterface := page["totalPages"]
		totalPages, _ := totalPagesInterface.(float64)

		spacecraftsInterface := result["spacecrafts"]
		crafts, _ := spacecraftsInterface.([]interface{})

		for _, craftInterface := range crafts {
			craft := craftInterface.(map[string]interface{})
			spacecrafts = append(spacecrafts, craft)
		}

		pageNumber++
		if pageNumber >= int(totalPages) {
			break
		}
	}

	for _, data := range spacecrafts {
		uid, _ := data["uid"].(string)
		jsonString, _ := json.Marshal(data)
		request := esapi.IndexRequest{Index: "stsc", DocumentID: uid, Body: strings.NewReader(string(jsonString))}
		request.Do(context.Background(), es)
	}
	print(len(spacecrafts), " spacecraft read\n")
}

func main() {
	fmt.Println(es.Info())
	LoadData()
}
