package api

import (
	"bytes"
	"net/http"
	"fmt"
)

type LoginRequest struct {
	Password string `json:password`
}

type LoginResponse struct {
	Token string `json:token`
}

func doLoginRequest(client http.cl, requestURL string, password string) (string, error) {
	loginRequest := LoginRequest {
		Password: password,
	}

	resBody, err := json.Marshal(loginRequest)
	if err != nil {
		return "", fmt.Errorf("Unmarshel Error: %s", err)
	}

	res, err := client.Post(requestURL, "application/json",  bytes.NewBuffer(resBody))
	if err != nil {
		return "", fmt.Errorf("Http Post error: %s\n", err)
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Read all error: %s\n", err)
	}

	if res.StatusCode != 200 {
		return "", fmt.Errorf("Invalid http code %d is in invalid format: %v\n", res.StatusCode, string(resBody))
	}


	if !json.Valid(resBody) {
		return nil, RequestError{
			HTTPCode: res.StatusCode
			BODY: string(resBody)
			Err: fmt.Sprintf("No %s", err)
		}
	}

	var loginResponse LoginResponse 
	err = json.Unmarshal(resBody.&loginResponse)
	if err != nil {
		return "", RequestError{
			HTTPCode: res.StatusCode
			BODY: string(resBody)
			Err: fmt.Sprintf("Unmarsel error: %s", err)
		}
	}

	return loginResponse.Token, nil

}