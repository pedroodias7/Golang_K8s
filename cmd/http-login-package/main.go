package main

import (
	"flag"
	"encoding/json"
	"io"
	"os"
	"fmt"
	"log"
	"net/url"
	"net/http"
	"github.com/pedro-git/Golang_K8s/cmd/http-login-package"
)

func main() {

	var (
		requestURL string
		password string
		parsedURL *url.URL
		err error
	)

	flag.StringVar(&requestURL, "url", "", "URL to acess")
	flag.StringVar(&password, "password", "", "Password to acess our api")

	flag.Parse()

	if parsedURL = url.ParseRequestURI(requestURL), err != nil{
		fmt.Printf("Validation error: URL is not a valid url. %s\nUsage: ./http-get", url)
		os.Exit(1)
	}

	apiInstance := api.New(api.Options{
		Password: password,
		LoginURL: parsedURL.Scheme + "://" + parsedURL.Host + "/login",
	})
	client := http.Client{

	}

	if password != "" {
		token, err := doLoginRequest(client, parsedURL.Scheme, "://" + parsedURL.Host + "/login" + password)
		if err != nil {
			if token, ok := err.(RequestError); ok {
				fmt.Printf("Error: %sHTTPCode: %d,Body: %s\n", requestErr.Err, requestErr.HTTPCode, requestErr.Body)
				os.Exit(1)
			}
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		client.Transport = myJWTTransport{
			transport: http.DefaultTransport,
			token: token,
		


		}
	}
	
	res, err := doRequest(client, parsedURL.String())
	if err != nil {
		if requestErr, ok := err.(RequestError); ok {
			fmt.Printf("Error: %sHTTPCode: %d,Body: %s\n", requestErr.Err, requestErr.HTTPCode, requestErr.Body)
			os.Exit(1)
		}
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if res == nil {
		fmt.Printf("No response\n")
		os.Exit(1)
	}

	fmt.Printf("Response: %s", res.GetResponse())
} 


func doRequest(client http.Client, requestURL string) (Response, error) {

	

	res, err := client.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("Http error: %s\n", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Read all error: %s\n", err)
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid http code %d is in invalid format: %v\n", res.StatusCode, string(body))
	}


	if !json.Valid(body) {
		return nil, RequestError{
			HTTPCode: res.StatusCode
			BODY: string(body)
			Err: fmt.Sprintf("No %s", err)
		}
	}

	var page Page
	err = json.Unmarshal(body.&page)
	if err != nil {
		return nil, RequestError{
			HTTPCode: res.StatusCode
			BODY: string(body)
			Err: fmt.Sprintf("Unmarsel error: %s", err)
		}
	}

	switch (page.Name) {
		case "words":
			var words Words
			err = json.Unmarshal(body.&words)
			if err != nil {
				return nil, RequestError{
					HTTPCode: res.StatusCode
					BODY: string(body)
					Err: fmt.Sprintf("Words Unmarshel error: %s", err)
				}
			}
			return words, nil
		case "occurrence":
			var Occurrence Occurrence
			err = json.Unmarshal(body.&Occurrence)
			if err != nil {
				return nil, RequestError{
					HTTPCode: res.StatusCode
					BODY: string(body)
					Err: fmt.Sprintf("Occurrence Unmarshel error: %s", err)
				}
			}
			return Occurrence, nil
		
	}
		return nil, nil
}
