package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
)

var es, _ = elasticsearch.NewDefaultClient()

func Exit() {
	fmt.Println("Goodbye!")
	os.Exit(0)
}

func Get(reader *bufio.Scanner) {
	fmt.Print("Enter spacecraft ID: ")
	reader.Scan()
	id := reader.Text()
	request := esapi.GetRequest{Index: "stsc", DocumentID: id}
	response, _ := request.Do(context.Background(), es)
	var results map[string]interface{}
	json.NewDecoder(response.Body).Decode(&results)
	Print(results["_source"].(map[string]interface{}))
}

func Print(spacecraft map[string]interface{}) {
	name := spacecraft["name"]
	status := ""
	if spacecraft["status"] != nil {

		status = "- status: " + spacecraft["status"].(string)
	}
	fmt.Println(name, status)
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("0) Exit")
		fmt.Println("1) Load spacecraft")
		fmt.Println("2) Get spacecraft")
		fmt.Print("Enter option: ")
		reader.Scan()
		option := reader.Text()
		if option == "0" {
			Exit()
		} else if option == "1" {
			LoadData()
		} else if option == "2" {
			Get(reader)
		} else {
			fmt.Println("Invalid option")
		}
	}
}
