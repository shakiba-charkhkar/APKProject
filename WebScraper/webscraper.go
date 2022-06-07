package webscraper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	esapi "github.com/elastic/go-elasticsearch/v8/esapi"
)

type WebData struct {
	IP  string `json:"ip"`
	URL string `json:"url"`
}

func ReadWebData() {
	fmt.Println("Read web Data")
	var URL string = "https://www.dan.me.uk/torlist/?exit"
	response, err := http.Get(URL) //use package "net/http"

	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	// Copy data from the response to standard output
	// n, err1 := io.Copy(os.Stdout, response.Body) //use package "io" and "os"
	// if err != nil {
	// 	fmt.Println(err1)
	// 	return
	// }
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	result := strings.Split(bodyString, "\n")
	fmt.Println(bodyString)
	StorageData(result, URL)

}
func StorageData(dataArr []string, url string) {
	log.SetFlags(0)

	var (
		r  map[string]interface{}
		wg sync.WaitGroup
	)

	// Initialize a client with the default settings.
	//
	// An `ELASTICSEARCH_URL` environment variable will be used when exported.
	//
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 1. Get cluster info
	//
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))

	// 2. Index documents concurrently
	//
	for i, title := range dataArr {
		wg.Add(1)

		go func(i int, title string) {
			defer wg.Done()

			// Build the request body.
			webIp := &WebData{
				IP:  title,
				URL: url}
			data, err := json.Marshal(webIp)
			if err != nil {
				log.Fatalf("Error marshaling document: %s", err)
			}

			// Set up the request object.
			req := esapi.IndexRequest{
				Index: "test6",
				//DocumentID: strconv.Itoa(i + 1),
				Body:    bytes.NewReader(data),
				Refresh: "true",
			}

			// Perform the request with the client.
			res, err := req.Do(context.Background(), es)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
			} else {
				// Deserialize the response into a map.
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Printf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version.
					log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
				}
			}
		}(i, title)
	}
	wg.Wait()

	log.Println(strings.Repeat("=", 37))

}
