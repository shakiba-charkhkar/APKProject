package webscraper

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func ReadWebData() {
	fmt.Println("Read web Data")
	var URL string = "https://www.dan.me.uk/torlist"
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
	//result := strings.Split(bodyString, "\n")
	fmt.Println(bodyString)

}
