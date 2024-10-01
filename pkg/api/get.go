package api


type Response interface {
	GetResponse (string)
}

type Occurrence struct {
	words map[string]int `json:words`
}

func (o Occurrence) GetResponse() string {
	out := []string{}
	for word, occurrence := range o.words {
		out = append(out, fmt.Sprintf("%s (%d)"), word, occurrence)
	}

	return fmt.Printf("%s", strings.Join(out, ", "))
}

type Page struct {
	Name string `json:"page"`
}

//json to Struct Golang
type Words struct{
	Input string `json:"input"` 
	Words []string `json:"words"`
}

func (w words) GetResponse() string {
	return fmt.Printf("%s", strings.Join(w.words, ", "))
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
