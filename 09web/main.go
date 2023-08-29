package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Read Webpage Content
func makeWebReq(URL string) {
	fmt.Println("LCO Web Request ----------")

	res, err := http.Get(URL)

	if err != nil {
		panic(err)
	}
	// Imp: Close response @ end
	defer res.Body.Close()

	fmt.Printf("Type of res: %T \n", res)

	// Reading Content of Website
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// Change bytes data to readable string
	content := string(data)
	fmt.Println(content)
}

// Operations on URL
func handleURL(myurl string) {
  fmt.Println("------  URL Ops ----------")

  // parse
  res, _ := url.Parse(myurl)

  fmt.Println("URL = " ,res)
  fmt.Println("Scheme = " ,res.Scheme)
  fmt.Println("Host = " ,res.Host)
  fmt.Println("Path = " ,res.Path)
  fmt.Println("Port = " ,res.Port())
  fmt.Println("RawQuery = " , res.RawQuery)

  // ---------------------------------------------------

  fmt.Println("Reading Queries ----- ")

  // Another better method to read Params
  // Stores all query as key-value pairs
  qparams := res.Query()

  fmt.Println("Course: ", qparams["course"])

  // Read all in 
  for _, val := range qparams {
    fmt.Println("Param is : ", val)
  }

}

// Construct new URL using given params
func constructURL(param1 string) {

  // create queries
  queryValues := url.Values{}
  queryValues.Add("user", param1)

  // Construct parts of url
  partsOfUrl := &url.URL{
    Scheme: "https",
    Host: "lco.dev",
    Path: "/tutorials",
    RawQuery: queryValues.Encode(),
  }

  // construct string
  var newUrl string = partsOfUrl.String()
  fmt.Println("Newly Created URL: ", newUrl)
}

func main() {

  URL1 := "https://lco.dev"
  URL2 := "https://lco.dev:5050/learn?course=react&id=2311"


	makeWebReq(URL1)
  handleURL(URL2)
  constructURL("keshavjha")
}
