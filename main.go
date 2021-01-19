package main

import (
	"fmt"
	"strconv"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"io"
)

func webserver(w http.ResponseWriter, r *http.Request) {
	// Generate unstructured JSON by GETing the api products/
	baseUrl := "https://reqres.in/api/products/"
	req, _ := http.NewRequest("GET", baseUrl, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

  // Get the total number of products
  // JSON data can be decoded to a map 
  var result map[string]interface{}
  json.Unmarshal([]byte(body), &result)
  var totalProducts int = int(result["total"].(float64))

  // Generate each product details and write it out
	for i := 1; i <= totalProducts; i++ {
	  url := baseUrl + strconv.Itoa(i)
    fmt.Println(url)
		req, _ := http.NewRequest("GET", url, nil)
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		output := "\n" + url + "\n" + string(body)
    io.WriteString(w, output)
  }
}

func main() {
	fmt.Println("Listening on port: 8000")

	//Web Server:
	http.HandleFunc("/", webserver)
	http.ListenAndServe(":8000", nil)
}

