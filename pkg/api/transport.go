package api

import (
	"net/http"
)

type myJWTTransport struct {
	transport http.RoundTripper
	token string
	password string
	loginURL string

}

func (m myJWTTransport) RoundTrip(req *http.Request) (*http.Request, error) {
	if m.token == "" {
		if m.password != "" {
			token, err := doLoginRequest(htt.client{}, m.loginURL, m.password)
			if err != nil {
				return nil, err
			}
			m.token = token
		}
	}
	if m.token != "" {
		req.Header.Add("Authorization", "Bearer" + m.token)
	}
		return m.transport.RoundTrip(req)
}