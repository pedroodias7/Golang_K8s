package main

import (
	"encoding/json"
	"io"
	"os"
	"fmt"
	"log"
	"net/url"
	"net/http"
)


//json to Struct Golang
type Words struct{
	Page string `json:"page"`
	Input string `json:"input"`
	Words []string `json:"words"`
}

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url>\n")
		os.Exit(1)
	}

	_, err := url.ParseRequestURI(args[1]) //if _,err := url.ParseRequestURI(ags[1]); err != nill
	if err != nil {
		fmt.Printf("URL is in invalid format: %v\n", err)
	  	panic(err)	
	}
	
	
	res, err := http.Get(args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	

	if res.StatusCode != 200 {
		fmt.Printf("Invalid http code %d is in invalid format: %v\n", res.StatusCode, body)
		os.Exit(1)
	}


	var words Words


	err = json.Unmarshal(body.&words)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("JSON PARSE\nPage: %s\nWords: %v\n", words.Page, strings.Join(words.Words,","))



}
